package squish

import (
	"net/http"
	"net/url"
	"regexp"
)

type Request struct {
	*http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: r,
	}
}

func (r *Request) Path() string {
	return r.Request.URL.Path
}

func (r *Request) Match(pattern string) (bool, url.Values) {
	values := make(url.Values)

	matcher := regexp.MustCompile(pathToRegexpString(pattern))
	matches := matcher.FindAllStringSubmatch(r.Path(), -1)
	if matches != nil {
		names := matcher.SubexpNames()
		for i := 1; i < len(names); i++ {
			name := names[i]
			value := matches[0][i]
			if len(name) > 0 && len(value) > 0 {
				values.Set(name, value)
			}
		}

		return true, values
	}

	return false, values

}
