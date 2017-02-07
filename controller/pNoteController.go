package controller

import (
	"io"
	"net/http"
	"github.com/Anteoy/liongo/dao/mysql"
	"github.com/Anteoy/liongo/modle"
	//"github.com/Anteoy/liongo/utils" TODO mgo session 共用问题
	//"github.com/Anteoy/liongo/service"
	"fmt"
	//"os"
	"github.com/Anteoy/liongo/service"
	"regexp"
	"log"
	"github.com/Anteoy/blackfriday"
	"html"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/Anteoy/liongo/dao/mongo"
)

type PNoteController struct{}



func (pNoteController *PNoteController)Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids := r.Form["id"]
	if ids == nil {
		io.WriteString(w, "请输入账号")
		return
	}
	id := ids[0]
	passwds := r.Form["passwd"]
	if passwds == nil {
		io.WriteString(w, "请输入密码")
		return
	}
	passwd := passwds[0]
	user := mysql.GetUserForEmail(id)
	if user != nil && user.Password == passwd {
		//fmt.Fprint(w,"<h1>login success!!!</h1>")
		http.ServeFile(w, r, "../views/serve/pnotelist.html") //ok
	} else {
		fmt.Fprint(w,"<h1>login faild!!!用户名或密码不正确！！！</h1>")
		http.ServeFile(w, r, "./static/html/login.html")
	}
	//var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题" TODO mongo session 共用问题
	////pNoteService.DealNoteUpload(ss)
	//pNoteService := new(service.PNoteService)
	//yp := new(utils.YamlParser)
	//yamlData := yp.Parse("../resources")
	//pNoteService.GetNoteByName(ss,yamlData,w,r)


}

//通过参数在List页面向详情页面跳转
func (p *PNoteController) GetNote(w http.ResponseWriter,r  *http.Request)  {
	r.ParseForm()
	links := r.Form["link"]
	if links == nil {
		io.WriteString(w, "无法路由，参数link为空！！！")
		return
	}
	fmt.Println(links[0])
	//创建html文件路径
	targetFile := service.PUBLISH + "/notes/" + links[0] +".html"
	fmt.Println(targetFile)
	http.ServeFile(w, r, targetFile)
}

//获取笔记md文件并存入mongo

func (pNoteController *PNoteController) DataTomongo(notemd *modle.Note){

}

//处理pnote上传
func (pNoteController *PNoteController) PNCommit(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	titles := r.Form["title"]
	title := titles[0]
	if len(title) == 0 {
		io.WriteString(w, "请输入标题")//TODO
		return
	}
	fmt.Println(title)
	contents := r.Form["content"]
	content := contents[0]
	if len(content) == 0 {
		io.WriteString(w, "请输入正文")//TODO
		return
	}
	fmt.Println(content)

	//md处理
	// 判断是否为空
	if len(content) == 0 {
		log.Fatal("md is nil")
		return
	}
	//处理md为html
	//markdown字符串转为ASCII html代码
	htmlByte := blackfriday.MarkdownCommon([]byte(content))
	//反转义实体如“& lt;”成为“<” 把byte转位strings
	htmlStr := html.UnescapeString(string(htmlByte))
	//正则匹配并替换
	re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
	htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
	fmt.Println(htmlStr)

	//构造struct
	//timestamp
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
	//Time
	//时间解析
	time, terr := time.Parse("2006-01-02 15:04:05", "2017-01-10 20:12:00")
	if terr != nil {
		log.Println(terr)
	}
	//装配struct
	note := &modle.Note{Name: title, Content: htmlStr,Title: title,Date: tm.Format("2006-01-02 03:04:05"),Time:time}
	//获取session
	var ch chan *mgo.Session= make(chan *mgo.Session,1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session//must init
	sess = <- (chan *mgo.Session)(ch)//must do
	c :=sess.DB("liongo").C("note")//获取数据
	//存入mongo
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	return

}