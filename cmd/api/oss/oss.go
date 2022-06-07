package oss

import (
	"bytes"
	"douyin/v1/pkg/constants"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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

func PutImage(sourceFile string, targetFile string) error {
	return ossBucket.PutObjectFromFile(targetFile, sourceFile)
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println("wrong") //这里换成我们的err
		return err
	}
	err = imaging.Save(img, "./out.jpeg") //这里换成我们的输出地址
	if err != nil {
		fmt.Println("wrong") //这里换成我们的err
		return err
	}
	return nil
}
