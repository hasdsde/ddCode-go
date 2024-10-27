package utils

import (
	"archive/zip"
	"bufio"
	"io"
	"os"
	"path"
	"path/filepath"
)

// MakeFileByLineStr 组织按行写入一个文件
func MakeFileByLineStr(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
	return nil
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// Unzip 解压缩 zip
func Unzip(zipPath, dstDir string) ([]string, error) {
	// open zip file
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	fileNames := make([]string, len(reader.File))
	for i, file := range reader.File {
		fileNames[i] = file.Name
		if err := unzipFile(file, dstDir); err != nil {
			return nil, err
		}
	}
	return fileNames, nil
}

func unzipFile(file *zip.File, dstDir string) error {
	// create the directory of file
	// nolint
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	// open the file
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer func(rc io.ReadCloser) {
		err = rc.Close()
		if err != nil {
			panic(err)
		}
	}(rc)
	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(w *os.File) {
		err = w.Close()
		if err != nil {
			panic(err)
		}
	}(w)
	// save the decompressed file content
	// nolint
	_, err = io.Copy(w, rc)
	return err
}

// OpenFileForWrite 打开（或创建）一个文件准备写入，返回值符合 io.Writer 接口
func OpenFileForWrite(path string) (io.Writer, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		// 如果打开文件出错，返回错误
		return nil, err
	}
	return file, nil
}
func GetFileABS(prefix, fileName string) string {
	if filepath.IsAbs(fileName) {
		return fileName
	}
	abs, err := filepath.Abs(filepath.Join(prefix, fileName))
	if err != nil {
		panic(err)
	}
	return abs
}
