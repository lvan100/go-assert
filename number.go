/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package assert

import (
	"fmt"

	"github.com/lvan100/go-assert/internal"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// NumberAssertion encapsulates a number value and a test handler for making assertions on the number.
type NumberAssertion[T Number] struct {
	t internal.T
	v T
}

// ThatNumber returns a NumberAssertion for the given testing object and number value.
func ThatNumber[T Number](t internal.T, v T) *NumberAssertion[T] {
	return &NumberAssertion[T]{
		t: t,
		v: v,
	}
}

// Equal asserts that the number value is equal to the expected value.
func (a *NumberAssertion[T]) Equal(expect T, msg ...string) {
	a.t.Helper()
	if a.v != expect {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotEqual asserts that the number value is not equal to the expected value.
func (a *NumberAssertion[T]) NotEqual(expect T, msg ...string) {
	a.t.Helper()
	if a.v == expect {
		str := fmt.Sprintf("got (%T) %v but expect not (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// GreaterThan asserts that the number value is greater than the expected value.
func (a *NumberAssertion[T]) GreaterThan(expect T, msg ...string) {
	a.t.Helper()
	if a.v <= expect {
		str := fmt.Sprintf("got (%T) %v but expect greater than (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// GreaterOrEqual asserts that the number value is greater than or equal to the expected value.
func (a *NumberAssertion[T]) GreaterOrEqual(expect T, msg ...string) {
	a.t.Helper()
	if a.v < expect {
		str := fmt.Sprintf("got (%T) %v but expect greater than or equal to (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// LessThan asserts that the number value is less than the expected value.
func (a *NumberAssertion[T]) LessThan(expect T, msg ...string) {
	a.t.Helper()
	if a.v >= expect {
		str := fmt.Sprintf("got (%T) %v but expect less than (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// LessOrEqual asserts that the number value is less than or equal to the expected value.
func (a *NumberAssertion[T]) LessOrEqual(expect T, msg ...string) {
	a.t.Helper()
	if a.v > expect {
		str := fmt.Sprintf("got (%T) %v but expect less than or equal to (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// IsZero asserts that the number value is zero.
func (a *NumberAssertion[T]) IsZero(msg ...string) {
	a.t.Helper()
	if a.v != 0 {
		str := fmt.Sprintf("got (%T) %v but expect zero", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// NotZero asserts that the number value is not zero.
func (a *NumberAssertion[T]) NotZero(msg ...string) {
	a.t.Helper()
	if a.v == 0 {
		str := fmt.Sprintf("got (%T) %v but expect not zero", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsPositive asserts that the number value is positive.
func (a *NumberAssertion[T]) IsPositive(msg ...string) {
	a.t.Helper()
	if a.v <= 0 {
		str := fmt.Sprintf("got (%T) %v but expect positive", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsNegative asserts that the number value is negative.
func (a *NumberAssertion[T]) IsNegative(msg ...string) {
	a.t.Helper()
	if a.v >= 0 {
		str := fmt.Sprintf("got (%T) %v but expect negative", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsNonNegative asserts that the number value is non-negative.
func (a *NumberAssertion[T]) IsNonNegative(msg ...string) {
	a.t.Helper()
	if a.v < 0 {
		str := fmt.Sprintf("got (%T) %v but expect non-negative", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsNonPositive asserts that the number value is non-positive.
func (a *NumberAssertion[T]) IsNonPositive(msg ...string) {
	a.t.Helper()
	if a.v > 0 {
		str := fmt.Sprintf("got (%T) %v but expect non-positive", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// Between asserts that the number value is between the lower and upper bounds (inclusive).
func (a *NumberAssertion[T]) Between(lower, upper T, msg ...string) {
	a.t.Helper()
	if a.v < lower || a.v > upper {
		str := fmt.Sprintf("got (%T) %v but expect between (%T) %v and (%T) %v", a.v, a.v, lower, lower, upper, upper)
		fail(a.t, str, msg...)
	}
}

// NotBetween asserts that the number value is not between the lower and upper bounds (exclusive).
func (a *NumberAssertion[T]) NotBetween(lower, upper T, msg ...string) {
	a.t.Helper()
	if a.v >= lower && a.v <= upper {
		str := fmt.Sprintf("got (%T) %v but expect not between (%T) %v and (%T) %v", a.v, a.v, lower, lower, upper, upper)
		fail(a.t, str, msg...)
	}
}

// InDelta asserts that the number value is within the delta range of the expected value.
func (a *NumberAssertion[T]) InDelta(expect T, delta T, msg ...string) {
	a.t.Helper()
	diff := a.v - expect
	if diff < 0 {
		diff = -diff
	}
	if diff > delta {
		str := fmt.Sprintf("got (%T) %v is not within delta (%T) %v of (%T) %v", a.v, a.v, delta, delta, expect, expect)
		fail(a.t, str, msg...)
	}
}

// IsNaN asserts that the number value is NaN (Not a Number).
func (a *NumberAssertion[T]) IsNaN(msg ...string) {
	a.t.Helper()
	if !isNaN(a.v) {
		str := fmt.Sprintf("got (%T) %v but expect NaN", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsInf asserts that the number value is infinite.
func (a *NumberAssertion[T]) IsInf(sign int, msg ...string) {
	a.t.Helper()
	if !isInf(a.v, sign) {
		str := fmt.Sprintf("got (%T) %v but expect infinite with sign %d", a.v, a.v, sign)
		fail(a.t, str, msg...)
	}
}

// IsFinite asserts that the number value is finite.
func (a *NumberAssertion[T]) IsFinite(msg ...string) {
	a.t.Helper()
	if isNaN(a.v) || isInf(a.v, 0) {
		str := fmt.Sprintf("got (%T) %v but expect finite", a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// isNaN checks if the value is NaN.
func isNaN[T Number](v T) bool {
	switch any(v).(type) {
	case float32:
		return any(v).(float32) != any(v).(float32)
	case float64:
		return any(v).(float64) != any(v).(float64)
	default:
		return false
	}
}

// isInf checks if the value is infinite.
func isInf[T Number](v T, sign int) bool {
	switch any(v).(type) {
	case float32:
		f := any(v).(float32)
		return (sign >= 0 && f > maxFloat32) || (sign <= 0 && f < -maxFloat32)
	case float64:
		f := any(v).(float64)
		return (sign >= 0 && f > maxFloat64) || (sign <= 0 && f < -maxFloat64)
	default:
		return false
	}
}

const (
	maxFloat32 = 3.4028234663852886e+38
	maxFloat64 = 1.7976931348623157e+308
)
