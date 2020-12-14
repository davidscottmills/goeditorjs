package goeditorjs

import (
	"encoding/json"
	"fmt"
)

// BlockHeaderHandler is the default BlockHeaderHandler for EditorJS HTML generation
type BlockHeaderHandler struct{}

func (*BlockHeaderHandler) parse(editorJSBlock EditorJSBlock) (*Header, error) {
	header := &Header{}
	return header, json.Unmarshal(editorJSBlock.Data, header)
}

// Type "header"
func (*BlockHeaderHandler) Type() string {
	return "header"
}

// GenerateHTML generate html for HeaderBlocks
func (h *BlockHeaderHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<h%d>%s</h%d>", header.Level, header.Text, header.Level), nil
}

// BlockParagraphHandler is the default BlockParagraphHandler for EditorJS HTML generation
type BlockParagraphHandler struct{}

func (*BlockParagraphHandler) parse(editorJSBlock EditorJSBlock) (*Paragraph, error) {
	paragraph := &Paragraph{}
	return paragraph, json.Unmarshal(editorJSBlock.Data, paragraph)
}

// Type "paragraph"
func (*BlockParagraphHandler) Type() string {
	return "paragraph"
}

// GenerateHTML generate html for ParagraphBlocks
func (h *BlockParagraphHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	paragraph, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	if paragraph.Alignment != "left" {
		return fmt.Sprintf(`<p style="text-align:%s">%s</p>`, paragraph.Alignment, paragraph.Text), nil
	}

	return fmt.Sprintf(`<p>%s</p>`, paragraph.Text), nil
}

// BlockListHandler is the default BlockListHandler for EditorJS HTML generation
type BlockListHandler struct{}

func (*BlockListHandler) parse(editorJSBlock EditorJSBlock) (*List, error) {
	list := &List{}
	return list, json.Unmarshal(editorJSBlock.Data, list)
}

// Type "list"
func (*BlockListHandler) Type() string {
	return "list"
}

// GenerateHTML generate html for ListBlocks
func (h *BlockListHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	list, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	result := ""
	if list.Style == "ordered" {
		result = "<ol>%s</ol>"
	} else {
		result = "<ul>%s</ul>"
	}

	innerData := ""
	for _, s := range list.Items {
		innerData += fmt.Sprintf("<li>%s</li>", s)
	}

	return fmt.Sprintf(result, innerData), nil
}

// BlockCodeBoxHandler is the default BlockCodeBoxHandler for EditorJS HTML generation
type BlockCodeBoxHandler struct{}

func (*BlockCodeBoxHandler) parse(editorJSBlock EditorJSBlock) (*CodeBox, error) {
	codeBox := &CodeBox{}
	return codeBox, json.Unmarshal(editorJSBlock.Data, codeBox)
}

// Type "codeBox"
func (*BlockCodeBoxHandler) Type() string {
	return "codeBox"
}

// GenerateHTML generate html for CodeBoxBlocks
func (h *BlockCodeBoxHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	codeBox, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<pre><code class="%s">%s</code></pre>`, codeBox.Language, codeBox.Code), nil
}
