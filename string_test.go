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

func TestString_JsonEqual(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, `{"a":0,"b":1}`).JsonEqual(`{"b":1,"a":0}`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"invalid character 'h' in literal true (expecting 'r')"})
		assert.String(g, `this is an error`).JsonEqual(`[{"b":1},{"a":0}]`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"invalid character 'h' in literal true (expecting 'r')"})
		assert.String(g, `{"a":0,"b":1}`).JsonEqual(`this is an error`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (string) {\"a\":0,\"b\":1} but expect (string) [{\"b\":1},{\"a\":0}]"})
		assert.String(g, `{"a":0,"b":1}`).JsonEqual(`[{"b":1},{"a":0}]`)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (string) {\"a\":0} but expect (string) {\"a\":1}; param (index=0)"})
		assert.String(g, `{"a":0}`).JsonEqual(`{"a":1}`, "param (index=0)")
	})
}

func TestString_Matches(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, "this is an error").Matches("this is an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"invalid pattern"})
		assert.String(g, "this is an error").Matches("an error \\")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\""})
		assert.String(g, "there's no error").Matches("an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\"; param (index=0)"})
		assert.String(g, "there's no error").Matches("an error", "param (index=0)")
	})
}

func TestString_EqualFold(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, "hello, world!").EqualFold("Hello, World!")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't equal fold to 'xxx'"})
		assert.String(g, "hello, world!").EqualFold("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't equal fold to 'xxx'; param (index=0)"})
		assert.String(g, "hello, world!").EqualFold("xxx", "param (index=0)")
	})
}

func TestString_HasPrefix(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, "hello, world!").HasPrefix("hello")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have prefix 'xxx'"})
		assert.String(g, "hello, world!").HasPrefix("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have prefix 'xxx'; param (index=0)"})
		assert.String(g, "hello, world!").HasPrefix("xxx", "param (index=0)")
	})
}

func TestString_HasSuffix(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, "hello, world!").HasSuffix("world!")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have suffix 'xxx'"})
		assert.String(g, "hello, world!").HasSuffix("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have suffix 'xxx'; param (index=0)"})
		assert.String(g, "hello, world!").HasSuffix("xxx", "param (index=0)")
	})
}

func TestString_Contains(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.String(g, "hello, world!").Contains("hello")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't contain substr 'xxx'"})
		assert.String(g, "hello, world!").Contains("xxx")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't contain substr 'xxx'; param (index=0)"})
		assert.String(g, "hello, world!").Contains("xxx", "param (index=0)")
	})
}
