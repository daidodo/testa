package main

import (
	"testing"

	"github.com/daidodo/testa/assert"
)

func myTest(t *testing.T, e, a int) {
	assert.Caller(1).Equal(t, e, a)
}

func TestA(t *testing.T) {
	myTest(t, 1, 1)
	myTest(t, 10, 10)
	myTest(t, 100, 101)
}
