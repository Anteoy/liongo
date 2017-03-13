package utils

import (
	"github.com/Anteoy/go-gypsy/yaml"
	"log"
	"os"
	"strings"
)

type YamlParser struct{}

//指定配置文件
var YAML_FILES = [3]string{"config.yml", "pages.yml", "nav.yml"}

//返回所有配置文件的key value map
func (yp *YamlParser) Parse(root string) map[string]interface{} {

	//配置文件map
	var yamlFilesConfig = make(map[string]interface{})

	//如果没有后缀/则添加
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	//循环获取每个配置文件信息
	for _, yamlFile := range YAML_FILES {
		path := root + yamlFile
		if !IsExists(path) {
			log.Panic(path + " file not found!")
			os.Exit(1)
		}
		//读取yml
		config, err := yaml.ReadFile(path)
		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}
		//放入map yamlFilesConfig中
		yamlFilesConfig[yamlFile] = config
	}

	return yamlFilesConfig
}
