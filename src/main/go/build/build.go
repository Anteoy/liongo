package build

import (
	"log"
	"main/go/utils"
	"os"
)
const (
	RENDER_DIR = "../resources"
)
func Build() {
	//publish
	if !utils.IsExists(PUBLICSH_DIR) {
		//创建777权限目录
		err := os.Mkdir(PUBLICSH_DIR, 0777)
		if err != nil {
			log.Panic("create publish dir error -- " + err.Error())
		}
	}
	//开始生成渲染文件
	var rf = new(RenderFactory)
	rf.Render(RENDER_DIR)
	//复制assets
	err := utils.CopyDir(RENDER_DIR+"/assets", PUBLICSH_DIR+"/assets")
	if err != nil {
		log.Println(err)
	}
	log.Println("blog process ok！")

}
