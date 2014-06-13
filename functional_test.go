package squish

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestAnd(t *testing.T) {
	res := And(MockRequest(), True, True, True)
	assert.True(t, res)

	res = And(MockRequest(), True, False, True)
	assert.False(t, res)
}

func TestOr(t *testing.T) {
	res := Or(MockRequest(), True, True, True)
	assert.True(t, res)

	res = Or(MockRequest(), True, False, True)
	assert.True(t, res)

	res = Or(MockRequest(), False, False, False)
	assert.False(t, res)
}

func TestEvery(t *testing.T) {
	res := Every(MockRequest(), True, True, True)
	assert.True(t, res)

	res = Every(MockRequest(), True, False, True)
	assert.False(t, res)
}

func TestSome(t *testing.T) {
	res := Or(MockRequest(), True, True, True)
	assert.True(t, res)

	res = Or(MockRequest(), True, False, True)
	assert.True(t, res)

	res = Or(MockRequest(), False, False, False)
	assert.False(t, res)
}

func TestNone(t *testing.T) {
	res := None(MockRequest(), True, True, True)
	assert.False(t, res)

	res = None(MockRequest(), True, False, True)
	assert.False(t, res)

	res = None(MockRequest(), False, False, False)
	assert.True(t, res)
}
