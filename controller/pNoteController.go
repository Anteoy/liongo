package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/Anteoy/liongo/dao/mysql"
	"github.com/Anteoy/liongo/model"
	//"github.com/Anteoy/liongo/utils" TODO mgo session 共用问题
	//"github.com/Anteoy/liongo/service"
	"fmt"
	//"os"
	"encoding/json"
	"html"
	"regexp"
	"time"

	"github.com/Anteoy/blackfriday"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/dao/mongo"
	logrus "github.com/Anteoy/liongo/utils/logrus"
	_ "github.com/Anteoy/liongo/utils/memory" //must this is the first init ,it cost any time
	"github.com/Anteoy/liongo/utils/session"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"github.com/Anteoy/liongo/utils"
)

type PNoteController struct{}

//初始化session
var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewSessionManager("memory", "goSessionid", 3600)
	go globalSessions.GC()
}

func (pNoteController *PNoteController) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//start session
	sess := globalSessions.SessionStart(w, r)
	logrus.Debugf("Login() sessionID 为： %s", sess.SessionID())
	logrus.Debugf("Login() sess.Get(sess.SessionID()) 为： %v", sess.Get(sess.SessionID()))
	//test
	//if sess.Get(sess.SessionID()) == nil {
	//	io.WriteString(w, "no sesssion")
	//	return
	//}
	//value设值为sessionid
	if sess.Get(sess.SessionID()) == nil {
		sess.Set(sess.SessionID(), sess.SessionID())
	}
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
		http.ServeFile(w, r, "../views/serve/pnotelist.html") //ok
	} else {
		fmt.Fprint(w, "<h1>login faild!!!用户名或密码不正确！！！</h1>")
		http.ServeFile(w, r, "./static/html/login.html")
	}
	//var ss string = "#For my memory scan ##1. 右上角资源监测 避免公司低配资源占用问题" TODO mongo session 共用问题
	////pNoteService.DealNoteUpload(ss)
	//pNoteService := new(service.PNoteService)
	//yp := new(utils.YamlParser)
	//yamlData := yp.Parse("../resources")
	//pNoteService.GetNoteByName(ss,yamlData,w,r)

}

type LoginRreq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (pNoteController *PNoteController) LoginR(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//start session
	sess := globalSessions.SessionStart(w, r)
	logrus.Debugf("Login() sessionID 为： %s", sess.SessionID())
	logrus.Debugf("Login() sess.Get(sess.SessionID()) 为： %v", sess.Get(sess.SessionID()))
	if r.Method == "OPTIONS" {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	req := LoginRreq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	log.Printf("收到登录请求，参数为：%+v\n", req)
	user := mysql.GetUserForEmail(req.UserName)
	if user != nil && user.Password == req.PassWord {
		token,err := utils.GenToken(1);
		if err != nil {
			s := CommonReturnModel{
				Code:    "500",
				Message: "服务器生成token失败",
			}
			b, _ := json.Marshal(s)
			w.Write(b)
			return
		}
		fmt.Printf("%s,%+v\n",token,err)
		s := LoginResModel{
			Code:    "200",
			Message: "登录成功",
			Token: token,
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return //ok
	} else {
		s := CommonReturnModel{
			Code:    "403",
			Message: "用户名或密码错误",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
}

//通过参数在List页面向详情页面跳转
func (p *PNoteController) GetNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	links := r.Form["link"]
	if links == nil {
		io.WriteString(w, "无法路由，参数link为空！！！")
		return
	}
	//创建html文件路径
	targetFile := PUBLISH_DIR + "/notes/" + links[0] + ".html"
	http.ServeFile(w, r, targetFile)
}

//获取笔记md文件并存入mongo

func (pNoteController *PNoteController) DataTomongo(notemd *model.Note) {

}

type CommonReturnModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginResModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token string `json:"token"`
}

//处理pnote上传
func (pNoteController *PNoteController) PNCommit(w http.ResponseWriter, r *http.Request) {

	//start session
	sessA := globalSessions.SessionStart(w, r)
	//test
	//if sess.Get(sess.SessionID()) == nil {
	//	io.WriteString(w, "no sesssion")
	//	return
	//}
	//value设值为sessionid
	if sessA.Get(sessA.SessionID()) == nil {
		s := CommonReturnModel{
			Code:    "502",
			Message: `session 失效，請重新登录！`,
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}

	r.ParseForm()
	titles := r.Form["title"]
	title := titles[0]
	if len(title) == 0 {
		io.WriteString(w, "请输入标题") //TODO
		return
	}
	contents := r.Form["content"]
	content := contents[0]
	if len(content) == 0 {
		io.WriteString(w, "请输入正文") //TODO
		return
	}

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

	//构造struct
	//timestamp
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	//Time
	//时间解析 strconv.FormatInt(time.Now().Unix(),10) base：进位制（2 进制到 36 进制 这种格式不行
	//time, terr := time.Parse("2006-01-02 15:04:05", strconv.FormatInt(time.Now().Unix(),10))
	//时间解析
	time, terr := time.Parse("2006-01-02 15:04:05", tm.Format("2006-01-02 03:04:05"))
	if terr != nil {
		logrus.Error(terr)
	}
	//装配struct
	note := &model.Note{Name: title, Content: htmlStr, Title: title, Date: tm.Format("2006-01-02 03:04:05"), Time: time}
	//获取session
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note") //获取数据
	//存入mongo
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	s := CommonReturnModel{
		Code:    "200",
		Message: `上传成功`,
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return

}

type RPNCommitReq struct {
	Title string `json:"title"`
	Token string `json:"token"`
	Content string `json:"content"`
}

func (pNoteController *PNoteController) RPNCommit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	if r.Method == "OPTIONS" {
		return
	}
	//start session
	sessA := globalSessions.SessionStart(w, r)
	//test
	//if sess.Get(sess.SessionID()) == nil {
	//	io.WriteString(w, "no sesssion")
	//	return
	//}
	//value设值为sessionid
	if sessA.Get(sessA.SessionID()) == nil {
		// todo tmp close
		//s := CommonReturnModel{
		//	Code:    "502",
		//	Message: `session 失效，請重新登录！`,
		//}
		//b, _ := json.Marshal(s)
		//w.Write(b)
		//return
	}
	s := CommonReturnModel{
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := RPNCommitReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Title == "" {
		s = CommonReturnModel{
			Code:    "406",
			Message: "请输入标题..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	title := req.Title
	if req.Content == "" {
		s = CommonReturnModel{
			Code:    "406",
			Message: "请输入正文..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	token := req.Token
	if req.Token == "" {
		s = CommonReturnModel{
			Code:    "406",
			Message: "token为空...",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	if !utils.ValidateToken(token){
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	content := req.Content

	//处理md为html
	//markdown字符串转为ASCII html代码
	htmlByte := blackfriday.MarkdownCommon([]byte(content))
	//反转义实体如“& lt;”成为“<” 把byte转位strings
	htmlStr := html.UnescapeString(string(htmlByte))
	//正则匹配并替换
	re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
	htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)

	//构造struct
	//timestamp
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	//Time
	//时间解析 strconv.FormatInt(time.Now().Unix(),10) base：进位制（2 进制到 36 进制 这种格式不行
	//time, terr := time.Parse("2006-01-02 15:04:05", strconv.FormatInt(time.Now().Unix(),10))
	//时间解析
	time, terr := time.Parse("2006-01-02 15:04:05", tm.Format("2006-01-02 03:04:05"))
	if terr != nil {
		logrus.Error(terr)
	}
	//装配struct
	note := &model.Note{Name: title, Content: htmlStr, Title: title, Date: tm.Format("2006-01-02 03:04:05"), Time: time}
	//获取session
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note") //获取数据
	//存入mongo
	err = c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	s = CommonReturnModel{
		Code:    "200",
		Message: `上传成功`,
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return

}
