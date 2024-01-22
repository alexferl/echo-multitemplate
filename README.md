# echo-multitemplate [![Go Report Card](https://goreportcard.com/badge/github.com/alexferl/echo-multitemplate)](https://goreportcard.com/report/github.com/alexferl/echo-multitemplate) [![codecov](https://codecov.io/gh/alexferl/echo-multitemplate/branch/master/graph/badge.svg)](https://codecov.io/gh/alexferl/echo-multitemplate)

This is a custom HTML renderer to support multiple templates, ie. more than one `*template.Template` for the
[Echo](https://github.com/labstack/echo) framework

## Installing
```shell
go get github.com/alexferl/echo-multitemplate
```

### Code example
See [examples/simple/example.go](examples/simple/example.go)

```go
package main

import (
	"net/http"

	"github.com/alexferl/echo-multitemplate"
	"github.com/labstack/echo/v4"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	e := echo.New()
	e.Renderer = createMyRender()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{"title": "Index"})
	})

	e.GET("/article", func(c echo.Context) error {
		return c.Render(http.StatusOK, "article", echo.Map{"title": "Article"})
	})

	e.Logger.Fatal(e.Start("localhost:1323"))
}
```

## Credits
Port of [gin-contrib/multitemplate](https://github.com/gin-contrib/multitemplate).
