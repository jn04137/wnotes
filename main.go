package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"com/jn04137/wnotes/views/layout"
	"com/jn04137/wnotes/views/editor"
)

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return renderView(c, layout.Layout(layout.Content("Something here")))
	})
	e.GET("/editor", func(c echo.Context) error {
		return renderView(c, layout.Layout(editor.Editor()))
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
