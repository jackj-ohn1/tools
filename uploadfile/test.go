package myself

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

//封装一个上传文件的包
func UploadFile(name string, url string) {
	buf, contenttype := ReadLocal(name)
	req, err := http.NewRequest("POST", url, buf)
	Error(err)
	req.Header.Set("Content-Type", contenttype)
	/*
		可自己添加需要的表头
		...
	*/
	resp, err := http.DefaultClient.Do(req)
	Error(err)
	body, err := ioutil.ReadAll(resp.Body)
	Error(err)
	fmt.Println(string(body))
}

//从本地读取文件
func ReadLocal(name string) (io.Reader, string) {
	file, err := os.Open(name)
	suffix := strings.Split(name, ".")[1]
	Error(err)
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	contype := writer.FormDataContentType()
	form, err := writer.CreateFormFile("file", "local"+"."+suffix)
	Error(err)
	_, err = io.Copy(form, file)
	Error(err)
	writer.Close()
	defer file.Close()
	return buf, contype
}
