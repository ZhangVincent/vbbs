package markdown

import (
	"sync"
	"vbbs/pkg/html"

	"github.com/88250/lute"
	"vbbs/pkg/simple/common/strs"
)

var (
	engine *lute.Lute
	once   sync.Once
)

func getEngine() *lute.Lute {
	once.Do(func() {
		engine = lute.New(func(lute *lute.Lute) {
			// lute.SetToC(true)
			lute.SetSanitize(true)
			lute.SetGFMTaskListItem(true)
		})
	})
	return engine
}

func ToHTML(markdownStr string) string {
	if strs.IsBlank(markdownStr) {
		return ""
	}
	return getEngine().MarkdownStr("", markdownStr)
}

func GetSummary(markdownStr string, summaryLen int) string {
	htmlStr := ToHTML(markdownStr)
	return html.GetSummary(htmlStr, summaryLen)
}
