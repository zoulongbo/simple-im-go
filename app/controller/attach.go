package controller

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"im/app/util"
	"im/config"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const filePath  = "./public/img"

func Upload(w http.ResponseWriter, r *http.Request) {
	uploadLocal(w, r)
}

func init(){
	_, err :=os.Stat(filePath)
	if err != nil {
		os.MkdirAll(filePath, os.ModePerm)
	}
}

/**
	1、 存储位置  ./public/img		必须创建好
	2、 url格式  ./public/img/xxx.png
 */
func uploadLocal(writer http.ResponseWriter, request *http.Request) {
	//获取源文件
	file, header, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//获取文件类型
	ext := ".png"
	oFileName := header.Filename
	//默认取后缀
	tmp := strings.Split(oFileName, ".")
	if len(tmp) > 1 {
		ext = tmp[len(tmp) - 1]
	}
	//若指定类型 取类型
	fileType := request.FormValue("filetype")
	if len(fileType) > 0 {
		ext = fileType
	}
	//文件名
	fileName := fmt.Sprintf("%d-%d%s", time.Now().Unix(), rand.Int31(), "." + ext)
	//创建新文件
	dstFile, osErr := os.Create(filePath + "/" + fileName)
	if osErr != nil {
		util.RespFail(writer, osErr.Error())
		return
	}
	//将源文件copy到新文件
	_, ioErr := io.Copy(dstFile, file)
	if ioErr != nil {
		util.RespFail(writer, ioErr.Error())
	}
	//将文件路径替换为url
	url := "/public/img/" + fileName
	util.RespOk(writer, url, "")
	return
}

func UploadOss(writer http.ResponseWriter, request * http.Request) {
	//获得上传文件
	srcFile, header, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//获取文件类型
	ext := ".png"
	oFileName := header.Filename
	//默认取后缀
	tmp := strings.Split(oFileName, ".")
	if len(tmp) > 1 {
		ext = tmp[len(tmp) - 1]
	}
	//若指定类型 取类型
	fileType := request.FormValue("filetype")
	if len(fileType) > 0 {
		ext = fileType
	}
	//文件名
	fileName := fmt.Sprintf("%d-%d%s", time.Now().Unix(), rand.Int31(), "." + ext)
	// 创建OSSClient实例。
	client, err := oss.New(config.ALiOssEedPoint, config.ALiOssAccessKeyId, config.ALiOssAccessKeySecret)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.ALiOssBucket)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 上传字符串。
	err = bucket.PutObject(fileName, srcFile, storageType, objectAcl)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	url := "https://" + config.ALiOssBucket + "." + config.ALiOssEedPoint + "/" + fileName
	util.RespOk(writer, url, "上传成功")
	return
}