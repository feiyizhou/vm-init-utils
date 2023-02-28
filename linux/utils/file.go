package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	FileMode0755 = 0755
	FileMode0644 = 0644
)

// IsExist 文件是否存在，true-存在，false-不存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// ReadFileToLines 按行读取文件为字符串数组
func ReadFileToLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var strArr []string
	for {
		lineBytes, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		strArr = append(strArr, string(lineBytes))
	}
	return strArr, err
}

// WriteToFile 写内容到文件中
func WriteToFile(content interface{}, path string) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, FileMode0755); err != nil {
			return err
		}
	}
	if err := ioutil.WriteFile(path, bytes, FileMode0644); err != nil {
		return err
	}
	return nil
}
