package constant

const (
	VERSION = "0.1.0"
	USAGE = `
liongo is a static site generator in Go

Usage:

        liongo command [args...]

The commands are:

	build	        			build and generate site.
	run					run the site of blog
		--note				run with the own note
	new	[]				new blog ,generate the new site
	version         			print liongo version

`

	INDEX_TPL = "index"
	BLOG_LIST_TPL = "blog"
	POSTS_TPL    = "posts"
	PAGES_TPL    = "pages"
	ARCHIVE_TPL  = "archive"
	CLASSIFY_TPL = "classify"
	PNOTELOGIN_TPL = "pnotelogin"
	PNOTELIST_TPL  = "pnotelist"

	POST_DIR = "posts"
	PUBLISH_DIR = "../views/serve"
	RENDER_DIR = "../resources"

	COMMON_HEADER_FILE = "header.tpl"
	COMMON_FOOTER_FILE = "footer.tpl"
)

