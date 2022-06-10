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

// Init 初始化视频和图片oss客户端
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

// VideoOss 存储create_video视频流数据处理中间信息
type VideoOss struct {
	SourceFile string
	TargetFile string
}

// ImgOss 存储create_video视频转图片处理的中间信息
type ImgOss struct {
	VideoFile  string
	SourceFile string
	TargetFile string
}

// VideoList 并发存储create_video视频流数据处理中间信息列表
var VideoList = struct {
	sync.RWMutex
	M []VideoOss
}{M: make([]VideoOss, 0)}

// ImgList 并发存储create_video视频转图片处理的中间信息列表
var ImgList = struct {
	sync.RWMutex
	M []ImgOss
}{M: make([]ImgOss, 0)}

// ConsumeVideo 消费需要进行视频流处理的信息
func ConsumeVideo() {
	for {
		time.Sleep(time.Second / 2)
		var source string
		var target string
		if len(VideoList.M) != 0 {
			// 取出一条数据进行处理
			VideoList.Lock()
			source = VideoList.M[0].SourceFile
			target = VideoList.M[0].TargetFile
			VideoList.M = VideoList.M[1:]
			VideoList.Unlock()
			// 将视频信息存入OSS文件服务器
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

// ConsumeImg 消费视频转图片的中间信息
func ConsumeImg() {
	for {
		time.Sleep(time.Second / 2)
		img := ImgOss{
			VideoFile:  "",
			SourceFile: "",
			TargetFile: "",
		}
		if len(ImgList.M) != 0 {
			// 取出一条数据进行处理
			ImgList.Lock()
			img.VideoFile = ImgList.M[0].VideoFile
			img.SourceFile = ImgList.M[0].SourceFile
			img.TargetFile = ImgList.M[0].TargetFile
			ImgList.M = ImgList.M[1:]
			ImgList.Unlock()
			// 从视频流中解析出图片，出错之后再将数据塞回去
			if err := ExampleReadFrameAsJpeg(img.VideoFile, 1, img.SourceFile); err != nil {
				ImgList.Lock()
				ImgList.M = append(ImgList.M, img)
				ImgList.Unlock()
				continue
			}
			// 将图片上传到图片服务器，出错之后再将数据塞回去
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

// handleError 错误处理函数
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// Init 将一个文件传到OSS服务器中
func PutFile(SourceFile string, TargetFile string) error {
	return ossBucket.PutObjectFromFile(TargetFile, SourceFile)
}

// ExampleReadFrameAsJpeg 将视频文件转化成图片文件
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
