package service

import (
	cst "github.com/Anteoy/liongo/constant"
	base "github.com/Anteoy/liongo/service/impl/base"
	"github.com/Anteoy/liongo/service/impl/pnote"
	ifbase "github.com/Anteoy/liongo/service/interface/base"
	ifpnote "github.com/Anteoy/liongo/service/interface/pnote"
	"github.com/Anteoy/liongo/utils"
)

type BaseFactory struct{}

//Generate 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (baseFactory *BaseFactory) Generate(rootDir string) {
	//博客初始化处理 获得yamlmapdata
	yp := utils.YamlParser{}
	cst.YamlData = yp.Parse(rootDir)

	//加载posts文件下md文件到slice
	var dispose ifbase.Dispose = &base.ProcessGetPostsArticle{}
	dispose.Dispose(rootDir)
	//处理分类
	dispose = &base.ProcessGetClassifies{}
	dispose.Dispose(rootDir)
	//处理导航 这里示例github
	dispose = new(base.ProcessGetNavBarList)
	dispose.Dispose(rootDir)
	//处理index.html
	dispose = new(base.ProcessIndexPage)
	dispose.Dispose(rootDir)
	//处理blog.html
	dispose = new(base.ProcessBlogListPage)
	dispose.Dispose(rootDir)
	//ProcessEveryArticlePage 按日期生成每一个文章html
	dispose = new(base.ProcessEveryArticlePage)
	dispose.Dispose(rootDir)
	//ProcessEveryArticlePage 生成按日期归档页面 /archive.html(Tag Date)
	dispose = new(base.ProcessArchiveDatePage)
	dispose.Dispose(rootDir)
	//ProcessUserDefinedPages 根据pages.yml生成用户自定义的html 需要用户在resources/pages下提供和pages.yaml配置文件id相同的页面内容.md文件
	dispose = new(base.ProcessUserDefinedPages)
	dispose.Dispose(rootDir)
	//生成分类的html /classify.html
	dispose = new(base.ProcessClassifyPage)
	dispose.Dispose(rootDir)

	//pNote初始化处理
	//生成ProcessPnoteloginPage.html
	dispose = new(base.ProcessPnoteloginPage)
	dispose.Dispose(rootDir)
	var disposePnote ifpnote.DisposePnote = &pnote.PreProcessNotes{}
	disposePnote.DisposePnote("str")
	//根据pre获取的notes进行生成pnotelist.html操作
	disposePnote = new(pnote.GeneratorPnotelist)
	disposePnote.DisposePnote(rootDir)
	pNoteService := new(pnote.PNoteService)
	//pNoteService.DealNoteUpload(ss)
	//从mgo中搜集并生成所有notes 单独html文件
	pNoteService.GetNotesFromMongo(cst.YamlData, nil, nil)

}

type SearchFactory struct{}

func (baseFactory *SearchFactory) Generate(rootDir string) {
	//博客初始化处理 获得yamlmapdata
	yp := utils.YamlParser{}
	cst.YamlData = yp.Parse(rootDir)

	//加载posts文件下md文件到slice
	var dispose ifbase.Dispose = &base.ProcessGetPostsArticle{}
	dispose.Dispose(rootDir)
}
