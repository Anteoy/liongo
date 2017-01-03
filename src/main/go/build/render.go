package build

import (
	"main/go/utils"
	//myconst "main/go/constant"
)

type RenderFactory struct{}

const (
	INDEX_TPL    = "index"
	TAG_TPL      = "tag"
	POSTS_TPL    = "posts"
	PAGES_TPL    = "pages"
	RSS_TPL      = "rss"
	CATEGORY_TPL = "category"
	ARCHIVE_TPL  = "archive"
)

const (
	POST_DIR     = "posts"
	PUBLICSH_DIR = "publish"
)

const (
	COMMON_HEADER_FILE = "header.tpl"
	COMMON_FOOTER_FILE = "footer.tpl"
	
)

//pre process posts pages
func (self *RenderFactory) PreProcessPosts(root string, yamls map[string]interface{}) error {
	return nil
}

//root 资源文件的相对路径resources yamlData 读取配置文件的键值对
func (self *RenderFactory) Render(root string) {
	yp := new(utils.YamlParser)
	yamlData := yp.Parse(root)
	self.PreProcessPosts(root,yamlData)
}
