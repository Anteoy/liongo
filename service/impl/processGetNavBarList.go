package impl

import (
	. "github.com/Anteoy/liongo/constant"
	"strconv"
	"github.com/Anteoy/go-gypsy/yaml"
	"log"
	"strings"
)

type ProcessGetNavBarList struct {}

//生成自定义多余页面导航条 存入navBarList 数组
//这里配置的示例github导航.
func (processAddNav *ProcessGetNavBarList)Dispose(dir string)  {

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	yCfg := YamlData["nav.yml"]
	var cfg *yaml.File
	if value, ok := yCfg.(*yaml.File); ok {
		cfg = value
	}
	ct, err := cfg.Count("")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < ct; i++ {
		//导航标签名
		name, errn := cfg.Get("[" + strconv.Itoa(i) + "].label")
		if nil != errn {
			log.Println(errn)
		}
		//导航标签链接 支持相对路径
		href, errh := cfg.Get("[" + strconv.Itoa(i) + "].href")
		if nil != errh {
			log.Println(errh)
		}
		//a标签target属性 默认为在当前页面打开
		target, errt := cfg.Get("[" + strconv.Itoa(i) + "].target")
		if nil != errt {
			log.Println(errt)
		}

		nav := NavConfig{name, href, target}
		NavBarList = append(NavBarList, nav)

	}
	log.Println(NavBarList)
}
