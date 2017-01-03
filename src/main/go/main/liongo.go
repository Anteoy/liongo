package main

import (
	"flag"
	"fmt"
	"os"
	Build "../build"
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
		Usage()
		os.Exit(1)
	}
	//通过第一个参数进行识别
	switch args[0] {
	default:
		Usage()
		os.Exit(1)
	case "build":
		Build.Build()
	case "run":
		if argsLength == 2 {
			httpAddr = args[1]
		}
		fmt.Println("Listen at ", httpAddr)
		//Build.run(httpAddr)
	case "version":
		fmt.Print("liongo version " + VERSION)
	}
}

func Usage() {
	fmt.Println(USAGE)
}
