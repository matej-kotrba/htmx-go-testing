package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
    tmpl *template.Template
}

func newTemplate() *Template {
    return &Template{
        tmpl: template.Must(template.ParseGlob("views/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.tmpl.ExecuteTemplate(w, name, data)
}

type Count struct {
    Count int
}

func main() {
	count := Count{Count: 0}

	e := echo.New()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/add", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.Logger.Fatal(e.Start(":1323"))
}