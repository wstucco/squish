package squish

import (
	"net/http"
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestRequest(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "http://example.com", nil)
	r := NewRequest(httpRequest)

	assert.Equal(t, r.Request, httpRequest)
}
