在上一篇文章《学会 Gin 中间件，这些操作你也能随心所欲！》中，我们深入探讨了 Gin 中间件的原理与实战，了解了如何通过中间件增强应用的灵活性与可扩展性。今天来看一个新东西 Gin 的模板引擎。

在前后端分离的开发模式兴起之前，模板引擎（如 ASP、JSP、PHP 等等）曾是 Web 开发中的核心角色。

传统的服务端渲染（SSR, Server-Side Rendering）模式中，后端框架不仅负责业务逻辑处理，还直接渲染 HTML 页面并返回给客户端。通过模板引擎，开发者可以在 HTML 文件中嵌入动态数据，实现页面内容的动态渲染。例如，根据用户身份展示不同的页面内容，或者通过循环渲染商品列表。这种方式极大简化了页面渲染的复杂度，同时减少了前后端数据交互的成本。

在前后端分离逐渐成为主流之前，模板引擎在以下场景中尤为常见：

- 内容管理系统（CMS）：文章页面动态渲染。
- 电商平台：商品详情页与订单确认页。
- 管理后台：数据表格与报表生成。

尽管如今前后端分离模式让前端承担了更多的渲染任务，但在需要 SEO 优化或快速生成静态页面时，服务端渲染依然具备不可替代的优势。而在 Web 开发的世界里，页面渲染与静态资源管理同样至关重要。无论是高效的模板引擎，还是优化的 CSS、JS 文件加载，Gin 都提供了强大的支持。  

## Gin 模板引擎

在 Gin 框架中，默认支持的是 Go 标准库的 `html/template` 模板引擎，但同时也可以通过适配其他模板引擎来满足不同的需求。

### 原生模板引擎

Gin 框架默认集成了 Go 标准库中的 `html/template` 模板引擎，具备强大的功能，并且天然支持 HTML 转义，能有效防止 XSS 攻击。

### 核心特性

- 自动 HTML 转义：原生防御 XSS 攻击，安全渲染动态内容
- 模板继承：支持 `{{ define }}` 和 `{{ template }}` 实现布局复用
- 严格上下文检查：变量访问错误在编译期暴露

### 加载模板文件

使用 `LoadHTMLGlob()` 或 `LoadHTMLFiles()` 加载模板文件。

1. 全局加载目录下的所有模板文件

    ```go
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    // router.LoadHTMLGlob("templates/**/*.html")  // 加载多级目录模板
    ```

2. 加载指定的模板文件

    ```go
    router.LoadHTMLFiles("templates/index.html", "templates/detail.html")
    ```

### 渲染模板

通过 `c.HTML()` 渲染模板，并传入模板数据。
下面示例的目录结构：

```tree
.
├── go.mod
├── go.sum
├── main.go
└── templates
    └── index.html
```

`main.go` 文件的代码如下：

```go
package main

import "github.com/gin-gonic/gin"

func main() {
 router := gin.Default()
 router.LoadHTMLGlob("templates/*")
 // router.LoadHTMLGlob("templates/**/*.html")  // 加载多级目录模板

 router.GET("/", func(c *gin.Context) {
  // 文件路径是基于加载模板文件下的路径
  c.HTML(200, "index.html", gin.H{
   "title":    "Gin Template Demo",
   "username": "长林啊",
  })
 })

 router.Run(":8080")
}
```

`index.html` 文件的内容：

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.title}}</title>
</head>

<body>
    <h2>this is home page</h2>
    <h4>username: {{.username}}</h4>
</body>

</html>
```

> 忽略具体语法，先看效果，语法后面再详细介绍！

启动服务后，在浏览访问 `http://localhost:8080/` 效果如下：
![](https://files.mdnice.com/user/8213/f18a0fbd-e873-4599-9a1c-5317289e655f.png)

## 模板语法

以下是一些常用的模板语法，帮助你快速掌握如何在 Gin 中使用模板渲染。

### 自定义分隔符

在 Gin 的模板引擎 `html/template` 中，你可以自定义模板分隔符，避免与 Vue、JS 等前端框架的 `{}` 语法冲突。例如，可以将 `{{` 和 `}}` 替换为 `[[` 和 `]]`。使用 `SetDelims()` 方法来自定义左右分隔符：

```go
// 设置模板解析
tmpl := template.Must(template.New("").Delims("[[", "]]").ParseGlob("templates/**/*"))
router.SetHTMLTemplate(tmpl)

// 渲染页面
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "Title": "自定义模板分隔符",
        "Msg":   "Gin 模板支持自定义分隔符",
    })
})
```

在 HTML 模板中使用新分隔符，默认的 `{{` 和 `}}` 变成 `[[` 和 `]]`：

```html
<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <title>[[ .Title ]]</title>
</head>
<body>
    <h1>[[ .Msg ]]</h1>
</body>
</html>
```

- 避免与 Vue (`{{ }}`) 语法冲突。
- 适用于 JavaScript 代码块中包含 `{}` 的情况。

### 注释

模板中的注释不会被渲染到页面上。

```html
{{/* 这是单行注释 */}}

{{/*
多行注释
不会被渲染到输出
*/}}
```

### 变量插值

在模板中，使用 `{{ .VariableName }}` 来插入动态变量。

在上面的 HTML 中，`{{.title}}` 和 `{{.username}}` 就是变量插值：

```html
<title>{{.title}}</title>
<h4>username: {{.username}}</h4>
```

最后渲染出来，就是:

```html
<title>Gin Template Demo</title>
<h4>username: 长林啊</h4>
```

### 条件判断—— `if`

`if` 用于判断条件，如果条件为真，则执行 `if` 块中的内容。

```go
router.GET("/", func(c *gin.Context) {
    // 文件路径是基于加载模板文件下的路径
    c.HTML(200, "index.html", gin.H{
        "title":     "Gin Template Demo",
        "username":  "长林啊",
        "isAdmin":   false,
        "isLoginIn": true,
    })
})
```

在 index.html 中使用：

```html
{{ if .IsAdmin }}
<button>管理面板</button>
{{ else if .IsSignin }}
<button>个人中心</button>
{{ else }}
<button>登录</button>
{{ end }}
```

渲染效果：

![](https://files.mdnice.com/user/8213/7b89f8e9-5dac-4f7f-b791-c9b068bbeed0.png)

### 循环遍历——`range`

使用 `range` 遍历数组或切片。可以获取元素和索引。

```go
router.GET("/", func(c *gin.Context) {
  // 文件路径是基于加载模板文件下的路径
  c.HTML(200, "index.html", gin.H{
    // ...
   "skills":    []string{"Go", "Gin", "MySQL", "Redis"},
  })
 })
```

index.html 中的内容：

```html
{{ if .skills }}
<h3>skills:</h3>
<ul>
    {{ range $index, $item := .skills }}
    <li class="item-{{ $index }}">{{ $item }}</li>
    {{ end }}
</ul>
{{ else }}
<p>暂无技能</p>
{{ end }}
```

效果如下：

![](https://files.mdnice.com/user/8213/93414e9b-eca1-44ad-b13f-8575498e1f19.png)

在上面的例子中，`.skills` 是一个字符串数组，模板会遍历并输出每个元素。如果没有项，则输出“暂无技能”。

> 遍历`map`、字符串和切片也是一样的；这里就不一一演示了！

### 变量声明

必须使用 `$` 为前缀来声明变量，作用域仍然是块级作用域，若在 `range/with` 等块内声明，仅限块内使用。

```html
<!-- 变量声明 -->
{{$length := len .skills }}
<p>技能总数: {{ $length }}</p>
```

除了使用 `$` 为前缀外，其它的跟 Go 语言中使用没啥区别，比如：

```go
// 多变量并行赋值
{{ $name, $age := "张三", 25 }}
// 变量重声明
{{ $counter := 0 }}
{{ $counter = add $counter 1 }}  <!-- 修改值而非重新声明 -->
```

### 函数调用

Go 的 `html/template` 允许注册自定义函数，并在模板中调用。在 Gin 中注册自定义函数，**注册自定义函数必须要在加载模板之前**：

```go
// 注册函数 必须要在 加载模板之前！！！
router.SetFuncMap(template.FuncMap{
    "formatDate": func(t time.Time) string {
        return t.Format("2006-01-02 15:04:05")
    },
})

router.LoadHTMLGlob("templates/*")
// router.LoadHTMLGlob("templates/**/*.html")  // 加载多级目录模板

router.GET("/", func(c *gin.Context) {
    // 文件路径是基于加载模板文件下的路径
    c.HTML(200, "index.html", gin.H{
        "title":     "Gin Template Demo",
        "username":  "长林啊",
        "isAdmin":   false,
        "isLoginIn": true,
        "skills":    []string{"Go", "Gin", "MySQL", "Redis"},
        "date":      time.Now(),
    })
})
```

在 index.html 中使用：

```html
<time>{{formatDate .date}}</time>
```

效果如下：

![](https://files.mdnice.com/user/8213/8821eeb4-ddcf-4e5b-b430-eab299b8c84c.png)

### 转义与安全输出

`html/template` 默认会对输出的内容进行 HTML 转义，这样可以防止 XSS 攻击。如果你需要输出原始 HTML，可以使用 `safeHTML` 函数。

- **默认输出会进行 HTML 转义**

  ```html
  <p>{{ .content }}</p>
  ```

  如果 `content` 包含 `<script>alert('XSS')</script>`，它会被安全地转义为：

  ```html
  &lt;script&gt;alert('XSS')&lt;/script&gt;
  ```
  
- 如果需要输出原始 HTML 使用 `safeHTML` 函数：

  ```html
  <p>{{ .content | safeHTML }}</p>
  ```
  
### 管道符

管道符用于将一个值传递给一个函数，并可以链式调用多个函数；从右向左依次执行。

```html
<p>{{ .name | upper }}</p>
```

`upper` 是一个自定义函数，它将字符串转换为大写。如果 `name` 是 "hello"，则渲染结果为：

```html
<p>HELLO</p>
```

### 字符串函数

Go 的 `html/template` 支持多种字符串处理函数，常见的有：

- `upper` - 转换为大写
- `lower` - 转换为小写
- `title` - 转换为标题格式（每个单词的首字母大写）
- `trim` - 去除前后空格
- `replace` - 替换字符串

### 多行文本

在模板中，如果你想插入多行文本，可以使用 `-` 来去掉换行和额外的空格。

```html
{{- /* No extra space or newlines here */ -}}
<p>{{ .content }}</p>
```

### 数学运算

Go 模板引擎不直接支持数学运算，但你可以通过自定义函数来实现：
> 在前端生态中有一个很出名的 [lodash](https://lodash.com/) 库，这是一个功能完备的工具库，Go 语言中与之相匹配的也有一个，那就是 [samber/lo](https://github.com/samber/lo) 库：<https://github.com/samber/lo>

- `lo.Sum` 可以快速对数组或切片求和。
- `lo.Max` 和 `lo.Min` 可以获取最大值和最小值。
- `lo.Contains` 用于判断字符串切片是否包含某个元素。

```go
package main

import (
 "html/template"
 "time"

 "github.com/gin-gonic/gin"
 "github.com/samber/lo"
)

func main() {
 router := gin.Default()

 // 注册函数
 router.SetFuncMap(template.FuncMap{
  "formatDate": func(t time.Time) string {
   return t.Format("2006-01-02 15:04:05")
  },
  "add":      lo.Sum[int],         // 求和
  "max":      lo.Max[int],         // 最大值
  "min":      lo.Min[int],         // 最小值
  "contains": lo.Contains[string], // 判断是否包含
 })

 router.LoadHTMLGlob("templates/*")
 // router.LoadHTMLGlob("templates/**/*.html")  // 加载多级目录模板

 router.GET("/", func(c *gin.Context) {
  // 文件路径是基于加载模板文件下的路径
  c.HTML(200, "index.html", gin.H{
   "title":     "Gin Template Demo",
   "username":  "长林啊",
   "isAdmin":   false,
   "isLoginIn": true,
   "skills":    []string{"Go", "Gin", "MySQL", "Redis"},
   "date":      time.Now(),
   "content":   "<script>alert('XSS')</script>",
   "numbers":   []int{3, 5, 7, 2},
   "word":      "hello",
   "words":     []string{"hello", "world", "gin"},
  })
 })

 router.Run(":8080")
}
```

比如：

![](https://files.mdnice.com/user/8213/d555aa58-3730-48c9-9956-1a66969e3dc5.png)
通过这种方式，你可以在 Gin 模板中轻松使用 `samber/lo` 强大的功能，大大提升了模板的表达能力。

### 模板继承与嵌套

模板继承允许在一个基础模板中定义共享内容，子模板可以复用这些内容。templates 下结构如下：

```tree
.
├── index.html // 这个是前面示例的内容
├── layout
│   ├── content.html
│   ├── footer.html
│   └── header.html
└── pages
    └── detail.html
```

下面的示例的入口文件是 `/pages/detail.html` 文件，首先来定义 layout 下的模板：

- `header.html` 文件的内容：

  ```html
  {{ define "header.html" }}
  <header>
      <h1>这是页面头部</h1>
      <nav>
          <a href="/">首页</a> |
          <a href="/about">关于</a> |
          <a href="/contact">联系我们</a>
      </nav>
  </header>
  {{ end }}
  ```

- `footer.html` 文件的内容：

  ```html
  {{ define "footer.html" }}
  <footer>
      <p>© 2025 Gin Demo. All rights reserved.</p>
  </footer>
  {{ end }}
  ```

- `content.html` 文件的内容：

  ```html
  {{ define "header.html" }}
  <header>
      <h1>这是页面头部</h1>
      <nav>
          <a href="/">首页</a> |
          <a href="/about">关于</a> |
          <a href="/contact">联系我们</a>
      </nav>
  </header>
  {{ end }}
  ```

接着就是在 `pages/detail.html` 中来整合这三个模板：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>

<body>
    <p>detail</p>
    {{ template "header.html" . }}
    <div class="content">
        {{ template "content.html" . }}
    </div>
    {{ template "footer.html" . }}
</body>
</html>
```

注册路由的核心代码：

```go
// 路由
router.GET("/detail", func(c *gin.Context) {
    c.HTML(http.StatusOK, "detail.html", gin.H{
        "Title": "Gin 模板继承",
    })
})
```

在浏览器中访问 `http://localhost:8080/detail` 效果如下：

![](https://files.mdnice.com/user/8213/70d66d1c-d7c8-4cd6-aee0-88ecdfd2eb03.png)

## 实战演练

上面演示了模板的嵌套使用，下面来看看怎么在模板中引入静态资源！我们以 `axios` 作为请求库、`tailwindcss` 作为样式库(这两文件都放在 static 目录下)，来获取一个产品列表作为示例，文件目录如下：

```tree
.
├── go.mod
├── go.sum
├── main.go
├── static
│   └── js
│       ├── axios.min.js
│       └── tailwindcssv4.js
├── templates
│   ├── index.html
│   ├── layout
│   │   ├── content.html
│   │   ├── footer.html
│   │   └── header.html
│   └── pages
│       ├── detail.html
│       └── products.html
```

最后的效果如下：

- 列表页
![](https://files.mdnice.com/user/8213/b3536490-a49b-420b-99eb-3e538ad6dc27.jpg)
- 点击产品后的效果：
![](https://files.mdnice.com/user/8213/8f84e166-4f3c-41de-a3a9-866ab1f0e5da.jpg)

知道最终要做一个什么东西后，我们先理一下具体实施步骤：

1. 下载 axios 和 tailwindcss 的静态资源。
2. 在 main.go 中注册静态资源。
3. 在 pages 中添加 products.html 文件。
4. 在 main.go 中添加 products 路由。
5. 做页面效果的 UI 开发（UI 是我通过 cursor 生成的）。
6. 使用 axios 请求产品信息，接口地址：[https://fakestoreapi.com/products](https://fakestoreapi.com/products)；并做分页相关逻辑。
7. 在浏览器中访问 `http://localhost:8080/products`。

因为篇幅原因，这里就不一一贴代码演示了，就将第二步和第四步的核心代码展示出来：

```go
// other code...
router.Static("/static", "./static")

// other code...
router.GET("/products", func(c *gin.Context) {
    c.HTML(http.StatusOK, "products.html", nil)
})
```

可以在 [https://github.com/clin211/gin-learn/tree/main/05-template](https://github.com/clin211/gin-learn/tree/main/05-template) 中获取所有源码！

## 总结

Gin 的模板引擎基于 Go 的 `html/template`，以安全渲染为核心，通过自动 HTML 转义防御 XSS 攻击，同时支持 `safeHTML` 可控原始输出；

提供变量插值、条件判断、循环遍历等动态语法，结合自定义函数与模板继承机制（如 `define`/`template`），可高效复用布局与组件，简化页面开发；

尤其适用于**服务端渲染（SSR）**场景及混合开发模式，实现动态内容与静态资源的高效协同。
