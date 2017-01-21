package service

import (
	"log"
	"../utils"
	"os"
)
const (
	RENDER_DIR = "../resources"
)
func Build() {
	//publish
	if !utils.IsExists(PUBLISH) {
		//创建777权限目录
		err := os.Mkdir(PUBLISH, 0777)
		if err != nil {
			log.Panic("create publish dir error -- " + err.Error())
		}
	}
	//开始生成渲染文件
	var rf = new(BaseFactory)
	rf.Generate(RENDER_DIR)
	//复制assets
	err := utils.CopyDir(RENDER_DIR+"/assets", PUBLISH +"/assets")
	if err != nil {
		log.Println(err)
	}
	//复制网站图标自定义文件
	err = utils.CopyDir(RENDER_DIR+"/images/icon", PUBLISH)
	//复制网站images
	err = utils.CopyDir(RENDER_DIR+"/images", PUBLISH+"/images")
	err = utils.CopyDir(RENDER_DIR+"/css", PUBLISH+"/css")
	log.Println("blog process ok！")

}
