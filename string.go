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

// StringAssertion encapsulates a string value and a test handler for making assertions on the string.
type StringAssertion struct {
	t internal.T
	v string
}

// ThatString returns a StringAssertion for the given testing object and string value.
func ThatString(t internal.T, v string) *StringAssertion {
	return &StringAssertion{
		t: t,
		v: v,
	}
}

// Length reports a test failure if the actual string's length is not equal to the expected length.
func (a *StringAssertion) Length(length int, msg ...string) *StringAssertion {
	a.t.Helper()
	if len(a.v) != length {
		str := fmt.Sprintf("length mismatch, got (%T) %v with length %d but expect length %d", a.v, a.v, len(a.v), length)
		fail(a.t, str, msg...)
	}
	return a
}

// Equal reports a test failure if the actual string is not equal to the expected string.
func (a *StringAssertion) Equal(expect string, msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v != expect {
		str := fmt.Sprintf("strings are not equal, got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
	return a
}

// NotEqual reports a test failure if the actual string is equal to the given string.
func (a *StringAssertion) NotEqual(expect string, msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v == expect {
		str := fmt.Sprintf("strings are equal, got (%T) %v but expect not", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// JsonEqual unmarshals both the actual and expected JSON strings into generic interfaces,
// then reports a test failure if their resulting structures are not deeply equal.
// If either string is invalid JSON, the test will fail with the unmarshal error.
func (a *StringAssertion) JsonEqual(expect string, msg ...string) {
	a.t.Helper()
	var gotJson interface{}
	if err := json.Unmarshal([]byte(a.v), &gotJson); err != nil {
		str := fmt.Sprintf("invalid JSON in got value, got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
		return
	}
	var expectJson interface{}
	if err := json.Unmarshal([]byte(expect), &expectJson); err != nil {
		str := fmt.Sprintf("invalid JSON in expect value, got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
		return
	}
	if !reflect.DeepEqual(gotJson, expectJson) {
		str := fmt.Sprintf("JSON structures are not equal, got (%T) %v but expect (%T) %v", a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
	}
}

// Matches reports a test failure if the actual string does not match the given regular expression.
func (a *StringAssertion) Matches(expr string, msg ...string) {
	a.t.Helper()
	if ok, err := regexp.MatchString(expr, a.v); err != nil || !ok {
		str := fmt.Sprintf("string does not match the pattern, got (%T) %v but expect to match regex %q", a.v, a.v, expr)
		fail(a.t, str, msg...)
	}
}

// EqualFold reports a test failure if the actual string and the given string
// are not equal under Unicode case-folding.
func (a *StringAssertion) EqualFold(s string, msg ...string) {
	a.t.Helper()
	if !strings.EqualFold(a.v, s) {
		str := fmt.Sprintf("strings are not equal under case-folding, got (%T) %v but expect (%T) %v", a.v, a.v, s, s)
		fail(a.t, str, msg...)
	}
}

// HasPrefix fails the test if the actual string does not start with the specified prefix.
func (a *StringAssertion) HasPrefix(prefix string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.HasPrefix(a.v, prefix) {
		str := fmt.Sprintf("string does not start with the specified prefix, got (%T) %v but expect to have prefix %q", a.v, a.v, prefix)
		fail(a.t, str, msg...)
	}
	return a
}

// HasSuffix fails the test if the actual string does not end with the specified suffix.
func (a *StringAssertion) HasSuffix(suffix string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.HasSuffix(a.v, suffix) {
		str := fmt.Sprintf("string does not end with the specified suffix, got (%T) %v but expect to have suffix %q", a.v, a.v, suffix)
		fail(a.t, str, msg...)
	}
	return a
}

// Contains fails the test if the actual string does not contain the specified substring.
func (a *StringAssertion) Contains(substr string, msg ...string) *StringAssertion {
	a.t.Helper()
	if !strings.Contains(a.v, substr) {
		str := fmt.Sprintf("string does not contain the specified substring, got (%T) %v but expect to contain substring %q", a.v, a.v, substr)
		fail(a.t, str, msg...)
	}
	return a
}

// IsEmpty reports a test failure if the actual string is not empty.
func (a *StringAssertion) IsEmpty(msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v != "" {
		str := fmt.Sprintf("string is not empty, got (%T) %v but expect empty string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsNotEmpty reports a test failure if the actual string is empty.
func (a *StringAssertion) IsNotEmpty(msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v == "" {
		str := fmt.Sprintf("string is empty, got (%T) %v but expect non-empty string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsBlank reports a test failure if the actual string is not blank (i.e., contains non-whitespace characters).
func (a *StringAssertion) IsBlank(msg ...string) *StringAssertion {
	a.t.Helper()
	if strings.TrimSpace(a.v) != "" {
		str := fmt.Sprintf("string contains non-whitespace characters, got (%T) %v but expect blank string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsNotBlank reports a test failure if the actual string is blank (i.e., empty or contains only whitespace characters).
func (a *StringAssertion) IsNotBlank(msg ...string) *StringAssertion {
	a.t.Helper()
	if strings.TrimSpace(a.v) == "" {
		str := fmt.Sprintf("string is blank, got (%T) %v but expect non-blank string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsLowerCase reports a test failure if the actual string contains any uppercase characters.
func (a *StringAssertion) IsLowerCase(msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v != strings.ToLower(a.v) {
		str := fmt.Sprintf("string contains uppercase characters, got (%T) %v but expect lowercase string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsUpperCase reports a test failure if the actual string contains any lowercase characters.
func (a *StringAssertion) IsUpperCase(msg ...string) *StringAssertion {
	a.t.Helper()
	if a.v != strings.ToUpper(a.v) {
		str := fmt.Sprintf("string contains lowercase characters, got (%T) %v but expect uppercase string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsNumeric reports a test failure if the actual string contains any non-numeric characters.
func (a *StringAssertion) IsNumeric(msg ...string) *StringAssertion {
	a.t.Helper()
	for _, r := range a.v {
		if r < '0' || r > '9' {
			str := fmt.Sprintf("string contains non-numeric characters, got (%T) %v but expect numeric string", a.v, a.v)
			fail(a.t, str, msg...)
			break
		}
	}
	return a
}

// IsAlpha reports a test failure if the actual string contains any non-alphabetic characters.
func (a *StringAssertion) IsAlpha(msg ...string) *StringAssertion {
	a.t.Helper()
	for _, r := range a.v {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			str := fmt.Sprintf("string contains non-alphabetic characters, got (%T) %v but expect alphabetic string", a.v, a.v)
			fail(a.t, str, msg...)
			break
		}
	}
	return a
}

// IsAlphaNumeric reports a test failure if the actual string contains any non-alphanumeric characters.
func (a *StringAssertion) IsAlphaNumeric(msg ...string) *StringAssertion {
	a.t.Helper()
	for _, r := range a.v {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			str := fmt.Sprintf("string contains non-alphanumeric characters, got (%T) %v but expect alphanumeric string", a.v, a.v)
			fail(a.t, str, msg...)
			break
		}
	}
	return a
}

// IsEmail reports a test failure if the actual string is not a valid email address.
func (a *StringAssertion) IsEmail(msg ...string) *StringAssertion {
	a.t.Helper()
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok, err := regexp.MatchString(emailRegex, a.v); err != nil || !ok {
		str := fmt.Sprintf("string is not a valid email, got (%T) %v but expect valid email address", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsURL reports a test failure if the actual string is not a valid URL.
func (a *StringAssertion) IsURL(msg ...string) *StringAssertion {
	a.t.Helper()
	urlRegex := `^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`
	if ok, err := regexp.MatchString(urlRegex, a.v); err != nil || !ok {
		str := fmt.Sprintf("string is not a valid URL, got (%T) %v but expect valid URL", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsIP reports a test failure if the actual string is not a valid IP address.
func (a *StringAssertion) IsIP(msg ...string) *StringAssertion {
	a.t.Helper()
	ipRegex := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	if ok, err := regexp.MatchString(ipRegex, a.v); err != nil || !ok {
		str := fmt.Sprintf("string is not a valid IP, got (%T) %v but expect valid IP address", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsHex reports a test failure if the actual string is not a valid hexadecimal number.
func (a *StringAssertion) IsHex(msg ...string) *StringAssertion {
	a.t.Helper()
	hexRegex := `^[0-9a-fA-F]+$`
	if ok, err := regexp.MatchString(hexRegex, a.v); err != nil || !ok {
		str := fmt.Sprintf("string is not a valid hexadecimal, got (%T) %v but expect valid hexadecimal number", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsBase64 reports a test failure if the actual string is not a valid Base64 encoded string.
func (a *StringAssertion) IsBase64(msg ...string) *StringAssertion {
	a.t.Helper()
	base64Regex := `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`
	if ok, err := regexp.MatchString(base64Regex, a.v); err != nil || !ok {
		str := fmt.Sprintf("string is not a valid Base64, got (%T) %v but expect valid Base64 encoded string", a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}
