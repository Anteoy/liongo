package router

import (
	"fmt"
	"github.com/Anteoy/liongo/controller"
	"net/http"
)

func Router() {
	fmt.Println("starting api router !!!")
	pNoteController := new(controller.PNoteController)
	http.HandleFunc("/login", pNoteController.Login)
	http.HandleFunc("/loginR", pNoteController.LoginR)
	http.HandleFunc("/notes", pNoteController.GetNote)
	//路由上传接口
	http.HandleFunc("/PNCommit", pNoteController.PNCommit)
	http.HandleFunc("/RPNCommit", pNoteController.RPNCommit)
	http.HandleFunc("/upload", pNoteController.UploadBlog)
}
