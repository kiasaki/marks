package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/kiasaki/batbelt/mst"
	"github.com/russross/blackfriday"
)

const layoutContents = `
{{define "layout"}}
<!DOCTYPE html>
<html>
  <head>
	<title>Marks - {{.Title}}</title>
	<link rel="stylesheet" href="/styles.css" />
  </head>
  <body>
	<nav>
	  <header>
		<h1>Marks</h1>
		<a href="/new">New note</a>
	  </header>
	  <ul>
		{{$id := .Id}}
		{{range .Notes}}
		  <li class="{{if eq $id .Id}}active{{end}}">
			<a href="/n/{{.Id}}">{{.Title}}</a>
		  </li>
		{{end}}
	  </ul>
	  <footer>
		<a href="http://github.com/kiasaki/marks">Marks</a>
		is open source software.
	  </footer>
	</nav>

	<section>
	  <form action="{{.PostbackURL}}" method="post">
		<header>
		  <input type="text" name="title" class="title" value="{{.Title}}" placeholder="Note title" />
		  <div class="buttons">
			{{template "buttons" .}}
		  </div>
		</header>
		<div class="contents">
		  {{template "contents" .}}
		<div>
	  </form>
	</section>
  </body>
</html>
{{end}}
`

const editorContents = `
  <textarea id="noteBody" name="body">{{.Body}}</textarea>
  <link href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/4.12.0/codemirror.min.css" rel="stylesheet" />
  <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/4.12.0/codemirror.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/4.12.0/mode/markdown/markdown.min.js"></script>
  <script>
    var editor = CodeMirror.fromTextArea(noteBody, {
	  lineNumbers: true,
	  lineWrapping: true,
	  autofocus: true,
	  viewportMargin: Infinity,
	  mode: 'markdown',
	});
  </script>
`

const newContents = `
{{define "buttons"}}
  <button type="submit" class="btn">Save</button>
{{end}}
{{define "contents"}}
` + editorContents + `
{{end}}
`

const viewContents = `
{{define "buttons"}}
  <a href="/n/{{.Id}}/delete" class="btn btn-default" id="deleteBtn">Delete</a>
  <a href="/n/{{.Id}}/edit" class="btn">Edit</a>
{{end}}
{{define "contents"}}
  <div class="body">
	%s
  </div>
  <script>
	deleteBtn.onclick = function(e) {
	  (!confirm('Are you sure?')) && e.preventDefault();
	};
  </script>
{{end}}
`

const editContents = `
{{define "buttons"}}
  <a href="/n/{{.Id}}" class="btn btn-default">Cancel</a>
  <button type="submit" class="btn">Save</button>
{{end}}
{{define "contents"}}
` + editorContents + `
{{end}}
`

var newPageTempate *template.Template
var editPageTempate *template.Template

func init() {
	var err error
	newPageTempate, err = loadLayoutTemplate().Parse(newContents)
	mst.MustNotErr(err)
	editPageTempate, err = loadLayoutTemplate().Parse(editContents)
	mst.MustNotErr(err)
}

func loadLayoutTemplate() *template.Template {
	t, err := template.New("page").Parse(layoutContents)
	mst.MustNotErr(err)
	return t
}

func executeTemplate(t *template.Template, p Page) ([]byte, error) {
	var out bytes.Buffer
	err := t.ExecuteTemplate(&out, "layout", p)
	if err != nil {
		return []byte{}, err
	}
	return out.Bytes(), nil
}

func RenderNewPage(p Page) ([]byte, error) {
	return executeTemplate(newPageTempate, p)
}

func RenderViewPage(p Page) ([]byte, error) {
	t := loadLayoutTemplate()
	body := string(blackfriday.MarkdownCommon([]byte(p.Body)))
	t, err := t.Parse(fmt.Sprintf(viewContents, body))
	if err != nil {
		return []byte{}, err
	}
	return executeTemplate(t, p)
}

func RenderEditPage(p Page) ([]byte, error) {
	return executeTemplate(editPageTempate, p)
}
