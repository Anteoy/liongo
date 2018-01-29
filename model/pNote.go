package model

import (
	"time"
)

//mongo note struct.
type Note struct {
	ID      string     `json:"_id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Title   string    `json:"title"`
	Date    string    `json:"date"`
	Time    time.Time `json:"time"`
	Origin  string 	  `json:"origin"`
}

//提供note list html
type NoteBase struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

//缓存mongo note数据信息
//某年的note
type YearNote struct {
	Year      string                //如2017
	Months    []*MonthNote          // 如1,2,3月
	Monthsmap map[string]*MonthNote //如 1月的MonthNote
}

//某月的note
type MonthNote struct {
	Month     string     //几月
	MonthEn   time.Month //// A Month specifies a month of the year (January = 1, ...).
	NotesBase []*NoteBase
}

type MonthNotesl []*MonthNote

//用于sort
func (m MonthNotesl) Len() int {
	return len(m)
}
func (m MonthNotesl) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MonthNotesl) Less(i, j int) bool {
	return m[i].MonthEn > m[j].MonthEn
}

//所有的note储存
type YearNotesl []*YearNote

//用于sort
func (y YearNotesl) Len() int {
	return len(y)
}
func (y YearNotesl) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}
func (y YearNotesl) Less(i, j int) bool {
	return y[i].Year > y[j].Year
}
type Notesl []*Note
type NotesByDate struct {
	Notesl
}

//sort.Sort() 入参需覆写提供如下方法
func (a Notesl) Len() int                { return len(a) }
func (a Notesl) Swap(i, j int)           { a[i], a[j] = a[j], a[i] }
func (a NotesByDate) Less(i, j int) bool { return a.Notesl[i].Time.After(a.Notesl[j].Time) }
