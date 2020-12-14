package goeditorjs

import "encoding/json"

// parseEditorJSON parses editorJS data
func parseEditorJSON(editorJSData string) (*EditorJS, error) {
	result := &EditorJS{}
	err := json.Unmarshal([]byte(editorJSData), result)
	if err != nil {
		return nil, err
	}
	return result, err
}
