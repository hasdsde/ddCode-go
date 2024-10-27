package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	cacheDir = "./tmp"
)

func TestCreateCacheFile(t *testing.T) {
	targets := NewStringSet(
		"10.11.12.35:80",
		"10.10.11.218:32000",
		"10.10.11.177:80",
		"a.com:8080",
	)

	fileName := fmt.Sprintf("%s/%s-%s.txt", cacheDir, fmt.Sprint(1001), time.Now().Format(StrTimeFormat))
	assert.NoError(t, MakeFileByLineStr(
		fileName,
		targets.List(),
	))
	defer os.Remove(fileName)

	// 验证存在
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Error(err)
		return
	}
	assert.NoError(t, err)
	// 验证内容
	rf, err := os.Open(fileName)
	if err != nil {
		assert.NoError(t, err)
	}
	defer rf.Close()
	fd, err := ioutil.ReadAll(rf)
	assert.NoError(t, err)
	for _, ta := range strings.Split(string(fd), "\n") {
		if ta != "" {
			assert.True(t, targets.Has(ta), ta)
		}
	}
}

func TestGetExt(t *testing.T) {
	type Case struct {
		name     string
		filename string
		ext      string
	}
	cases := []Case{
		{
			name:     "usually_txt",
			filename: "usually.txt",
			ext:      ".txt",
		},
		{
			name:     "path_txt",
			filename: "this/is/a/path/usually.txt",
			ext:      ".txt",
		},
	}
	for _, cc := range cases {
		ext := GetExt(cc.filename)
		assert.Equal(t, cc.ext, ext)
	}
}
