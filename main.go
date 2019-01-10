package main

import (
	_ "file-upload-download/routers"
	"fmt"
	"os"

	"github.com/astaxie/beego"
)

// pathExists PathExists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// createDic createDic
func createDic(dicName string) {
	exist, err := pathExists(dicName)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", dicName)
	} else {
		fmt.Printf("no dir![%v]\n", dicName)
		// 创建文件夹
		err := os.Mkdir(dicName, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

func main() {
	createDic("./static/files")
	beego.Run()
}
