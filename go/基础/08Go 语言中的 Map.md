`map` 类型是一个无序的键值对集合；由键（`key`）和值（`value`）组成；在一个 `map` 中，键是唯一的，在集合中不能有两个相同的键。

## 声明语法

```go
var name map[key_type]value_type
```

- name 为 `map` 的变量名
- key_type 为键的类型
- value_type 为键对应的值类型

> Tips:
>
> - `key` 与 `value` 的类型可以相同，也可以不相同
> - `key` 的类型必须支持 `==` 和 `!=` 两种比较操作；如果没有办法比较 `map` 中的`key` 是否相同，那么这些 `key` 就不能作为 `map` 的 `key`
> - 函数类型、`map` 类型和切片只支持与 `nil` 的比较，而**不支持同类型两个变量的比较**，否则会编译报错
> - 如果没有显示地赋予 `map` 变量初始值，则 `map` 类型变量的默认值为 `nil`
> - `map` 类型的变量必须显示初始化后才能使用，否则会导致程序进程异常退出

## 显示赋值 `map` 类型的两种方式

### 复合字面值

```go
package main

import "fmt"

func main() {
	m := map[int]string{} // 或者 var m map[int]string = map[int]string{}
	m[1] = "go programming"
	fmt.Printf("m[1]: %s", m[1]) // m[1]: go programming
}
```

显示初始化 `map` 类型变量 m，虽然 m 中没有任何键值对，但变量 m 也不等同于初值 `nil` 的 `map` 的变量，此时对其操作不会引起程序异常。

> 在Go语言中，`var books map[string]interface{}` 和 `var books map[string]interface{}{}`的区别是：
>
> - `var books map[string]interface{}`声明了一个名为 books 的 `map` 变量，但没有为其分配内存空间。这意味着books 是一个空的 `map`，无法进行任何操作。
> - `var books map[string]interface{}{}`声明了一个名为 books 的 `map` 变量，并为其分配了内存空间。这意味着books 是一个空的 `map`，但可以进行操作，如添加、删除和检索键值对。

### 使用 `make` 内置函数

分为不显示指定容量和显示指定初始容量两种；`map` 类型的容量不会受限于它的初始容量，当其中的键值对数量超过初始容量后，Go 运行时会自动增加 `map` 类型的容量，保证后续键值对的正常插入。

```go
package main

import "fmt"

func main() {
	var books = make(map[string]float64)
	var notes = make(map[string]float64, 8) // 初始容量为8
	books["en"] = 23.34
	notes["en"] = 3.0
	fmt.Printf("books en: %f, notes en: %f\n", books["en"], notes["en"]) // books en: 23.340000, notes en: 3.000000
}
```

### `map` 变量的传递开销

`map` 是引用类型，也就是 `map` 作为参数被传递给函数或方法的时候,实质上传递的只是一个"描述符"，不是值拷贝，所以开销是固定的，函数内部对 `map` 类型参数的修改在函数外部也是可见的。

```go
package main

import "fmt"

func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(m) // map[key1:1 key2:2]
	foo(m)
	fmt.Println(m) // map[key1:11 key2:12]
}
```

在 main 函数中定义了一个 m 的 map，传递给 foo 函数进行修改值，将 11 赋值给key1，将 12 赋值给 key2，然后再次打印 main 函数中的 m 时，m 中的 key1 和 key2 已经被修改。

## `map` 的基本操作

### 插入新的键值对

> 如果插入的 key 不存在则新建一个 key 并赋值，如果存在，则替换。

```go
package main

import "fmt"

func main() {
	m := map[int]string{}
	m[1] = "go programming"
	fmt.Printf("m[1]: %s\n", m[1]) // m[1]: go programming
	m[1] = "replace origion value"
	fmt.Printf("替换后的值：%s\n", m[1]) // 替换后的值：replace origion value
}
```

### 获取键值对数量

> 通过内置函数 `len` 获取当前遍历那个已经存储的键值对数量。**不能对 `map` 变量调用 `cap` 来获取当前容量**。

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(len(m)) // 2
	m["key3"] = 3
	fmt.Println(len(m)) // 3
}
```

### 查找和数据读取

判断某个 key 是否存在；使用 “comma ok” 惯用法对 `map` 进行键查找和键值读取操作。

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)
	v, ok := m["key1"]
	if !ok {
		// "key1"不在map中
	}

	// "key1"在map中，v将被赋予"key1"键对应的value
	fmt.Printf("v: %d", v) // v: 0 int 类型零值
}
```

### 删除数据

> 使用内置函数 `delete` 来删除 `map` 中的数据；`delete` 函数是从 `map`中删除键的唯一方法；即使传给 `delete` 的键在 `map` 中并不存在，`delete` 函数的执行也不会失败，也不会抛出运行时的异常。

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"key1": 13,
		"key2": 22,
	}

	fmt.Println(m) // map[key1:13 key2:22]
	// 第一个参数是 map 类型变量，第二个参数就是要删除的键
	delete(m, "key2") // 删除"key2"
	fmt.Println(m)    // map[key1:13]
}
```

### 遍历 `map` 中的键值数据

```go
package main

import "fmt"

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}
```

多次遍历

```go
package main

import "fmt"

func iteration(m map[int]int) {
	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	for i := 0; i < 3; i++ {
		iteration(m)
	}
}

// 输出结果
// { [1, 11] [2, 12] [3, 13] }
// { [1, 11] [2, 12] [3, 13] }
// { [3, 13] [1, 11] [2, 12] }
```

> 对同一个`map`做多次遍历的时候,每次遍历元素的次序都不相同。

Tips:

- 不要依赖 `map` 的元素遍历顺序。
- `map` 不是线程安全的，不支持并发读写。
- 不要尝试获取 `map` 中元素（value）的地址。

### 实操：实现一个对 `map` 进行稳定次序遍历的方法

> 思路：用一个有序结构存储 key， 如 `slice`，然后遍历这个`slice`，用 key 获取值。

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]string = map[int]string{
		1: "value-1",
		2: "value-2",
		3: "value-3",
	}

	for i := 0; i < 3; i++ {
		itrater(m)
	}
}

func itrater(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("{ ")
	for _, key := range keys {
		fmt.Printf("[%d, %s] ", key, m[key])
	}
	fmt.Printf("}\n")
}

// 输出结果
// { [1, value-1] [2, value-2] [3, value-3] }
// { [1, value-1] [2, value-2] [3, value-3] }
// { [1, value-1] [2, value-2] [3, value-3] }
```
