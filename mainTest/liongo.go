package main

import (
	"fmt"
	"github.com/Anteoy/liongo/service"
)

const VERSION = "0.0.1"

const (
	USAGE = `
liongo is a static site generator in Go

Usage:

        liongo command [args...]

The commands are:

	build	        			build and generate site.
	version         			print liongo version

`
)

var httpAddr = ":8080"

func main() {
	pNoteService := new(service.PNoteService)
	var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题"
	pNoteService.DealNoteUpload(ss)
}

func UseInfo() {
	fmt.Println(USAGE)
}
