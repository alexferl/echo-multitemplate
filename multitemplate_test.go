package multitemplate

import (
	"context"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequestWithContext(context.Background(), method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createFromFile() Render {
	r := New()
	r.AddFromFiles("index", "fixtures/base.html", "fixtures/article.html")

	return r
}

func createFromGlob() Render {
	r := New()
	r.AddFromGlob("index", "fixtures/global/*")

	return r
}

func createFromString() Render {
	r := New()
	r.AddFromString("index", "Welcome to {{ .name }} template")

	return r
}

func createFromStringsWithFuncs() Render {
	r := New()
	r.AddFromStringsFuncs("index", template.FuncMap{}, `Welcome to {{ .name }} {{template "content"}}`, `{{define "content"}}template{{end}}`)

	return r
}

func createFromFilesWithFuncs() Render {
	r := New()
	r.AddFromFilesFuncs("index", template.FuncMap{}, "fixtures/welcome.html", "fixtures/content.html")

	return r
}

func TestMissingTemplateOrName(t *testing.T) {
	r := New()
	tmpl := template.Must(template.New("test").Parse("Welcome to {{ .name }} template"))
	assert.Panics(t, func() {
		r.Add("", tmpl)
	}, "template name cannot be empty")

	assert.Panics(t, func() {
		r.Add("test", nil)
	}, "template can not be nil")
}

func TestAddFromFiles(t *testing.T) {
	e := echo.New()
	e.Renderer = createFromFile()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"title": "Test Multiple Template"})
	})

	w := performRequest(e, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is article template\n", w.Body.String())
}

func TestAddFromGlob(t *testing.T) {
	e := echo.New()
	e.Renderer = createFromGlob()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"title": "Test Multiple Template"})
	})

	w := performRequest(e, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is login template\n", w.Body.String())
}

func TestAddFromString(t *testing.T) {
	e := echo.New()
	e.Renderer = createFromString()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"name": "index"})
	})

	w := performRequest(e, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromStringsFruncs(t *testing.T) {
	e := echo.New()
	e.Renderer = createFromStringsWithFuncs()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"name": "index"})
	})

	w := performRequest(e, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromFilesFruncs(t *testing.T) {
	e := echo.New()
	e.Renderer = createFromFilesWithFuncs()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"name": "index"})
	})

	w := performRequest(e, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template\n", w.Body.String())
}

func TestDuplicateTemplate(t *testing.T) {
	assert.Panics(t, func() {
		r := New()
		r.AddFromString("index", "Welcome to {{ .name }} template")
		r.AddFromString("index", "Welcome to {{ .name }} template")
	})
}
