package utils

import "hash/fnv"

// Fnv32Hash 字符串转换 hash 32 位 fnv 算法
func Fnv32Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Fnv64Hash 字符串转换 hash 34 位 fnv 算法
func Fnv64Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
