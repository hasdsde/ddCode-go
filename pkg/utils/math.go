package utils

const (
	MININT64 = -922337203685477580
	MAXINT64 = 9223372036854775807
)

func Max(nums ...int64) int64 {
	var maxNum int64 = MININT64
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func Min(nums ...int64) int64 {
	var minNum int64 = MAXINT64
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func Sum(nums ...int64) int64 {
	var sumNum int64 = 0
	for _, num := range nums {
		sumNum += num
	}
	return sumNum
}
