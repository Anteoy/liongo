package service

import (
	"github.com/Anteoy/blackfriday"
	"qiniupkg.com/x/log.v7"
	"github.com/Anteoy/liongo/modle"
	"github.com/Anteoy/liongo/dao/mongo"
	"github.com/Anteoy/go-gypsy/yaml"
	"regexp"
	"html"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"os"
	"net/http"
)

type PNoteService struct{}

//处理Note上传
func (p *PNoteService) DealNoteUpload(md string)  error {
	// 判断是否为空
	if len(md) == 0 {
		log.Fatal("md is nil")
		return nil
	}
	//处理md为html
	//markdown字符串转为ASCII html代码
	htmlByte := blackfriday.MarkdownCommon([]byte(md))
	//反转义实体如“& lt;”成为“<” 把byte转位strings
	htmlStr := html.UnescapeString(string(htmlByte))
	//正则匹配并替换
	re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
	htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
	//装配struct
	note := &modle.Note{Name: "test1", Content: htmlStr}
	fmt.Printf(note.Content)
	c := mongo.Session.DB("liongo").C("note")
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
//从Mongo中拉取具体note
func (p *PNoteService) GetNoteByName(name string,yamls map[string]interface{},w http.ResponseWriter, r *http.Request) error {
	if len(name) == 0 {
		fmt.Println("传入Note name为空，请检查！！！")
		return nil
	}
	//从mongo中获取noteinfo
	//获取连接
	c := mongo.Session.DB("liongo").C("note")
	//获取数据
	note := modle.Note{}
	err := c.Find(bson.M{"name": "test1"}).One(&note)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	fmt.Println(note.Content)
	fmt.Println(note.Name)
	//new 模板对象
	t := template.New("pSpecificNote.tpl")
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	//向模板中注入函数
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	t.Funcs(template.FuncMap{"get": cfg.Get})

	//openfile := "../resources/templates/default/pSpecificNote.tpl"
	//
	//if !isExists(openfile) {
	//	log.Println(openfile + " can not be found!")
	//	os.Exit(1)
	//}


	//从模板文件解析
	t, errp := t.ParseFiles("/root/IdeaProjects/liongo/src/github.com/Anteoy/liongo/resources/templates/default/pSpecificNote.tpl")
	if errp != nil {
		log.Error(errp)
		panic(err)
	}
	//创建html文件
	targetFile := PUBLISH + "/notes/" + note.Name+".html"
	fout, err := os.Create(targetFile)
	m := map[string]interface{}{"fi": note,"nav": navBarList, "cats": classifies}
	//执行模板的merge操作，输出到fout
	t.Execute(fout, m)
	http.ServeFile(w, r, targetFile)
	defer mongo.Session.Close()
	defer fout.Close()
	return nil

}
