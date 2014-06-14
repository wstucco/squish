package squish

// And and Or are aliases for Every and Some
func And(matchers ...RouteFunc) RouteFunc {
	return Every(matchers...)
}

func Or(matchers ...RouteFunc) RouteFunc {
	return Some(matchers...)
}

func Every(matchers ...RouteFunc) RouteFunc {
	return func(r *Request) bool {
		for _, f := range matchers {
			if !f(r) {
				return false
			}
		}

		return true
	}
}

func Some(matchers ...RouteFunc) RouteFunc {
	return func(r *Request) bool {
		for _, f := range matchers {
			if f(r) {
				return true
			}
		}

		return false
	}
}

// None is the opposite of Some
func None(matchers ...RouteFunc) RouteFunc {
	return Not(Some(matchers...))
}

func Not(matcher RouteFunc) RouteFunc {
	return func(r *Request) bool {
		return !matcher(r)
	}
}
