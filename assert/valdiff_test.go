package assert

import (
	"reflect"
	"testing"
)

func TestWriteDiffStringValue(t *testing.T) {
	eq := func(x1, x2 string, s1, s2 string) {
		var d ValueDiffer
		d.writeDiffValuesString(reflect.ValueOf(x1), reflect.ValueOf(x2))
		Caller(1).Equal(t, s1, d.String(0), "%s\n%s", d.String(0), d.String(1))
		Caller(1).Equal(t, s2, d.String(1), "%s\n%s", d.String(0), d.String(1))
	}
	eq("abcaadef", "accaa", "\"a\x1b[41mb\x1b[0mcaa\x1b[41mdef\x1b[0m\"", "\"a\x1b[41mc\x1b[0mcaa\"")
	eq("This is\x83这是 Chinese 中文！", "This is    这不是Chinase 汉文吗?", "\"This is\x1b[41m\\x83\x1b[0m这\x1b[41m是 \x1b[0mChin\x1b[41me\x1b[0mse \x1b[41m中\x1b[0m文\x1b[41m！\x1b[0m\"", "\"This is\x1b[41m    \x1b[0m这\x1b[41m不是\x1b[0mChin\x1b[41ma\x1b[0mse \x1b[41m汉\x1b[0m文\x1b[41m吗?\x1b[0m\"")
}
