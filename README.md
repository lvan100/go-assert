# go-assert

**go-assert** 是一个简洁、高效且可读性极强的 Go 单元测试断言库，旨在帮助开发者编写更健壮、自然流畅的测试代码。

## ✨ 特性亮点

- 💬 **自然语言风格**：`That(got).Equal(expect)` 更易读、更易懂
- 🔄 **明确参数语义**：清晰区分 `got` 与 `expect`，不再弄反顺序
- ⚡ **高性能实现**：泛型支持，避免过度反射带来的性能损耗
- ✅ **丰富断言能力**：涵盖常用断言类型，满足大多数测试场景

## 📦 安装方式

```bash
go get github.com/lvan100/go-assert
```

## 🤔 为什么选择 go-assert？

在使用传统断言库（如 `testify`）时，我们常常会困惑哪个是实际值（got），哪个是期望值（expect）。  
**go-assert** 通过自然语言风格的 API 设计消除了这种困扰，让测试代码更易于理解和维护。

此外，它还借鉴了多个优秀库的优点，并以更现代的方式实现，比如泛型支持、低反射开销等。

## 🧪 快速上手

### ✅ 简洁断言（函数式）

适用于简单值判断：

```go
assert.True(t, isValid, "should be true")
assert.False(t, isClosed)
assert.Nil(t, result)
assert.NotNil(t, user, "user should not be nil")
assert.Panic(t, func () { panic("oops") }, "oops")
```

### 🔗 链式断言（更语义化）

#### That：适用于任意值

```go
assert.That(t, got).Equal(expect)
assert.That(t, got).NotEqual(expect)
assert.That(t, got).Same(expect)         // 同一实例
assert.That(t, got).NotSame(expect)
assert.That(t, got).TypeOf(MyStruct{})
assert.That(t, got).Implements((*io.Reader)(nil))

assert.That(t, got).Has(field)
assert.That(t, got).Contains(item)

assert.That(t, got).InSlice(slice)
assert.That(t, got).NotInSlice(slice)
assert.That(t, got).InMapKeys(mapVar)
assert.That(t, got).InMapValues(mapVar)
```

#### ThatError：专为 `error` 设计

```go
assert.ThatError(t, err).Matches("timeout")
```

#### ThatString：字符串专用断言器

```go
assert.ThatString(t, got).Equal("hello")
assert.ThatString(t, got).NotEqual("bye")
assert.ThatString(t, got).JsonEqual(`{"a":1}`)
assert.ThatString(t, got).Matches("^he.*")
assert.ThatString(t, got).EqualFold("Hello")
assert.ThatString(t, got).HasPrefix("he")
assert.ThatString(t, got).HasSuffix("lo")
assert.ThatString(t, got).Contains("ell")
```

## 💡 设计理念

- 🧠 **语义明确**：`got` 和 `expect` 顺序固定，减少思考负担
- 🧩 **断言器分工清晰**：`That` / `ThatError` / `ThatString` 分别适用于通用值、错误、字符串
- 🛡️ **泛型保障类型安全**：提升 IDE 支持和运行稳定性
- 🧰 **丰富断言方法**：满足从基本值到复杂结构的各种需求

## ✅ 示例测试

```go
func TestLogin(t *testing.T) {
user, err := Login("admin", "1234")

assert.ThatError(t, err).Matches("invalid password")
assert.That(t, user).NotNil("user should not be nil")
assert.ThatString(t, user.Name).HasPrefix("admin")
}
```

## 📜 License

MIT License —— 免费使用，欢迎贡献！

## 🚀 让测试更自然、更清晰、更强大 —— 快使用 `go-assert` 吧！

