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
)

type RenderFactory struct{}
//
type Artic []*ArticleConfig

const (
	INDEX_TPL    = "index"
	TAG_TPL      = "tag"
	POSTS_TPL    = "posts"
	PAGES_TPL    = "pages"
	RSS_TPL      = "rss"
	CATEGORY_TPL = "category"
	ARCHIVE_TPL  = "archive"
)

const (
	POST_DIR     = "posts"
	PUBLICSH_DIR = "publish"
)

type MonthArchives []*MonthArchive

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
var (
	articles        Artic
	articleListSize int = 5000
	rss             []RssConfig
	rssListSize     int = 10
	navBarList      []NavConfig
	allTags         map[string]Tag
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

func parseTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//get theme template
	themeFold, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml")
		os.Exit(1)
	}

	file := root + "templates/" + themeFold + "/" + tpl + ".tpl"
	if !isExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": unescaped})

	headerTpl := root + "templates/" + themeFold + "/common/" + COMMON_HEADER_FILE
	footerTpl := root + "templates/" + themeFold + "/common/" + COMMON_FOOTER_FILE

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

func parseXMLTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//get theme template
	themeFold, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml")
		os.Exit(1)
	}
	file := root + "/templates/" + themeFold + "/" + tpl + ".tpl"
	if !isExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	t.Funcs(template.FuncMap{"xmlheader": xmlHeader})

	t, err := t.ParseFiles(file)
	if err != nil {
		log.Println("parse " + tpl + " Template error!" + err.Error())
		os.Exit(1)
	}

	log.Println("parse " + tpl + " Template complete!")
	return t
}

func isExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//渲染生成index文件
func (self *RenderFactory) RenderIndex(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := parseTemplate(root, INDEX_TPL, cfg)

	targetFile := PUBLICSH_DIR + "/index.html"
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
func (self *RenderFactory) PreProcessPosts(root string, yamls map[string]interface{}) error {
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
	//获取文件夹下文件信息数组
	fileInfos, err := ioutil.ReadDir(root + POST_DIR)
	if err != nil {
		log.Println(err)
	}
	//对所有md文件进行遍历处理
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			log.Println("begin process article -- " + fileInfo.Name())
			fileName := fileInfo.Name()
			//change markdown
			mardownStr, fi, err := processArticleFile(root+POST_DIR+"/"+fileName, fileName)
			//create post html file
			if err != nil {
				log.Println("preprocess article file error!")
				os.Exit(1)
			}
			trName := strings.TrimSuffix(fileName, ".md")
          		p := processArticleUrl(fi)
			//deal markdown markdown字符串转为ASCII 转为html代码
			htmlByte := blackfriday.MarkdownCommon([]byte(mardownStr))
			//init other article infos
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
			//sort by date
			addAndSortArticles(fi)



		}
	}
	generateCategories()
	generateTags()
	generateNavBar(yamls)
	return nil
}

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

func generateTags() {
	allTags = make(map[string]Tag)
	for _, ar := range articles {
		for _, tg := range ar.Tags {
			//log.Println(tg)
			t, ok := allTags[tg.Name]
			if ok {
				t.Articles = append(t.Articles, ArticleBase{ar.Link, ar.Title})
				t.Length = len(t.Articles)
				allTags[tg.Name] = t
			} else {
				art := ArticleBase{ar.Link, ar.Title}
				arts := make([]ArticleBase, 0)
				arts = append(arts, art)
				allTags[tg.Name] = Tag{tg.Name, arts, 1}
			}
		}
	}
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

func generateCategories() {
	categories = make(map[string]Category)
	for _, ar := range articles {
		c, ok := categories[ar.Category]
		if ok {
			c.Articles = append(c.Articles, ArticleBase{ar.Link, ar.Title})
			c.Length = len(c.Articles)
			categories[ar.Category] = c
		} else {
			art := ArticleBase{ar.Link, ar.Title}
			arts := make([]ArticleBase, 0)
			arts = append(arts, art)
			categories[ar.Category] = Category{ar.Category, arts, 1}
		}
	}
}

//process posts,get article title,post date 返回md正文，读取md配置组合，error
func processArticleFile(filePath, fileName string) (string, ArticleConfig, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)

	}
	defer f.Close()
	rd := bufio.NewReader(f)
	//对article计数
	var ct int = 0
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
				ct++
			}
			if ct == 2 {
				if content != "---" {
					markdownStr += content + "\n"
				}
			} else {
				yamlStr += content + "\n"
			}

		}

	}//把md中---配置部分交于yaml进行处理
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
//sort articles by date
func addAndSortArticles(arInfo ArticleConfig) {
	//log.Println(len(articles))
	artLen := len(articles)
	if artLen < articleListSize {
		articles = append(articles, &arInfo)
	}
	log.Println(len(articles))
	sort.Sort(ByDate{articles})
}

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
	//log.Println(navBarList)
}

//root 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (self *RenderFactory) Render(root string) {
	yp := new(utils.YamlParser)
	yamlData := yp.Parse(root)
	self.PreProcessPosts(root,yamlData)
	self.RenderIndex(root, yamlData)
}
