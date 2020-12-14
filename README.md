# davidscottmills/goeditorjs

An extensible library that converts [editor.js](https://editorjs.io/) data into HTML or markdown.

## Installation

```bash
go get github.com/davidscottmills/goeditorjs
```

## Usage

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/davidscottmills/goeditorjs"
)

func main() {
	content, err := ioutil.ReadFile("editorjs_output.json")
	if err != nil {
		log.Fatal(err)
	}

	ejs := string(content)

    // HTML
    // Get the HTML engine
    htmlEngine := goeditorjs.NewHTMLEngine()
    // Register the handlers you wish to use
	htmlEngine.RegisterBlockHandlers(
		&goeditorjs.BlockHeaderHandler{},
		&goeditorjs.BlockParagraphHandler{},
		&goeditorjs.BlockListHandler{},
		&goeditorjs.BlockCodeBoxHandler{},
	)
    // Generate the html
	html, err := htmlEngine.GenerateHTML(ejs)
	if err != nil {
		log.Fatal(err)
    }

    // Do something with the html output. In this case, write it to a file.
	err = ioutil.WriteFile("editorjs.html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

    // Generate markdown and save it to a file
    // Get the markdown engine
	markdownEngine := goeditorjs.NewMarkdownEngine()
    // Register the handlers you wish to use
	markdownEngine.RegisterBlockHandlers(
		&goeditorjs.BlockHeaderHandler{},
		&goeditorjs.BlockParagraphHandler{},
		&goeditorjs.BlockListHandler{},
		&goeditorjs.BlockCodeBoxHandler{},
    )
    // Generate the markdown
	md, err := markdownEngine.GenerateMarkdown(ejs)
	if err != nil {
		log.Fatal(err)
	}

    // Do something with the md output. In this case, write it to a file.
	err = ioutil.WriteFile("editorjs.md", []byte(md), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Using a Custom Handler

You can create and use your own handler in either engine by implementing the required interface and registering it.
This package provides two interfaces for handlers.

- `HTMLBlockHandler`

  ```go
  type MarkdownBlockHandler interface {
      Type() string // Type returns the type the block handler supports as a string
      GenerateHTML(editorJSBlock EditorJSBlock) (string, error) // Return associated HTML
  }
  ```

- `MarkdownBlockHandler`
  ```go
  type MarkdownBlockHandler interface {
      Type() string // Type returns the type the block handler supports as a string
      GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) // Return associated markdown
  }
  ```

If you're only planning to use the HTMLEngine, then you only need to implement the `HTMLBlockHandler` interface. The same goes for markdown.

Once you've met the required interface, register the handler for use in the engine.

```go
htmlEngine := goeditorjs.NewHTMLEngine()
// Register the handlers you wish to use
htmlEngine.RegisterBlockHandlers(
    &MyCustomBlockHandler{},
)
```

Below is an example of how the header handle is implemented.

```go
package header

import (
	"encoding/json"
	"fmt"
)

// BlockHeaderHandler is the default BlockHeaderHandler for EditorJS HTML generation
type BlockHeaderHandler struct {
    // Notice that you could put some configurable options in this struct and then use them in your handler
}

// Header represents header data from EditorJS
type Header struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

func (*BlockHeaderHandler) parse(editorJSBlock EditorJSBlock) (*Header, error) {
	header := &Header{}
	return header, json.Unmarshal(editorJSBlock.Data, header)
}

// Type "header"
func (*BlockHeaderHandler) Type() string {
	return "header"
}

// GenerateHTML generates html for HeaderBlocks
func (h *BlockHeaderHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<h%d>%s</h%d>", header.Level, header.Text, header.Level), nil
}

// GenerateMarkdown generates markdown for HeaderBlocks
func (h *BlockHeaderHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", strings.Repeat("#", header.Level), header.Text), nil
}
```

## TODO
- The markdown output for codebox is really messy HTML and doesn't display correctly. Fix this.