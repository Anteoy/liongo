package impl

import (
	. "github.com/Anteoy/liongo/constant"
)

type ProcessGetClassifies struct{}

//分类first处理方法 放入Classifies map中
func (processClassifies *ProcessGetClassifies) Dispose(dir string) {
	Classifiesm = make(map[string]Category) //key为类别，value为Category包含名字，article数组和长度的Category结构体
	//对获取到的所有articles进行处理
	for _, ar := range Articlesl {
		//获取里面的分类 如果map中存在则添加，如果不存在则新建
		category, ok := Classifiesm[ar.Classify]
		if ok {
			//新建articlebase
			category.Articles = append(category.Articles, ArticleBase{ar.Link, ar.Title})
			category.Length = len(category.Articles)
			Classifiesm[ar.Classify] = category
		} else { //当此分类为新分类时
			articleBase := ArticleBase{ar.Link, ar.Title}
			artsl := make([]ArticleBase, 0)
			artsl = append(artsl, articleBase)
			Classifiesm[ar.Classify] = Category{ar.Classify, artsl, 1} //则新建一个并注入
		}
	}
}
