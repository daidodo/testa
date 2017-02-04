package assert

import (
	"fmt"
	"io"
)

type FeatureBuf struct {
	w    io.Writer
	code string
	Tab  int
}

func (b *FeatureBuf) Writef(format string, a ...interface{}) *FeatureBuf {
	return b.writeNormalString(fmt.Sprintf(format, a...))
}

func (b *FeatureBuf) Write(a ...interface{}) *FeatureBuf {
	return b.writeNormalString(fmt.Sprint(a...))
}

func (b *FeatureBuf) Highlightf(format string, a ...interface{}) *FeatureBuf {
	return b.writeHighlightString(fmt.Sprintf(format, a...))
}

func (b *FeatureBuf) Highlight(a ...interface{}) *FeatureBuf {
	return b.writeHighlightString(fmt.Sprint(a...))
}

func (b *FeatureBuf) Plainf(format string, a ...interface{}) *FeatureBuf {
	b.w.Write([]byte(fmt.Sprintf(format, a...)))
	return b
}

func (b *FeatureBuf) Plain(a ...interface{}) *FeatureBuf {
	b.w.Write([]byte(fmt.Sprint(a...)))
	return b
}

func (b *FeatureBuf) NL() *FeatureBuf {
	b.Plain("\n")
	for i := 0; i < b.Tab; i++ {
		b.Plain("\t")
	}
	return b
}

func (b *FeatureBuf) Finish() {
	b.normal()
}
func (b *FeatureBuf) writeNormalString(s string) *FeatureBuf {
	if len(s) > 0 {
		b.normal()
		b.w.Write([]byte(s))
	}
	return b
}

func (b *FeatureBuf) writeHighlightString(s string) *FeatureBuf {
	if len(s) > 0 {
		b.highlight()
		b.w.Write([]byte(s))
	}
	return b
}

const kEND = "\033[0m"
const kRED = "\033[41m"

func (b *FeatureBuf) normal() {
	if b.code == "" {
		return
	}
	b.w.Write([]byte(kEND))
	b.code = ""
}

func (b *FeatureBuf) red() {
	if b.code == kRED {
		return
	} else if b.code != "" {
		b.w.Write([]byte(kEND))
	}
	b.w.Write([]byte(kRED))
	b.code = kRED
}

func (b *FeatureBuf) highlight() {
	b.red()
}
