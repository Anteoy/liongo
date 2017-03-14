package impl

import (
	"strings"
	"io/ioutil"
	"os"
	"regexp"
	"github.com/Anteoy/liongo/utils"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/service"
	"github.com/Anteoy/go-gypsy/yaml"
	"time"
	"log"
	"bufio"
	"io"
	"strconv"
	"html"
	"github.com/Anteoy/blackfriday"
	"sort"
)

type ProcessPosts struct {}

////导航 struct
////导航标识名 链接 a标签target属性
//type NavConfig struct {
//	Name   string
//	Href   string
//	Target string
//}
//
//
//type TagConfig struct {
//	Name         string
//	ArticleTitle string
//	ArticleLink  string
//}
//
////预处理完整article结构体
//type ArticleConfig struct {
//	Title     string      //标题
//	Date      string      //时间
//	ShortDate string      //简短时间 不用精确比较
//	Category  string      //所属分类
//	Tags      []TagConfig //所属标签
//	Abstract  string      //摘要
//	Author    string      //作者
//	Time      time.Time   //精确时间
//	Link      string      //博客链接
//	Content   string      //完整内容
//	Nav       []NavConfig
//}
//
////基础博客结构体 用于分类和标签结构体组装
////链接地址 标题
//type ArticleBase struct {
//	Link  string
//	Title string
//}
//
////分类结构体
////分类名 ArticleBase数组 数组length
//type Category struct {
//	Name     string
//	Articles []ArticleBase
//	Length   int
//}
//
//type ByDate struct {
//	Artic
//}
//
////sort.Sort() 入参需覆写提供如下方法
//func (a Artic) Len() int      { return len(a) }
//func (a Artic) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
//
//// the time instant t is after u 意思为 i的时间是否在j的后面
//func (a ByDate) Less(i, j int) bool { return a.Artic[i].Time.After(a.Artic[j].Time) }
//
//
//type Artic []*ArticleConfig
//
//
//
//var(
//	articles        Artic                          //完整信息
//	articleListSize int                     = 5000 //博文不能超过5000
//	navBarList      []NavConfig                    //导航条数组
//
//)


//加载posts文件夹下的md博文信息
func (processPosts *ProcessPosts)Dispose(dir string)  {

	yp := new(utils.YamlParser)
	yamls := yp.Parse(dir)

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	//获取config.yml键值对节点信息
	yCfg := yamls["config.yml"]

	//Comma-ok断言 不推荐下面第二种方式 下面第二种如果转换失败会直接panic
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//cfg = yCfg.((*yaml.File))
	//存放article的常量数组 固定size 5000
	articles = make([]*ArticleConfig, 0, articleListSize)
	//读取posts下article开始
	// returns
	// a list of directory entries sorted by filename.
	//获取文件夹下文件信息数组 return []FileInfo
	fileInfos, err := ioutil.ReadDir(dir + POST_DIR)
	if err != nil {
		log.Println(err)
	}
	//对所有md文件进行遍历处理
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			log.Println("begin process article -- " + fileInfo.Name())
			fileName := fileInfo.Name()
			//获取posts文件夹下md文件信息 原始markdownstring config fi = ArticleConfig
			mardownStr, fi, err := processArticleFile(dir +POST_DIR+"/"+fileName, fileName)
			if err != nil {
				log.Println("preprocess article file error!")
				os.Exit(1)
			}
			//去掉文件.md后缀
			trName := strings.TrimSuffix(fileName, ".md")
			//根据.md中配置生成年月日文件路径字符串 返回html前一级路径
			p := processArticleUrl(fi)
			log.Println(p)
			//markdown字符串转为ASCII html代码 []byte(mardownStr) string强转为[]byte
			htmlByte := blackfriday.MarkdownCommon([]byte(mardownStr))
			//反转义实体如“& lt;”成为“<” 把byte转位strings
			htmlStr := html.UnescapeString(string(htmlByte))
			re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
			htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
			//增加正文和链接 组装ArticleConfig fi
			fi.Content = htmlStr
			fi.Link = p + trName + ".html"
			//if abstract is empty,auto gen it
			if fi.Abstract == "" {
				var limit int = 1000
				rs := []rune(htmlStr)
				if len(rs) < 1000 {
					limit = len(rs)
				}
				//组装ArticleConfig摘要
				abstract := utils.SubStr(htmlStr, 0, limit)
				fi.Abstract = utils.TrimHTML(abstract)
			}
			if fi.Author == "" {
				//从配置文件获取author
				author, cerr := cfg.Get("meta.author")
				if cerr != nil {
					log.Println(cerr)
				}
				fi.Author = author
			}
			//添加article到articles 并对此进行排序 传入前面获取的fi(ArticleConfig) TODO 提高效率
			addAndSortArticles(fi)
		}
	}
	////分类预处理
	//generateClassifies()
	//生成自定义多余页面导航条 存入navBarList 数组
	////这里配置的示例github导航
	//generateNavBar(yamls)
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

	arInfo := ArticleConfig{title, date, shortDate, cat, tags, abstract, author, t, "", "", navBarList}

	//log.Println(markdownStr)
	return markdownStr, arInfo, nil

}

//ArticleConfig 解析获取的.md --- --- 中配置文件struct
//string 根据年月日生成article路径
func processArticleUrl(ar ArticleConfig) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}

//添加article到articles 并对此进行排序
func addAndSortArticles(arInfo ArticleConfig) {
	//log.Println(len(articles))
	artLen := len(articles)
	//articleListSize 初始长度
	if artLen < articleListSize {
		articles = append(articles, &arInfo)
	}
	log.Println(len(articles))
	sort.Sort(ByDate{articles})
}


