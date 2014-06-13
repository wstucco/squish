package squish

// And and Or are aliases for Every and Some
func And(r *Request, matchers ...RouteFunc) bool {
	return Every(r, matchers...)
}

func Or(r *Request, matchers ...RouteFunc) bool {
	return Some(r, matchers...)
}

func Every(r *Request, matchers ...RouteFunc) bool {
	for _, f := range matchers {
		if !f(r) {
			return false
		}
	}

	return true
}

func Some(r *Request, matchers ...RouteFunc) bool {
	for _, f := range matchers {
		if f(r) {
			return true
		}
	}

	return false
}

// None is the opposite of Some
func None(r *Request, matchers ...RouteFunc) bool {
	return !Some(r, matchers...)
}
