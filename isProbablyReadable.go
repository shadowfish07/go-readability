package readability

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseStyle(style string) (map[string]string, error) {
	result := make(map[string]string)

	rules := strings.Split(style, ";")

	for _, value := range rules {
		token := strings.Split(value, ":")
		if len(token) != 2 {
			return nil, errors.New("style is not valid")
		}
		result[strings.TrimSpace(token[0])] = strings.TrimSpace(token[1])
	}

	return result, nil
}

func (r *Readability) IsProbablyReadable() bool {
	defaultOptions := struct {
		minScore          int
		minContentLength  int
		visibilityChecker func(*goquery.Document) bool
	}{20, 140, func(doc *goquery.Document) bool {
		style, exist := doc.Attr("style")
		styleMap, err := parseStyle(style)

		fmt.Println(style, exist, styleMap["display"], err)
		return true
	}}
	defaultOptions.visibilityChecker(goquery.NewDocumentFromNode(r.doc.Find("div").First().Nodes[0]))
	return true
}
