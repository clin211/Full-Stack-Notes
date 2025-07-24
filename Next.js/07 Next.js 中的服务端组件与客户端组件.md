在上一篇文章《[掌握 Next.js 渲染机制：如何在 CSR、SSR、SSG 和 ISR 中做出最佳选择](https://mp.weixin.qq.com/s/V3FadXse_MIXOZqSBaHAMg)》中，我们深入探讨了 Next.js 提供的多种渲染机制，包括客户端渲染（CSR）、服务端渲染（SSR）、静态站点生成（SSG）和增量静态再生（ISR），并分析了它们在不同场景下的优势和适用性。这些渲染机制为开发者提供了强大的工具，以满足现代 Web 应用对于性能、SEO 和用户体验的高标准要求。

继上次深入剖析 Next.js 的渲染机制之后，我们将继续探索 Next.js 的另一个核心概念：**服务端组件与客户端组件**。在构建 Web 应用时，选择合适的渲染策略是至关重要的，而组件的类型选择则是实现这些策略的基础。服务端组件和客户端组件在渲染机制中扮演着不同的角色，它们直接影响到应用的加载速度、SEO 表现以及开发体验。

接下来，我们进一步深入探讨服务端组件与客户端组件的工作原理，将探讨以下几个核心问题：

- 服务端组件与客户端组件的定义和区别是什么？
- 它们各自有哪些优势和适用场景？
- 如何在 Next.js 中合理的应用服务端组件和客户端组件？

## 一、React Server Components
2020 年 12 月 21 日，React 官方发布了一篇介绍 [React Server Components](https://react.dev/blog/2020/12/21/data-fetching-with-react-server-components) 的文章，同时由 React 团队工程师 Dan Abramov 和 Lauren Tan 主讲，分享了一段约 1 小时的演讲。演讲中详细介绍了 React Server Components 的背景以及其使用方式。
### 什么是 React Server Components
React Server Components (RSC) 是 React 引入的一种新概念，旨在通过在服务端渲染组件来提升性能并优化开发体验。与传统的客户端渲染或服务端渲染（Server-Side Rendering, SSR）不同，React Server Components 的核心理念是：

- 在服务器端执行组件的逻辑，并将生成的 UI 片段（HTML）直接发送给客户端。
- 减少客户端 JavaScript 的数量，将渲染的工作移到服务器端以提升性能。

#### 特点
1. 无需发送 JavaScript 到客户端：Server Components 在服务器端完全处理逻辑，不需要在客户端运行，因此不生成额外的 JavaScript。
2. 高效的数据获取：Server Components 可以直接与数据库、API 或文件系统交互，无需通过客户端发送请求，从而减少网络延迟。
3. 流式渲染 (Streaming Rendering)：允许服务器将组件生成的部分内容逐步发送到客户端，用户可以在数据加载完成之前看到页面的部分内容。
4. 与客户端组件协作：RSC 可以与传统的 React 客户端组件（Client Components）无缝协作，混合使用来满足不同场景的需求。

#### 与传统渲染方式的区别
| 特性                | React Server Components | 客户端渲染 (CSR)                   | 服务端渲染 (SSR)           |
| ------------------- | ----------------------- | ---------------------------------- | -------------------------- |
| **运行环境**        | 服务端                  | 浏览器                             | 服务端                     |
| **是否有运行时 JS** | 无运行时 JavaScript     | 需要完整的 React JavaScript 运行时 | 有一定运行时 JavaScript    |
| **适用场景**        | 静态内容或无交互的页面  | 动态页面或复杂交互                 | SEO 优化或初次渲染性能优化 |
| **数据获取方式**    | 直接从服务端获取数据    | 通过浏览器请求 API                 | 通过服务端请求 API         |

#### 使用场景
- **静态内容**： 不需要交互的页面，例如展示产品列表、博客文章或营销页面。
- **高性能需求**： 需要减少客户端 JS 负载的应用，例如大型电商网站或高流量博客。
- **复杂数据逻辑**： 数据需要在服务器端整合的场景，例如多数据源的聚合展示。

#### 优势
- 减少客户端开销：客户端不需要加载和执行与 RSC 相关的 JavaScript，显著提升性能。
- 更快的首屏渲染：通过流式渲染，用户可以更快地看到页面内容，即使后续部分内容仍在加载。
- 开发体验优化：直接在服务端编写逻辑，可以省去前后端切换的复杂性，例如不需要手动管理 API 请求。

#### 局限性
- 不支持客户端特有的功能：例如无法访问 `window` 或 `document`，这些功能需要交由客户端组件处理。
- 学习曲线：与传统的 React 应用开发相比，需要学习如何在组件之间合理分工。
- 环境依赖：需要 Node.js 支持，比如与 Next.js 等框架结合使用。

## 二、Next.js 中的服务端组件
在上一篇文章中，我们对服务端组件有了详细的介绍，这里就不赘述了，如果有不了解的，可以翻阅文章《[掌握 Next.js 渲染机制：如何在 CSR、SSR、SSG 和 ISR 中做出最佳选择](https://mp.weixin.qq.com/s/V3FadXse_MIXOZqSBaHAMg)》

## 三、React Server Component VS Server-side Rendering
表面上看，RSC 和 SSR 非常相似，都发生在服务端，都涉及到渲染，目的都是更快的呈现内容。但实际上，这两个技术概念是相互独立的。


React Server Components (RSC) VS Server-side Rendering (SSR)

| 特性                     | React Server Components (RSC)                   | Server-side Rendering (SSR)                   |
| ------------------------ | ----------------------------------------------- | --------------------------------------------- |
| **渲染阶段**             | 服务端负责渲染 UI 内容，无需生成客户端运行时 JS | 服务端生成完整的 HTML，包含 React 的运行时 JS |
| **运行时依赖**           | 无运行时 JavaScript，仅生成 HTML                | 包含 React 运行时，支持客户端动态交互         |
| **交互逻辑**             | 无法处理客户端的动态逻辑和交互                  | 支持完整的 React 生命周期和交互功能           |
| **适用场景**             | 静态内容、不需要动态交互的页面                  | 需要 SEO 支持，同时具备动态交互的页面         |
| **性能**                 | 更轻量，减少客户端 JS 体积，提升首屏渲染速度    | 相对较重，需要传输运行时 JS                   |
| **数据获取方式**         | 数据直接从服务端获取，并与渲染逻辑紧密结合      | 数据在服务端获取后生成 HTML 返回客户端        |
| **流式渲染 (Streaming)** | 支持，允许分段传输内容到客户端，提升用户体验    | 通常需要等待完整页面生成后返回                |
| **复杂性**               | 需要明确区分服务端组件和客户端组件              | 开发逻辑与传统 React 类似，学习成本较低       |
| **浏览器 API 支持**      | 不支持（如 `window`、`document`）               | 支持，运行在客户端时可以访问浏览器 API        |
| **开发工具支持**         | 需要框架（如 Next.js）的支持                    | 广泛支持（如 Next.js、Nuxt 等）               |

---

### 区别
- **React Server Components** 更适合静态内容展示、减少客户端负载，并提升首屏渲染性能；但不支持交互逻辑，需与客户端组件配合使用。
- **Server-side Rendering** 提供完整的 React 功能，适合既需要 SEO 又需要客户端动态交互的场景，但在性能上可能稍逊于 RSC。

结合实际应用需求，可以灵活选择 RSC 或 SSR，或者将两者结合以达到最佳效果。

## 四、Next.js 中的客户端组件
在 Next.js 中，客户端组件（Client Components）是为处理动态交互和运行时 JavaScript 提供的组件类型，与服务端组件（Server Components）配合使用，可以灵活实现现代 Web 应用的需求。
### 客户端组件的概念
客户端组件通过文件顶部的 `'use client';` 声明来标识。
#### 定义什么是客户端组件
客户端组件在浏览器中运行，包含 React 的运行时 JavaScript，支持 React 的所有功能（如 `useState`、`useEffect` 等 Hook）。这些组件可以处理**动态交互**和**浏览器特有的功能**（如 DOM 操作、事件处理、地理位置和 localStorage 等等）。

#### 在 Next.js 中使用客户端组件
```jsx
'use client'; // 声明为客户端组件

import { useState } from 'react';

export default function Counter () {
    const [count, setCount] = useState(0);

    return (
        <div>
            <p>Current Count: {count}</p>
            <button onClick={() => setCount(count + 1)}>Increment</button>
        </div>
    );
}

```
### 客户端组件的特点
- 运行时依赖：必须在浏览器中执行，依赖 React 的运行时。
- 支持浏览器 API：可以使用 `window`、`document` 等浏览器特定的功能。
- 动态交互能力：允许使用状态管理（如 `useState`）和副作用（如 `useEffect`）来处理用户交互。

在 App Router 中，默认所有组件都是 Server 组件；`"use client";` 用于声明服务端和客户端组件的边界，定义后，文件及其导入的模块（包括子组件）都会成为客户端 bundle 的一部分。
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/djUzPf-465bfe49-7d4a-4071-aabf-55535705ff3f.png)
上图显示，如果在嵌套组件（如 toggle.js)中使用 `onClick` 和 `useState`，但未定义 `"use client";` 指令，会导致错误。解决办法就是在 toggle.js 文件的顶部声明 `"use client";`。


## 五、Next.js 中的服务端组件与客户端组件

### 如何选择

#### 服务端场景
- **SEO 优化**：需要通过服务端渲染以提高搜索引擎优化（SEO）效果。
- **静态内容渲染**：页面内容基本不需要交互或客户端状态管理，例如展示产品详情、博客文章等。
- **数据预加载**：页面数据可以在服务端加载，减少客户端的计算量和加载时间。
- **性能优化**：适合减少客户端负担，减少 JavaScript 执行，特别是在低性能设备上。
> - 在服务端渲染，生成完整的 HTML 后发送到客户端。
> - 不支持客户端状态管理，如 `useState` 和 `useEffect`。
> - 有利于首次渲染的性能，适用于静态内容和展示。
#### 客户端场景
- **动态交互**：需要用户输入、事件处理或界面更新的页面，如表单提交、动态内容显示等。
- **客户端状态管理**：需要在客户端存储和管理状态，常使用 React 的 `useState` 和 `useEffect`。
- **依赖浏览器特性**：例如，使用浏览器 API（如 `localStorage`、`window` 等）来执行任务。
- **需要局部更新的组件**：只需要渲染和更新特定部分，而不是整个页面。

> - 在客户端渲染，允许更复杂的交互。
> - 支持客户端状态管理，使用 React Hook（`useState`, `useEffect` 等）。
> - 增加初始加载负担，但提供更灵活的用户交互。

### 渲染环境
- 服务端组件：默认情况下，Next.js 中的页面是服务端渲染的，适合用来处理静态内容和 SEO。
- 客户端组件：通过 `"use client"` 指令，可以将某些组件标记为客户端组件，从而实现客户端渲染。

### 交替使用服务端组件和客户端组件
**服务端组件可以直接导入客户端组件，但客户端组件并不能导入服务端组件。**

1. 服务端组件渲染静态内容  
服务端组件负责渲染页面的静态部分。它们可以从数据库或 API 获取数据，并将渲染结果传递到客户端，而不需要发送 JavaScript 到浏览器。

2. 客户端组件处理动态交互  
客户端组件可以在服务端组件中嵌套，它们会在浏览器中加载并处理用户的交互。这样，你可以将需要动态更新的部分（例如按钮点击、表单提交等）单独放在客户端组件中。

3. 通过 `"use client"` 声明客户端组件  
在 Next.js 中，通过在组件的顶部添加 `'use client'` 声明，可以标记一个组件为客户端组件。只有标记为 `'use client'` 的组件，才会在浏览器中执行 JavaScript 代码。

下面我们来做一个验证；在 `/home` 路由中引入一个客户端组件，在客户端组件中引入服务端组件，看看会是什么样的效果：
> 在演示之前先创建一个 Next.js 的项目，命令 `npx create-next-app@latest`，选择配置如下图：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/n4udp1-ed961f96-be8e-40c2-867f-12af1d8321ed.png)
> package.json 的内容如下：
> ```json
> {
> "name": "server-component-and-client-component",
> "version": "0.1.0",
> "private": true,
> "scripts": {
>  "dev": "next dev --turbopack",
>  "build": "next build",
>  "start": "next start",
>  "lint": "next lint"
> },
> "dependencies": {
>  "react": "19.0.0-rc-66855b96-20241106",
>  "react-dom": "19.0.0-rc-66855b96-20241106",
>  "next": "15.0.3"
> },
> "devDependencies": {
>  "typescript": "^5",
>  "@types/node": "^20",
>  "@types/react": "^18",
>  "@types/react-dom": "^18",
>  "postcss": "^8",
>  "tailwindcss": "^3.4.1",
>  "eslint": "^8",
>  "eslint-config-next": "15.0.3"
> }
> }
> ```
> 准备好环境之后，我们就开始进入正题！

- 创建服务端组件
  ```jsx
  // src/components/server-components/example.tsx
  import React from 'react'
  
  async function ServerComponentExample() {
      const res = await fetch('https://jsonplaceholder.typicode.com/todos/1')
      const data = await res.json()
  
      return (
          <div>
              <h2>Server Component Example</h2>
              <p>{data.title}</p>
          </div>
      )
  }
  
  export default ServerComponentExample
  ```

- 创建客户端组件
  ```jsx
  // src/components/client-components/example.tsx
  
  'use client';
  import React from 'react'
  import ServerComponentExample from '../server-components/example'
  
  function ClientComponentExample() {
      return (
          <div>
              <h2>Client Component Example</h2>
              <ServerComponentExample />
          </div>
      )
  }
  
  export default ClientComponentExample
  ```
- 在 `/home` 路由中引入客户端组件
  ```jsx
  // src/app/home/page.tsx
  import React from 'react'
  import ClientComponentExample from '@/components/client-components/example'
  
  function HomePage() {
      return (
          <div>
              <h1>Home Page</h1>
              <p>This is the home page.</p>
              <ClientComponentExample />
          </div>
      )
  }
  
  export default HomePage
  ```
- 在浏览器中访问 `/home` 路由；如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/5qrofP-63d1a613-d215-42b3-a622-cbb0f0a8d209.png)
> 上面的代码可以在 [https://github.com/clin211/next-awesome/commit/b79c11507aae174319e5d04158d369b216ac417d](https://github.com/clin211/next-awesome/commit/b79c11507aae174319e5d04158d369b216ac417d) 中找到！

**虽然 Next.js 不支持客户端组件中使用服务端组件，但支持将服务器组件作为 Props 传递给客户端组件**。
我们来改一下上面的例子：
- 修改客户端组件
```jsx
// src/components/client-components/example.tsx

'use client';
import React from 'react'

function ClientComponentExample({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <div>
            <h2>Client Component Example</h2>
            {children}
        </div>
    )
}

export default ClientComponentExample
```
- 修改 `/home` 文件
```jsx
// src/app/home/page.tsx
import React from 'react'
import ClientComponentExample from '@/components/client-components/example'
import ServerComponentExample from '@/components/server-components/example'

function HomePage() {
    return (
        <div>
            <h1>Home Page</h1>
            <p>This is the home page.</p>
            <ClientComponentExample>
                <ServerComponentExample />
            </ClientComponentExample>
        </div>
    )
}

export default HomePage
```

在浏览器中访问 [http://localhost:3000/home](http://localhost:3000/home) 效果如下：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/zvFoQ6-3ea33b2e-6692-4138-877c-5a7cba0b5d5e.png)


### 服务端组件是怎么渲染的
在服务端，Next.js 使用 React API 编排渲染，渲染工作会根据路由和 Suspense 拆分成多个块（chunks），每个块分两步进行渲染：

- React 将服务端组件渲染成一个特殊的数据格式称为 React Server Component Payload (RSC Payload)
- Next.js 使用 RSC Payload 和客户端组件代码在服务端渲染 HTML
> RSC payload 中包含如下这些信息：
>
> - 服务端组件的渲染结果
> - 客户端组件占位符和引用文件
> - 从服务端组件传给客户端组件的数据

在客户端：

- 加载渲染的 HTML 快速展示一个非交互界面（Non-interactive UI）
- RSC Payload 会被用于协调（reconcile）客户端和服务端组件树，并更新 DOM
- JavaScript 代码被用于水合客户端组件，使应用程序具有交互性（Interactive UI）


### 最佳实践

#### 服务端组件间的数据共享

当在服务端获取数据的时候，有可能出现多个组件共用同一个数据的情况。这种情况下，你不需要使用 React Context（因为服务端组件用不了），也不需要通过 props 传递数据，直接在需要的组件中请求数据即可。这是因为 React 拓展了 fetch 的功能，添加了[记忆缓存功能](https://nextjs.org/docs/app/building-your-application/caching#request-memoization)，相同的请求和参数，返回的数据会做缓存。如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/J68Ylm-1319c447-170f-4eb3-99be-6c9ce8631f9e.png)

看一个官方的例子：
```jsx
async function getItem() {
  // `fetch` 函数会自动缓存结果
  const res = await fetch('https://.../item/1')
  return res.json()
}
 
// 该函数被调用两次，但只在第一次执行
const item = await getItem() // 缓存未命中
 
// 第二次调用可能在路由中的任何地方
const item = await getItem() // 缓存命中
```
虽然 `getItem()` 函数被调用两次，但只有第一次调用才会执行，后面的调用都是使用缓存，当然这个缓存也是有一定条件限制的，比如仅适用于请求 `GET` 中的方法 `fetch`，其他方法（例如 `POST` 和 `DELETE`）不会被缓存。


#### 将服务器专用代码排除在客户端环境之外
由于 JavaScript 模块可以在服务器和客户端组件模块之间共享，因此原本只打算在服务器上运行的代码可能会潜入客户端。比如下面这段代码：
```jsx
export async function getData() {
    const res = await fetch('https://external-service.com/data', {
        headers: {
            authorization: process.env.API_KEY,
        },
    })

    return res.json()
}
```

因为 Next.js 环境变量的规则， 没有 NEXT_PUBLIC 前缀的环境变量都是服务端私有变量，为了防止环境变量泄露给客户端，Next.js 会将私有环境变量替换为空字符串！即使这段代码在客户端中执行，它也不会按预期运行。

为了防止这种情况，可以使用 `server-only` 包来做隔离（客户端如果使用服务端模块则会抛出构建错误）。

- 安装 [server-only](https://www.npmjs.com/package/server-only)

  ```sh
  npm install server-only
  ```
- 然后将包导入到任何包含仅限服务器代码的模块中：
  ```js
  import 'server-only'
  
  export async function getData() {
      const res = await fetch('https://external-service.com/data', {
          headers: {
              authorization: process.env.API_KEY,
          },
      })
  
      return res.json()
  }
  ```

现在，任何导入 `getData()` 的客户端组件都会收到构建时错误，以保证该模块只能在服务器上使用。

> 相应的包 [client-only](https://www.npmjs.com/package/client-only) 可用于标记包含仅限客户端代码的模块 - 例如，访问 `window` 对象的代码。

#### 使用第三方包
由于服务器组件是 React 的新功能， React 生态里的很多包可能还没有跟上在客户端组件添加 `"use client"` 指令；就有可能存在一些潜在的问题。

比如你使用了一个导出 `<Carousel />` 组件的 acme-carousel 包。这个组件使用了 `useState`，但是它并没有在文件顶部声明 `"use client"`。

如果你在客户端组件中使用 `<Carousel />`，它将按预期工作：
```jsx
'use client'

import { useState } from 'react'
import { Carousel } from 'acme-carousel'

export default function Gallery() {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <div>
            <button onClick={() => setIsOpen(true)}>View pictures</button>

            {/* Works, since Carousel is used within a Client Component */}
            {isOpen && <Carousel />}
        </div>
    )
}
```
如果你在服务端组件中使用它，则会报错：
```
import { Carousel } from 'acme-carousel'

export default function Page() {
    return (
        <div>
            <p>View pictures</p>

            {/* Error: `useState` can not be used within Server Components */}
            <Carousel />
        </div>
    )
}
```
因为 Next.js 不知道 `<Carousel />` 组件只能在客户端组件中使用，这个问题也好处理，那就我们在封装一个客户端组件，然后把它包在这个客户端组件中；比如在 `conponents/carousel.tsx` 中写入：
```js
'use client'

import { Carousel } from 'acme-carousel'

export default Carousel
```
这么修改后，我们也可以在服务端组件中使用 `<Carousel />` 组件。

#### 使用 Context Provider
上下文提供者通常在应用程序的根部渲染，以共享全局关注点，例如切换主题色、暗黑/浅色模式。

由于 React Context 不支持服务端组件，因此尝试在应用程序的根部创建上下文会导致错误，比如在 layout 中使用 React Context：
```jsx
import { createContext } from 'react'

//  createContext is not supported in Server Components
export const ThemeContext = createContext({})

export default function RootLayout({ children }) {
    return (
        <html>
            <body>
                <ThemeContext.Provider value="dark">{children}</ThemeContext.Provider>
            </body>
        </html>
    )
}
```
写入上面的代码，IDE 就提示：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/IEBFNx-37bf8185-1433-4c3e-87e6-cfe68dadc99f.png)

修复这个问题也比较简单，那就是在客户端组件中进行渲染这对逻辑：
```jsx
'use client'

import { createContext } from 'react'

export const ThemeContext = createContext({})

export default function ThemeProvider({
    children,
}: {
    children: React.ReactNode
}) {
    return <ThemeContext.Provider value="dark">{children}</ThemeContext.Provider>
}
```

然后再在 layout.tsx 中使用：
```jsx
import ThemeProvider from './theme-provider'

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang='en'>
            <body>
                <ThemeProvider>{children}</ThemeProvider>
            </body>
        </html>
    )
}
```
这样应用程序中的所有其他客户端组件都将能够使用 ThemeProvider Context 啦！

## 六、Next.js 中客户端组件的最佳实践
### 客户端组件尽可能下移
为了减少客户端 JavaScript 包的大小，建议将客户端组件从组件树中向下移动。比如，你有一个包含一些静态元素（Logo）和一个交互式的使用状态的搜索框（Search）的布局，就没有必要让整个布局都成为客户端组件，将交互的逻辑部分抽离成一个客户端组件（比如 `<SearchBar />`），让布局成为一个服务端组件：
```jsx
// SearchBar 是一个客户端组件
import SearchBar from '@/components/client-components/searchbar'
// Logo 是一个服务端组件
import Logo from '@/components/server-components/logo';

// 布局组件默认为服务端组件
export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <>
            <nav>
                <Logo />
                <SearchBar />
            </nav>
            <main>{children}</main>
        </>
    )
}
```
> 本文中所有的代码都能在 [https://github.com/clin211/next-awesome/tree/server-component-and-client-component](https://github.com/clin211/next-awesome/tree/server-component-and-client-component) 上面找到！
### 从服务端组件到客户端组件传递的数据需要序列化
- 如果在服务端组件（Server Component）中获取数据，将数据以 Props 方式传递给客户端组件（Client Component）；则传递的数据需要做序列化。
- 如果不能序列化，可以使用第三方库在客户端获取数据，或者在服务端使用路由处理器（Route Handler）获取数据。

> 注意：  
> React 序列化的限制意味着某些类型的数据（如函数、Symbol 等）无法直接传递给客户端组件。如果需要传递这些类型的数据，可以考虑使用其他方法，比如在客户端获取数据或使用服务端渲染等解决方案。
## 七、总结
在现代 Web 应用开发中，服务端组件（Server Components） 和 客户端组件（Client Components） 的合理使用，是提升性能、优化用户体验以及开发效率的关键。通过理解它们的特点和适用场景，可以更好地根据项目需求灵活选择组件类型。

### 服务端组件的特点与优势：
- 无需运行时 JavaScript，适合静态内容渲染。
- 高效的数据获取，直接从服务端获取数据，减少网络延迟。
- 流式渲染，提升首屏加载速度。
- 更适用于SEO 优化和低动态交互需求的页面。
### 客户端组件的特点与优势：
- 支持 动态交互 和 浏览器 API。
- 适用于需要 客户端状态管理 和 复杂交互逻辑 的页面。
- 通过 useState 和 useEffect 等 Hook，能够灵活处理用户行为。
### 如何交替使用服务端组件和客户端组件：
- 服务端组件 渲染静态内容，减少客户端 JavaScript 负载。
- 客户端组件 处理用户交互，支持动态更新。
- 通过 'use client'; 声明，明确组件运行环境，确保功能正确性。
- 结合 动态导入，延迟加载非必要的客户端组件，进一步优化性能。
### 最佳实践：
- 动态性选择：静态内容优先使用服务端组件，动态交互使用客户端组件。
- 流式渲染：提升用户体验，让静态内容尽早显示。
- 性能优化：通过服务端减少客户端负担，通过动态导入降低初始加载压力。
- 模块化组合：将服务端组件和客户端组件结合，构建高效、可维护的页面结构。
#### 应用场景总结：
- 对于 SEO 要求高、内容动态性低的页面，推荐服务端组件。
- 对于交互频繁、需要客户端逻辑支持的页面，推荐客户端组件。
- 混合使用时，应根据项目需求权衡性能与交互需求，灵活设计组件架构。
- 通过合理应用服务端组件和客户端组件，可以充分发挥 Next.js 的能力，构建高性能、用户友好的现代 Web 应用。

『参考资料』
- Client Components：[https://nextjs.org/docs/app/building-your-application/rendering/client-components](https://nextjs.org/docs/app/building-your-application/rendering/client-components)
- Server Components：[https://nextjs.org/docs/app/building-your-application/rendering/server-components#server-rendering-strategies](https://nextjs.org/docs/app/building-your-application/rendering/server-components#server-rendering-strategies)
- Server and Client Composition Patterns：[https://nextjs.org/docs/app/building-your-application/rendering/composition-patterns](https://nextjs.org/docs/app/building-your-application/rendering/composition-patterns)
- Reusing data across multiple functions：[https://nextjs.org/docs/app/building-your-application/data-fetching/fetching#reusing-data-across-multiple-functions](https://nextjs.org/docs/app/building-your-application/data-fetching/fetching#reusing-data-across-multiple-functions)
- Request Memoization：[https://nextjs.org/docs/app/building-your-application/caching#request-memoization](https://nextjs.org/docs/app/building-your-application/caching#request-memoization)