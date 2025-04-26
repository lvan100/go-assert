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
	"github.com/lvan100/go-assert/internal"
)

// ErrorAssertion assertion for type string.
type ErrorAssertion struct {
	t internal.T
	v error
}

// ThatError returns an assertion for type string.
func ThatError(t internal.T, v error) *ErrorAssertion {
	return &ErrorAssertion{
		t: t,
		v: v,
	}
}

// Matches assertion failed when got doesn't match expr expression.
func (a *ErrorAssertion) Matches(expr string, msg ...string) {
	a.t.Helper()
	if a.v == nil {
		fail(a.t, "expect not nil error", msg...)
		return
	}
	matches(a.t, a.v.Error(), expr, msg...)
}
