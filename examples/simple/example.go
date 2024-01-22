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
