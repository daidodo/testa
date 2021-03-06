# Testa
[![GoDoc](https://godoc.org/github.com/daidodo/testa/assert?status.svg)](https://godoc.org/github.com/daidodo/testa/assert)
[![Report Card](https://goreportcard.com/badge/github.com/daidodo/testa)](https://goreportcard.com/badge/github.com/daidodo/testa)
[![Build Status](https://travis-ci.org/daidodo/testa.svg?branch=master)](https://travis-ci.org/daidodo/testa)

Testa is a Go library that aimed at easing testing development.

Package *testa/assert* is a powerful unit testing tool. Many features are unique compared to other testing tools you can find.

## Usage
* Install:
```
$ go get github.com/daidodo/testa
```
* Import:
```{.go}
import "github.com/daidodo/testa/assert"
```
* Code:
```{.go}
func TestA(t *testing.T) {
	assert.Equal(t, 100, 100, "100 != 100?")
}
```
* Integrate with Vim

  * Install plugin ["daidodo/Improved-AnsiEsc"](https://github.com/daidodo/Improved-AnsiEsc)
  * Add this line to `.vimrc`:
```
 au BufReadPost * if getbufvar(winbufnr(0), "&buftype") == "quickfix" | set nospell | call AnsiEsc#AnsiEsc(0) | endif
```
If you're using *vim-go* and `:GoTest` command, you may find some messy codes in the diagnosis window. Basically these instructions are meant to enable *AnsiEsc* so Vim can show highlighted text properly.

## Package Assert
*testa/assert* helps you write unit tests easily and efficiently. Full documentations are available [here](https://godoc.org/github.com/daidodo/testa/assert).

Aside from compatibility with Package testing, one-line assertion and message, various and user-friendly APIs, there are other exceptional features you may find useful:

### Highlighted formatted information
Tools tend to provide "full" or even messy information when an assertion fails, simply because they don't understand it.

But *testa/assert* tries to understand the information, provide with what is only necessary, make it well readable, and  **highlight** the key part of it.
```{.go}
func TestA(t *testing.T) {
	var a, b [5][5]int
	// ...
	assert.Equal(t, a, b, "a == b?")
}
```
![image](https://github.com/daidodo/testa/blob/master/res/1.jpg)

More on this:
```{.go}
func TestA(t *testing.T) {
	m1 := make(map[int8]string)
	m2 := make(map[uint]string)
	// ...
	assert.Equal(t, m1, m2)
}
```
![image](https://github.com/daidodo/testa/blob/master/res/2.jpg)

### Caller
Normally, when an assertion fails, the calling information of the assert statement is shown:
```{.go}
 1 package example
 2
 3 import (
 4     "testing"
 5
 6     "github.com/daidodo/testa/assert"
 7 )
 8
 9 func myTest(t *testing.T, e, a int) {
10     assert.Equal(t, e, a)
11 }
12
13 func TestA(t *testing.T) {
14     myTest(t, 1, 1)
15     myTest(t, 10, 10)
16     myTest(t, 100, 101) // This is the part we want to see!
17 }
```
will produce:
```
example_test.go:10: in example.myTest:
	Expect: 100
	Actual: 101
```
But, as you may notice, line 16 inside *TestA* is the real interesting part.

To add it to our diagnosis information, **assert.Caller** can help you:
```{.go}
 1 package example
 2
 3 import (
 4     "testing"
 5
 6     "github.com/daidodo/testa/assert"
 7 )
 8
 9 func myTest(t *testing.T, e, a int) {
10     assert.Caller(1).Equal(t, e, a) // Show callers information 1 level up the chain
11 }
12
13 func TestA(t *testing.T) {
14     myTest(t, 1, 1)
15     myTest(t, 10, 10)
16     myTest(t, 100, 101)	// This is the part we want to see!
17 }
```
will now produce:
```
example_test.go:16: in example.TestA:
example_test.go:10: in example.myTest:
	Expect: 100
	Actual: 101
```

### EqualValue
Apart from *assert.Equal*, **assert.EqualValue** compares objects by values only, regardless of their types. So `int(100)` is equal to `uint(100)` in value, but not in type. This is *never* a trivial task as someone might think using `reflect.Value.Convert`. 

As a counter-example, [stretchr/testify/assert](https://github.com/stretchr/testify) has a simplified version of *EqualValues* that may cause confusion:
```{.go}
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestB(t *testing.T) {
	a := int32(-1000000000)
	b := int8(0)
	assert.EqualValues(t, a, b) //This is a BUG!!!
}
```
```
PASS
```
The reason is obvious: after converting `int32(-1000000000)` to `int8`, it is `0`.

And there are tons of other subtle corners along the way, and many results may seem surprising to some people.

*testa/assert* implements *EqualValue* from the scratch, using common knowledge and intuition, with regards to *reflect.DeepEqual*.

The general rules are:
* Boolean is comparable only to Boolean;
* Math objects (signed/unsigned integers, floats, complexes) are compared mathematically, e.g. `uint8(255) != int8(-1)`, `int(1) == complex64(1+0i)`;
* Different types of pointers are not equal to each other; But pointers are comparable to `unsafe.Pointer`;
* Array and slice objects are equal in value if: a) they are both `nil`; or b) they both have zero length and their elements' types are convertible; or c) they have the same length and all corresponding elements are equal in value;
* Maps are equal in value if: a) they are both `nil`; or b) they both have zero length and their keys and elements' types are both convertible, respectively; or c) they have the same length and all keys are **deeply equal** and the corresponding elements are equal in value;
* Structs are equal in value if they have the same type and all corresponding fields are equal in value;
* As an exception, array or slice of `byte` or `rune` can compare to `string`;
* All other objects are equal in value only when they are deeply equal defined by *reflect.DeepEqual*.


## License

Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
