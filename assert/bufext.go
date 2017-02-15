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
	"io"
)

type tFeatureBuf struct {
	w    io.Writer
	pl   bytes.Buffer
	code string
	Tab  int
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
