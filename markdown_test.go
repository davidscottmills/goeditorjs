package goeditorjs_test

import (
	"errors"
	"testing"

	"github.com/davidscottmills/goeditorjs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockMarkdownBlockHandler struct {
	mock.Mock
	typeName string
}

func (m *mockMarkdownBlockHandler) GenerateMarkdown(editorJSBlock goeditorjs.EditorJSBlock) (string, error) {
	args := m.Called(editorJSBlock)
	return args.String(0), args.Error(1)
}

func (m *mockMarkdownBlockHandler) Type() string {
	return m.typeName
}

func Test_Markdown(t *testing.T) {
	eng := goeditorjs.NewMarkdownEngine()
	require.NotNil(t, eng)
	require.NotNil(t, eng.BlockHandlers)
}

func Test_MarkdownEngine_RegisterBlockHandler(t *testing.T) {
	bh1 := &mockMarkdownBlockHandler{typeName: "header"}
	bh2 := &mockMarkdownBlockHandler{typeName: "list"}
	eng := &goeditorjs.MarkdownEngine{BlockHandlers: make(map[string]goeditorjs.MarkdownBlockHandler)}
	eng.RegisterBlockHandlers(bh1, bh2)
	require.Equal(t, eng.BlockHandlers["header"], bh1)
	require.Equal(t, eng.BlockHandlers["list"], bh2)
}

func Test_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	eng := &goeditorjs.MarkdownEngine{BlockHandlers: make(map[string]goeditorjs.MarkdownBlockHandler)}
	_, err := eng.GenerateMarkdown(``)
	require.Error(t, err)
}

func Test_GenerateMarkdown_NoHandler_Should_Err(t *testing.T) {
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.MarkdownEngine{BlockHandlers: make(map[string]goeditorjs.MarkdownBlockHandler)}
	_, err := eng.GenerateMarkdown(editorJSData)
	require.Error(t, err)
	require.True(t, errors.Is(err, goeditorjs.ErrBlockHandlerNotFound))
}

func Test_GenerateMarkdown_Returns_Err_From_Handler(t *testing.T) {
	bh := &mockMarkdownBlockHandler{}
	mockErr := errors.New("Mock Error")
	bh.On("GenerateMarkdown", mock.Anything).Return("", mockErr)
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.MarkdownEngine{BlockHandlers: make(map[string]goeditorjs.MarkdownBlockHandler)}
	eng.BlockHandlers["header"] = bh
	_, err := eng.GenerateMarkdown(editorJSData)
	require.Error(t, err)
	require.Equal(t, mockErr, err)
	bh.AssertCalled(t, "GenerateMarkdown", mock.Anything)
}

func Test_GenerateMarkdown_Result_Includes_Handler_Result(t *testing.T) {
	bh := &mockMarkdownBlockHandler{}
	handlerResult := "# Hello World"
	bh.On("GenerateMarkdown", mock.Anything).Return(handlerResult, nil)
	editorJSData := `{"time": 1607709186831,"blocks": [{"type": "header","data": {"text": "Heading 1","level": 1}}],"version": "2.19.1"}`
	eng := &goeditorjs.MarkdownEngine{BlockHandlers: make(map[string]goeditorjs.MarkdownBlockHandler)}
	eng.BlockHandlers["header"] = bh
	result, _ := eng.GenerateMarkdown(editorJSData)
	require.Contains(t, result, handlerResult)
	bh.AssertCalled(t, "GenerateMarkdown", mock.Anything)
}
