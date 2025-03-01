package impl

import (
	"bufio"
	"fmt"
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/utils"
	"github.com/russross/blackfriday/v2"
	"html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ProcessGetPostsArticle struct{}

//加载posts文件夹下的md博文信息
func (processPosts *ProcessGetPostsArticle) Dispose(dir string) {

	//添加后缀
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	//获取config.yml键值对节点信息
	yCfg := YamlData["config.yml"]

	//Comma-ok断言 不推荐下面第二种方式 下面第二种如果转换失败会直接panic
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	//cfg = yCfg.((*yaml.File))
	//存放article的常量数组 数组的长度并且也是切片的初始长度为0 capacity（容量） 5000
	Articlesl = make([]*ArticleConfig, 0, ArticleListSize)
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
			log.Println("begin process article ： " + fileInfo.Name())
			fileName := fileInfo.Name()
			//获取posts文件夹下md文件信息 原始markdownstring config ArticleConfig
			//返回md正文，读取md配置组合，error，组装配置到articleConfig
			mardownStr, articleConfig, err := processArticleFile(dir+POST_DIR+"/"+fileName, fileName)
			if err != nil {
				log.Println("preprocess article file error!" + fileName)
				os.Exit(1)
			}
			//去掉文件.md后缀
			//htmlName := strings.TrimSuffix(fileName, ".md")
			//根据.md中配置生成年月日文件路径字符串 返回html前一级路径(年月日路径字符串)
			p := processArticleUrl(articleConfig)
			log.Println(p)
			//markdown字符串转为ASCII html代码 []byte(mardownStr) string强转为[]byte
			htmlByte := blackfriday.Run([]byte(mardownStr))
			//反转义实体如“& lt;”成为“<” 把byte转位strings
			htmlStr := html.UnescapeString(string(htmlByte))
			re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
			htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
			//增加正文和链接 组装ArticleConfig
			articleConfig.Content = htmlStr
			articleConfig.Link = p + articleConfig.Id + ".html"
			//装配摘要Abstract
			if articleConfig.Abstract == "" {
				var limit int = 1000
				rs := []rune(htmlStr) //int32
				if len(rs) < 1000 {
					limit = len(rs)
				}
				//组装ArticleConfig摘要
				abstract := utils.SubStr(htmlStr, 0, limit)
				articleConfig.Abstract = utils.TrimHTML(abstract)
			}
			//装配作者
			if articleConfig.Author == "" {
				//从配置文件获取author
				author, cerr := cfg.Get("meta.author")
				if cerr != nil {
					log.Println(cerr)
				}
				articleConfig.Author = author
			}
			//添加article到articles 并对此进行排序 传入前面获取的ArticleConfig TODO 提高效率
			addAndSortArticles(articleConfig)
		}
	}

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
			if content == "---" && flag < 2 {
				flag++
			}
			if flag == 2 {
				if content != "---" {
					//获取article 正文markdownStr
					markdownStr += content + "\n"
				}
			} else if flag < 2 {
				//获取article正文前配置信息
				yamlStr += content + "\n"
			}

		}

	}
	//把md中---配置部分交于yaml进行处理（md中配置也基于yaml） 去掉所有---\n
	config := yaml.Config(strings.Replace(yamlStr, "---\n", "", -1))
	//获取md中配置说明信息
	title, err := config.Get("title")
	//获取时间
	date, err := config.Get("date")
	//获取id
	id, err := config.Get("id")
	fmt.Println(id)
	//获取标签
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

	arInfo := ArticleConfig{title, date, shortDate, cat, tags, abstract, author, t, "", "", NavBarsl, id}

	//log.Println(markdownStr)
	return markdownStr, arInfo, nil

}

//ArticleConfig 解析获取的.md --- --- 中配置文件struct
//string 根据年月日生成article路径
//int 转为 string strconv.Itoa.
func processArticleUrl(ar ArticleConfig) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}

//添加article到articles 并对此进行排序
func addAndSortArticles(arInfo ArticleConfig) {
	//log.Println(len(articles))
	artLen := len(Articlesl)
	//slice最大长度
	if artLen < ArticleListSize {
		//必须使用引用
		Articlesl = append(Articlesl, &arInfo)
	}
	log.Println(len(Articlesl))
	sort.Sort(ByDate{Articlesl})
}
