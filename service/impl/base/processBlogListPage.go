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

type ProcessBlogListPage struct{}

//渲染生成/blog.html
func (processBlogList *ProcessBlogListPage) Dispose(dir string) {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//var cfg = yCfg.(*yaml.File)
	t := ParseTemplate(dir, BLOG_LIST_TPL, cfg)

	targetFile := PUBLISH_DIR + "/blog.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		logrus.Error("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	m := map[string]interface{}{"ar": Articlesl[:], "nav": NavBarsl, "cats": Classifiesm}
	exErr := t.Execute(fout, m)
	if exErr != nil {
		log.Fatal(exErr)
	}

}
