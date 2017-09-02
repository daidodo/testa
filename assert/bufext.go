/*
* Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package assert

import (
	"bytes"
	"fmt"
	"io"
)

type tFeatureBuf struct {
	w    io.Writer
	pl   bytes.Buffer
	code string
	Tab  int
	PH   bool // last plain highlight flag
}

func (b *tFeatureBuf) Writef(hl bool, format string, a ...interface{}) *tFeatureBuf {
	return b.Write(hl, fmt.Sprintf(format, a...))
}

func (b *tFeatureBuf) Write(hl bool, a ...interface{}) *tFeatureBuf {
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

func (b *tFeatureBuf) Normalf(format string, a ...interface{}) *tFeatureBuf {
	return b.Writef(false, format, a...)
}

func (b *tFeatureBuf) Normal(a ...interface{}) *tFeatureBuf {
	return b.Write(false, a...)
}

func (b *tFeatureBuf) Highlightf(format string, a ...interface{}) *tFeatureBuf {
	return b.Writef(true, format, a...)
}

func (b *tFeatureBuf) Highlight(a ...interface{}) *tFeatureBuf {
	return b.Write(true, a...)
}

func (b *tFeatureBuf) Plainf(format string, a ...interface{}) *tFeatureBuf {
	return b.Plain(fmt.Sprintf(format, a...))
}

func (b *tFeatureBuf) Plain(a ...interface{}) *tFeatureBuf {
	b.pl.WriteString(fmt.Sprint(a...))
	return b
}

func (b *tFeatureBuf) NL() *tFeatureBuf {
	b.Normal("\n")
	for i := 0; i < b.Tab; i++ {
		b.Normal("\t")
	}
	return b
}

func (b *tFeatureBuf) Finish() {
	b.normal()
}

func (b *tFeatureBuf) writeString(s string) *tFeatureBuf {
	if len(s) > 0 {
		b.w.Write([]byte(s))
	}
	return b
}

func (b *tFeatureBuf) flushPlain() {
	b.PH = b.code != ""
	b.writeString(b.pl.String())
	b.pl.Reset()
}

const kEND = "\033[0m"
const kRED = "\033[41m"

func (b *tFeatureBuf) normal() {
	defer b.flushPlain()
	if b.code == "" {
		return
	}
	b.writeString(kEND)
	b.code = ""
}

func (b *tFeatureBuf) highlight() {
	b.flushPlain()
	if b.code == kRED {
		return
	} else if b.code != "" {
		b.writeString(kEND)
	}
	b.writeString(kRED)
	b.code = kRED
}
