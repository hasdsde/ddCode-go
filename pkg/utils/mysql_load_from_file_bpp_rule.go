package utils

import (
	"context"
	"ddCode-server/global"
	"encoding/csv"
	"encoding/json"
	"errors"
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

const timeFormat = "2006-01-02 15:04:05"

func ExcelToDB[T any](ctx context.Context, db *gorm.DB, excelFilePath, tableName string, batchSize, cacheSize int) error {
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
	var errMsg string
	pdChan := make(chan T, cacheSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var batch []T
		batchSize := batchSize
		for pd := range pdChan {
			batch = append(batch, pd)
			if errMsg != "" {
				batch = nil
			}
			if len(batch) >= batchSize {
				// 执行批量插入
				if error := db.Table(tableName).WithContext(ctx).Create(&batch).Error; error != nil {
					fmt.Printf("Error saving to database: %v\n", error)
					continue
				}
				batch = nil // 清空批量
			}
		}
		// 插入剩余的数据
		if len(batch) > 0 {
			if err1 := db.Table(tableName).Create(&batch).Error; err1 != nil {
				fmt.Printf("Error saving remaining data to database: %v\n", err1)
			} else {
				fmt.Println("Final batch inserted successfully")
			}
		}
	}()

	// 跳过标题行
	var wgRows sync.WaitGroup
	header := rows[0]
	expectedHeader := []string{"事件ID", "事件名称", "黑产类型", "操作人角色", "操作人行为", "事件定义", "事件动作", "事件模板"}
	if len(header) != len(expectedHeader) {
		fmt.Println("The number of columns in the header is not as expected")
		return fmt.Errorf("The number of columns in the header is not as expected")
	}
	for i := range header {
		if header[i] != expectedHeader[i] {
			fmt.Println("The header does not match expectations")
			return fmt.Errorf("The header does not match expectations")
		}
	}

	for j, row := range rows[1:] {
		var intValue int64
		wgRows.Add(1)
		go func(j int, row []string) {
			defer wgRows.Done()
			var instance T
			val := reflect.ValueOf(&instance).Elem()
			for i := 0; i <= 7; i++ {
				fieldVal := val.Field(i)
				var tempValue interface{}
				// 根据字段类型处理数据
				switch fieldVal.Interface().(type) {
				case string:
					tempValue = row[i]
				case int, int32, int64:
					intValue, err = strconv.ParseInt(row[i], 10, 64)
					if err != nil {
						fmt.Printf("EXCEL TO DB ERROR %s\n", err)
					}
					tempValue = intValue
				case time.Time:
					parsedTime, err := time.Parse(timeFormat, row[i])
					if err != nil {
						fmt.Printf("EXCEL TO DB TIME PARSE ERROR: %v\n", err)
						break
					}
					tempValue = parsedTime
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
					tempValue = datatypes.JSON(jsonBytes)
				}

				if i == 0 && intValue == 0 {
					errMsg = fmt.Sprintf("Line %d Type 3 does not match", i)
				}
				if i == 2 {
					switch tempValue {
					case "虚假贷款诈骗", "虚假投资诈骗", "虚假返利诈骗", "虚假博彩诈骗", "虚假货币传销诈骗", "仿冒类诈骗",
						"公检法类诈骗", "刷单诈骗", "裸聊诈骗", "资格证办理诈骗", "银行账户钓鱼诈骗", "邮箱账户钓鱼诈骗", "纳税信息诈骗":
						// 符合条件，不进行编辑
					default:
						errMsg = fmt.Sprintf("Line %d Type 3 does not match", i)
					}
				}
				if i == 3 {
					switch tempValue {
					case "管理员操作", "用户操作", "代理操作":
						// 符合条件，不进行编辑
					default:
						errMsg = fmt.Sprintf("Line %d Type 4 does not match", i)
						global.Logger.Errorf("service excel to db err: %s", errMsg)
					}
				}
				if i == 4 {
					switch tempValue {
					case "管理员登录", "用户登录", "代理登录":
						// 符合条件，不进行编辑
					default:
						errMsg = fmt.Sprintf("Line %d Type 5 does not match", i)
						global.Logger.Errorf("service excel to db err: %s", errMsg)
					}
				}
				fieldVal.Set(reflect.ValueOf(tempValue))
			}
			pdChan <- instance
		}(j, row)
	}

	go func() {
		wgRows.Wait() // 等待所有行处理goroutine完成
		close(pdChan) // 所有goroutine完成后关闭通道
	}()
	wg.Wait()
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func CSVToDB[T any](ctx context.Context, db *gorm.DB, csvFilePath, tableName string, batchSize, cacheSize int) error {
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
