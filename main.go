package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	//"bufio"
	//"fmt"
	"index/suffixarray"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("PDF結合")
	w.Resize(fyne.NewSize(600, 480))
	hello := widget.NewLabel("csvを読み込んでください。")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			dialog.ShowFolderOpen(func(fyne.ListableURI, error) {}, w)
		}),
	))

	w.ShowAndRun()

	//// TODO ビルド
	//// GOOS=windows GOARCH=amd64 go build main.go
	//// ビルドした後にwinでの動作検証

	//// AAB000001 AAB000002
	//// 結合する型番
	//var modelNum = input_arr
	//// 結合するPDFファイルのパスを取得
	//var targetPDF = getPDFPath(modelNum)
	//// 結合
	//merge(targetPDF)
}

func contains(src string, words []string) bool {
	index := suffixarray.New([]byte(src))
	for _, x := range words {
		if len(index.Lookup([]byte(x), -1)) > 0 {
			return true
		}
	}
	return false
}

func getPDFPath(modelNum []string) []string {
	inFiles := []string{}
	filepath.Walk("./test_pdf/", func(path string, info os.FileInfo, err error) error {
		if contains(info.Name(), modelNum) && strings.Contains(info.Name(), "pdf") {
			inFiles = append(inFiles, path)
		}
		return nil
	})
	return inFiles
}

func merge(input []string) {
	var config = pdfcpu.NewDefaultConfiguration()
	api.MergeAppendFile(input, "out.pdf", config)
}
