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
	//博客初始化处理
	yp := new(utils.YamlParser)
	YamlData = yp.Parse(rootDir)

	//ProcessPosts
	var rf Dispose= new(ProcessGetPostsArticle)
	rf.Dispose(rootDir)
	//ProcessClassifies
	rf = new(ProcessGetClassifies)
	rf.Dispose(rootDir)
	//ProcessAddNav
	rf = new(ProcessGetNavBarList)
	rf.Dispose(rootDir)
	//ProcessIndex
	rf = new(ProcessIndexPage)
	rf.Dispose(rootDir)
	//ProcessBlogList
	rf = new(ProcessBlogListPage)
	rf.Dispose(rootDir)
	//ProcessEveryArticlePage 按日期生成每一个文章html
	rf = new(ProcessEveryArticlePage)
	rf.Dispose(rootDir)
	//ProcessEveryArticlePage 生成按日期归档页面 /archive.html(Tag Date)
	rf = new(ProcessEveryArticlePage)
	rf.Dispose(rootDir)
	//ProcessUserDefinedPages 根据pages.yml生成用户自定义的html
	rf = new(ProcessUserDefinedPages)
	rf.Dispose(rootDir)
	//生成分类的html /classify.html
	rf = new(ProcessClassifyPage)
	rf.Dispose(rootDir)
	//生成ProcessPnoteloginPage.html
	rf = new(ProcessClassifyPage)
	rf.Dispose(rootDir)

	//pNote初始化处理
	var pNoteService = new(PNoteService)
	pNoteService.PreProcessNotes()
	pNoteService.GeneratorPnotelist(rootDir, YamlData)
}
