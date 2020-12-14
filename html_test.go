package goeditorjs_test

import (
	"errors"
	"testing"

	"github.com/davidscottmills/goeditorjs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockBlockHandler struct {
	mock.Mock
	typeName string
}

func (m *mockBlockHandler) GenerateHTML(editorJSBlock goeditorjs.EditorJSBlock) (string, error) {
	args := m.Called(editorJSBlock)
	return args.String(0), args.Error(1)
}

func (m *mockBlockHandler) Type() string {
	return m.typeName
}

func Test_ParseEditorJSON(t *testing.T) {
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	editorJS, err := goeditorjs.ParseEditorJSON(editorJSData)
	require.NoError(t, err)
	require.Len(t, editorJS.Blocks, 1)
}

func Test_ParseEditorJSON_Err_Empty(t *testing.T) {
	editorJSData := ``

	_, err := goeditorjs.ParseEditorJSON(editorJSData)
	require.Error(t, err)
}

func Test_NewHTMLEngine(t *testing.T) {
	eng := goeditorjs.NewHTMLEngine()
	require.NotNil(t, eng)
	require.NotNil(t, eng.BlockHandlers)
}

func Test_RegisterBlockHandler(t *testing.T) {
	bh := &mockBlockHandler{typeName: "header"}
	eng := &goeditorjs.HTMLEngine{BlockHandlers: make(map[string]goeditorjs.HTMLBlockHandler)}
	eng.RegisterBlockHandler(bh)
	require.Equal(t, eng.BlockHandlers["header"], bh)
}

func Test_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	eng := &goeditorjs.HTMLEngine{BlockHandlers: make(map[string]goeditorjs.HTMLBlockHandler)}
	_, err := eng.GenerateHTML(``)
	require.Error(t, err)
}

func Test_GenerateHTML_NoHandler_Should_Err(t *testing.T) {
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.HTMLEngine{BlockHandlers: make(map[string]goeditorjs.HTMLBlockHandler)}
	_, err := eng.GenerateHTML(editorJSData)
	require.Error(t, err)
	require.True(t, errors.Is(err, goeditorjs.ErrBlockHandlerNotFound))
}

func Test_GenerateHTML_Returns_Err_From_Handler(t *testing.T) {
	bh := &mockBlockHandler{}
	mockErr := errors.New("Mock Error")
	bh.On("GenerateHTML", mock.Anything).Return("", mockErr)
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.HTMLEngine{BlockHandlers: make(map[string]goeditorjs.HTMLBlockHandler)}
	eng.BlockHandlers["header"] = bh
	_, err := eng.GenerateHTML(editorJSData)
	require.Error(t, err)
	require.Equal(t, mockErr, err)
	bh.AssertCalled(t, "GenerateHTML", mock.Anything)
}

func Test_GenerateHTML_Result_Includes_Handler_Result(t *testing.T) {
	bh := &mockBlockHandler{}
	handlerResult := "<h1>Hello World</h1>"
	bh.On("GenerateHTML", mock.Anything).Return(handlerResult, nil)
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.HTMLEngine{BlockHandlers: make(map[string]goeditorjs.HTMLBlockHandler)}
	eng.BlockHandlers["header"] = bh
	result, _ := eng.GenerateHTML(editorJSData)
	require.Contains(t, result, handlerResult)
	bh.AssertCalled(t, "GenerateHTML", mock.Anything)
}
