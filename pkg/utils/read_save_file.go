package utils

import (
	"bufio"
	"ddCode-server/global"
	"fmt"
	"os"
	"regexp"

	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func ReadSaveFile(file *multipart.FileHeader, ctx *gin.Context) ([]string, error) {
	str := []string{""}
	dst := "./static/uploads/" + file.Filename
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		global.Logger.Errorf(fmt.Sprintf("文件保存失败: %s", err.Error()))
		return str, err
	}
	// 读取文件内容
	// 将文件内容读取为字符串
	contentBytes, err := os.ReadFile(dst)
	if err != nil {
		global.Logger.Errorf(fmt.Sprintf("无法读取文件内容: %s", err.Error()))
		return str, err
	}
	fileContentStr := string(contentBytes)
	// 根据文件类型选择相应的处理方式
	fileExtension := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])
	switch fileExtension {
	case "xlsx", "xls":
		str, err = extractIPFromExcel(fileContentStr)
	case "txt", "csv", "json", "xml":
		str, err = extractIPFromText(fileContentStr)
	default:
		err = fmt.Errorf("不支持的文件格式: %s", fileExtension)
		return str, err
	}
	if err != nil {
		global.Logger.Errorf(fmt.Sprintf("文件处理失败: %s", err.Error()))
		return str, err
	}
	return str, nil
}
func extractIPFromExcel(fileContentStr string) ([]string, error) {
	excelFile, err := xlsx.OpenBinary([]byte(fileContentStr))
	if err != nil {
		return nil, fmt.Errorf("无法打开 Excel 文件: %s", err.Error())
	}

	var ipAddresses []string

	for _, sheet := range excelFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				ip := cell.String() // 假设 IP 地址或 IP 段所在的单元格类型为字符串
				if !isValidIPFormat(ip) {
					continue
				}
				ipAddresses = append(ipAddresses, ip)
			}
		}
	}
	return ipAddresses, nil
}

func extractIPFromText(fileContentStr string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(fileContentStr))
	var ipAddresses []string
	for scanner.Scan() {
		line := scanner.Text()
		// 这里可以根据文件的格式进行相应的处理，比如使用正则表达式提取 IP 地址或 IP 段
		// 将提取到的 IP 地址或 IP 段进行处理，比如存储到数组或进行其他操作
		if !isValidIPFormat(line) {
			continue
		}
		ipAddresses = append(ipAddresses, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件出错: %s", err.Error())
	}
	return ipAddresses, nil
}

// isValidIPFormat 检测给定的字符串是否为有效的IPv4、CIDR、IPv4+端口或IPv4范围格式。
func isValidIPFormat(ip string) bool {
	// IPv4 地址的正则表达式。
	ipv4Regex := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	// CIDR 表达式的正则表达式。
	cidrRegex := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(?:\/(?:3[0-2]|[12]?[0-9]))$`)
	// IPv4+端口的正则表达式。
	ipPortRegex := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):(?:6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$`)
	// IPv4范围的正则表达式。
	ipRangeRegex := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)-(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	switch {
	case ipv4Regex.MatchString(ip):
		return true
	case cidrRegex.MatchString(ip):
		return true
	case ipPortRegex.MatchString(ip):
		return true
	case ipRangeRegex.MatchString(ip):
		return true
	default:
		return false
	}
}
