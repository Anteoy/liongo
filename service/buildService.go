package service

import (
	"log"
	"os"

	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/utils"
	"github.com/Anteoy/liongo/utils/logrus"
)

func Build() {
	//../views/serve
	if utils.IsExists(PUBLISH_DIR) {
		//创建777权限目录
		err := os.RemoveAll(PUBLISH_DIR)
		if err != nil {
			log.Panic("remove all publish dir error -- " + err.Error())
		}
	}
	//创建777权限目录
	err := os.MkdirAll(PUBLISH_DIR, 0777)
	if err != nil {
		log.Panic("create publish dir error -- " + err.Error())
	}
	//开始生成渲染文件
	var rf = new(BaseFactory)
	rf.Generate(RENDER_DIR)
	//cp resources
	err = utils.CopyDir(RENDER_DIR+"/prettify", PUBLISH_DIR+"/prettify")
	err = utils.CopyDir(RENDER_DIR+"/js", PUBLISH_DIR+"/js")
	err = utils.CopyDir(RENDER_DIR+"/css", PUBLISH_DIR+"/css")
	if err != nil {
		logrus.Error(err)
	}
	//复制网站图标自定义文件
	err = utils.CopyDir(RENDER_DIR+"/images/icon", PUBLISH_DIR)
	//复制网站images
	err = utils.CopyDir(RENDER_DIR+"/images", PUBLISH_DIR+"/images")
	err = utils.CopyDir(RENDER_DIR+"/css", PUBLISH_DIR+"/css")
	//复制pnote upload commit.html
	err = utils.CopyDir(RENDER_DIR+"/html", PUBLISH_DIR+"/protohtml")
	logrus.Debug("blog process ok！")
}
