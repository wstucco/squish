package squish

import (
	"fmt"
	"net/http"
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestRequest(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "http://example.com", nil)
	r := NewRequest(httpRequest)

	assert.Equal(t, r.Request, httpRequest)
}

func TestRequest_Path(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "http://example.com/test/path", nil)
	r := NewRequest(httpRequest)

	assert.Equal(t, r.Path(), "/test/path")
}

func TestRequest_Match(t *testing.T) {
	tests := [][]string{
		{
			"/",
			"/",
			"",
		},
		{
			"/:foo/?",
			"/bar",
			"foo=bar",
		},
		{
			"/:foo/?",
			"/bar/", // with trailing slash
			"foo=bar",
		},
		{
			"/categories/:category_id/posts/:id",
			"/categories/foo/posts/bar",
			"category_id=foo&id=bar",
		},
		{
			"/pages/:page_path*",
			"/pages/foo/bar/baz",
			"page_path=foo%2Fbar%2Fbaz",
		},
		{
			"/pages/:page.html",
			"/pages/foo.html",
			"page=foo",
		},
	}

	for _, opts := range tests {
		routePath := opts[0]
		requestPath := opts[1]
		expectedQuery := opts[2]

		httpRequest, _ := http.NewRequest("GET", fmt.Sprintf("http://example.com%s", requestPath), nil)
		r := NewRequest(httpRequest)

		assert.Equal(t, r.Path(), requestPath)

		match, values := r.Match(routePath)
		assert.True(t, match)
		assert.Equal(t, values.Encode(), expectedQuery)
	}
}
