package impl

import (
	"strings"
	"os"
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"log"
)

type ProcessEveryArticlePage struct {}

//根据日期生成每一个article的html文件
func (processEveryArticlePage *ProcessEveryArticlePage)Dispose(root string)  {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//var cfg = yCfg.(*yaml.File)
	log.Println(cfg.Get("title"))
	t := ParseTemplate(root, POSTS_TPL, cfg)
	for _, fileInfo := range Articles {
		//根据时间生成日期类目录 /yyyy/MM/dd
		p := processArticleUrl(*fileInfo)
		if !IsExists(PUBLISH_DIR + "/articles/" + p) {
			os.MkdirAll(PUBLISH_DIR +"/articles/"+p, 0777)
		}
		targetFile := PUBLISH_DIR + "/articles/" + fileInfo.Link
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		m := map[string]interface{}{"fi": fileInfo, "nav": NavBarList, "cats": Classifies}
		t.Execute(fout, m)
	}
}
