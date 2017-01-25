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
		http.ServeFile(w, r, "./static/html/index.html")
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