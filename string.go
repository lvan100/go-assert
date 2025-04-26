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

package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/lvan100/go-assert/internal"
)

// StringAssertion assertion for type string.
type StringAssertion struct {
	t internal.T
	v string
}

// ThatString returns an assertion for type string.
func ThatString(t internal.T, v string) *StringAssertion {
	return &StringAssertion{
		t: t,
		v: v,
	}
}

// Equal assertion failed when got and expect are not `deeply equal`.
func (a *StringAssertion) Equal(expect string, msg ...string) {
	a.t.Helper()
	if !reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// NotEqual assertion failed when got and expect are `deeply equal`.
func (a *StringAssertion) NotEqual(expect string, msg ...string) {
	a.t.Helper()
	if reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf("got (%T) %v but expect not (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// JsonEqual assertion failed when got and expect are not `json equal`.
func (a *StringAssertion) JsonEqual(expect string, msg ...string) {
	a.t.Helper()
	var gotJson interface{}
	if err := json.Unmarshal([]byte(a.v), &gotJson); err != nil {
		fail(a.t, err.Error(), msg...)
		return
	}
	var expectJson interface{}
	if err := json.Unmarshal([]byte(expect), &expectJson); err != nil {
		fail(a.t, err.Error(), msg...)
		return
	}
	if !reflect.DeepEqual(gotJson, expectJson) {
		str := fmt.Sprintf("got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

func matches(t internal.T, got string, expr string, msg ...string) {
	t.Helper()
	if ok, err := regexp.MatchString(expr, got); err != nil {
		fail(t, "invalid pattern", msg...)
	} else if !ok {
		str := fmt.Sprintf("got %q which does not match %q", got, expr)
		fail(t, str, msg...)
	}
}

// Matches assertion failed when got doesn't match expr expression.
func (a *StringAssertion) Matches(expr string, msg ...string) {
	a.t.Helper()
	matches(a.t, a.v, expr, msg...)
}

// EqualFold assertion failed when v doesn't equal to `s` under Unicode case-folding.
func (a *StringAssertion) EqualFold(s string, msg ...string) {
	a.t.Helper()
	if !strings.EqualFold(a.v, s) {
		fail(a.t, fmt.Sprintf("'%s' doesn't equal fold to '%s'", a.v, s), msg...)
	}
}

// HasPrefix assertion failed when v doesn't have prefix `prefix`.
func (a *StringAssertion) HasPrefix(prefix string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.HasPrefix(a.v, prefix) {
		fail(a.t, fmt.Sprintf("'%s' doesn't have prefix '%s'", a.v, prefix), msg...)
	}
	return a
}

// HasSuffix assertion failed when v doesn't have suffix `suffix`.
func (a *StringAssertion) HasSuffix(suffix string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.HasSuffix(a.v, suffix) {
		fail(a.t, fmt.Sprintf("'%s' doesn't have suffix '%s'", a.v, suffix), msg...)
	}
	return a
}

// Contains assertion failed when v doesn't contain substring `substr`.
func (a *StringAssertion) Contains(substr string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.Contains(a.v, substr) {
		fail(a.t, fmt.Sprintf("'%s' doesn't contain substr '%s'", a.v, substr), msg...)
	}
	return a
}
