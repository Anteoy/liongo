package modle

import (
	"time"
)

//mongo note struct
//last config
//---
//date: 2017-01-20 20:12:00
//title: java构造函数以及static关键
//---
type Note struct {
	Name     string  `json:"name"`
	Content  string  `json:"content"`
	Title    string	 `json:"title"`
	Date     string  `json:"date"`
	Time     time.Time `json:"time"`
}



