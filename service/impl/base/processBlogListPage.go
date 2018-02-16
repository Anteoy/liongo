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
	"fmt"
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
	fmt.Println("============1",len(Articlesl))
	//TODO
	if len(Articlesl)%10 != 0 {
		totalPage++
	}
	//每页显示个数
	pageSize := 10
	//当前需要渲染的articlesl
	for i := 0;i < totalPage;i++ {
		targetFile := PUBLISH_DIR + "/blog_"+ strconv.Itoa(i+1) + ".html"
		var pre string
		var next string
		//update count + 1
		if os.Getenv("liongo_env") == "online"{
			pre = "http://anteoy.site" + "/blog_"+ strconv.Itoa(i-1+1) + ".html"
			next = 	"http://anteoy.site" + "/blog_"+ strconv.Itoa(i+1+1) + ".html"
		}else {
			pre = "http://127.0.0.1:8080" + "/blog_"+ strconv.Itoa(i-1+1) + ".html"
			next = 	"http://127.0.0.1:8080" + "/blog_"+ strconv.Itoa(i+1+1) + ".html"
		}
		fout, err := os.Create(targetFile)
		if err != nil {
			logrus.Error("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		start := i * pageSize
		var end int
		//the last
		if (i+1)== totalPage{
			end = len(Articlesl)
		}else{
			end = (i+1)*pageSize
		}
		curArticle := Articlesl[start:end]
		var display0 string
		var display1 string
		if i == 0 {
			display0 = "none"
		}else{
			display0 = ""
		}
		if (i+1)== totalPage {
			display1 = "none"
		}else {
			display1 = ""
		}
		m := map[string]interface{}{
		"ar": curArticle[:],
		"nav": NavBarsl,
		"cats": Classifiesm,
		"pre":pre,
		"next": next,
		"i":i+1,
		"total":totalPage,
		"display0": display0,
		"display1":display1,
		}
		exErr := t.Execute(fout, m)
		if exErr != nil {
			panic(exErr)
			log.Fatal(exErr)
		}
	}
}
