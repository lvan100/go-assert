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

// Package assert provides some useful assertion methods.
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
	t.Error(strings.Join(args, "; "))
}

// True assertion failed when got is false.
func True(t internal.T, got bool, msg ...string) {
	t.Helper()
	if !got {
		fail(t, "got false but expect true", msg...)
	}
}

// False assertion failed when got is true.
func False(t internal.T, got bool, msg ...string) {
	t.Helper()
	if got {
		fail(t, "got true but expect false", msg...)
	}
}

// isNil reports v is nil, but will not panic.
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

// Nil assertion failed when got is not nil.
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

// NotNil assertion failed when got is nil.
func NotNil(t internal.T, got interface{}, msg ...string) {
	t.Helper()
	if isNil(reflect.ValueOf(got)) {
		fail(t, "got nil but expect not nil", msg...)
	}
}

// Panic assertion failed when fn doesn't panic or not match expr expression.
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

// Error assertion failed when got `error` doesn't match expr expression.
func Error(t internal.T, got error, expr string, msg ...string) {
	t.Helper()
	if got == nil {
		fail(t, "expect not nil error", msg...)
		return
	}
	matches(t, got.Error(), expr, msg...)
}

// ThatAssertion assertion for type any.
type ThatAssertion struct {
	t internal.T
	v interface{}
}

// That returns an assertion for type any.
func That(t internal.T, v interface{}) *ThatAssertion {
	return &ThatAssertion{
		t: t,
		v: v,
	}
}

// Equal assertion failed when got and expect are not `deeply equal`.
func (a *ThatAssertion) Equal(expect interface{}, msg ...string) {
	a.t.Helper()
	if !reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotEqual assertion failed when got and expect are `deeply equal`.
func (a *ThatAssertion) NotEqual(expect interface{}, msg ...string) {
	a.t.Helper()
	if reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect not (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// Same assertion failed when got and expect are not same.
func (a *ThatAssertion) Same(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v != expect {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotSame assertion failed when got and expect are same.
func (a *ThatAssertion) NotSame(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v == expect {
		str := fmt.Sprintf("expect not (%T) %v", expect, expect)
		fail(a.t, str, msg...)
	}
}

// TypeOf assertion failed when got and expect are not same type.
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

// Implements assertion failed when got doesn't implement expect.
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

// Has assertion failed when got is not has expect.
func (a *ThatAssertion) Has(expect interface{}, msg ...string) {
	a.t.Helper()

	m := reflect.ValueOf(a.v).MethodByName("Has")
	if !m.IsValid() {
		str := fmt.Sprintf("method 'Has' not found on type %T", a.v)
		fail(a.t, str, msg...)
		return
	}

	ret := m.Call([]reflect.Value{reflect.ValueOf(expect)})
	if !ret[0].Bool() {
		str := fmt.Sprintf("got (%T) %v not has (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// InSlice assertion failed when got is not in expect array & slice.
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

// NotInSlice assertion failed when got is in expect array & slice.
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

// SubInSlice assertion failed when got is not sub in expect array & slice.
func (a *ThatAssertion) SubInSlice(expect interface{}, msg ...string) {
	a.t.Helper()

	v1 := reflect.ValueOf(a.v)
	if v1.Kind() != reflect.Array && v1.Kind() != reflect.Slice {
		str := fmt.Sprintf("unsupported got value (%T) %v", a.v, a.v)
		fail(a.t, str, msg...)
		return
	}

	v2 := reflect.ValueOf(expect)
	if v2.Kind() != reflect.Array && v2.Kind() != reflect.Slice {
		str := fmt.Sprintf("unsupported expect value (%T) %v", expect, expect)
		fail(a.t, str, msg...)
		return
	}

	for i := 0; i < v1.Len(); i++ {
		for j := 0; j < v2.Len(); j++ {
			if reflect.DeepEqual(v1.Index(i).Interface(), v2.Index(j).Interface()) {
				return
			}
		}
	}

	str := fmt.Sprintf("got (%T) %v is not sub in (%T) %v", a.v, a.v, expect, expect)
	fail(a.t, str, msg...)
}

// InMapKeys assertion failed when got is not in keys of expect map.
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

// InMapValues assertion failed when got is not in values of expect map.
func (a *ThatAssertion) InMapValues(expect interface{}, msg ...string) {
	a.t.Helper()

	switch v := reflect.ValueOf(expect); v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if reflect.DeepEqual(a.v, v.MapIndex(key).Interface()) {
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
