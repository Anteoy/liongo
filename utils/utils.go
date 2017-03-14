package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"github.com/Anteoy/go-gypsy/yaml"
	. "github.com/Anteoy/liongo/constant"
	"html/template"
	"strings"
)

func TrimHTML(str string) string {
	if str == "" {
		return str
	}
	re, _ := regexp.Compile(`<[\s\S]+?(>|$)`)
	newstr := re.ReplaceAllString(str, "")
	return newstr
}

func SubStr(str string, start, end int) string {
	if start < 0 {
		log.Panic("start position is wrong!")
	}
	if end > len(str) {
		log.Panic("end positon is wrong!")
	}
	if start > end {
		log.Panic("wrong position!")
	}

	rs := []rune(str)
	return string(rs[start:end])
}

/**
检测文件是否存在 Stat返回fileInfo
*/
func IsExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//复制文件
func CopyFile(src, dst string) (w int64, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()
	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer dstf.Close()
	return io.Copy(dstf, f)
}

//递归复制目录以及其文件
func CopyDir(source, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			_, err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

func XmlEscapString(str string) string {
	str = strings.Replace(str, `<pre class="prettyprint linenums">`, "@@PRE_BEGIN", -1)
	str = strings.Replace(str, `</pre>`, "@@PRE_END", -1)
	str = template.HTMLEscapeString(str)
	str = strings.Replace(str, "@@PRE_BEGIN", `<pre class="prettyprint linenums">`, -1)
	str = strings.Replace(str, "@@PRE_END", `</pre>`, -1)
	return str
}

//转义
func Unescaped(str string) interface{} {
	re := regexp.MustCompile(`<pre class="prettyprint linenums">([\s\S]*?)</pre>`)
	str = re.ReplaceAllStringFunc(str, XmlEscapString)
	return template.HTML(str)

}

//传入路径和配置信息 返回一个template config.yml tpl 主题所使用的或自定义tmp名 融合footer header body为一个tpl
func ParseTemplate(root, tpl string, cfg *yaml.File) *template.Template {
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

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}
