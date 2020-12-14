package goeditorjs_test

import (
	"fmt"
	"testing"

	"github.com/davidscottmills/goeditorjs"
	"github.com/stretchr/testify/require"
)

func Test_BlockHeaderHandler_GenerateHTML(t *testing.T) {
	bhh := &goeditorjs.HeaderHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"text": "Heading","level": 1}`, expectedResult: "<h1>Heading</h1>"},
		{data: `{"text": "Heading","level": 2}`, expectedResult: "<h2>Heading</h2>"},
		{data: `{"text": "Heading","level": 3}`, expectedResult: "<h3>Heading</h3>"},
		{data: `{"text": "Heading","level": 4}`, expectedResult: "<h4>Heading</h4>"},
		{data: `{"text": "Heading","level": 5}`, expectedResult: "<h5>Heading</h5>"},
		{data: `{"text": "Heading","level": 6}`, expectedResult: "<h6>Heading</h6>"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "header", Data: jsonData}
		html, _ := bhh.GenerateHTML(ejsBlock)
		require.Equal(t, td.expectedResult, html)
	}
}

func Test_BlockHeaderHandler_GenerateMarkdown(t *testing.T) {
	bhh := &goeditorjs.HeaderHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"text": "Heading","level": 1}`, expectedResult: "# Heading"},
		{data: `{"text": "Heading","level": 2}`, expectedResult: "## Heading"},
		{data: `{"text": "Heading","level": 3}`, expectedResult: "### Heading"},
		{data: `{"text": "Heading","level": 4}`, expectedResult: "#### Heading"},
		{data: `{"text": "Heading","level": 5}`, expectedResult: "##### Heading"},
		{data: `{"text": "Heading","level": 6}`, expectedResult: "###### Heading"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "header", Data: jsonData}
		md, _ := bhh.GenerateMarkdown(ejsBlock)
		require.Equal(t, td.expectedResult, md)
	}
}

func Test_BlockParagraphHandler_GenerateHTML_Left(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	html, _ := bph.GenerateHTML(ejsBlock)
	require.Equal(t, "<p>paragraph</p>", html)
}

func Test_BlockParagraphHandler_GenerateHTML_Center_Right(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	testData := []struct {
		alignment string
		data      string
	}{
		{alignment: "center", data: `{"text": "paragraph","alignment": "center"}`},
		{alignment: "right", data: `{"text": "paragraph","alignment": "right"}`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
		html, _ := bph.GenerateHTML(ejsBlock)
		require.Equal(t, fmt.Sprintf(`<p style="text-align:%s">paragraph</p>`, td.alignment), html)
	}
}

func Test_BlockParagraphHandler_GenerateMarkdown_Left(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	md, _ := bph.GenerateMarkdown(ejsBlock)
	require.Equal(t, "paragraph", md)
}

func Test_BlockParagraphHandler_GenerateMarkdown_Center_Right(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	testData := []struct {
		alignment string
		data      string
	}{
		{alignment: "center", data: `{"text": "paragraph","alignment": "center"}`},
		{alignment: "right", data: `{"text": "paragraph","alignment": "right"}`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
		md, _ := bph.GenerateMarkdown(ejsBlock)
		require.Equal(t, fmt.Sprintf(`<p style="text-align:%s">paragraph</p>`, td.alignment), md)
	}
}

func Test_BlockListHandler_GenerateHTML(t *testing.T) {
	blh := &goeditorjs.ListHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"style": "ordered", "items": ["one", "two", "three"]}`,
			expectedResult: "<ol><li>one</li><li>two</li><li>three</li></ol>"},
		{data: `{"style": "unordered", "items": ["one", "two", "three"]}`,
			expectedResult: "<ul><li>one</li><li>two</li><li>three</li></ul>"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "list", Data: jsonData}
		html, _ := blh.GenerateHTML(ejsBlock)
		require.Equal(t, td.expectedResult, html)
	}
}

func Test_BlockListHandler_GenerateMarkdown(t *testing.T) {
	blh := &goeditorjs.ListHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"style": "ordered", "items": ["one", "two", "three"]}`,
			expectedResult: "1. one\n1. two\n1. three"},
		{data: `{"style": "unordered", "items": ["one", "two", "three"]}`,
			expectedResult: "- one\n- two\n- three"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "list", Data: jsonData}
		md, _ := blh.GenerateMarkdown(ejsBlock)
		require.Equal(t, td.expectedResult, md)
	}
}

func Test_BlockCodeBoxHandler_GenerateHTML(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "func main(){fmt.Println(\"HelloWorld\")}"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := `<pre><code class="go">func main(){fmt.Println("HelloWorld")}</code></pre>`
	html, _ := bcbh.GenerateHTML(ejsBlock)
	require.Equal(t, expectedResult, html)
}

func Test_BlockCodeBoxHandler_GenerateMarkdown(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "func main(){fmt.Println(\"HelloWorld\")}"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := "```go\nfunc main(){fmt.Println(\"HelloWorld\")}\n```"
	md, _ := bcbh.GenerateMarkdown(ejsBlock)
	require.Equal(t, expectedResult, md)
}

func Test_BlockCodeBoxHandler_GenerateMarkdown_Clean(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "<span class=\"hljs-keyword\"><span class=\"hljs-keyword\">package</span></span> main<div><br></div><div>import <span class=\"hljs-string\"><span class=\"hljs-string\">\"fmt\"</span></span></div><div><br></div><div><span class=\"hljs-function\"><span class=\"hljs-keyword\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span></span></span><span class=\"hljs-function\"> </span><span class=\"hljs-title\"><span class=\"hljs-function\"><span class=\"hljs-title\">main</span></span></span><span class=\"hljs-params\"><span class=\"hljs-function\"><span class=\"hljs-params\">()</span></span></span></span> {</div><div>  fmt.Println(<span class=\"hljs-string\"><span class=\"hljs-string\">\"Hello World\"</span></span>)</div><div>}</div>"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello World\")\n}\n```"
	md, _ := bcbh.GenerateMarkdown(ejsBlock)
	require.Equal(t, expectedResult, md)
}

func Test_BlockHeaderHandler_Type(t *testing.T) {
	h := &goeditorjs.HeaderHandler{}
	require.Equal(t, "header", h.Type())
}

func Test_BlockParagraphHandler_Type(t *testing.T) {
	h := &goeditorjs.ParagraphHandler{}
	require.Equal(t, "paragraph", h.Type())
}

func Test_BlockListHandler_Type(t *testing.T) {
	h := &goeditorjs.ListHandler{}
	require.Equal(t, "list", h.Type())
}

func Test_BlockCodeBoxHandler_Type(t *testing.T) {
	h := &goeditorjs.CodeBoxHandler{}
	require.Equal(t, "codeBox", h.Type())
}
