package assert

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	kSucc = iota
	kAssertFail
)

type printer struct {
	res int
	msg *string
	t   *testing.T
}

func (p printer) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	if p.res == kSucc {
		return
	}
	n, err = fmt.Fprint(w, fmt.Sprint(w, p.Sprint(a...)))
	if p.res == kAssertFail {
		p.failNow()
	}
	return
}

func (p printer) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if p.res == kSucc {
		return
	}
	n, err = fmt.Fprint(w, p.Sprintf(format, a...))
	if p.res == kAssertFail {
		p.failNow()
	}
	return
}

func (p printer) Print(a ...interface{}) (n int, err error) {
	return p.Fprint(os.Stdout, a...)
}

func (p printer) Printf(format string, a ...interface{}) (n int, err error) {
	return p.Fprintf(os.Stdout, format, a...)
}

func (p printer) sprint(a ...interface{}) string {
	return p.format(fmt.Sprint(a...))
}

func (p printer) sprintf(format string, a ...interface{}) string {
	return p.format(fmt.Sprintf(format, a...))
}

func (p printer) format(s string) string {
	const heading = "\t"
	if heading == "" {
		return s
	}
	var buf bytes.Buffer
	for i, l := range strings.Split(s, "\n") {
		if i > 0 {
			buf.WriteString("\n")
			if l != "" {
				buf.WriteString(heading)
			}
		}
		buf.WriteString(l)
	}
	return buf.String()
}

func (p printer) failNow() {

}
