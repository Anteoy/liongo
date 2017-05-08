package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/Anteoy/liongo/utils/logrus"

	. "github.com/Anteoy/liongo/constant"
	Build "github.com/Anteoy/liongo/service"

	"strings"

	"github.com/Anteoy/liongo/controller"
	"github.com/Anteoy/liongo/newPosts"
	"github.com/Anteoy/liongo/utils"
)

var httpPort = ":8080"

func main() {
	year := fmt.Sprintf("%v", "827827")
	fmt.Println(year)
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
		if argsLength == 2 {
			httpPort = args[1]
		}
		if argsLength == 3 && strings.EqualFold(args[1], "-p") {
			httpPort = ":" + args[2]
		}
		if argsLength == 2 && strings.EqualFold(args[1], "--note") {
			log.Debug("starting run with note !!!")
			pNoteService := new(Build.PNoteService)
			//pNoteService.DealNoteUpload(ss)
			yp := new(utils.YamlParser)
			yamlData := yp.Parse("../resources")
			//从mgo中搜集并生成所有notes 单独html文件
			pNoteService.GetNotesFromMongo(yamlData, nil, nil)
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
		err := http.ListenAndServe(":8080", nil) // TODO httpAddr
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
		fmt.Print("liongo version " + VERSION)
	default:
		UseInfo()
		os.Exit(1)
	}

}

func UseInfo() {
	fmt.Println(USAGE)
}
