package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"index/suffixarray"
	"os"
	"path/filepath"
	"strings"
)

func contains(src string, words []string) bool {
	index := suffixarray.New([]byte(src))
	for _, x := range words {
		if len(index.Lookup([]byte(x), -1)) > 0 {
			return true
		}
	}
	return false
}

func GetPDFPath(modelNum []string) []string {
	var inFiles []string
	filepath.Walk(filepath.FromSlash("./test_pdf/"), func(path string, info os.FileInfo, err error) error {
		if contains(info.Name(), modelNum) && strings.Contains(info.Name(), "pdf") {
			inFiles = append(inFiles, path)
		}
		return nil
	})
	return inFiles
}

func Merge(input []string) {
	var config = pdfcpu.NewDefaultConfiguration()
	api.MergeAppendFile(input, "out.pdf", config)
}
