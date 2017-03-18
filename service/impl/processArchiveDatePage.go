package impl

import (
	"strings"
	"os"
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	"log"
	"sort"
	"fmt"
)

type ProcessArchiveDatePage struct{}
//渲染生成archives时间归档静态文件 archive.html
func (processArchiveDatePage *ProcessArchiveDatePage)Dispose(dir string)  {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	yCfg := YamlData["config.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}

	t := ParseTemplate(dir, ARCHIVE_TPL, cfg)
	targetFile := PUBLISH_DIR + "/archive.html"
	//创建targetFile
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	//时间归档处理
	generateArchive()
	//log.Println(allArchive)
	m := map[string]interface{}{"archives": AllArchive, "nav": NavBarList,"cats": Classifiesm}
	t.Execute(fout, m)
}

//按照日期生成分类
func generateArchive() {
	YearArchivemap = make(map[string]*YearArchive)
	//读取所有articlesl 获取完整YearArchivemap MonthsArchiveMap
	for _, ar := range Articlesl {//排序好的ArticleConfig指针数组
		y, m, _ := ar.Time.Date()//获取当前的ArticleConfig year和month
		year := fmt.Sprintf("%v", y)
		month := m.String() // annotation // String returns the English name of the month ("January", "February", ...).
		yArchivestru := YearArchivemap[year]
		if yArchivestru == nil {//判断map中是否有此yearArchive日期分类
			yArchivestru = &YearArchive{year, make([]*MonthArchive, 0), make(map[string]*MonthArchive)}
			YearArchivemap[year] = yArchivestru//放入新的以年分类的Key
		}
		//以上面同样的规则处理mArchive 新建MonthArchive放入yArchivestru.MonthsArchiveMap[month]
		mArchivestru := yArchivestru.MonthsArchiveMap[month]
		if mArchivestru == nil {//是否存在月份小分类
			mArchivestru = &MonthArchive{month, m, make([]*ArticleBase, 0)}
			yArchivestru.MonthsArchiveMap[month] = mArchivestru//新建并赋值于yArchive，内层嵌套
		}
		//上面未填充的ArticleBase 这里进行补充植入
		mArchivestru.Articles = append(mArchivestru.Articles, &ArticleBase{ar.Link, ar.Title})//年月下嵌入此article

	}
	// make AllArchive 初始化
	AllArchive = make(YearArchivesl, 0)
	//月份排序
	for _, yArchive := range YearArchivemap {
		monthCollect := make(MonthArchivesl, 0)
		for _, mArchive := range yArchive.MonthsArchiveMap {//获取内部months
			monthCollect = append(monthCollect, mArchive)
		}
		sort.Sort(monthCollect)//月份排序
		yArchive.MonthsArchiveMap = nil //months map[string]*MonthArchive TODO
		//放入有序的monthCollect  到 yArchive.Months
		yArchive.Months = monthCollect
		//yArchive全部依次放入需要排序的AllArchive
		AllArchive = append(AllArchive, yArchive)//放入此年的yArchive到allArchive
	}
	//年份排序
	sort.Sort(AllArchive)//对year进行排序
}
