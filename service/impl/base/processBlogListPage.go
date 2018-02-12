package impl

import (
	"log"
	"os"
	"strings"

	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"github.com/Anteoy/liongo/utils/logrus"
	"strconv"
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

	//处理分页
	totalPage := len(Articlesl)/10
	//每页显示个数
	pageSize := 10
	//当前需要渲染的articlesl
	for i := 0;i < totalPage;i++ {
		targetFile := PUBLISH_DIR + "/blog_"+ strconv.Itoa(i) + ".html"
		pre := "http://127.0.0.1:8080" + "/blog_"+ strconv.Itoa(i-1) + ".html"
		next := 	"http://127.0.0.1:8080" + "/blog_"+ strconv.Itoa(i+1) + ".html"
		fout, err := os.Create(targetFile)
		if err != nil {
			logrus.Error("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		start := i * pageSize
		end := (i+1)*pageSize
		curArticle := Articlesl[start:end]
		m := map[string]interface{}{
		"ar": curArticle[:],
		"nav": NavBarsl,
		"cats": Classifiesm,
		"pre":pre,
		"next": next,
		"i":i+1,
		"total":totalPage,
		}
		exErr := t.Execute(fout, m)
		if exErr != nil {
			log.Fatal(exErr)
		}
	}
}
