package pnote

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	. "github.com/Anteoy/liongo/model"
	. "github.com/Anteoy/liongo/utils"
	"github.com/Anteoy/liongo/utils/logrus"
)

type GeneratorPnotelist struct{}

//根据pre获取的notes进行生成pnotelist.html操作
func (generatorPnotelist *GeneratorPnotelist) DisposePnote(root string) {
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	yCfg := YamlData["config.yml"]
	var cfg = yCfg.(*yaml.File)
	t := ParseTemplate(root, PNOTELIST_TPL, cfg)
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
	for index, value := range allNotesl {
		logrus.Debug("notes[%d]=%d \n", index, value)
		//fmt.Println(value.Months)
		for _, value1 := range value.Months {
			logrus.Debug(value1.Month)
			for _, value2 := range value1.NotesBase {
				logrus.Debug(value2.Link + " " + value2.Title)
			}
		}
		logrus.Debug(value.Monthsmap)
		logrus.Debug(value.Year) //.Year
	}
	m := map[string]interface{}{"archives": allNotesl, "nav": NavBarsl} ////注意 这里如果传入参数有误 将会影响到tmp生成的完整性 如footer等 并且此时程序不会报错 但会产生意想不到的结果
	t.Execute(fout, m)
}

//根据有序notesl生成yearNotesmap allNotesl
func generatePnotelist() error {
	yearNotesmap = make(map[string]*YearNote) //初始化存储某年notes的map
	for _, iter := range notesl {             //排序好的Note指针数组
		y, m, _ := iter.Time.Date() //获取当前的note year和month
		year := fmt.Sprintf("%v", y)
		month := m.String()         // annotation // String returns the English name of the month ("January", "February", ...).
		yNote := yearNotesmap[year] //作为Key储存在yearNotesmap中
		if yNote == nil {           //判断是否有此yearNote日期分类 如果没有则
			//新建一个存入
			//某年的note
			yNote = &YearNote{year, make([]*MonthNote, 0), make(map[string]*MonthNote)}
			yearNotesmap[year] = yNote //放入新的以年分类的Key
		}
		//确认当前note的月份是否在yNote的months节点中存在
		mNote := yNote.Monthsmap[month]
		if mNote == nil { //是否存在月份小分类
			//test
			oo := &NoteBase{"test.do", "test"}
			logrus.Debug(oo.Link)
			//不存在则新建立一个并放如其中
			mNote = &MonthNote{month, m, make([]*NoteBase, 0)} //这里开始用m 一直报错undefined,,,m是最近定义了 不会编译为model
			yNote.Monthsmap[month] = mNote                     //新建并赋值于yNote，内层嵌套
		}
		mNote.NotesBase = append(mNote.NotesBase, &NoteBase{iter.Title, iter.Title}) //年月下嵌入此article TODO 暂时使用title作为Link标识
	}

	allNotesl = make(YearNotesl, 0)
	//对notes内部使用yNote.months进行排序
	for _, yNote := range yearNotesmap {
		//实例化MonthNotes
		monthCollect := make(MonthNotesl, 0)
		//把某年的month全部放入这个数组中 某年的所有notes
		for _, mNote := range yNote.Monthsmap { //获取内部months
			monthCollect = append(monthCollect, mNote)
		}
		sort.Sort(monthCollect) //月份排序
		yNote.Monthsmap = nil   //invalid exchange
		yNote.Months = monthCollect
		allNotesl = append(allNotesl, yNote) //放入此年的yArchive到allArchive yNote已有序
		//fix sort
		sort.Sort(allNotesl)
	}
	return nil
}
