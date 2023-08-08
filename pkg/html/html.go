package html

import (
	"strings"

	"vbbs/pkg/text"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"vbbs/pkg/simple/common/strs"
)

func GetSummary(htmlStr string, summaryLen int) string {
	if summaryLen <= 0 || strs.IsEmpty(htmlStr) {
		return ""
	}
	return text.GetSummary(GetHtmlText(htmlStr), summaryLen)
}

// GetHtmlText 获取html文本
func GetHtmlText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return doc.Text()
}
