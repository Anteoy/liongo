package service

import (
	"fmt"
	"github.com/Anteoy/blackfriday"
	"github.com/Anteoy/go-gypsy/yaml"
	"github.com/Anteoy/liongo/dao/mongo"
	"github.com/Anteoy/liongo/modle"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/utils"
	m "github.com/Anteoy/liongo/modle"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html"
	"html/template"
	"net/http"
	"os"
	"qiniupkg.com/x/log.v7"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type PNoteService struct{}

type MonthNotes []*MonthNote
//用于sort
func (m MonthNotes) Len() int {
	return len(m)
}
func (m MonthNotes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MonthNotes) Less(i, j int) bool {
	return m[i].month > m[j].month
}

//缓存mongo note数据信息
//某年的note
type YearNote struct {
	Year   string                //如2017
	Months []*MonthNote          // 如1,2,3月
	months map[string]*MonthNote //如 1月的MonthNote
}
//某月的note
type MonthNote struct {
	Month     string     //几月
	month     time.Month //// A Month specifies a month of the year (January = 1, ...).
	NotesBase []*m.NoteBase
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
type NotesByDate struct {
	Notes
}
//sort.Sort() 入参需覆写提供如下方法
func (a Notes) Len() int                 { return len(a) }
func (a Notes) Swap(i, j int)            { a[i], a[j] = a[j], a[i] }
func (a NotesByDate) Less(i, j int) bool { return a.Notes[i].Time.After(a.Notes[j].Time) }

//初始化待用变量
var (
	notes         Notes
	allNotes      YearNotes                   //所有使用年分类的Notes
	yearNotesmap  map[string]*YearNote        //某年所有Note map
	notesListSize int                  = 5000 //最大slice
)

//处理数据库中note 对struct进行装配排序生成到struct
func (p *PNoteService) PreProcessNotes() error {

	//存放article的常量数组 固定最大size 1000
	notes = make([]*m.Note, 0, notesListSize) //make 初始化notes
	//从mongo中获取noteinfo
	//获取连接
	// c := mongo.Session.DB("liongo").C("note")
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note")
	//获取全部数据
	err := c.Find(bson.M{}).All(&notes)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	sort.Sort(NotesByDate{notes})
	for index, value := range notes {
		fmt.Printf("notes[%d]=%d \n", index, value)
		//TODO 处理 url
		processNoteUrl(*value)
	}
	//defer mongo.Session.Close()
	defer sess.Close()
	return nil
}

//string 根据年月日生成note link
func processNoteUrl(ar m.Note) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}

//根据pre获取的notes 进行生成.html操作
func (p *PNoteService) GeneratorNotes() error {
	//根据不同日期生成不同
	return nil
}

//根据pre获取的notes进行生成pnotelist.html操作
func (p *PNoteService) GeneratorPnotelist(root string, yamls map[string]interface{}) error {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := yamls["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := testparseTemplate(root, PNOTELIST_TPL, cfg)
	targetFile := PUBLISH_DIR + "/pnotelist.html"
	//创建targetFile
	fout, err := os.Create(targetFile)
	if err != nil {
		log.Println("create file " + targetFile + " error!")
		os.Exit(1)
	}
	defer fout.Close()
	//时间归档处理
	generatePnotelist()
	//debug pipei tmp
	for index, value := range allNotes {
		fmt.Printf("notes[%d]=%d \n", index, value)
		//fmt.Println(value.Months)
		for _, value1 := range value.Months {
			fmt.Println(value1.Month)
			for _, value2 := range value1.NotesBase {
				fmt.Println(value2.Link + " " + value2.Title)
			}
		}
		fmt.Println(value.months)
		fmt.Println(value.Year) //.Year
	}
	m := map[string]interface{}{"archives": allNotes, "nav": NavBarsl, "cats": Classifiesm} ////注意 这里如果传入参数有误 将会影响到tmp生成的完整性 如footer等 并且此时程序不会报错 但会产生意想不到的结果
	exErr := t.Execute(fout, m)
	return exErr
}

//传入路径和配置信息 返回一个template config.yml tpl 主题所使用的或自定义tmp名 融合footer header body为一个tpl
func testparseTemplate(root, tpl string, cfg *yaml.File) *template.Template {
	//默认default
	themeFolder, errt := cfg.Get("theme")
	if errt != nil {
		log.Println("get theme error!check config.yml at the theme value!")
		os.Exit(1)
	}

	file := root + "templates/" + themeFolder + "/" + tpl + ".tpl"
	if !IsExists(file) {
		log.Println(file + " can not be found!")
		os.Exit(1)
	}
	log.Println(cfg.Get)
	t := template.New(tpl + ".tpl")
	t.Funcs(template.FuncMap{"get": cfg.Get})
	t.Funcs(template.FuncMap{"unescaped": Unescaped})

	headerTpl := root + "templates/" + themeFolder + "/common/" + COMMON_HEADER_FILE
	footerTpl := root + "templates/" + themeFolder + "/common/" + COMMON_FOOTER_FILE

	if !IsExists(headerTpl) {
		log.Println(headerTpl + " can not be found!")
		os.Exit(1)
	}

	if !IsExists(footerTpl) {
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

//根据时间生成有序的notes list
func generatePnotelist() error {
	yearNotesmap = make(map[string]*YearNote) //初始化存储某年notes的map
	for _, iter := range notes {              //排序好的Note指针数组
		y, m, _ := iter.Time.Date() //获取当前的note year和month
		year := fmt.Sprintf("%v", y)
		month := m.String()         // annotation // String returns the English name of the month ("January", "February", ...).
		yNote := yearNotesmap[year] //作为Key储存在yearNotesmap中
		if yNote == nil {           //判断是否有此yearNote日期分类 如果没有则
			//新建一个存入
			//某年的note
			//type YearNote struct {
			//	Year   string //如2017
			//	Months []*MonthNote // 如1,2,3月
			//	months map[string]*MonthNote //如 1月的MonthNote
			//}
			yNote = &YearNote{year, make([]*MonthNote, 0), make(map[string]*MonthNote)}
			yearNotesmap[year] = yNote //放入新的以年分类的Key
		}
		//确认当前note的月份是否在yNote的months节点中存在
		mNote := yNote.months[month]
		if mNote == nil { //是否存在月份小分类
			//不存在则新建立一个并放如其中
			oo := &modle.NoteBase{"test.do", "test"}
			fmt.Println(oo.Link)
			mNote = &MonthNote{month, m, make([]*modle.NoteBase, 0)} //这里开始用m 一直报错undefined,,,m是最近定义了 不会编译为model
			yNote.months[month] = mNote                              //新建并赋值于yNote，内层嵌套
		}
		mNote.NotesBase = append(mNote.NotesBase, &modle.NoteBase{iter.Title, iter.Title}) //年月下嵌入此article TODO 暂时使用title作为Link标识

	}
	allNotes = make(YearNotes, 0)
	//对notes内部使用yNote.months进行排序
	for _, yNote := range yearNotesmap {
		//实例化MonthNotes
		monthCollect := make(MonthNotes, 0)
		//把某年的month全部放入这个数组中
		for _, mNote := range yNote.months { //获取内部months
			monthCollect = append(monthCollect, mNote)
		}
		sort.Sort(monthCollect)            //月份排序
		yNote.months = nil                 //months map[string]*MonthArchive TODO
		yNote.Months = monthCollect        //放入archives struct中Months节点 Months []*MonthArchive 再植入yNote的Months
		allNotes = append(allNotes, yNote) //放入此年的yArchive到allArchive
	}
	return nil
}

//处理Note上传
func (p *PNoteService) DealNoteUpload(md string) error {
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
	//时间解析
	time, terr := time.Parse("2006-01-02 15:04:05", "2017-01-10 20:12:00")
	if terr != nil {
		log.Println(terr)
	}
	fmt.Println(time)
	//装配struct
	note := &modle.Note{Name: "test1", Content: htmlStr, Title: "title1", Date: "2017-01-10 20:12:00", Time: time}
	fmt.Printf(note.Content)
	c := mongo.Session.DB("liongo").C("note")
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//从Mongo中拉取具体note 拉取并生成所有Notes
func (p *PNoteService) GetNoteByName(yamls map[string]interface{}, w http.ResponseWriter, r *http.Request) error {
	//从mongo中获取noteinfo
	//获取连接
	//c := mongo.Session.DB("liongo").C("note")

	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note") //获取数据
	note := modle.Note{}
	err := c.Find(bson.M{"name": "test1"}).One(&note)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	fmt.Println(note.Content)
	fmt.Println(note.Name)
	pnoteservice := new(PNoteService)
	notes := pnoteservice.QueryAllFromMgo()
	for index, value := range *notes { //* 遍历所有的mgo notes
		fmt.Printf("notes[%d]=%d \n", index, value)
		//new 模板对象
		t := template.New("pSpecificNote.tpl")
		yCfg := yamls["config.yml"]
		var cfg = yCfg.(*yaml.File)
		//向模板中注入函数
		t.Funcs(template.FuncMap{"unescaped": Unescaped})
		t.Funcs(template.FuncMap{"get": cfg.Get})

		//openfile := "../resources/templates/default/pSpecificNote.tpl"
		//
		//if !isExists(openfile) {
		//	log.Println(openfile + " can not be found!")
		//	os.Exit(1)
		//}

		//从模板文件解析
		t, errp := t.ParseFiles("../resources/templates/default/pSpecificNote.tpl")
		if errp != nil {
			log.Error(errp)
			panic(err)
		}
		//创建前检查notes文件夹是否存在
		if !IsExists(PUBLISH_DIR+"/notes") {
			//创建777权限目录
			err := os.Mkdir(PUBLISH_DIR+"/notes", 0777)
			if err != nil {
				log.Panic("create publish dir error -- " + err.Error())
			}
		}
		//创建html文件
		targetFile := PUBLISH_DIR + "/notes/" + value.Title + ".html"
		fout, err := os.Create(targetFile)
		if errp != nil {
			log.Error(errp)
			panic(err)
		}
		m := map[string]interface{}{"fi": value, "nav": NavBarsl, "cats": Classifiesm}
		//执行模板的merge操作，输出到fout
		t.Execute(fout, m)
	}
	////new 模板对象 TODO 取消测试版本
	//t := template.New("pSpecificNote.tpl")
	//yCfg := yamls["config.yml"]
	//var cfg = yCfg.(*yaml.File)
	////向模板中注入函数
	//t.Funcs(template.FuncMap{"unescaped": unescaped})
	//t.Funcs(template.FuncMap{"get": cfg.Get})
	//
	////openfile := "../resources/templates/default/pSpecificNote.tpl"
	////
	////if !isExists(openfile) {
	////	log.Println(openfile + " can not be found!")
	////	os.Exit(1)
	////}
	//
	//
	////从模板文件解析
	//t, errp := t.ParseFiles("/root/IdeaProjects/liongo/src/github.com/Anteoy/liongo/resources/templates/default/pSpecificNote.tpl")
	//if errp != nil {
	//	log.Error(errp)
	//	panic(err)
	//}
	////创建html文件
	//targetFile := PUBLISH + "/notes/" + note.Name+".html"
	//fout, err := os.Create(targetFile)
	//m := map[string]interface{}{"fi": note,"nav": navBarList, "cats": classifies}
	////执行模板的merge操作，输出到fout
	//t.Execute(fout, m)
	//http.ServeFile(w, r, targetFile)
	//defer mongo.Session.Close()
	//defer fout.Close()
	defer sess.Close()
	return nil

}

//查询mongo中所有数据
func (p *PNoteService) QueryAllFromMgo() *[]modle.Note {
	//从mongo中获取noteinfo
	//获取连接
	// c := mongo.Session.DB("liongo").C("note")
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note")
	//获取数据
	notes := make([]modle.Note, 100)
	err := c.Find(bson.M{}).All(&notes)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	for index, value := range notes {
		fmt.Printf("notes[%d]=%d \n", index, value)
	}
	//defer mongo.Session.Close() TODO
	defer sess.Close()
	return &notes //&dizhi
}
