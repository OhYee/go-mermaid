package mermaid

import (
	"context"
	"fmt"
	"strings"

	_ "embed"

	"github.com/chromedp/chromedp"
)

//go:generate wget https://cdn.bootcdn.net/ajax/libs/mermaid/8.9.2/mermaid.min.js

//go:embed mermaid.min.js
var mermaidjs string

const content = `data:text/html,<!DOCTYPE html>
<html lang="en">
	<head><meta charset="utf-8"></head>
	<body></body>
</html>
`

func Render(source string) (result string) {
	ctx, cancel := chromedp.NewContext(
		// allocCtx,
		context.Background(),
		// chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	source = strings.ReplaceAll(source, "`", "\\`") // 转义字符
	source = strings.TrimSpace(source)              // 清除首尾空格

	var ok bool
	if err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf(content)),
		chromedp.Evaluate(mermaidjs, &ok),                                                   // 加载 mermaid.js
		chromedp.Evaluate(fmt.Sprintf("mermaid.render('mermaid', `%s`);", source), &result), // 生成 svg
	); err != nil {
		result = "<p style='color:red'>render error</p>"
	}
	return
}
