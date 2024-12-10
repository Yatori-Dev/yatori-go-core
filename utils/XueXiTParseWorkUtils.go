package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// QuestionSet 结构体用于存储每个 div.Py-mian1 的原始 HTML 内容
type QuestionSet struct {
	ID   string // 可以根据需要添加其他字段，例如题目ID
	HTML string
}

// ParseQuestionSets 函数用于提取所有 div.Py-mian1 节点的原始 HTML 内容
func ParseQuestionSets(doc *goquery.Document) []QuestionSet {
	questionNodes := doc.Find("div.Py-mian1")
	var questionSets []QuestionSet

	questionNodes.Each(func(i int, questionNode *goquery.Selection) {
		// 获取 div.Py-mian1 的 data 属性（如果有）
		dataAttr, exists := questionNode.Attr("data")
		if !exists {
			dataAttr = fmt.Sprintf("question_%d", i+1)
		}

		// 提取原始 HTML 内容
		htmlContent, _ := questionNode.Html()

		// 添加到结果切片
		questionSets = append(questionSets, QuestionSet{
			ID:   dataAttr,
			HTML: htmlContent,
		})
	})

	return questionSets
}
