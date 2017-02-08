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

func (b *FeatureBuf) Writef(hl bool, format string, a ...interface{}) *FeatureBuf {
	return b.Write(hl, fmt.Sprintf(format, a...))
}

func (b *FeatureBuf) Write(hl bool, a ...interface{}) *FeatureBuf {
	if s := fmt.Sprint(a...); len(s) > 0 {
		if hl {
			b.highlight()
		} else {
			b.normal()
		}
		b.writeString(s)
	}
	return b
}

func (b *FeatureBuf) Normalf(format string, a ...interface{}) *FeatureBuf {
	return b.Writef(false, format, a...)
}

func (b *FeatureBuf) Normal(a ...interface{}) *FeatureBuf {
	return b.Write(false, a...)
}

func (b *FeatureBuf) Highlightf(format string, a ...interface{}) *FeatureBuf {
	return b.Writef(true, format, a...)
}

func (b *FeatureBuf) Highlight(a ...interface{}) *FeatureBuf {
	return b.Write(true, a...)
}

func (b *FeatureBuf) Plainf(format string, a ...interface{}) *FeatureBuf {
	return b.writeString(fmt.Sprintf(format, a...))
}

func (b *FeatureBuf) Plain(a ...interface{}) *FeatureBuf {
	return b.writeString(fmt.Sprint(a...))
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

func (b *FeatureBuf) writeString(s string) *FeatureBuf {
	if len(s) > 0 {
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
	b.writeString(kEND)
	b.code = ""
}

func (b *FeatureBuf) red() {
	if b.code == kRED {
		return
	} else if b.code != "" {
		b.writeString(kEND)
	}
	b.writeString(kRED)
	b.code = kRED
}

func (b *FeatureBuf) highlight() {
	b.red()
}
