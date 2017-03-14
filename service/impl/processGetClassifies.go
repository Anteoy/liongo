package impl

import (
	. "github.com/Anteoy/liongo/constant"
)

type ProcessGetClassifies struct {}

//分类first处理方法 放入Classifies map中
func (processClassifies *ProcessGetClassifies)Dispose(dir string)  {
	Classifies = make(map[string]Category) //key为类别，value为Category包含名字，article数组和长度的Category结构体
	//对获取到的所有articles进行处理
	for _, ar := range Articles {
		c, ok := Classifies[ar.Category]
		if ok {
			c.Articles = append(c.Articles, ArticleBase{ar.Link, ar.Title})
			c.Length = len(c.Articles)
			Classifies[ar.Category] = c
		} else { //当此分类为新分类时
			art := ArticleBase{ar.Link, ar.Title}
			arts := make([]ArticleBase, 0)
			arts = append(arts, art)
			Classifies[ar.Category] = Category{ar.Category, arts, 1} //则新建一个并注入
		}
	}
}
