package main

import (
	// "fmt"
	"html/template"
	"io"

	// "net/http"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	template *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		template: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	count int
}

func main() {

	e := echo.New()
	// e.Use(middleware.Logger())
	count := Count{count: 0}
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		count.count++
		return c.Render(200, "index", count)
	})

	e.Logger.Fatal(e.Start(":3000"))
	// fmt.Println("hello world")

	// http.HandleFunc("/",handler)

	// if err := http.ListenAndServe(":4000",nil); err != nil{
	// 	fmt.Println(err)
	// }
}

// func handler(w http.ResponseWriter, r *http.Request){
// 	fmt.Printf("om")
// }
