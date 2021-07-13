package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
			dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
				p := file.URI().Path()
				modelNum := ParseCsv(p)
				//// 結合するPDFファイルのパスを取得
				var targetPDF = GetPDFPath(modelNum)
				//// 結合
				Merge(targetPDF)
			}, w)
		}),
	))

	w.ShowAndRun()

	//// TODO ビルド
	//// GOOS=windows GOARCH=amd64 go build main.go
	//// ビルドした後にwinでの動作検証
}
