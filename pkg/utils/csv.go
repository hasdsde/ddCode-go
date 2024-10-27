package utils

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// ReadCsv 读取 csv 文件，
// file 文件句柄
// allowEmp 是否允许为空
func ReadCsv(file *os.File, allowEmp bool) ([][]string, error) {
	w := csv.NewReader(file)
	arr := make([][]string, 0)
	for {
		data, err := w.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return [][]string{}, err
		}
		if !allowEmp && len(data) == 0 {
			continue
		}
		arr = append(arr, data)
	}
	return arr, nil
}

// ReadCsvByBytes 读取 csv 文件，
// file 文件句柄
// allowEmp 是否允许为空
func ReadCsvByBytes(file []byte, allowEmp bool) ([][]string, error) {
	w := csv.NewReader(bytes.NewReader(file))
	arr := make([][]string, 0)
	for {
		data, err := w.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return [][]string{}, err
		}
		if !allowEmp && len(data) == 0 {
			continue
		}
		arr = append(arr, data)
	}
	return arr, nil
}
