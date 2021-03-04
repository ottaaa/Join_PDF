package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// // 今のディレクトリの場所を取ってくるやつ
	// dir, err := os.Getwd()

	// //　表示する
	// fmt.Println(dir)
	// fmt.Println(err)

	// ファイルの中身
	fileInfos, err := ioutil.ReadDir("./test_pdf")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(fileInfos)

	// PDFがあるディレクトリで↑をして絞り込みたい
	for _, fileInfo := range fileInfos {
		// pdfのみ出力する
		if strings.Contains(fileInfo.Name(), ".pdf") {
			// 型番で検索
			if strings.Contains(fileInfo.Name(), "") {
				fmt.Println(fileInfo.Name())
			}
		}
	}
}
