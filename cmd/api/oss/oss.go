package oss

import (
	"bytes"
	"douyin/v1/pkg/constants"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"sync"
	"time"
)

type VideoOss struct {
	SourceFile string
	TargetFile string
}

type ImgOss struct {
	VideoFile  string
	SourceFile string
	TargetFile string
}

var VideoList = struct {
	sync.RWMutex
	M []VideoOss
}{M: make([]VideoOss, 0)}

var ImgList = struct {
	sync.RWMutex
	M []ImgOss
}{M: make([]ImgOss, 0)}

func ConsumeVideo() {
	for {
		time.Sleep(time.Second / 2)
		var source string
		var target string
		if len(VideoList.M) != 0 {
			VideoList.Lock()
			source = VideoList.M[0].SourceFile
			target = VideoList.M[0].TargetFile
			VideoList.M = VideoList.M[1:]
			VideoList.Unlock()
			err := PutFile(source, target)
			// 出错之后再将数据塞回去
			if err != nil {
				VideoList.Lock()
				v := VideoOss{
					SourceFile: source,
					TargetFile: target,
				}
				VideoList.M = append(VideoList.M, v)
				VideoList.Unlock()
				continue
			}
			fmt.Println(source + "  into  " + target)
		}
	}
}

func ConsumeImg() {
	for {
		time.Sleep(time.Second / 2)
		img := ImgOss{
			VideoFile:  "",
			SourceFile: "",
			TargetFile: "",
		}
		if len(ImgList.M) != 0 {
			ImgList.Lock()
			img.VideoFile = ImgList.M[0].VideoFile
			img.SourceFile = ImgList.M[0].SourceFile
			img.TargetFile = ImgList.M[0].TargetFile
			ImgList.M = ImgList.M[1:]
			ImgList.Unlock()
			// 出错之后再将数据塞回去
			if err := ExampleReadFrameAsJpeg(img.VideoFile, 1, img.SourceFile); err != nil {
				ImgList.Lock()
				ImgList.M = append(ImgList.M, img)
				ImgList.Unlock()
				continue
			}
			// 出错之后再将数据塞回去
			if err := PutFile(img.SourceFile, img.TargetFile); err != nil {
				ImgList.Lock()
				ImgList.M = append(ImgList.M, img)
				ImgList.Unlock()
				continue
			}
		}
	}
}

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

func PutFile(SourceFile string, TargetFile string) error {
	return ossBucket.PutObjectFromFile(TargetFile, SourceFile)
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int, imgPath string) error {
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
	err = imaging.Save(img, imgPath) //这里换成我们的输出地址
	if err != nil {
		fmt.Println("wrong") //这里换成我们的err
		return err
	}
	return nil
}
