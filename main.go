package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 今のディレクトリの場所を取ってくるやつ
	dir, err := os.Getwd()

	//　表示する
	fmt.Println(dir)
	fmt.Println(err)

	// ファイルの中身
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("******************************")

	// PDFがあるディレクトリで↑をして絞り込みたい
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
