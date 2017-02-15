package reflect

import (
	"reflect"
	"testing"

	"github.com/daidodo/testa/assert"
	assert2 "github.com/stretchr/testify/assert"
)

type Kind uint

//func TestWriteDiffPkgTypes(t *testing.T) {
//    cs := []struct {
//        v1, v2 reflect.Value
//        s1, s2 string
//    }{
//        {v1: reflect.ValueOf(reflect.Int), v2: reflect.ValueOf(Kind(100)), s1: "reflect.Kind(int)", s2: "\x1b[41mexample\x1b[0m/reflect.Kind(100)"},
//    }
//    for i, c := range cs {
//        var d assert.tValueDiffer
//        d.WriteDiff(c.v1, c.v2, 0)
//        assert.Equal(t, c.s1, d.String(0), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
//        assert.Equal(t, c.s2, d.String(1), "i=%v, s2\n%v\n%v", i, d.String(0), d.String(1))
//    }
//}

func TestWriteDiffPkgTypes1(t *testing.T) {
	assert.Equal(t, reflect.Int, Kind(100), "Msg")
}

func TestWriteDiffPkgTypes2(t *testing.T) {
	assert2.Equal(t, reflect.Int, Kind(100))
}
