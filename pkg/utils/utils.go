package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

func DeepCopy(src, dst interface{}) error {
	ds, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(ds, dst)
}

// MD5
func GenerateMD5(s []byte) [16]byte {
	return md5.Sum(s)
}
func GenerateMD5ToHex(s []byte) string {
	return fmt.Sprintf("%x", GenerateMD5(s))
}

// base64加密
func Base64Encode(src []byte) []byte {
	coder := base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}

// base64解密
func Base64Decode(src []byte) ([]byte, error) {
	coder := base64.NewEncoding(base64Table)
	return coder.DecodeString(string(src))
}

// SplitArray 将arr均分成n组
func SplitArray[T any](arr []T, num int) [][]T {
	if num <= 1 {
		return [][]T{arr}
	}
	max := len(arr)
	if max < num {
		segmens := make([][]T, 0, max)
		for _, item := range arr {
			segmens = append(segmens, []T{item})
		}
		return segmens
	}
	segmens := make([][]T, 0)
	quantity := max / num
	end := 0
	endGroup := []T{}
	for i := 1; i <= num; i++ {
		qu := i * quantity
		items := make([]T, quantity)
		_ = copy(items, arr[i-1+end:qu])
		segmens = append(segmens, items)
		if i == num {
			endGroup = arr[qu:]
		}
		end = qu - i
	}
	for i, item := range endGroup {
		segmens[i] = append(segmens[i], item)
	}
	return segmens
}

// SequentialSplitArray 按序将 arr 拆分成 num 组
func SequentialSplitArray[T any](arr []T, num int) [][]T {
	segmens := make([][]T, num)
	for i, item := range arr {
		index := i % num
		if segmens[index] == nil {
			segmens[index] = make([]T, 0)
		}
		segmens[index] = append(segmens[index], item)
	}
	return segmens
}

func GetDateTime(typeTime int) string {
	var str string
	switch typeTime {
	case 1:
		str = time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
	case 2:
		str = time.Now().AddDate(0, -3, 0).Format("2006-01-02 15:04:05")
	case 3:
		str = time.Now().AddDate(0, -6, 0).Format("2006-01-02 15:04:05")
	default:
		str = time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")
	}
	return str
}
