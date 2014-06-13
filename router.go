package squish

import (
	"fmt"
	"net/http"
)

// module reference to the router
var router *Router

type Router struct {
	routes []*Route
}

func init() {
	router = &Router{
		routes: make([]*Route, 0),
	}
}

func New() *Router {
	return router
}

func (r *Router) ServeHTTP(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	for _, route := range r.routes {
		if route.Match(NewRequest(httpRequest)) {
			fmt.Fprintf(httpResponseWriter, httpRequest.RequestURI)
			fmt.Println(httpRequest.RequestURI)
		}
	}
}

func (r *Router) AddRoute(f RouteFunc) {
	r.routes = append(r.routes, NewRoute(f))
}

func (r *Router) Run() {
	http.ListenAndServe(":8080", r)
}

// run router
func Run() {
	router.Run()
}
