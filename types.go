package goeditorjs

import (
	"encoding/json"
	"errors"
)

// editorJS rpresents the Editor JS data
type editorJS struct {
	Blocks []EditorJSBlock `json:"blocks"`
}

// EditorJSBlock type
type EditorJSBlock struct {
	Type string `json:"type"`
	// Data is the Data for an editorJS block in the form of RawMessage ([]byte). It is left up to the Handler to parse the Data field
	Data json.RawMessage `json:"data"`
}

var (
	//ErrBlockHandlerNotFound is returned from GenerateHTML when the HTML engine doesn't have a registered handler
	//for that type and the HTMLEngine is set to return on errors.
	ErrBlockHandlerNotFound = errors.New("Handler not found for block type")
)

// header represents header data from EditorJS
type header struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

// paragraph represents header data from EditorJS
type paragraph struct {
	Text      string `json:"text"`
	Alignment string `json:"alignment"`
}

// list represents header data from EditorJS
type list struct {
	Style string   `json:"style"`
	Items []string `json:"items"`
}

// codeBox represents header data from EditorJS
type codeBox struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}
