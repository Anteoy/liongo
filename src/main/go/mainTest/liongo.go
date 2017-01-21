package main

import (
	"flag"
	"fmt"
	"os"
	Build "../service"
	"net/http"
	"log"

	"../newPosts"
	"strings"
	"main/go/controller"
)

const VERSION = "0.0.1"

const (
	USAGE = `
liongo is a static site generator in Go

Usage:

        liongo command [args...]

The commands are:

	build	        			build and generate site.
	version         			print liongo version

`
)

var httpAddr = ":8080"

func main() {
	flag.Parse()
	args := flag.Args()
	argsLength := len(args)
	fmt.Println(argsLength)
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
		Build.Build()
		if argsLength == 2 {
			httpAddr = args[1]
		}
		if argsLength == 3 && strings.EqualFold(args[1], "-p") {
			httpAddr = ":"+args[2]
		}
		if argsLength == 2 && strings.EqualFold(args[1],"--note"){
			fmt.Println("starting run with note !!!")
			pNoteController:=new(controller.PNoteController)
			http.HandleFunc("/login", pNoteController.Login)
			//http.HandleFunc("/lionnote", func() {//TODO
			//
			//})
		}
		fmt.Println("Listen at ", httpAddr)
		http.Handle("/", http.FileServer(http.Dir("../views/serve")))
		err := http.ListenAndServe(":8080",nil)// TODO httpAddr
		if err!=nil{
			log.Fatal("Start error",err)
		}
	case "new":
		args2 := args[1]
		//如果第二个参数为空 则直接返回并输出提示信息
		if args2 == "" && len(args2)==0 {
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
