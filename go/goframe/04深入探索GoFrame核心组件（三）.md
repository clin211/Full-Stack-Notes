大家好，我是长林啊！一名全栈开发者，同时也是一个 Go、Rust 爱好者；致力于终shen学习和技术分享。

在《GoFrame入门指南：环境搭建及框架初识》中，我们掌握了GoFrame的基本概念，学习了环境搭建步骤和常用命令。在《深入探索GoFrame核心组件（一）》中，我们详细讨论了对象管理、调试模式、命令管理、配置管理、日志和错误处理等关键模块。《深入探索GoFrame核心组件（二）》则深入研究了数据校验、类型转换和缓存管理。接下来，我们将继续探讨模板引擎、国际化（I18N），一起解锁更多GoFrame的高级使用技巧。

为了本文的代码演示和实操，我们先创建一个新的 goframe 项目，操作如下：
```sh
gf init corex -u
```
<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/TwdnkV-9f32937b-4316-4c92-b7d2-a657ceaa5d89.png" />

> "corex" 可以被理解为核心组件的代号或项目名称。它传达了这个组件在整个系统中具有重要性和核心功能的意义。

创建完成后使用自己顺手的 IDE 工具打开，然后就继续往下推进。

## 模板引擎

模板引擎特点：
- 简单、易用、强大；
- 支持多模板目录搜索；
- 支持 layout 模板设计；
- 支持模板视图对象单例模式；
- 与配置管理模块原生集成，使用方便；
- 底层采用了二级缓存设计，性能高效；
- 新增模板标签及大量的内置模板变量、模板函数；
- 支持模板文件修改后自动更新缓存机制，对开发过程更友好；
- `define`/`template` 标签支持跨模板调用（同一模板路径包括子目录下的模板文件）；
- `include` 标签支持任意路径的模板文件引入；

### 通用视图管理

通用视图管理即使用原生的模板引擎gview模块来实现模板管理，包括模板读取展示，模板变量渲染等等。可以使用通过方法 `gview.Instance` 来获取视图单例对象，并可以按照单例对象名称进行获取。同时也可以通过对象管理器 `g.View()` 来获取默认的单例 `gview` 对象。

### 目录配置方法

GoFrame 框架的模板引擎支持多目录自动搜索功能。使用 `SetPath` 可以设定模板目录为唯一的地址，而 `AddPath` 方法则允许添加多个搜索目录。模板引擎会按照添加的目录顺序依次搜索，找到第一个匹配的文件路径后停止。如果所有搜索目录都没有找到匹配的模板文件，则搜索失败。

> 默认目录配置—— `gview` 视图对象初始化时，默认会自动添加以下模板文件搜索目录：
> - 当前工作目录及其下的 `template` 目录
> - 当前可执行文件所在目录及其下的 `template` 目录
> - 当前 `main` 源代码包所在目录及其下的template目录

### 修改模板目录
Goframe 提供了三种修改模板目录的方式：
1. 单例模式获取全局 `View` 对象，通过 `SetPath` 方法手动修改；
    ```go
    g.View().SetPath("/opt/template")
    ```
2. 修改命令行启动参数——`gf.gview.path`；
    ```sh
    ./main --gf.gview.path=/opt/template/
    ```
3. 修改指定的环境变量——`GF_GVIEW_PATH`；
    - 启动时修改环境变量：
      ```sh
      $ GF_GVIEW_PATH=/opt/config/; gf run main.go
      ```
    - 使用genv模块来修改环境变量：
      ```go
      genv.Set("GF_GVIEW_PATH", "/opt/template")
      ```
### 自动检测更新
模板引擎使用了缓存机制，第一次读取模板文件后会将其缓存到内存，提升执行效率。同时，引擎支持自动检测更新机制，能够即时监控并刷新外部修改的模板文件缓存。
> 目前，GoFrame的模板引擎不支持像前端框架Vue、React那样的热更新功能。因此，每次修改 `.tpl` 文件后，需要手动刷新浏览器，这一点使用体验较差。

### 模板配置

#### 配置文件

视图组件支持配置文件，当使用 `g.View` (单例名称)获取 View 单例对象时，将会自动通过默认的配置管理对象获取对应的 `View` 配置。默认情况下会读取 `viewer.` 单例名称配置项，当该配置项不存在时，将会读取 `viewer` 配置项。

完整配置文件配置项及说明如下，其中配置项名称不区分大小写：
```yml
[viewer]
    Paths       = ["/var/www/template"] # 模板文件搜索目录路径，建议使用绝对路径。默认为当前程序工作路径
    DefaultFile = "index.html"          # 默认解析的模板引擎文件。默认为"index.html"
    Delimiters  =  ["${", "}"]          # 模板引擎变量分隔符号。默认为 ["{{", "}}"]
    AutoEncode  = false                 # 是否默认对变量内容进行XSS编码。默认为false
    [viewer.Data]                       # 自定义的全局Key-Value键值对，将在模板解析中可被直接使用到
        Key1 = "Value1"
        Key2 = "Value2"
```

- 默认配置项
  ```yml
  [viewer]
      paths       = ["template", "/var/www/template"]
      defaultFile = "index.html"
      delimiters  =  ["${", "}"]
      [viewer.data]
          name    = "gf"
          version = "1.10.0"
  ```
  
- 多个配置项
  ```yml
  [viewer]
      paths       = ["template", "/var/www/template"]
      defaultFile = "index.html"
      delimiters  =  ["${", "}"]
      [viewer.data]
          name    = "gf"
          version = "1.10.0"
      [viewer.view1]
          defaultFile = "layout.html"
          delimiters  = ["${", "}"]
      [viewer.view2]
          defaultFile = "main.html"
          delimiters  = ["#{", "}"]
  ```
  

通过单例对象名称获取对应配置的 View 单例对象:
```go
// 对应 viewer.view1 配置项
v1 := g.View("view1")
// 对应 viewer.view2 配置项
v2 := g.View("view2")
// 对应默认配置项 viewer
v3 := g.View("none")
// 对应默认配置项 viewer
v4 := g.View()
```
#### 配置方法
- 可以通过 `SetConfig` 及 `SetConfigWithMap` 来设置。
- 也可以使用 View 对象的 `Set*`方法进行特定配置的设置。
- 主要注意的是，配置项在 View 对象执行视图解析之前设置，避免并发安全问题。

使用 `SetConfigWithMap` 方法通过 `Key-Value` 键值对来设置/修改 View的特定配置(其中Key的名称即是Config这个struct中的属性名称，并且不区分大小写，单词间也支持使用-/_/空格符号连接)
```go
package main

import (
    "context"
    "fmt"

    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gview"
)

func main() {
    view := gview.New()
    view.SetConfigWithMap(g.Map{
        "Paths":       []string{"template"},
        "DefaultFile": "index.html",
        "Delimiters":  []string{"${", "}"},
        "Data": g.Map{
            "name":    "gf",
            "version": "1.10.0",
        },
    })
    result, err := view.ParseContent(context.TODO(), "hello ${.name}, version: ${.version}")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

`DefaultFile` 表示默认解析的模板文件，键名也可以使用
- `defaultFile`
- `default-File`
- `default_file`
- `default file`


看到上面的示例介绍是不是云里雾里？你是不是也有这些问题：
- 在哪里配置？
- 怎么配置？

下面就来用一个实际的例子还演示一下。就以上面配置项的例子来处理，文档中给的示例示例是 toml 格式的，因为我在项目用的是 yml，所以需要写yml的格式。如下：
```yml
viewer:
  paths:
    - "/resource/template"
  defaultFile: "index.html"
  delimiters:
    - "${"
    - "}"
  data:
    name: "gf"
    version: "1.10.0"
```
然后在项目的根目录 `/resource/template` 下创建三个 HTML 文件：index.html。其内容如下：
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    this is template
    <p>name: ${.nickname}</p>
    <p>age: ${.age}</p>
</body>

</html>
```
然后再cmd中注册一个路由：
```go
s.BindHandler("/index", func(r *ghttp.Request) {
    r.Response.WriteTpl("index.html", data)
})
```
完整代码如下：
```go
// internal/cmd/cmd.go
package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"corex/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			data := map[string]any{
				"nickname": "forest",
				"age":      "18",
			}
			s.BindHandler("/index", func(r *ghttp.Request) {
				r.Response.WriteTpl("index.html", data)
			})
			s.Run()
			return nil
		},
	}
)
```

使用命令 `gf run main.go` 运行项目：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/K4KRiq-36a67ff0-3142-4fd2-ab57-18ae5f947137.png" />

在浏览器中访问 `loclahost:8000/index` 效果如下：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/aZjxeg-145d761d-1cdc-4be2-b080-2a449678f4a6.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />

上面的示例代码可以在仓库 [clin211/goframe-practice/tree/template-config](https://github.com/clin211/goframe-practice/commit/4e19a65d380fb386d104b01607f7ee71764d2dd2) 中找到。

> 为什么他们的模板解析没有生效？为什么页面上直接显示他们写的标签原样而不是解析后的结果？
>
> - 检查你的配置文件中是否正确设置了模板标签。常见问题是 `delimiters` 设置为 `["${", "}"]`，但实际模板中使用的是 `["{{", "}}"]`。
>
> - 确保配置文件中的标签设置与实际模板代码一致，以确保模板能够正确解析和渲染。

### 模板标签

模板引擎默认使用 `{{` 和 `}}` 作为左右闭合标签。开发者可以通过 `gview.SetDelimiters` 方法设置自定义的模板闭合标签。

在模板中，使用 `.` 来访问当前对象的值（模板局部变量）。

使用 `$` 来引用当前模板根级的上下文（模板全局变量）。

使用 `$var` 来访问特定的模板变量。

> 下面的代码还是沿用之前在config中的配置 ${} 定界符

准备工作：
- 注册一个新的路由 'template-tags'
  ```go
  s.BindHandler("/template-tags", func(r *ghttp.Request) {
      r.Response.WriteTpl("template-tags.html", map[string]any{})
  })
  ```
- 创建一个新的 html 页面——template-tags.html
  ```html
  <!DOCTYPE html>
  <html lang="en">
  
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>模板标签</title>
  </head>
  
  <body>
  
  </body>
  
  </html>
  ```

#### 模板中支持Go语言表达式
```go
<h2> 模板中支持的 Go 符号</h2>
<ul>
    <li>${"string"} // 一般 string</li>
    <li>${`raw string`} // 原始 string</li>
    <li>${'c'} // byte</li>
    <li>${print nil} // nil 也被支持</li>
</ul>
```
效果如下：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/6O9Yh5-34004ac8-7040-44c5-83c9-5ed0e4ca5287.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />

#### 模板中的 pipeline
可以是上下文的变量输出，也可以是函数通过管道传递的返回值。语法如下：
```go
${. | FuncA | FuncB | FuncC}
```
当 pipeline 的值等于:
- `false`或 `0`
- `nil` 的指针或 `interface`
- 长度为0的 `array`, `slice`, `map`, `string`

那么这个 pipeline 被认为是空。

#### 条件判断
if判断时，pipeline为空时，相当于判断为false。
- 单行
```go
{{if pipeline}}...{{end}}
```
- 多行
```go
{{if .condition}}
    ...
{{else if .condition2}}
    ...
{{else}}
    ...
{{end}}
```

- 嵌套
```go
{{if .condition}}
    ...
{{else}}
	{{if .condition2}}
        ...
    {{end}}
{{end}}
```
示例如下：
```html
<h2> 模板中的判断</h2>
<p>
    ${if .Name}
    hello ${.Name}
    ${else}
    default name
    ${end}
</p>

<h3>判断嵌套示例</h3>
${if .User}
<p>用户名: ${.User.Name}</p>
${if .User.Admin}
<p>角色: 管理员</p>
${else}
<p>角色: 普通用户</p>
${end}
${else}
<p>用户未登录</p>
${end}
```

#### 循环
```go
{{range pipeline}} {{.}} {{end}}
```
pipeline 支持的类型为 `slice`, `map`, `channel`。

注意：在 range循环内部，`.` 符号会被覆盖为以上类型的子元素（局部变量）。如果想在循环中访问外部变量（全局变量），请加上 `$` 符号，如：`{{$.Session.Name}}`

此外，对应的值长度为 0 时，range 不会执行，`.` 不会改变。

在 HTML 中添加如下内容：
```html
<h2>用户列表</h2>
<ul>
    ${range .users}
    <li>
        名字: ${.Name}, 年龄: ${.Age}
        ${if .Admin}
        (管理员)
        ${end}
    </li>
    ${else}
    <li>没有用户</li>
    ${end}
</ul>
```
在cmd中添加数据：
```go
s.BindHandler("/template-tags", func(r *ghttp.Request) {
    // 创建用户数据列表
    var users = []struct {
        Name  string
        Age   int
        Admin bool
    }{
        {Name: "Alice", Age: 30, Admin: true},
        {Name: "Bob", Age: 25, Admin: false},
        {Name: "Charlie", Age: 28, Admin: false},
    }

    r.Response.WriteTpl("template-tags.html", g.Map{
        "data":  data,
        "users": users,
    })
})
```
效果如下：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/AqP2fc-a4f875c8-8e59-436d-a68e-cf734627be4e.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />

#### with...end
`with...end` 是 GoFrame 模板引擎提供的一种语法结构，用于简化模板中访问结构体字段的代码。它允许在一个块中指定一个对象，并在块内部直接访问该对象的字段，而不需要重复指定对象名称。

使用 with ... end 结构，可以**将一个对象绑定到一个作用域内，在作用域内可以直接访问该对象的字段，而无需重复使用点操作符来指定对象名称**。
```html
<h2>用户信息</h2>
${with .user}
<p>名字: ${.Name}</p>
<p>年龄: ${.Age}</p>
${if .Admin}
<p>角色: 管理员</p>
${else}
<p>角色: 普通用户</p>
${end}
${else}
<p>用户不存在</p>
${end}
```
在上面的模板中，我们使用了 with .User 结构来指定当前作用域的对象为 .User，这意味着在 with ... end 块内部可以直接访问 User 结构体的字段，比如 Name、Age 和 Admin。如果 .User 为 nil 或未定义，则会执行 else 分支显示 "用户不存在"。

然后 在cmd 中添加一个包含 User 结构体的数据给模板如下：
```go
r.Response.WriteTpl("template-tags.html", g.Map{
    "data":  data,
    "users": users,
    "user":  User{Name: "forest", Age: 18, Admin: true},
})
```
效果如下：
<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/le3WD3-1471d1c3-92fb-45f5-9f85-b5dd38d145e6.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />


#### define
在 GoFrame 的模板引擎中，`define` 是用来**定义和命名模板块的关键字，可以将一段模板代码片段定义为一个可重用的块，并在需要的地方引用它**。这种方式可以提高代码的复用性和可维护性。

使用 `define` 关键字可以将一段模板内容定义为一个块，语法结构如下：
```html
{{define "blockName"}}
    <!-- 模板内容 -->
{{end}}
```
- "blockName" 是块的名称，用于标识和引用该块。
- {{end}} 用于结束块的定义。

定义好的模板块可以在其他地方通过 template 关键字引用和渲染。

假设我们要定义一个头部导航栏的模板块 header，包含网站的标题和导航链接，在template 目录下创建一个 components 的目录，然后创建 `header.html` 文件，并添加以下信息：

```html
${define "header"}
<header>
    <h1>${.title}</h1>
    <nav>
        <ul>
            <li><a href="/">首页</a></li>
            <li><a href="/about">关于我们</a></li>
            <li><a href="/contact">联系我们</a></li>
        </ul>
    </nav>
</header>
${end}
```
header 块定义了一个简单的网站头部结构，包括网站标题和导航链接。然后，我们可以在其他模板文件中使用 template 关键字来引用和渲染这个块，如下所示：
```html
<h2>模板中的应用</h2>
${template "header" .}

<section>
    <!-- 其他页面内容 -->
</section>
```
> `define` 标签需要结合 template 标签一起使用，并且支持跨模板使用（在同一模板目录/子目录下有效，原理是使用的 `ParseFiles` 方法解析模板文件）

通过 `${template "header" .}` 引用了名为 header 的模板块，并将当前上下文 `.` 传递给模板，使得模板可以访问当前页面的数据。

然后在 cmd 中添加一个 title 的数据：
```go
r.Response.WriteTpl("template-tags.html", g.Map{
    "data":  data,
    "users": users,
    "title": "GoFrame Template", // ----- 添加这行
    "user":  User{Name: "forest", Age: 18, Admin: true},
})
```

在浏览器中访问，效果如下：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/nB0nti-ad002aa6-09ea-451a-85ef-5274664b8f2f.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />

> **注意事项:**
> - 模板块的作用域： **定义的模板块只在当前模板文件中有效**，除非使用**模板继承或包含的方式在其他模板中使用**。
> - 命名规范： **块名称应当具有描述性**，以便于理解和维护代码。
> - 代码复用： 使用 `define` 和 `template` 可以有效地提高模板代码的复用性，**避免重复编写相似的代码片段**。

#### template
```html
{{template "模板名称" pipeline}}
```
- `{{template "模板名称" pipeline}}` 中的 `"模板名称"` 是你想**要引用的模板的名称**。
- pipeline 是一个用于传递数据的表达式，可以是**变量**、**字段**、**函数**等，**它将作为被引用模板的数据源**。

> 注意：**template 标签参数为模板名称，而不是模板文件路径，template 标签不支持模板文件路径**。

#### include
在 GoFrame 框架中，模板的 `include` 功能允许在一个模板文件中引用和嵌入另一个模板文件的内容，从而实现模板的复用和模块化。通过引用其他模板文件来实现页面模块的组合。

```html
{{include "模板文件名(需要带完整文件名后缀)" pipeline}}
```
在模板中可以使用 `include` 标签载入其他模板（任意路径），模板文件名支持相对路径以及文件的系统绝对路径。如果想要把当前模板的模板变量传递给子模板(嵌套模板)，可以这样：
```
{{include "模板文件名(需要带完整文件名后缀)" .}}
```
与 template 标签的区别是：
- include 仅支持文件路径，不支持模板名称。
- tempalte 标签仅支持模板名称，不支持文件路径。

#### 注释
允许多行文本注释，不允许嵌套。
```go
{{/*
comment content
support new line
*/}}
```
上面的所有代码都可以在 [clin211/goframe-practice/tree/template-tags](https://github.com/clin211/goframe-practice/tree/template-tags) 中找到。

### 模板函数
在 GoFrame 框架的模板引擎中，支持一些内置的函数和操作符，用于在模板中进行数据操作和控制流处理。

为了演示下面的内容，在 cmd 中新注册一个路由 `/func`，代码如下：
```go
// ...
s.BindHandler("/func", func(r *ghttp.Request) {
    r.Response.WriteTpl("template-func.html", g.Map{
        "LoggedIn": true,
        "Admin":    true,
    })
})
// ...
```
1. `and`: 逻辑与操作符，用于判断多个条件是否同时成立。`and` 会**逐一判断每个参数，将返回第一个为空的参数，否则就返回最后一个非空参数。**
    ```html
    <!DOCTYPE html>
    <html lang="en">
    
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>模板函数</title>
    </head>
    
    <body>
        ${if and .LoggedIn .Admin}
        <!-- 显示管理员特有内容 -->
        <p>管理员权限内容</p>
        ${end}
    </body>
    
    </html>
    ```

    然后再 `/func` 路由中添加数据：
    ```go
    r.Response.WriteTpl("template-func.html", g.Map{
        "LoggedIn": true,
        "Admin":    true,
    })
    ```
    效果如下：
    
    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/af13wf-a3ad742b-2a8f-4dfe-a59c-7013a23e5940.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />  


2. `call`: 调用指定的模板函数。`call` 可以调用函数，并传入参数；调用的函数需要返回 1 个值或者 2 个值，返回两个值时，第二个值用于返回 `error` 类型的错误。返回的错误不等于 `nil` 时，执行将终止。

    > - `.FuncName`: 要调用的**模板函数的名称**。
    > - `"arg1"`, `"arg2"`: **作为参数传递给模板函数的参数，可以是变量、常量或表达式**。
    
    下面来做一个根据当前时间戳格式化为 YYYY:MM:DD HH:mm:ss 的形式：
    - 创建一个路由(examples/func/call/main.go)
      ```go
      package main
    
      import (
          "time"
    
          "github.com/gogf/gf/frame/g"
          "github.com/gogf/gf/net/ghttp"
      )
    
      // formatDate 函数用于将时间格式化为指定的字符串格式
      func formatDate(date time.Time) string {
          // 使用Format方法将时间格式化为"2006-01-02 15:04:05"的形式
          return date.Format("2006-01-02 15:04:05")
      }
    
      func main() {
          // 创建一个GoFrame的服务器实例
          s := g.Server()
    
          // 为服务器绑定一个处理根路径("/")的处理器
          s.BindHandler("/", func(r *ghttp.Request) {
              // 当请求根路径时，使用WriteTpl方法渲染模板并返回响应
              // 将当前时间作为参数"date"传递给模板，并提供formatDate函数用于时间格式化
              r.Response.WriteTpl("cell.html", g.Map{
                  "date":       time.Now(), // 获取当前时间
                  "formatDate": formatDate, // 将formatDate函数传递给模板
              })
          })
    
          // 设置服务器监听的端口号为 8199
          s.SetPort(8199)
    
          // 启动服务器，监听并处理HTTP请求
          s.Run()
      }
      ```
    - 创建模版
      ```html
      <!-- examples/func/call/template/cell.html -->
      <!DOCTYPE html>
      <html lang="en">
    
      <head>
          <meta charset="UTF-8">
          <title>Call Function Example</title>
      </head>
    
      <body>
          <div>
              <!-- 调用 formatDate 函数，并传递参数 "2024-07-10" -->
              <p>格式化后的日期：{{call .formatDate .date}}</p>
          </div>
      </body>
    
      </html>
      ```
    - 启动项目 `go run main.go`，然后在浏览器中访问 http://localhost:8199/ 查看效果
    
      <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/PI6xV0-29083c7f-bbdd-482a-8e14-524301021e65.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />

3. `index`：读取指定类型对应下标的值
    - 创建一个完整代码如下：
      ```go
      // examples/func/index/main.go
      package main
      
      import (
          "github.com/gogf/gf/frame/g"
          "github.com/gogf/gf/net/ghttp"
      )
      
      func main() {
          // 创建一个GoFrame的服务器实例
          s := g.Server()
      
          // 为服务器绑定一个处理根路径("/")的处理器
          s.BindHandler("/", func(r *ghttp.Request) {
              // 当请求根路径时，使用WriteTpl方法渲染模板并返回响应
              // 将当前时间作为参数"date"传递给模板，并提供formatDate函数用于时间格式化
              r.Response.WriteTpl("index.html", g.Map{
                  "Map":    map[string]any{"key": "this is map"},
                  "Slice":  []string{"this is slice 1", "this is slice 2"},
                  "Array":  [2]string{"this is array 1", "this is array 2"},
                  "String": "hello features",
              })
          })
      
          // 设置服务器监听的端口号为 8199
          s.SetPort(8199)
      
          // 启动服务器，监听并处理HTTP请求
          s.Run()
      }
      ```
    - 创建模板如下：
      ```html
      <!DOCTYPE html>
      <html lang="en">
      
      <head>
          <meta charset="UTF-8">
          <meta name="viewport" content="width=device-width, initial-scale=1.0">
          <title>index</title>
      </head>
      
      <body>
          <p>读取 map `{{.Map}}` 的 key 值:{{index .Map "key"}}</p>
          <p>读取 slice `{{.Slice}}` 的第一个值:{{index .Slice 0}}</p>
          <p>读取 array `{{.Array}}` 的第一个值:{{index .Array 0}}</p>
          <p>读取 string `{{.String}}` 的第二个值:{{printf "%c" (index .String 1)}}</p>
      </body>
      
      </html>
      ```
    - 运行项目 `go run main.go`，然后在浏览器中访问 http://localhost:8199/ 查看效果
    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/AP1d5B-efd0175c-0334-42c2-9a51-3cf080611166.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />
    
4. `len`：返回对应类型的长度，支持类型：`map`, `slice`, `array`, `string`, `chan`。 
    - HTML 的内容
      ```html
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">len 方法</h2>
          <p>len 方法返回 string--{{.len.string }} 的长度：{{len .len.string}}</p>
          <p>len 方法返回 slice--{{.len.slice}} 的长度：{{len .len.slice}}</p>
          <p>len 方法返回 map--{{.len.map}} 的长度：{{len .len.map}}</p>
          <p>len 方法返回 array--{{.len.array}} 的长度：{{len .len.array}}</p>
          <p>len 方法返回 chan--{{printf "%T" .len.chan}} 的长度：{{len .len.chan}}</p>
      </div>
      ```
    - 逻辑
      ```go
      "len": map[string]any{
          "string": "hello features",
          "array":  [5]string{"a", "b", "c", "d", "e"},
          "slice":  []string{"a", "b", "c", "d", "e"},
          "map":    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
          "chan":   make(chan int, 10),
      },
      ```
    - 效果如下
      <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/52Lve2-ebf0751a-c3ff-4954-accb-2dbb16436ed6.png" />

5. `not`：返回输入参数的否定值。
    ```go
    {{if not .Var}}
    // 执行为空操作(.Var为空, 如: nil, 0, "", 长度为0的slice/map)
    {{else}}
    // 执行非空操作(.Var不为空)
    {{end}}
    ```
    - HTML 代码
      ```html
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">not 方法</h2>
          {{if not .len.string}}
          <p>.len.string 的值为空</p>
          {{else}}
          <p>.len.string 的值为：{{.len.string}}</p>
          {{end}}
      </div>
      ```
    - 沿用上面 len 方法的数据
    - 查看效果
    
    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/XMDI54-ceb3be13-cdb1-4eed-92bb-6c73ca02f296.png" />

6. `or`：会逐一判断每个参数，将返回第一个非空的参数，否则就返回最后一个参数。
    - HTML 代码
      ```html
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">or 方法</h2>
          <p>{{or (len .len.chan) .len.string}}</p>
      </div>
      ```
    - 仍然沿用上面 len 方法的数据
    - 查看效果如下
    
      <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/aRgJR9-6f4a30e3-b31d-4e7e-83e5-3f378dd51a28.png" /> 

7. `print` 同 `fmt.Sprint`、`printf` 同 `fmt.Sprintf`、`println` 同 `fmt.Sprintln`。
    - HTML 代码
      ```html
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">print/printf/println 方法</h2>
          <p>{{print .len.string}}</p>
          <p>{{printf "%v" .len.string}}</p>
          <p>{{println .len.string}}</p>
      </div>
      ```
    - 仍然沿用上面 len 方法的数据
    - 查看效果
    
    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/zVjzgQ-ca678697-c7d5-49fb-ae0e-77e58304688d.png" />

8. `urlquery`：函数允许将 URL Query 文本转义成安全文本；`{{urlquery "http://johng.cn"}}` 将返回 `http%3A%2F%2Fjohng.cn`。
9. 这类函数一般配合在if中使用的方法
    - `eq`：相等；支持多个参数；比如：`{{eq arg1 arg2 arg3 arg4}}` 逐一判断每个参数；与 `arg1==arg2 || arg1==arg3 || arg1==arg4` 逻辑判断相同。
    - `ne`：不相等
    - `lt`：小于
    - `le`：小于等于
    - `gt`：大于
    - `ge`：大于等于
    
    下面就来实际演示一下这些方法：
    - 创建一个模版
      ```html
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">eq/ne/lt/le/gt/ge 方法</h2>
          <table class="w-full table-auto rounded-lg border-collapse table">
              <thead class="text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  <tr class="text-gray-800 border border-gray-400">
                      <th class="px-4 py-2 border border-gray-400">Operation</th>
                      <th class="px-4 py-2 border border-gray-400">description</th>
                      <th class="px-4 py-2 border border-gray-400">Example</th>
                      <th class="px-4 py-2 border border-gray-400">Result</th>
                  </tr>
              </thead>
              <tbody class="text-gray-700">
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">"a" == "a"</td>
                      <td class="px-4 py-2 border border-gray-400">eq "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq "a" "a"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">"1" == "1"</td>
                      <td class="px-4 py-2 border border-gray-400">eq "1" "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq "1" "1"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">1 == 1</td>
                      <td class="px-4 py-2 border border-gray-400">eq 1 "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq 1 "1"}}</td>
                  </tr>
      
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"1" != "1"</td>
                      <td class="px-4 py-2 border border-gray-400">ne 1 "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne 1 "1"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"a" != "a"</td>
                      <td class="px-4 py-2 border border-gray-400">ne "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne "a" "a"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"a" != "b"</td>
                      <td class="px-4 py-2 border border-gray-400">ne "a" "b"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne "a" "b"}}</td>
                  </tr>
      
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">"1" < "2" </td>
                      <td class="px-4 py-2 border border-gray-400">lt 1 "2"</td>
                      <td class="px-4 py-2 border border-gray-400">{{lt 1 "2"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">2 < 2</td>
                      <td class="px-4 py-2 border border-gray-400">lt 2 2</td>
                      <td class="px-4 py-2 border border-gray-400">{{lt 2 2}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">"a" < "b" </td>
                      <td class="px-4 py-2 border border-gray-400">lt "a" "b"</td>
                      <td class="px-4 py-2 border border-gray-400">{{le "a" "b"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">"1" >= "2"</td>
                      <td class="px-4 py-2 border border-gray-400">ge 1 "2"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge 1 "2"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">2 >= 1</td>
                      <td class="px-4 py-2 border border-gray-400">ge 2 1</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge 2 1}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">"a" >= "a"</td>
                      <td class="px-4 py-2 border border-gray-400">ge "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge "a" "a"}}</td>
                  </tr>
              </tbody>
          </table>
      </div>
      ```

    - 查看效果
      <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/KHgmLQ-150913f9-0e7b-4215-a773-eb484b3fa984.png" />
      
      
      

上面这个几个示例的完整代码如下：
- HTML 的代码
  ```html
  <!DOCTYPE html>
  <html lang="en">
  
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>剩下的其他方法</title>
      <script src="https://cdn.tailwindcss.com"></script>
      <style>
          thead {
              background-color: #f7fafc;
              /* 表头背景色 */
          }
  
          tbody tr:nth-child(odd) {
              background-color: #edf2f7;
              /* 奇数行背景色 */
          }
      </style>
  </head>
  
  <body>
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">len 方法</h2>
          <p>len 方法返回 string--{{.len.string }} 的长度：{{len .len.string}}</p>
          <p>len 方法返回 slice--{{.len.slice}} 的长度：{{len .len.slice}}</p>
          <p>len 方法返回 map--{{.len.map}} 的长度：{{len .len.map}}</p>
          <p>len 方法返回 array--{{.len.array}} 的长度：{{len .len.array}}</p>
          <p>len 方法返回 chan--{{printf "%T" .len.chan}} 的长度：{{len .len.chan}}</p>
      </div>
  
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">not 方法</h2>
          {{if not .len.string}}
          <p>.len.string 的值为空</p>
          {{else}}
          <p>.len.string 的值为：{{.len.string}}</p>
          {{end}}
      </div>
  
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">or 方法</h2>
          <p>{{or (len .len.chan) .len.string}}</p>
      </div>
  
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">print/printf/println 方法</h2>
          <p>{{print .len.string}}</p>
          <p>{{printf "%v" .len.string}}</p>
          <p>{{println .len.string}}</p>
      </div>
  
      <div class="container mx-auto px-4 py-2">
          <h2 class="text-xl font-bold mb-4">eq/ne/lt/le/gt/ge 方法</h2>
          <table class="w-full table-auto rounded-lg border-collapse table">
              <thead class="text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  <tr class="text-gray-800 border border-gray-400">
                      <th class="px-4 py-2 border border-gray-400">Operation</th>
                      <th class="px-4 py-2 border border-gray-400">description</th>
                      <th class="px-4 py-2 border border-gray-400">Example</th>
                      <th class="px-4 py-2 border border-gray-400">Result</th>
                  </tr>
              </thead>
              <tbody class="text-gray-700">
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">"a" == "a"</td>
                      <td class="px-4 py-2 border border-gray-400">eq "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq "a" "a"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">"1" == "1"</td>
                      <td class="px-4 py-2 border border-gray-400">eq "1" "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq "1" "1"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">eq</td>
                      <td class="px-4 py-2 border border-gray-400">1 == 1</td>
                      <td class="px-4 py-2 border border-gray-400">eq 1 "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{eq 1 "1"}}</td>
                  </tr>
  
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"1" != "1"</td>
                      <td class="px-4 py-2 border border-gray-400">ne 1 "1"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne 1 "1"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"a" != "a"</td>
                      <td class="px-4 py-2 border border-gray-400">ne "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne "a" "a"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ne</td>
                      <td class="px-4 py-2 border border-gray-400">"a" != "b"</td>
                      <td class="px-4 py-2 border border-gray-400">ne "a" "b"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ne "a" "b"}}</td>
                  </tr>
  
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">"1" < "2" </td>
                      <td class="px-4 py-2 border border-gray-400">lt 1 "2"</td>
                      <td class="px-4 py-2 border border-gray-400">{{lt 1 "2"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">2 < 2</td>
                      <td class="px-4 py-2 border border-gray-400">lt 2 2</td>
                      <td class="px-4 py-2 border border-gray-400">{{lt 2 2}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">lt</td>
                      <td class="px-4 py-2 border border-gray-400">"a" < "b" </td>
                      <td class="px-4 py-2 border border-gray-400">lt "a" "b"</td>
                      <td class="px-4 py-2 border border-gray-400">{{le "a" "b"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">"1" >= "2"</td>
                      <td class="px-4 py-2 border border-gray-400">ge 1 "2"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge 1 "2"}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">2 >= 1</td>
                      <td class="px-4 py-2 border border-gray-400">ge 2 1</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge 2 1}}</td>
                  </tr>
                  <tr>
                      <td class="px-4 py-2 border border-gray-400">ge</td>
                      <td class="px-4 py-2 border border-gray-400">"a" >= "a"</td>
                      <td class="px-4 py-2 border border-gray-400">ge "a" "a"</td>
                      <td class="px-4 py-2 border border-gray-400">{{ge "a" "a"}}</td>
                  </tr>
              </tbody>
          </table>
      </div>
  </body>
  
  </html>
  ```
- Go 代码
  ```go
  // examples/func/index/main.go
  package main
  
  import (
      "github.com/gogf/gf/frame/g"
      "github.com/gogf/gf/net/ghttp"
  )
  
  func main() {
      // 创建一个GoFrame的服务器实例
      s := g.Server()
  
      // 为服务器绑定一个处理根路径("/")的处理器
      s.BindHandler("/", func(r *ghttp.Request) {
          // 当请求根路径时，使用WriteTpl方法渲染模板并返回响应
          // 将当前时间作为参数"date"传递给模板，并提供formatDate函数用于时间格式化
          r.Response.WriteTpl("other.html", g.Map{
              "len": map[string]any{
                  "string": "hello features",
                  "array":  [5]string{"a", "b", "c", "d", "e"},
                  "slice":  []string{"a", "b", "c", "d", "e"},
                  "map":    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
                  "chan":   make(chan int, 10),
              },
          })
      })
  
      // 设置服务器监听的端口号为 8199
      s.SetPort(8199)
  
      // 启动服务器，监听并处理HTTP请求
      s.Run()
  }
  ```

#### 内置函数
- plus/minus/times/divide
- text
- htmlencode/encode/html
- htmldecode/decode
- urlencode/url
- urldecode
- date
- compare
- replace
- substr
- strlimit
- concat
- hidestr
- highlight
- toupper/tolower
- nl2br
- dump
- map
- maps
- json/xml/ini/yaml/yamli/toml

这些函数的用法比较简单，而且官方文档写得也比较详细，我这里就不在一一去复述了，不知道怎么用的时候，可以以看看[官方文档](https://goframe.org/pages/viewpage.action?pageId=1114270)。

#### 自定义函数

自定义函数的作用非常广泛，它们可以极大地增强模板的灵活性和功能性。以下是一些自定义函数的主要作用：
- 数据操作：自定义函数可以对传递给模板的数据进行处理，比如格式化、转换数据类型、计算等。
- 逻辑处理：可以在模板中实现一些简单的逻辑判断，比如根据条件显示不同的内容。
- HTML 属性生成：根据数据动态生成 HTML 属性，例如类名、样式或数据属性。
- 国际化：实现多语言支持，根据用户的语言偏好显示相应的文本。
- 安全性增强：例如，实现自定义的 HTML 转义函数，以防止跨站脚本（XSS）攻击。
- URL 生成：生成 URL，包括查询参数的拼接等。
- 模板继承：创建可重用的模板组件，如页眉、页脚、导航栏等。
- 复杂计算：执行更复杂的数学或逻辑计算，这些计算在模板中直接编写可能不够直观或容易出错。
- 数据聚合：对数据进行聚合操作，如求和、平均值、最大/最小值等。
- 美化输出：美化或格式化输出结果，如日期时间格式化、数字格式化等。

下面以自定义字符串的大写和小写的方法为例，演示如何自定义函数。

HTML 代码：
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>自定义函数</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
    <h2 class="text-xl font-bold mb-4">自定义函数——普通方式传参</h2>
    <p>uppercase "string"——{{uppercase "string"}}</p>
    <p>lowercase "string"——{{lowercase "STRING"}}</p>
    <h2 class="text-xl font-bold mb-4">自定义函数——通过管道方式传参</h2>
    <p>"string" | uppercase——{{"string" | uppercase}}</p>
    <p>"STRING" | lowercase——{{"STRING" | lowercase}}</p>
</body>

</html>
```

Go 代码：
```go
package main

import (
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

func main() {
	// 创建一个GoFrame的服务器实例
	s := g.Server()

	// 为模板绑定自定义函数
	view := g.View()
	view.BindFunc("uppercase", uppercase)
	view.BindFunc("lowercase", lowercase)

	// 为服务器绑定一个处理根路径("/")的处理器
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("custom.html", g.Map{})
	})

	// 设置服务器监听的端口号为 8199
	s.SetPort(8199)

	// 启动服务器，监听并处理HTTP请求
	s.Run()
}
```
在这个例子中，我们定义了一个 uppercase 和 lowercase 两个自定义函数，它将被模板用来处理字符串。然后我们使用 BindFunc 方法将这两个函数绑定到模板引擎中，这样在渲染模板时就可以像使用内置函数一样使用它。

在终端中运行 `go run main.go`命令后，在浏览器中访问 `http://localhost:8199` 查看效果如下：

<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/o3aVJI-a9d4511b-4be1-4208-9e8f-80a3e66c6ba8.png" style="border-radius: 12px;border:1px solid rgba(145, 109, 213, 0.5);" />


自定义函数使得模板不仅仅是简单的数据展示，而是可以进行复杂的数据处理和逻辑判断，从而创建更加动态和交互性强的 web 页面。

上面的所有代码都可以在 [clin211/goframe-practice/tree/template-functions](https://github.com/clin211/goframe-practice/tree/template-functions) 中查看。

## i18n 国际化
i18n 是 "internationalization" 的缩写，它是一个过程，通过这个过程，软件应用程序的设计和程序代码可以适应不同的语言和地区设置。

GoFrame框架提供了常用的I18N国际化组件，由gi18n模块实现。

### 配置管理
在应用程序中配置本地化设置，包括默认语言、支持的语言列表以及如何根据用户的地区或偏好选择语言等等。

#### 文件格式
gi18n 国际化组件支持框架通用的五种配置文件格式：`xml/ini/yaml/toml/json/properties`。同样的，和配置管理模块一样，框架推荐使用toml文件格式。

#### 读取路径
默认情况下 gi18n 会自动查找并读取**当前项目源码根目录**（或者当前PWD运行目录下）下的以下目录：
- manifest/i18n
- i18n

默认将查找到的目录作为国际化转译文件存储目录。开发者也可以通过 `SetPath` 方法自定义 i18n 文件的存储目录路径。

#### 文件存储
在 i18n 文件夹中，你可以通过两种方式组织国际化文件：

1. 直接使用语言代码命名文件，例如 `en.toml`、`ja.toml`、`zh-CN.toml`。
2. 创建语言代码为名称的子目录，然后在这些目录中放置自定义命名的配置文件，如 `en/editor.toml、zh-CN/editor.toml` 等。

gi18n 能够智能识别并加载这些文件，无论是单一文件还是目录结构。

> 国际化文件和目录的命名也可以由开发者自定义，建议遵循 ISO 语言代码标准。详见[ISO 639-1 WIKI](https://zh.wikipedia.org/wiki/ISO_639-1)。

- 可以通过单独的 i18n 文件区分不同的语言
  ```tree
  └── i18n
      ├── en.toml
      ├── ja.toml
      ├── ru.toml
      ├── zh-CN.toml
      └── zh-TW.toml
  ```
- 也可以通过不同的目录名称区分不同的语言
  ```tree
  └── i18n
      ├── en
      │   ├── hello.toml
      │   └── world.toml
      ├── ja
      │   ├── hello.yaml
      │   └── world.yaml
      ├── ru
      │   ├── hello.ini
      │   └── world.ini
      ├── zh-CN
      │   ├── hello.json
      │   └── world.json
      └── zh-TW
          ├── hello.xml
          └── world.xml
  ```
- 还可以不同语言可以存在不同文件格式
  ```tree
  └── i18n
      ├── en.toml
      ├── ja.yaml
      ├── ru.ini
      ├── zh-CN.json
      └── zh-TW.xml
  ```
  
> 资源管理器默认情况下会从 gres 配置管理器中检索 manifest/i18n 及i18n目录，或者开发者设置的 i18n 目录路径。

### 使用介绍
#### 对象创建
- goframe 推荐在大多数场景下使用 `g.I18n` 单例对象
  ```go
  g.I18n().T(context.TODO(), "{#hello} {#world}")
  ```

- 也可以通过 `gi18n.New()` 方法创建独立的 i18n 对象，然后开发者自行进行管理。
  ```go
  i18n := gi18n.New() 
  i18n.T(context.TODO(), "{#hello} {#world}")
  ```

#### 语言设置

设置转译语言有两种方式，一种是通过 `SetLanguage` 方法设置当前 I18N 对象统一的转译语言，另一种是通过上下文设置当前执行转译的语言。

- `SetLanguage` 通过 `g.I18n().SetLanguage("zh-CN")` 即可设置当前转译对象的转译语言，随后使用该对象都将使用 zh-CN 进行转译。
  > 需要注意的是，组件的配置方法往往都不是并发安全的，该方法也同样如此，需要在程序初始化时进行设置，随后不能在运行时进行更改。
  
- `WithLanguage` 方法可以创建一个新的上下文变量，并临时设置您当前转译的语言，由于该方法作用于Context上下文，因此是并发安全的，常用于运行时转译语言设置
  ```go
  ctx := gi18n.WithLanguage(context.TODO(), "zh-CN")
  i18n.Translate(ctx, `hello`)
  ```
  用于将转译语言设置到上下文变量中，并返回一个新的上下文变量，该变量可用于后续的转译方法。
  
#### 常用方法
- `T` 方法
  
  `T` 方法是 `Translate` 方法的别名，也是大多数时候我们推荐使用的方法名称。`T` 方法可以给定关键字名称，也可以直接给定模板内容，将会被自动转译并返回转译后的字符串内容。

  方法的定义：
  ```go
  // T translates <content> with configured language and returns the translated content.
  func T(ctx context.Context, content string)
  ```
  `T` 方法允许通过第二个参数指定目标语言进行翻译，需使用标准的国际化语言代码（如 en, ja, ru, zh-CN, zh-TW 等）。若未指定，则使用默认语言。
  
- `Tf` 方法

  `Tf` 为 `TranslateFormat` 的别名，该方法支持格式化转译内容，字符串格式化语法参考标准库 fmt 包的 `Sprintf` 方法。
  
  方法的定义：
  ```go
  // Tf translates, formats and returns the <format> with configured language
  // and given <values>.
  func Tf(ctx context.Context, format string, values ...interface{}) string
  ```
  
- 上下文设置转译语言

#### I18N与视图引擎
gi18n 默认已经集成到了 GoFrame 框架的视图引擎中，直接在模板文件/内容中使用 `gi18n` 的关键字标签即可。我们同样可以通过上下文变量的形式来设置当前请求的转译语言。

以下是如何在 GoFrame v2 中结合使用 i18n 和 template 的示例。完整代码可查看 [i18n-example](https://github.com/clin211/goframe-practice/tree/i18n-example)
1. 创建项目，我以最小闭环的方式来实现这个功能，所以我就不用 `gf cli` 来创建项目，直接创建一个 `i18n-template` 的目录，然后进入这个目录，然后初始化 mod：
  ```sh
  mkdir i18n-template

  cd i18n-template
  
  go mod init goframe-i18n
  ```

2. 项目结构
  ```tree
  .
  ├── i18n/
  │   ├── en.json
  │   └── zh.json
  ├── template/
  │   └── index.html
  ├── go.mod
  ├── go.sum
  └── main.go
  ```

3. 配置 i18n 文件

    在 i18n 目录下创建语言文件 en.json 和 zh.json：

    - en.json：

      ```json
      {
          "hello": "Hello, world!"
      }
      ```
    - zh.json：

      ```json
      {
          "hello": "你好，世界！"
      }
      ```

4. 创建模板文件

    在 template 目录下创建模板文件 `index.html`：
    ```html
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>i18n Example</title>
    </head>
    <body>
        <h1>{{ i18n "hello" }}</h1>
    </body>
    </html>
    ```
    
5. 编写 main.go

    在 `main.go` 中，加载 `i18n` 配置，并在模板中注册自定义函数 `i18n`：
    ```go
    package main
    
    import (
        "github.com/gogf/gf/v2/frame/g"
        "github.com/gogf/gf/v2/i18n/gi18n"
        "github.com/gogf/gf/v2/net/ghttp"
        "github.com/gogf/gf/v2/os/gctx"
        "github.com/gogf/gf/v2/os/gview"
    )
    
    func main() {
        ctx := gctx.New()
        s := g.Server()
    
        // 设置 i18n 配置
        i18n := gi18n.New(gi18n.Options{
            Path: "i18n",
        })
        i18n.SetLanguage("en") // 默认语言
    
        // 自定义函数
        funcMap := map[string]interface{}{
            "i18n": func(key string) string {
                return i18n.Translate(ctx, key)
            },
        }
    
        // 添加自定义模板引擎并注册函数
        view := gview.New()
        view.BindFuncMap(funcMap)
        s.SetView(view)
    
        // 注册路由
        s.BindHandler("/", func(r *ghttp.Request) {
            lang := r.Get("lang").String()
            if lang != "" {
                i18n.SetLanguage(lang)
            }
    
            r.Response.WriteTpl("index.html", nil)
        })
    
        s.SetPort(8199)
        // 启动服务器
        s.Run()
    }
    ```
    
6. 启动服务器

    在项目根目录下运行 `go run main.go` 命令启动服务器：

    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/85wJfc-958660be-d228-4290-a986-f77a176f7f22.png" /> 

7. 访问网页查看效果

    打开浏览器访问 `http://localhost:8199`，没有设置lang参数的时候，你将看到默认的语言内容:

    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/oZfikt-3d429dfb-cd82-4e60-848c-aca318f17a14.png" style="border-radius: 16px;border:1px solid rgba(145, 109, 213, 0.5);" />

    如果在访问的是 `http://localhost:8199/?lang=zh`，看到的内容则是中文:

    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/tkK8eQ-588de006-7b96-47e5-a9b6-4b3539da8a4a.png" style="border-radius: 16px;border:1px solid rgba(145, 109, 213, 0.5);" />

## 总结

本文详细讲解了 GoFrame 框架中的模板引擎，展示了大量内置函数和自定义函数的使用技巧，并通过实际示例进行演示。最后，我们结合 i18n 功能，演示了如何在路由处理中动态设置语言，实现多语言支持的动态网页开发。