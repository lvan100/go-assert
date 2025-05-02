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
	"errors"
	"fmt"
	"strings"

	"github.com/lvan100/go-assert/internal"
)

// ErrorAssertion provides assertion methods for values of type error.
// It is used to perform validations on error values in test cases.
type ErrorAssertion struct {
	t internal.T
	v error
}

// ThatError returns a new ErrorAssertion for the given error value.
func ThatError(t internal.T, v error) *ErrorAssertion {
	return &ErrorAssertion{
		t: t,
		v: v,
	}
}

// IsNil reports a test failure if the error is not nil.
func (a *ErrorAssertion) IsNil(msg ...string) {
	a.t.Helper()
	if a.v != nil {
		fail(a.t, "expect nil error, got: "+a.v.Error(), msg...)
	}
}

// IsNotNil reports a test failure if the error is nil.
func (a *ErrorAssertion) IsNotNil(msg ...string) {
	a.t.Helper()
	if a.v == nil {
		fail(a.t, "expect not nil error", msg...)
	}
}

// Is reports a test failure if the error is not the same as the given error.
func (a *ErrorAssertion) Is(target error, msg ...string) {
	a.t.Helper()
	if !errors.Is(target, a.v) {
		fail(a.t, "expect error: "+target.Error()+", got: "+a.v.Error(), msg...)
	}
}

// IsNot reports a test failure if the error is the same as the given error.
func (a *ErrorAssertion) IsNot(target error, msg ...string) {
	a.t.Helper()
	if errors.Is(target, a.v) {
		fail(a.t, "expect error not to be: "+target.Error(), msg...)
	}
}

// As checks if the error can be converted to the target type.
func (a *ErrorAssertion) As(target interface{}, msg ...string) {
	a.t.Helper()
	if !errors.As(a.v, &target) {
		fail(a.t, "expect error to be of type: "+fmt.Sprintf("%T", target), msg...)
	}
}

// ContainsMessage reports a test failure if the error message does not contain the given substring.
func (a *ErrorAssertion) ContainsMessage(substring string, msg ...string) {
	a.t.Helper()
	if a.v == nil {
		fail(a.t, "expect not nil error", msg...)
		return
	}
	if !strings.Contains(a.v.Error(), substring) {
		fail(a.t, "expect error message to contain: "+substring+", got: "+a.v.Error(), msg...)
	}
}

// Matches reports a test failure if the error string does not match the given expression.
// It expects a non-nil error and uses the provided expression (typically a regex)
// to validate the error message content. Optional custom failure messages can be provided.
func (a *ErrorAssertion) Matches(expr string, msg ...string) {
	a.t.Helper()
	if a.v == nil {
		fail(a.t, "expect not nil error", msg...)
		return
	}
	matches(a.t, a.v.Error(), expr, msg...)
}
