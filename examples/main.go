package main

import (
	"io/ioutil"
	"log"

	"github.com/davidscottmills/goeditorjs"
)

func main() {
	content, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}

	ejs := string(content)

	// Generate HTML and save it to a file
	htmlEngine := goeditorjs.NewHTMLEngine()
	htmlEngine.RegisterBlockHandlers(
		&goeditorjs.HeaderHandler{},
		&goeditorjs.ParagraphHandler{},
		&goeditorjs.ListHandler{},
		&goeditorjs.CodeBoxHandler{},
	)
	html, err := htmlEngine.GenerateHTML(ejs)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("editorjs.html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Generate markdown and save it to a file
	markdownEngine := goeditorjs.NewMarkdownEngine()
	markdownEngine.RegisterBlockHandlers(
		&goeditorjs.HeaderHandler{},
		&goeditorjs.ParagraphHandler{},
		&goeditorjs.ListHandler{},
		&goeditorjs.CodeBoxHandler{},
	)
	md, err := markdownEngine.GenerateMarkdown(ejs)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("editorjs.md", []byte(md), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
