package web

import "log"

func (w *Web) StartServer() {
	if !w.routesInit {
		log.Fatal("Routes uninitialized!")
		return
	}
	err := w.router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}
