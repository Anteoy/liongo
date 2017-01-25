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
	//pNoteService.DealNoteUpload(ss)
	yp := new(utils.YamlParser)
	yamlData := yp.Parse("../resources")
	pNoteService.GetNoteByName(yamlData,nil,nil)//从mgo中搜集并生成所有notes
	//pNoteService.QueryAll()
	//pNoteService.PreProcessNotes()
}

func UseInfo() {
	fmt.Println(USAGE)
}
