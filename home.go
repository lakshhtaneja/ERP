package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lakshhtaneja/ERP/internal"
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
	edit := internal.Peter

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/account", func(c echo.Context) error {
		return c.Render(200, "account", edit)
	})

	e.GET("/account/edit", func(c echo.Context) error {
		return c.Render(200, "edit", edit)
	})

	e.GET("/new", func(c echo.Context) error {
		return c.Render(200, "notes", nil)
	})

	e.POST("/account/edit", func(c echo.Context) error {
		edit.Name = c.FormValue("name")
		edit.Email = c.FormValue("email")
		edit.Phone = c.FormValue("phone")
		edit.Address = c.FormValue("address")
		return c.Render(200, "studentDetails", edit)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
