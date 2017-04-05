package service

import (
	"github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/service/impl/base"
	. "github.com/Anteoy/liongo/service/impl/pnote"
	. "github.com/Anteoy/liongo/service/interface/pnote"
	. "github.com/Anteoy/liongo/service/interface/base"
)

type BaseFactory struct{}


//root 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (baseFactory *BaseFactory) Generate(rootDir string) {
	//博客初始化处理 获得yamlmapdata
	yp := utils.YamlParser{}
	YamlData = yp.Parse(rootDir)

	//加载posts文件下md文件到slice
	var dispose Dispose= &ProcessGetPostsArticle{}
	dispose.Dispose(rootDir)
	//处理分类
	dispose = &ProcessGetClassifies{}
	dispose.Dispose(rootDir)
	//处理导航 这里示例github
	dispose = new(ProcessGetNavBarList)
	dispose.Dispose(rootDir)
	//处理index.html
	dispose = new(ProcessIndexPage)
	dispose.Dispose(rootDir)
	//处理blog.html
	dispose = new(ProcessBlogListPage)
	dispose.Dispose(rootDir)
	//ProcessEveryArticlePage 按日期生成每一个文章html
	dispose = new(ProcessEveryArticlePage)
	dispose.Dispose(rootDir)
	//ProcessEveryArticlePage 生成按日期归档页面 /archive.html(Tag Date)
	dispose = new(ProcessArchiveDatePage)
	dispose.Dispose(rootDir)
	//ProcessUserDefinedPages 根据pages.yml生成用户自定义的html 需要用户在resources/pages下提供和pages.yaml配置文件id相同的页面内容.md文件
	dispose = new(ProcessUserDefinedPages)
	dispose.Dispose(rootDir)
	//生成分类的html /classify.html
	dispose = new(ProcessClassifyPage)
	dispose.Dispose(rootDir)

	//pNote初始化处理
	//生成ProcessPnoteloginPage.html
	dispose = new(ProcessPnoteloginPage)
	dispose.Dispose(rootDir)
	var disposePnote DisposePnote = &PreProcessNotes{}
	disposePnote.DisposePnote("str")
	//根据pre获取的notes进行生成pnotelist.html操作
	disposePnote = new(GeneratorPnotelist)
	disposePnote.DisposePnote(rootDir)
}
