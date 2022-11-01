package readability

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

type Options struct {
	maxElementsToParse int
}

type Readability struct {
	options Options
	doc     *goquery.Document
}

// 奇妙的可选参数实现方式
type ModOption func(*Options) *Options

func New(doc *goquery.Document, modOption ModOption) *Readability {
	options := Options{
		maxElementsToParse: 0,
	}

	if modOption != nil {
		modOption(&options)
	}

	return &Readability{
		options: options,
		doc:     doc,
	}
}

func (r *Readability) unwrapNoScriptImages() {

}

func (r *Readability) Parse() (string, error) {
	// 如果有配置的话，避免解析太大的文件
	if r.options.maxElementsToParse > 0 {
		// TODO 这里不一定指的是 _doc.getElementsByTagName("*").length
		if tagsCount := r.doc.Length(); tagsCount > r.options.maxElementsToParse {
			return "", errors.New("Aborting parsing document;" + string(tagsCount) + " elements found")
		}
	}

	r.unwrapNoScriptImages()
	return "tmp", nil
}
