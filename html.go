package goeditorjs

import (
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

// NewHTMLEngine creates a new HTMLEngine
func NewHTMLEngine() *HTMLEngine {
	bhs := make(map[string]HTMLBlockHandler)
	return &HTMLEngine{BlockHandlers: bhs}
}

// RegisterBlockHandlers registers or overrides a block handlers for blockType given by HTMLBlockHandler.Type()
func (htmlEngine *HTMLEngine) RegisterBlockHandlers(handlers ...HTMLBlockHandler) {
	for _, bh := range handlers {
		htmlEngine.BlockHandlers[bh.Type()] = bh
	}
}

// GenerateHTML generates html from the editorJS using configured set of HTML handlers
func (htmlEngine *HTMLEngine) GenerateHTML(editorJSData string) (string, error) {
	result := ""
	ejs, err := parseEditorJSON(editorJSData)
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
