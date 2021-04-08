package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	// brew install tesseract-lang
	// tessdata https://tesseract-ocr.github.io/tessdoc/Data-Files.html
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/Users/admin/Desktop/gray.jpeg")
	client.SetLanguage( "chi_sim")
	text, _ := client.Text()
	fmt.Println(text)
}
