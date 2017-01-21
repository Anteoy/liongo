package controller

import (
	"io"
	"net/http"
	"fmt"
	"github.com/Anteoy/liongo/src/main/go/dao/mysql"
	"github.com/Anteoy/liongo/src/main/go/modle"
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
		fmt.Fprint(w,"<h1>login success!!!</h1>")
		//http.ServeFile(w, r, "./static/html/index.html")
	} else {
		fmt.Fprint(w,"<h1>login faild!!!用户名或密码不正确！！！</h1>")
		//http.ServeFile(w, r, "./static/html/login.html")
	}

}

//获取笔记md文件并存入mongo

func (pNoteController *PNoteController) DataTomongo(notemd *modle.Note){

}