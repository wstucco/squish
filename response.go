package squish

import "net/http"

type Response interface {
	http.ResponseWriter
}

type response struct {
	http.ResponseWriter
}

func NewResponse(r http.ResponseWriter) *response {
	return &response{
		ResponseWriter: r,
	}
}

func (r *response) Write(data []byte) (n int, err error) {
	// w.written = true
	// w.bodyWritten = true

	return r.ResponseWriter.Write(data)
}
