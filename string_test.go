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

package assert_test

import (
	"testing"

	"github.com/lvan100/go-assert"
	"github.com/lvan100/go-assert/internal"
)

func TestString_Equal(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "0").Equal("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings not equal:
    got: (string) "0"
 expect: (string) "1"`})
		assert.ThatString(g, "0").Equal("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings not equal:
    got: (string) "0"
 expect: (string) "1"
message: param (index=0)`})
		assert.ThatString(g, "0").Equal("1", "param (index=0)")
	})
}

func TestString_NotEqual(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "0").NotEqual("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings are equal:
    got: (string) "0"
 expect: not equal to "0"`})
		assert.ThatString(g, "0").NotEqual("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings are equal:
    got: (string) "0"
 expect: not equal to "0"
message: param (index=0)`})
		assert.ThatString(g, "0").NotEqual("0", "param (index=0)")
	})
}

func TestString_JSONEqual(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, `{"a":0,"b":1}`).JSONEqual(`{"b":1,"a":0}`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`invalid JSON in got value:
    got: (string) "this is an error"
 expect: (string) "[{\"b\":1},{\"a\":0}]"
  error: invalid character 'h' in literal true (expecting 'r')`})
		assert.ThatString(g, `this is an error`).JSONEqual(`[{"b":1},{"a":0}]`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`invalid JSON in expect value:
    got: (string) "{\"a\":0,\"b\":1}"
 expect: (string) "this is an error"
  error: invalid character 'h' in literal true (expecting 'r')`})
		assert.ThatString(g, `{"a":0,"b":1}`).JSONEqual(`this is an error`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`JSON structures are not equal:
    got: (string) "{\"a\":0,\"b\":1}"
 expect: (string) "[{\"b\":1},{\"a\":0}]"`})
		assert.ThatString(g, `{"a":0,"b":1}`).JSONEqual(`[{"b":1},{"a":0}]`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`JSON structures are not equal:
    got: (string) "{\"a\":0}"
 expect: (string) "{\"a\":1}"
message: param (index=0)`})
		assert.ThatString(g, `{"a":0}`).JSONEqual(`{"a":1}`, "param (index=0)")
	})
}

func TestString_Matches(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "this is an error").Matches("this is an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"string does not match the pattern:\n    got: (string) \"this is an error\"\n expect: to match regex \"an error (\"\n  error: error parsing regexp: missing closing ): `an error (`"})
		assert.ThatString(g, "this is an error").Matches("an error (")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not match the pattern:
    got: (string) "there's no error"
 expect: to match regex "an error"`})
		assert.ThatString(g, "there's no error").Matches("an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not match the pattern:
    got: (string) "there's no error"
 expect: to match regex "an error"
message: param (index=0)`})
		assert.ThatString(g, "there's no error").Matches("an error", "param (index=0)")
	})
}

func TestString_EqualFold(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello, world!").EqualFold("Hello, World!")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings are not equal under case-folding:
    got: (string) "hello, world!"
 expect: (string) "xxx"`})
		assert.ThatString(g, "hello, world!").EqualFold("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`strings are not equal under case-folding:
    got: (string) "hello, world!"
 expect: (string) "xxx"
message: param (index=0)`})
		assert.ThatString(g, "hello, world!").EqualFold("xxx", "param (index=0)")
	})
}

func TestString_HasPrefix(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello, world!").HasPrefix("hello")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not start with the specified prefix:
    got: (string) "hello, world!"
 expect: to have prefix "xxx"`})
		assert.ThatString(g, "hello, world!").HasPrefix("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not start with the specified prefix:
    got: (string) "hello, world!"
 expect: to have prefix "xxx"
message: param (index=0)`})
		assert.ThatString(g, "hello, world!").HasPrefix("xxx", "param (index=0)")
	})
}

func TestString_HasSuffix(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello, world!").HasSuffix("world!")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not end with the specified suffix:
    got: (string) "hello, world!"
 expect: to have suffix "xxx"`})
		assert.ThatString(g, "hello, world!").HasSuffix("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not end with the specified suffix:
    got: (string) "hello, world!"
 expect: to have suffix "xxx"
message: param (index=0)`})
		assert.ThatString(g, "hello, world!").HasSuffix("xxx", "param (index=0)")
	})
}

func TestString_Contains(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello, world!").Contains("hello")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not contain the specified substring:
    got: (string) "hello, world!"
 expect: to contain substring "xxx"`})
		assert.ThatString(g, "hello, world!").Contains("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string does not contain the specified substring:
    got: (string) "hello, world!"
 expect: to contain substring "xxx"
message: param (index=0)`})
		assert.ThatString(g, "hello, world!").Contains("xxx", "param (index=0)")
	})
}

func TestString_IsEmpty(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "").IsEmpty()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not empty:
    got: (string) "hello"
 expect: empty string`})
		assert.ThatString(g, "hello").IsEmpty()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not empty:
    got: (string) "hello"
 expect: empty string
message: param (index=0)`})
		assert.ThatString(g, "hello").IsEmpty("param (index=0)")
	})
}

func TestString_IsNotEmpty(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello").IsNotEmpty()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is empty:
    got: (string) ""
 expect: non-empty string`})
		assert.ThatString(g, "").IsNotEmpty()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is empty:
    got: (string) ""
 expect: non-empty string
message: param (index=0)`})
		assert.ThatString(g, "").IsNotEmpty("param (index=0)")
	})
}

func TestString_IsBlank(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "   ").IsBlank()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-whitespace characters:
    got: (string) "hello"
 expect: blank string`})
		assert.ThatString(g, "hello").IsBlank()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-whitespace characters:
    got: (string) "hello"
 expect: blank string
message: param (index=0)`})
		assert.ThatString(g, "hello").IsBlank("param (index=0)")
	})
}

func TestString_IsNotBlank(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello").IsNotBlank()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is blank:
    got: (string) "   "
 expect: non-blank string`})
		assert.ThatString(g, "   ").IsNotBlank()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is blank:
    got: (string) "   "
 expect: non-blank string
message: param (index=0)`})
		assert.ThatString(g, "   ").IsNotBlank("param (index=0)")
	})
}

func TestString_IsLowerCase(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "hello").IsLowerCase()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains uppercase characters:
    got: (string) "Hello"
 expect: lowercase string`})
		assert.ThatString(g, "Hello").IsLowerCase()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains uppercase characters:
    got: (string) "Hello"
 expect: lowercase string
message: param (index=0)`})
		assert.ThatString(g, "Hello").IsLowerCase("param (index=0)")
	})
}

func TestString_IsUpperCase(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "HELLO").IsUpperCase()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains lowercase characters:
    got: (string) "Hello"
 expect: uppercase string`})
		assert.ThatString(g, "Hello").IsUpperCase()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains lowercase characters:
    got: (string) "Hello"
 expect: uppercase string
message: param (index=0)`})
		assert.ThatString(g, "Hello").IsUpperCase("param (index=0)")
	})
}

func TestString_IsNumeric(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "12345").IsNumeric()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-numeric characters:
    got: (string) "123a45"
 expect: numeric string`})
		assert.ThatString(g, "123a45").IsNumeric()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-numeric characters:
    got: (string) "123a45"
 expect: numeric string
message: param (index=0)`})
		assert.ThatString(g, "123a45").IsNumeric("param (index=0)")
	})
}

func TestString_IsAlpha(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "abcdef").IsAlpha()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-alphabetic characters:
    got: (string) "abc123"
 expect: alphabetic string`})
		assert.ThatString(g, "abc123").IsAlpha()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-alphabetic characters:
    got: (string) "abc123"
 expect: alphabetic string
message: param (index=0)`})
		assert.ThatString(g, "abc123").IsAlpha("param (index=0)")
	})
}

func TestString_IsAlphaNumeric(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "abc123").IsAlphaNumeric()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-alphanumeric characters:
    got: (string) "abc@123"
 expect: alphanumeric string`})
		assert.ThatString(g, "abc@123").IsAlphaNumeric()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string contains non-alphanumeric characters:
    got: (string) "abc@123"
 expect: alphanumeric string
message: param (index=0)`})
		assert.ThatString(g, "abc@123").IsAlphaNumeric("param (index=0)")
	})
}

func TestString_IsEmail(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "test@example.com").IsEmail()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid email:
    got: (string) "invalid-email"
 expect: valid email address`})
		assert.ThatString(g, "invalid-email").IsEmail()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid email:
    got: (string) "invalid-email"
 expect: valid email address
message: param (index=0)`})
		assert.ThatString(g, "invalid-email").IsEmail("param (index=0)")
	})
}

func TestString_IsURL(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "https://www.example.com").IsURL()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid URL:
    got: (string) "invalid-url"
 expect: valid URL`})
		assert.ThatString(g, "invalid-url").IsURL()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid URL:
    got: (string) "invalid-url"
 expect: valid URL
message: param (index=0)`})
		assert.ThatString(g, "invalid-url").IsURL("param (index=0)")
	})
}

func TestString_IsIP(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "192.168.1.1").IsIP()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid IP:
    got: (string) "invalid-ip"
 expect: valid IP address`})
		assert.ThatString(g, "invalid-ip").IsIP()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid IP:
    got: (string) "invalid-ip"
 expect: valid IP address
message: param (index=0)`})
		assert.ThatString(g, "invalid-ip").IsIP("param (index=0)")
	})
}

func TestString_IsHex(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "abcdef123456").IsHex()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid hexadecimal:
    got: (string) "abcdefg"
 expect: valid hexadecimal number`})
		assert.ThatString(g, "abcdefg").IsHex()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid hexadecimal:
    got: (string) "abcdefg"
 expect: valid hexadecimal number
message: param (index=0)`})
		assert.ThatString(g, "abcdefg").IsHex("param (index=0)")
	})
}

func TestString_IsBase64(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.ThatString(g, "SGVsbG8gd29ybGQ=").IsBase64()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid Base64:
    got: (string) "invalid-base64"
 expect: valid Base64 encoded string`})
		assert.ThatString(g, "invalid-base64").IsBase64()
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{`string is not a valid Base64:
    got: (string) "invalid-base64"
 expect: valid Base64 encoded string
message: param (index=0)`})
		assert.ThatString(g, "invalid-base64").IsBase64("param (index=0)")
	})
}
