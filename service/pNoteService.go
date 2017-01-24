package service

import (
	"github.com/Anteoy/blackfriday"
	"qiniupkg.com/x/log.v7"
	"github.com/Anteoy/liongo/modle"
	"github.com/Anteoy/liongo/dao/mongo"
	"github.com/Anteoy/go-gypsy/yaml"
	m "github.com/Anteoy/liongo/modle"
	"regexp"
	"html"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"os"
	"net/http"
	"time"
	"sort"
)

type PNoteService struct{}

//缓存mongo note数据信息

//某年的note
type YearNote struct {
	Year   string //如2017
	Months []*MonthNote // 如1,2,3月
	months map[string]*MonthNote //如 1月的MonthNote
}
//某月的note
type MonthNote struct {
	Month    string //几月
	month    time.Month //// A Month specifies a month of the year (January = 1, ...).
	Notes []*modle.Note
}
//所有的note储存
type YearNotes []*YearNote
//用于sort
func (y YearNotes) Len() int {
	return len(y)
}

func (y YearNotes) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}

func (y YearNotes) Less(i, j int) bool {
	return y[i].Year > y[j].Year
}

type Notes []*m.Note

type NotesByDate struct{
	Notes
}

//sort.Sort() 入参需覆写提供如下方法
func (a Notes) Len() int            { return len(a) }
func (a Notes) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a NotesByDate) Less(i, j int) bool { return a.Notes[i].Time.After(a.Notes[j].Time) }
//添加article到articles 并对此进行排序 每次传入一个Note
func AddAndSortNotes(noteInfo m.Note) {
	artLen := len(notes)
	if artLen < notesListSize {
		notes = append(notes, &noteInfo)
	}
	log.Println(len(notes))
	sort.Sort(NotesByDate{notes})
}

//初始化待用变量
var (
	notes        Notes
	allNotes YearNotes //所有使用年分类的Notes
	yearNotesmap map[string]*YearArchive //某年所有Note map
	notesListSize int = 5000 //最大slice
)

//处理数据库中note 对struct进行装配排序生成到struct
func (p *PNoteService) PreProcessNotes() error{

	//存放article的常量数组 固定最大size 1000
	notes = make([]*m.Note, 0, notesListSize) //make 初始化notes
	//从mongo中获取noteinfo
	//获取连接
	c := mongo.Session.DB("liongo").C("note")
	//获取全部数据
	err := c.Find(bson.M{}).All(&notes)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	sort.Sort(NotesByDate{notes})
	return nil
}

//处理Note上传
func (p *PNoteService) DealNoteUpload(md string)  error {
	// 判断是否为空
	if len(md) == 0 {
		log.Fatal("md is nil")
		return nil
	}
	//处理md为html
	//markdown字符串转为ASCII html代码
	htmlByte := blackfriday.MarkdownCommon([]byte(md))
	//反转义实体如“& lt;”成为“<” 把byte转位strings
	htmlStr := html.UnescapeString(string(htmlByte))
	//正则匹配并替换
	re := regexp.MustCompile(`<pre><code>([\s\S]*?)</code></pre>`)
	htmlStr = re.ReplaceAllString(htmlStr, `<pre class="prettyprint linenums">${1}</pre>`)
	//装配struct
	note := &modle.Note{Name: "test1", Content: htmlStr}
	fmt.Printf(note.Content)
	c := mongo.Session.DB("liongo").C("note")
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
//从Mongo中拉取具体note
func (p *PNoteService) GetNoteByName(name string,yamls map[string]interface{},w http.ResponseWriter, r *http.Request) error {
	if len(name) == 0 {
		fmt.Println("传入Note name为空，请检查！！！")
		return nil
	}
	//从mongo中获取noteinfo
	//获取连接
	c := mongo.Session.DB("liongo").C("note")
	//获取数据
	note := modle.Note{}
	err := c.Find(bson.M{"name": "test1"}).One(&note)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	fmt.Println(note.Content)
	fmt.Println(note.Name)
	//new 模板对象
	t := template.New("pSpecificNote.tpl")
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	//向模板中注入函数
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	t.Funcs(template.FuncMap{"get": cfg.Get})

	//openfile := "../resources/templates/default/pSpecificNote.tpl"
	//
	//if !isExists(openfile) {
	//	log.Println(openfile + " can not be found!")
	//	os.Exit(1)
	//}


	//从模板文件解析
	t, errp := t.ParseFiles("/root/IdeaProjects/liongo/src/github.com/Anteoy/liongo/resources/templates/default/pSpecificNote.tpl")
	if errp != nil {
		log.Error(errp)
		panic(err)
	}
	//创建html文件
	targetFile := PUBLISH + "/notes/" + note.Name+".html"
	fout, err := os.Create(targetFile)
	m := map[string]interface{}{"fi": note,"nav": navBarList, "cats": classifies}
	//执行模板的merge操作，输出到fout
	t.Execute(fout, m)
	http.ServeFile(w, r, targetFile)
	defer mongo.Session.Close()
	defer fout.Close()
	return nil

}

//查询mongo中所有数据
func (p *PNoteService) QueryAll(){
	//从mongo中获取noteinfo
	//获取连接
	c := mongo.Session.DB("liongo").C("note")
	//获取数据
	notes := make([]modle.Note,100)
	err := c.Find(bson.M{}).All(&notes)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	for index,value := range notes{
		fmt.Printf("notes[%d]=%d \n", index, value)
	}
	defer mongo.Session.Close()
}
