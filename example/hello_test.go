package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	m1 := make(map[int]string)
	m2 := make(map[int]string)
	var s1, s2 []rune
	for i := 0; i < 20; i++ {
		s1 = append(s1, rune('a'+i))
		s2 = append(s2, rune('b'+i))
		m1[i] = string(s1)
		m2[i] = string(s2)
	}
	assert.Equal(t, m1, m2, "You should not see this")
}
