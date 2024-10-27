package oss

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	Bucket            = "xxx"
	Base64ImagePrefix = "data:image/png;base64,"
)

type Config struct {
	AccessKey string
	SecretKey string
	Regions   string
	Addr      string
}

type OptionFunc func(*OSS)

func WithLog(log *log.Logger) OptionFunc {
	return func(g *OSS) {
		g.log = log
	}
}

type OSS struct {
	svc *session.Session // 亚马逊s3库操作器
	log *log.Logger
}

// NewOSS
// addr 确保地址包含HTTP前缀
func NewOSS(addr, access, secret, regions string, ops ...OptionFunc) (*OSS, error) {
	g := &OSS{
		log: log.Default(),
	}
	cres := credentials.NewStaticCredentials(access, secret, "")
	cfg := aws.NewConfig().WithRegion(regions).WithEndpoint(addr).WithCredentials(cres).WithS3ForcePathStyle(true)
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}
	g.svc = sess
	for _, op := range ops {
		op(g)
	}
	return g, nil
}

func (g *OSS) PutS3ObjectWithReader(ctx context.Context, bucket, key, contentType string, data io.Reader) (string, error) {
	uploader := s3manager.NewUploader(g.svc)
	result, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        data,
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

func (g *OSS) PutS3Object(ctx context.Context, bucket, fileName, contentType string, data []byte) (string, error) {
	uploader := s3manager.NewUploader(g.svc)
	result, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		ContentType: aws.String(contentType),
		Body:        bytes.NewReader(data),
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
func (g *OSS) GetS3Object(ctx context.Context, bucket, key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(g.svc)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.DownloadWithContext(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return buf.Bytes(), err
}

func (g *OSS) GetS3ObjectWithWriter(ctx context.Context, bucket, key string, w io.WriterAt) error {
	downloader := s3manager.NewDownloader(g.svc)
	_, err := downloader.DownloadWithContext(ctx, w, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

// UploadImg 上传图片
func UploadImg(file []byte, fileName, accessKey, secretKey, addr, bucket string, ctx context.Context) (string, error) {
	g, err := NewOSS(addr, accessKey, secretKey, bucket)
	if err != nil {
		return "", fmt.Errorf("创建OSS client失败:%s", err.Error())
	}
	// 这个不能解码ico
	// 	图片不进行压缩，实在是太糊了
	// mg, _, err := image.Decode(bytes.NewReader(file))
	// f err != nil {
	// 	return "", err
	//
	// ar compressedImageBytes bytes.Buffer
	// rr = jpeg.Encode(&compressedImageBytes, img, nil)
	// f err != nil {
	// 	return "", err
	//
	key, err := g.PutS3Object(ctx, bucket, fileName, "application/octet-stream", file)
	if err != nil {
		return "", err
	}
	return key, nil
}

// DownloadImg 下载图片并转换为base64,记得去掉前缀
func DownloadImg(fileName, accessKey, secretKey, addr, bucket string, ctx context.Context) (string, error) {
	g, err := NewOSS(addr, accessKey, secretKey, bucket)
	if err != nil {
		return "", err
	}
	file, err := g.GetS3Object(ctx, bucket, fileName)
	if err != nil {
		return "", err
	}
	return Base64ImagePrefix + base64.StdEncoding.EncodeToString(file), nil
}
