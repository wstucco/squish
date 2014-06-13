package squish

type RouteFunc func(*Request) bool

type Route struct {
	match RouteFunc
}

func NewRoute(f RouteFunc) *Route {
	return &Route{
		match: f,
	}
}

func (r *Route) Match(req *Request) bool {
	return r.match(req)
}
