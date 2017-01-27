package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

func True(t *testing.T, actual bool, messages ...interface{}) {
	if actual {
		return
	}
	fail(t, true, actual, messages...)
}

func False(t *testing.T, actual bool, messages ...interface{}) {
	if !actual {
		return
	}
	fail(t, false, actual, messages...)
}

func Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(t, expected, actual, messages...)
}

func fail(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	//var buf FeatureBuf
	//if kHOOK {
	//    buf.WriteRune('\t')
	//} else {
	//    buf.WriteRune('\n')
	//}
	function, file, line := codeInfo()
	//if len(function) > 0 {
	//    m = buf.Writef("%v:%v: in %v:", file, line, function)
	//} else {
	//    m = fmt.Writef("%v:%v:", file, line)
	//}

	var content []string
	var m string
	if len(function) > 0 {
		m = fmt.Sprintf("%v:%v: in %v:", file, line, function)
	} else {
		m = fmt.Sprintf("%v:%v:", file, line)
	}
	if kHOOK {
		content = append(content, "\t"+m)
	} else {
		content = append(content, "\n"+m)
	}
	content = append(content, expectAndActual(expected, actual)...)
	if m := formatMsg(messages...); len(m) > 0 {
		content = append(content, m)
	}
	if kHOOK {
		content[len(content)-1] += "\n"
		output := strings.Join(content, "\n\t\t")
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = append([]byte(nil), output...)
		tt.su.Unlock()
		t.FailNow()
	} else {
		t.Fatal(strings.Join(content, "\n"))
	}
}

func codeInfo() (function, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		if fp := runtime.FuncForPC(pc); fp != nil {
			function = lastPartOf(fp.Name())
		}
		file = lastPartOf(file)
	} else {
		file = "???"
		line = 1
	}
	return
}

func lastPartOf(str string) string {
	if index := strings.LastIndex(str, "/"); index >= 0 {
		return str[index+1:]
	} else if index = strings.LastIndex(str, "\\"); index >= 0 {
		return str[index+1:]
	}
	return str
}

func expectAndActual(expected, actual interface{}) []string {
	t := reflect.TypeOf(expected)
	if t != reflect.TypeOf(actual) {
		return []string{
			fmt.Sprintf("expected:\t%v", toTypeStr(expected)),
			fmt.Sprintf("  actual:\t%v", toTypeStr(actual)),
		}
	}
	switch t.Kind() {
	case reflect.Array:
		e, a := diffStr(expected, actual)
		return []string{
			fmt.Sprintf("expected:\t%v", e),
			fmt.Sprintf("  actual:\t%v", a),
		}
	}
	return []string{
		fmt.Sprintf("expected:\t%v", toStr(expected)),
		fmt.Sprintf("  actual:\t%v", toStr(actual)),
	}
}

func toTypeStr(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr:
		return fmt.Sprintf("%T(%#v)", val, val)
	}
	return fmt.Sprintf("%T(%v)", val, val)
}

func toStr(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr:
		return fmt.Sprintf("%#v", val)
	}
	return fmt.Sprintf("%v", val)
}

func diffStr(var1, var2 interface{}) (string, string) {
	v1, v2 := reflect.ValueOf(var1), reflect.ValueOf(var2)
	if v1.Type() != v2.Type() {
		panic(fmt.Sprintf("Should be the same type, but var1 is %T, var2 is %T", var1, var2))
	}
	var b1, b2 []byte
	switch v1.Kind() {
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ { // v1.Len() == v2.Len()
			e1, e2 := v1.Index(i), v2.Index(i)
			if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
				b1 = appendValueStr(b1, toStr(e1), false)
				b2 = appendValueStr(b2, toStr(e2), false)
			} else {
				b1 = appendValueStr(b1, toStr(e1), true)
				b2 = appendValueStr(b2, toStr(e2), true)
			}
		}
		return string(b1), string(b2)
	}
	return toStr(var1), toStr(var2)
}

func appendValueStr(buf []byte, str string, highlight bool) []byte {
	if len(str) < 1 {
		return buf
	}
	sp := len(buf) > 0
	if highlight {
		buf = append(buf, 033, '[', '3', '1', 'm')
	}
	if sp {
		buf = append(buf, ' ')
	}
	buf = append(buf, str...)
	if highlight {
		buf = append(buf, 033, '[', '0', 'm')
	}
	return buf
}

func formatMsg(m ...interface{}) string {
	if len(m) < 1 {
		return ""
	}
	if s, ok := m[0].(string); ok {
		return fmt.Sprintf(s, m[1:]...)
	}
	return fmt.Sprint(m...)
}
