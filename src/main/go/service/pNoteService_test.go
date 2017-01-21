package service

import "testing"

func TestPNoteService_DealNoteUpload(t *testing.T) {
	pNoteService := new(PNoteService)
	var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题"
	pNoteService.DealNoteUpload(ss)
}