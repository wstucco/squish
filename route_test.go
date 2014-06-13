package squish

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

var F = func(*Request) bool {
	return true
}

var timesHandlerIsExecute = 0
var Handler = func(req *Request, res Response) {
	timesHandlerIsExecute++
}

var handlers = []HttpHandlerFunc{Handler, Handler, Handler}

func TestNewRoute(t *testing.T) {
	r := NewRoute(F, handlers...)
	assert.Type(t, "*squish.Route", r)
	assert.Type(t, "squish.RouteFunc", r.match)
	assert.Equal(t, 3, len(r.handlers))
}

func TestRoute_Match(t *testing.T) {
	r := NewRoute(F, handlers...)
	req := MockRequest()

	assert.Equal(t, r.Match(req), F(req))
	assert.True(t, r.Match(req))
}

func TestRoute_Execute(t *testing.T) {
	r := NewRoute(F, handlers...)

	r.Execute(nil, nil)
	assert.Equal(t, 3, timesHandlerIsExecute)
}
