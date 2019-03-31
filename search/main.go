package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Anteoy/liongo/constant"
	cst "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/service"
	"io/ioutil"
	"net/http"
	"strings"
)

type ArticleConfigSearch struct {
	Title    string `json:"title"`    //标题
	Date     string `json:"date"`     //时间
	Classify string `json:"classify"` //所属分类
	Abstract string `json:"abstract"` //摘要
	Author   string `json:"author"`   //作者
	Link     string `json:"link"`     //博客链接
	Content  string `json:"content"`  //完整内容
	Id       string `json:"id"`       //博客唯一id
}
type ArticleConfigslSearch []*ArticleConfigSearch

var ArticleslSearch ArticleConfigslSearch

//注意需要用Post 不能用Delete delete url 跟id可以删除其中一个id
func DeleteByPost(esURL string) {

	url := "http://" + esURL + "/articles/article/_delete_by_query"
	payload := strings.NewReader("{\n" +
		"    \"query\" : { \n" +
		"        \"match_all\" : {}\n" +
		"    }\n" +
		"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

var hValue string

func init() {
	const (
		defaulth = "127.0.0.1:9200"
		usage    = "like host:port"
	)
	flag.StringVar(&hValue, "h", defaulth, usage)
}

func main() {
	defer fmt.Println("the latest defer printed")
	flag.Parse()
	esURL := hValue
	println(esURL)
	//开始解析
	var rf = new(service.SearchFactory)
	rf.Generate(constant.RENDER_DIR)
	for _, ar := range cst.Articlesl {
		arSearch := &ArticleConfigSearch{
			Title:    ar.Title,
			Date:     ar.Date,
			Classify: ar.Classify,
			Abstract: ar.Abstract,
			Author:   ar.Author,
			Link:     ar.Link,
			Content:  ar.Content,
			Id:       ar.Id,
		}
		ArticleslSearch = append(ArticleslSearch, arSearch)
	}

	DeleteByPost(esURL)

	for _, ar := range ArticleslSearch {
		data, _ := json.Marshal(ar)
		println(string(data))
		url := "http://" + esURL + "/articles/article/" + ar.Id
		payload := strings.NewReader(string(data))
		req, _ := http.NewRequest("PUT", url, payload)
		req.Header.Add("Content-Type", "application/json")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res)
		fmt.Println(string(body))

		//resp, _ := http.Post("http://127.0.0.1:9200/articles/article", "application/json", strings.NewReader(string(data)))
		//defer resp.Body.Close()
		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
	}

}
