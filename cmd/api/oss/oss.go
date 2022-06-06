package oss

import (
	"douyin/v1/pkg/constants"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"mime/multipart"
	"os"
)

var ossBucket *oss.Bucket

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func Init() {
	client, err := oss.New(constants.ENDPOINT, constants.ACCESSId, constants.AccessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	ossBucket, err = client.Bucket(constants.BucketName)
	if err != nil {
		handleError(err)
	}
}

func PutFeed(fd *multipart.FileHeader, objectName string) error {
	fh, err := fd.Open()
	if err != nil {
		return err
	}
	feed := io.Reader(fh)
	return ossBucket.PutObject(objectName, feed)
}
