package main

import (
	"github.com/wsand02/jxbscare/web"
)

func main() {
	w := web.InitWeb()
	w.InitRoutes()
	w.StartServer()
}
