package impl

import (
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"log"
	"os"
	"strings"
)

type ProcessIndexPage struct{}

func (processIndex *ProcessIndexPage) Dispose(dir string) {

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	pEiz := YamlData["config.yml"]
	var cfg = pEiz.(*yaml.File)
	t := ParseTemplate(dir, INDEX_TPL, cfg)

	targetFile := PUBLISH_DIR + "/index.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	m := map[string]interface{}{"nav": NavBarsl, "cats": Classifiesm}
	exErr := t.Execute(fout, m)
	if exErr != nil {
		log.Fatal(exErr) //TODO
	}
}
