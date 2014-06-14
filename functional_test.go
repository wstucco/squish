package squish

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestAnd(t *testing.T) {
	f := And(True, True, True)
	assert.True(t, f(MockRequest()))

	f = And(True, False, True)
	assert.False(t, f(MockRequest()))
}

func TestOr(t *testing.T) {
	f := Or(True, True, True)
	assert.True(t, f(MockRequest()))

	f = Or(True, False, True)
	assert.True(t, f(MockRequest()))

	f = Or(False, False, False)
	assert.False(t, f(MockRequest()))
}

func TestEvery(t *testing.T) {
	f := Every(True, True, True)
	assert.True(t, f(MockRequest()))

	f = Every(True, False, True)
	assert.False(t, f(MockRequest()))
}

func TestSome(t *testing.T) {
	f := Or(True, True, True)
	assert.True(t, f(MockRequest()))

	f = Or(True, False, True)
	assert.True(t, f(MockRequest()))

	f = Or(False, False, False)
	assert.False(t, f(MockRequest()))
}

func TestNone(t *testing.T) {
	f := None(True, True, True)
	assert.False(t, f(MockRequest()))

	f = None(True, False, True)
	assert.False(t, f(MockRequest()))

	f = None(False, False, False)
	assert.True(t, f(MockRequest()))
}

func TestNot(t *testing.T) {
	f := Not(None(True, True, True))
	assert.True(t, f(MockRequest()))

	f = Not(Every(True, False, True))
	assert.True(t, f(MockRequest()))

	f = Not(Or(False, False, False))
	assert.True(t, f(MockRequest()))
}
