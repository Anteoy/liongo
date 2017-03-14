package impl

import (
	"strings"
	"os"
	. "github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/go-gypsy/yaml"
	"log"
)

type ProcessPnoteloginPage struct {}

//生成pnote login.html
func (processPnoteloginPage *ProcessPnoteloginPage)Dispose(root string)  {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//var cfg = yCfg.(*yaml.File)
	t := ParseTemplate(root, PNOTELOGIN_TPL, cfg)
	targetFile := PUBLISH_DIR + "/pnotelogin.html"
	//创建targetFile
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	exErr := t.Execute(fout, nil)
	if exErr != nil {
		log.Fatal(exErr) //TODO
	}
}
