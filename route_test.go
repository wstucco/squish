package squish

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

var F = func(*Request) bool {
	return true
}

func TestRoute(t *testing.T) {

	r := NewRoute(F)
	assert.Type(t, "squish.RouteFunc", r.match)

	req := MockRequest()
	assert.Equal(t, r.Match(req), F(req))
}
