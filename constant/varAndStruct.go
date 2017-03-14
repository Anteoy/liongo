package constant

import "time"

type Artic []*ArticleConfig

type MonthArchives []*MonthArchive

//用于sort
func (m MonthArchives) Len() int {
	return len(m)
}

func (m MonthArchives) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MonthArchives) Less(i, j int) bool {
	return m[i].MonthEn > m[j].MonthEn
}


//定制新增页面page内容 存储定制page的页面信息 用md转换为页面
type CustomPage struct {
	Id      string //md id
	Title   string //md article title
	Content string //md content
}

type YearArchives []*YearArchive

//用于sort
func (y YearArchives) Len() int {
	return len(y)
}

func (y YearArchives) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}

func (y YearArchives) Less(i, j int) bool {
	return y[i].Year > y[j].Year
}

var (
	Articles        Artic                          //完整信息
	ArticleListSize int                     = 5000 //博文不能超过5000
	NavBarList      []NavConfig                    //导航条数组
	Classifies      map[string]Category            //分类map
	Pages           []*CustomPage                  //新增定制页面数组 包含页面id.md md title 和md content
	YearArchivemap  map[string]*YearArchive        //key year value *YearArchive
	AllArchive      YearArchives                   //[]*YearArchive
	YamlData map[string]interface{}		       //Yaml 数据map

)

// Year 年 Months []*MonthArchive
type YearArchive struct {
	Year             string
	Months           []*MonthArchive
	MonthsArchiveMap map[string]*MonthArchive //months
}

// MonthArchive
type MonthArchive struct {
	Month    string
	MonthEn  time.Month
	Articles []*ArticleBase
}

//基础标签结构体
//结构体标识名 []ArticleBase数组 []ArticleBase长度
type Tag struct {
	Name     string
	Articles []ArticleBase
	Length   int
}

//基础博客结构体 用于分类和标签结构体组装
//链接地址 标题
type ArticleBase struct {
	Link  string
	Title string
}

//分类结构体
//分类名 ArticleBase数组 数组length
type Category struct {
	Name     string
	Articles []ArticleBase
	Length   int
}

//预处理完整article结构体
type ArticleConfig struct {
	Title     string      //标题
	Date      string      //时间
	ShortDate string      //简短时间 不用精确比较
	Category  string      //所属分类
	Tags      []TagConfig //所属标签
	Abstract  string      //摘要
	Author    string      //作者
	Time      time.Time   //精确时间
	Link      string      //博客链接
	Content   string      //完整内容
	Nav       []NavConfig
}

type TagConfig struct {
	Name         string
	ArticleTitle string
	ArticleLink  string
}

//导航 struct
//导航标识名 链接 a标签target属性
type NavConfig struct {
	Name   string
	Href   string
	Target string
}

type ByDate struct {
	Artic
}

//sort.Sort() 入参需覆写提供如下方法
func (a Artic) Len() int      { return len(a) }
func (a Artic) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// the time instant t is after u 意思为 i的时间是否在j的后面
func (a ByDate) Less(i, j int) bool { return a.Artic[i].Time.After(a.Artic[j].Time) }

