package squish

import "net/http"

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

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	request := NewRequest(req)
	response := NewResponse(res)
	for _, route := range r.routes {
		if route.Match(request) {
			route.Execute(request, response)
		}
	}
}

func (r *Router) AddRoute(f RouteFunc, handlers ...HttpHandlerFunc) {
	r.routes = append(r.routes, NewRoute(f, handlers...))
}

func (r *Router) Run() {
	http.ListenAndServe(":8080", r)
}

// run router
func Run() {
	router.Run()
}
