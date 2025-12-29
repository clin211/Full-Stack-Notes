# 结构体

## 初识

结构体是**由一系列具有相同类型或不同类型**的数据构成的数据集合。结构体是由**0个或多个任意类型的值聚合成的实体**，每个值都可以被称为结构体的成员。

### 特性

结构体的成员也可以被称为"字段"，具有以下特性：

- 字段**拥有自己的类型和值**。
- 字段**名必须唯一**。
- 字段的**类型也可以是结构体**；甚至是**字段所在结构体的类型的指针类型**。
- 字段的**首字母决定其可访问性**。

### 自定义类型

#### 用 `type` 关键字

```go
type T S // 定义一个新类型 T
```

示例:[传送门](https://go.dev/play/p/1TB_x-R6N2J)

```go
package main

import (
	"fmt"
)

// 新类型 T1 是基于 Go 原生类型 int 定义的新自定义类型
type T1 int

// 新类型 T2 则是基于上面定义的类型 T1，定义的新类型
type T2 T1

func main() {
	var t1 T1 = 42
	var t2 T2 = 10

	fmt.Printf("Type of t1: %T\n", t1) // Type of t1: main.T1
	fmt.Printf("Type of t2: %T\n", t2) // Type of t2: main.T2

	t2 = T2(33)
	fmt.Printf("t2: %d\n", t2) // T2: 33
}
```

如果一个新类型是基于某个 Go 原生类型定义的，那么我们就叫 Go 原生类型为新类型的**底层类型**（Underlying Type)

#### 类型别名

```go
type T = S // type alias
```

实际上，T 和 S 是同一种类型。

示例：[传送门](https://go.dev/play/p/Jvb_tsUoREp)

```go
package main

import (
	"fmt"
)

func main() {
	type T = string

	var s string = "hello"
	var t T = s
	fmt.Printf("%T\n", t) // string
}
```

#### 自定义类型和类型别名的区别

1、自定义类型：使用 `type` 定义一个新的类型，它*与现有的类型没有直接关系*，即使底层类型相同；这意味着**自定义类型不能直接赋值给其他类型，也不能比较**；自定义类型可以。[传送门](https://go.dev/play/p/Qlcpf1hzn__Q)

```go
package main

import "fmt"

type MyInt int

func main() {
	var num MyInt = 42
	var i int = 10

	// 自定义类型不能直接赋值给其底层类型
	i = num // 编译报错：cannot use num (variable of type MyInt) as int value in assignment  解决方案: 显示类型转换: i = int(num)

	// 自定义类型不能直接比较
	fmt.Println(num == i) // 编译报错：invalid operation: num == i (mismatched types MyInt and int)  解决方案: 同上，int(num) == i
}
```

2、类型别名：使用 `type` 定义一个类型别名，它*与现有的类型是完全相同的*，**类型别名可以直接赋值给其底层类型**，也可以直接比较。[传送门](https://go.dev/play/p/KCLb4gqydxR)

```go
package main

import "fmt"

type MyIntAlias = int

func main() {
	var num MyIntAlias = 42
	var i int = 10

	// 类型别名可以直接赋值给其底层类型
	i = num

	// 类型别名可以直接比较
	fmt.Printf("比较结果：%t\n", num == i) // 比较结果：true
}
```

## 定义一个结构体类型

### 声明形式

一个名为 `T` 的结构体类型，定义中 `struct` 关键字后面的大括号包裹的内容就是一个类型字面值。

```go
type T struct {
	Field1 T1
	Field2 T2
	... ...
	FieldN Tn
}
```

> 说明：
>
> - `T` 表示自定义结构体的名称，在同一包内不能包含重复的类型名。
> - `struct{}` 表示结构体类型；`type T struct{}` 可以理解为将 struct 结构体定义为类型名的类型。

### 定义一个空结构体

```go
type Empty struct{} // Empty是一个不包含任何字段的空结构体类型

var s Empty
println(unsafe.Sizeof(s)) // 0
```

输出的空结构体类型变量的大小为 0，也就是说，空结构体类型变量的内存占用为 0。基于空结构体类型内存零开销这样的特性，我们在日常 Go 开发中会经常使用空结构体类型元素，作为一种“事件”信息进行 Goroutine 之间的通信。

### 使用其他结构体作为自定义结构体中字段的类型

示例代码:[传送门](https://go.dev/play/p/s5bSFZ0ry1Y)

```go
package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	address := Address{
		City:    "New York",
		Country: "USA",
	}

	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: address,
	}

	fmt.Printf("person name: %s\n", person.Name)                       // person name: John Doe
	fmt.Printf("person age: %d\n", person.Age)                         // person age: 30
	fmt.Printf("person city: %s\n", person.Address.City)               // person city: New York
	fmt.Printf("person address country: %s\n", person.Address.Country) // person address country: USA
}
```

## 结构体变量的声明与初始化

### 零值可用

在Go语言中，当声明一个变量但没有显式赋值时，这个变量会被初始化为其对应类型的零值。_对于结构体类型来说，使用结构体的零值作为初始值意味着将所有字段都设置为其类型的零值_。

下面是一个示例代码来说明结构体的零值初始化：[传送门](https://go.dev/play/p/GJDr97b2oUu)

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person

	fmt.Printf("Name: %s\n", p.Name) // Name: ---> string 类型的零值空 字符串
	fmt.Printf("Age: %d\n", p.Age)   // Age: 0 ---> int 类型的零值 0
}
```

在上面的示例中，定义了一个 `Person` 结构体，它包含了 `Name` 和 `Age` 两个字段。然后，在 `main` 函数中声明并初始化一个 `Person` 类型的变量 `p`，但没有显式赋值。

由于没有给 `p` 赋值，所以它会被初始化为 `Person` 类型的零值。**对于`string`类型的`Name`字段，其零值是空字符串`""`；对于`int`类型的`Age`字段，其零值是`0`**。

因此，打印 `p.Name` 时，输出为空字符串；打印 `p.Age` 时，输出为`0`。这就是结构体的零值初始化，即使用结构体的零值作为其初始值。

### 使用复合字面值

1、对结构体变量进行显示初始化，按顺序依次给每个结构体字段进行赋值（不推荐）。[传送门](https://go.dev/play/p/7v1-N8XQFRG)

```go
package main

import "fmt"

type T struct {
	F1 int
	F2 string
	f3 int
	F4 int
	F5 int
}

func main() {

	var t = T{11, "hello", 13, 14, 15}
	fmt.Println(t) // {11 hello 13 14 15}
}
```

> 这种方式有一些潜在的弊端：
>
> 1. **依赖字段顺序**：使用显示初始化时，需要按照结构体定义中字段的顺序给字段赋值。如果结构体的字段很多，或者字段顺序发生变化，就需要小心确保赋值的正确顺序。
> 2. **可读性下降**：当结构体具有大量字段时，逐个指定字段值可能会导致代码变得冗长和难以阅读。

2、_Go 推荐我们用 "field:value" 形式的符合字面值，对结构体类型进行显示初始化_。[传送门](https://go.dev/play/p/bgPoUMoawsp)

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Printf("Name: %s\n", p.Name) // Name: Alice
	fmt.Printf("Age: %d\n", p.Age)   // Age: 25
}
```

3、使用特定的构造函数

可以使用特定的构造函数来声明和初始化结构体。构造函数是一个普通的函数，用于创建并初始化结构体对象，并返回对象的指针或引用。

下面是一个示例代码，展示了如何使用构造函数声明和初始化结构体：[传送门](https://go.dev/play/p/D4U8TlnPDLj)

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	p := NewPerson("Alice", 25)

	fmt.Printf("Name: %s\n", p.Name) // Name: Alice
	fmt.Printf("Age: %d\n", p.Age)   // Age: 25
}
```

在上面的示例中，定义了一个名为 `Person` 的结构体类型，包含了 `Name` 和 `Age` 两个字段。然后，又定义了一个名为 `NewPerson` 的构造函数，该函数接收 `name` 和 `age` 作为参数，并返回一个指向 `Person` 结构体的指针。

在`main`函数中，使用 `NewPerson` 构造函数来创建并初始化一个 `Person` 对象，将其赋值给变量 `p`。通过构造函数，我们可以在创建结构体对象的同时对其字段进行初始化。

使用这种方式也有其利弊；如下：

> 使用构造函数的方式有以下优点：
>
> 1. **封装复杂逻辑**：构造函数可以封装一些复杂的逻辑，例如对字段进行验证或设置默认值等，以确保结构体对象的合法性。
> 2. **提供灵活初始化选项**：构造函数可以接收不同的参数组合，以提供灵活的初始化选项，使得创建结构体对象更加方便。
> 3. **隐藏实现细节**：通过使用构造函数，可以隐藏结构体的内部实现细节，使得代码更具模块化和封装性。
>
> 然而，使用构造函数的方式也存在一些潜在的弊端：
>
> 1. **需要显式调用构造函数**：相对于使用字面值初始化结构体对象的方式，使用构造函数需要显式调用函数来创建对象，增加了一些额外的代码。
> 2. **需要定义额外的构造函数**：如果结构体有多个初始化逻辑或需要提供不同的初始化选项，可能需要定义多个构造函数，增加了一些代码复杂性。

_如果结构体的构造函数需要接收变长参数，可以使用切片或可变参数的方式来实现_。下面是一个示例代码，展示了如何使用构造函数的方式声明结构体并初始化，并接收变长参数：[传送门](https://go.dev/play/p/xHCHP2RHZm2)

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Tags []string
}

func NewPerson(name string, age int, tags ...string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		Tags: tags,
	}
}

func main() {
	p1 := NewPerson("Alice", 25)
	p2 := NewPerson("Bob", 30, "tag1", "tag2")

	fmt.Printf("Name: %s\n", p1.Name) // Name: Alice
	fmt.Printf("Age: %d\n", p1.Age)   // Age: 25
	fmt.Printf("Tags: %v\n", p1.Tags) // Tags: []

	fmt.Printf("Name: %s\n", p2.Name) // Name: Bob
	fmt.Printf("Age: %d\n", p2.Age)   // Age: 30
	fmt.Printf("Tags: %v\n", p2.Tags) // Tags: [tag1 tag2]
}
```

在上面的示例中，将 `NewPerson` 构造函数修改为接收变长参数 `tags ...string` 。在构造函数内部，将传入的 `name` 和 `age`参数赋值给结构体的对应字段，而 `tags` 参数则被视为一个切片，将其作为 `Person` 结构体的 `Tags` 字段进行初始化。

使用构造函数声明结构体并初始化时，可以根据需要传入不同数量的变长参数。在示例中，分别调用了两次 `NewPerson`构造函数，第一次没有传入任何 `tags` 参数，第二次传入了两个 `tags` 参数。

这种方式的好处是可以根据需求传入不同数量的参数，并将它们作为切片或可变参数进行处理。它提供了更大的灵活性，使得结构体初始化更加方便。

需要注意的是，在**使用变长参数的构造函数时，需要根据实际情况进行参数传递，确保传递的参数类型和顺序与构造函数的定义一致**。

## 进阶

在 Go 语言中，没有像 Java、C# 等面向对象编程语言那样的显式继承机制，但可以使用组合和接口来实现类似的效果。

### 组合

通过在一个结构体中嵌入另一个结构体，可以实现组合关系。嵌入的结构体可以访问被嵌入结构体的字段和方法，从而实现了一种类似继承的效果。

下面是一个示例代码，展示了如何使用组合来实现类似继承的效果：[传送门](https://go.dev/play/p/7ATPOeq3xGL)

```go
package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("Animal speaks")
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}

	fmt.Println(d.Name)  // Buddy
	fmt.Println(d.Breed) // Labrador
	d.Speak()            // Animal speaks
}
```

在上面的示例中，定义了一个 `Animal` 结构体，其中包含一个 `Name` 字段和一个 `Speak` 方法。然后，我们定义了一个 `Dog` 结构体，其中嵌入了 `Animal` 结构体，并添加了一个 `Breed` 字段。

通过组合，`Dog`结构体继承了 `Animal` 结构体的字段和方法。我们可以通过访问 `Dog` 结构体的字段和调用 `Dog` 结构体的方法来间接访问和使用 `Animal` 结构体的字段和方法。

### 接口

通过定义接口，可以实现多态，使不同的类型可以被统一处理。**一个类型只需要实现了接口中定义的方法，就可以被视为该接口的实现类型**。

下面是一个示例代码，展示了如何使用接口来实现类似继承的效果：[传送门](https://go.dev/play/p/rztKZSAd13_u)

```go
package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Speak() {
	fmt.Println("Dog barks")
}

type Cat struct {
	Name  string
	Color string
}

func (c *Cat) Speak() {
	fmt.Println("Cat meows")
}

func main() {
	animals := []Animal{
		&Dog{Name: "Buddy", Breed: "Labrador"},
		&Cat{Name: "Whiskers", Color: "Gray"},
	}

	for _, animal := range animals {
		animal.Speak()
	}
}
```

在上面的示例中，定义了一个 `Animal` 接口，其中包含一个 `Speak` 方法。然后，又定义了 `Dog` 和 `Cat` 两个结构体，并分别实现了 `Speak` 方法。

通过将 `Dog` 和 `Cat` 类型的实例赋值给 `Animal` 类型的变量，我们可以将它们视为 `Animal` 接口的实现类型。在循环中，统一调用 `Speak` 方法，不管是 `Dog` 类型还是 `Cat` 类型，它们都能正确地执行各自的行为。

需要注意的是，在使用组合和接口实现类似继承的效果时，有一些差异和限制：

1. **组合的限制**：通过组合来实现类似继承的效果时，无法直接访问被嵌入结构体的私有字段和方法。只能通过嵌入结构体的公开字段和方法间接访问。
2. **方法重写**：通过组合或接口实现的方法，可以在子类型中进行重写，以改变其行为。这样可以实现多态性。
3. **父子类型转换**：通过组合或接口实现的父类型可以被转换为子类型，但需要进行显式的类型断言或类型转换操作。

### 结构体字段标签

在 Go 语言中，结构体字段标签（Struct Tag）是一种用于为结构体字段附加元数据的机制。**字段标签是一个字符串**，可以在**结构体字段的定义中使用反引号括起来，位于字段类型和字段名之间**。

结构体字段标签的主要作用是为结构体的字段提供额外的信息，例如字段的序列化格式、数据库映射、表单验证等。_标签字符串可以被反射机制读取和解析_，以便在运行时根据标签的信息进行相应的处理。

下面是一个示例代码，展示了如何使用结构体字段标签：

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender,omitempty"`
}

func main() {
	p := Person{
		Name:   "Forest",
		Age:    24,
		Gender: "",
	}

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, field.Tag.Get("json"))
	}
}
```

执行结果:

```sh
Field: Name, Tag: name
Field: Age, Tag: age
Field: Gender, Tag: gender,omitempty
```

在上面的示例中，定义了一个名为 `Person` 的结构体，其中包含了三个字段 `Name`、`Age`和 `Gender`。在每个字段的后面使用了反引号括起来的字符串，即结构体字段标签。

在 `main` 函数中，使用 `reflect` 包来获取结构体的类型信息，并遍历每个字段。通过调用 `Field` 方法获取字段的信息，然后使用 `Tag` 方法获取字段的标签。在输出中，我们打印了字段的名称和标签的值。

在示例中，结构体字段标签使用了 `json` 作为标签名称，并为每个字段指定了相应的标签值。这些标签值可以用于指定字段在进行JSON序列化和反序列化时的名称和行为。

> 通过结构体字段标签，我们可以实现以下功能：
>
> 1.  **序列化和反序列化**：可以使用结构体标签指定字段在序列化为JSON、XML等格式时的名称，并可选择性地忽略某些字段。
> 2.  **数据库映射**：可以使用结构体标签指定字段在数据库中的列名称、数据类型等信息，以便进行ORM操作。
> 3.  **表单验证**：可以使用结构体标签指定字段的验证规则，例如最大长度、正则表达式等。

需要注意的是，结构体字段标签的解析需要依赖反射机制，而反射对性能有一定的影响。因此，在实际使用中，应谨慎使用结构体字段标签，避免过度使用导致性能下降。

#### Tag 序列化的格式有以下几种

- JSON 格式

    ```go
    type Person struct {
     Name   string  `json:"name"`
     Age    int     `json:"age,string"`
     Height float64 `json:"height,number"`
     Email  string  `json:"email,omitempty"`
    }
    ```

    可以使用 `json` 标签来指定字段在 JSON 序列化和反序列化时的名称和行为。常用的标签选项有：
    - `omitempty`：如果字段的值为空值（零值或空引用），则在序列化时忽略该字段。
    - `string`：将字段的值转换为 JSON 字符串。
    - `number`：将字段的值转换为 JSON 数字。
    - `omitempty,number`：如果字段的值为空值，则在序列化时忽略该字段；否则，将字段的值转换为 JSON 数字。

    示例：[传送门](https://go.dev/play/p/orvih3kwelA)

    ```go
    package main

    import (
     "encoding/json"
     "fmt"
    )

    type Person struct {
     Name   string  `json:"name"`
     Age    int     `json:"age,string"`
     Height float64 `json:"height,number"`
     Email  string  `json:"email,omitempty"`
    }

    func main() {
     p := Person{
      Name:   "Forest",
      Age:    24,
      Height: 170.5,
      Email:  "",
     }

     // JSON 序列化
     jsonData, err := json.Marshal(p)
     if err != nil {
      fmt.Println("JSON serialization error:", err)
      return
     }

     fmt.Println(string(jsonData)) // {"name":"Forest","age":"24","height":170.5}

     // JSON 反序列化
     var p2 Person
     err = json.Unmarshal(jsonData, &p2)
     if err != nil {
      fmt.Println("JSON deserialization error:", err)
      return
     }

     fmt.Println(p2) // {Forest 25 170.5 }
    }
    ```

- XML 格式

    ```go
    type Person struct {
        Name   string `xml:"name"`
        Age    int    `xml:"age"`
        Gender string `xml:"gender,omitempty"`
    }
    ```

    可以使用`xml`标签来指定字段在 XML 序列化和反序列化时的名称和行为。常用的标签选项有：
    - `attr`：将字段序列化为 XML 元素的属性而不是子元素。
    - `omitempty`：如果字段的值为空值（零值或空引用），则在序列化时忽略该字段。

    示例：[传送门](https://go.dev/play/p/6ih2tFZh8zC)

    ```go
    package main

    import (
     "encoding/xml"
     "fmt"
    )

    type Person struct {
     Name   string `xml:"name"`
     Age    int    `xml:"age,omitempty"`
     Height int    `xml:"height,attr"`
    }

    func main() {
     p := Person{
      Name:   "Forest",
      Age:    0,
      Height: 170,
     }

     // XML 序列化
     xmlData, err := xml.Marshal(p)
     if err != nil {
      fmt.Println("XML serialization error:", err)
      return
     }

     fmt.Println(string(xmlData)) // <Person height="170"><name>Forest</name></Person>

     // XML 反序列化
     var p2 Person
     err = xml.Unmarshal(xmlData, &p2)
     if err != nil {
      fmt.Println("XML deserialization error:", err)
      return
     }

     fmt.Println(p2) // {Forest 0 170}
    }
    ```

- CSV 格式

    可以使用`csv`标签来指定字段在 CSV 序列化和反序列化时的名称和行为。

    ```go
    type Person struct {
        Name   string `csv:"name"`
        Age    int    `csv:"age"`
        Gender string `csv:"gender"`
    }
    ```

    示例：[传送门](https://go.dev/play/p/2XLLZry_MQz)

    ```go
    package main

    import (
     "encoding/csv"
     "fmt"
     "strings"
    )

    type Person struct {
     Name   string `csv:"name"` // 指定字段在CSV序列化和反序列化时的名称为"name"
     Age    int    `csv:"age"`  // 指定字段在CSV序列化和反序列化时的名称为"age"
     Height int    `csv:"-"`    // 忽略该字段在CSV序列化和反序列化时的行为
    }

    func main() {
     p := Person{
      Name:   "Forest",
      Age:    24,
      Height: 170,
     }

     // CSV 序列化
     csvData, err := toCSV(p)
     if err != nil {
      fmt.Println("CSV serialization error:", err)
      return
     }

     fmt.Println(csvData) // name,age,height,Forest,24,170

     // CSV 反序列化
     p2, err := fromCSV(csvData)
     if err != nil {
      fmt.Println("CSV deserialization error:", err)
      return
     }

     fmt.Println(p2) // {age 0 0}
    }

    // toCSV将给定的Person结构体转换为CSV格式的字符串
    func toCSV(p Person) (string, error) {
     fields := make([]string, 0)
     values := make([]string, 0)

     // 将name字段添加到fields和values切片中
     fields = append(fields, "name")
     values = append(values, p.Name)

     // 将age字段添加到fields和values切片中
     fields = append(fields, "age")
     values = append(values, fmt.Sprintf("%d", p.Age))

     // 如果Height字段不为0，则将height字段添加到fields和values切片中
     if p.Height != 0 {
      fields = append(fields, "height")
      values = append(values, fmt.Sprintf("%d", p.Height))
     }

     // 创建一个空的字符串切片来保存CSV记录
     record := make([]string, 0)
     // 将fields切片的内容追加到record切片中
     record = append(record, fields...)
     // 将values切片的内容追加到record切片中
     record = append(record, values...)

     // 创建一个strings.Builder来构建CSV数据
     w := &strings.Builder{}
     // 创建一个csv.Writer，并将其与strings.Builder关联
     csvWriter := csv.NewWriter(w)
     // 将record作为CSV记录写入csv.Writer
     if err := csvWriter.Write(record); err != nil {
      return "", err
     }
     // 刷新csv.Writer以确保所有数据都写入strings.Builder
     csvWriter.Flush()
     // 检查csv.Writer是否有错误
     if err := csvWriter.Error(); err != nil {
      return "", err
     }

     // 返回strings.Builder中的CSV数据作为字符串
     return w.String(), nil
    }

    // fromCSV将给定的CSV格式字符串转换回Person结构体
    func fromCSV(csvData string) (Person, error) {
     // 创建一个csv.Reader，将其与给定的CSV数据关联
     r := csv.NewReader(strings.NewReader(csvData))
     // 读取CSV数据中的一行记录
     record, err := r.Read()
     if err != nil {
      return Person{}, err
     }

     // 创建一个空的Person结构体
     p := Person{}
     // 遍历CSV记录中的每个字段和值对
     for i := 0; i < len(record); i += 2 {
      field := record[i]
      value := record[i+1]

      // 根据字段的名称将值分配给Person结构体的相应字段
      switch field {
      case "name":
       p.Name = value
      case "age":
       fmt.Sscanf(value, "%d", &p.Age)
      }
     }

     // 返回解析后的Person结构体
     return p, nil
    }
    ```
