package main

import (
	"fmt"
	"index/suffixarray"
	"os"
	"path/filepath"
)

func main() {

	var modelNum =	[]string{"AAA", "AAB"}
	filepath.Walk("./test_pdf/", func(path string, info os.FileInfo, err error) error {
		if contains(info.Name(), modelNum){
			fmt.Println(info.Name())
		}
		return nil        
	})
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
