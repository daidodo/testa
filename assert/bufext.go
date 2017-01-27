package assert

import (
	"bytes"
	"fmt"
)

type FeatureBuf struct {
	buf  bytes.Buffer
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
	b.buf.WriteString(fmt.Sprintf(format, a...))
	return b
}

func (b *FeatureBuf) Plain(a ...interface{}) *FeatureBuf {
	b.buf.WriteString(fmt.Sprint(a...))
	return b
}

func (b *FeatureBuf) NL() *FeatureBuf {
	b.Plain("\n")
	for i := 0; i < b.Tab; i++ {
		b.Plain("\t")
	}
	return b
}

func (b *FeatureBuf) String() string {
	b.normal()
	return b.buf.String()
}

func (b *FeatureBuf) Bytes() []byte {
	b.normal()
	return b.buf.Bytes()
}

func (b *FeatureBuf) writeNormalString(s string) *FeatureBuf {
	if len(s) > 0 {
		b.normal()
		b.buf.WriteString(s)
	}
	return b
}

func (b *FeatureBuf) writeHighlightString(s string) *FeatureBuf {
	if len(s) > 0 {
		b.highlight()
		b.buf.WriteString(s)
	}
	return b
}

const kEND = "\033[0m"
const kRED = "\033[41m"

func (b *FeatureBuf) normal() {
	if b.code == "" {
		return
	}
	b.buf.WriteString(kEND)
	b.code = ""
}

func (b *FeatureBuf) red() {
	if b.code == kRED {
		return
	} else if b.code != "" {
		b.buf.WriteString(kEND)
	}
	b.buf.WriteString(kRED)
	b.code = kRED
}

func (b *FeatureBuf) highlight() {
	b.red()
}

/*
type featureBuf struct {
	bytes.Buffer
	code string
}

func NewFeatureBuffer() *featureBuf {
	return &featureBuf{}
}

func (fb *featureBuf) Write(p []byte) (n int, err error) {
	if len(p) < 1 {
		return 0, nil
	}
	fb.normal()
	return fb.Buffer.Write(p)
}

func (fb *featureBuf) WriteByte(c byte) error {
	fb.normal()
	return fb.Buffer.WriteByte(c)
}

func (fb *featureBuf) WriteRune(r rune) (n int, err error) {
	fb.normal()
	return fb.Buffer.WriteRune(r)
}

func (fb *featureBuf) WriteString(s string) (n int, err error) {
	if len(s) < 1 {
		return 0, nil
	}
	fb.normal()
	return fb.Buffer.WriteString(s)
}

func (fb *featureBuf) WriteTo(w io.Writer) (n int64, err error) {
	fb.normal()
	return fb.Buffer.WriteTo(w)

}

func (fb *featureBuf) HighLightWrite(p []byte) (n int, err error) {
	if len(p) < 1 {
		return 0, nil
	}
	fb.highlight()
	return fb.Buffer.Write(p)
}

func (fb *featureBuf) HighLightWriteByte(c byte) error {
	fb.highlight()
	return fb.Buffer.WriteByte(c)
}

func (fb *featureBuf) HighLightWriteRune(r rune) (n int, err error) {
	fb.highlight()
	return fb.Buffer.WriteRune(r)
}

func (fb *featureBuf) HighLightWriteString(s string) (n int, err error) {
	if len(s) < 1 {
		return 0, nil
	}
	fb.highlight()
	return fb.Buffer.WriteString(s)
}

func (fb *featureBuf) Bytes() []byte {
	fb.normal()
	return fb.Buffer.Bytes()
}

func (fb *featureBuf) String() string {
	fb.normal()
	return fb.Buffer.String()
}

const kEND = "\033[0m"
const kRED = "\033[31m"

func (fb *featureBuf) normal() {
	if fb.code == "" {
		return
	}
	fb.Buffer.WriteString(kEND)
	fb.code = ""
}

func (fb *featureBuf) red() {
	if fb.code == kRED {
		return
	} else if fb.code != kEND {
		fb.Buffer.WriteString(kEND)
	}
	fb.Buffer.WriteString(kRED)
	fb.code = kRED
}

func (fb *featureBuf) highlight() {
	fb.red()
}
*/
