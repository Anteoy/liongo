package service

import (
	"github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/service/impl"
	. "github.com/Anteoy/liongo/service/interface"
)

type BaseFactory struct{}


//root 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (baseFactory *BaseFactory) Generate(rootDir string) {
	//博客初始化处理 获得yamlmapdata
	yp := utils.YamlParser{}
	YamlData = yp.Parse(rootDir)

	//加载posts文件下md文件到slice
	var rf Dispose= &ProcessGetPostsArticle{}
	rf.Dispose(rootDir)
	//处理分类
	rf = &ProcessGetClassifies{}
	rf.Dispose(rootDir)
	//处理导航 这里示例github
	rf = new(ProcessGetNavBarList)
	rf.Dispose(rootDir)
	//处理index.html
	rf = new(ProcessIndexPage)
	rf.Dispose(rootDir)
	//处理blog.html
	rf = new(ProcessBlogListPage)
	rf.Dispose(rootDir)
	//ProcessEveryArticlePage 按日期生成每一个文章html
	rf = new(ProcessEveryArticlePage)
	rf.Dispose(rootDir)
	//ProcessEveryArticlePage 生成按日期归档页面 /archive.html(Tag Date)
	rf = new(ProcessArchiveDatePage)
	rf.Dispose(rootDir)
	//ProcessUserDefinedPages 根据pages.yml生成用户自定义的html 需要用户在resources/pages下提供和pages.yaml配置文件id相同的页面内容.md文件
	rf = new(ProcessUserDefinedPages)
	rf.Dispose(rootDir)
	//生成分类的html /classify.html
	rf = new(ProcessClassifyPage)
	rf.Dispose(rootDir)
	//生成ProcessPnoteloginPage.html
	rf = new(ProcessPnoteloginPage)
	rf.Dispose(rootDir)

	//pNote初始化处理
	var pNoteService = new(PNoteService)
	pNoteService.PreProcessNotes()
	pNoteService.GeneratorPnotelist(rootDir, YamlData)
}
