1. 如何实现一个元素的水平垂直居中
    - 绝对定位方法

        ```css
        .box {
          position: absolute;
          top: 50%;
          bottom: 50%;
          width: 50px;
          height: 50px;
          transform: translate(-50%, -50%);
        }
        ```

    - flex
    - grid
  
2. 网站开发中，如何实现图片的懒加载
    - 位置计算 + 滚动事件 (Scroll) + DataSet API
    - getBoundingClientRect API + Scroll with Throttle + DataSet API
    - IntersectionObserver API + DataSet API
    - LazyLoading属性

3. 什么是防抖和节流，它们的应用场景有哪些?
    - 防抖：防止抖动，单位时间内事件触发会被重置，避免事件被误伤触发多次。代码实现重在清零 clearTimeout。防抖可以比作等电梯，只要有一个人进来，就需要再等一会儿。业务场景有避免登录按钮多次点击的重复提交。
    - 节流：控制流量，单位时间内事件只能触发一次，与服务器端的限流 (Rate Limit) 类似。代码实现重在开锁关锁 timer=timeout; timer=null。节流可以比作过红绿灯，每等一个红灯时间就可以过一批。

4. React 的组件间通信都有哪些形式？
   - 父传子：在 React 中，父组件调用子组件时可以将要传递给子组件的数据添加在子组件的属性中，在子组件中通过 props 属性进行接收。这个就是父组件向子组件通信。
   - 子传父：React 是单向数据流，数据永远只能自上向下进行传递。当子组件中有些数据需要向父级进行通信时，需要在父级中定义好回调，将回调传递给子组件，子组件调用父级传递过来的回调方法进行通信。
   - 跨组件通信 - context。使用 context API，可以在组件中向其子孙级组件进行信息传递。

5. 如何使用 react hooks 实现一个计数器的组件？

```js
const useCountDown = (num) => {
    const [seconds, setSecond] = useState(num)

    useEffect(() => {
        setTimeout(() => {
            if (seconds > 0) {
                setSecond(c => c - 1);
            }
        }, 1000);
    }, [seconds]);

    return [seconds, setSecond]
}
```

6. React 中如何实现路由懒加载？

- 使用 React.lazy 和 Suspense

7. React 的生命周期函数都有哪些，分别有什么作用？
React 的生命周期已经经历了 3 次改动，推荐查看《[深入学习 React：组件、状态与事件处理的完整指南](https://mp.weixin.qq.com/s/PFnnK9SoYURwuxMUwAlZqA)》，这篇文章中对 react 生命周期有详细的讲解！

8. 说一下 React Hooks 在平时开发中需要注意的问题和原因?

- 不要在循环，条件或嵌套函数中调用 Hook，必须始终在 React 函数的顶层使用 Hook
- 使用 useState 时候，使用 push，pop，splice 等直接更改数组对象的坑
- 不要滥用 useContext
- 善用 useCallback

9. useEffect 中如何使用 async/await?
    可以选择再包装一层 async 函数，置于 useEffect 的回调函数中，变相使用 `async/await`:

    ```jsx
    async function fetchMyAPI() {
      let response = await fetch('api/data')
      response = await res.json()
      dataSet(response)
    }

    useEffect(() => {
      fetchMyAPI();
    }, []);
    ```

10. React 中，`cloneElement` 与 `createElement` 各是什么，有什么区别?
    1. createElement：

      - 第一个参数是 type 简单来说就是各种 标签名字（包括 HTML 自身的，还有 React 组件名字）
      - 第二个参数是传入的属性
      - 第三个参数以及之后的参数就是作为组件的子组件 JSX 编写的代码就是转换为这个方法，一般用了 JSX 写法都不会再需要自己直接调用 该方法
    2. cloneElement

       - 第一个参数是 一个 React 元素
       - 新添加的属性会并入原有的属性 一般配合 `React.children.map` 使用，如用于动态地给子组件添加更多 `props` 信息、样式

    更深一点的原因在于，React 元素是 不可变对象 例如 `props.children` 获取到的只是一个 描述符，不能直接修改它的任何属性，只能读取他的信息。 所以我们可以选择拷贝它们，然后再修改、添加

11. 如何使用 react hooks 实现 useFetch 请求数据?
    可以参考 <https://www.robinwieruch.de/react-hooks-fetch-data/>

12. 有使用过vue吗？说说你对vue的理解?
    这个问题至少要包含以下几个方面：
    - 出现的背景
    - 什么是 Vue
    - MVVM
    - Virtual DOM
    - 组件化
    - 指令系统
    - 跟传统的开发的区别
    - 生态

13. `v-show` 和 `v-if` 有什么区别？使用场景分别是什么？
    `v-show` 总是会进行编译和渲染的工作 - 它只是简单的在元素上添加了 `display: none;` 的样式。`v-show` 具有较高的初始化性能成本上的消耗，但是使得转换状态变得很容易。

    相比之下，`v-if` 才是真正「有条件」的：它的加载是惰性的，因此，若它的初始条件是 `false`，它就不会做任何事情。这对于初始加载时间来说是有益的，当条件为 `true` 时，`v-if` 才会编译并渲染其内容。切换 `v-if` 下的块儿内容实际上时销毁了其内部的所有元素，比如说处于 `v-if` 下的组件实际上在切换状态时会被销毁并重新生成，因此，切换一个较大 `v-if` 块儿时会比`v-show` 消耗的性能多。

    如果需要非常频繁地切换，则使用 `v-show` 较好，如果在运行时条件很少改变，则使用 `v-if` 较好

14. `v-if` 和 `v-for` 的优先级是什么？
v-show 总是会进行编译和渲染的工作 - 它只是简单的在元素上添加了 display: none; 的样式。v-show 具有较高的初始化性能成本上的消耗，但是使得转换状态变得很容易。
相比之下，v-if 才是真正「有条件」的：它的加载是惰性的，因此，若它的初始条件是 false，它就不会做任何事情。这对于初始加载时间来说是有益的，当条件为 true 时，v-if 才会编译并渲染其内容。切换 v-if 下的块儿内容实际上时销毁了其内部的所有元素，比如说处于v-if下的组件实际上在切换状态时会被销毁并重新生成，因此，切换一个较大v-if块儿时会比v-show消耗的性能多。

1. vue3.0 中为什么要使用 Proxy，它相比以前的实现方式有什么改进？
    - 可以提高实例初始化启动速度，优化数据响应式系统，由全部监听改为惰性监听（lazy by default)。
    - 数据响应式系统全语言特性支持，添加数组索引修改监听，对象的属性增加和删除。

16. 你对SPA单页面的理解，它的优缺点分别是什么？如何实现SPA应用呢?
    单页面应用（SPA，Single Page Application）是一种通过动态更新网页内容而不重新加载整个页面来创建应用程序的技术。SPA 的核心思想是将所有用户操作都在一个单独的 HTML 页面内进行，使用 JavaScript 和 AJAX 请求动态加载和显示内容。
    **优点**:
    1. 更快的用户体验：

        - 流畅的交互：只加载一次 HTML、CSS 和 JavaScript，后续交互时只更新部分内容，减少页面重新加载的时间。
        - 即刻响应：用户操作后，不需要等待页面重新加载，应用可以即时响应用户的操作。
    2. 更好的用户体验和性能：

        - 无刷新体验：通过动态加载和渲染，提供无刷新体验，提升用户体验。
        - 减少网络请求：可以减少与服务器的全页请求，节省带宽。
    3. 前后端分离：

        - 开发效率：前端和后端可以并行开发，前端开发者可以专注于用户界面，而后端开发者处理数据和业务逻辑。
    4. 客户端路由：

        - 导航控制：可以使用客户端路由控制 URL，模拟传统多页应用的页面导航。
    5. 更好的离线体验：

        - 离线支持：通过服务工作者（Service Workers）等技术，可以提供离线支持和缓存策略。

    **缺点**:
    1. 首次加载速度较慢：

        - 初始化加载：由于需要加载所有的 JavaScript、CSS 资源，首屏加载时间可能较长。
        - 资源包体积：较大的 JavaScript 包可能导致首次加载性能问题。
    2. SEO 问题：

        - 搜索引擎索引：传统搜索引擎难以索引 JavaScript 渲染的内容，可能影响 SEO。虽然现代搜索引擎在这方面有所改进，但仍需注意。
    3. 浏览器历史管理复杂：

        - 前进和后退：管理浏览器的前进和后退按钮的行为需要额外处理。
    4. JavaScript 依赖：

        - 依赖 JavaScript：如果 JavaScript 代码出现问题，整个应用可能会崩溃。
    5. 内存管理：

        - 内存泄漏：长时间运行的 SPA 应用可能出现内存泄漏问题，因为页面未重新加载，内存不会被自动释放。

    如何实现SPA应用呢？可以选择目前比较热门的解决方案：Vue、React、Svelte 等方案进行开发项目

17. SPA首屏加载速度慢的怎么解决？
    1. 资源压缩和优化
    2. 动态导入和路由级别 code splitting
    3. 服务端渲染
    4. 移除不要的依赖和将非关键代码延迟加载
    5. 对资源预加载和预获取
    6. 减少重绘和重排，在需要时异步渲染内容，避免阻塞主线程
18. 如何实现一个简单的 Promise?
    1. 状态管理：

       - state 用于表示 Promise 的当前状态。
       - value 和 reason 分别存储成功和失败的值。
    2. resolve 和 reject 方法：

        resolve 和 reject 用于更改 Promise 的状态并通知注册的回调。
    3. then 方法：

       - then 用于注册成功和失败的回调。
       - 如果 Promise 已经完成或拒绝，立即调用相应的回调。
       - 如果 Promise 还在等待中，将回调函数存储在 `onResolvedCallbacks` 和 `onRejectedCallbacks` 中。
    4. resolvePromise 函数：

       - 用于处理 then 中返回的值。
       - 如果返回值是一个 Promise，则递归调用 then。
    5. catch 方法：

        作为 then 的简写形式，用于处理失败情况。

    可以微信公众号留言或者私信我，给你发源码，这个源码确实太长了
19. 在前端开发中，如何获取浏览器的唯一标识?
    - 绘制 canvas，获取 base64 的 dataurl，对 dataurl 这个字符串进行 md5 摘要计算，得到指纹信息
    - Fingerprintjs 是一种高级技术，通过收集浏览器和设备的各种信息（如屏幕分辨率、插件信息等）来生成唯一标识。可以结合 localStorage 和 sessionStorage 以及其他方法来增强唯一性。
20. js 中什么是 softbind，如何实现?
    softBind 允许你将一个函数绑定到特定的对象，但它与 `Function.prototype.bind()` 不同。如果绑定后的函数是作为其他对象的方法调用的，softBind 会尊重该调用对象，而不会强制使用绑定对象。

    ```js
    if (!Function.prototype.softBind) {
        Function.prototype.softBind = function(obj, ...rest) {
            const fn = this
            const bound = function(...args) {
                const o = !this || this === (window || global) ? obj : this
                return fn.apply(o, [...rest, ...args])
            }
            bound.prototype = Object.create(fn.prototype)
            return bound
        }
    }
    ```

21. http 状态码中 301，302和307有什么区别?
    - 301 当网站的资源位置永久变更时，可以使用 301。比如，网站域名从 <http://example.com> 迁移到 <http://new-example.com，应该使用> 301。
    - 302 当资源的地址只是临时变化时，使用 302。比如，某个页面的临时维护期间跳转到另一个页面，维护结束后又恢复到原来的 URL。
    - 307 与 302 类似，表示资源临时移动，但 307 确保请求方法不会改变，因此在处理非 GET 请求（如 POST）时，307 更加安全。
22. http2 与 http1.1 有什么改进?
    - 多路复用
    - 安全性
    - 高效的长连接
    - ...
23. 前端如何进行多分支部署?
    - 部署到不同的服务器或存储服务
    - 单服务器，docker + 多环境配置
    - 根据请求头中的自定义属性 + NGINX 的条件判断和反向代理功能将路由映射到不同的服务器或不同的文件目录
24. 你们的前端代码上线部署一次需要多长时间，需要人为干预吗?
    这个问题就需要根据实际的情况去回答，可以适当的往自动化上面去靠，前提是要结合实际
25. npm i 与 npm ci 的区别是什么?
    `npm install` 是最常用的命令，用于安装项目中的依赖包。如果项目中没有 node_modules 目录或新的依赖项被添加到 package.json 中，它会将依赖项安装到项目的 node_modules 目录。

    主要特点：
       - 灵活性：npm install 会基于 package.json 文件安装最新的依赖版本，除非 package-lock.json 存在，它会优先根据 package-lock.json 安装精确版本。
       - package-lock.json：如果在安装依赖的过程中有新依赖或者版本变化，npm install 会更新 package-lock.json 文件。
       - 新增依赖：运行 npm install <package> 可以新增一个依赖，并自动更新 package.json 和 package-lock.json。
       - 不存在 package-lock.json 时：会根据 package.json 中的版本范围来安装符合条件的最新依赖版本，生成新的 package-lock.json。

    适用场景：
       - 开发中使用频繁，特别是在你可能新增、更新或删除依赖时。
       - 如果你需要在开发过程中手动管理或调整依赖项的版本。

    `npm ci` 是用于持续集成（CI）场景的命令，专注于快速、确定性地安装依赖，确保开发环境和生产环境之间的一致性。

    主要特点：
       - 严格性：npm ci 只依赖于 package-lock.json 文件，不参考 package.json 的版本范围。它会直接使用 package-lock.json 中锁定的版本进行安装。
       - 更快的安装：相比 npm install，npm ci 更快，因为它跳过了一些依赖解析和版本匹配的步骤，直接安装 package-lock.json 中精确锁定的依赖。
       - 清空 node_modules：在安装之前，npm ci 会首先删除当前的 node_modules 文件夹，并进行全新的依赖安装，确保安装环境的干净整洁。
       - 不可更新锁文件：与 npm install 不同，npm ci 绝不会更新 package-lock.json。如果 package-lock.json 与 package.json 不匹配，安装会失败。

    适用场景：
       - 持续集成/持续部署（CI/CD）系统：在 CI/CD 流水线中，确保依赖的版本与开发时完全一致，并且每次安装都是干净的状态。
       - 确定性依赖管理：当你需要确保所有人安装的依赖版本完全一致时，如团队开发或构建生产环境。

26. 有没有用 npm 发布过 package，如何发布?
项目结构合理且包含必要的文件
    - package.json：包含包的元数据，如包名、版本号、描述等。
      - name：包名，必须唯一，遵循 npm 包命名规范。
      - version：版本号，必须遵循 SemVer 规范。
      - main：入口文件，指定包的主要文件路径。
      - keywords：关键词，有助于用户搜索你的包。
      - repository：项目的仓库地址，帮助其他开发者找到源码。
      - scripts：可以定义常用的脚本命令，如测试和构建。
    - README.md：项目说明文件，帮助用户理解你的包和如何使用它。
    - LICENSE：软件许可文件，决定别人如何合法使用你的包。
    - npm 账号登录
    - 发布包
    - 版本更新（版本号规范）
    - 取消发布
  
27. 在 Node 应用中如何利用多核心 CPU 的优势?
    - Cluster: Node.js 是单线程的，但是可以通过 Cluster 模块来创建多个子进程，每个进程可以共享同一个端口，从而实现多核心的负载均衡。
    - Worker Threads: 对于需要在应用程序中执行 CPU 密集型任务时，可以使用 Worker Threads 模块。它允许在不同的线程中运行 JavaScript 代码，从而避免阻塞主线程。
    - Child Processes：使用 child_process 模块可以启动外部的命令或者是创建新的 Node.js 子进程，适合那些需要大量并行任务的应用场景。

28. node 中如何查看函数异步调用栈?
    - Error().stack
    - async_hooks
    - 利用 v8 的 asyncStackTraces
    - longjohn

29. Code Splitting 的原理是什么?
    - 动态导入 (import())： 使用动态导入将模块懒加载，Webpack 会自动创建一个独立的代码块并按需加载。
    - 多入口文件： 可以为 Webpack 配置多个入口（entry），每个入口都会打包成一个独立的 bundle。
    - SplitChunksPlugin： Webpack 内置了 SplitChunksPlugin，它可以自动将重复的依赖或较大的第三方库（如 react, lodash）提取到单独的文件中，避免在多个 bundle 中重复加载相同的代码。

30. 你使用过哪些前端性能分析工具?

    最常见且实用的性能工具有两个：
    - lighthouse: 可在 chrome devtools 直接使用，根据个人设备及网络对目标网站进行分析，并提供各种建议
    - webpagetest: 分布式的性能分析工具，可在全球多个区域的服务器资源为你的网站进行分析，并生成相应的报告
