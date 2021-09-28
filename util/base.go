package util

import (
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
)

func IF(f bool, a interface{}, b interface{}) interface{} {
	if f {
		return a
	}
	return b
}
func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func CreatePathIfNotExists(path string) error {
	if !IsExists(path) {
		if err := os.MkdirAll(path, 0666); err != nil {
			return err
		}
	}
	return nil
}
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// GetMPFDContentType 根据gin文档，FileName不可靠，使用文件头特征识别mime
func GetMPFDContentType(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	buffer := make([]byte, 512) // 前512字节
	if _, err = src.Read(buffer); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
func IsInStringSlice(key string, str []string) bool {
	for _, s := range str {
		if key == s {
			return true
		}
	}
	return false
}
