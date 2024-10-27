package oss

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/nfnt/resize"
	"github.com/stretchr/testify/assert"
)

const (
	accessKey = "Zb87HDzeKRR0BqZM0D8h"
	secretKey = "BP1DKwgF8me3icnRVPEt1GDmPCGRWriT7bbtShzL"
	bucket    = "xxx"
	addr      = "http://124.70.56.154:9000"
)

func TestOSS(t *testing.T) {
	g, err := NewOSS(addr, accessKey, secretKey, bucket)

	assert.NoError(t, err)
	// 写入
	local, err := g.PutS3Object(context.Background(), "xxx", accessKey, "application/octet-stream", []byte("hello world"))
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s/%s/%s", addr, "xxx", accessKey), local)

	// 读取
	data, err := g.GetS3Object(context.Background(), "xxx", accessKey)
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), data)

	buf := aws.NewWriteAtBuffer([]byte{})
	err = g.GetS3ObjectWithWriter(context.Background(), "xxx", accessKey, buf)
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), buf.Bytes())
}

func TestImages(t *testing.T) {
	g, err := NewOSS(addr, accessKey, secretKey, bucket)
	assert.NoError(t, err)

	imageBytes, err := readImageFile("./mayi.png")
	assert.NoError(t, err)
	originalSize := len(imageBytes)

	compressedImageBytes, err := compressImage(imageBytes)
	assert.NoError(t, err)

	local, err := uploadCompressedImage(g, compressedImageBytes)
	assert.NoError(t, err)
	expectedURL := fmt.Sprintf("%s/%s/%s", addr, "xxx", "mayi.jpg")
	assert.Equal(t, expectedURL, local)

	downloadedImageBytes, err := downloadCompressedImage(g)
	assert.NoError(t, err)

	assertCompressionRatio(t, originalSize, downloadedImageBytes)

	err = saveDownloadedImage(downloadedImageBytes)
	assert.NoError(t, err)
}

func readImageFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func compressImage(imageBytes []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(uint(img.Bounds().Dx()/2), 0, img, resize.Lanczos3)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func uploadCompressedImage(g *OSS, compressedImageBytes []byte) (string, error) {
	return g.PutS3Object(context.Background(), "xxx", "mayi.jpg", "application/octet-stream", compressedImageBytes)
}

func downloadCompressedImage(g *OSS) ([]byte, error) {
	return g.GetS3Object(context.Background(), "xxx", "mayi.jpg")
}

func assertCompressionRatio(t *testing.T, originalSize int, downloadedImageBytes []byte) {
	storedSize := len(downloadedImageBytes)
	compressionRatio := calculateCompressionRatio(originalSize, storedSize)
	fmt.Println("Compression Ratio:", compressionRatio)
}

func calculateCompressionRatio(originalSize, storedSize int) float64 {
	if originalSize == 0 {
		return 0
	}
	return float64(storedSize) / float64(originalSize)
}

func saveDownloadedImage(downloadedImageBytes []byte) error {
	return ioutil.WriteFile("downloaded_compressed_image.jpg", downloadedImageBytes, 0644)
}

func CalculateCompressionRatio(originalSize, storedSize int) float64 {
	if originalSize == 0 {
		return 0
	}
	return float64(storedSize) / float64(originalSize)
}

func TestUploadImg(t *testing.T) {
	ctx := context.Background()
	file, err := os.Open("./mayi.png")
	if err != nil {
		t.Errorf("打开失败:%s", err)
	}
	readAll, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	key, err := UploadImg(readAll, "mayi.png", accessKey, secretKey, addr, Bucket, ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(key)
}

func TestDownloadImg(t *testing.T) {
	ctx := context.Background()
	key := "mayi.png"
	img, err := DownloadImg(key, accessKey, secretKey, addr, bucket, ctx)
	if err != nil {
		t.Errorf("下载失败:%s", err)
	}
	file, err := os.Create("test.png")
	if err != nil {
		t.Errorf("创建失败:%s", err)
	}
	// 去除前缀
	img = strings.Split(img, ",")[1]
	decodeString, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		t.Errorf("解码失败:%s", err)
	}
	file.Write(decodeString)
	defer file.Close()
}
