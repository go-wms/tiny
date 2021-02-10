package file

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
)

func Save(imageUrl string) (string ,error) {
	fileName := path.Base(imageUrl)


	res, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)

	filePath := "tmp/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	_, err = io.Copy(writer, reader)
	return filePath,err
}
