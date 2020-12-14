package goeditorjs_test

import (
	"fmt"
	"testing"

	"github.com/davidscottmills/goeditorjs"
	"github.com/stretchr/testify/require"
)

func Test_BlockHeaderHandler_GenerateHTML(t *testing.T) {
	bhh := &goeditorjs.BlockHeaderHandler{}
	testData := []struct {
		level int
		data  string
	}{
		{level: 1, data: `{"text": "Heading","level": 1}`},
		{level: 2, data: `{"text": "Heading","level": 2}`},
		{level: 3, data: `{"text": "Heading","level": 3}`},
		{level: 4, data: `{"text": "Heading","level": 4}`},
		{level: 5, data: `{"text": "Heading","level": 5}`},
		{level: 6, data: `{"text": "Heading","level": 6}`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "header", Data: jsonData}
		html, _ := bhh.GenerateHTML(ejsBlock)
		require.Equal(t, fmt.Sprintf("<h%d>Heading</h%d>", td.level, td.level), html)
	}
}

func Test_BlockParagraphHandler_GenerateHTML_Left(t *testing.T) {
	bph := &goeditorjs.BlockParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	html, _ := bph.GenerateHTML(ejsBlock)
	require.Equal(t, "<p>paragraph</p>", html)
}

func Test_BlockParagraphHandler_GenerateHTML_Center_Right(t *testing.T) {
	bph := &goeditorjs.BlockParagraphHandler{}
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

func Test_BlockListHandler_GenerateHTML(t *testing.T) {
	blh := &goeditorjs.BlockListHandler{}
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

func Test_BlockCodeBoxHandler_GenerateHTML(t *testing.T) {
	bcbh := &goeditorjs.BlockCodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "func main(){fmt.Println(\"HelloWorld\")}"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := `<pre><code class="go">func main(){fmt.Println("HelloWorld")}</code></pre>`
	html, _ := bcbh.GenerateHTML(ejsBlock)
	require.Equal(t, expectedResult, html)
}

func Test_BlockHeaderHandler_Type(t *testing.T) {
	h := &goeditorjs.BlockHeaderHandler{}
	require.Equal(t, "header", h.Type())
}

func Test_BlockParagraphHandler_Type(t *testing.T) {
	h := &goeditorjs.BlockParagraphHandler{}
	require.Equal(t, "paragraph", h.Type())
}

func Test_BlockListHandler_Type(t *testing.T) {
	h := &goeditorjs.BlockListHandler{}
	require.Equal(t, "list", h.Type())
}

func Test_BlockCodeBoxHandler_Type(t *testing.T) {
	h := &goeditorjs.BlockCodeBoxHandler{}
	require.Equal(t, "codeBox", h.Type())
}
