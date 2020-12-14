package goeditorjs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseEditorJSON(t *testing.T) {
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	editorJS, err := parseEditorJSON(editorJSData)
	require.NoError(t, err)
	require.Len(t, editorJS.Blocks, 1)
}

func Test_parseEditorJSON_Err_Empty(t *testing.T) {
	editorJSData := ``

	_, err := parseEditorJSON(editorJSData)
	require.Error(t, err)
}
