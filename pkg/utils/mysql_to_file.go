package utils

import (
	"bytes"
	"ddCode-server/global"
	"ddCode-server/internal"
	"ddCode-server/pkg/storage/oss"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
)

const ExportTag = "export"

func ExportToCSV(data interface{}, filePath string) error {
	file, err := OpenFileForWrite(filePath) // 打开文件
	if err != nil {
		panic(err)
	}
	defer func() {
		f := file.(*os.File)
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()
	_, bomErr := file.Write([]byte{'\xEf', '\xBB', '\xBF'})
	if bomErr != nil {
		return bomErr
	}
	writer := csv.NewWriter(file) // 创建CSV写入器
	defer writer.Flush()
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("expected a slice, got %s", v.Kind()) // 检查数据是否为切片类型
	}
	var header []string
	if v.Len() > 0 {
		firstElem := v.Index(0)
		t := firstElem.Type()
		for i := 0; i < firstElem.NumField(); i++ {
			field := t.Field(i)
			columnName := field.Tag.Get(ExportTag)
			if columnName == "" {
				columnName = field.Name
			}
			header = append(header, columnName)
		}
	}
	if err := writer.Write(header); err != nil {
		return err // 写入表头
	}
	for i := 0; i < v.Len(); i++ {
		var record []string
		elem := v.Index(i)
		for j := 0; j < elem.NumField(); j++ {
			valueField := elem.Field(j)
			var valueStr string
			if valueField.Type() == reflect.TypeOf(internal.LocalTime{}) {
				valueStr = valueField.Interface().(internal.LocalTime).StrValue()
			} else {
				valueStr = fmt.Sprintf("%v", valueField.Interface())
			}
			record = append(record, valueStr) // 将每个字段的值转换为字符串
		}
		if err := writer.Write(record); err != nil {
			return err // 写入数据行
		}
	}
	return nil
}
func ExportToXLSX(data interface{}, filePath string) error {
	file, err := OpenFileForWrite(filePath) // 打开文件
	if err != nil {
		panic(err)
	}
	defer func() {
		f := file.(*os.File)
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()
	f := excelize.NewFile() // 创建新的XLSX文件
	sheetName := Sheet1
	index, _ := f.NewSheet(sheetName) // 创建新的工作表
	f.SetActiveSheet(index)           // 设置活动工作表

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("expected a slice, got %s", v.Kind()) // 检查数据是否为切片类型
	}
	if v.Len() > 0 {
		firstElem := v.Index(0)
		t := firstElem.Type()
		for i := 0; i < firstElem.NumField(); i++ {
			field := t.Field(i)
			columnName := field.Tag.Get(ExportTag)
			if columnName == "" {
				columnName = field.Name
			}
			cell, _ := excelize.CoordinatesToCellName(i+1, 1)
			err = f.SetCellValue(sheetName, cell, columnName) // 写入表头
			if err != nil {
				fmt.Println(err)
			}
		}

		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			for j := 0; j < elem.NumField(); j++ {
				cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
				valueField := elem.Field(j)
				var value interface{}
				if valueField.Type() == reflect.TypeOf(internal.LocalTime{}) {
					value = valueField.Interface().(internal.LocalTime).StrValue()
				} else {
					value = valueField.Interface()
				}

				// TODO: 这里可能影响到的全局
				strValue := fmt.Sprintf("%s", value)
				if strings.HasPrefix(strValue, oss.Base64ImagePrefix) {
					imgStr := strings.Split(strValue, ",")[1]
					reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgStr))
					img, err := jpeg.Decode(reader) //nolint:all
					if err != nil {
						// TODO: 图片为icon会出错
						fmt.Println(err)
						continue
					}
					buf := new(bytes.Buffer)
					err = jpeg.Encode(buf, img, nil)
					if err != nil {
						return err
					}
					err = f.AddPictureFromBytes(sheetName, cell, &excelize.Picture{
						Extension: ".png",
						File:      buf.Bytes(),
						Format: &excelize.GraphicOptions{
							ScaleX:          0.1,
							ScaleY:          0.1,
							LockAspectRatio: true,
						},
					})
					if err != nil {
						return err
					}
				} else {
					err = f.SetCellValue(sheetName, cell, value) // 写入数据
				}
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	// 保存XLSX文件
	if err := f.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}

// ExportMapToXLSX 将 []map[string]interface{} 导出为 xlsx 文件
func ExportMapToXLSX(data []map[string]interface{}, filePath string, fields []string, bppMap map[string]string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return err
	}
	_, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	file := excelize.NewFile() // 创建新的 XLSX 文件
	sheetName := Sheet1
	index, _ := file.NewSheet(sheetName) // 创建新的工作表
	file.SetActiveSheet(index)           // 设置活动工作表

	if len(data) == 0 {
		return fmt.Errorf("no data to export")
	}

	// 写入表头
	headerRow := 1
	headers := make([]string, 0) // 转换为有序数组
	for columnIndex, fieldName := range fields {
		cellName, _ := excelize.CoordinatesToCellName(columnIndex+1, headerRow)
		err = file.SetCellValue(sheetName, cellName, bppMap[strings.ReplaceAll(fieldName, "`", "")])
		headers = append(headers, strings.ReplaceAll(fieldName, "`", ""))
		if err != nil {
			global.Logger.Infof("写入表头错误:%v", err)
			return err
		}
	}

	// 写入数据
	for rowIndex, item := range data {
		for columnIndex, fieldName := range headers {
			cellName, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
			err = file.SetCellValue(sheetName, cellName, item[fieldName])
			if err != nil {
				global.Logger.Infof("写入数据错误:%v", err)
				return err
			}
		}
	}

	// 保存 XLSX 文件
	if err := file.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}

// ExportMapToCSV 将 []map[string]interface{} 导出为 CSV 文件
func ExportMapToCSV(data []map[string]interface{}, filePath string, fields []string, bppMap map[string]string) error {
	// 创建 CSV 文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入表头
	_, bomErr := file.Write([]byte{'\xEf', '\xBB', '\xBF'})
	if bomErr != nil {
		return bomErr
	}
	// 创建 CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(data) == 0 {
		return fmt.Errorf("no data to export")
	}

	// 中文表头
	header := make([]string, 0)
	for _, field := range fields {
		header = append(header, bppMap[strings.ReplaceAll(field, "`", "")])
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	// 写入数据
	for _, item := range data {
		row := make([]string, 0)
		for _, field := range fields {
			row = append(row, fmt.Sprintf("%v", item[strings.ReplaceAll(field, "`", "")]))
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
