在本文中，我将讨论面试中关于 React 概念的最常见问题及其简短答案。

## JavaScript 概念和问题

### JS 中的原始类型和非原始类型：
原始数据类型包括数字、字符串、布尔值、NULL、无穷大和符号。非原始数据类型是对象。JavaScript 数组和函数也是对象。点击此处了解更多信息

### Javascript 中的浅拷贝和深拷贝是什么？
浅拷贝用于“平面”对象，深拷贝用于“嵌套”对象。

### 什么是扁平化？
“扁平”对象是指仅包含原始值的对象。例如：[1, 2, 3, 4, 5]

### 什么是嵌套对象？
嵌套对象是指包含非原始值的对象。例如：[“laptop”, {value: 5000}]

### 做浅拷贝的方法有哪些？
要创建浅拷贝，我们可以使用以下方法：
- 扩展语法 `[...arr], {...obj}`
- `Object.assign()`
- `Object.from()`
- `Object.create()`
- `Array.prototype.concat()`

### 如何在 JS 中进行深复制？
要创建深层副本，我们可以使用：
- `JSON.parse(JSON.stringify())`
- `structuredClone()`
- 第三方库，例如 [Lodash](https://lodash.com/)、[radash](https://radash-docs.vercel.app/docs/getting-started)
有关浅拷贝和深拷贝的[更多详细信息，请参阅此博客](https://makimo.com/blog/shallow-and-deep-copies-in-javascript/#:~:text=We%20have%20two%20kinds%20of,that%20contain%20only%20primitive%20values.&text=Nested%20objects%20mean%20objects%20that%20contain%20non%2Dprimitive%20values.)

### 什么是对象原型？
原型是 JavaScript 对象相互继承特性的机制。请查看 [MDN 的官方文档](https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/Objects/Object_prototypes)

### “map”、“filter”和“reduce”之间有什么区别？
`Map`、`Filter` 和 `Reduce` 是数组方法，可帮助以各种方式创建新数组。
- `Map` 返回原始数组中的信息片段数组。在回调函数中，返回您希望成为新数组一部分的数据。

- `Filter`：根据自定义条件返回原始数组的子集。在回调函数中，返回一个布尔值来确定每个项目是否将包含在新数组中。

- `Reduce`：可用于返回几乎任何内容。它通常用于返回单个数字，例如数组的总和

### 什么是“高阶函数”？
它们是可以将其他函数作为参数或将其作为结果返回的函数。（顺便说一下，Map、Reduce 和 Filter 是“高阶函数”

### call()、apply() 和 bind() 之间有什么区别？
`call`、`bind`和`apply` 三种方法都将参数 设置 `this` 为函数。

- `call`：绑定this值，调用函数，并允许您传递参数列表。

- `apply`：绑定this值，调用函数，并允许您将参数作为数组传递。

- `bind`：绑定this值，返回新函数，并允许您传入参数列表。

### 扩展运算符和剩余运算符有什么区别？
JavaScript 中的扩展运算符将数组和字符串中的值扩展为单个元素，而 rest 运算符将用户指定的数据的值放入 JavaScript 数组中。例如，请点击[此处](https://www.freecodecamp.org/news/javascript-rest-vs-spread-operators/)

## React 概念
### react 中的元素 (Element) 和组件 (Component) 有什么区别?
在 React 中，元素是虚拟 DOM 中显示内容的最小表示。而组件是定义特定 UI 或功能的可重用且独立的代码单元。组件可以返回 React 元素

[文章](https://www.geeksforgeeks.org/what-is-the-difference-between-element-and-component/)

### ReactJS 中的 Inceptors 是什么？
在 Web 开发中，拦截器可以指拦截和处理请求或响应的中间件函数或方法。Axios 等库提供了请求和响应拦截器，可以灵活地处理 HTTP 请求

### 什么是“refs”和“forwardRef”？
React 中的 Refs 提供了一种访问在 render 方法中创建的 DOM 节点或 React 元素的方法

forwardRef 是一个允许 React 组件将 ref 传递给子组件的函数。当你需要从父组件访问子组件的实例或 DOM 节点时，它很有用。

### 什么是 React Fiber？

Fiber Reconciler 是从 React 16 开始引入的，是 React 中的新协调算法。术语 Fiber 指的是 React 的数据结构（或）架构，源自“fiber”——DOM 树节点的表示。

### 什么是 HOC (高阶组件)？
高阶组件是 React 中的一种模式，其中函数接受一个组件并返回具有附加功能的新组件。它是一种重用组件逻辑、在组件之间共享代码以及应用横切关注点的方法。

### 什么是延迟加载？
延迟加载是一种延迟加载模块或资源直到需要时才加载的技术。它可以帮助减少初始包大小并提高应用程序的性能。在 React 上下文中，延迟加载通常与React.lazy异步加载组件的功能一起使用。

### 什么是虚拟 DOM？
虚拟 DOM 是 React 中的一个概念，它维护实际 DOM 的内存表示。当应用程序发生变化时，React 首先更新虚拟 DOM，然后计算更新实际 DOM 的最有效方法，最后应用这些更改

### 什么是 React Fragment?
React Fragment 是 React 中的一种轻量级语法，它允许您对多个元素进行分组，而无需向 DOM 添加额外的节点。当您需要从组件的 render 方法返回多个元素但又不想引入额外的包装 div 或任何其他元素时，它特别有用。

<>在 React 中，Fragments 可以使用和</>语法或来表示<React.Fragment>。

### React 中的上下文是什么?
Context 是 React 中的一种数据共享机制。你可以问这个问题来评估应聘者对组件如何传递和共享数据的了解程度。应聘者应该能够解释什么是 context 以及它们的理想用例

### 什么是道具钻孔？
当你必须通过多层组件传递数据时，即使某些中间组件不需要这些数据，也会发生 Props 钻取。它会使代码更难维护和理解。

### 如何解决 React 中的 Props Drilling?
解决 props 钻研的一个方法是使用React Context或 Redux 等状态管理库

### 什么是错误边界？
错误边界是 React 组件，可在其子组件树中的任何位置捕获 JavaScript 错误。然后，错误边界组件可以记录这些错误并显示后备 UI，而不会导致整个组件树崩溃。您可以将错误边界视为组件的捕获块。