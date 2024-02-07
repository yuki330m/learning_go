package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 { // ファイル名が指定されている？
		log.Fatal("ファイル名が指定されていません")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		break
	}
}
