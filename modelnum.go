package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"os"
)

func ParseCsv(path string) []string {
	fl, _ := os.Open(path)
	defer fl.Close()
	r := csv.NewReader(transform.NewReader(fl, japanese.ShiftJIS.NewDecoder()))

	rows, _ := r.ReadAll()
	header := rows[0]
	modelNumIdx := 0
	for idx, headerStr := range header {
		if headerStr == "商品コード" {
			modelNumIdx = idx
		}
	}
	var modelNums []string
	for _, row := range rows[1:] {
		modelNums = append(modelNums, row[modelNumIdx])
	}
	fmt.Println(modelNums)
	return modelNums
}
