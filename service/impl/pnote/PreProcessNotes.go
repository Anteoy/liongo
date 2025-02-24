package pnote

import (
	"github.com/Anteoy/liongo/dao/mongo"
	. "github.com/Anteoy/liongo/model"
	log "github.com/Anteoy/liongo/utils/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"strconv"
)

type PreProcessNotes struct{}

//初始化待用变量
var (
	notesl        Notesl
	yearNotesmap  map[string]*YearNote        //key年所有Note map
	allNotesl     YearNotesl                  //所有使用年分类的Notes
	notesListSize int                  = 5000 //最大slice
)

//处理数据库中note 对struct进行装配排序生成到struct
func (preProcessNotes *PreProcessNotes) DisposePnote(str string) {
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
	return
}

//string 根据年月日生成note link
func processNoteUrl(ar Note) string {
	y := strconv.Itoa(ar.Time.Year())
	m := strconv.Itoa(int(ar.Time.Month()))
	d := strconv.Itoa(ar.Time.Day())
	return y + "/" + m + "/" + d + "/"
}
