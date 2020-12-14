package goeditorjs

// Header represents header data from EditorJS
type Header struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

// Paragraph represents header data from EditorJS
type Paragraph struct {
	Text      string `json:"text"`
	Alignment string `json:"alignment"`
}

// List represents header data from EditorJS
type List struct {
	Style string   `json:"style"`
	Items []string `json:"items"`
}

// CodeBox represents header data from EditorJS
type CodeBox struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}
