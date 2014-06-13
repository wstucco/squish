package squish

var mockRequest *Request

func MockRequest() *Request {
	if mockRequest == nil {
		mockRequest = NewRequest(nil)
	}

	return mockRequest
}

var F = func(*Request) bool {
	return true
}

var True = func(*Request) bool {
	return true
}

var False = func(*Request) bool {
	return false
}

var timesHandlerIsExecute = 0
var Handler = func(req *Request, res Response) {
	timesHandlerIsExecute++
}

var handlers = []HttpHandlerFunc{Handler, Handler, Handler}
