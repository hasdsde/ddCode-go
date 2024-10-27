package utils

import (
	"encoding/base64"
	"io"
)

type IfTrue interface {
	~bool | []string | any
}

func If[T IfTrue](b bool, trueVal, falseVal T) T {
	if b {
		return trueVal
	}
	return falseVal
}

func SliceSepByStep[T interface{}](originalSlice []T, step int) [][]T {
	length := len(originalSlice) / step
	if len(originalSlice)%step != 0 {
		length++
	}

	slicedSlice := make([][]T, length)
	for i := 0; i < length; i++ {
		start := i * step
		end := start + step
		if end > len(originalSlice) {
			end = len(originalSlice)
		}
		slicedSlice[i] = originalSlice[start:end]
	}
	return slicedSlice
}
func ReadToBase64(reader io.Reader) (string, error) {
	data, err := ReadToBytes(reader)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func ReadToBytes(reader io.Reader) (data []byte, err error) {
	buf := make([]byte, 1024)
	data = make([]byte, 0, 4096)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	return
}
