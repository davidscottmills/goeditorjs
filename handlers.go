package goeditorjs

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// HeaderHandler is the default HeaderHandler for EditorJS HTML generation
type HeaderHandler struct{}

func (*HeaderHandler) parse(editorJSBlock EditorJSBlock) (*header, error) {
	header := &header{}
	return header, json.Unmarshal(editorJSBlock.Data, header)
}

// Type "header"
func (*HeaderHandler) Type() string {
	return "header"
}

// GenerateHTML generates html for HeaderBlocks
func (h *HeaderHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<h%d>%s</h%d>", header.Level, header.Text, header.Level), nil
}

// GenerateMarkdown generates markdown for HeaderBlocks
func (h *HeaderHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", strings.Repeat("#", header.Level), header.Text), nil
}

// ParagraphHandler is the default ParagraphHandler for EditorJS HTML generation
type ParagraphHandler struct{}

func (*ParagraphHandler) parse(editorJSBlock EditorJSBlock) (*paragraph, error) {
	paragraph := &paragraph{}
	return paragraph, json.Unmarshal(editorJSBlock.Data, paragraph)
}

// Type "paragraph"
func (*ParagraphHandler) Type() string {
	return "paragraph"
}

// GenerateHTML generates html for ParagraphBlocks
func (h *ParagraphHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	paragraph, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	if paragraph.Alignment != "left" {
		return fmt.Sprintf(`<p style="text-align:%s">%s</p>`, paragraph.Alignment, paragraph.Text), nil
	}

	return fmt.Sprintf(`<p>%s</p>`, paragraph.Text), nil
}

// GenerateMarkdown generates markdown for ParagraphBlocks
func (h *ParagraphHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	paragraph, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	if paragraph.Alignment != "left" {
		// Native markdown doesn't support alignment, so we'll use html instead.
		return fmt.Sprintf(`<p style="text-align:%s">%s</p>`, paragraph.Alignment, paragraph.Text), nil
	}

	return paragraph.Text, nil
}

// ListHandler is the default ListHandler for EditorJS HTML generation
type ListHandler struct{}

func (*ListHandler) parse(editorJSBlock EditorJSBlock) (*list, error) {
	list := &list{}
	return list, json.Unmarshal(editorJSBlock.Data, list)
}

// Type "list"
func (*ListHandler) Type() string {
	return "list"
}

// GenerateHTML generates html for ListBlocks
func (h *ListHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
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

// GenerateMarkdown generates markdown for ListBlocks
func (h *ListHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	list, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	listItemPrefix := ""
	if list.Style == "ordered" {
		listItemPrefix = "1. "
	} else {
		listItemPrefix = "- "
	}

	results := []string{}
	for _, s := range list.Items {
		results = append(results, listItemPrefix+s)
	}

	return strings.Join(results, "\n"), nil
}

// CodeBoxHandler is the default CodeBoxHandler for EditorJS HTML generation
type CodeBoxHandler struct{}

func (*CodeBoxHandler) parse(editorJSBlock EditorJSBlock) (*codeBox, error) {
	codeBox := &codeBox{}
	return codeBox, json.Unmarshal(editorJSBlock.Data, codeBox)
}

// Type "codeBox"
func (*CodeBoxHandler) Type() string {
	return "codeBox"
}

// GenerateHTML generates html for CodeBoxBlocks
func (h *CodeBoxHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	codeBox, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<pre><code class="%s">%s</code></pre>`, codeBox.Language, codeBox.Code), nil
}

// GenerateMarkdown generates markdown for CodeBoxBlocks
func (h *CodeBoxHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	codeBox, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	codeBox.Code = strings.ReplaceAll(codeBox.Code, "<div>", "\n")
	codeBox.Code = removeHTMLTags(codeBox.Code)

	return fmt.Sprintf("```%s\n%s\n```", codeBox.Language, codeBox.Code), nil
}

func removeHTMLTags(in string) string {
	// regex to match html tag
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}

// RawHTMLHandler is the default raw handler for EditorJS HTML generation
type RawHTMLHandler struct{}

// Type "raw"
func (*RawHTMLHandler) Type() string {
	return "raw"
}

// GenerateHTML generates html for rawBlocks
func (h *RawHTMLHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	return h.raw(editorJSBlock)

}

// GenerateMarkdown generates markdown for rawBlocks
func (h *RawHTMLHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	return h.raw(editorJSBlock)
}

func (h *RawHTMLHandler) raw(editorJSBlock EditorJSBlock) (string, error) {
	raw := &raw{}
	err := json.Unmarshal(editorJSBlock.Data, raw)
	if err != nil {
		return "", err
	}

	return raw.HTML, nil
}
