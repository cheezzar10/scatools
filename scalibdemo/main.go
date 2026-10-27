package main

import (
	"github.com/cheezzar10/scatools/scalibdemo/imagescan"
	"os"
)

func main() {
	err := imagescan.ScanImage("debian-jdk:trixie")
	if err != nil {
		os.Exit(1)
	}
}
