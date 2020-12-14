package goeditorjs

import (
	"encoding/json"
	"errors"
	"fmt"
)

// HTMLEngine is the engine that creates the HTML from EditorJS blocks
type HTMLEngine struct {
	BlockHandlers map[string]HTMLBlockHandler
}

// HTMLBlockHandler is an interface for a plugable EditorJS HTML generator
type HTMLBlockHandler interface {
	Type() string // Type returns the type the block handler supports as a string
	GenerateHTML(editorJSBlock EditorJSBlock) (string, error)
}

// EditorJS rpresents the Editor JS data
type EditorJS struct {
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

// NewHTMLEngine creates a new HTMLEngine
func NewHTMLEngine() *HTMLEngine {
	bhs := make(map[string]HTMLBlockHandler)
	return &HTMLEngine{BlockHandlers: bhs}
}

// RegisterBlockHandler registers or overrides a new block handler for blockType
func (htmlEngine *HTMLEngine) RegisterBlockHandler(bh HTMLBlockHandler) {
	htmlEngine.BlockHandlers[bh.Type()] = bh
}

// GenerateHTML generates html from the editorJS using configured set of HTML Generators
func (htmlEngine *HTMLEngine) GenerateHTML(editorJSData string) (string, error) {
	result := ""
	ejs, err := ParseEditorJSON(editorJSData)
	if err != nil {
		return "", err
	}
	for _, block := range ejs.Blocks {
		if generator, ok := htmlEngine.BlockHandlers[block.Type]; ok {
			html, err := generator.GenerateHTML(block)
			if err != nil {
				return result, err
			}
			result += html
		} else {
			return "", fmt.Errorf("%w, Block Type: %s", ErrBlockHandlerNotFound, block.Type)
		}
	}

	return result, nil
}

// ParseEditorJSON parses editorJS data
func ParseEditorJSON(editorJSData string) (*EditorJS, error) {
	result := &EditorJS{}
	err := json.Unmarshal([]byte(editorJSData), result)
	if err != nil {
		return nil, err
	}
	return result, err
}
