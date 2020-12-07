package template

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// BannerGenerator 横幅生成器接口
type BannerGenerator interface {
	setTitle(string)
	ShowBanner(string)
}

// Template 模板
type Template struct {
	BannerGenerator

	title string
}

func newTemplate(b BannerGenerator) *Template {
	return &Template{
		BannerGenerator: b,
	}
}

func (t *Template) setTitle(title string) {
	t.title = title
}

// ShowBanner 展示横幅
func (t *Template) ShowBanner(title string) {
	t.BannerGenerator.setTitle(title)
	fmt.Println(t.title)
}

// DotsStyle 简单地将msg首字母大写，并在其之前和之后输出10个点
type DotsStyle struct {
	*Template
}

// NewDotsStyle .
func NewDotsStyle() *DotsStyle {
	d := &DotsStyle{}
	d.Template = newTemplate(d)
	return d
}

func (d *DotsStyle) setTitle(title string) {
	title = strings.Title(title)
	title = fmt.Sprintf(".......... %s ..........", title)
	d.title = title
}

// ASCIIStyle 生成ASCII码艺术字符
type ASCIIStyle struct {
	*Template
}

// NewASCIIStyle .
func NewASCIIStyle() *ASCIIStyle {
	a := &ASCIIStyle{}
	a.Template = newTemplate(a)
	return a
}

func (a *ASCIIStyle) setTitle(title string) {
	title = figure.NewFigure(title, "", true).String()
	a.title = title
}
