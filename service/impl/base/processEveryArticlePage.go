package impl

import (
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"log"
	"os"
	"strings"
)

type ProcessEveryArticlePage struct{}

//根据日期生成每一个article的html文件
func (processEveryArticlePage *ProcessEveryArticlePage) Dispose(root string) {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//var cfg = yCfg.(*yaml.File)
	t := ParseTemplate(root, POSTS_TPL, cfg)
	for _, articleConfig := range Articlesl {
		//根据时间生成日期类目录 /yyyy/MM/dd
		p := processArticleUrl(*articleConfig)
		if !IsExists(PUBLISH_DIR + "/articles/" + p) {
			os.MkdirAll(PUBLISH_DIR+"/articles/"+p, 0777)
		}
		targetFile := PUBLISH_DIR + "/articles/" + articleConfig.Link
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		m := map[string]interface{}{"fi": articleConfig, "nav": NavBarsl, "cats": Classifiesm}
		t.Execute(fout, m)
	}
}
