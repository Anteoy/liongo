package router

import (
	"github.com/Anteoy/liongo/controller"
	"net/http"
	"fmt"
)

func Router() {
	fmt.Println("starting api router !!!")
	pNoteController := new(controller.PNoteController)
	http.HandleFunc("/login", pNoteController.Login)
	http.HandleFunc("/notes", pNoteController.GetNote)
	//路由上传接口
	http.HandleFunc("/PNCommit", pNoteController.PNCommit)
	http.HandleFunc("/RPNCommit", pNoteController.RPNCommit)
}
