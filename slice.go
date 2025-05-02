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
	"cmp"
	"fmt"

	"github.com/lvan100/go-assert/internal"
)

type SliceAssertion[T cmp.Ordered] struct {
	t internal.T
	v []T
}

func ThatSlice[T cmp.Ordered](t internal.T, v []T) *SliceAssertion[T] {
	return &SliceAssertion[T]{
		t: t,
		v: v,
	}
}

// Len asserts that the slice has the expected length.
func (a *SliceAssertion[T]) Len(length int, msg ...string) {
	a.t.Helper()
	if len(a.v) != length {
		str := fmt.Sprintf("got length %d but expect length %d", len(a.v), length)
		fail(a.t, str, msg...)
	}
}

// IsEmpty asserts that the slice is empty.
func (a *SliceAssertion[T]) IsEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) != 0 {
		str := fmt.Sprintf("got %v is not empty", a.v)
		fail(a.t, str, msg...)
	}
}

// IsNotEmpty asserts that the slice is not empty.
func (a *SliceAssertion[T]) IsNotEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) == 0 {
		str := fmt.Sprintf("got %v is empty", a.v)
		fail(a.t, str, msg...)
	}
}

// IsNil asserts that the slice is nil.
func (a *SliceAssertion[T]) IsNil(msg ...string) {
	a.t.Helper()
	if a.v != nil {
		str := fmt.Sprintf("got %v is not nil", a.v)
		fail(a.t, str, msg...)
	}
}

// IsNotNil asserts that the slice is not nil.
func (a *SliceAssertion[T]) IsNotNil(msg ...string) {
	a.t.Helper()
	if a.v == nil {
		str := fmt.Sprintf("got %v is nil", a.v)
		fail(a.t, str, msg...)
	}
}

// Zero asserts that the slice is nil or empty.
func (a *SliceAssertion[T]) Zero(msg ...string) {
	a.t.Helper()
	if a.v != nil && len(a.v) != 0 {
		str := fmt.Sprintf("got %v is not nil or empty", a.v)
		fail(a.t, str, msg...)
	}
}

// NotZero asserts that the slice is not nil and not empty.
func (a *SliceAssertion[T]) NotZero(msg ...string) {
	a.t.Helper()
	if a.v == nil || len(a.v) == 0 {
		str := fmt.Sprintf("got %v is nil or empty", a.v)
		fail(a.t, str, msg...)
	}
}

// Contains asserts that the slice contains the expected element.
func (a *SliceAssertion[T]) Contains(element T, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if v == element {
			return
		}
	}
	str := fmt.Sprintf("got %v does not contain %v", a.v, element)
	fail(a.t, str, msg...)
}

// NotContains asserts that the slice does not contain the expected element.
func (a *SliceAssertion[T]) NotContains(element T, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if v == element {
			str := fmt.Sprintf("got %v contains %v", a.v, element)
			fail(a.t, str, msg...)
			return
		}
	}
}

// SubSlice asserts that the slice contains the expected sub-slice.
func (a *SliceAssertion[T]) SubSlice(sub []T, msg ...string) {
	a.t.Helper()
	if len(sub) == 0 {
		return
	}
	for i := 0; i <= len(a.v)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if a.v[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return
		}
	}
	str := fmt.Sprintf("got %v does not contain sub-slice %v", a.v, sub)
	fail(a.t, str, msg...)
}

// NotSubSlice asserts that the slice does not contain the expected sub-slice.
func (a *SliceAssertion[T]) NotSubSlice(sub []T, msg ...string) {
	a.t.Helper()
	if len(sub) == 0 {
		return
	}
	for i := 0; i <= len(a.v)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if a.v[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			str := fmt.Sprintf("got %v contains sub-slice %v", a.v, sub)
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasPrefix asserts that the slice starts with the specified prefix.
func (a *SliceAssertion[T]) HasPrefix(prefix []T, msg ...string) {
	a.t.Helper()
	if len(prefix) > len(a.v) {
		str := fmt.Sprintf("got length %d is less than prefix length %d", len(a.v), len(prefix))
		fail(a.t, str, msg...)
		return
	}
	for i := range prefix {
		if a.v[i] != prefix[i] {
			str := fmt.Sprintf("got element %v at index %d does not match prefix element %v", a.v[i], i, prefix[i])
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasSuffix asserts that the slice ends with the specified suffix.
func (a *SliceAssertion[T]) HasSuffix(suffix []T, msg ...string) {
	a.t.Helper()
	if len(suffix) > len(a.v) {
		str := fmt.Sprintf("got length %d is less than suffix length %d", len(a.v), len(suffix))
		fail(a.t, str, msg...)
		return
	}
	offset := len(a.v) - len(suffix)
	for i := range suffix {
		if a.v[offset+i] != suffix[i] {
			str := fmt.Sprintf("got element %v at index %d does not match suffix element %v", a.v[offset+i], offset+i, suffix[i])
			fail(a.t, str, msg...)
			return
		}
	}
}

// Equal asserts that the slice is equal to the expected slice.
func (a *SliceAssertion[T]) Equal(expect []T, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf("got length %d but expect length %d", len(a.v), len(expect))
		fail(a.t, str, msg...)
		return
	}
	for i := range a.v {
		if a.v[i] != expect[i] {
			str := fmt.Sprintf("got element %v at index %d but expect %v", a.v[i], i, expect[i])
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotEqual asserts that the slice is not equal to the expected slice.
func (a *SliceAssertion[T]) NotEqual(expect []T, msg ...string) {
	a.t.Helper()
	if len(a.v) == len(expect) {
		equal := true
		for i := range a.v {
			if a.v[i] != expect[i] {
				equal = false
				break
			}
		}
		if equal {
			str := fmt.Sprintf("got %v but expect not %v", a.v, expect)
			fail(a.t, str, msg...)
		}
	}
}

// IsIncreasing asserts that the slice is strictly increasing.
func (a *SliceAssertion[T]) IsIncreasing(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] >= a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is not greater than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsNonIncreasing asserts that the slice is not strictly increasing.
func (a *SliceAssertion[T]) IsNonIncreasing(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] < a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is greater than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsDecreasing asserts that the slice is strictly decreasing.
func (a *SliceAssertion[T]) IsDecreasing(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] <= a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is not less than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsNonDecreasing asserts that the slice is not strictly decreasing.
func (a *SliceAssertion[T]) IsNonDecreasing(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] > a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is less than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsSorted asserts that the slice is sorted in ascending order.
func (a *SliceAssertion[T]) IsSorted(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] > a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is greater than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsSortedDescending asserts that the slice is sorted in descending order.
func (a *SliceAssertion[T]) IsSortedDescending(msg ...string) {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] < a.v[i] {
			str := fmt.Sprintf("got element %v at index %d is less than %v at index %d", a.v[i], i, a.v[i-1], i-1)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsUnique asserts that all elements in the slice are unique.
func (a *SliceAssertion[T]) IsUnique(msg ...string) {
	a.t.Helper()
	seen := make(map[T]bool)
	for _, v := range a.v {
		if seen[v] {
			str := fmt.Sprintf("got duplicate element %v", v)
			fail(a.t, str, msg...)
			return
		}
		seen[v] = true
	}
}

// IsUniqueBy asserts that all elements in the slice are unique based on a custom function.
func (a *SliceAssertion[T]) IsUniqueBy(fn func(T) interface{}, msg ...string) {
	a.t.Helper()
	seen := make(map[interface{}]bool)
	for _, v := range a.v {
		key := fn(v)
		if seen[key] {
			str := fmt.Sprintf("got duplicate element %v", v)
			fail(a.t, str, msg...)
			return
		}
		seen[key] = true
	}
}

// All asserts that all elements in the slice satisfy the given condition.
func (a *SliceAssertion[T]) All(fn func(T) bool, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if !fn(v) {
			str := fmt.Sprintf("got element %v does not satisfy the condition", v)
			fail(a.t, str, msg...)
			return
		}
	}
}

// Any asserts that at least one element in the slice satisfies the given condition.
func (a *SliceAssertion[T]) Any(fn func(T) bool, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if fn(v) {
			return
		}
	}
	str := fmt.Sprintf("no element in %v satisfies the condition", a.v)
	fail(a.t, str, msg...)
}

// None asserts that no element in the slice satisfies the given condition.
func (a *SliceAssertion[T]) None(fn func(T) bool, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if fn(v) {
			str := fmt.Sprintf("got element %v satisfies the condition", v)
			fail(a.t, str, msg...)
			return
		}
	}
}
