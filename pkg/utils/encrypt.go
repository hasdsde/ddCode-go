// nolint
package utils

import (
	"crypto/sha1" // nosec
	"fmt"
)

func GetSha1(password string) string {
	h := sha1.New() // md5加密类似md5.New()
	// 写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(password))
	// 这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	// SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	str := fmt.Sprintf("%x", bs)
	return str
}
