package impl

import (
	"strings"
	"os"
	"bufio"
	"io"
	. "github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/go-gypsy/yaml"
	"strconv"
	"log"
	"github.com/Anteoy/blackfriday"
	"html"
)

type ProcessUserDefinedPages struct{}

//生成用户自定义的pages
func (processUserDefinedPages *ProcessUserDefinedPages)Dispose(dir string)  {
	//判断结尾是否/
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//获取配置文件页面信息
	getPagesInfo(YamlData)

	t := ParseTemplate(dir, PAGES_TPL, cfg)

	//Pages为前面解析好的CustomPage数组
	for _, p := range Pages {
		p.Id = strings.TrimSuffix(p.Id, " ")
		filePath := dir + "pages/" + p.Id + ".md"
		if !IsExists(filePath) {
			log.Println(filePath + " is not found!")
			os.Exit(1)
		}
		f, err := os.Open(filePath)
		if err != nil {
			log.Println(err)

		}
		defer f.Close()
		rd := bufio.NewReader(f)
		var markdownStr string
		//以行读取md文件
		for {
			buf, _, err := rd.ReadLine()

			if err == io.EOF {
				break
			} else {
				content := string(buf)
				markdownStr += content + "\n"
			}

		}

		//转化位二进制html
		htmlByte := blackfriday.MarkdownCommon([]byte(markdownStr))
		//转化位htmlstrings
		htmlStr := html.UnescapeString(string(htmlByte))
		//-1无限制完全替换
		htmlStr = strings.Replace(htmlStr, "<pre><code", `<pre class="prettyprint linenums"`, -1)
		htmlStr = strings.Replace(htmlStr, `</code>`, "", -1)

		p.Content = htmlStr //设置markdown文章内容
		//log.Println(p.Content)
		if !IsExists(PUBLISH_DIR + "/pages/") {
			os.MkdirAll(PUBLISH_DIR +"/pages/", 0777)
		}
		targetFile := PUBLISH_DIR + "/pages/" + p.Id + ".html"
		//创建target html
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		//p .md article信息 nav 自定义的额外导航条信息 暂移除 "cats": categories "newly":articles[:NEWLY_ARTICLES_COUNT-1]
		m := map[string]interface{}{"p": p, "nav": NavBarList}
		t.Execute(fout, m)
	}
}

//获取配置文件页面信息 最后植入指定pages中
func getPagesInfo(yamls map[string]interface{}) {
	yCfg := yamls["pages.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//统计配置pages.yml中配置个数
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		//strconv.Itoa转换int i为string
		//获取配置的id和title yaml配置使用数组- /pages/id.md
		id, erri := cfg.Get("[" + strconv.Itoa(i) + "].id")
		log.Println("[" + strconv.Itoa(i) + "].id")
		if nil != erri {
			log.Println(erri)
		}
		//指定页面的内容title
		title, errt := cfg.Get("[" + strconv.Itoa(i) + "].title")
		if nil != errt {
			log.Println(errt)
		}
		log.Println("[" + strconv.Itoa(i) + "].title")
		page := CustomPage{id, title, ""}
		//追加到pages
		Pages = append(Pages, &page)
	}

}
