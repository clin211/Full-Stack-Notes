# fmt

> `fmt` 包实现了类似 C 语言 `printf` 和 `scanf` 的格式化 I/O 功能，但更加简洁。它是 Go 中最常用的标准库之一，无论是调试输出、日志记录还是字符串格式化，都离不开它。本文基于 Go 1.25 源码，带你深入理解 `fmt` 包的设计与使用。

## 整体概览

`fmt` 包的源码由四个核心文件组成：

| 文件 | 职责 |
| ---- | ---- |
| `print.go` | 输出类函数（Printf/Sprintf/Fprint 等）及核心格式化逻辑 |
| `format.go` | 底层格式化引擎（整数、浮点数、字符串的格式化） |
| `scan.go` | 输入类函数（Scanf/Sscan/Fscan 等） |
| `errors.go` | `Errorf` 函数及错误包装（`%w`） |

先看一张全景图：

```text
fmt 包
├── 输出（Printing）
│   ├── Print / Println / Printf        → 标准输出
│   ├── Sprint / Sprintln / Sprintf     → 返回字符串
│   ├── Fprint / Fprintln / Fprintf     → 写入 io.Writer
│   ├── Append / Appendln / Appendf     → 追加到 []byte
│   └── Errorf                          → 创建 error
├── 输入（Scanning）
│   ├── Scan / Scanln / Scanf           → 从标准输入读取
│   ├── Sscan / Sscanln / Sscanf        → 从字符串读取
│   └── Fscan / Fscanln / Fscanf        → 从 io.Reader 读取
└── 核心接口
    ├── Stringer    (String() string)
    ├── GoStringer  (GoString() string)
    ├── Formatter   (Format(State, rune))
    └── Scanner     (Scan(ScanState, rune) error)
```

## 输出函数家族

`fmt` 包提供了三组输出函数，每组都有 `f`（格式化）、无后缀（默认）和 `ln`（换行）三个变体：

### Printf / Sprintf / Fprintf

这三个函数接受格式化字符串，是最灵活的输出方式：

```go
// Printf 输出到标准输出
n, err := fmt.Printf("name: %s, age: %d\n", "Forest", 25)

// Sprintf 返回格式化后的字符串
s := fmt.Sprintf("name: %s, age: %d", "Forest", 25)

// Fprintf 写入任意 io.Writer
n, err := fmt.Fprintf(os.Stderr, "error: %v\n", err)
```

从源码看它们的关系（`print.go`）：

```go
func Printf(format string, a ...any) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}

func Sprintf(format string, a ...any) string {
    p := newPrinter()
    p.doPrintf(format, a)
    s := string(p.buf)
    p.free()
    return s
}
```

`Printf` 只是 `Fprintf` 的包装，传入 `os.Stdout`。`Sprintf` 则将结果写入了内部的 buffer 后转为字符串返回。

### Print / Sprint / Fprint

不带格式化，每个参数默认使用 `%v` 输出：

```go
fmt.Print("hello", " ", "world")        // 输出: hello world
fmt.Print(42, " ", true)                // 输出: 42 true
s := fmt.Sprint("value: ", 42)          // s = "value: 42"
```

注意：当两个相邻参数都不是字符串时，`Print` 会在它们之间自动插入空格。从源码（`print.go:1200`）可以看到这个逻辑：

```go
func (p *pp) doPrint(a []any) {
    prevString := false
    for argNum, arg := range a {
        isString := arg != nil && reflect.TypeOf(arg).Kind() == reflect.String
        if argNum > 0 && !isString && !prevString {
            p.buf.writeByte(' ')
        }
        p.printArg(arg, 'v')
        prevString = isString
    }
}
```

### Println / Sprintln / Fprintln

与 `Print` 类似，但始终在参数之间插入空格，并在末尾追加换行符：

```go
fmt.Println("hello", "world")    // 输出: hello world\n
fmt.Println(1, 2, 3)             // 输出: 1 2 3\n
```

### Append / Appendf / Appendln（Go 1.19+）

追加到已有的 `[]byte` 切片，在需要高性能拼接场景下很有用：

```go
buf := make([]byte, 0, 64)
buf = fmt.Appendf(buf, "name: %s, age: %d", "Forest", 25)
// buf = []byte("name: Forest, age: 25")
```

## 格式化动词

格式化动词以 `%` 开头，决定了值以何种形式输出。以下按类型分类。

### 通用动词

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%v` | 默认格式 | `fmt.Sprintf("%v", 42)` → `"42"` |
| `%+v` | 添加字段名（struct） | `fmt.Sprintf("%+v", person)` → `{Name:Forest Age:25}` |
| `%#v` | Go 语法表示 | `fmt.Sprintf("%#v", 42)` → `"42"` |
| `%T` | 类型名 | `fmt.Sprintf("%T", 42)` → `"int"` |
| `%%` | 字面量 `%` | `fmt.Sprintf("100%%")` → `"100%"` |

`%#v` 对于不同类型有不同的表现：

```go
fmt.Sprintf("%#v", []int{1, 2, 3})  // "[]int{1, 2, 3}"
fmt.Sprintf("%#v", map[string]int{"a": 1})
// "map[string]int{"a":1}"
```

### 布尔值

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%t` | 输出 `true` 或 `false` | `fmt.Sprintf("%t", true)` → `"true"` |

### 整数

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%b` | 二进制 | `fmt.Sprintf("%b", 10)` → `"1010"` |
| `%c` | 对应 Unicode 字符 | `fmt.Sprintf("%c", 65)` → `"A"` |
| `%d` | 十进制 | `fmt.Sprintf("%d", 42)` → `"42"` |
| `%o` | 八进制 | `fmt.Sprintf("%o", 8)` → `"10"` |
| `%O` | 八进制带 `0o` 前缀 | `fmt.Sprintf("%O", 8)` → `"0o10"` |
| `%q` | 单引号字符字面量 | `fmt.Sprintf("%q", 65)` → `"'A'"` |
| `%x` | 十六进制（小写） | `fmt.Sprintf("%x", 255)` → `"ff"` |
| `%X` | 十六进制（大写） | `fmt.Sprintf("%X", 255)` → `"FF"` |
| `%U` | Unicode 格式 | `fmt.Sprintf("%U", '林')` → `"U+6797"` |

源码中整数格式化的核心逻辑在 `format.go` 的 `fmtInteger` 方法中，它通过从右到左的方式填充数字到 buffer：

```go
// format.go:197 简化逻辑
func (f *fmt) fmtInteger(u uint64, base int, isSigned bool, verb rune, digits string) {
    // ...省略负数和精度处理...
    i := len(buf)
    switch base {
    case 10:
        for u >= 10 {
            i--
            next := u / 10
            buf[i] = byte('0' + u - next*10)
            u = next
        }
    case 16:
        for u >= 16 {
            i--
            buf[i] = digits[u&0xF]
            u >>= 4
        }
    // case 8, 2 ...
    }
}
```

### 浮点数与复数

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%f` | 小数点，无指数 | `fmt.Sprintf("%f", 3.14)` → `"3.140000"` |
| `%F` | `%f` 的同义词 | 同上 |
| `%e` | 科学计数法（小写 e） | `fmt.Sprintf("%e", 3.14)` → `"3.140000e+00"` |
| `%E` | 科学计数法（大写 E） | `fmt.Sprintf("%E", 3.14)` → `"3.140000E+00"` |
| `%g` | 自动选择 `%e` 或 `%f` | `fmt.Sprintf("%g", 3.14)` → `"3.14"` |
| `%G` | 自动选择 `%E` 或 `%F` | `fmt.Sprintf("%G", 3.14)` → `"3.14"` |
| `%b` | 无小数科学计数法 | `fmt.Sprintf("%b", 3.14)` → `"1417889010936558p-48"` |
| `%x` | 十六进制浮点 | `fmt.Sprintf("%x", 3.14)` → `"0x1.91eb851eb851fp+01"` |

复数以 `(real+imagi)` 格式输出：

```go
fmt.Sprintf("%f", 1+2i) // "(1.000000+2.000000i)"
```

从源码（`print.go:470`）可以看到，复数的格式化是对实部和虚部分别调用 `fmtFloat`：

```go
func (p *pp) fmtComplex(v complex128, size int, verb rune) {
    p.buf.writeByte('(')
    p.fmtFloat(real(v), size/2, verb)
    p.fmt.plus = true  // 虚部总是带符号
    p.fmtFloat(imag(v), size/2, verb)
    p.buf.writeString("i)")
}
```

### 字符串与字节切片

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%s` | 原始字符串 | `fmt.Sprintf("%s", "hello")` → `"hello"` |
| `%q` | 双引号转义字符串 | `fmt.Sprintf("%q", "hello")` → `"\"hello\""` |
| `%x` | 十六进制（小写） | `fmt.Sprintf("%x", "abc")` → `"616263"` |
| `%X` | 十六进制（大写） | `fmt.Sprintf("%X", "abc")` → `"616263"` |

### 指针

| 动词 | 说明 | 示例 |
| ---- | ---- | ---- |
| `%p` | 十六进制地址，带 `0x` | `fmt.Sprintf("%p", &x)` → `"0xc0000b2008"` |

`%p` 也可以和 `%b`、`%d`、`%o`、`%x`、`%X` 配合使用，将指针值当作整数格式化。

## 宽度与精度

宽度（width）和精度（precision）在动词前面指定，用 `.` 分隔：

```text
%[flags][width][.precision]verb
```

示例：

```go
fmt.Sprintf("%f", 3.14159)      // "3.141590"    默认宽度，默认精度
fmt.Sprintf("%9f", 3.14159)     // " 3.141590"   宽度 9，默认精度
fmt.Sprintf("%.2f", 3.14159)    // "3.14"        默认宽度，精度 2
fmt.Sprintf("%9.2f", 3.14159)   // "    3.14"    宽度 9，精度 2
fmt.Sprintf("%9.f", 3.14159)    // "        3"   宽度 9，精度 0
```

宽度和精度都可以用 `*` 代替，从参数中获取：

```go
fmt.Sprintf("%*.*f", 9, 2, 3.14159) // 等价于 "%9.2f" → "    3.14"
```

宽度和精度的行为取决于数据类型：

| 类型 | 宽度含义 | 精度含义 |
| ---- | ---- | ---- |
| 整数 | 最小输出宽度（右对齐，左补空格） | 最少数字位数（前导补零） |
| 浮点数 | 最小输出宽度 | 小数位数（`%g/%G` 为有效数字位数） |
| 字符串 | 最小输出宽度 | 最大截取长度（按 rune 计数） |

## 标志（Flags）

| 标志 | 说明 |
| ---- | ---- |
| `+` | 总是输出正负号；对 `%q` 保证纯 ASCII 输出 |
| `-` | 左对齐（右侧填充空格） |
| `#` | 备用格式：`%#b` 加 `0b`，`%#o` 加 `0`，`%#x` 加 `0x`；`%#v` 输出 Go 语法；`%#U` 输出字符名 |
| ` `（空格） | 正数前留空格；`% x` 十六进制字节间加空格 |
| `0` | 用 `0` 填充而非空格 |

```go
fmt.Sprintf("%+d", 42)      // "+42"
fmt.Sprintf("% d", 42)      // " 42"
fmt.Sprintf("%#x", 255)     // "0xff"
fmt.Sprintf("%#b", 10)      // "0b1010"
fmt.Sprintf("%-10d|", 42)   // "42        |"
fmt.Sprintf("%010d", 42)    // "0000000042"
```

## 显式参数索引

使用 `[n]` 语法可以显式指定第 n 个参数（从 1 开始计数）：

```go
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)        // "22 11"
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)   // " 12.00"  等价于 "%6.2f"
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)     // "16 17 0x10 0x11"
```

这在需要重复使用同一个参数或国际化（调整参数顺序）时非常有用。

## 核心接口

`fmt` 包定义了四个核心接口，让自定义类型可以控制自己的格式化行为：

### Stringer

```go
type Stringer interface {
    String() string
}
```

最常用的接口。当一个值实现了 `String()` 方法，`fmt` 在使用 `%v`、`%s` 等动词时会自动调用它：

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Forest", Age: 25}
    fmt.Println(p)       // Forest(25)
    fmt.Printf("%v\n", p) // Forest(25)
}
```

从源码（`print.go:656`）可以看到 `Stringer` 的调用时机：当格式化动词是 `v`、`s`、`x`、`X`、`q` 时，会检查参数是否实现了 `error` 或 `Stringer` 接口：

```go
switch verb {
case 'v', 's', 'x', 'X', 'q':
    switch v := p.arg.(type) {
    case error:
        handled = true
        defer p.catchPanic(p.arg, verb, "Error")
        p.fmtString(v.Error(), verb)
        return
    case Stringer:
        handled = true
        defer p.catchPanic(p.arg, verb, "String")
        p.fmtString(v.String(), verb)
        return
    }
}
```

**注意优先级**：`error` 接口优先于 `Stringer` 接口。如果一个类型同时实现了 `Error()` 和 `String()`，在使用 `%v` 时会调用 `Error()`。

**避免无限递归**：在 `String()` 方法内部调用 `fmt.Sprintf` 时，不要把接收者自身当作参数传给 `%s` 等动词，否则会无限递归：

```go
// 错误：会导致无限递归
type X string
func (x X) String() string { return fmt.Sprintf("<%s>", x) }

// 正确：先转换为 string
type X string
func (x X) String() string { return fmt.Sprintf("<%s>", string(x)) }
```

### GoStringer

```go
type GoStringer interface {
    GoString() string
}
```

在使用 `%#v` 时被调用，返回 Go 语法表示：

```go
func (p Person) GoString() string {
    return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, p.Age)
}

fmt.Printf("%#v\n", Person{"Forest", 25})
// Person{Name: "Forest", Age: 25}
```

### Formatter

```go
type Formatter interface {
    Format(f State, verb rune)
}
```

最灵活的接口。实现了 `Formatter` 的类型可以完全自定义格式化行为，且优先级最高（高于 `Stringer` 和 `GoStringer`）。

`State` 接口提供了对格式化状态（宽度、精度、标志）的访问：

```go
type State interface {
    Write(b []byte) (n int, err error)
    Width() (wid int, ok bool)
    Precision() (prec int, ok bool)
    Flag(c int) bool
}
```

示例：

```go
type Hex int

func (h Hex) Format(f fmt.State, verb rune) {
    switch verb {
    case 'x', 's':
        fmt.Fprintf(f, "%x", int(h))
    case 'X':
        fmt.Fprintf(f, "%X", int(h))
    case 'd':
        fmt.Fprintf(f, "%d", int(h))
    default:
        fmt.Fprintf(f, "Hex(%d)", int(h))
    }
}

func main() {
    h := Hex(255)
    fmt.Printf("%x\n", h)  // ff
    fmt.Printf("%X\n", h)  // FF
    fmt.Printf("%d\n", h)  // 255
    fmt.Printf("%v\n", h)  // Hex(255)
}
```

`FormatString` 函数（Go 1.21+）可以在 `Format` 方法中重建原始格式化指令：

```go
func (h Hex) Format(f fmt.State, verb rune) {
    directive := fmt.FormatString(f, verb) // 例如 "%08x"
    fmt.Fprintf(f, "directive=%s value=%x", directive, int(h))
}
```

### 接口优先级总结

当多个接口同时实现时，`fmt` 按以下顺序处理（`print.go:621`）：

1. `Formatter` 接口 — 最高优先级
2. `%#v` + `GoStringer` 接口
3. `error` 接口（当动词接受字符串时）
4. `Stringer` 接口（当动词接受字符串时）

## Errorf 与错误包装

`Errorf` 用于创建格式化的错误消息，配合 `%w` 动词可以实现错误包装：

```go
var ErrNotFound = errors.New("not found")

func getUser(id int) error {
    return fmt.Errorf("get user %d: %w", id, ErrNotFound)
}

func main() {
    err := getUser(42)
    fmt.Println(errors.Is(err, ErrNotFound)) // true

    var target *ErrNotFound
    fmt.Println(errors.As(err, &target))     // true
}
```

从源码（`errors.go`）可以看到，`%w` 会触发 `Unwrap` 机制的实现：

- 单个 `%w`：返回 `*wrapError`，实现 `Unwrap() error`
- 多个 `%w`：返回 `*wrapErrors`，实现 `Unwrap() []error`

```go
// errors.go 源码
func Errorf(format string, a ...any) error {
    p := newPrinter()
    p.wrapErrs = true
    p.doPrintf(format, a)
    s := string(p.buf)
    switch len(p.wrappedErrs) {
    case 0:
        err = errors.New(s)
    case 1:
        err = &wrapError{msg: s, err: a[p.wrappedErrs[0]].(error)}
    default:
        // 多个 %w 返回 wrapErrors
        err = &wrapErrors{s, errs}
    }
}
```

## Scanning（输入扫描）

`fmt` 包的扫描函数是输出函数的逆操作，从输入中读取数据并存储到变量中：

### Scan / Scanln / Scanf

```go
// Scan: 从标准输入读取，换行视为空格
var name string
var age int
fmt.Scan(&name, &age)

// Scanln: 从标准输入读取，遇到换行停止
fmt.Scanln(&name, &age)

// Scanf: 按格式字符串读取
fmt.Scanf("name: %s age: %d", &name, &age)
```

### Sscan / Sscanln / Sscanf

从字符串读取，常用于解析固定格式的内容：

```go
var s string
var i int
fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
// s = "12345", i = 67
```

### 三组 Scan 的区别

| 函数 | 换行处理 | 特点 |
| ---- | ---- | ---- |
| `Scan` / `Fscan` / `Sscan` | 换行视为空格 | 持续读取直到所有参数填满 |
| `Scanln` / `Fscanln` / `Sscanln` | 换行停止扫描 | 最后一个参数后必须跟换行或 EOF |
| `Scanf` / `Fscanf` / `Sscanf` | 换行必须匹配格式 | 按格式字符串精确解析 |

### Scanner 接口

```go
type Scanner interface {
    Scan(state ScanState, verb rune) error
}
```

自定义类型可以实现此接口来控制扫描行为。`ScanState` 提供了底层读取能力：

```go
type ScanState interface {
    ReadRune() (r rune, size int, err error)
    UnreadRune() error
    SkipSpace()
    Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
    Width() (wid int, ok bool)
    Read(buf []byte) (n int, err error)
}
```

## 内部实现原理

### 对象池优化

`fmt` 包使用 `sync.Pool` 来复用 `pp`（printer）对象，避免每次格式化都分配新内存：

```go
// print.go:146
var ppFree = sync.Pool{
    New: func() any { return new(pp) },
}

func newPrinter() *pp {
    p := ppFree.Get().(*pp)
    p.panicking = false
    p.erroring = false
    p.wrapErrs = false
    p.fmt.init(&p.buf)
    return p
}
```

`pp` 结构体是格式化的核心，它持有格式化过程中的所有状态：

```go
// print.go:120
type pp struct {
    buf        buffer         // 输出缓冲区
    arg        any            // 当前参数
    value      reflect.Value  // reflect.Value 参数
    fmt        fmt            // 底层格式化器
    reordered  bool           // 是否使用了显式参数索引
    goodArgNum bool           // 参数索引是否有效
    panicking  bool           // 防止 panic 无限递归
    erroring   bool           // 防止 Error/String 方法循环调用
    wrapErrs   bool           // 是否启用 %w 错误包装
    wrappedErrs []int         // %w 对应的参数位置
}
```

当 buffer 超过 64KB 时，free 方法会丢弃 buffer 而不是放回池中，防止内存膨胀：

```go
func (p *pp) free() {
    if cap(p.buf) > 64*1024 {
        p.buf = nil
    } else {
        p.buf = p.buf[:0]
    }
    // ...
    ppFree.Put(p)
}
```

### doPrintf 解析流程

`doPrintf` 是格式化输出的核心方法，它的解析流程如下：

```text
1. 遍历格式字符串，查找 % 字符
2. 解析标志（+、-、#、空格、0）
3. 解析显式参数索引 [n]
4. 解析宽度（数字或 *）
5. 解析精度（.数字或 .*）
6. 获取动词（verb）
7. 调用 printArg 进行实际格式化
8. 检查多余参数
```

对于简单格式（小写字母动词，无宽度和精度），源码中有快速路径优化：

```go
// print.go:1060 快速路径
if 'a' <= c && c <= 'z' && argNum < len(a) {
    switch c {
    case 'w':
        p.wrappedErrs = append(p.wrappedErrs, argNum)
        fallthrough
    case 'v':
        p.fmt.sharpV = p.fmt.sharp
        p.fmt.sharp = false
        p.fmt.plusV = p.fmt.plus
        p.fmt.plus = false
    }
    p.printArg(a[argNum], rune(c))
    argNum++
    i++
    continue formatLoop
}
```

### printArg 类型分发

`printArg` 方法（`print.go:681`）根据参数的实际类型选择合适的格式化路径：

```text
printArg(arg, verb)
├── nil → "<nil>"
├── %T → reflect.TypeOf(arg).String()
├── %p → fmtPointer()
├── bool → fmtBool()
├── int/int8/.../int64 → fmtInteger()
├── uint/uint8/.../uint64 → fmtInteger()
├── float32/float64 → fmtFloat()
├── complex64/complex128 → fmtComplex()
├── string → fmtString()
├── []byte → fmtBytes()
├── reflect.Value → 特殊处理
└── default → handleMethods() → printValue()
```

对于复合类型（struct、map、slice 等），`printValue` 通过反射递归处理每个元素。

## 格式化错误提示

当格式化动词与参数类型不匹配时，`fmt` 会输出描述性的错误信息：

```go
fmt.Sprintf("%d", "hi")     // %!d(string=hi)
fmt.Sprintf("hi%d")         // hi%!d(MISSING)
fmt.Sprintf("hi", "guys")   // hi%!(EXTRA string=guys)
fmt.Sprintf("%*s", 4.5, "hi") // %!(BADWIDTH)hi
```

所有错误都以 `%!` 开头，便于在输出中快速识别。

## 实战技巧

### 调试利器 `%+v` 和 `%#v`

```go
type Config struct {
    Host string
    Port int
    TLS  bool
}

c := Config{Host: "localhost", Port: 8080, TLS: true}

fmt.Printf("%v\n", c)   // {localhost 8080 true}
fmt.Printf("%+v\n", c)  // {Host:localhost Port:8080 TLS:true}
fmt.Printf("%#v\n", c)  // main.Config{Host:"localhost", Port:8080, TLS:true}
```

### 格式化 JSON

```go
data := map[string]any{"name": "Forest", "age": 25}
fmt.Sprintf("%#v", data)
// map[string]interface {}{"age":25, "name":"Forest"}
```

### 自定义类型的简洁输出

```go
type Money int64 // 单位：分

func (m Money) String() string {
    return fmt.Sprintf("¥%.2f", float64(m)/100)
}

fmt.Println(Money(999)) // ¥9.99
```

### 错误链的最佳实践

```go
func LoadConfig(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("load config %s: %w", path, err)
    }
    // ...
    return nil
}
```

使用 `%w` 包装错误，使得调用方可以用 `errors.Is` 和 `errors.As` 检查错误链。

## 小结

`fmt` 包虽小但五脏俱全：

- **输出**：`Printf` 系列通过格式化动词和标志实现灵活输出
- **输入**：`Scanf` 系列是输出函数的逆操作
- **接口**：`Stringer`、`GoStringer`、`Formatter` 让自定义类型控制格式化行为
- **错误**：`Errorf` + `%w` 是 Go 错误链的基石
- **性能**：`sync.Pool` 复用 printer 对象，快速路径优化常见格式

掌握 `fmt` 包是 Go 开发的基本功，理解其内部机制有助于写出更优雅的代码。
