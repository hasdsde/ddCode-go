package utils

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	"gorm.io/datatypes"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func ExcelToGangDB[T any](ctx context.Context, db *gorm.DB, excelFilePath, tableName string, batchSize, cacheSize int) error {
	f, err := excelize.OpenFile(excelFilePath)
	if err != nil {
		return fmt.Errorf("failed to open excel file: %v", err)
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return fmt.Errorf("failed to get rows from sheet: %v", err)
	}
	var wg sync.WaitGroup
	pdChan := make(chan T, cacheSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		var batch []T
		batchSize := batchSize
		for pd := range pdChan {
			batch = append(batch, pd)
			if len(batch) >= batchSize {
				// 执行批量插入
				if err := db.Table(tableName).WithContext(ctx).Create(&batch).Error; err != nil {
					fmt.Printf("Error saving to database: %v\n", err)
					continue
				}
				batch = nil // 清空批量
			}
		}
		// 插入剩余的数据
		if len(batch) > 0 {
			if err := db.Table(tableName).Create(&batch).Error; err != nil {
				fmt.Printf("Error saving remaining data to database: %v\n", err)
			} else {
				fmt.Println("Final batch inserted successfully")
			}
		}
	}()

	// 跳过标题行
	var wgRows sync.WaitGroup
	for i, row := range rows[1:] {
		wgRows.Add(1)
		go func(i int, row []string) {
			defer wgRows.Done()
			var instance T
			val := reflect.ValueOf(&instance).Elem()
			for i := 0; i < val.NumField(); i++ {
				fieldVal := val.Field(i)

				// 根据字段类型处理数据
				switch fieldVal.Interface().(type) {
				case string:
					fieldVal.SetString(row[i])
				case int, int32, int64:
					intValue, err := strconv.ParseInt(row[i], 10, 64)
					if err != nil {
						fmt.Printf("EXCEL TO DB ERROR %s\n", err)
						continue
					}
					fieldVal.SetInt(intValue)
				case float32, float64:
				case time.Time:
					parsedTime, err := time.Parse(timeFormat, row[i])
					if err != nil {
						fmt.Printf("EXCEL TO DB TIME PARSE ERROR: %v\n", err)
						break
					}
					fieldVal.Set(reflect.ValueOf(parsedTime))

				case sql.NullTime:

				case datatypes.JSON:
					var details interface{}
					err := json.Unmarshal([]byte(row[i]), &details)
					if err != nil {
						fmt.Printf("EXCEL TO DB JSON UNMARSHAL Error: %v, 数据: %s\n\n", err, row[i])
						return
					}
					jsonBytes, err := json.Marshal(details)
					if err != nil {
						fmt.Println("EXCEL TO DB JSON MARSHAL Error:", err)
						return
					}
					fieldVal.Set(reflect.ValueOf(datatypes.JSON(jsonBytes)))
				}
			}
			pdChan <- instance
		}(i, row)
	}
	go func() {
		wgRows.Wait() // 等待所有行处理goroutine完成
		close(pdChan) // 所有goroutine完成后关闭通道
	}()
	wg.Wait()

	return nil
}

func CSVToGangDB[T any](ctx context.Context, db *gorm.DB, csvFilePath, tableName string, batchSize, cacheSize int) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	pdChan := make(chan T, cacheSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		var batch []T
		batchSize := batchSize
		for pd := range pdChan {
			batch = append(batch, pd)
			if len(batch) >= batchSize {
				if err := db.Table(tableName).Create(&batch).Error; err != nil {
					fmt.Printf("Error saving to database: %v\n", err)
					continue
				}
				batch = nil
			}
		}
		if len(batch) > 0 {
			if err := db.Table(tableName).Create(&batch).Error; err != nil {
				fmt.Printf("Error saving remaining data to database: %v\n", err)
			}
		}
	}()

	var wgRows sync.WaitGroup
	for i, row := range rows[1:] {
		wgRows.Add(1)
		go func(i int, row []string) {
			defer wgRows.Done()
			var instance T
			val := reflect.ValueOf(&instance).Elem()
			for i := 0; i < val.NumField(); i++ {
				fieldVal := val.Field(i)

				switch fieldVal.Interface().(type) {
				case string:
					fieldVal.SetString(row[i])
				case int, int32, int64:
					intValue, err := strconv.ParseInt(row[i], 10, 64)
					if err != nil {
						fmt.Printf("CSV TO DB ERROR %s\n", err)
						continue
					}
					fieldVal.SetInt(intValue)
				case float32, float64:
					floatValue, err := strconv.ParseFloat(row[i], 64)
					if err != nil {
						fmt.Printf("CSV TO DB ERROR %s\n", err)
						continue
					}
					fieldVal.SetFloat(floatValue)
				case time.Time:
					parsedTime, err := time.Parse(timeFormat, row[i])
					if err != nil {
						fmt.Printf("CSV TO DB TIME PARSE ERROR: %v\n", err)
						break
					}
					fieldVal.Set(reflect.ValueOf(parsedTime))
				case datatypes.JSON:
					var details interface{}
					err := json.Unmarshal([]byte(row[i]), &details)
					if err != nil {
						fmt.Printf("CSV TO DB JSON UNMARSHAL Error: %v, 数据: %s\n", err, row[i])
						return
					}
					jsonBytes, err := json.Marshal(details)
					if err != nil {
						fmt.Println("CSV TO DB JSON MARSHAL Error:", err)
						return
					}
					fieldVal.Set(reflect.ValueOf(datatypes.JSON(jsonBytes)))
				}
			}
			pdChan <- instance
		}(i, row)
	}

	go func() {
		wgRows.Wait() // 等待所有行处理goroutine完成
		close(pdChan) // 所有goroutine完成后关闭通道
	}()
	wg.Wait()

	return nil
}
