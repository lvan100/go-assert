/*
 * Copyright 2024 The Go-Spring Authors.
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

// Package assert provides assertion helpers for testing,
// offering both functional and fluent assertion styles.
package assert

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lvan100/go-assert/internal"
)

func fail(t internal.T, str string, msg ...string) {
	t.Helper()
	args := append([]string{str}, msg...)
	t.Error(strings.Join(args, ", "))
}

// True asserts that got is true. It reports an error if the value is false.
func True(t internal.T, got bool, msg ...string) {
	t.Helper()
	if !got {
		fail(t, "got false but expect true", msg...)
	}
}

// False asserts that got is false. It reports an error if the value is true.
func False(t internal.T, got bool, msg ...string) {
	t.Helper()
	if got {
		fail(t, "got true but expect false", msg...)
	}
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return v.IsNil()
	default:
		return !v.IsValid()
	}
}

// Nil asserts that got is nil. It reports an error if the value is not nil.
func Nil(t internal.T, got interface{}, msg ...string) {
	t.Helper()
	// Why can't we use got==nil to judge？Because if
	// a := (*int)(nil)        // %T == *int
	// b := (interface{})(nil) // %T == <nil>
	// then a==b is false, because they are different types.
	if !isNil(reflect.ValueOf(got)) {
		str := fmt.Sprintf("got (%T) %v but expect nil", got, got)
		fail(t, str, msg...)
	}
}

// NotNil asserts that got is not nil. It reports an error if the value is nil.
func NotNil(t internal.T, got interface{}, msg ...string) {
	t.Helper()
	if isNil(reflect.ValueOf(got)) {
		fail(t, "got nil but expect not nil", msg...)
	}
}

// Panic asserts that fn panics and the panic message matches expr.
// It reports an error if fn does not panic or if the recovered message does not satisfy expr.
func Panic(t internal.T, fn func(), expr string, msg ...string) {
	t.Helper()
	str := recovery(fn)
	if str == "<<SUCCESS>>" {
		fail(t, "did not panic", msg...)
	} else {
		matches(t, str, expr, msg...)
	}
}

func recovery(fn func()) (str string) {
	defer func() {
		if r := recover(); r != nil {
			str = fmt.Sprint(r)
		}
	}()
	fn()
	return "<<SUCCESS>>"
}

// ThatAssertion wraps a test context and a value for fluent assertions.
type ThatAssertion struct {
	t internal.T
	v interface{}
}

// That creates a ThatAssertion for the given value v and test context t.
func That(t internal.T, v interface{}) *ThatAssertion {
	return &ThatAssertion{
		t: t,
		v: v,
	}
}

// Equal asserts that the wrapped value v is deeply equal to expect.
// It reports an error if the values are not deeply equal.
func (a *ThatAssertion) Equal(expect interface{}, msg ...string) {
	a.t.Helper()
	if !reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotEqual asserts that the wrapped value v is not deeply equal to expect.
// It reports an error if the values are deeply equal.
func (a *ThatAssertion) NotEqual(expect interface{}, msg ...string) {
	a.t.Helper()
	if reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect not (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// Same asserts that the wrapped value v and expect are the same (using Go ==).
// It reports an error if v != expect.
func (a *ThatAssertion) Same(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v != expect {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotSame asserts that the wrapped value v and expect are not the same (using Go !=).
// It reports an error if v == expect.
func (a *ThatAssertion) NotSame(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v == expect {
		str := fmt.Sprintf("expect not (%T) %v", expect, expect)
		fail(a.t, str, msg...)
	}
}

// TypeOf asserts that the type of the wrapped value v is assignable to the type of expect.
// It supports pointer to interface types.
// It reports an error if the types are not assignable.
func (a *ThatAssertion) TypeOf(expect interface{}, msg ...string) {
	a.t.Helper()

	e1 := reflect.TypeOf(a.v)
	e2 := reflect.TypeOf(expect)
	if e2.Kind() == reflect.Ptr && e2.Elem().Kind() == reflect.Interface {
		e2 = e2.Elem()
	}

	if !e1.AssignableTo(e2) {
		str := fmt.Sprintf("got type (%s) but expect type (%s)", e1, e2)
		fail(a.t, str, msg...)
	}
}

// Implements asserts that the type of the wrapped value v implements the interface type of expect.
// The expect parameter must be an interface or pointer to interface.
// It reports an error if v does not implement the interface.
func (a *ThatAssertion) Implements(expect interface{}, msg ...string) {
	a.t.Helper()

	e1 := reflect.TypeOf(a.v)
	e2 := reflect.TypeOf(expect)
	if e2.Kind() == reflect.Ptr {
		if e2.Elem().Kind() == reflect.Interface {
			e2 = e2.Elem()
		} else {
			fail(a.t, "expect should be interface", msg...)
			return
		}
	}

	if !e1.Implements(e2) {
		str := fmt.Sprintf("got type (%s) but expect type (%s)", e1, e2)
		fail(a.t, str, msg...)
	}
}

// Has asserts that the wrapped value v has a method named 'Has' that returns true when passed expect.
// It reports an error if the method does not exist or returns false.
func (a *ThatAssertion) Has(expect interface{}, msg ...string) {
	a.t.Helper()

	m := reflect.ValueOf(a.v).MethodByName("Has")
	if !m.IsValid() {
		str := fmt.Sprintf("method 'Has' not found on type %T", a.v)
		fail(a.t, str, msg...)
		return
	}

	if m.Type().NumOut() != 1 || m.Type().Out(0).Kind() != reflect.Bool {
		fail(a.t, "method 'Has' must return only a bool", msg...)
		return
	}

	ret := m.Call([]reflect.Value{reflect.ValueOf(expect)})
	if !ret[0].Bool() {
		str := fmt.Sprintf("got (%T) %v not has (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// Contains asserts that the wrapped value v has a method named 'Contains' that returns true when passed expect.
// It reports an error if the method does not exist or returns false.
func (a *ThatAssertion) Contains(expect interface{}, msg ...string) {
	a.t.Helper()

	m := reflect.ValueOf(a.v).MethodByName("Contains")
	if !m.IsValid() {
		str := fmt.Sprintf("method 'Contains' not found on type %T", a.v)
		fail(a.t, str, msg...)
		return
	}

	if m.Type().NumOut() != 1 || m.Type().Out(0).Kind() != reflect.Bool {
		fail(a.t, "method 'Contains' must return only a bool", msg...)
		return
	}

	ret := m.Call([]reflect.Value{reflect.ValueOf(expect)})
	if !ret[0].Bool() {
		str := fmt.Sprintf("got (%T) %v not contains (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// InSlice asserts that the wrapped value v is present in the provided slice or array.
// It reports an error if expect is not a slice/array or if v is not found.
func (a *ThatAssertion) InSlice(expect interface{}, msg ...string) {
	a.t.Helper()

	v := reflect.ValueOf(expect)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		str := fmt.Sprintf("unsupported expect value (%T) %v", expect, expect)
		fail(a.t, str, msg...)
		return
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(a.v, v.Index(i).Interface()) {
			return
		}
	}

	str := fmt.Sprintf("got (%T) %v is not in (%T) %v", a.v, a.v, expect, expect)
	fail(a.t, str, msg...)
}

// NotInSlice asserts that the wrapped value v is not present in the provided slice or array.
// It reports an error if expect is not a slice/array, if types do not match, or if v is found.
func (a *ThatAssertion) NotInSlice(expect interface{}, msg ...string) {
	a.t.Helper()

	v := reflect.ValueOf(expect)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		str := fmt.Sprintf("unsupported expect value (%T) %v", expect, expect)
		fail(a.t, str, msg...)
		return
	}

	e := reflect.TypeOf(a.v)
	if e != v.Type().Elem() {
		str := fmt.Sprintf("got type (%s) doesn't match expect type (%s)", e, v.Type())
		fail(a.t, str, msg...)
		return
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(a.v, v.Index(i).Interface()) {
			str := fmt.Sprintf("got (%T) %v is in (%T) %v", a.v, a.v, expect, expect)
			fail(a.t, str, msg...)
			return
		}
	}
}

// InMapKeys asserts that the assertion’s value is one of the keys in the provided map.
// It fails the test if the expected value is not a map or if the actual value
// does not match any key in the map.
func (a *ThatAssertion) InMapKeys(expect interface{}, msg ...string) {
	a.t.Helper()

	switch v := reflect.ValueOf(expect); v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if reflect.DeepEqual(a.v, key.Interface()) {
				return
			}
		}
	default:
		str := fmt.Sprintf("unsupported expect value (%T) %v", expect, expect)
		fail(a.t, str, msg...)
		return
	}

	str := fmt.Sprintf("got (%T) %v is not in keys of (%T) %v", a.v, a.v, expect, expect)
	fail(a.t, str, msg...)
}

// InMapValues asserts that the assertion’s value is one of the values in the provided map.
// It fails the test if the expected value is not a map or if the actual value
// does not match any value in the map.
func (a *ThatAssertion) InMapValues(expect interface{}, msg ...string) {
	a.t.Helper()

	switch v := reflect.ValueOf(expect); v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key).Interface()
			if reflect.DeepEqual(a.v, val) {
				return
			}
		}
	default:
		str := fmt.Sprintf("unsupported expect value (%T) %v", expect, expect)
		fail(a.t, str, msg...)
		return
	}

	str := fmt.Sprintf("got (%T) %v is not in values of (%T) %v", a.v, a.v, expect, expect)
	fail(a.t, str, msg...)
}
