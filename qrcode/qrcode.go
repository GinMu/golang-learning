package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {

	qr, err := qrcode.New("https://github.com/GinMu", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.Black
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./golang_qrcode.png")
	}
}
