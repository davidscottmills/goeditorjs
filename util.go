package goeditorjs

import "encoding/json"

// parseEditorJSON parses editorJS data
func parseEditorJSON(editorJSData string) (*editorJS, error) {
	result := &editorJS{}
	err := json.Unmarshal([]byte(editorJSData), result)
	if err != nil {
		return nil, err
	}
	return result, err
}
