大家好，我是长林啊！一个 Go、Rust 爱好者，同时也是一名全栈开发者；致力于终生学习和技术分享。

在构建现代 Web 应用时，导航是连接用户界面的关键纽带。React Router 作为 React 生态中的核心路由库，为开发者提供了强大的工具来实现 SPA（单页应用）的导航逻辑。它不仅简化了页面间的跳转，还支持动态路由匹配、懒加载和状态管理集成，让应用的导航更加灵活和高效。

## 初识 React Router
React Router 是一个用于 React 应用程序的路由库，它允许你以声明式的方式来定义应用的导航结构。

### 介绍 React Router的重要性
React Router 的重要性在于它为构建单页应用（SPA）提供了一个强大而灵活的导航解决方案。以下是 React Router 的几个关键重要性点：

- **用户体验**：React Router 允许应用在不重新加载页面的情况下进行页面跳转，提供了无缝的用户体验。

- **应用结构**：它帮助开发者以组件化的方式组织应用的视图，使得应用的结构更加清晰和模块化。

- **动态路由**：React Router 支持动态路由，可以根据URL参数动态渲染组件，这在处理用户输入和API数据时非常有用。

- **导航控制**：提供了编程式导航和声明式导航的方式，使得开发者可以更灵活地控制应用的导航流程。

- **状态同步**：React Router 能够与 React 的状态管理库（如Redux或Context API）集成，同步路由状态与应用状态。

- **性能优化**：通过懒加载和代码分割，React Router 有助于提高应用的加载速度和运行效率。

- **SEO友好**：对于需要进行搜索引擎优化的应用，React Router 支持服务器端渲染，有助于提高SEO效果。

- **社区支持**：React Router 有着庞大的社区支持，提供了大量的教程、插件和第三方集成方案。

- **安全性**：React Router 提供了路由保护机制，可以防止未授权的路由访问，增强应用的安全性。

- **跨平台兼容性**：React Router 不仅限于 Web 应用，还可以与 React Native 等其他React平台集成，提供跨平台的导航解决方案。
### 概述 React Router 在现代Web应用中的作用
- **增强用户体验**：通过实现无缝页面跳转，React Router提升了用户交互的流畅性。

- **促进代码组织**：它通过组件化路由，帮助开发者以模块化的方式组织代码，提高应用的可维护性。

- **支持动态内容**：React Router允许根据URL动态加载内容，为构建数据驱动的应用提供了便利。

- **提高性能**：通过懒加载和代码分割，它有助于减少初始加载时间和提高应用性能。

- **保障安全性**：提供了路由保护功能，确保应用的导航逻辑安全且符合业务规则。

- **改善SEO**：支持服务器端渲染，有助于提高应用的搜索引擎优化效果。

- **灵活集成**：React Router可以与多种状态管理和UI库集成，提供一致的开发体验。

> 下文有不少示例演示，我们就用 vite 创建一个新的 React 项目吧！创建的命令如下：
> ```sh
> $ npm create vite@latest react-router-tutorial -- --template react
> ```
> 创建完成之后，用自己熟悉的 IDE 工具打开，并在终端中运行命令启动项目。

## 安装 React Router
可以选择自己熟悉的 Node.js 包管理工具，建议在同一个项目中只使用一种包管理工具，混合使用可能会导致一些包依赖出问题；建议直接使用 `pnpm` 包管理工具。
```sh
$ pnpm add react-router-dom
```
## 路由
React Router 提供了多种创建路由的方式，在 v6.4 又引入了4中新的创建路由的方式：
- `createBrowserRouter` 它使用 DOM History API 来更新 URL 并管理历史记录堆栈。
- `createMemoryRouter`
- `createHashRouter`
- `createStaticRouter`
当然原来声明式的创建路由的方式仍然还是保留了：
- `<BrowserRouter>`
- `<MemoryRouter>`
- `<HashRouter>`
- `<NativeRouter>`
- `<StaticRouter>`
这四种声明式的创建路由不支持 react-router 新增的一些 [Data](https://reactrouter.com/en/main/routers/picking-a-router#data-apis) 相关的 API 的使用，官方也建议所有 Web 项目使用 `createBrowserRouter` 的方式创建路由。

### 创建路由的方法详解
#### createBrowserRouter 的使用
它还支持 v6.4 数据 API，如 loaders、actions、fetchers 等。
```jsx
const router = createBrowserRouter([
  {
    path: "/",
    element: <ComponentA />,
    loader: rootLoader,
    children: [
      {
        path: "team",
        element: <ComponentB />,
        loader: teamLoader,
      },
    ],
  },
]);
```

`createBrowserRouter` 的类型:
```js
function createBrowserRouter (
    routes: RouteObject[],
    opts?: {
        basename?: string;
        future?: FutureConfig;
        hydrationData?: HydrationState;
        window?: Window;
    }
): RemixRouter;
```
- `routes`：Route 对象的数组，在 children 属性上有嵌套路由。
- `basename`：应用程序的基名，用于无法部署到域根目录而只能部署到子目录的情况。
  ```jsx
  createBrowserRouter(routes, {
      basename: "/app",
  });
  ```
- `future`：为路由器启用的一组可选的 Future Flags
  ```jsx
  const router = createBrowserRouter(routes, {
      future: {
          // Normalize `useNavigation()`/`useFetcher()` `formMethod` to uppercase
          v7_normalizeFormMethod: true,
      },
  });
  ```
  目前可用的 future flags 如下：
  | Flag	Description      | 说明                                                       |
  | ------------------------ | ---------------------------------------------------------- |
  | `v7_fetcherPersist`      | 延迟活动的 `fetcher` 清理，直到它们返回到 `idle` 状态      |
  | `v7_normalizeFormMethod` | 将 `useNavigation().formMethod` 规范化为大写的 `HTTP` 方法 |
  | `v7_partialHydration`    | 支持服务端渲染应用程序的部分水合功能                       |
  | `v7_prependBasename`     | 将路由的基名添加到 `navigate/fetch` 路径的前面             |
  | `v7_relativeSplatPath`   | 修复 `splat` 路由中相对路径解析的错误                      |
- `hydrationData`：在进行服务器渲染并选择退出自动水合时， `hydrationData` 选项允许您从服务器渲染器中传递水合数据。
  ```jsx
  const router = createBrowserRouter(routes, {
      hydrationData: {
          loaderData: {
              // [routeId]: serverLoaderData
          },
          // may also include `errors` and/or `actionData`
      },
  });
  ```
- `window`：对于浏览器 devtool 插件或测试等环境来说，使用与全局 window 不同的窗口非常有用。

#### createHashRouter
如果您无法配置 Web 服务器以将所有流量导向 React Router 应用程序，则此路由器非常有用。
> 不建议使用 Hash 路由！

功能与 createBrowserRouter 并无二致。
```jsx
const router = createHashRouter([
  {
    path: "/",
    element: <ComponentA />,
    loader: rootLoader,
    children: [
      {
        path: "team",
        element: <ComponentB />,
        loader: teamLoader,
      },
    ],
  },
]);
```
#### createMemoryRouter
内存路由器不使用浏览器的历史记录，而是在内存中管理自己的历史记录堆栈。它**主要用于测试和组件开发工具**（如 Storybook），但也可用于在任何非浏览器环境中运行 React Router。

```jsx
import * as React from "react";
import {
  RouterProvider,
  createMemoryRouter,
} from "react-router-dom";
import {
  render,
  waitFor,
  screen,
} from "@testing-library/react";
import "@testing-library/jest-dom";
import CalendarEvent from "./routes/event";

test("event route", async () => {
  const FAKE_EVENT = { name: "test event" };
  const routes = [
    {
      path: "/events/:id",
      element: <CalendarEvent />,
      loader: () => FAKE_EVENT,
    },
  ];

  const router = createMemoryRouter(routes, {
    initialEntries: ["/", "/events/123"], // 历史记录
    initialIndex: 1, // 初始化索引
  });

  render(<RouterProvider router={router} />);

  await waitFor(() => screen.getByRole("heading"));
  expect(screen.getByRole("heading")).toHaveTextContent(
    FAKE_EVENT.name
  );
});
```

参数除了 `initialIndex` 和 `initialEntries` 外其它参数与 `createBorwserRouter` 并无二致。
- `initialEntries`：历史记录堆栈中的初始条目。可以使用历史记录堆栈中已有的多个位置来启动测试（或应用）（用于测试后退导航等）
  ```jsx
  createMemoryRouter(routes, {
      initialEntries: ["/", "/events/123"],
  });
  ```
- `initialIndex`：历史堆栈中要呈现的初始索引。从特定条目开始测试。它默认为 中的最后一个条目 `initialEntries`。
  ```jsx
  createMemoryRouter(routes, {
    initialEntries: ["/", "/events/123"],
    initialIndex: 1, // start at "/events/123"
  });
  ```

#### createStaticHandler 
`createStaticHandler` 用于在服务器端通过 呈现应用程序之前在服务器（即 Node `<StaticRouterProvider>` 或其他 Javascript 运行时）上执行数据获取和提交。

```jsx
import {
  createStaticHandler,
  createStaticRouter,
  StaticRouterProvider,
} from "react-router-dom/server";

// ...

const routes = [
  {
    path: "/",
    loader: exampleLoader,
    Component: Root,
    ErrorBoundary: ComponentA,
  },
];

export async function renderHtml(req) {
  let { query, dataRoutes } = createStaticHandler(routes);
  let fetchRequest = createFetchRequest(req);
  let context = await query(fetchRequest);

  // If we got a redirect response, short circuit and let our Express server
  // handle that directly
  if (context instanceof Response) {
    throw context;
  }

  let router = createStaticRouter(dataRoutes, context);
  return ReactDOMServer.renderToString(
    <React.StrictMode>
      <StaticRouterProvider
        router={router}
        context={context}
      />
    </React.StrictMode>
  );
}
```
`routes` 和 `basename` 与 `createBrowserRouter` 是一样的，

- `handler.query()` 方法接受 Fetch 请求，执行路由匹配，并根据请求执行所有相关的路由 `action/loader` 方法，返回context值包含呈现请求的 HTML 文档所需的所有信息（路由级别 `actionData`、`loaderData`、`errors` 等）。如果任何匹配的路由返回或抛出重定向响应，`query()` 则将以 Fetch 的形式返回该重定向 `Response`。如果请求被中止，`query` 将抛出错误，例如 `Error("query() call aborted: GET /path")`。如果你想抛出本机 `AbortSignal.reason`（默认情况下为 `DOMException`），你可以选择加入 `future.v7_throwAbortReason` 未来标志。

  > DOMException 是在 Node 17 中添加的，因此你必须在 Node 17 或更高版本上才能正常工作。

- `opts.requestContext` 如果您需要将信息从服务器传递到 Remix action/loader，您可以使用它来执行此操作，`opts.requestContext` 它将显示在上下文参数中的操作/加载器中。
  ```jsx
  const routes = [{
    path: '/',
    loader({ request, context }) {
      // Access `context.dataFormExpressMiddleware` here
    },
  }];
  
  export async function render(req: express.Request) {
    let { query, dataRoutes } = createStaticHandler(routes);
    let remixRequest = createFetchRequest(request);
    let staticHandlerContext = await query(remixRequest, {
      // Pass data from the express layer to the remix layer here
      requestContext: {
        dataFromExpressMiddleware: req.something
      }
   });
   ...
  }
  ```
  
- `opts.routeId`

  如果你需要调用一个与 URL 不完全对应的特定路由操作/加载器（例如，父路由加载器），你可以指定`routeId`：
  ```jsx
  staticHandler.queryRoute(new Request("/parent/child"), {
    routeId: "parent",
  });
  ```
  
- `opts.requestContext`

  如果您需要将信息从服务器传递到 Remix `action/loader` 中，您可以使用 进行传递，`opts.requestContext` 它将显示在上下文参数中的 `action/loader` 中。


#### createStaticRouter

`createStaticRouter` 是利用数据路由器在服务器（即Node或其他 Javascript 运行时）上进行渲染时，可以使用它。

```jsx
import {
  createStaticHandler,
  createStaticRouter,
  StaticRouterProvider,
} from "react-router-dom/server";

// ...

const routes = [
  {
    path: "/",
    loader: exampleLoader,
    Component: ComponentA,
    ErrorBoundary: RootErrorBoundary,
  },
];

export async function renderHtml(req) {
  let { query, dataRoutes } = createStaticHandler(routes);
  let fetchRequest = createFetchRequest(req);
  let context = await query(fetchRequest);

  // If we got a redirect response, short circuit and let our Express server
  // handle that directly
  if (context instanceof Response) {
    throw context;
  }

  let router = createStaticRouter(dataRoutes, context);
  return ReactDOMServer.renderToString(
    <React.StrictMode>
      <StaticRouterProvider
        router={router}
        context={context}
      />
    </React.StrictMode>
  );
}
```

#### RouterProvider
`RouterProvider` 是 react-router-dom 提供的一个组件，用于将路由配置传递给整个应用，使得应用中的所有组件都可以访问和使用这些路由信息。

主要功能：
- 提供路由上下文：`RouterProvider` 创建并提供一个路由上下文，使应用中的任何组件都可以方便地访问路由信息和导航功能。
- 管理路由状态：它负责管理路由的状态和更新，包括当前路径、导航历史等。
- 加载数据：结合路由加载器 (loader)，`RouterProvider` 可以在路由切换时预加载数据，提高应用的性能和用户体验。

以下是一个完整的示例，展示了如何使用 `RouterProvider` 配置和提供路由：
```jsx
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

// 定义组件
function Home() {
  return <h2>Home Page</h2>;
}

function About() {
  return <h2>About Page</h2>;
}

// 创建路由实例
const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/about",
    element: <About />,
  },
]);

// 提供路由上下文
ReactDOM.createRoot(document.getElementById('root')).render(
  <RouterProvider router={router} />
);
```

除了在上面的示例中用到的 `router` 参数外，还有 `fallbaclElement`，它用于在路由加载过程中显示一个备用的 UI 元素。这在需要加载数据或组件时非常有用，可以提供一个良好的用户体验，例如显示一个加载指示器或占位符，直到实际的内容加载完毕。例如：
```jsx
<RouterProvider
  router={router}
  fallbackElement={<SpinnerOfDoom />}
/>
```

#### StaticRouterProvider

`StaticRouterProvider` 是 react-router-dom 提供的一个组件，主要用于在服务器端渲染（SSR）环境中进行路由配置。与客户端渲染不同，服务器端渲染需要在服务器上执行路由匹配和组件渲染，然后将渲染好的 HTML 发送到客户端。


**主要用途**：
- 服务器端渲染（SSR）：StaticRouterProvider 适合在服务器端环境中使用，配合 React 的服务器端渲染功能，实现同构应用。
- 静态路由配置：它使用静态路由配置，不依赖于浏览器的历史记录和导航功能。
**主要特性**：
- 静态上下文：`StaticRouterProvider` 使用静态上下文进行路由匹配，这在服务器端环境中是必要的，因为没有浏览器的历史记录或导航功能。
- 便于 SSR：它简化了在服务器端进行路由匹配和渲染的过程。

```jsx
import {
  createStaticHandler,
  createStaticRouter,
  StaticRouterProvider,
} from "react-router-dom/server";

const routes = [
  {
    path: "/",
    loader: rootLoader,
    Component: Root,
    ErrorBoundary: RootBoundary,
  },
];

export async function renderHtml(req) {
  let { query, dataRoutes } = createStaticHandler(routes);
  let fetchRequest = createFetchRequest(req);
  let context = await query(fetchRequest);

  // If we got a redirect response, short circuit and let our Express server
  // handle that directly
  if (context instanceof Response) {
    throw context;
  }

  let router = createStaticRouter(dataRoutes, context);
  return ReactDOMServer.renderToString(
    <React.StrictMode>
      <StaticRouterProvider
        router={router}
        context={context}
      />
    </React.StrictMode>
  );
}
```

这个方法的主要参数：
- `context`：从createStaticHandler().query()调用返回的内容，其中包含了请求所获取的所有数据。
- `router`：这是通过以下方式创建的路由器 `createStaticRouter`
- `hydrate`：默认情况下，`<StaticRouterProvider>` 将把所需的水合数据字符串化到标签 `window.__staticRouterHydrationData` 上` <script>`，该标签将被读取并自动水合 `createBrowserRouter()`。如果希望手动进行更高级的水合，您可以通过 来 `hydrate={false}` 禁用此自动水合。在客户端，您可以将自己的水合传递 `hydrationData` 给 `createBrowserRouter`。
- `nonce`：当利用自动补水时，您可以提供一个 `nonce` 要呈现到 `<script>` 标签上并与内容安全策略一起使用的值。
  
### 路由组件

#### BrowserRouter（浏览器组件）

将当前位置存储在浏览器的地址栏中，并使用浏览器的内置历史记录堆栈进行导航。
```jsx
declare function BrowserRouter(
    props: BrowserRouterProps
): React.ReactElement;

interface BrowserRouterProps {
    basename?: string;
    children?: React.ReactNode;
    future?: FutureConfig;
    window?: Window;
}
```

**参数解释**：
- `window`：`BrowserRouter` 默认使用当前文档的 `defaultView`，但它也可用于跟踪另一个窗口的 URL 的变化，例如 `<iframe>`。
- `future`：一组可选的未来标志可供启用。
  ```jsx
  function App() {
    return (
      <BrowserRouter future={{ v7_startTransition: true }}>
        <Routes>{/*...*/}</Routes>
      </BrowserRouter>
    );
  }
  ```
- `basename`:一个用于**指定基础路径**的属性。它主要用于当你的应用部署在一个子路径（而不是根路径）时，确保路由能够正确匹配和导航。

  **主要用途**：
  1. 支持子路径部署：当你的应用部署在非根路径（例如 https://example.com/myapp）时，basename 可以帮助你的应用正确处理路由。
  2. 确保 URL 正确：通过设置 `basename`，你可以确保所有链接和导航行为都基于指定的基础路径。
  
  ```jsx
  import * as React from "react";
  import * as ReactDOM from "react-dom";
  import { createBrowserRouter, RouterProvider, Link, Outlet } from 'react-router-dom';
  
  // 定义组件
  function Home() {
    return <h2>Home Page</h2>;
  }
  
  function About() {
    return <h2>About Page</h2>;
  }
  
  function App() {
    return (
      <div>
        <nav>
          <Link to="/">Home</Link>
          <Link to="/about">About</Link>
        </nav>
        <Outlet />
      </div>
    );
  }
  
  // 创建路由器实例并设置 basename
  const router = createBrowserRouter(
    [
      {
        path: "/",
        element: <App />,
        children: [
          {
            path: "/",
            element: <Home />,
          },
          {
            path: "/about",
            element: <About />,
          },
        ],
      },
    ],
    {
      basename: "/myapp",  // 设置基础路径
    }
  );
  
  // 渲染应用
  ReactDOM.createRoot(document.getElementById("root")).render(
    <RouterProvider router={router} />
  );
  ```
  在这个示例中，假设你的应用部署在 https://example.com/myapp，设置了 basename: "/myapp" 后，Link 组件和路由匹配都会基于这个基础路径进行。导航到 /about 实际上会导航到 https://example.com/myapp/about。
  
#### HashRouter（哈希路由）
在 react-router-dom 中，`HashRouter` 是一种用于管理路由的组件，它使用 URL 的哈希部分（即 # 后面的部分）来保持 UI 与 URL 同步。`HashRouter` 适用于那些服务器端不处理路由的应用，例如静态文件服务，或在某些需要兼容旧版本浏览器的场景中。

```ts
declare function HashRouter(
  props: HashRouterProps
): React.ReactElement;

interface HashRouterProps {
  basename?: string;
  children?: React.ReactNode;
  future?: FutureConfig;
  window?: Window;
}
```

**主要用途**：
- 兼容性：`HashRouter` 可以在不支持 HTML5 历史记录 API 的旧版本浏览器中使用。
- 简单部署：当你的服务器不处理路由时，`HashRouter` 是一个简便的解决方案，因为它不需要服务器端的配置来支持不同的 URL。

**基本特性**：
- 使用 URL 哈希部分进行路由。
- URL 形式为 http://example.com/#/your/path。
- 哈希部分的变化不会触发服务器请求，仅会触发客户端的路由变化。
  
```jsx
import { HashRouter, Route, Routes, Link } from 'react-router-dom';

// 定义组件
function Home() {
  return <h2>Home Page</h2>;
}

function About() {
  return <h2>About Page</h2>;
}

function App() {
  return (
    <div>
      <nav>
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </div>
  );
}

// 使用 HashRouter 包裹应用
ReactDOM.createRoot(document.getElementById('root')).render(
  <HashRouter>
    <App />
  </HashRouter>
);
```
在这个示例中，`HashRouter` 使用 URL 的哈希部分进行路由，所以当你点击导航链接时，URL 会变为 http://example.com/#/about 或 http://example.com/#/，而不需要服务器进行任何配置或处理。

#### MemoryRouter
`<MemoryRouter>` 将其位置**存储在数组内部**。与 `<BrowserHistory>` 和 `<HashHistory>`  不同它不依赖于外部源，例如浏览器中的历史记录堆栈。这使其成为需要完全控制历史记录堆栈的场景（例如测试）的理想选择。

```jsx
declare function MemoryRouter(
  props: MemoryRouterProps
): React.ReactElement;

interface MemoryRouterProps {
  basename?: string;
  children?: React.ReactNode;
  initialEntries?: InitialEntry[];
  initialIndex?: number;
  future?: FutureConfig;
}
```

这个组件一般用于写单元测试，比如：
```jsx
import * as React from "react";
import { create } from "react-test-renderer";
import {
  MemoryRouter,
  Routes,
  Route,
} from "react-router-dom";

describe("My app", () => {
  it("renders correctly", () => {
    let renderer = create(
      <MemoryRouter initialEntries={["/users/mjackson"]}>
        <Routes>
          <Route path="users" element={<Users />}>
            <Route path=":id" element={<UserProfile />} />
          </Route>
        </Routes>
      </MemoryRouter>
    );

    expect(renderer.toJSON()).toMatchSnapshot();
  });
});
```
NativeRouter 组件实在 React Native 中使用的路由导航的工具，就不在本系列文章中涉及；因为 React Native 也是一个比较大的概念。


## Route

### Route 组件
Route 是 react-router-dom 中的一个核心组件，用于定义应用中的各个路由及其关联的组件。它的主要作用是根据当前的 URL 匹配相应的路径，并渲染对应的组件。

**主要特性**：
- 路径匹配：Route 组件根据指定的路径（path）来匹配 URL。
- 组件渲染：当路径匹配成功时，渲染对应的组件。
- 嵌套路由：支持嵌套路由，通过 Outlet 组件实现子路由的嵌套渲染。

**基本用法**：
创建 Route 有两种方式，一种是通过 createBrowserRouter 的方式创建，另一种是通过声明式的方式创建，如下：
- createBrowserRouter 函数的方式创建
  ```ts
  const router = createBrowserRouter([
    {
      // it renders this element
      element: <Team />,
  
      // when the URL matches this segment
      path: "teams/:teamId",
  
      // with this data loaded before rendering
      loader: async ({ request, params }) => {
        return fetch(
          `/fake/api/teams/${params.teamId}.json`,
          { signal: request.signal }
        );
      },
  
      // performing this mutation when data is submitted to it
      action: async ({ request }) => {
        return updateFakeTeam(await request.formData());
      },
  
      // and renders this element in case something went wrong
      errorElement: <ErrorBoundary />,
    },
  ]);
  ```
- 声明式创建
  ```jsx
  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route
        element={<Team />}
        path="teams/:teamId"
        loader={async ({ params }) => {
          return fetch(
            `/fake/api/teams/${params.teamId}.json`
          );
        }}
        action={async ({ request }) => {
          return updateFakeTeam(await request.formData());
        }}
        errorElement={<ErrorBoundary />}
      />
    )
  );
  ```

Route 组件的参数说明：
```ts
interface RouteObject {
  path?: string;
  index?: boolean;
  children?: RouteObject[];
  caseSensitive?: boolean;
  id?: string;
  loader?: LoaderFunction;
  action?: ActionFunction;
  element?: React.ReactNode | null;
  hydrateFallbackElement?: React.ReactNode | null;
  errorElement?: React.ReactNode | null;
  Component?: React.ComponentType | null;
  HydrateFallback?: React.ComponentType | null;
  ErrorBoundary?: React.ComponentType | null;
  handle?: RouteObject["handle"];
  shouldRevalidate?: ShouldRevalidateFunction;
  lazy?: LazyRouteFunction<RouteObject>;
}
```
- path：与 URL 匹配的路径模式以确定此路由是否匹配 URL、链接 href 或表单操作。
- index：通俗理解就是，被标记为 index 的路由就是默认子路由。
  ```jsx
  <Route path="/teams" element={<Teams />}>
      <Route index element={<TeamsIndex />} />
      <Route path=":teamId" element={<Team />} />
  </Route>
  ```
- children：用于嵌套路由的场景。
- caseSensitive：是否区分大小写。在组件中添加此属性则为严格匹配：
  ```jsx
  <Route caseSensitive path="/wEll-aCtuA11y" /> // 不匹配 "well-actua11y"
  ```
- loader：路由渲染之前被调用的回调，并且通过 `useLoaderData` 给路由提供数据。
  ```js
  <Route
    path="/teams/:teamId"
    loader={({ params }) => {
      return fetchTeam(params.teamId);
    }}
  />;
  
  function Team() {
    let team = useLoaderData();
    // ...
  }
  ```
- action：当从 form、fetcher 或 submission 将提交发送到路由时，就调用路由操作。
  ```jsx
  <Route
    path="/teams/:teamId"
    action={({ request }) => {
      const formData = await request.formData();
      return updateTeam(formData);
    }}
  />
  ```
- element/Component
  如果要创建 React 元素，请使用element：
  ```jsx
  <Route path="/for-sale" element={<Properties />} />
  ```
  否则使用ComponentReact Router 将为您创建 React Element：
  ```jsx
  <Route path="/for-sale" Component={Properties} />
  ```
- errorElement/ErrorBoundary
  当路由在 `loader` 或 `action` 中渲染时抛出异常时，该 组件将代替正常组件进行渲染。
  ```jsx
  <Route
      path="/for-sale"
      // if this throws an error while rendering
      element={<Properties />}
      // or this while loading properties
      loader={() => loadProperties()}
      // or this while creating a property
      action={async ({ request }) =>
        createProperty(await request.formData())
      }
      // then this element will render
      errorElement={<ErrorBoundary />}
  />
  ```
  
- hydrateFallbackElement/HydrateFallback

  如果您正在使用服务器端渲染并且正在利用部分水合，那么您可以在应用程序的初始水合期间指定要为非水合路线渲染的元素/组件。
  
- handle

  任何特定于应用程序的数据。
  
- lazy

  为了使您的应用程序包保持较小并且支持路由的代码拆分，每个路由都可以提供一个异步函数来解析路由定义中与路由不匹配的部分（`loader`、`action`、`Component/element`、`ErrorBoundary/errorElement`等）。
  
  每个lazy函数通常会返回动态导入的结果:
  ```jsx
  let routes = createRoutesFromElements(
    <Route path="/" element={<Layout />}>
      <Route path="a" lazy={() => import("./a")} />
      <Route path="b" lazy={() => import("./b")} />
    </Route>
  );
  ```
  如果是在惰性路由模块中，导出想要为路由定义的属性：
  ```jsx
  export async function loader({ request }) {
    let data = await fetchData(request);
    return json(data);
  }
  
  export function Component() {
    let data = useLoaderData();
  
    return (
      <>
        <h1>You made it!</h1>
        <p>{data}</p>
      </>
    );
  }
  ```
### action

action 属性允许你在路径匹配时定义一个处理函数，用于处理表单提交等动作，并返回一个处理结果。

**主要用途**：
- 处理表单提交：当表单提交到某个路由时，action 可以处理提交的数据。
- 处理用户交互：处理用户在特定路由上的交互，比如按钮点击等。
每当应用向您的路线发送非获取提交（`post`、`put`、`patch`、`delete`）时，就会调用操作。这可能以几种方式发生：

```jsx
// forms
<Form method="post" action="/songs" />;
<fetcher.Form method="put" action="/songs/123/edit" />;

// imperative submissions
let submit = useSubmit();
submit(data, {
  method: "delete",
  action: "/songs/123",
});
fetcher.submit(data, {
  method: "patch",
  action: "/songs/123/edit",
});
```
- params 解析动态路由的参数，比如下面示例的 `projectId`
  ```jsx
  <Route
      path="/projects/:projectId/delete"
      action={({ params }) => {
          return fakeDeleteProject(params.projectId);
      }}
  />
  ```
- request 从路由中获取请求实例，最常见的用例是从请求中解析 `FormData`:
  ```jsx
  <Route
      action={async ({ request }) => {
          let formData = await request.formData();
          // ...
      }}
  />
  ```
  
#### 抛出错误：
```jsx
<Route
  action={async ({ params, request }) => {
    const res = await fetch(
      `/api/properties/${params.id}`,
      {
        method: "put",
        body: await request.formData(),
      }
    );
    if (!res.ok) throw res;
    return { ok: true };
  }}
/>
```

#### 根据不同的行为进行处理
```jsx
async function action({ request }) {
  let formData = await request.formData();
  let intent = formData.get("intent");

  if (intent === "edit") {
    await editSong(formData);
    return { ok: true };
  }

  if (intent === "add") {
    await addSong(formData);
    return { ok: true };
  }

  throw json(
    { message: "Invalid intent" },
    { status: 400 }
  );
}

function Component() {
  let song = useLoaderData();

  // When the song exists, show an edit form
  if (song) {
    return (
      <Form method="post">
        <p>Edit song lyrics:</p>
        {/* Edit song inputs */}
        <button type="submit" name="intent" value="edit">
          Edit
        </button>
      </Form>
    );
  }

  // Otherwise show a form to add a new song
  return (
    <Form method="post">
      <p>Add new lyrics:</p>
      {/* Add song inputs */}
      <button type="submit" name="intent" value="add">
        Add
      </button>
    </Form>
  );
}
```
这段代码展示了如何使用 action 属性处理表单提交，并根据提交的“意图”执行不同的操作（编辑或添加歌曲）。同时，组件根据数据的存在与否渲染相应的表单，提供了一个动态的用户界面。


以下是一个更详细的示例，展示了如何使用 action 属性处理表单提交，并在提交后进行重定向。
```jsx
import React from 'react';
import { createBrowserRouter, RouterProvider, Route, Form, redirect } from 'react-router-dom';

// 定义 action 函数
const action = async ({ request }) => {
  const formData = await request.formData();
  const username = formData.get('username');
  // 处理提交的数据，比如保存到数据库
  console.log('Username:', username);
  // 返回一个重定向
  return redirect('/success');
};

// 定义表单组件
function FormComponent() {
  return (
    <Form method="post">
      <label>
        Username:
        <input type="text" name="username" />
      </label>
      <button type="submit">Submit</button>
    </Form>
  );
}

// 定义成功页面组件
function Success() {
  return <h2>Form submitted successfully!</h2>;
}

// 创建路由器
const router = createBrowserRouter([
  {
    path: '/',
    element: <FormComponent />,
    action: action,
  },
  {
    path: '/success',
    element: <Success />,  
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
```
### errorElement
在加载器（`loaders`）、动作（`actions`）或组件渲染过程中抛出异常时，路由（`Routes`）将不会走正常的渲染路径（即 `<Route element>`），而是会走错误路径（即 `<Route errorElement>`），并且可以通过 `useRouteError` 获取到该错误。

```jsx
<Route
  path="/invoices/:id"
  // if an exception is thrown here
  loader={loadInvoice}
  // here
  action={updateInvoice}
  // or here
  element={<Invoice />}
  // this will render instead of `element`
  errorElement={<ErrorBoundary />}
/>;

function Invoice() {
  return <div>Happy {path}</div>;
}

function ErrorBoundary() {
  let error = useRouteError();
  console.error(error);
  // Uncaught ReferenceError: path is not defined
  return <div>Dang!</div>;
}
```
#### 冒泡

当路由没有 时 `errorElement`，错误将通过父路由冒泡。在你的路由树顶层放置一个 `errorElement`，你就可以在一个地方处理应用中的几乎所有错误。或者你也可以在每个路由上都放置 `errorElement`，

#### 手动抛出异常
在errorElement处理意外错误的同时，它还可以用来处理预判到的异常。特别是在 loader 和action 中，执行结果是不受你控制的外部数据，数据可能不存在、服务补可用或用户无权访问它。等等，都可以自定义异常并抛出。

```jsx
<Route
  path="/properties/:id"
  element={<PropertyForSale />}
  errorElement={<PropertyError />}
  loader={async ({ params }) => {
    const res = await fetch(`/api/properties/${params.id}`);
    if (res.status === 404) {
      throw new Response("Not Found", { status: 404 });
    }
    const home = await res.json();
    const descriptionHtml = parseMarkdown(
      data.descriptionMarkdown
    );
    return { home, descriptionHtml };
  }}
/>
```

#### 捕获异常
所有抛出的自定义异常都会通过 `useRouteError` 获取到，但如果你抛出一个 `Response` 的异常，React Router 会在返回给你的组件渲染之前自动解析响应数据。

然后，可以通过 `isRouteErrorResponse` 检查这种特定类型的异常。结合 react-router-dom 包的 `json`,可以根据相关信息处理其边界情况。
```jsx
import { json } from "react-router-dom";

function loader() {
  const stillWorksHere = await userStillWorksHere();
  if (!stillWorksHere) {
    throw json(
      {
        sorry: "You have been fired.",
        hrEmail: "hr@bigco.com",
      },
      { status: 401 } // 状态码为 401
    );
  }
}

function ErrorBoundary() {
  const error = useRouteError();

  // 如果是 response 错误且是 401 则渲染特殊的 DOM
  if (isRouteErrorResponse(error) && error.status === 401) {
    // the response json is automatically parsed to
    // `error.data`, you also have access to the status
    return (
      <div>
        <h1>{error.status}</h1>
        <h2>{error.data.sorry}</h2>
        <p>
          Go ahead and email {error.data.hrEmail} if you
          feel like this is a mistake.
        </p>
      </div>
    );
  }

  // rethrow to let the parent error boundary handle it
  // when it's not a special case for this route
  throw error;
}
```

有了 `isRouteErrorResponse` 后，则可以在根路由上创建一个通用的错误边界处理其常见的错误问题：
```jsx
function RootBoundary() {
  const error = useRouteError();

  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      return <div>This page doesn't exist!</div>;
    }

    if (error.status === 401) {
      return <div>You aren't authorized to see this</div>;
    }

    if (error.status === 503) {
      return <div>Looks like our API is down</div>;
    }

    if (error.status === 418) {
      return <div>🫖</div>;
    }
  }

  return <div>Something went wrong</div>;
}
```

### hydrateFallbackElement
它允许你在服务器端渲染（SSR）时提供一个备用的 React 元素，当服务器端渲染的页面在客户端进行水合（hydration）时，如果遇到某些组件或元素无法在服务器端渲染，就会使用这个备用元素。

```jsx
let router = createBrowserRouter(
  [
    {
      id: "root",
      path: "/",
      loader: rootLoader,
      Component: Root,
      children: [
        {
          id: "invoice",
          path: "invoices/:id",
          loader: loadInvoice,
          Component: Invoice,
          HydrateFallback: InvoiceSkeleton,
        },
      ],
    },
  ],
  {
    future: {
      v7_partialHydration: true,
    },
    hydrationData: {
      root: {
        /*...*/
      },
      // No hydration data provided for the `invoice` route
    },
  }
);
```

## 组件
### Await
await 组件是一个在 React Router v6 引入的新特性。它的作用是处理路由加载时的异步逻辑，比如数据获取或懒加载组件。

```jsx
import {
  defer,
  Route,
  useLoaderData,
  Await,
} from "react-router-dom";

// given this route
<Route
  loader={async () => {
    let book = await getBook();
    let reviews = getReviews(); // not awaited
    return defer({
      book,
      reviews, // this is a promise
    });
  }}
  element={<Book />}
/>;

function Book() {
  const {
    book,
    reviews, // this is the same promise
  } = useLoaderData();
  return (
    <div>
      <h1>{book.title}</h1>
      <p>{book.description}</p>
      <React.Suspense fallback={<ReviewsSkeleton />}>
        <Await
          // and is the promise we pass to Await
          resolve={reviews}
        >
          <Reviews />
        </Await>
      </React.Suspense>
    </div>
  );
}
```

### Form
Form 组件是纯 HTML表单的包装器，用于模拟浏览器以进行客户端路由和数据变更。
```jsx
import { Form } from "react-router-dom";

function NewEvent() {
  return (
    <Form method="post" action="/events">
      <input type="text" name="title" />
      <input type="text" name="description" />
      <button type="submit">Create</button>
    </Form>
  );
}
```
- action 表单提交到的 URL，与HTML 表单操作类似。唯一的区别在于默认操作。对于 HTML 表单，默认为完整 URL。对于 `<Form>`，默认为上下文中最接近的路由的相对 URL。
  ```jsx
  function ProjectsLayout() {
    return (
      <>
        <Form method="post" />
        <Outlet />
      </>
    );
  }
  
  function ProjectsPage() {
    return <Form method="post" />;
  }
  
  <DataBrowserRouter>
    <Route
      path="/projects"
      element={<ProjectsLayout />}
      action={ProjectsLayout.action}
    >
      <Route
        path=":projectId"
        element={<ProjectsPage />}
        action={ProjectsPage.action}
      />
    </Route>
  </DataBrowserRouter>;
  ```
  
- method 执行请求的方法，除了 `get` 和 `post` 外，还支持 `put`、`patch` 和 `delete`；默认 `get`。

- navigate：可以指定 `<Form navigate={false}>` ，让表单跳过导航，在内部使用 `fetcher`。这基本上是 `useFetcher() + <fetcher.Form>` 的简写。

- fetcherKey：在使用非导航 Form 时，也可选择通过 `<Form navigate={false} fetcherKey="my-key">` 指定自己的取值器密钥。

- replace：替换当前路由
  ```jsx
  <Form replace />
  ```
- relative：默认情况下，路径是相对于路由层次结构而言的，因此 `..` 会向上移动一级 `Route`。有时，你可能会发现有一些匹配的 URL 模式没有嵌套的意义，这时你更愿意使用相对路径路由。
  ```jsx
  <Form to="../some/where" relative="path">
  ```
  
- reloadDocument：指示表单跳过 React Router，使用浏览器内置行为提交表单。
> 这个方法一般用于 remix 框架，否则还是乖乖的使用原生的 form 标签吧！
> ```jsx
> <Form reloadDocument />
> ```
> 建议使用 `<form>` ，这样可以获得默认和相对 `action` 的好处，除此之外，它与普通 HTML 表单相同。

- state：可用于为存储在历史状态中的新位置设置一个有状态的值。
  ```jsx
  <Form
      method="post"
      action="new-path"
      state={{ some: "value" }}
  />
  ```
  可以在 "新路径 "路由上访问该状态值：
  ```jsx
  const { state } = useLocation();
  ```
  
- preventScrollReset：如果使用 `<ScrollRestoration>`，则可以防止在表单操作重定向到新位置时，滚动位置被重置到窗口顶部。
  ```jsx
  <Form method="post" preventScrollReset={true} />
  ```
  
- unstable_viewTransition：该属性通过在 [`document.startViewTransition()`](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/startViewTransition) 中封装最终状态更新，为该导航启用了视图转换。如果需要为该视图转换应用特定样式，还需要利用 [`unstable_useViewTransitionState()`](https://reactrouter.com/en/main/hooks//use-view-transition-state)
  > 该属性是一个实验性的 API，因为原生的 `startViewTransition` 是一个实验性的 API。

### Link

这个组件分为 RN 版本的 WEB 版本，我们这系列着看 WEB 版本。

Link 就是一个能跳转的 `<a>` 标签，用户可以通过点击或轻点它来导航到另一个页面。在 react-router-dom 中， `<Link>` 会渲染一个可访问的 `<a>` 元素，该元素带有一个真正的 `href`，指向它所链接的资源。这

```jsx
declare function Link(props: LinkProps): React.ReactElement;

interface LinkProps
  extends Omit<
    React.AnchorHTMLAttributes<HTMLAnchorElement>,
    "href"
  > {
  to: To;
  preventScrollReset?: boolean;
  relative?: "route" | "path";
  reloadDocument?: boolean;
  replace?: boolean;
  state?: any;
  unstable_viewTransition?: boolean;
}

type To = string | Partial<Path>;

interface Path {
  pathname: string;
  search: string;
  hash: string;
}
```
简单示例如下：
```jsx
import * as React from "react";
import { Link } from "react-router-dom";

function UsersIndexPage({ users }) {
  return (
    <div>
      <h1>Users</h1>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            <Link to={user.id}>{user.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
```
相对 `<Link to>` 值（不以 `/` 开头）是相对于父路由解析的，这意味着它建立在渲染该 `<Link>` 的路由所匹配的 URL 路径之上。它可能包含 `..` ，以链接到层级更高的路由。在这种情况下， `..` 的工作原理与命令行函数 cd 完全相同；每 `..` 删除父路径中的一段。

> 当当前 URL 以 `/` 结尾时，带有 `..` 的 `<Link to>` 与普通 `<a href>` 的行为不同。 `<Link to>` 会忽略尾部斜线，并为每个 `..` 删除一个 URL 段。但是，当当前 URL 以 `/` 结尾时， `<a href>` 值处理 `..` 的方式与其不同。

此组件的的参数跟 Form 中的参数差不多，这里就不再介绍了。

### NavLink
`<NavLink>` 是一种特殊的 `<Link>` ，它知道自己是否处于 "激活"、"待定 "或 "过渡 "状态。这在几种不同的情况下都很有用：

- 在创建导航菜单（如面包屑或一组选项卡）时，您希望显示当前选择了哪个选项卡
- 它为屏幕阅读器等辅助技术提供了有用的背景信息
- 它提供了一个 "过渡 "值，可让您对视图转换进行更精细的控制

```jsx
import { NavLink } from "react-router-dom";

<NavLink
  to="/messages"
  className={({ isActive, isPending }) =>
    isPending ? "pending" : isActive ? "active" : ""
  }
>
  Messages
</NavLink>;
```

#### 默认 active 类
默认情况下，当组件处于活动状态时，active会向其添加一个类，以便您可以使用 CSS 来设置其样式。
```jsx
<nav id="sidebar">
  <NavLink to="/messages" />
</nav>
```
```css
#sidebar a.active {
    color: red;
}
```

#### className
该 `className` prop 的工作方式与普通的 `className` 类似，也可以向其传递一个函数，以根据链接的活动状态和待处理状态自定义所应用的 `className`。
```jsx
<NavLink
  to="/messages"
  className={({ isActive, isPending, isTransitioning }) =>
    [
      isPending ? "pending" : "",
      isActive ? "active" : "",
      isTransitioning ? "transitioning" : "",
    ].join(" ")
  }
>
  Messages
</NavLink>
```
#### style
`style` 属性的工作方式与普通样式属性类似，也可以通过一个函数，根据链接的活动和待定状态自定义应用的样式。

```jsx
<NavLink
  to="/messages"
  style={({ isActive, isPending, isTransitioning }) => {
    return {
      fontWeight: isActive ? "bold" : "",
      color: isPending ? "red" : "black",
      viewTransitionName: isTransitioning ? "slide" : "",
    };
  }}
>
  Messages
</NavLink>
```
#### children
可以传递一个呈现属性作为子元素，以便根据活动和待定状态自定义 `<NavLink>` 的内容，这对更改内部元素的样式非常有用。

```jsx
<NavLink to="/tasks">
  {({ isActive, isPending }) => (
    <span className={isActive ? "active" : ""}>Tasks</span>
  )}
</NavLink>
```

#### end
`end` 属性更改了 `active` 和 `pending` 状态的匹配逻辑，使其只匹配到导航链接 `to` 路径的 "末端"。如果 URL 长于 to ，将不再被视为激活状态。

| Link                           | Current URL | isActive |
| ------------------------------ | ----------- | -------- |
| `<NavLink to="/tasks" />`      | / tasks     | true     |
| `<NavLink to="/tasks" /> `     | /tasks/123  | true     |
| `<NavLink to="/tasks" end />`  | /tasks      | true     |
| `<NavLink to="/tasks" end />`  | /tasks/123  | false    |
| `<NavLink to="/tasks/" end />` | /tasks      | false    |
| `<NavLink to="/tasks/" end />` | /tasks/     | true     |

关于根路由链接的说明

`<NavLink to="/">` 是一个特例，因为每个 URL 都匹配 `/` 。为了避免默认情况下每条路由都匹配，它实际上忽略了 `end` 属性，只在根路由上匹配。

#### caseSensitive
添加 `caseSensitive` 属性后，匹配逻辑会发生变化，变得区分大小写。

| Link                                         | URL         | isActive |
| -------------------------------------------- | ----------- | -------- |
| `<NavLink to="/SpOnGe-bOB" />`               | /sponge-bob | true     |
| `<NavLink to="/SpOnGe-bOB" caseSensitive />` | /sponge-bob | false    |

#### reloadDocument 此属性可用于跳过客户端路由，让浏览器正常处理转换（如同 `<a href>` ）。

#### `unstable_viewTransition` 用以动画过渡场景下的，现在还是一个实验特性。谨慎使用！ 

### Navigate
`<Navigate>` 元素在渲染时会改变当前位置。它是 `useNavigate` 的组件包装器，并接受与 `props` 相同的参数。

```jsx
declare function Navigate(props: NavigateProps): null;

interface NavigateProps {
  to: To;
  replace?: boolean;
  state?: any;
  relative?: RelativeRoutingType;
}
```

示例如下：
```jsx
import * as React from "react";
import { Navigate } from "react-router-dom";

class LoginForm extends React.Component {
  state = { user: null, error: null };

  async handleSubmit(event) {
    event.preventDefault();
    try {
      let user = await login(event.target);
      this.setState({ user });
    } catch (error) {
      this.setState({ error });
    }
  }

  render() {
    let { user, error } = this.state;
    return (
      <div>
        {error && <p>{error.message}</p>}
        {user && (
          <Navigate to="/dashboard" replace={true} />
        )}
        <form
          onSubmit={(event) => this.handleSubmit(event)}
        >
          <input type="text" name="username" />
          <input type="password" name="password" />
        </form>
      </div>
    );
  }
```

#### Outlet
父路由元素中应使用 `<Outlet>` 来呈现其子路由元素。这样就可以在呈现子路由时显示嵌套用户界面。如果父路由完全匹配，则会呈现子索引路由；如果没有索引路由，则不会呈现任何内容。

```js
interface OutletProps {
  context?: unknown;
}
declare function Outlet(
  props: OutletProps
): React.ReactElement | null; 
```

示例如下：
```jsx
function Dashboard() {
  return (
    <div>
      <h1>Dashboard</h1>

      {/* This element will render either <DashboardMessages> when the URL is
          "/messages", <DashboardTasks> at "/tasks", or null if it is "/"
      */}
      <Outlet />
    </div>
  );
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />}>
        <Route
          path="messages"
          element={<DashboardMessages />}
        />
        <Route path="tasks" element={<DashboardTasks />} />
      </Route>
    </Routes>
  );
}
```
定义一个简单的路由结构，其中包含一个 Dashboard 组件和一些子路由。这里的 Dashboard 组件作为父路由，DashboardMessages 和 DashboardTasks 作为子路由。

Dashboard 是一个简单的 React 组件，包含一个标题 "Dashboard" 和一个 `<Outlet />` 组件。

`<Outlet />` 是 React Router 提供的组件，用于在父路由中渲染匹配的子路由组件。当 URL 匹配子路由的路径时，子路由组件会被渲染到 `<Outlet />` 位置。这样就实现了在不同路径下渲染不同的组件。

#### Routes 与 Route
在应用程序中的任何地方，`<Routes>` 都会匹配当前位置的一组子路由。
```jsx
interface RoutesProps {
    children?: React.ReactNode;
    location?: Partial<Location> | string;
}

<Routes location>
    <Route />
</Routes>;
```

> 如果使用的是 `createBrowserRouter` 这样的数据路由器创建路由，使用 `Routes` 组件的情况并不常见，因为作为 `<Routes>` 树的后代的一部分定义的路由无法利用 `RouterProvider` 应用程序可用的数据 API。

每当位置发生变化时， `<Routes>` 就会查看其所有子路由，找出最匹配的路由，并渲染用户界面的该分支。 `<Route>` 元素可以嵌套，以表示嵌套的用户界面，这也与嵌套的 URL 路径相对应。父路由通过呈现 `<Outlet>`来呈现其子路由。

```jsx
<Routes>
<Route path="/" element={<Dashboard />}>
  <Route
    path="messages"
    element={<DashboardMessages />}
  />
  <Route path="tasks" element={<DashboardTasks />} />
</Route>
<Route path="about" element={<AboutPage />} />
</Routes>
```

#### ScrollRestoration
将在加载程序完成后，模拟浏览器在位置更改时的滚动恢复功能，以确保滚动位置恢复到正确位置，甚至跨域滚动。

只需呈现其中一个，建议在应用程序的根路由中呈现：
```jsx
import { ScrollRestoration } from "react-router-dom";

function RootRouteComponent() {
  return (
    <div>
      {/* ... */}
      <ScrollRestoration />
    </div>
  );
}
```

- getKey：可选属性，用于定义 React Router 恢复滚动位置时应使用的键。
  ```jsx
  <ScrollRestoration
    getKey={(location, matches) => {
      // default behavior
      return location.key;
    }}
  />
  ```
  默认情况下，它使用 `location.key` ，在没有客户端路由的情况下模拟浏览器的默认行为。用户可以在堆栈中多次导航到相同的 URL，每个条目都有自己的滚动位置来还原。

  有些应用可能希望覆盖这一行为，并根据其他内容恢复位置。例如，一个社交应用程序有四个主要页面：
  - "/home"
  - "/messages"
  - "/notifications"
  - "/search"
  
  如果用户从 `"/home"` 开始，向下滚动一点，点击导航菜单中的 "信息"，然后点击导航菜单中的 "主页"（而不是返回按钮！），历史堆栈中就会出现三个条目：
  
  ```jsx
  1. /home
  2. /messages
  3. /home
  ```
  
  默认情况下，React Router（和浏览器）会为 1 和 3 存储两个不同的滚动位置，即使它们的 URL 相同。这意味着当用户从 2 → 3 浏览时，滚动位置会移到顶部，而不是恢复到 1 中的位置。

  这里一个可靠的产品决策是，无论用户如何到达（返回按钮或新链接点击），都要保持他们在主页上的滚动位置。为此，您需要使用 `location.pathname` 作为关键字。
  ```jsx
  <ScrollRestoration
      getKey={(location, matches) => {
          return location.pathname;
      }}
  />
  ```
  或者，您可能只想对某些路径使用路径名，而对其他路径使用正常行为：
  ```jsx
  <ScrollRestoration
    getKey={(location, matches) => {
      const paths = ["/home", "/notifications"];
      return paths.includes(location.pathname)
        ? // home and notifications restore by pathname
          location.pathname
        : // everything else by location like the browser
          location.key;
    }}
  />
  ```
  
- 防止滚动重置

  当导航创建新的滚动键时，滚动位置会重置为页面顶部。您可以防止链接和表单出现 "滚动到顶部 "行为：
  ```jsx
  <Link preventScrollReset={true} />
  <Form preventScrollReset={true} />
  ```
  
- 滚动闪烁

  如果没有 Remix 这样的服务器端渲染框架，在初始页面加载时可能会出现一些滚动闪烁。这是因为 React Router 无法还原滚动位置，直到您的 JS 捆绑包下载完毕、数据加载完毕、整个页面渲染完毕（如果您正在渲染一个旋转器，视口很可能不是保存滚动位置时的大小）。

  服务器渲染框架可以防止滚动闪烁，因为它们可以在首次加载时发送一个完整的文档，因此可以在页面首次渲染时恢复滚动。
  
## Hooks
官方提供了很多 Hook，其实上面我们也用到了不少，下面就不一一演示了：
- `useActionData`：提供上一次导航 `action` 结果的返回值，如果没有提交，则提供 `undefined` 。这个 Hook 最常用的情况是表单验证错误。
- `useAsyncError`：从最近的 `<Await>` 组件返回拒绝值。
- `useAsyncValue`：从最近的 `<Await>` 父组件返回已解析的数据。
- `useBeforeUnload`：该钩子只是 `window.onbeforeunload` 的一个辅助工具。在用户离开页面之前，将重要的应用程序状态保存在页面上（如浏览器的本地存储）可能会很有用。这样，如果用户回来，就可以恢复任何状态信息（恢复表单输入值等）。
- `useBlocker`：通过 `useBlocker` 钩子，可以阻止用户从当前位置导航，并为他们提供自定义用户界面，让他们确认导航。
- `useFetcher`：请求数据用的
- `useFetchers`：返回所有不带 load 、 submit 或 Form 属性正在进行的 fetchers 数组，但不包括它们的 load ， submit 或 Form 属性（不能让父组件试图控制其子组件的行为！根据实际经验，我们知道这是很愚蠢的做法）。
- `useFormAction`：自动根据上下文解析当前路由的默认操作和相对操作。
- `useHref`：返回一个 URL，可用于链接到给定的 to 位置，即使在 React Router 之外也是如此。
- `useInRouterContext`：如果组件是在 `<Router>` 的上下文中呈现，则 `useInRouterContext` 钩子返回 `true` ，否则返回 `false` 。这对某些需要知道自己是否在 React Router 应用程序上下文中呈现的第三方扩展很有用。
- `useLinkClickHandler`：返回一个用于导航的点击事件处理程序。
- `useLinkPressHandler`：返回一个用于自定义 `<Link>` 导航的按压事件处理程序。
- `useLoaderData`：提供路由 loader 返回的值。
- `useLocation`：返回当前 `location` 对象。
- `useMatch`：返回给定路径上的路由相对于当前位置的匹配数据。
- `useMatches`：返回页面上匹配的当前路由。
- `useNavigate`：会返回一个函数，让你以编程方式导航。
- `useNavigation`：该钩子会告诉你关于页面导航的一切信息，以便在数据突变时建立待定的导航指示器和优化的用户界面。例如：
  - 全局加载指示器
  - 在发生突变时禁用表单
  - 在提交按钮上添加繁忙指示器
  - 在服务器上创建新记录时优化的显示新记录
  - 在更新记录时优化的显示记录的新状态
  
- `useNavigationType`：返回当前的导航类型或用户是如何进入当前页面的；可以是通过历史堆栈上的弹出、推送或替换操作。
- `useOutlet`：返回子路由在该路由层次结构中的元素。`<Outlet>` 内部使用此钩子来呈现子路由。
- `useOutletContext`：父路由通常会管理状态或其他你希望与子路由共享的值。
- `useParams`：钩子会返回一个由 `<Route path>` 匹配的当前 URL 动态参数的键/值对组成的对象。
- `unstable_usePrompt`：钩子允许您在导航离开当前位置前通过 `window.confirm` 提示用户进行确认。
- `useResolvedPath`：此钩子根据当前位置的路径名解析给定 `to` 值中位置的 `pathname`。
- `useRevalidator`：此钩子允许您以任何理由重新验证数据。React Router 会在调用操作后自动重新验证数据，但您也可能出于其他原因（如焦点返回窗口时）需要重新验证数据。
- `useRouteError`：在 `errorElement` 中，该钩子会返回在操作、加载器或渲染过程中抛出的任何响应。
- `useRouteLoaderData`：这个钩子可以让当前呈现的路由数据在树中的任何位置都可用。这对于树中较深位置的组件需要更远位置路由的数据，以及父路由需要树中较深位置子路由的数据时非常有用。
- `useRoutes`：钩子的功能上等同于 `<Routes>`，但它使用 JavaScript 对象而不是 `<Route>` 元素 元素来定义路由。`useRoutes` 的返回值要么是一个有效的 React 元素，可以用来呈现路由树；要么是 `null` （如果没有匹配的元素）。
- `useSearchParams`：用于读取和修改当前位置 URL 中的查询字符串。
- `useSubmit`：`<Form>` 的命令式版本，让程序员代替用户提交表单。
- `unstable_useViewTransitionState`: 当指定位置有活动视图转换时，此 Hook 会返回 true 。这可用于对元素应用更精细的样式，以进一步自定义视图转换。这要求通过 `Link` （或 `Form`, `navigate` 或 `submit` 调用）上的 `unstable_viewTransition` 中启用指定导航的视图转换。

## 请求方法
- json
  快捷方式：
  ```js
  new Response(JSON.stringify(someValue), {
      headers: {
          "Content-Type": "application/json; utf-8",
      },
  });
  ```
  通常用于 loader
  ```js
  import { json } from "react-router-dom";
  
  const loader = async () => {
    const data = getSomeData();
    return json(data);
  };
  ```

- redirect

  由于可以在 `loaders` 和 `actions` 中返回或抛出响应，因此可以使用 `redirect` 重定向到另一个路由。

- redirectDocument

  这是 `redirect` 的一个小封装，它将触发文档级重定向到新位置，而不是客户端导航。

- replace

  这是一个围绕重定向的封装，它将使用 `history.replaceState` 代替 `history.pushState` 触发客户端重定向到新位置。

## 实用工具
- `createRoutesFromChildren`：其实它是 `createRoutesFromElements` 的别名。
- `createRoutesFromElements`：是一个从 `<Route>` 元素创建路由对象的辅助工具。
- `createSearchParams`：是对 `new URLSearchParams(init)` 的轻量级封装，增加了对使用具有数组值的对象的支持。该函数与 `useSearchParams` 内部使用的从 `URLSearchParamsInit` 值创建 `URLSearchParams` 对象的函数相同。
- `defer`：该实用程序允许您通过传递承诺而不是解析值来延迟从loader返回的值。
- `generatePath`：将一组参数插值为路由路径字符串，其中包含 `:id` 和 `*` 占位符。
- `isRouteErrorResponse`：如果路由错误是路由错误响应，则返回 `true`。
- `Location`：React Router 中的 "位置 "一词指的是 history 库中的 Location 接口。
- `matchPath`：将路由路径模式与 URL 路径名进行匹配，并返回匹配信息。
- `matchRoutes`：针对一组路由与给定的 `location` 运行路由匹配算法，查看哪些路由（如果有）匹配。如果发现匹配，就会返回一个 `RouteMatch` 对象数组，每个匹配路由对应一个对象。
- `renderMatches`：`renderMatches` 会将 `matchRoutes()` 的结果渲染为一个 React 元素。
- `resolvePath`：`resolvePath` 将给定的 `To` 值解析为具有绝对 `pathname` 的实际 `Path` 对象。每当您需要知道相对 `To` 值的确切路径时，这个功能就非常有用。例如， `<Link> 使用该函数来了解其指向的实际 URL。
  
## 总结

React Router 是单页应用（SPA）中管理 URL 和视图映射关系的重要工具。它在不同版本中不断演进，提供了更强大、灵活的路由管理功能。当前主流版本的 React Router 具备了多个核心组件和功能，使得路由配置更加简单和直观。

1. 核心组件与功能
- `BrowserRouter` 和 `HashRouter` 是两个基础的路由器，分别适用于不同的环境和需求。
- `Routes` 和 `Route` 是定义路由和嵌套路由的基本单元，结合 `Link` 和 `NavLink`，可以轻松创建导航和管理路由的激活状态。
- `useNavigate` 和 `useLocation` 作为钩子，提供了编程式导航和获取当前路由信息的能力。
- `Outlet` 用于处理嵌套路由的内容渲染，是实现多级路由结构的关键。

2. 高级用法
- 动态路由 可以通过 `URL` 参数和查询参数实现个性化页面展示。
- 使用 懒加载 可以提高应用性能，结合 `lazy` 和 `Suspense` 组件，能够在用户访问特定页面时按需加载代码。
- `Navigate` 组件用于页面重定向，保证用户在特定条件下访问正确的页面。
- 路由守卫 可以通过自定义钩子或组件实现，确保用户访问受保护的路由时经过验证。

3. 状态管理的结合

React Router 与 Redux 或 Context API 等状态管理工具的结合，确保应用在路由切换时保持状态的一致性。处理好路由切换时的组件状态问题，有助于提升用户体验。

4. 常见问题
在开发中，可能会遇到路径基准（basename）设置、组件不更新以及 404 页面处理等问题。通过正确的配置和使用，React Router 能有效避免这些常见坑。

5. 性能优化
为了进一步提升应用的性能，React Router 提供了诸如组件懒加载、缓存策略等优化手段，减少不必要的重新渲染，提升页面响应速度。