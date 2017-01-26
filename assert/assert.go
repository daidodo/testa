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
	function, file, line := codeInfo()
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
			fmt.Sprintf("expected:\t%v", toStr(expected, true)),
			fmt.Sprintf("  actual:\t%v", toStr(actual, true)),
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
		fmt.Sprintf("expected:\t%v", toStr(expected, false)),
		fmt.Sprintf("  actual:\t%v", toStr(actual, false)),
	}
}

func toStr(val interface{}, withType bool) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr:
		if withType {
			return fmt.Sprintf("%T(%#v)", val, val)
		}
		return fmt.Sprintf("%#v", val)
	}
	if withType {
		return fmt.Sprintf("%T(%v)", val, val)
	}
	return fmt.Sprintf("%v", val)
}

func diffStr(v1, v2 interface{}) (s1, s2 string) {

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
