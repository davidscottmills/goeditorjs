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

// paragraph represents paragraph data from EditorJS
type paragraph struct {
	Text      string `json:"text"`
	Alignment string `json:"alignment"`
}

// list represents list data from EditorJS
type list struct {
	Style string   `json:"style"`
	Items []string `json:"items"`
}

// codeBox represents code box data from EditorJS
type codeBox struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// raw represents raw html data from EditorJS
type raw struct {
	HTML string `json:"html"`
}

// image represents image data from EditorJS
type image struct {
	File           file   `json:"file"`
	Caption        string `json:"caption"`
	WithBorder     bool   `json:"withBorder"`
	WithBackground bool   `json:"withBackground"`
	Stretched      bool   `json:"stretched"`
}

type file struct {
	URL string `json:"url"`
}
