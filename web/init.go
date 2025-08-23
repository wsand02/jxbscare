package web

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/wsand02/jxbscare/web/routes"
)

type Web struct {
	router     *gin.Engine
	routesInit bool
	f          embed.FS
}

//go:embed templates/*
var embedFS embed.FS

func New() *Web {
	w := initWeb()
	w.f = embedFS
	w.initTempl()
	w.initRoutes()
	return w
}

func initWeb() *Web {
	router := gin.Default()
	w := Web{router: router}
	w.routesInit = false
	return &w
}

func (w *Web) initTempl() {
	templ := template.Must(template.New("").ParseFS(w.f, "templates/*.tmpl"))
	w.router.SetHTMLTemplate(templ)
}

func (w *Web) initRoutes() {
	routes.InitRoutes(w.router)
	w.routesInit = true
}
