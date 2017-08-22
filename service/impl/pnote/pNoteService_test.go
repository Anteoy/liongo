package pnote

import (
	"testing"
	"time"

	"github.com/Anteoy/liongo/utils/logrus"
)

func TestPNoteService_DealNoteUpload(t *testing.T) {
	time, terr := time.Parse("2006-01-02 15:04:05", "2017-01-20 20:12:00")
	if terr != nil {
		logrus.Error(terr)
	}
	pNoteService := new(PNoteService)
	var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题"
	pNoteService.DealNoteUpload(ss)
}
