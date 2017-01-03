package build

import (
	"log"
	"main/go/utils"
	myconst "main/go/constant"
	"os"
)

func Build() {
	//publish
	if !utils.IsExists(myconst.PUBLICSH_DIR) {
		//创建777权限目录
		err := os.Mkdir(myconst.PUBLICSH_DIR, 0777)
		if err != nil {
			log.Panic("create publish dir error -- " + err.Error())
		}
	}
	//开始生成渲染文件
	var rf = new(RenderFactory)
	rf.Render(myconst.RENDER_DIR)
	//copy res
	err := utils.CopyDir(myconst.RENDER_DIR+"/assets", PUBLICSH_DIR+"/assets")
	if err != nil {
		log.Println(err)
	}
	log.Println("blog process ok！")

}
