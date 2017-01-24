package service

import (
	"testing"
	"time"
	"log"
	"fmt"
)

func TestPNoteService_DealNoteUpload(t *testing.T) {
	time, terr := time.Parse("2006-01-02 15:04:05", "2017-01-20 20:12:00")
	if terr != nil {
		log.Println(terr)
	}
	fmt.Println(time)
	pNoteService := new(PNoteService)
	var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题"
	pNoteService.DealNoteUpload(ss)
}