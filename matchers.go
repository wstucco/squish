package squish

import (
	"fmt"

	"regexp"
)

// route functions shortcuts
func Url(url string) {
	router.AddRoute(func(r *Request) bool {
		return _url(url, r.RequestURI)
	})
}

// match a Url using a regex
func _url(path string, url string) bool {
	matcher := regexp.MustCompile(pathToRegexpString(path))

	matches := matcher.FindAllStringSubmatch(url, -1)
	// if matches != nil {
	// 	names := route.PathRegexp.SubexpNames()
	// 	for i := 1; i < len(names); i++ {
	// 		name := names[i]
	// 		value := matches[0][i]
	// 		if len(name) > 0 && len(value) > 0 {
	// 			values.Set(name, value)
	// 		}
	// 	}

	// 	return values, true
	// }

	// return values, false

	return matches != nil
}

func pathToRegexpString(pattern string) string {
	var re *regexp.Regexp
	regexpString := pattern

	// Dots
	re = regexp.MustCompile(`([^\\])\.`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf(`%s\.`, string(m[0]))
	})

	// Wildcard names
	re = regexp.MustCompile(`:[^/#?()\.\\]+\*`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf("(?P<%s>.+)", m[1:len(m)-1])
	})

	re = regexp.MustCompile(`:[^/#?()\.\\]+`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:len(m)])
	})

	return fmt.Sprintf(`\A%s\z`, regexpString)
}
