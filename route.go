package squish

type HttpMethod string

type HttpHandlerFunc func(*Request, Response)

type RouteFunc func(*Request) bool

type Route struct {
	match    RouteFunc
	handlers []HttpHandlerFunc
}

func NewRoute(f RouteFunc, handlers ...HttpHandlerFunc) *Route {
	return &Route{
		match:    f,
		handlers: handlers,
	}
}

func (r *Route) Match(req *Request) bool {
	return r.match(req)
}

func (r *Route) Execute(req *Request, res Response) error {

	for _, handler := range r.handlers {
		handler(req, res)
	}

	return nil
}
