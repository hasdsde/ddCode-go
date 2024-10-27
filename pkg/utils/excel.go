package utils

import (
	"bytes"
	"errors"

	"github.com/duke-git/lancet/fileutil"
	"github.com/xuri/excelize/v2"
)

const Sheet1 = "Sheet1"

// ReadExcelFile 读取excel
// filePath  文件相对路径
// sheetName 标签页名称
// offsetNum 跳过行数 -1 为不跳过行数
func ReadExcelFile(filePath, sheetName string, offsetRowNum int) (content [][]string, err error) {
	if offsetRowNum < 0 {
		err = errors.New("offsetRowNum 不能小于0")
		return
	}
	isFileExist := fileutil.IsExist(filePath)
	if !isFileExist {
		err = errors.New("文件不存在")
		return
	}
	if sheetName == "" {
		sheetName = Sheet1
	}
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	rows, err := f.GetRows(sheetName)

	for key, row := range rows {
		// 跳过行数
		if offsetRowNum > 0 && key+1 <= offsetRowNum {
			continue
		}
		content = append(content, row)
	}
	return
}

// ReadExcelBytes 通过 bytes 数组获取Excel内容
func ReadExcelBytes(body []byte, sheetName string, offsetRowNum int, sheetIndex ...int) (content [][]string, err error) {
	if offsetRowNum < 0 {
		err = errors.New("offsetRowNum 不能小于0")
		return
	}
	if len(body) == 0 {
		return
	}
	if sheetName == "" {
		sheetName = Sheet1
	}
	reader := bytes.NewReader(body)
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return [][]string{}, err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			return
		}
	}()

	var sheetNames []string
	// 指定 sheet index
	if len(sheetIndex) > 0 {
		list := f.GetSheetList()
		for _, idx := range sheetIndex {
			for i, sheet := range list {
				if i == idx {
					sheetNames = append(sheetNames, sheet)
					break
				}
			}
		}
	} else {
		sheetNames = append(sheetNames, sheetName)
	}

	for _, sheet := range sheetNames {
		rows, err := f.GetRows(sheet)
		if err != nil {
			return [][]string{}, err
		}
		for key, row := range rows {
			// 跳过行数
			if offsetRowNum > 0 && key+1 <= offsetRowNum {
				continue
			}
			content = append(content, row)
		}
	}

	return
}
