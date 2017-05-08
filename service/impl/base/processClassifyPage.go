package impl

import (
	"log"
	"os"
	"strings"

	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"github.com/Anteoy/liongo/utils/logrus"
)

type ProcessClassifyPage struct{}

//生成classify.html
func (processClassifyPage *ProcessClassifyPage) Dispose(dir string) {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//var cfg = yCfg.(*yaml.File)

	t := ParseTemplate(dir, CLASSIFY_TPL, cfg)
	targetFile := PUBLISH_DIR + "/classify.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		logrus.Error("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	m := map[string]interface{}{"cats": Classifiesm, "nav": NavBarsl}
	exErr := t.Execute(fout, m)
	if exErr != nil {
		log.Fatal(exErr)
	}
}
