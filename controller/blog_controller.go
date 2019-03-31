package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Anteoy/go-gypsy/yaml"
	"github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/service"
	"github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/utils"
	"github.com/Anteoy/liongo/utils/logrus"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ReqUploadBlog struct {
	Title   string `json:"title"`
	Token   string `json:"token"`
	Content string `json:"content"`
}

func (pNoteController *PNoteController) UploadBlog(w http.ResponseWriter, r *http.Request) {

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
	s := CommonReturnModel{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := RPNCommitReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		s = CommonReturnModel{
			Code:    "407",
			Message: "参数不正确...",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
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
	if !utils.ValidateToken(token) {
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	content := req.Content
	name := constant.RENDER_DIR + "/" + constant.POST_DIR + "/" + title + ".md"
	outputFile, outputError := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	rs, err := outputWriter.WriteString(content)
	fmt.Printf("WriteString res:%d,%v\n", rs, err)
	err = outputWriter.Flush()
	if err != nil {
		s = CommonReturnModel{
			Code:    "501",
			Message: "处理文件失败...",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	service.Build()
	s = CommonReturnModel{
		Code:    "200",
		Message: `上传成功`,
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return

}

type ReqDeleteBlog struct {
	Title string `json:"title"`
	Token string `json:"token"`
}

func (pNoteController *PNoteController) DeleteBlog(w http.ResponseWriter, r *http.Request) {

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
	s := CommonReturnModel{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := ReqDeleteBlog{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		s = CommonReturnModel{
			Code:    "407",
			Message: "参数不正确...",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
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
	if !utils.ValidateToken(token) {
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	name := constant.RENDER_DIR + "/" + constant.POST_DIR + "/" + req.Title + ".md"
	err = os.Remove(name) //删除文件test.txt
	if err != nil {
		fmt.Println("file remove Error!")
		fmt.Printf("%s", err)
		s = CommonReturnModel{
			Code:    "508",
			Message: "operate err.." + err.Error(),
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	} else {
		fmt.Print("file remove OK!")
	}

	service.Build()
	s = CommonReturnModel{
		Code:    "200",
		Message: `删除成功`,
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return

}

type ReqGetBolg struct {
	Title string `json:"title"`
	Token string `json:"token"`
}

func (pNoteController *PNoteController) GetBlog(w http.ResponseWriter, r *http.Request) {

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
	s := CommonReturnModel{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := ReqGetBolg{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		s = CommonReturnModel{
			Code:    "407",
			Message: "参数不正确...",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
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
	if !utils.ValidateToken(token) {
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	name := constant.RENDER_DIR + "/" + constant.POST_DIR + "/" + req.Title + ".md"
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		fmt.Printf("%s", err)
		s = CommonReturnModel{
			Code:    "508",
			Message: "operate ioutil ReadFile err.." + err.Error(),
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	fmt.Print("file read OK!")

	s = CommonReturnModel{
		Code:    "200",
		Message: `获取成功`,
		Data:    string(buf),
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return
}

func (p *PNoteController) GetSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchKey := r.Form["search"]
	if searchKey == nil {
		io.WriteString(w, "参数search为空！！！")
		return
	}
	queryStr := "{\n" +
		"  \"query\": {\n" +
		"    \"dis_max\": {\n" +
		"      \"queries\": [\n" +
		"        {\n" +
		"          \"match\": {\n" +
		"            \"title\": {\n" +
		"              \"query\": \"芝麻\",\n" +
		"              \"minimum_should_match\": \"50%\",\n" +
		"              \"boost\": 4\n" +
		"            }\n" +
		"          }\n" +
		"        },\n" +
		"        {\n" +
		"          \"match\": {\n" +
		"            \"content\": {\n" +
		"              \"query\": \"芝麻\",\n" +
		"              \"minimum_should_match\": \"75%\",\n" +
		"              \"boost\": 4\n" +
		"            }\n" +
		"          }\n" +
		"        },\n" +
		"        {\n" +
		"          \"match\": {\n" +
		"            \"author\": {\n" +
		"              \"query\": \"芝麻\",\n" +
		"              \"minimum_should_match\": \"100%\",\n" +
		"              \"boost\": 2\n" +
		"            }\n" +
		"          }\n" +
		"        }\n" +
		"      ],\n" +
		"      \"tie_breaker\": 0.3\n" +
		"    }\n" +
		"  },\n" +
		"  \"highlight\" : {\n" +
		"            \"fields\" : {\n" +
		"                \"title\" : {},\n" +
		"                \"content\": {}\n" +
		"            }\n" +
		"        }\n" +
		"}"
	queryStr = strings.Replace(queryStr, "芝麻", searchKey[0], -1)
	println(queryStr)
	url := "http://" + SEARCH_URL + "/articles/article/_search"
	payload := strings.NewReader(queryStr)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
	js, _ := simplejson.NewJson(body) //反序列化
	println(js.Get("hits").Get("total"))
	println(js.Get("hits").Get("total"))
	total, _ := js.Get("hits").Get("total").Int()
	println(total)
	datas, _ := js.Get("hits").Get("hits").Array()
	var Articlesl constant.ArticleConfigsl
	for i, _ := range datas {
		article := &constant.ArticleConfig{}
		artjs := js.Get("hits").Get("hits").GetIndex(i)
		article.Title, _ = artjs.Get("_source").Get("title").String()
		article.Date, _ = artjs.Get("_source").Get("date").String()
		article.Classify, _ = artjs.Get("_source").Get("classify").String()
		article.Abstract, _ = artjs.Get("_source").Get("abstract").String()
		article.Author, _ = artjs.Get("_source").Get("author").String()
		article.Link, _ = artjs.Get("_source").Get("link").String()
		article.Content, _ = artjs.Get("_source").Get("content").String()
		article.Id, _ = artjs.Get("_source").Get("id").String()
		Articlesl = append(Articlesl, article)
		println(article.Id, article.Title, article.Content)
	}
	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	t := ParseTemplate("../resources/", SEARCH_TPL, cfg)
	targetFile := PUBLISH_DIR + "/search.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		logrus.Error("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	m := map[string]interface{}{"ar": Articlesl[:], "nav": NavBarsl, "cats": Classifiesm, "search": searchKey[0], "total": total}
	exErr := t.Execute(fout, m)
	if exErr != nil {
		log.Fatal(exErr)
	}
	http.Redirect(w, r, "/search.html", http.StatusFound)
}
