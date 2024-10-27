package utils

import (
	"os"
	"testing"
)

func TestReadCsvByBytes(t *testing.T) {
	f, err := os.Open("/Users/sunlin/Downloads/20230717.csv")
	if err != nil {
		t.Fatal(err)
	}
	content, err := ReadCsv(f, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)

	bytes, err := os.ReadFile("/Users/sunlin/Downloads/20230717.csv")
	if err != nil {
		t.Fatal(err)
	}
	content1, err := ReadCsvByBytes(bytes, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content1)
}
