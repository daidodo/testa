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

func Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(t, expected, actual, messages...)
}

func fail(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	file, line := fileLine()
	content := []string{
		fmt.Sprintf("\n\t%v:%v:", file, line),
	}
	content = append(content, expectAndActual(expected, actual)...)
	if m := formatMsg(messages...); len(m) > 0 {
		content = append(content, m)
	}
	if kHOOK {
		content[len(content)-1] += "\n"
		output := strings.Join(content, "\n\t")
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = append([]byte(nil), output...)
		tt.su.Unlock()
		t.FailNow()
	} else {
		t.Fatal(strings.Join(content, "\n"))
	}
}

func fileLine() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return file, line
}

func expectAndActual(expected, actual interface{}) []string {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return []string{
			fmt.Sprintf("expected:\t%T(%#v)", expected, expected),
			fmt.Sprintf("  actual:\t%T(%#v)", actual, actual),
		}
	}
	return []string{
		fmt.Sprintf("expected:\t%#v", expected),
		fmt.Sprintf("  actual:\t%#v", actual),
	}
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
