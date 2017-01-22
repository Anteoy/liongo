package main

import (
	"fmt"
	"github.com/Anteoy/liongo/service"
	"github.com/Anteoy/liongo/utils"
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
	//pNoteService.DealNoteUpload(ss)
	yp := new(utils.YamlParser)
	yamlData := yp.Parse("../resources")
	pNoteService.GetNoteByName(ss,yamlData,nil,nil)
}

func UseInfo() {
	fmt.Println(USAGE)
}
