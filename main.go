package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {

  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", WelcomePage)
  m.Get("/todos", TodosIndexPage)
  m.Get("/todos/:id", TodosShowPage)

}

func WelcomePage(r render.Render) {
  r.HTML(200, "dummy", nil)
}

func TodosIndexPage(r render.Render) {
  r.HTML(200, "dummy", nil)
}

func TodosShowPage(r render.Render) {
  r.HTML(200, "dummy", nil)
}
