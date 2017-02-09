package main

import (
	"fmt"
	"testing"

	"github.com/daidodo/testa/assert"
)

func TestSlice(t *testing.T) {
	a := []int{1, 21, 3, 4, 5}
	b := a
	assert.Equal(t, b, a)
	c := []int{1, 2, 3, 4, 6, 7, 8, 10, 11, 12, 13}
	assert.Equal(t, c, a, "a=%T(%v) is not %T(%v)", a, a, c, c)
}

func TestArray(t *testing.T) {
	a := [...]int{1, 21, 3, 4, 5}
	b := a
	assert.Equal(t, b, a)
	c := [...]int{1, 2, 3, 4, 6}
	assert.Equal(t, c, a, "a=%T(%v) is not %T(%v)", a, a, c, c)
}

func TestInt(t *testing.T) {
	var a int
	a = 'A'
	fmt.Printf("\t%v\n", a)
	assert.Equal(t, 'A', a, "a=%v is not 'A'", a)
}

func TestMap(t *testing.T) {
	m1 := make(map[int]string)
	m2 := make(map[uint]string)
	var s1, s2 []rune
	for i := 0; i < 3; i++ {
		s1 = append(s1, rune('a'+i))
		s2 = append(s2, rune('b'+i))
		m1[i] = string(s1)
		m2[uint(i)] = string(s2)
	}
	assert.Equal(t, m1, m2)
}
