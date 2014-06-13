package squish

import (
	assert "github.com/pilu/miniassert"

	"testing"
)

func TestNew(t *testing.T) {
	r := New()

	assert.Type(t, "*squish.Router", r)
	assert.Type(t, "[]*squish.Route", r.routes)

	assert.Equal(t, 0, len(router.routes))

	emptyArray := make([]*Route, 0)
	assert.Equal(t, emptyArray, router.routes)
}

func TestRouter_AddRoute(t *testing.T) {
	r := New()

	r.AddRoute(F)
	assert.Equal(t, 1, len(router.routes))
	assert.Type(t, "*squish.Route", router.routes[0])
}
