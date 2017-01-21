package newPosts

import (
	"log"
	"os"
	//"io"
	"fmt"
	"time"
	//"io/ioutil"
	"io"
)

type AddFactory struct {
}

type SampleMD struct {
	date  string
	title string
}

func check(e error) {
	if e != nil {
		log.Println(e)
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//处理liongo new resources生成标准post
func (addFactory *AddFactory) New(title string) {
	var f *os.File
	var err error
	fileName := /*"/root/IdeaProjects/liongo/src/main/go/resources/posts/zhoudazhuang.md"*/ "../resources/posts/" + title + ".md"
	if checkFileIsExist(fileName) { //如果文件存在
		f, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) //打开文件 并可追加文件内容
		f.Name()
		fmt.Println("文件存在" + f.Name())
	} else {
		f, err = os.Create(fileName) //创建文件
		fmt.Println("文件不存在")
	}
	check(err)
	line0 := "---" + "\n"
	line1 := "date: " + time.Now().Format("2006-01-02 15:04:05") + "\n"
	line2 := "title: " + title + "\n"
	line3 := "categories:" + "\n"
	line4 := "    #- golang" + "\n"
	line5 := "tags:" + "\n"
	line6 := "    #- golang" + "\n"
	line7 := "---" + "\n"
	//var d1 = []byte(line1);

	//log.Println(line1)
	_, err1 := io.WriteString(f, line0) //写入文件(字符串)
	_, err1 = io.WriteString(f, line1)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line2)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line3)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line4)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line5)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line6)  //写入文件(字符串)
	_, err1 = io.WriteString(f, line7)  //写入文件(字符串)
	//err1 := ioutil.WriteFile(fileName, d1, 0666)  //写入文件(字节数组)

	defer f.Close()
	check(err1)
	//fmt.Printf("写入 %d 个字节n", n);
	//log.Println(fileName)
}

//处理sample.md文件 提供于new指令
func processSample() {

}
