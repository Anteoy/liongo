package pnote

import (
	log "github.com/Anteoy/liongo/utils/logrus"
	"html"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"sort"
	"time"

	"github.com/russross/blackfriday/v2"
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	"github.com/Anteoy/liongo/dao/mongo"
	"github.com/Anteoy/liongo/model"
	. "github.com/Anteoy/liongo/model"
	. "github.com/Anteoy/liongo/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PNoteService struct{}

//处理数据库中note 对struct进行装配排序生成到struct
func (p *PNoteService) PreProcessNotes() error {

	//存放article的常量数组 固定最大size 1000
	notesl = make([]*Note, 0, notesListSize) //make 初始化notes
	//从mongo中获取noteinfo
	//获取连接
	// c := mongo.Session.DB("liongo").C("note")
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note")
	//获取全部数据
	err := c.Find(bson.M{}).All(&notesl)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	sort.Sort(NotesByDate{notesl})
	for index, value := range notesl {
		log.Printf("notes[%d]=%d \n", index, value)
		//TODO 处理 url
		processNoteUrl(*value)
	}
	//defer mongo.Session.Close()
	defer sess.Close()
	return nil
}

//根据pre获取的notes 进行生成.html操作
func (p *PNoteService) GeneratorNotes() error {
	//根据不同日期生成不同
	return nil
}

//处理Note上传 test use
func (p *PNoteService) DealNoteUpload(md string) error {
	// 判断是否为空
	if len(md) == 0 {
		log.Fatal("md is nil")
		return nil
	}
	//处理md为html
	//markdown字符串转为ASCII html代码
	htmlByte := blackfriday.Run([]byte(md))
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
	// log.Println(time)
	//装配struct
	note := &model.Note{Name: "test1", Content: htmlStr, Title: "title1", Date: "2017-01-10 20:12:00", Time: time}
	// log.Printf(note.Content)
	c := mongo.Session.DB("liongo").C("note")
	err := c.Insert(&note)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//从Mongo中拉取具体note 拉取并生成所有Notes
func (p *PNoteService) GetNotesFromMongo(yamls map[string]interface{}, w http.ResponseWriter, r *http.Request) error {
	//从mongo中获取noteinfo
	//获取连接
	//c := mongo.Session.DB("liongo").C("note")

	//var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	//go mongo.GetMongoSession(ch)
	////var sess *mgo.Session
	//sess := <-(ch)                   //must do
	//c := sess.DB("liongo").C("note") //获取数据
	//note := modle.Note{}
	//err := c.Find(bson.M{"name": "test1"}).One(&note)
	//if err != nil {
	//	log.Error(err)
	//	panic(err)
	//}
	//pnoteservice := new(PNoteService)
	//notes := pnoteservice.QueryAllFromMgo()
	for index, value := range notesl { //* 遍历所有的mgo notes
		log.Printf("notes[%d]=%d \n", index, value)
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
		t, errp := t.ParseFiles("../resources/tpl/pSpecificNote.tpl")
		if errp != nil {
			log.Error(errp)
			panic(errp)
		}
		//检查PUBLISH_DIR是否存在
		if !IsExists(PUBLISH_DIR) {
			err := os.MkdirAll(PUBLISH_DIR, 0777)
			if err != nil {
				log.Panic("create publish dir error -- " + err.Error())
			}
		}
		//创建前检查notes文件夹是否存在
		if !IsExists(PUBLISH_DIR + "/notes") {
			//创建777权限目录
			err := os.Mkdir(PUBLISH_DIR+"/notes", 0777)
			if err != nil {
				log.Panic("create publish dir + /notes error -- " + err.Error())
			}
		}
		//创建html文件
		targetFile := PUBLISH_DIR + "/notes/" + value.Title + ".html"
		fout, err := os.Create(targetFile)
		if errp != nil {
			log.Error(errp)
			panic(err)
		}
		m := map[string]interface{}{"fi": value, "nav": NavBarsl}
		//执行模板的merge操作，输出到fout
		t.Execute(fout, m)
	}
	//defer sess.Close()
	return nil

}

//查询mongo中所有数据
func (p *PNoteService) QueryAllFromMgo() *[]model.Note {
	//从mongo中获取noteinfo
	//获取连接
	// c := mongo.Session.DB("liongo").C("note")
	var ch chan *mgo.Session = make(chan *mgo.Session, 1)
	go mongo.GetMongoSession(ch)
	var sess *mgo.Session            //must init
	sess = <-(chan *mgo.Session)(ch) //must do
	c := sess.DB("liongo").C("note")
	//获取数据
	notes := make([]model.Note, 100)
	err := c.Find(bson.M{}).All(&notes)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	for index, value := range notes {
		log.Printf("notes[%d]=%d \n", index, value)
	}
	//defer mongo.Session.Close() TODO
	defer sess.Close()
	return &notes //&dizhi
}
