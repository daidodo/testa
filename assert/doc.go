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
	Package assert provides useful tools for unit testing in Go. Many features are unique comparing
	to other testing tools you can find.


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
		16     myTest(t, 100, 101) // This is the part we want to notice!
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
		16     myTest(t, 100, 101)	// This is the part we want to notice!
		17 }
	will now produce:
		example_test.go:16: in example.TestA:
		example_test.go:10: in example.myTest:
			Expect: 100
			Actual: 101
*/
package assert
