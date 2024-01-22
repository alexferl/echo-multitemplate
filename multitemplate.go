package multitemplate

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Render map[string]*template.Template

var (
	_ echo.Renderer = Render{}
	_ Renderer      = Render{}
)

// New instance
func New() Render {
	return make(Render)
}

// Add new template
func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	if _, ok := r[name]; ok {
		panic(fmt.Sprintf("template %s already exists", name))
	}
	r[name] = tmpl
}

// AddFromFiles supply add template from files
func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromGlob supply add template from global path
func (r Render) AddFromGlob(name, glob string) *template.Template {
	tmpl := template.Must(template.ParseGlob(glob))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromString supply add template from strings
func (r Render) AddFromString(name, templateString string) *template.Template {
	tmpl := template.Must(template.New(name).Parse(templateString))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromStringsFuncs supply add template from strings
func (r Render) AddFromStringsFuncs(name string, funcMap template.FuncMap, templateStrings ...string) *template.Template {
	tmpl := template.New(name).Funcs(funcMap)

	for _, ts := range templateStrings {
		tmpl = template.Must(tmpl.Parse(ts))
	}

	r.Add(name, tmpl)
	return tmpl
}

// AddFromFilesFuncs supply add template from file callback func
func (r Render) AddFromFilesFuncs(name string, funcMap template.FuncMap, files ...string) *template.Template {
	tname := filepath.Base(files[0])
	tmpl := template.Must(template.New(tname).Funcs(funcMap).ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

func (r Render) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var t *template.Template
	t = r[name]
	return t.Execute(w, data)
}
