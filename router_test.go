package squish

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	assert "github.com/pilu/miniassert"

	"testing"
)

func TestNew(t *testing.T) {
	r := New()

	assert.Type(t, "*squish.Router", r)
	assert.Type(t, "[]*squish.Route", r.routes)

	assert.Equal(t, 0, len(r.routes))

	emptyArray := make([]*Route, 0)
	assert.Equal(t, emptyArray, r.routes)
}

func TestRouter_AddRoute(t *testing.T) {
	r := New()

	r.AddRoute(F)
	assert.Equal(t, 1, len(r.routes))
	assert.Type(t, "*squish.Route", r.routes[0])
}

func TestRouter_ServeHTTP(t *testing.T) {
	r := New()
	request, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	assert.Equal(t, "", string(recorder.Body.Bytes()))
	assert.Equal(t, http.StatusOK, recorder.Code)

	matcher := func(req *Request) bool {
		return req.URL.Path == "/test"
	}

	handler := func(req *Request, res Response) {
		fmt.Fprint(res, "test body")
	}

	r.AddRoute(matcher, handler)

	request, _ = http.NewRequest("GET", "/test", nil)
	recorder = httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
	assert.Equal(t, "test body", string(recorder.Body.Bytes()))
}
