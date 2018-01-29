package controller

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/Anteoy/liongo/utils"
	"os"
	"fmt"
	"bufio"
	"github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/service"
)

type ReqUploadBlog struct {
	Title string `json:"title"`
	Token string `json:"token"`
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
	s := CommonReturnModel{
	}

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
	name := constant.RENDER_DIR +"/"+ constant.POST_DIR+"/"+title + ".md"
	outputFile, outputError := os.OpenFile( name, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	rs,err := outputWriter.WriteString(content)
	fmt.Printf("WriteString res:%d,%v\n",rs,err)
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

type ReqDeleteBlog struct{
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
	s := CommonReturnModel{
	}

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
	if !utils.ValidateToken(token){
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	name := constant.RENDER_DIR +"/"+ constant.POST_DIR+"/"+req.Title + ".md"
	err = os.Remove(name)               //删除文件test.txt
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
	s := CommonReturnModel{
	}

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
	if !utils.ValidateToken(token){
		s = CommonReturnModel{
			Code:    "403",
			Message: "无效token..",
		}
		b, _ := json.Marshal(s)
		w.Write(b)
		return
	}
	name := constant.RENDER_DIR +"/"+ constant.POST_DIR+"/"+req.Title + ".md"
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
		Data: string(buf),
	}
	b, _ := json.Marshal(s)
	w.Write(b)
	return
}