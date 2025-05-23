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
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"testing"

	"github.com/lvan100/go-assert"
	"github.com/lvan100/go-assert/internal"
	"go.uber.org/mock/gomock"
)

func runCase(t *testing.T, f func(g *internal.MockT)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	g := internal.NewMockT(ctrl)
	g.EXPECT().Helper().AnyTimes()
	f(g)
}

func TestTrue(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.True(g, true)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got false but expect true"})
		assert.True(g, false)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got false but expect true, param (index=0)"})
		assert.True(g, false, "param (index=0)")
	})
}

func TestFalse(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.False(g, false)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got true but expect false"})
		assert.False(g, true)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got true but expect false, param (index=0)"})
		assert.False(g, true, "param (index=0)")
	})
}

func TestNil(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.Nil(g, nil)
	})
	runCase(t, func(g *internal.MockT) {
		assert.Nil(g, (*int)(nil))
	})
	runCase(t, func(g *internal.MockT) {
		var a []string
		assert.Nil(g, a)
	})
	runCase(t, func(g *internal.MockT) {
		var m map[string]string
		assert.Nil(g, m)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 3 but expect nil"})
		assert.Nil(g, 3)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 3 but expect nil, param (index=0)"})
		assert.Nil(g, 3, "param (index=0)")
	})
}

func TestNotNil(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.NotNil(g, 3)
	})
	runCase(t, func(g *internal.MockT) {
		a := make([]string, 0)
		assert.NotNil(g, a)
	})
	runCase(t, func(g *internal.MockT) {
		m := make(map[string]string)
		assert.NotNil(g, m)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got nil but expect not nil"})
		assert.NotNil(g, nil)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got nil but expect not nil, param (index=0)"})
		assert.NotNil(g, nil, "param (index=0)")
	})
}

func TestPanic(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.Panic(g, func() { panic("this is an error") }, "an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"did not panic"})
		assert.Panic(g, func() {}, "an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"invalid pattern"})
		assert.Panic(g, func() { panic("this is an error") }, "an error \\")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\""})
		assert.Panic(g, func() { panic("there's no error") }, "an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\", param (index=0)"})
		assert.Panic(g, func() { panic("there's no error") }, "an error", "param (index=0)")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\""})
		assert.Panic(g, func() { panic(errors.New("there's no error")) }, "an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"there's no error\" which does not match \"an error\""})
		assert.Panic(g, func() { panic(bytes.NewBufferString("there's no error")) }, "an error")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got \"[there's no error]\" which does not match \"an error\""})
		assert.Panic(g, func() { panic([]string{"there's no error"}) }, "an error")
	})
}

func TestThat_Equal(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, 0).Equal(0)
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, []string{"a"}).Equal([]string{"a"})
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, struct {
			text string
		}{text: "a"}).Equal(struct {
			text string
		}{text: "a"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (struct { Text string }) {a} but expect (struct { Text string \"json:\\\"text\\\"\" }) {a}"})
		assert.That(g, struct {
			Text string
		}{Text: "a"}).Equal(struct {
			Text string `json:"text"`
		}{Text: "a"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (struct { text string }) {a} but expect (struct { msg string }) {a}"})
		assert.That(g, struct {
			text string
		}{text: "a"}).Equal(struct {
			msg string
		}{msg: "a"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 0 but expect (string) 0"})
		assert.That(g, 0).Equal("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 0 but expect (string) 0, param (index=0)"})
		assert.That(g, 0).Equal("0", "param (index=0)")
	})
}

func TestThat_NotEqual(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, "0").NotEqual(0)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got ([]string) [a] but expect not ([]string) [a]"})
		assert.That(g, []string{"a"}).NotEqual([]string{"a"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (string) 0 but expect not (string) 0"})
		assert.That(g, "0").NotEqual("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (string) 0 but expect not (string) 0, param (index=0)"})
		assert.That(g, "0").NotEqual("0", "param (index=0)")
	})
}

func TestThat_Same(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, "0").Same("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 0 but expect (string) 0"})
		assert.That(g, 0).Same("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 0 but expect (string) 0, param (index=0)"})
		assert.That(g, 0).Same("0", "param (index=0)")
	})
}

func TestThat_NotSame(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, "0").NotSame(0)
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"expect not (string) 0"})
		assert.That(g, "0").NotSame("0")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"expect not (string) 0, param (index=0)"})
		assert.That(g, "0").NotSame("0", "param (index=0)")
	})
}

func TestThat_TypeOf(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, new(int)).TypeOf((*int)(nil))
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got type (string) but expect type (fmt.Stringer)"})
		assert.That(g, "string").TypeOf((*fmt.Stringer)(nil))
	})
}

func TestThat_Implements(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		assert.That(g, errors.New("error")).Implements((*error)(nil))
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"expect should be interface"})
		assert.That(g, new(int)).Implements((*int)(nil))
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got type (*int) but expect type (io.Reader)"})
		assert.That(g, new(int)).Implements((*io.Reader)(nil))
	})
}

type Node struct{}

func (t *Node) Has(key string) (bool, error) {
	return false, nil
}

func (t *Node) Contains(key string) (bool, error) {
	return false, nil
}

type Tree struct {
	Keys []string
}

func (t *Tree) Has(key string) bool {
	return slices.Contains(t.Keys, key)
}

func (t *Tree) Contains(key string) bool {
	return slices.Contains(t.Keys, key)
}

func TestThat_Has(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"method 'Has' not found on type int"})
		assert.That(g, 1).Has("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"method 'Has' must return only a bool"})
		assert.That(g, &Node{}).Has("2")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (*assert_test.Tree) &{[]} not has (string) 2"})
		assert.That(g, &Tree{}).Has("2")
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, &Tree{Keys: []string{"1"}}).Has("1")
	})
}

func TestThat_Contains(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"method 'Contains' not found on type int"})
		assert.That(g, 1).Contains("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"method 'Contains' must return only a bool"})
		assert.That(g, &Node{}).Contains("2")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (*assert_test.Tree) &{[]} not contains (string) 2"})
		assert.That(g, &Tree{}).Contains("2")
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, &Tree{Keys: []string{"1"}}).Contains("1")
	})
}

func TestThat_InSlice(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"unsupported expect value (string) 1"})
		assert.That(g, 1).InSlice("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 1 is not in ([]string) [1]"})
		assert.That(g, 1).InSlice([]string{"1"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int64) 1 is not in ([]int64) [3 2]"})
		assert.That(g, int64(1)).InSlice([]int64{3, 2})
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, int64(1)).InSlice([]int64{3, 2, 1})
		assert.That(g, "1").InSlice([]string{"3", "2", "1"})
	})
}

func TestThat_NotInSlice(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"unsupported expect value (string) 1"})
		assert.That(g, 1).NotInSlice("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got type (int) doesn't match expect type ([]string)"})
		assert.That(g, 1).NotInSlice([]string{"1"})
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (string) 1 is in ([]string) [3 2 1]"})
		assert.That(g, "1").NotInSlice([]string{"3", "2", "1"})
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, int64(1)).NotInSlice([]int64{3, 2})
	})
}

func TestThat_InMapKeys(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"unsupported expect value (string) 1"})
		assert.That(g, 1).InMapKeys("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 1 is not in keys of (map[string]string) map[1:1]"})
		assert.That(g, 1).InMapKeys(map[string]string{"1": "1"})
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, int64(1)).InMapKeys(map[int64]int64{3: 1, 2: 2, 1: 3})
		assert.That(g, "1").InMapKeys(map[string]string{"3": "1", "2": "2", "1": "3"})
	})
}

func TestThat_InMapValues(t *testing.T) {
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"unsupported expect value (string) 1"})
		assert.That(g, 1).InMapValues("1")
	})
	runCase(t, func(g *internal.MockT) {
		g.EXPECT().Error([]interface{}{"got (int) 1 is not in values of (map[string]string) map[1:1]"})
		assert.That(g, 1).InMapValues(map[string]string{"1": "1"})
	})
	runCase(t, func(g *internal.MockT) {
		assert.That(g, int64(1)).InMapValues(map[int64]int64{3: 1, 2: 2, 1: 3})
		assert.That(g, "1").InMapValues(map[string]string{"3": "1", "2": "2", "1": "3"})
	})
}
