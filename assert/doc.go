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

/*
	Package assert provides useful tools for unit testing in Go. Many features are unique compared
	to other testing tools you can find.


	Highlighted information

	Tools tend to provide "full" information when an assertion fails, simply because they don't
	understand it.

	But testa/assert tries to understand the information, provide with what is only necessary, and
	highlight the key part of it.


	Caller

	Normally, when an assertion fails, the calling information of that assert statement is shown:
		example_test.go
		 1 package main
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
	will produce:
		example_test.go:10: in example.myTest:
			Expect: 100
			Actual: 101
	But, as you may notice, line 16 of TestA is the real interesting part.

	To add that part to our diagnosis information, Caller can help you:
		example_test.go
		 1 package main
		 2
		 3 import (
		 4     "testing"
		 5
		 6     "github.com/daidodo/testa/assert"
		 7 )
		 8
		 9 func myTest(t *testing.T, e, a int) {
		10     assert.Caller(1).Equal(t, e, a) // Show calling information 1 level up the chain
		11 }
		12
		13 func TestA(t *testing.T) {
		14     myTest(t, 1, 1)
		15     myTest(t, 10, 10)
		16     myTest(t, 100, 101)	// This is the part we want to see!
		17 }
	will now produce:
		example_test.go:16: in example.TestA:
		example_test.go:10: in example.myTest:
			Expect: 100
			Actual: 101


	EqualValue

	Different from assert.Equal, assert.EqualValue compares objects by values only, regardless of
	their types. So int(100) is equal to uint(100) in value, but not in type.

	testa/assert implements EqualValue from the scratch, using intuitive and common sense, with
	regards to reflect.DeepEqual. The general rules are:

	Boolean is comparable only to Boolean;

	Math objects (signed/unsigned integers, floats, complexes) are compared mathematically, e.g.
	uint8(255) != int8(-1), int(1) == complex64(1+0i);

	Different types of pointers are not equal to each other; But pointers are comparable to
	unsafe.Pointer;

	Array and slice objects are equal in value if: a) they are both nil; or b) they both have zero
	length and their elements' types are convertible; or c) they have the same length and all
	corresponding elements are equal in value;

	Maps are equal in value if: a) they are both nil; or b) they both have zero length and their
	keys and elements' types are both convertible, respectively; or c) they have the same length
	and all keys are DEEPLY equal and the corresponding elements are equal in value;

	Structs are equal in value if they have the same type and all corresponding fields are equal in
	value;

	As an exception, array or slice of byte or rune can compare to string;

	All other objects are equal in value only when they are deeply equal defined by
	reflect.DeepEqual.
*/
package assert

//TODO:
//Contain/NotContain
//ContainValue/NotContainValue
//GreaterThan/LessThan
//GreaterEqual/LessEqual
//Empty/NotEmpty
//Error/NoError/EqualError
