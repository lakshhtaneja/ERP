package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Renderer = NewTemplates()
	e.Use(middleware.Logger())
	e.Static("/public", "public")

	e.Renderer = NewTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/account", func(c echo.Context) error {
		return c.Render(200, "account", nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
