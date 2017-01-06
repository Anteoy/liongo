package build

import (
	"bufio"
	"github.com/Anteoy/blackfriday"
	"github.com/Anteoy/go-gypsy/yaml"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
    	"regexp"
	"main/go/utils"
	"fmt"
)

type BaseFactory struct{}
//
type Artic []*ArticleConfig

const (
	INDEX_TPL = "welcome"
	BLOG_LIST_TPL = "index"
	POSTS_TPL    = "posts"
	PAGES_TPL    = "pages"
	ARCHIVE_TPL  = "archive"
)

const (
	POST_DIR     = "posts"
	PUBLISH = "publish"
)

type MonthArchives []*MonthArchive
//用于sort
func (m MonthArchives) Len() int {
	return len(m)
}

func (m MonthArchives) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MonthArchives) Less(i, j int) bool {
	return m[i].month > m[j].month
}

const (
	COMMON_HEADER_FILE = "header.tpl"
	COMMON_FOOTER_FILE = "footer.tpl"
	
)
type RssConfig struct {
	Title  string
	Link   string
	Author string
	Date   string
	Desc   string
}

type CustomPage struct {
	Id      string
	Title   string
	Content string
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
	articles        Artic
	articleListSize int = 5000
	navBarList      []NavConfig
	categories      map[string]Category
	pages           []*CustomPage
	archives        map[string]*YearArchive
	allArchive      YearArchives

	NEWLY_ARTICLES_COUNT = 6
	INDEX_ARTICLES_SHOW_COUNT = 15
)

type YearArchive struct {
	Year   string
	Months []*MonthArchive
	months map[string]*MonthArchive
}

type MonthArchive struct {
	Month    string
	month    time.Month
	Articles []*ArticleBase
}

//传入路径和配置信息 返回一个template config.yml tpl 主题所使用的或自定义tmp名 融合footer header body为一个tpl
func parseTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//默认default
	themeFolder, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml at the theme value!")
		os.Exit(1)
	}

	file := root + "templates/" + themeFolder + "/" + tpl + ".tpl"
	if !isExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	log.Println("test")
	log.Println(cfg.Get)
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": unescaped})

	headerTpl := root + "templates/" + themeFolder + "/common/" + COMMON_HEADER_FILE
	footerTpl := root + "templates/" + themeFolder + "/common/" + COMMON_FOOTER_FILE

	if !isExists(headerTpl) {
		log.Println(headerTpl + " can not be found!")
		os.Exit(1)
	}

	if !isExists(footerTpl) {
		log.Println(footerTpl + " can not be found!")
		os.Exit(1)
	}

	t, err := t.ParseFiles(file, headerTpl, footerTpl)
	if err != nil {
		log.Println("parse " + tpl + " Template error!" + err.Error())
		os.Exit(1)
	}

	log.Println("parse " + tpl + " Template complete!")
	return t
}

func (baseFactory *BaseFactory) RenderIndex (root string,yamls map[string]interface{}) error {
	if !strings.HasSuffix(root,"/"){
		root += "/"
	}
	pEiz := yamls["config.yml"]
	var cfg = pEiz.(*yaml.File)
	t := parseTemplate(root,INDEX_TPL,cfg)

	targetFile := PUBLISH + "/index.html"
	fout,err :=os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	m := map[string]interface{}{"nav": navBarList,"cats": categories}
	exErr := t.Execute(fout, m)
	return exErr
}

func isExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//渲染生成index文件
func (baseFactory *BaseFactory) RenderBlogList(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := parseTemplate(root, BLOG_LIST_TPL, cfg)

	targetFile := PUBLISH + "/blog.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	if len(articles)<INDEX_ARTICLES_SHOW_COUNT{
		INDEX_ARTICLES_SHOW_COUNT = len(articles)
	}
	if len(articles)<NEWLY_ARTICLES_COUNT{
		NEWLY_ARTICLES_COUNT = len(articles)
	}

	m := map[string]interface{}{"ar": articles[:INDEX_ARTICLES_SHOW_COUNT], "nav": navBarList,"cats": categories,"newly":articles[:NEWLY_ARTICLES_COUNT]}
	exErr := t.Execute(fout, m)
	return exErr
}

//pre process posts pages
func (baseFactory *BaseFactory) PreProcessPosts(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	//获取config.yml键值对节点信息
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	//存放article的常量数组 固定size 5000
	articles = make([]*ArticleConfig, 0, articleListSize)
	//读取posts下article开始
	// returns
	// a list of directory entries sorted by filename.
	//获取文件夹下文件信息数组 return []FileInfo
	fileInfos, err := ioutil.ReadDir(root + POST_DIR)
	if err != nil {
		log.Println(err)
	}
	//对所有md文件进行遍历处理
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			log.Println("begin process article -- " + fileInfo.Name())
			fileName := fileInfo.Name()
			//获取posts文件夹下md文件信息 原始markdownstring config
			mardownStr, fi, err := processArticleFile(root+POST_DIR+"/"+fileName, fileName)
			if err != nil {
				log.Println("preprocess article file error!")
				os.Exit(1)
			}
			//去掉文件.md后缀
			trName := strings.TrimSuffix(fileName, ".md")
			//根据.md中配置生成年月日文件路径字符串
          		p := processArticleUrl(fi)
			log.Println(p)
			//markdown字符串转为ASCII html代码
			htmlByte := blackfriday.MarkdownCommon([]byte(mardownStr))
			//反转义实体如“& lt;”成为“<” 把byte转位strings
			htmlStr := html.UnescapeString(string(htmlByte))
		        re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
		    	htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
			//增加正文和链接
			fi.Content = htmlStr
			fi.Link = p + trName + ".html"
			//if abstract is empty,auto gen it
			if fi.Abstract == "" {
				var limit int = 1000
				rs := []rune(htmlStr)
				if len(rs) < 1000 {
					limit = len(rs)
				}

				abstract := utils.SubStr(htmlStr, 0, limit)
				fi.Abstract = utils.TrimHTML(abstract)
			}
			if fi.Author == "" {
				author, cerr := cfg.Get("meta.author")
				if cerr != nil {
					log.Println(cerr)
				}
				fi.Author = author
			}
			//添加article到articles 并对此进行排序 传入前面获取的fi(ArticleConfig)
			addAndSortArticles(fi)



		}
	}
	//生成自定义多余页面导航条 存入navBarList 数组
	generateNavBar(yamls)
	return nil
}

//生成日期归档静态文件
func (baseFactory *BaseFactory) RenderPosts(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	log.Println(cfg.Get("title"))
	t := parseTemplate(root, POSTS_TPL, cfg)
	for _,fileInfo := range articles {
		//create dir /yyyy/MM/dd
		p := processArticleUrl(*fileInfo)
		if !isExists(PUBLISH + "/articles/" + p) {
			os.MkdirAll(PUBLISH +"/articles/"+p, 0777)
		}
		targetFile := PUBLISH + "/articles/" + fileInfo.Link
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		m := map[string]interface{}{"fi": fileInfo,"nav": navBarList, "cats": categories,"newly":articles[:NEWLY_ARTICLES_COUNT-1]}
		t.Execute(fout, m)
	}


	return nil
}

//渲染生成archives分类静态文件
func (baseFactory *BaseFactory) RenderArchives(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)

	t := parseTemplate(root, ARCHIVE_TPL, cfg)
	targetFile := PUBLISH + "/archive.html"
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()

	generateArchive()
	//log.Println(allArchive)
	m := map[string]interface{}{"archives": allArchive, "nav": navBarList,"cats": categories,"newly":articles[:NEWLY_ARTICLES_COUNT-1]}
	exErr := t.Execute(fout, m)
	return exErr

}

//渲染pages下自定义导航 默认pages/about.md
func (baseFactory *BaseFactory) RenderPages(root string, yamls map[string]interface{}) error {
	//判断结尾是否/
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	//获取配置文件页面信息
	getPagesInfo(yamls)

	t := parseTemplate(root, PAGES_TPL, cfg)

	//pages为前面解析好的CustomPage数组
	for _, p := range pages {
		p.Id = strings.TrimSuffix(p.Id, " ")
		filePath := root + "pages/" + p.Id + ".md"
		if !isExists(filePath) {
			log.Println(filePath + " is not found!")
			os.Exit(1)
		}
		f, err := os.Open(filePath)
		if err != nil {
			log.Println(err)

		}
		defer f.Close()
		rd := bufio.NewReader(f)
		var markdownStr string
		//以行读取md文件
		for {
			buf, _, err := rd.ReadLine()

			if err == io.EOF {
				break
			} else {
				content := string(buf)
				markdownStr += content + "\n"
			}

		}

		//转化位二进制html
		htmlByte := blackfriday.MarkdownCommon([]byte(markdownStr))
		//转化位htmlstrings
		htmlStr := html.UnescapeString(string(htmlByte))
		//-1无限制完全替换
		htmlStr = strings.Replace(htmlStr, "<pre><code", `<pre class="prettyprint linenums"`, -1)
		htmlStr = strings.Replace(htmlStr, `</code>`, "", -1)

		p.Content = htmlStr//设置markdown文章内容
		log.Println(p.Content)
		if !isExists(PUBLISH + "/pages/") {
			os.MkdirAll(PUBLISH +"/pages/", 0777)
		}
		targetFile := PUBLISH + "/pages/" + p.Id + ".html"
		//创建target html
		fout, err := os.Create(targetFile)
		if err != nil {
			log.Println("create file " + targetFile + " error!")
			os.Exit(1)
		}
		defer fout.Close()
		//p .md article信息 nav 自定义的额外导航条信息 暂移除 "cats": categories "newly":articles[:NEWLY_ARTICLES_COUNT-1]
		m := map[string]interface{}{"p": p, "nav": navBarList}
		t.Execute(fout, m)
	}

	return nil
}

//获取配置文件页面信息
func getPagesInfo(yamls map[string]interface{}) {
	yCfg := yamls["pages.yml"]
	var cfg = yCfg.(*yaml.File)
	//统计配置pages.yml中配置个数
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		//strconv.Itoa转换为字符串拼接
		//获取配置的id和title
		id, erri := cfg.Get("[" + strconv.Itoa(i) + "].id")
		log.Println("[" + strconv.Itoa(i) + "].id")
		log.Println(id)
		if nil != erri {
			log.Println(erri)
		}

		title, errt := cfg.Get("[" + strconv.Itoa(i) + "].title")
		if nil != errt {
			log.Println(errt)
		}
		log.Println("[" + strconv.Itoa(i) + "].title")
		log.Println(title)
		page := CustomPage{id, title, ""}
		//追加到pages
		pages = append(pages, &page)
	}

}
//生成分类
func generateArchive() {
	archives = make(map[string]*YearArchive)
	for _, ar := range articles {
		y, m, _ := ar.Time.Date()
		year := fmt.Sprintf("%v", y)
		month := m.String()
		yArchive := archives[year]
		if yArchive == nil {
			yArchive = &YearArchive{year, make([]*MonthArchive, 0), make(map[string]*MonthArchive)}
			archives[year] = yArchive
		}
		mArchive := yArchive.months[month]
		if mArchive == nil {
			mArchive = &MonthArchive{month, m, make([]*ArticleBase, 0)}
			yArchive.months[month] = mArchive
		}
		mArchive.Articles = append(mArchive.Articles, &ArticleBase{ar.Link, ar.Title})

	}
	allArchive = make(YearArchives, 0)
	//sort by time
	for _, yArchive := range archives {
		monthCollect := make(MonthArchives, 0)
		for _, mArchive := range yArchive.months {
			monthCollect = append(monthCollect, mArchive)
		}
		sort.Sort(monthCollect)
		yArchive.months = nil
		yArchive.Months = monthCollect
		allArchive = append(allArchive, yArchive)
	}
	sort.Sort(allArchive)
}
//ArticleConfig 解析获取的.md --- --- 中配置文件
//string 根据年月日生成article路径
func processArticleUrl(ar ArticleConfig) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}

type Tag struct {
	Name     string
	Articles []ArticleBase
	Length   int
}


type ArticleBase struct {
	Link  string
	Title string
}

type Category struct {
	Name     string
	Articles []ArticleBase
	Length   int
}
//获取posts文件夹下md文件信息 原始markdownstring config
//filePath posts下文件全路径 fileName 文件名
//process posts,get article title,post date 返回md正文，读取md配置组合，error
func processArticleFile(filePath, fileName string) (string, ArticleConfig, error) {
	//打开文件
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)

	}
	defer f.Close()
	rd := bufio.NewReader(f)
	//flag主要标识程序处理md文件中 --- --- 读取各格式文件问题
	var flag int = 0
	var yamlStr, markdownStr string
	for {
		buf, _, err := rd.ReadLine()
		if err == io.EOF {
			//读取完毕
			break
		} else {
			//使用行级对md文件进行标识
			content := string(buf)
			if content == "---" {
				flag++
			}
			if flag == 2 {
				if content != "---" {
					//获取article 正文markdownStr
					markdownStr += content + "\n"
				}
			} else {
				//获取article正文前配置信息
				yamlStr += content + "\n"
			}

		}

	}
	//把md中---配置部分交于yaml进行处理（md中配置也基于yaml） 去掉所有---\n
	config := yaml.Config(strings.Replace(yamlStr, "---\n", "", -1))
	//获取md中配置说明信息
	title, err := config.Get("title")
	date, err := config.Get("date")
	tagCount, err := config.Count("tags")
	if err != nil {
		log.Println(err)
	}
	//处理md中配置tag
	var tags []TagConfig
	//去掉后缀
	trName := strings.TrimSuffix(fileName, ".md")
	for i := 0; i < tagCount; i++ {
		tagName, err := config.Get("tags[" + strconv.Itoa(i) + "]")
		if err != nil {
			log.Println("generate Tags error " + err.Error())
		}
		tags = append(tags, TagConfig{tagName, title, trName + ".html"})
	}
	//获取配置中分类
	cat, err := config.Get("categories[0]")
	abstract, err := config.Get("abstract")
	author, err := config.Get("author")
	//获取article 时间
	t, terr := time.Parse("2006-01-02 15:04:05", date)
	if terr != nil {
		log.Println(terr)
	}
	//log.Println(t)

	shortDate := t.UTC().Format("Jan 2, 2006")

	arInfo := ArticleConfig{title, date,shortDate, cat, tags, abstract, author, t, "", "", navBarList}

	//log.Println(markdownStr)
	return markdownStr, arInfo, nil

}

type ArticleConfig struct {
	Title    string
	Date     string
	ShortDate string
	Category string
	Tags     []TagConfig
	Abstract string
	Author   string
	Time     time.Time
	Link     string
	Content  string
	Nav      []NavConfig
}

type TagConfig struct {
	Name         string
	ArticleTitle string
	ArticleLink  string
}
//导航 struct
type NavConfig struct {
	Name   string
	Href   string
	Target string
}

type ByDate struct{ Artic }
//sort.Sort() 入参需覆写提供如下方法
func (a Artic) Len() int            { return len(a) }
func (a Artic) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a.Artic[i].Time.After(a.Artic[j].Time) }
//添加article到articles 并对此进行排序
func addAndSortArticles(arInfo ArticleConfig) {
	//log.Println(len(articles))
	artLen := len(articles)
	if artLen < articleListSize {
		articles = append(articles, &arInfo)
	}
	log.Println(len(articles))
	sort.Sort(ByDate{articles})
}

//转义
func unescaped(str string) interface{} {
	re := regexp.MustCompile(`<pre class="prettyprint linenums">([\s\S]*?)</pre>`)
	str = re.ReplaceAllStringFunc(str,xmlEscapString)
	return template.HTML(str)

}
func xmlHeader(blank string) string {
	return blank + `<?xml version="1.0" encoding="utf-8"?>`
}
func xmlEscapString(str string) string{
	str = strings.Replace(str,`<pre class="prettyprint linenums">`,"@@PRE_BEGIN",-1)
	str = strings.Replace(str,`</pre>`,"@@PRE_END",-1)
	str = template.HTMLEscapeString(str)
	str = strings.Replace(str,"@@PRE_BEGIN",`<pre class="prettyprint linenums">`,-1)
	str = strings.Replace(str,"@@PRE_END",`</pre>`,-1)
	return str
}
//生成自定义多余页面导航条 存入navBarList 数组
func generateNavBar(yamls map[string]interface{}) {
	yCfg := yamls["nav.yml"]
	var cfg = yCfg.(*yaml.File)
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		name, errn := cfg.Get("[" + strconv.Itoa(i) + "].label")
		if nil != errn {
			log.Println(errn)
		}
		href, errh := cfg.Get("[" + strconv.Itoa(i) + "].href")
		if nil != errh {
			log.Println(errh)
		}
		target, errt := cfg.Get("[" + strconv.Itoa(i) + "].target")
		if nil != errt {
			log.Println(errt)
		}

		nav := NavConfig{name, href, target}
		navBarList = append(navBarList, nav)

	}
	log.Println(navBarList)
}

//root 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (baseFactory *BaseFactory) Render(root string) {
	yp := new(utils.YamlParser)
	yamlData := yp.Parse(root)
	baseFactory.PreProcessPosts(root,yamlData)
	baseFactory.RenderIndex(root,yamlData)
	baseFactory.RenderBlogList(root, yamlData)
	baseFactory.RenderPosts(root, yamlData)
	baseFactory.RenderArchives(root, yamlData)
	baseFactory.RenderPages(root, yamlData)//pages/about.md
}
