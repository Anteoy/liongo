package main

import (
	"flag"
	"fmt"
	cst "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/controller"
	"github.com/Anteoy/liongo/newPosts"
	Build "github.com/Anteoy/liongo/service"
	log "github.com/Anteoy/liongo/utils/logrus"
	"net/http"
	"os"
	"strings"
)

var httpPort = ":8080"

func main() {
	flag.Parse()
	args := flag.Args()
	argsLength := len(args)
	//判断输入命令长度
	if argsLength == 0 || argsLength > 3 {
		UseInfo()
		os.Exit(1)
	}
	//通过第一个参数进行识别
	switch args[0] {
	case "build":
		Build.Build()
	case "run":
		if argsLength == 3 && strings.EqualFold(args[1], "-p") {
			httpPort = ":" + args[2]
		}
		if argsLength == 2 && strings.EqualFold(args[1], "--note") {
			log.Debug("starting run with note !!!")
			pNoteController := new(controller.PNoteController)
			http.HandleFunc("/login", pNoteController.Login)
			http.HandleFunc("/notes", pNoteController.GetNote)
			//路由上传接口
			http.HandleFunc("/PNCommit", pNoteController.PNCommit)
			//http.HandleFunc("/lionnote", func() {//TODO
			//
			//})
		}
		Build.Build()
		log.Debug("Listen at ", httpPort)
		http.Handle("/", http.FileServer(http.Dir("../views/serve")))
		err := http.ListenAndServe(httpPort, nil)
		if err != nil {
			log.Fatal("Start error", err)
		}
	case "new":
		args2 := args[1]
		//如果第二个参数为空 则直接返回并输出提示信息
		if args2 == "" && len(args2) == 0 {
			UseInfo()
			os.Exit(1)
		}
		addFactory := new(newPosts.AddFactory)
		addFactory.New(args2)
	case "version":
		fmt.Print("liongo version " + cst.VERSION)
	default:
		UseInfo()
		os.Exit(1)
	}

}

func UseInfo() {
	fmt.Println(cst.USAGE)
}
