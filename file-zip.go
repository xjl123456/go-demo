package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	src := "/tmp/data-nfs/unicom123/2/test-freeze1/qqqq/LATEST/functions/Knative/hello-world-go"
	zip := "/tmp/zip"
	err := CreateZipFile(src, zip)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateZipFile(srcfile, zipfile string) error {
	zipFile, err := os.Create(zipfile)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	writer := zip.NewWriter(zipFile)
	defer writer.Close()
	err = filepath.Walk(srcfile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//log.Println(path, info, err)
		infoHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		infoHeader.Method = zip.Deflate
		//heeder必须是相对路径，所以取出头部的路径标识
		infoHeader.Name = strings.TrimPrefix(path, string(filepath.Separator))
		//检测如果是文件夹，加上路径分隔符
		if info.IsDir() {
			infoHeader.Name += string(filepath.Separator)
		}
		w, err := writer.CreateHeader(infoHeader)
		if err != nil {
			return err
		}
		//打开文件
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		//把文件放进去
		_, err = io.Copy(w, file)
		//todo 这个错误暂时不看
		//if err != nil {
		//	return err
		//}

		return nil
	})
	log.Println(err)
	return nil
}
