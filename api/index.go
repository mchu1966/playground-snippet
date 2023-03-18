package handler

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

var (
	router = bunrouter.New()
)

func routers(g *bunrouter.Group) {
	g.GET("/users/:id", debugHandler)
	g.GET("/users/current", debugHandler)
	g.GET("/users/*path", debugHandler)
}

func init() {
	router.WithGroup("/api", routers)
}

func debugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
