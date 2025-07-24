> 上一篇文章《Next.js 全栈框架入门：从零搭建你的全栈应用》，介绍了如何使用 Next.js 官方脚手架工具 [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app) 快速搭建一个全栈应用，并详细演示了如何手动创建和配置一个 Next.js 项目，包括安装依赖、配置 TypeScript 环境，以及创建基本的文件系统路由。
>
> 这篇文章我们来看看 Next.js 中的路由系统，路由是 Next.js 应用的核心。Next.js 有两套路由解决方案，Next.js v13 之前只有 Pages Router，Next.js v13 开始引入了 App Router，从 v13.4 开始，App Router 已成为默认的路由方案。两套方案在新版本中都是兼容的，官方也比较推荐 App Router，接下面我们也是基于 App Router 进行探讨！
>
> > 如果一个 Next.js 应用中，既包含 App Router，又包含 Pages Router 时，App Router 的优先级高于 Pages Router！如果两者解析为同一个 URL，会导致构建错误。
>
>
> ## 文件夹与文件的作用
>
> Next.js 是基于文件系统的路由器；文件夹用于定义路由，路由是嵌套文件夹的单一路径，遵循文件系统层次结构，从根文件一直到包含 `page.tsx` 或者 `page.jsx` 的文件。
>
> > `page.tsx` 或者 `page.jsx` 是 Next.js 中一个特殊的文件，它导出一个 React 组件，作为该路由呈现的页面。
>
> ### 文件系统
>
> 直白点讲就是，一个文件就是一个路由。比如：
> - 在 Pages Router 中：
>   ```tree
>   /pages
>     ├── index.tsx         // 对应于 /
>     ├── about.tsx          // 对应于 /about
>     └── blog
>         ├── index.tsx      // 对应于 /blog
>         └── [id].tsx       // 对应于 /blog/:id
>   ```
>   - `pages/index.tsx` 对应的 `/` (入口)的路由。
>
>   - `pages/about.tsx` 对应的 `/about` 的路由。
>   - `pages/blog/index.tsx` 对应于 `/blog` 的路由。
>   - `pages/blog/[id].tsx` 对应于动态路由 是`/blog/:id`。
>   
> - 在 App Router 中：
>
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/wMob6i-d86d2e58-f9eb-480a-937e-37add03e58bd.png)
>   
>   上图展示了文件夹如何映射到 URL 片段；可以使用 `page.tsx` 文件为每个路由创建单独的用户界面。将上面 Pages Router 示例改成 App Router 后，目录结构如下：
>   ```tree
>   └── app
>     ├── page.ts s
>     ├── about
>     │   └── page.ts
>     └── blog
>         ├── [id]
>         │   └── page.ts
>         └── page.ts
>   ```
>   - `app/page.tsx` 对应的 `/` (入口)的路由。
>
>   - `app/about/page.tsx` 对应的 `/about` 的路由。
>   - `app/blog/page.tsx` 对应于 `/blog` 的路由。
>   - `app/blog/[id]/page.tsx` 对应于动态路由 是`/blog/:id`。
>   
> ### 文件约定
>
> > 特殊文件可用的 `.js`、`.jsx` 或 `.tsx` 文件扩展名。
>
> | 文件名       | 说明                            |
> | ------------ | ------------------------------- |
> | layout       | 布局 UI                         |
> | page         | 对应路由所呈现的页面            |
> | loading      | 正在加载当前路由或者子路由的 UI |
> | not-found    | 未找到当前路由及其子路由的 UI   |
> | error        | 当前路由及其子子路由的错误 UI   |
> | global-error | 全局错误用户界面                |
> | route        | 服务器端 API 端点               |
> | template     | 专门重新渲染的布局 UI           |
> | default      | 并行路由的后备 UI               |
>
> ### 组件的层次结构
> 路由片段的特殊文件中定义的 React 组件按照特定的层次结构进行渲染：
> - `layout.js`
> - `template.js`
> - `error.js`（React 错误边界）
> - `loading.js`（React Suspense 边界）
> - `not-found.js`（React 错误边界）
> - `page.js` 或嵌套 `layout.js`
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/xV2wtk-9202acc2-da3d-435a-843e-d01ed5859282.png)
> 在嵌套路由中，子片段的组件将嵌套在其父片段的组件内。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/RprGED-41c04cd6-e747-4d63-8106-a58677977e69.png)
>
> ## 路由约定
>
> 在版本 13 中，Next.js 引入了一个基于 React Server Components 构建的新 App Router，它支持共享布局、嵌套路由、加载状态、错误处理等。
>
> ### 项目搭建
>
> 在进入正题之前，我们先创建一个 Next.js 的项目，结合代码及浏览器效果能更好的理解各个路由规则！
> 使用命令 `npx create-next-app@latest nextjs-app-router --use-pnpm` 创建项目(`--use-pnpm` 表示使用 pnpm 创建项目，如果没有 pnpm，也可以使用 npm 或者 yarn)；可以选择你自己喜欢的技术栈，我的配置如下图：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/czGDWs-d550f83b-a778-46af-834c-08bb918be791.png)
>
> 在 VS Code 中打开后入下图：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/7YF3KJ-3c9be634-016c-4ced-8819-d63655d519a8.png)
>
> 在浏览器中打开后如下图：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/F1E0Hv-d435452c-3e7b-446f-a2d1-8b1911b80b2a.png)
>
>
> ### 路由文件
>
> | 文件名         | 支持文件              | 说明             |
> | -------------- | --------------------- | ---------------- |
> | `layout`       | `.js`、`.jsx`、`.tsx` | 布局             |
> | `page`         | `.js`、`.jsx`、`.tsx` | 页面             |
> | `loading`      | `.js`、`.jsx`、`.tsx` | 加载界面         |
> | `not-found`    | `.js`、`.jsx`、`.tsx` | 未找到用户界面   |
> | `error`        | `.js`、`.jsx`、`.tsx` | 错误用户界面     |
> | `global-error` | `.js`、`.jsx`、`.tsx` | 全局错误用户界面 |
> | `route`        | `.js`、`.ts`          | API 接口         |
> | `template`     | `.js`、`.jsx`、`.tsx` | 重新渲染布局     |
> | `default`      | `.js`、`.jsx`、`.tsx` | 并行路由回退页面 |
>
> 下面我们来逐个看看这些文件的具体表现形式是怎样的！
> #### **layout**
>
> 在当前路由及所有子路由下共享的 UI。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/1hXocV-4ab2d145-56b4-4c76-a627-b61ff2c6dac8.png)
>
> 在上图的目录结构中，`dashboard/layout.tsx` 是 `/dashboard/settings` 和 `/dashboard/analytics` 两个路由的通用布局。
>
> 我们来实践一下，在 `app/dashboard` 目录下创建一个 `layout.tsx` 和 `page.tsx` 文件，分别写入代码如下：
> ```jsx
> // app/dashboard/layout.tsx
> import { FC, PropsWithChildren } from 'react'
> 
> const DashboardLayout: FC<PropsWithChildren> = ({ children }) => {
>     return <section>
>         <nav>dashboard nav</nav>
>         {children}
>     </section>
> }
> 
> export default DashboardLayout
> ```
> ```jsx
> // dashboard/page.tsx
> function page() {
>     return (
>         <div>这里是 dashboard</div>
>     )
> }
> 
> export default page
> ```
> 在 dashboard 中分别创建 `settings/page.tsx` 和 `analytics/page.tsx` 文件，分别写入代码如下：
> ```jsx
> // settings/page.tsx
>   const page = () => {
>     return (
>         <div>这里是 dashboard/settings</div>
>     )
> }
> 
> export default page
> ```
>   ```jsx
> // analytics/page.tsx
> function page() {
>     return (
>         <div>这里是 dashboard/analytics</div>
>     )
> }
> 
> export default page
>   ```
>
> 当访问 `/dashboard` 时，效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/jBDSAW-3e0d255a-3ee4-457e-9ca2-5a7425e54837.png)
>
> 当访问 `/dashboard/settings` 时，效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/PTScfi-58144bce-b969-4964-add6-92d0a8a023cb.png)
>
> 当访问 `/dashboard/analytics` 时，效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/bD9h5m-95780566-c125-4d35-b7c3-a29f1daaf074.png)
>
> 根据上面效果，不难发现，同一个目录下，如果有 `layout` 和 `page`，`page` 会作为 `children` 参数传入 `layout` 中，也就是 `layout` 会包裹同层级的 `page`。
>
> #### **root layout**
>
> ![image](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/XSdrg2-64e9dcb6-40b2-4378-967a-1745832df461.png)
>
> layout 也是可以嵌套的！在我们创建项目的时候，Next.js 脚手架工具也会在 `app/` 目录下创建一个 `layout.tsx` 文件；也就是根布局，这个根布局还有一些特殊性需要注意：
> - 在 app 目录下必须包含一个 `layout` 布局文件。
> - 这个布局文件中必须包含 `<html>` 和 `<body>` 标签；其他布局文件不能包含这些标签。如果你要更改这些标签，不推荐直接修改，而是用 [Metadata API](https://nextjs.org/docs/app/api-reference/functions/generate-metadata) 来修改。
> - 可以使用路由组创建多个根布局。
> - 默认根布局是服务端组件，且不能设置为客户端组件。
>
> #### **page**
>
> 页面是用于呈现路由的用户界面的文件。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/CMXk3T-ea223d06-b9b1-4b60-a06f-1a07cd3e3399.png)
>
> 一个 page 对应一个路由；比如：`dashboard/page.tsx` 对应的是 `/dashboard` 路由，`/dashboard/settings/page.tsx` 对应的是 `/dashboard/settings` 路由。
>
> #### **loading**
> 用于展示加载界面的，作用于当前路由及所有子路由。这个功能的实现借助了 React 的 [Suspense API](https://react.dev/reference/react/Suspense)。它实现的效果就是当发生路由变化的时候，立刻展示 fallback UI，等加载完成后，展示数据驱动的页面。
> ```jsx
> <Suspense fallback={<Loading />}>
>     <SomeComponent />
> </Suspense>
> ```
> 我们在 `app/dashboard` 目录下创建一个 `loading.tsx` 文件。目录结构如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/WbSjeU-cf836fb3-0c49-45d6-99b6-df2e913ee018.png)
>
> 然后写入如下代码：
> ```jsx
> // app/dashboard/loading.tsx
> export default function DashboardLoading() {
>     return <>Loading dashboard...</>
> }
> ```
> 我们来修改 `app/dashboard/page.tsx` 使用定时器来模拟网络延迟，测试一下这个 loading 的功能；修改如下：
> ```jsx 
> // app/dashboard/page.tsx
> async function page() {
>     // 用于模拟网络延迟
>     await new Promise(resolve => setTimeout(resolve, 5000))
>     return (
>         <div>这里是 dashboard</div>
>     )
> }
> 
> export default page
> ```
> 当我输入 `http://localhost:3000/dashboard` 回车，能够直观看到页面有一个 loading 的文本（在真实业务场景中可以使用骨架屏等处理），效果如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/tFRG3B-938db7e7-f41c-480b-9906-38852f7fc907.gif)
>
> 上面这个例子我们是在 `/dashboard` 中写的，如果在 `/dashboard/settings` 中也需要特殊定制一个 loading 的页面，则可以在 `/dashboard/settings/loading.tsx` 中自定义加载页面的效果。
>
> #### **not-found**
> 顾名思义，当该路由不存在的时候展示的内容。基于上面的路由系统，我来访问一下 `http://localhost:3000/dashboard/address` 时，会出现如下效果：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/UwaaSz-b689f5e5-eb26-4038-a58f-7dc87f793b99.png)
>
> 上面这个效果是 Next.js 提供的默认效果；当然也可以自定义这个效果，只需要在 app 目录下创建一个 `not-found.tsx` 的文件，然后就可以自定义效果了。
>
> 这个也有一些要注意：
> - 当组件抛出 `notFound()` 函数的时候会展示这个界面。
> - 所请求的路由不存在的时候也会展示这个界面。
> - `not-found.tsx` 这个文件的页面不接受任何的 `props`。
> - 如果 `not-found.tsx` 放到了任何子文件夹下，它只能由 `notFound()` 函数手动触发。执行 `notFound()` 函数时，会由最近的 `not-found.tsx` 来处理。但如果直接访问不存在的路由，则都是由 `app/not-found.tsx` 来处理。
>
> #### **error**
> 这个文件用于捕获服务器组件和客户端组件中发生的意外错误时展示的 UI；也就是当发生错误时的展示 UI。其实它借助了 React 的 [`Error Boundary`](https://react.dev/reference/react/Component#catching-rendering-errors-with-an-error-boundary) 功能。简单来说，就是给 page.js 和 children 包了一层 `ErrorBoundary`。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/xaJcaF-12bceaef-4d4d-4e73-bed5-577b36315cc8.png)
>
> 我们在 `/dashboard` 中来演示一下 error 的效果，在 dashboard 目录下新建一个 error.tsx 文件，目录效果如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/UIlQ61-adfa8a1a-b599-4c0c-abe6-dcaee725ea41.png)
>
> `dashboard/error.tsx` 代码如下：
> ```jsx
> 'use client' // 错误组件必须是客户端组件
> 
> import { useEffect } from 'react'
> 
> export default function Error({ error, reset }: {
>     error: Error & { digest?: string }
>     reset: () => void
> }) {
>     useEffect(() => {
>         console.error(error)
>     }, [error])
> 
>     return (
>         <div>
>             <h2>Something went wrong!</h2>
>             <button
>                 onClick={
>                     // 尝试恢复
>                     () => reset()
>                 }
>             >
>                 Try again
>             </button>
>         </div>
>     )
> }
> ```
> 为了模拟 error 的效果，我们在同级的 page.tsx 中修改代码如下：
> ```jsx
> 'use client'
> 
> import { useState } from 'react';
> 
> async function page() {
>     // 用于模拟网络延迟
>     // await new Promise(resolve => setTimeout(resolve, 5000))
>     const [isErr, setIsErr] = useState(false);
>     const handleOnClickError = () => {
>         setIsErr(true)
>     }
> 
>     return (
>         <div>
>             <p>这里是 dashboard</p>
>             {isErr ? Error() :
>                 <button onClick={handleOnClickError}>Get Error</button>
>             }
>         </div>
>     )
> }
> 
> export default page
> ```
> 效果如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/IMMjyb-bbe0b9fd-c809-46d7-ba13-32d8e1b8cbf9.gif)
>
> 有时错误是暂时的，只需要重试就可以解决问题。所以 Next.js 会在 `error.js` 导出的组件中，传入 `reset()` 函数，帮助尝试从错误中恢复。该函数会触发重新渲染错误边界里的内容。如果成功，会替换展示重新渲染的内容。
>
> #### **global-error**
> 这个是用来专门处理根目录中的错误，也就是跟根目录 layout 同级的一个错误处理文件。
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/xV2wtk-9202acc2-da3d-435a-843e-d01ed5859282.png)
>
> 从图中也可以看出，`Layout` 和 `Template` 在 `ErrorBoundary` 的外面，如果 `Layout` 或者 `Template` 发生了错误，那就需要在父级的 `error.tsx` 中捕获错误。如果在顶层的话，Next.js 就提供了 `global-error` 的方案。 `global-error.tsx` 会包裹整个应用，而且当它触发的时候，它会替换掉根布局的内容。所以，`global-error.tsx` 中也要定义 `<html>` 和 `<body>` 标签。
>
> `app/global-error.tsx` 的代码如下：
> ```jsx
> 'use client'
> 
> export default function GlobalError({
>     error,
>     reset,
> }: {
>     error: Error & { digest?: string }
>     reset: () => void
> }) {
>     return (
>         <html lang="en">
>             <body>
>                 <h2>Something went wrong!</h2>
>                 <button onClick={() => reset()}>Try again</button>
>             </body>
>         </html>
>     )
> }
> ```
> > `global-error.tsx` 用来处理根布局和根模板中的错误，与 `app/error.tsx` 并不冲突。
>
> #### **route**
> 前后端分离架构中，客户端与服务端之间通过 API 接口来交互。这个“API 接口”在 Next.js 中成为路由处理程序。
>
> 在 Next.js 中，写路由处理程序，文件名必须是 `route.ts` 或者 `route.js`，且必须在 `app/` 目录下，还不能与 `page.tsx` 同级存在。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/Vn97Bb-7affac4e-a58c-4ee2-a27d-2c39bda51b37.png)
>
> 支持 `GET`、`POST`、`PUT`、`PATCH`、和 `DELETE` 方法；如果调用不受支持的方法，Next.js 将返回 `405 Method Not Allowed` 的响应。
> ```ts
> // app/api/xxx/route.ts
> export async function GET(request) {}
>  
> export async function HEAD(request) {}
>  
> export async function POST(request) {}
>  
> export async function PUT(request) {}
>  
> export async function DELETE(request) {}
>  
> export async function PATCH(request) {}
>  
> // 如果 `OPTIONS` 没有定义, Next.js 会自动实现 `OPTIONS`
> export async function OPTIONS(request) {}
> ```
>
> 下面我们就用 [jsonplaceholder](https://jsonplaceholder.typicode.com/) 的文章的增删改查来演示一下最常用几个的方法：
> - **GET**
>
>   以请求列表接口为例：
>   ```ts
>   // app/api/posts/route.ts
>   import { NextResponse } from 'next/server'
>
>   export async function GET() {
>       const res = await fetch('https://jsonplaceholder.typicode.com/posts')
>       const data = await res.json()
>
>       return NextResponse.json({ data })
>   }
>   ```
>   在浏览器访问 `http://localhost:3000/api/posts` 效果如下：
>
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/WS4mLx-efa158e3-cad4-4c77-a2d8-4ec2c72041fe.png)
>
> - **POST**
>
>   以创建新文章为例：
>   ```ts
>   export async function POST() {
>       const res = await fetch('https://jsonplaceholder.typicode.com/posts', {
>           method: 'POST',
>           body: JSON.stringify({
>               title: 'foo',
>               body: 'bar',
>               userId: 1,
>           }),
>           headers: {
>               'Content-type': 'application/json; charset=UTF-8',
>           },
>       })
>
>       const data = await res.json()
>       return NextResponse.json({ code: 200, data, message: 'success' })
>   }
>   ```
>   在 postman 中请求 `http://localhost:3000/api/posts` 后效果如下：
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/1eAUXj-57fcc147-78cd-4176-8c74-713d07b9d8a0.png)
>
> - **PUT**
>
>   以更新 id 为 1 的文章为例：
>   ```ts
>   export async function PUT() {
>       const res = await fetch('https://jsonplaceholder.typicode.com/posts/1', {
>           method: 'PUT',
>           body: JSON.stringify({
>               id: 1,
>               title: 'foo',
>               body: 'bar',
>               userId: 1,
>           }),
>           headers: {
>               'Content-type': 'application/json; charset=UTF-8',
>           },
>       })
>
>       const data = await res.json()
>       return NextResponse.json({ code: 200, data, message: 'success' })
>   }
>   ```
>   在 postman 中请求 `http://localhost:3000/api/posts` 后效果如下：
>
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/yZt1li-14f9de4a-e0b3-44a9-9dcf-018f15a0a33b.png)
>
> - **patch**
>
>   以更新 id 为 1 的文章的 title 为例：
>   ```ts
>   export async function PATCH() {
>       const res = await fetch('https://jsonplaceholder.typicode.com/posts/1', {
>           method: 'PATCH',
>           body: JSON.stringify({
>               id: 1,
>               title: 'foo',
>               body: 'bar',
>               userId: 1,
>           }),
>           headers: {
>               'Content-type': 'application/json; charset=UTF-8',
>           },
>       })
>
>       const data = await res.json()
>       return NextResponse.json({ code: 200, data, message: 'success' })
>   }
>   ```
>   在 postman 中请求 `http://localhost:3000/api/posts` 后效果如下：
>   
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/0ZYMAK-275456cd-52db-4b9a-9613-8042eb5aeb7b.png)
>
> - **delete**
>
>   以删除 id 为 1 的文章为例：
>   ```ts
>   export async function DELETE() {
>       const res = await fetch('https://jsonplaceholder.typicode.com/posts/1', {
>           method: 'DELETE',
>       })
>     
>       const data = await res.json()
>       return NextResponse.json({ code: 200, data, message: 'success' })
>   }
>   ```
>   在 postman 中请求 `http://localhost:3000/api/posts` 后效果如下：
>   
>   ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/JXc5iy-d204e890-6f96-45dd-85c2-a4d175f92bb0.png)
>
> 如果还有想要了解 FormData 或者 Stream 等相关的内容，可以阅读官方[路由处理程序（Route Handlers）](https://nextjs.org/docs/app/building-your-application/routing/route-handlers)。
> #### **template**
> 模板类似于布局，它也会传入每个子布局或者页面。但不会像布局那样维持状态。也就是模板在路由切换时会为每一个 `children` 创建一个实例。这就意味着在多个路由中共享一个模板，各路由间跳转的时候，将会重新挂载组件实例，重新创建 DOM 元素，不会保留原来的状态。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/urWwef-68ef50c9-ba63-4090-b0fe-add482ab31ed.png)
>
> 定义一个模板，必须以 `template` 为文件名，且默认导出一个 React 组件，这个组件接收一个 `children` 参数。如上图，我们再 app 目录下创建一个模板，并写入如下代码：
> ```jsx
> import { FC, PropsWithChildren } from 'react'
> 
> const RootTemplate: FC<PropsWithChildren> = ({ children }) => {
>     return <div>{children}</div>
> }
> 
> export default RootTemplate
> ```
>
> 如果在同一目录层级下既有 layout，又有 template 时，它们的关系如下：
> ```jsx
> <Layout>
>     {/* Note that the template is given a unique key. */}
>     <Template key={routeParam}>{children}</Template>
> </Layout>
> ```
>
> 在这些场景下，使用 template 比使用 layout 更适合：
> - 依赖于 `useEffect` 和 `useState` 的功能，比如记录页面访问数（维持状态就不会在路由切换时记录访问数了）、用户反馈表单（每次重新填写）等。
> - 更改框架的默认行为，举个例子，布局内的 `Suspense` 只会在布局加载的时候展示一次 fallback UI，当切换页面的时候不会展示。但是使用模板，`fallback` 会在每次路由切换的时候展示。
>
> #### **default**
>
> `default.js` 文件用于在[并行路由](https://nextjscn.org/docs/app/building-your-application/routing/parallel-routes)中渲染备用内容，当 Next.js 无法在完整页面加载后恢复[插槽](https://nextjscn.org/docs/app/building-your-application/routing/parallel-routes#%E6%8F%92%E6%A7%BD)的活动状态时使用。
>
> 在 软导航 期间，Next.js 会跟踪每个插槽的活动状态 (子页面)。然而，对于硬导航 (完整页面加载)，Next.js 无法恢复活动状态。在这种情况下，可以为不匹配当前 URL 的子页面渲染 `default.js` 文件。这块涉及到并行路由的概念，后续再演示！
>
> ## Next.js 中路由分类
>
> ### 嵌套路由
> | 文件名          | 说明                                                         |
> | --------------- | ------------------------------------------------------------ |
> | `folder`        | 路由片段；比如：`app/dashboard/page.tsx`                     |
> | `folder/folder` | 嵌套路由片段；比如：`app/dashboard/settings/page.tsx`和 `app/dashboard/analytics/page.tsx` |
>
> 上面的 `dashboard/settings` 和 `dashboard/analytics` 就是典型的嵌套路由，这里就不再重复演示了。
>
> ### 动态路由
>
> 在某些情况下，我们无法预先确定路由的具体地址，比如需要根据 URL 中的 `id` 参数来展示对应 `id` 的商品详情。由于商品种类繁多，不可能为每个商品单独定义一个路由。在这种情况下，动态路由就显得非常必要。
>
> | 文件名              | 说明                                                      |
> | ------------------- | --------------------------------------------------------- |
> | `[folderName]`      | 动态路由片段；比如：`app/post/[slug]/page.tsx`            |
> | `[...folderName]`   | 捕获所有路由片段；比如：`app/post/[...slug]/page.tsx`     |
> | `[[...folderName]]` | 可选的综合路由片段；比如：`app/post/[[...slug]]/page.tsx` |
>
> #### **`[folderName]`（动态片段）**
> 使用动态路由，你需要将文件夹的名字用方括号括住，比如 `[id]`、`[slug]`。这个路由的名字会作为 `params` prop 传给 layout、page、route 以及 `generateMetadata` 函数。
>
> 我们用 [jsonplaceholder](https://jsonplaceholder.typicode.com/guide/) 的文章接口为例，在 `app/` 目录下创建 `posts/page.tsx` （文章列表）和 `posts/[slug]/page.tsx` （文章详情）文件，分别写入以下内容：
> ```jsx
> // app/posts/page.tsx
> import Link from 'next/link';
> 
> interface Post {
>     userId: number;
>     id: number;
>     title: string;
>     body: string;
> }
> 
> const page = async () => {
>     const posts = await fetch('https://jsonplaceholder.typicode.com/posts');
> 
>     const data = await posts.json() as Post[];
> 
>     return (
>         <div className="p-4 max-w-screen-md mx-auto bg-#f5f5f5">
>             <div className="space-y-4">
>                 {data.map((post) => (
>                     <Link href={`/posts/${post.id}`} key={post.id}>
>                         <div className="bg-white shadow-md rounded-lg p-4 hover:shadow-lg transition-all">
>                             <div className="mb-2">
>                                 <h2 className="text-xl font-semibold">{post.title}</h2>
>                             </div>
>                             <div className="text-gray-700">
>                                 <p>{post.body}</p>
>                             </div>
>                         </div>
>                     </Link>
>                 ))}
>             </div>
>         </div>
>     )
> }
> 
> export default page
> ```
> ```jsx
> // app/posts/[slug]/page.tsx
> const page = async ({ params }: { params: { slug: string } }) => {
>     const { slug } = params
> 
>     const res = await fetch(`https://jsonplaceholder.typicode.com/posts/${slug}`)
>     const post = await res.json()
> 
>     return <div className="p-4 max-w-3xl mx-auto bg-white rounded-lg">
>         <div className="mb-4">
>             <h1 className="text-3xl font-semibold">{post.title}</h1>
>         </div>
>         <div className="text-lg text-gray-800">
>             <h5>My Post id: {slug}</h5>
>             <p>{post.body}</p>
>         </div>
>     </div>
> }
> 
> export default page
> ```
> 文章列表效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/nd8BM8-9184d025-d2e8-4ce0-bad5-f4b29cb98878.png)
>
> 文章详情效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/fLFwGK-7894c90c-9846-48f0-8761-16a20d48e2df.png)
>
> 路由与参数的对应关系：
> | 路由                        | URL 示例      | 参数             |
> | --------------------------- | ------------- | ---------------- |
> | `app/posts/[slug]/page/tsx` | `/posts/a`    | `{slug: 'a'}`    |
> | `app/posts/[slug]/page/tsx` | `/posts/1`    | `{slug: 1}`      |
> | `app/posts/[slug]/page/tsx` | `/posts/text` | `{slug: 'text'}` |
>
> #### **[...folderName]（捕获所有片段）**
> 捕获 folder 后面所有的路由片段。也就是说，`app/post/[...slug]/page.tsx` 会匹配 `/post/lifestyle`、`/post/lifestyle/travel` 和 `/post/lifestyle/travel/europe`；下面我们将实际演示：
> ```jsx
> // app/posts/[...lifestyle]/page.tsx
> const page = (params: any) => {
>     return (
>         <div>life style page: {JSON.stringify(params)}</div>
>     )
> }
> 
> export default page
> ```
> 效果如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/FKCDkB-8ccadfc9-a4bb-48b0-816b-4acb9807d7be.png)
>
> 路由与参数的对应关系：
> | 路由                           | URL 示例          | 参数                         |
> | ------------------------------ | ----------------- | ---------------------------- |
> | `app/posts/[...slug]/page/tsx` | `/posts/a`        | `{slug: ['a']}`              |
> | `app/posts/[...slug]/page/tsx` | `/posts/a/b`      | `{slug: ['a', 'b']}`         |
> | `app/posts/[...slug]/page/tsx` | `/posts/text/a/b` | `{slug: ['text', 'a', 'b']}` |
>
> #### **[[...folderName]]（可选的捕获所有片段）**
> 捕获所有段可以通过将参数包含在双方括号中来设为可选：[[...folderName]]。按照官方的解释，`app/shop/[[...slug]]/page.js` 除了匹配 `/shop` 之外，还匹配`/shop/clothes`、`/shop/clothes/tops` 和 `/shop/clothes/tops/t-shirts`
>
> 捕获所有片段和可选捕获所有片段的区别在于，可选的情况下，不带参数的路由也会被匹配 (上例中的 `/shop`)。
>
> 我们以 shop 为例，来演示一下；在 `app/` 下创建 `shop/[[...slug]]/page.tsx`，并写入一下代码：
> ```jsx
> // app/shop/[[...slug]]/page.tsx
> const page = (params: any) => {
>     return (
>         <div>posts [[...slug]] page：{JSON.stringify(params)}</div>
>     )
> }
> 
> export default page
> ```
> 访问 `http://localhost:3000/shop` 效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/dQuJtE-bbbd397f-553e-4e83-aeaf-b6d3398431d4.png)
>
> 访问 `http://localhost:3000/shop/clothes/top` 效果如下：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/ODDDSV-c7b06163-976c-40fb-94ea-022f1f562a39.png)
> 路由与参数的对应关系：
> | 路由                             | URL 示例         | 参数                         |
> | -------------------------------- | ---------------- | ---------------------------- |
> | `app/posts/[[...slug]]/page/tsx` | `/shop`          | `{}`                         |
> | `app/posts/[[...slug]]/page/tsx` | `/shop/a`        | `{slug: ['a']}`              |
> | `app/posts/[[...slug]]/page/tsx` | `/shop/a/b`      | `{slug: ['a', 'b']}`         |
> | `app/posts/[[...slug]]/page/tsx` | `/shop/text/a/b` | `{slug: ['text', 'a', 'b']}` |
>
> 
>
> ### 路由组（逻辑分组）
>
> | 文件名     | 说明                                                         |
> | ---------- | ------------------------------------------------------------ |
> | `(folder)` | 用于逻辑分组的使用场景；比如：`app/(auth)/login/page.tsx` 在对应的路由是 `/login` |
>
> 在 `app/` 下，目录名称通常会被映射到 URL 中，但你可以将文件夹标记为路由组，阻止文件夹名称被映射到 URL 中。
>
> > 使用路由组，可以将路由和项目文件按照逻辑进行分组，但不会影响 URL 路径结构。路由组可用于比如：
> > - 按站点、意图、团队等将路由分组。
> > - 在同一层级中创建多个布局，甚至是创建多个根布局。
>
> **创建路由组就是把文件夹用括号括起来**就可以了；比如下图中的 `(marketing)` 和 `(shop)`：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/qzYvn1-6448d900-3ca7-4474-a7cd-40921e9d82fe.png)
>
>
> #### 根据路由组创建单独的根布局
> 要创建多个根布局，删除顶级 `layout.tsx` 文件，并在每个路由组内添加一个 `layout.tsx` 文件。这对于将应用程序划分为具有完全不同 UI 或体验的部分很有用。需要在每个根布局中添加 `<html>` 和 `<body>` 标签。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/1uGRu0-9ba8d2ec-69bc-4ed2-bb36-41fd42563156.png)
>
> #### 根据路由组创建单独的布局
>
> 在上图中，虽然 (marketing) 和 (shop) 内的路由共享相同的 URL 层级结构，但你可以通过在它们的文件夹中添加 layout 文件为每个路由组创建不同的布局。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/oEuHfn-b4fc8ed1-af9e-4173-9d82-0e94397290ed.png)
>
> > Tips:
> > - **路由组的命名除了标识为一个组之外没有特殊意义**；它们不会影响 URL 路径。
> > - **包含路由组的路由不应该解析为与其他路由相同的 URL 路径**。例如：`(marketing)/about/page.tsx` 和 `(shop)/about/page.tsx` 都会解析为 `/about`；这就会导致错误。
> > - 如果你使用多个根布局而没有顶级 `layout.tsx` 文件，你的主页 `page.tsx` 文件应该定义在其中一个路由组中，例如：`app/(marketing)/page.tsx`。
> > - 在多个根布局之间的导航会触发完整的页面加载（而不是客户端导航）。例如，从使用 `app/(shop)/layout.tsx` 的 `/cart` 导航到使用 `app/(marketing)/layout.tsx` 的 `/blog` 将导致完整页面加载。这仅适用于多个根布局。
>
> ### 平行路由
> 平行路由可以使你在同一个布局中同时或者有条件的渲染一个或者多个页面（类似于 Vue 的插槽功能）。对于应用程序中高度动态的部分，如社交网站上的仪表盘和信息源，平行路由非常有用。
>
> 平行路由是使用命名插槽创建的，插槽是按照 `@folder` 约定定义的，比如在上图中就定义了两个插槽 `@team` 和 `@analytics`。
>
> #### 平行路由的对应关系
> | 文件名           | 说明                                               |
> | ---------------- | -------------------------------------------------- |
> | `@folder`        | 在同一个布局中同时或者有条件的渲染一个或者多个页面 |
> | `(.)folder`      | 表示匹配同一层级；比如：`app/@modal/(.)settings`   |
> | `(..)folder`     | 表示匹配上一层级                                   |
> | `(..)(..)folder` | 表示匹配上上层级                                   |
> | `(...)folder`    | 表示从根目录拦截                                   |
>
> 下面就来实践一下！
>
> #### 有条件的渲染
>
> 在一些后管理系统中，你可以通过使用平行路由基于某些条件 (如用户角色) 有条件地渲染路由。例如，为 `/admin` 或 `/user` 角色渲染不同的仪表盘页面：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/geArBq-b5a115e1-0cb1-4f22-9d5d-41833399a778.png)
>
> 在上图，插槽会通过 props 传入这个共享的 `layout` 中，然后 `layout` 从 `props` 中获取 `admin` 和 `user` 两个插槽的内容，并将其渲染。
>
> 除了条件渲染外，还可以并行渲染。比如，考虑一个仪表盘，你可以使用平行路由同时渲染 "team" 和 "analytics" 页面。
> ```jsx
> export default function Layout({
>     children,
>     team,
>     analytics,
> }: {
>     children: React.ReactNode;
>     analytics: React.ReactNode;
>     team: React.ReactNode;
> }) {
>     return (
>         <>
>             {children}
>             {team}
>             {analytics}
>         </>
>     );
> }
> ```
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/kVkSSY-e58be249-8054-4887-94ec-c1aab7a7045b.png)
>
> #### 标签组
> 在插槽内添加一个 `layout`，允许用户独立导航该插槽。
>
> 例如，`@analytics` 插槽有两个子页面：`/page-views` 和 `/visitors`。结构如下图：
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/XlNqeC-6f84f690-1552-4f5f-8d92-14393e6143e5.png)
> 在 `@analytics` 内创建一个 `layout` 文件，在两个页面之间共享标签：
> ```jsx
> import Link from "next/link";
> 
> export default function Layout({ children }: { children: React.ReactNode }) {
>     return (
>         <>
>             <nav>
>                 <Link href="/page-views">页面浏览量</Link>
>                 <Link href="/visitors">访问者</Link>
>             </nav>
>             <div>{children}</div>
>         </>
>     );
> }
> ```
>
> #### 独立的路由处理
>
> 平行路由可以独立流式传输，允许开发者为个路由定义独立的错误和加载状态：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/zTUauV-8a31f06f-a6a5-46e9-9e30-3fce5c128345.png)
>
>
> #### 子导航
>
> ```tree
> app
> ├── parallel-route
> │   ├── @analytics
> │   │   ├── page-views
> │   │   │   └── page.tsx
> │   │   ├── visitors
> │   │   │   └── page.tsx
> │   │   └── page.tsx
> │   └──  layout.tsx
> ```
> 平行路由跟路由组一样，不会影响 URL，所以 `app/parallel-route/@analytics/page-views/page.tsx` 对应的地址是 `/parallel-route/page-views`，`app/parallel-route/@analytics/visitors/page.tsx` 对应的地址是 `/parallel-route/visitors`。下面来分别实现这些页面的跳转和内容显示：
>
> - `app/parallel-route/layout.tsx` 的代码如下：
>   ```jsx
>   import Link from "next/link";
>     
>   export default function RootLayout({ analytics }: { children: React.ReactNode, analytics: React.ReactNode }) {
>       return (
>           <>
>               <nav className='flex items-center gap-4'>
>                   <Link href="/parallel-route">Home</Link>
>                   <br />
>                   <Link href="/parallel-route/page-views">Page Views</Link>
>                   <br />
>                   <Link href="/parallel-route/visitors">Visitors</Link>
>               </nav>
>               <h1>root layout</h1>
>               <div>
>                   {analytics}
>               </div>
>           </>
>       );
>   }
>   ```
> - `app/parallel-route/@analytics/page.tsx` 的代码如下：
>   ```jsx
>   import React from 'react'
>     
>   const page = () => {
>       return (
>           <div>这里是 @analytics/page.tsx </div>
>       )
>   }
>     
>   export default page
>   ```
> - `app/parallel-route/@analytics/page-views/page.tsx` 的代码如下：
>   ```jsx
>   import React from 'react'
>     
>   const page = () => {
>       return (
>           <div>page-views page</div>
>       )
>   }
>     
>   export default page
>   ```
> - `app/parallel-route/@analytics/visitors/page.tsx` 的代码如下：
>   ```jsx
>   import React from 'react'
>     
>   const page = () => {
>       return (
>           <div>visitors page</div>
>       )
>   }
>     
>   export default page
>   ```
>   
>
> 最后效果如下：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/Oq2Oww-373889b5-802c-4c31-b9f7-ae8cc993c6df.gif)
>
>
> ### 拦截路由
> 通俗一点就是允许你在当前路由拦截其他路由地址，并在当前路由中展示内容。
>
> 举个例子，在一个照片列表中，当点击信息流中的照片时，可以在模态框中显示该照片，覆盖在照片列表的上方。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/AJqA03-5875da16-4add-4d1b-8303-7dd26b86220e.png)
>
> 上图左侧就是在一个图片列表中，当点击一个图片信息 `/photo/123` 的时候，结果就是右边这种呈现形式。但当你将地址栏的链接分享出去的时候，就会得到下面的效果：
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/iHGutF-3d459905-c8c1-4355-9993-ba4ff6eee667.png)
>
> > 效果可以去 [https://dribbble.com/](https://dribbble.com/) 的 Explore inspiring designs 真实体验！
>
> #### 文件约定
> 在 Next.js 中，实现拦截路由需要你在命名文件夹的时候以 `(..)` 开头，其中：
>
> - `(.)` 表示匹配同一层级。
> - `(..)` 表示匹配上一层级。
> - `(..)(..)` 表示匹配上上层级。
> - `(...)` 表示匹配根目录。
>
> 但是要注意的是，这个匹配的是路由的层级而不是文件夹路径的层级，就比如路由组、平行路由这些不会影响 URL 的文件夹就不会被计算层级。
>
> 例如下图，`/feed/(..)photo` 对应的路由是 `/feed/photo`，要拦截的路由是 `/photo`，两者只差了一个层级，所以使用 `(..)`。
>
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/b8HvuJ-6602c8cf-b467-42f0-a370-ac41120a3cef.png)
>
> ## 总结
>
> 在 Next.js 中，路由系统分为两种主要类型：Pages Router 和 App Router。
>
> 这两者在路由的管理和结构上存在显著区别:
> - Pages Router：基于文件系统的路由，每个在 `pages/` 的文件自动对应一个路由。这种方式简单易用，适合快速开发。
> - App Router：提供更灵活的路由管理，支持更复杂的应用结构。通过组合和重用组件，开发者可以创建动态和复杂的路由体系。
>
>
> 在 App Router 中，有几个核心文件，各自承担特定的功能和作用，影响着路由的表现和行为：
>
> - `layout`：用于定义页面的布局结构，允许在多个页面中重用相同的布局。可以嵌套实现不同层级的布局。
>
> - `page`：表示具体的页面内容。每个 `page` 文件对应一个路由，负责渲染特定的视图。
> - `template`：用于定义可复用的模板，允许在多个页面中共享相同的结构和样式。
> - `loading`：在数据加载期间显示的占位符或加载动画，提升用户体验，确保用户在等待时不会感到空白。
> - `error (error/global-error)`：处理错误的组件。可以用于捕获并显示应用中的错误信息，提供友好的错误反馈。
> - `not-found`：当访问的路由不存在时显示的页面，通常用于处理 404 错误。
> - `route`：用于定义路由的具体行为和配置，支持更复杂的路由逻辑。
>
> 最后也详细的介绍了动态路由、路由组、平行路由和路由拦截。