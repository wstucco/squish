package squish

var mockRequest *Request

func MockRequest() *Request {
	if mockRequest == nil {
		mockRequest = NewRequest(nil)
	}

	return mockRequest
}
