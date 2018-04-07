package main

import (
	"flag"
	"fmt"
	cst "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/router"
	"github.com/Anteoy/liongo/service"
	log "github.com/Anteoy/liongo/utils/logrus"
	"net/http"
	_ "net/http/pprof"
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
		service.Build()
	case "run":
		if argsLength == 3 && strings.EqualFold(args[1], "-p") {
			httpPort = ":" + args[2]
		}
		if argsLength == 2 && strings.EqualFold(args[1], "--note") {
			log.Debug("starting run with note !!!")
			router.Router()
		}
		service.Build()
		log.Debug("Listen at ", httpPort)
		http.Handle("/", http.FileServer(http.Dir("../views/serve")))
		http.HandleFunc("/resources/images/site/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("r.URL.Path = %s,r.URL.Path[1:] = %s\n", r.URL.Path, r.URL.Path[1:])
			//相对路径 否则必须同级
			http.ServeFile(w, r, "../"+r.URL.Path[1:])
		})
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
		addFactory := new(service.AddFactory)
		addFactory.New(args2)
	case "version":
		fmt.Printf("liongo version '%s'\n", cst.VERSION)
	default:
		UseInfo()
		os.Exit(1)
	}

}

func UseInfo() {
	fmt.Println(cst.USAGE)
}
