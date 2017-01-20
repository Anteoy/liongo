package controller

import (
	"io"
	"net/http"
	"../dao/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		http.ServeFile(w, r, "./static/html/index.html")
	} else {
		http.ServeFile(w, r, "./static/html/login.html")
	}

}
