package utils

import "time"

const (
	PUBLICSH_DIR = "publish"
	RENDER_DIR = "../resources"
)

type ArticleConfig struct {
	Title    string
	Date     string
	ShortDate string
	Category string
	Tags     []TagConfig
	Abstract string
	Author   string
	Time     time.Time
	Link     string
	Content  string
	Nav      []NavConfig
}

type TagConfig struct {
	Name         string
	ArticleTitle string
	ArticleLink  string
}

type NavConfig struct {
	Name   string
	Href   string
	Target string
}