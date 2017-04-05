/*
 * Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or any
 * later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see accompanying file LICENSE.txt
 * or <http://www.gnu.org/licenses/>.
 */

package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

// NotEqualType asserts whether e and a have different types, regardless of their values.
func NotEqualType(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.NotEqualType(t, e, a, m...)
}

// NotEqualType asserts whether e and a have different types, regardless of their values.
func (c CallerT) NotEqualType(t *testing.T, e, a interface{}, m ...interface{}) {
	if reflect.TypeOf(e) != reflect.TypeOf(a) {
		return
	}
	fail(c, t, e, a, kNeT, m...)
}

// EqualType asserts whether e and a have the same type, regardless of their values.
func EqualType(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.EqualType(t, e, a, m...)
}

// EqualType asserts whether e and a have the same type, regardless of their values.
func (c CallerT) EqualType(t *testing.T, e, a interface{}, m ...interface{}) {
	if reflect.TypeOf(e) == reflect.TypeOf(a) {
		return
	}
	fail(c, t, e, a, kEqT, m...)
}

// NotEqualValue asserts whether e and a have unequal values, regardless of their types.
func NotEqualValue(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.NotEqualValue(t, e, a, m...)
}

// NotEqualValue asserts whether e and a have unequal values, regardless of their types.
func (c CallerT) NotEqualValue(t *testing.T, e, a interface{}, m ...interface{}) {
	if !isSameInValue(e, a) {
		return
	}
	fail(c, t, e, a, kNeV, m...)
}

// EqualValue asserts whether e and a have identical value, regardless of their types.
// For example, int(100) and uint(100) have the same value, but different types, so:
//		assert.NotEqual(t, int(100), uint(100))
//		assert.EqualValue(t, int(100), uint(100))
func EqualValue(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.EqualValue(t, e, a, m...)
}

// EqualValue asserts whether e and a have identical value, regardless of their types.
// For example, int(100) and uint(100) have the same value, but different types, so:
//		assert.NotEqual(t, int(100), uint(100))
//		assert.EqualValue(t, int(100), uint(100))
func (c CallerT) EqualValue(t *testing.T, e, a interface{}, m ...interface{}) {
	if isSameInValue(e, a) {
		return
	}
	fail(c, t, e, a, kEqV, m...)
}

// NotNil asserts whether a is not nil.
//
// It is DIFFERENT from assert.NotEqual against nil.
// For example:
//		var a chan int
//		assert.NotEqual(t, nil, a) // Success! But a IS nil
//		assert.NotNil(t, a)        // Fail
func NotNil(t *testing.T, a interface{}, m ...interface{}) {
	CallerT{1, 1}.NotNil(t, a, m...)
}

// NotNil asserts whether a is not nil.
//
// It is DIFFERENT from assert.NotEqual against nil.
// For example:
//		var a chan int
//		assert.Caller(1).NotEqual(t, nil, a) // Success! But a IS nil
//		assert.Caller(1).NotNil(t, a)        // Fail
func (c CallerT) NotNil(t *testing.T, a interface{}, m ...interface{}) {
	if !isNil(a) {
		return
	}
	fail(c, t, nil, a, kNNil, m...)
}

// Nil asserts whether a is nil.
//
// It is DIFFERENT from assert.Equal against nil.
// For example:
//		var a chan int
//		assert.Equal(t, nil, a) // Fail!
//		assert.Nil(t, a)        // Success
func Nil(t *testing.T, a interface{}, m ...interface{}) {
	CallerT{1, 1}.Nil(t, a, m...)
}

// Nil asserts whether a is nil.
//
// It is DIFFERENT from assert.Equal against nil.
// For example:
//		var a chan int
//		assert.Caller(1).Equal(t, nil, a) // Fail!
//		assert.Caller(1).Nil(t, a)        // Success
func (c CallerT) Nil(t *testing.T, a interface{}, m ...interface{}) {
	if isNil(a) {
		return
	}
	fail(c, t, nil, a, kNil, m...)
}

// True asserts that a is true.
func True(t *testing.T, a bool, m ...interface{}) {
	CallerT{1, 1}.True(t, a, m...)
}

// True asserts that a is true.
func (c CallerT) True(t *testing.T, a bool, m ...interface{}) {
	if a {
		return
	}
	fail(c, t, true, a, kEq, m...)
}

// False asserts that a is false.
func False(t *testing.T, a bool, m ...interface{}) {
	CallerT{1, 1}.False(t, a, m...)
}

// False asserts that a is false.
func (c CallerT) False(t *testing.T, a bool, m ...interface{}) {
	if !a {
		return
	}
	fail(c, t, false, a, kEq, m...)
}

// Equal asserts that e and a are exactly the same, both type and value.
func Equal(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.Equal(t, e, a, m...)
}

// Equal asserts that e and a are exactly the same, both type and value.
func (c CallerT) Equal(t *testing.T, e, a interface{}, m ...interface{}) {
	if reflect.DeepEqual(e, a) {
		return
	}
	fail(c, t, e, a, kEq, m...)
}

// NotEqual asserts that e and a are not the same, either type or value.
func NotEqual(t *testing.T, e, a interface{}, m ...interface{}) {
	CallerT{1, 1}.NotEqual(t, e, a, m...)
}

// NotEqual asserts that e and a are not the same, either type or value.
func (c CallerT) NotEqual(t *testing.T, e, a interface{}, m ...interface{}) {
	if !reflect.DeepEqual(e, a) {
		return
	}
	fail(c, t, e, a, kNe, m...)
}

// CallerT is useful for customizing calling information shown for assertions.
type CallerT struct {
	from, to int
}

// Caller changes calling information shown for assertions.
func Caller(lv int) CallerT {
	return CallerT{0, lv}
}

type tRes int

const (
	kEq tRes = iota
	kNe
	kNil
	kNNil
	kEqV
	kNeV
	kEqT
	kNeT
)

func fail(c CallerT, t *testing.T, expected, actual interface{}, res tRes, msg ...interface{}) {
	var buf bytes.Buffer
	b := tFeatureBuf{w: &buf, Tab: 0}
	writeCodeInfo(c, &b)
	b.Tab++
	switch res {
	case kEq:
		writeFailEq(&b, expected, actual)
	case kNe:
		writeFailNe(&b, actual)
	case kNil:
		writeFailNil(&b, actual)
	case kNNil:
		writeFailNNil(&b, actual)
	case kEqV:
		writeFailEqV(&b, expected, actual)
	case kNeV:
		writeFailNeV(&b, expected, actual)
	case kEqT:
		writeFailEqT(&b, expected, actual)
	case kNeT:
		writeFailNeT(&b, actual)
	}
	writeMessages(&b, msg...)
	b.Tab--
	b.Finish()
	flushLog(t, &buf)
	t.FailNow()
}

func writeFailEq(buf *tFeatureBuf, expected, actual interface{}) {
	var v tValueDiffer
	v.WriteDiff(reflect.ValueOf(expected), reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normalf("Expect:\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailNe(buf *tFeatureBuf, actual interface{}) {
	var v tValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("DIFF to Actual")
	buf.NL().Normalf("\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeFailNil(buf *tFeatureBuf, actual interface{}) {
	var v tValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight(nil)
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeFailNNil(buf *tFeatureBuf, actual interface{}) {
	var v tValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("NOT ", nil)
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeFailEqV(buf *tFeatureBuf, expected, actual interface{}) {
	var v tValueDiffer
	v.WriteValDiff(reflect.ValueOf(expected), reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normalf("Expect:\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailNeV(buf *tFeatureBuf, expected, actual interface{}) {
	var v tValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(expected), buf.Tab+1)
	v.WriteTypeValue(1, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("DIFF to Actual")
	buf.NL().Normalf("\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailEqT(buf *tFeatureBuf, expected, actual interface{}) {
	var v tValueDiffer
	v.WriteDiffT(reflect.TypeOf(expected), reflect.TypeOf(actual), buf.Tab+1)
	buf.NL().Normalf("Expect:\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailNeT(buf *tFeatureBuf, actual interface{}) {
	var v tValueDiffer
	v.WriteType(0, reflect.TypeOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("DIFF to Actual")
	buf.NL().Normalf("\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeAttrs(buf *tFeatureBuf, v tValueDiffer) {
	if v.Attrs[kOmitSame] {
		buf.NL().Normal("\t(").Highlight("Only diffs are shown").Normal(")")
	}
	if v.Attrs[kCompFunc] {
		buf.NL().Normal("\t(").Highlight("func can only be compared to nil").Normal(")")
	}
}

func writeCodeInfo(c CallerT, buf *tFeatureBuf) {
	narrow(&c.from, 0, 100)
	narrow(&c.to, c.from, 100)
	for find := false; c.to >= c.from; c.to-- {
		if find {
			buf.NL()
		}
		pc, file, line, ok := runtime.Caller(3 + c.to)
		if ok {
			buf.Normalf("%v:%v:", lastPartOf(file), line)
			if fp := runtime.FuncForPC(pc); fp != nil {
				buf.Normalf(" in %v:", lastPartOf(fp.Name()))
			}
			find = true
		} else if find || c.to == c.from {
			buf.Normal("???:1:")
		}
	}
}

func writeMessages(buf *tFeatureBuf, messages ...interface{}) {
	if len(messages) < 1 {
		return
	}
	var m, h string
	if s, ok := messages[0].(string); ok {
		m = fmt.Sprintf(s, messages[1:]...)
	} else {
		m = fmt.Sprint(messages...)
	}
	for i := 0; i < buf.Tab; i++ {
		h = h + "\t"
	}
	m = format(h, m)
	buf.Normal(m)
}

func flushLog(t *testing.T, buf *bytes.Buffer) {
	if kHOOK {
		buf.WriteByte('\n')
		output := buf.Bytes()
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = output
		tt.su.Unlock()
	} else {
		t.Log("\n" + buf.String())
	}
}

func narrow(i *int, min, max int) {
	if *i < min {
		*i = min
	}
	if *i > max {
		*i = max
	}
}

func lastPartOf(str string) string {
	if index := strings.LastIndex(str, "/"); index >= 0 {
		return str[index+1:]
	} else if index = strings.LastIndex(str, "\\"); index >= 0 {
		return str[index+1:]
	}
	return str
}

func format(h, s string) string {
	if h == "" {
		return s
	}
	var buf bytes.Buffer
	for _, l := range strings.Split(s, "\n") {
		buf.WriteString("\n")
		if l != "" {
			buf.WriteString(h)
		}
		buf.WriteString(l)
	}
	return buf.String()
}
