package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//获取私钥文件
func GetFileContent(address, dirname string) (string, error) {
	data, err := ioutil.ReadDir(dirname)//得到路径中所有文件名
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(strings.ToLower(v.Name()), strings.ToLower(address)[2:]) > 0 {
			//代表找到文件
			file, err := os.Open(dirname + "\\" + v.Name())
			if err != nil {
				fmt.Printf("Failed to open file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			data := make([]byte, 1024)
			count, err := file.Read(data)
			if err != nil || count < 1 {
				fmt.Printf("Failed to read file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			return string(data[:count]), nil
		}
	}
	return "", nil
}
