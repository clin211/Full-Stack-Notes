随着 Web 应用的复杂性不断提高，用户对页面加载速度和交互流畅度的期望也水涨船高。对于开发者来说，仅仅提升实际性能指标已经不够了，如何优化用户的“感知性能”也变得至关重要。加载动画、骨架屏以及逐步展现内容等技术，成为了改善用户体验的关键手段。

Next.js 作为一个强大的 React 全栈框架。引入了更高效的 Loading UI 和 Streaming 功能，让开发者可以更轻松地实现无缝加载体验。这些特性不仅提升了用户体验，还通过流式渲染减少了服务器和客户端的性能压力。

下面我们来感受下 Next.js 中的 Loading UI 和 Streaming，从基础概念到具体实现，再到实际应用场景，全面解读这些功能的核心原理和使用方法。

在正式进入 Loading UI 和 Streaming 之前，我们先来回顾一下 SSR 渲染（如果对这一块不熟悉的话，推荐去看看《[掌握 Next.js 渲染机制：如何在 CSR、SSR、SSG 和 ISR 中做出最佳选择](https://mp.weixin.qq.com/s/V3FadXse_MIXOZqSBaHAMg)》）。使用 SSR，简单来说，就是需要经过一系列的步骤，用户才能查看页面并与之交互。

具体这些步骤是：
- 首先，在服务器上获取页面的所有数据。
- 然后服务器呈现该页面的 HTML。
- 页面的 HTML、CSS 和 JavaScript 被发送到客户端。
- 使用生成的 HTML 和 CSS 显示非交互式用户界面。
- 最后，React对界面进行[水合（hydrate）](https://react.dev/reference/react-dom/client/hydrateRoot#hydrating-server-rendered-html)，使其具有交互性。

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/kc7HII-37646afd-4c51-4f7f-9c17-450a85d485d3-20241209232950014.png)

这些步骤是连续的、阻塞的。也就是服务器需等所有数据获取完成后才能渲染 HTML，客户端也需等所有组件代码加载完毕后才能对 UI 进行水合：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/Gc4wgd-5b433ffb-393a-4597-b856-1da4a7cf512b-20241209232950006.png)

React 18 为了解决上面这些问题，引入了 [Suspense](https://react.dev/reference/react/Suspense) 组件。

## Suspense

在 React 中，`Suspense` 是一个用于处理异步加载的组件，旨在简化代码和改善用户体验。它允许开发者定义组件在加载异步数据或资源时的备用 UI（通常是加载指示器）。

### 基本用法
```jsx
import React, { Suspense } from 'react';

function App () {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <OtherComponent />
        </Suspense>
    );
}
```
你可以将动态组件包装在 `Suspense` 中，然后向其传递一个 `fallback UI`，以便在动态组件加载时显示。如果数据请求缓慢，使用 Suspense 流式渲染该组件，不会影响页面其他部分的渲染，更不会阻塞整个页面。

下面我们写一个案例，如何将 `Suspense` 和 `use` 结合使用来优雅地处理异步操作，下面是核心代码（这里使用的是 react 19版本）：
```jsx
import { Suspense, use } from "react";

const todo = async () => {
  const res = await fetch("https://jsonplaceholder.typicode.com/todos/1");
  return await res.json();
}

interface TodoItem {
  title: string;
  completed: boolean;
  id: number;
}

const Todo = ({ promise }: { promise: Promise<TodoItem> }) => {
  const todoData = use(promise);
  return (
    <div>
      <h2>Todo</h2>
      <p>{todoData.title}</p>
    </div>
  );
}
export default function App() {
  return (
    <div>
      <h2>App</h2>
      <Suspense fallback={<p>Loading...</p>}>
        <Todo promise={todo()} />
      </Suspense>
    </div>
  );
}
```
效果如下：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/BjSWd1-fee6e138-1f24-401b-a089-8e7561d069c0-20241209232949993.gif)

完整代码可以查看 [https://github.com/clin211/react-awesome/tree/react19-use-suspense](https://github.com/clin211/react-awesome/tree/react19-use-suspense)。

### 在 Next.js 中使用 Suspense 组件

> 在正式演示之前先来创建下项目：
> ```sh
> npx create-next-app@latest --use-pnpm
> ```
> 配置如下图：
> ![创建项目配置选项](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/cViN7J-fdcacd9d-b6ae-4459-9d4d-eaed3a4040f2-20241209232950013.png)

我们以电商系统的产品详情的产品信息、用户评论和产品推荐等功能模块为例：

```jsx
import React, { Suspense } from 'react'

const sleep = (ms: number) => new Promise(r => setTimeout(r, ms));

// 产品信息
const ProductInfo = async () => {
    await sleep(2000); // 模拟异步操作
    return <h1>product info</h1>
}

// 产品评论
const ProductComments = async () => {
    await sleep(3000); // 模拟异步操作
    return <h1>product comments</h1>
}

// 产品推荐
const ProductRecommends = async () => {
    await sleep(5000); // 模拟异步操作
    return <h1>product recommends</h1>
}

export default function page() {
    return (
        <main>
            <Suspense fallback={<h2>Product Info Loading...</h2>}><ProductInfo /></Suspense>
            <Suspense fallback={<h2>Product Comments Loading...</h2>}><ProductComments /></Suspense>
            <Suspense fallback={<h2>Product Recommends Loading...</h2>}><ProductRecommends /></Suspense>
        </main>
    )
}
```
上面这段代码，通过模拟不同的加载时间，可以看到不同的加载状态，确保用户在等待时得到反馈。通过 `Suspense` 实现了异步组件加载时的过渡效果，每个异步组件都有独立的加载指示器。当每个组件的数据或内容加载完成后，相应的组件将被渲染。

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/kbPWPV-60442c32-6aa5-4f9e-b103-b76ee26834f0-20241209232950121.gif)

从上面的 GIF 图中，可以看出 `/product` 路由的加载时间变化，从一开始的 2.08s 到最后的 5.07s，上面的代码中，我们设置的最长时间就是 5s，然后查看网络面板查看 `/product` 路由的详细请求情况如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/trqs0w-1ae94cb0-3b5a-4550-bc7c-60ab29d3fe8e-20241209232950110.png)

其中最关键的就是响应头中 `transfer-encoding: chunked`，表示数据将以一系列分块的形式进行发送。

![来自 MDN 截图](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/kefqGx-1703aa6b-4d14-4188-9a87-2d10c54c6cff-20241209232949994.png)

分块传输编码只在 HTTP 协议1.1版本（HTTP/1.1）中提供！

![来自 MDN 截图](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/mlCuNB-6917e87c-fa95-46f3-8018-365826ac7999-20241209232950115.png)

通过使用 `Suspense`，可以获得以下好处：
- Streaming Server Rendering（流式渲染）：从服务器到客户端渐进式渲染 HTML
- Selective Hydration（选择性水合）：React 根据用户交互决定水合的优先级。

我们还可以通过 `Suspense` 嵌套来控制它的渲染顺序，比如按照： `A组件 --> B组件 --> C组件` 的顺序进渲染，应该怎么做呢？下面代码是在不考虑数据的前后依赖关系的情况下：
```jsx
<Suspense fallback={<h2>A Loading...</h2>}>
    <A />
    <Suspense fallback={<h2>B Loading...</h2>}>
        <B />
        <Suspense fallback={<h2>C Loading...</h2>}>
            <C />
        </Suspense>
    </Suspense>
</Suspense>
```
### Suspense 与 SEO
- Next.js 会等待 `generateMetadata` 中的数据获取完成，然后再将 UI 流式传输到客户端。这保证了流式响应的第一部分包括 `<head>` 标签。
- 由于流式渲染是在服务器端进行的，因此不会影响 SEO。

## Streaming

在 Next.js 中，Suspense 被称为 Streaming，也就是将页面的 HTML 拆分成多个 chunks，然后逐步将这些块从服务端发送到客户端。 

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/GAWLA8-70e5c5b6-98c0-4084-8bce-da50c10eb334-20241209232950107.png)

这样就可以更快的展现出页面的某些内容，而无需在渲染 UI 之前等待加载所有数据。提前发送的组件可以提前开始水合，这样当其他部分还在加载的时候，用户可以和已完成水合的组件进行交互，有效改善用户体验。

Streaming 可以有效的阻止耗时长的数据请求阻塞整个页面加载的情况。它还可以减少加载[第一个字节所需时间（TTFB）](https://web.dev/articles/ttfb)和[首次内容绘制（FCP）](https://developer.chrome.com/docs/lighthouse/performance/first-contentful-paint)，有助于缩短[交互时间（TTI）](https://developer.chrome.com/docs/lighthouse/performance/interactive)，尤其在速度慢的设备上。

传统的 SSR 的输出执行过程：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/ITtx8p-c7846641-ca59-48fe-b406-d423bf72b39b-20241209232950088.png)

使用 Streaming 后的执行过程：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/SDWg6E-a4b5d741-9982-4332-bafc-f4f4d6b0b506-20241209232950098.png)

在 Next.js 中有两种实现 Streaming 的方法：使用页面级别 `Loading File` 和 `<Suspense>`。

> 推荐阅读文章：
> - [What are React Server Components? Understanding the Future of React Apps](https://www.builder.io/blog/why-react-server-components#suspense-for-server-side-rendering)

### Suspense 在 SSR 中的缺点：

尽管JavaScript代码可以异步流式传输到浏览器，但最终用户仍需下载整个网页的代码。随着应用程序功能的增加，用户需要下载的代码量也会随之增长。这引发了一个重要问题：**用户是否真的需要下载如此多的数据？**

当前的方式要求所有React组件都在客户端进行水合，无论这些组件是否真正需要交互功能。这种做法可能会浪费资源，并延长加载时间和用户可交互时间。用户设备需要处理和渲染可能并不需要客户端交互的组件。这引发了另一个问题：**是否所有组件都需要水合，即便它们不需要客户端交互？**

尽管服务器在处理密集计算任务方面能力更强，但大部分JavaScript的执行仍发生在用户设备上。对于性能较弱的设备，这会显著降低体验。这又引发了一个重要问题：**是否应该让如此多的工作在用户设备上完成？**
## 总结
通过 Loading UI 和 Streaming，Next.js 提供了更优雅的加载体验，显著优化了用户的感知性能。同时，这些技术有效减轻了服务器和客户端的压力，为开发者实现复杂 Web 应用提供了更强大的工具支持。未来，随着 Web 性能优化的进一步发展，这些技术将成为提升 Web 体验的重要手段。

『参考资料』
- [Loading UI and Streaming](https://nextjs.org/docs/app/building-your-application/routing/loading-ui-and-streaming)：https://nextjs.org/docs/app/building-your-application/routing/loading-ui-and-streaming
- [Hydrating server-rendered HTML](https://react.dev/reference/react-dom/client/hydrateRoot#hydrating-server-rendered-html)：https://react.dev/reference/react-dom/client/hydrateRoot#hydrating-server-rendered-html
- [Suspense](https://react.dev/reference/react/Suspense)：https://react.dev/reference/react/Suspense
- [Transfer-Encoding](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Transfer-Encoding)：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Transfer-Encoding