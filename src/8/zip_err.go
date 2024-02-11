package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	nonZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(nonZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("zip形式ではありません")
	}
}
