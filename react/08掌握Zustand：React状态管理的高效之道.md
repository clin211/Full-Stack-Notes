在上一篇《React的状态管理：主流状态库的对比》的文章中，我们深入探讨了React生态中那些经典的状态管理库，特别是 Redux 这老牌库，能完成各种基本功能，并且有着庞大的中间件生态来扩展额外功能，但 redux 经常被人诟病它的使用繁琐。近两年，React 社区出现了很多新的状态管理库，zustand 算是其中最流行的一个；从 star (2024年9月11日) 数看，redux 有 60.8k，而 zustand 也有 46.2k 了，Zustand 是 2021 年 Star 增长最快的 React 状态管理库，也是最近这三年呼声最高的一个状态管理库，设计理念函数式，全面拥抱 hooks，API 设计简洁、优雅，对业务的侵入小，是一个现代化、高效且易于使用的状态管理解决方案，它正逐渐成为React 开发者的新宠儿。主要是学习的成本也不是很高，维护的心智负担也比较小。

![img](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/UYM6hc-bear.jpg)

## zustand 介绍

[Zustand](https://zustand.docs.pmnd.rs/getting-started/introduction) 是一个轻量级的状态管理库，用于 JavaScript 应用程序，特别是在 React 生态系统中。它提供了一个简单、可扩展的解决方案来管理应用程序的状态。

与其他状态管理解决方案（如 Redux ）相比，`Zustand` 旨在提供更简洁的 API 和更少的样板代码。它允许你创建一个全局状态存储，并且可以在应用程序的任何地方访问和更新这个状态，而不需要像 Redux 那样编写大量的 action creators 和 reducers。

### 为什么使用 zustand

- 学习曲线平缓。只需几行代码即可创建管理状态
- 无需提供 React 上下文（Context），这使得状态管理更灵活，且可以在任何组件中轻松访问状态。
- 体积相对较小，减少了项目的依赖和加载时间
- 原生支持 TypeScript，提供类型安全，增强了开发体验，便于维护和扩展。
- ……

### 它的出现解决了那些问题

Zustand 是一个为 React 应用程序提供状态管理的库，它旨在简化状态管理的过程，解决了许多开发者在使用其他状态管理库时遇到的问题。以下是 Zustand 解决的一些主要问题：

1. **简化状态管理**：Zustand 提供了一个简单直观的 API，使得创建和管理状态变得容易，无需复杂的配置。
2. **减少样板代码**：它减少了编写样板代码的需要，如在 Redux 中常见的 action 创建和 reducer 编写。
3. **全面支持 Hooks**：Zustand 完全基于 React Hooks，使得在函数组件中使用状态变得非常自然和方便。
4. **性能优化**：Zustand 通过使用 Proxy 来实现状态的响应式更新，只有当状态真正改变时，组件才会重新渲染，这有助于提高应用的性能。
5. **易于测试**：由于其简单性和基于 Hooks 的设计，使用 Zustand 的应用更容易进行单元测试和集成测试。
6. **减少心智负担**：Zustand 的设计哲学是减少开发者的心智负担，使得状态管理变得更加直观和容易理解。
7. **更好的类型支持**：Zustand 与 TypeScript 有很好的集成，提供了良好的类型推断和类型安全，减少了运行时错误的可能性。
8. **易于集成**：Zustand 可以很容易地与其他库和工具集成，如 Redux DevTools，使得调试和开发更加方便。
9. **支持中间件**：虽然不像 Redux 那样拥有庞大的中间件生态，Zustand 也支持中间件，允许开发者扩展其功能。
10. **轻量级**：Zustand 的体积相对较小，不会给应用带来过多的负担，这对于性能敏感或需要快速加载的应用来说是一个优势。

### 优势及特点

- **轻量级** ：Zustand 的整个代码库非常小巧，gzip 压缩后仅有 1KB，对项目性能影响极小。

- **简洁的 API** ：Zustand 提供了简洁明了的 API，能够快速上手并使用它来管理项目状态。   
- **基于 React Hooks**: Zustand 使用 React 的 Hooks 机制作为状态管理的基础。它通过创建自定义 Hook 来提供对状态的访问和更新。这种方式与函数式组件和钩子的编程模型紧密配合，使得状态管理变得非常自然和无缝。

- **易于集成** ：Zustand 可以轻松地与其他 React 库（如 Redux、MobX 等）共存，方便逐步迁移项目状态管理。

- **支持 TypeScript**：Zustand 支持 TypeScript，让项目更具健壮性。

- **灵活性**：Zustand 允许根据项目需求自由组织状态树，适应不同的项目结构。

- **可拓展性** : Zustand 也引入了中间件 (middleware) 的概念，允许你通过插件的方式扩展其功能。中间件可以用于处理日志记录、持久化存储、异步操作等需求，使得状态管理更加灵活和可扩展。

- **性能优化**: Zustand 在设计时非常注重性能。它采用了高效的状态更新机制，避免了不必要的渲染。同时，Zustand 还支持分片状态和惰性初始化，以提高大型应用程序的性能。

- **无副作用**: Zustand 鼓励无副作用的状态更新方式。它倡导使用 immer 库来处理不可变性，使得状态更新更具可预测性，也更易于调试和维护。

### Redux 与 zustand 对比

- redux

  ```tsx
  // redux-toolkit
  import { useSelector } from 'react-redux'
  import type { TypedUseSelectorHook } from 'react-redux'
  import { createSlice, configureStore } from '@reduxjs/toolkit'
  
  const countSlice = createSlice({
        name: 'count',
        initialState: { value: 0 },
        reducers: {
              incremented: (state, qty: number) => {
                // Redux Toolkit does not mutate the state, it uses the Immer library
                // behind scenes, allowing us to have something called "draft state".
                state.value += qty
              },
              decremented: (state, qty: number) => {
                state.value -= qty
              },
        },
  })
  
  const countStore = configureStore({ reducer: countSlice.reducer })
  
  const useAppSelector: TypedUseSelectorHook<typeof countStore.getState> = useSelector
  
  const useAppDispatch: () => typeof countStore.dispatch = useDispatch
  
  const Component = () => {
        const count = useAppSelector((state) => state.count.value)
        const dispatch = useAppDispatch()
        // ...
  }
  ```

  ```ts
  // redux
  import { createStore } from 'redux'
  import { useSelector, useDispatch } from 'react-redux'
  
  type State = {
    	count: number
  }
  
  type Action = {
        type: 'increment' | 'decrement'
        qty: number
  }
  
  const countReducer = (state: State, action: Action) => {
        switch (action.type) {
              case 'increment':
                	return { count: state.count + action.qty }
              case 'decrement':
                	return { count: state.count - action.qty }
              default:
                	return state
        }
  }
  
  const countStore = createStore(countReducer)
  
  const Component = () => {
        const count = useSelector((state) => state.count)
        const dispatch = useDispatch()
        // ...
  }
  ```

- zustand

  ```tsx
  import { create } from 'zustand'
  
  type State = {
    	count: number
  }
  
  type Actions = {
        increment: (qty: number) => void
        decrement: (qty: number) => void
  }
  
  const useCountStore = create<State & Actions>((set) => ({
        count: 0,
        increment: (qty: number) => set((state) => ({ count: state.count + qty })),
        decrement: (qty: number) => set((state) => ({ count: state.count - qty })),
  }))
  
  const Component = () => {
        const count = useCountStore((state) => state.count)
        const increment = useCountStore((state) => state.increment)
        const decrement = useCountStore((state) => state.decrement)
        // ...
  }
  ```

  ![image-20240909145946338](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/S2cMUD-image-20240909145946338.png)

  ![image-20240909145851733](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/S9haX5-image-20240909145851733.png)

### 对比分析

1. 状态定义：
   - Zustand：使用 `create` 函数直接定义状态和操作。
   - Redux：需要定义初始状态和 reducer 函数。

2. 状态更新：
   - Zustand：通过 `set` 方法直接更新状态，使用函数式更新。
   - Redux：通过 dispatching action 来更新状态，涉及 action 类型和 reducer。

3. 组件连接：
   - Zustand：直接在组件中使用 `useCountStore` 获取状态和操作。
   - Redux：使用 `useSelector` 获取状态，使用 `useDispatch` 调用操作。

4. 代码复杂性：
   - Zustand：代码结构简单，聚焦于状态和操作的定义。
   - Redux：需要定义多个部分（reducer、action、store），代码量较大。


## zustand 的使用

有了上面的铺垫，接下来看看如何在 react 中如何使用 zustand！本模块就新建一个 react 应用，以 todoMVC 应用为例来快速入门 zustand！

### 新建项目并启动

- 使用 vite 创建项目

  ```sh
  $ npx create-vite
  ```

  ![QQ_1725867113285](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/35nKL1-QQ_1725867113285.png)

- 在 IDE 中打开并安装依赖并启动

  在 vscode 中打开并安装依赖，如下图：

  ![QQ_1725867303615](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/PJjnKT-QQ_1725867303615.png)

- 使用命令 `pnpm dev`，在浏览器中查看效果图如下图：

  ![QQ_1725867504024](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/OEF4V6-QQ_1725867504024.png)

  ![QQ_1725867542270](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/XIljEO-QQ_1725867542270.png)

- 安装 zustand

  ```sh
  $ pnpm add zustand
  ```

  

在刚创建的项目中添加一个计数器的应用；其中显示数字和操作数字的按钮放在两个不同的组件中！具体操作步骤如下：

1. 在项目的src目录下创建一个 store 目录

2. 在 store 目录中，创建一个 useCounter 的 ts 文件

3. 在这个 ts 文件中写入如下代码：
  ```ts
  import { create } from 'zustand'
  
  type State = {
      count: number
  }
  
  type Action = {
      increment: () => void
      decrement: () => void
  }
  
  const useCounter = create<State & Action>((set) => ({
      count: 0,
      increment: () => set((state) => ({ count: state.count + 1 })),
      decrement: () => set((state) => ({ count: state.count - 1 })),
  }))
  
  
  export default useCounter
  ```

4. 在项目的 src 目录下创建 components/CounterButton.tsx 文件，代码如下：
  ```tsx
  import React from 'react'
  import useCounter from '../store/useCounter'
  
  interface Props {
      type: 'increment' | 'decrement'
  }
  export default function CounterButton(props: Props) {
      const increment = useCounter(state => state.increment)
      const decrement = useCounter(state => state.decrement)
      
      const handleClick = () => {
          if (props.type === 'increment') {
              increment()
          } else {
              decrement()
          }
      }
      
      return (
          <button onClick={handleClick}>{props.type === 'increment' ? 'increment' : 'decrement'}</button>
      )
  }
  ```

5. 在 App.tsx 中将原来 count 的逻辑替换成如下代码：

  ```tsx
  import reactLogo from './assets/react.svg'
  import viteLogo from '/vite.svg'
  import './App.css'
  import useCounter from './store/useCounter'
  import CounterButton from './assets/components/CounterButton'
  
  function App() {
    const count = useCounter(state => state.count)
  
    return (
      <>
        <div>
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <h1>Vite + React</h1>
        <div className="card">
          <CounterButton type='decrement' />
          <p>{count}</p>
          <CounterButton type='increment' />
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </>
    )
  }
  
  export default App
  ```

  效果如下：

  ![2024-09-09 16.46.42](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/X0Huzp-2024-09-09%2016.46.42.gif)

上面就是用 zustand 简单实现了一个跨组件计数的功能，接着我们看看 zustand 中的中间件的使用（这里不会去一一介绍中间件怎么使用）。

### zustand 的中间件

`Zustand`也提供了中间件的功能，可以用于在状态更新前后执行一些操作。

官方也提供了几个常用的中间：combine、devtools、immer、presist等，中间件是一个函数，它接收一个对象，其中包含了当前状态和一个更新状态的函数。中间件可以在状态更新前后执行一些操作，例如打印日志、发送网络请求等。

下面以 immer 为例，假如现在有一个复杂的状态，这个状态对象嵌套了多层级，我们称之为`nestedObject`。如下：

```js
const nestedObject = {
    a: {
        b: {
            c: {
                d: 0,
            },
        },
    },
};
```

如果我们想要更新这个状态，应该怎么做？

```js
const useStore = create((set) => ({
    nestedObject,
    updateState () {
        set(prevState => ({
            nestedObject: {
                ...prevState.nestedObject,
                a: {
                    ...prevState.nestedObject.a,
                    b: {
                        ...prevState.nestedObject.a.b,
                        c: {
                            ...prevState.nestedObject.a.b.c,
                            d: ++prevState.nestedObject.a.b.c.d,
                        },
                    },
                },
            },
        }));
    },
}));
```

这段代码中，不难发现以下问题：

- 可读性太差，深层嵌套不仅增加理解成本还增加维护成本
- 每一层都使用扩展运算符进行拷贝，对性能也有很大的影响
- 写了很多与修改目标属性无关的代码

修改四层就如此麻烦，那要是修改个七八层的数据，稍有不慎就搞错了，排查问题也极为头痛。其实，我们可以借助 [immer](https://immerjs.github.io/immer/zh-CN/) 来优化这个问题，最终上面的代码被优化后的为：

```js
import produce from 'immer';

const useStore = create((set) => ({
    nestedObject,
    updateState () {
        set(produce(draft => {
            ++draft.nestedObject.a.b.c.d;
        }));
    },
}));
```

代码一下子就干净、整洁多了，可读性也高了不少，但也增加了额外的理解成本，本阶不做 immer 的讲解，如果读者不熟悉的话，可以去官 [immer 官网](https://immerjs.github.io/immer/zh-CN/)查看，官网也有中文！

在 zustand 中，需要单独安装 immer 库才能搭配使用，不过 zustand 官方已经适配好了，使用姿势如下：

```ts
import { create } from 'zustand'
import { immer } from 'zustand/middleware/immer'

type State = {
    count: number
}

type Actions = {
    increment: (qty: number) => void
    decrement: (qty: number) => void
}

export const useCountStore = create<State & Actions> ()(
    immer((set) => ({
        count: 0,
        increment: (qty: number) =>
            set((state) => {
                state.count += qty
            }),
        decrement: (qty: number) =>
            set((state) => {
                state.count -= qty
            }),
    })),
)
```

另一种使用方式：

```ts
import { create } from 'zustand'
import { immer } from 'zustand/middleware/immer'

interface Todo {
    id: string
    title: string
    done: boolean
}

type State = {
    todos: Record<string, Todo>
}

type Actions = {
    toggleTodo: (todoId: string) => void
}

export const useTodoStore = create<State & Actions> (immer((set) => ({
    todos: {
        '82471c5f-4207-4b1d-abcb-b98547e01a3e': {
            id: '82471c5f-4207-4b1d-abcb-b98547e01a3e',
            title: 'Learn Zustand',
            done: false,
        },
        '354ee16c-bfdd-44d3-afa9-e93679bda367': {
            id: '354ee16c-bfdd-44d3-afa9-e93679bda367',
            title: 'Learn Jotai',
            done: false,
        },
        '771c85c5-46ea-4a11-8fed-36cc2c7be344': {
            id: '771c85c5-46ea-4a11-8fed-36cc2c7be344',
            title: 'Learn Valtio',
            done: false,
        },
        '363a4bac-083f-47f7-a0a2-aeeee153a99c': {
            id: '363a4bac-083f-47f7-a0a2-aeeee153a99c',
            title: 'Learn Signals',
            done: false,
        },
    },
    toggleTodo: (todoId: string) =>
        set((state) => {
            state.todos[todoId].done = !state.todos[todoId].done
        }),
})))
```

### zustand 中异步请求

```ts

import { create } from 'zustand';

interface User {
    userId: number
    id: number
    title: string
    completed: boolean
}

type State = {
    user: User | null
    loading: boolean
    error: string | null
}

type Actions = {
    fetchUser: (id: string) => void
}

const useUserStore = create<State & Actions>((set) => ({
    user: null,
    loading: false,
    error: null,
    fetchUser: async (id) => {
        set({ loading: true });
        try {
            const response = await fetch(`https://jsonplaceholder.typicode.com/todos/${id}`);
            if (!response.ok || response.status !== 200) return;
            const data = await response.json();
            set({ user: data, loading: false });
        } catch (error) {
            set({ error: error.message || '', loading: false });
        }
    },
}));

export default useUserStore;
```

在 react 中使用：

```tsx
import { useEffect } from 'react';
import useUserStore from './store/useUserStore';

function Sync() {
    const user = useUserStore(state => state.user);
    const fetchUser = useUserStore(state => state.fetchUser);

    useEffect(() => {
        fetchUser('1');
    }, []);

    return (
        <div>
            {JSON.stringify(user, null, 4)}
        </div>
    );
}

export default Sync;
```

在这个组件中，我们使用 `useEffect` hook 在组件挂载时调用 `fetchItems` 函数。当 `fetchItems` 函数完成时，它会更新 `user` 状态，这将触发组件重新渲染。

## shallow

细心的小伙伴可能发现了，上面的代码中写了两遍 `useUserStore`，如果有100个属性或方法，那是不是要写100次 `useUserStore` 呢？如果写成 `const {user, fetchUser} = useUserStore()`，会有什么问题呢？ 

我们先用一个设置主题和语言的例子来看看写成对象解构的方式有什么问题。

创建一个存放主题和语言类型的store：

```ts
import { create } from 'zustand';

interface State {
    theme: string;
    lang: string;
}

interface Action {
    setTheme: (theme: string) => void;
    setLang: (lang: string) => void;
}

const useConfigStore = create<State & Action>((set) => ({
    theme: 'light',
    lang: 'zh-CN',
    setLang: (lang: string) => set({ lang }),
    setTheme: (theme: string) => set({ theme }),
}));

export default useConfigStore;
```

分别创建两个组件，主题组件和语言类型组件：

````tsx
import useConfigureStore from '../store/useConfigureStore';

const Theme = () => {
    const { theme, setTheme } = useConfigureStore();
    console.log('theme render', theme);

    return (
        <div>
            <div>{theme}</div>
            <button
                onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}>
                切换主题
            </button>
        </div>
    );
};

export default Theme;
````

````tsx
import useConfigStore from '../store/useConfigureStore';

const Language = () => {
    const { lang, setLang } = useConfigStore();
    console.log('lang render', lang);
    return (
        <div>
            <div>{lang}</div>
            <button
                onClick={() => setLang(lang === 'zh-CN' ? 'en-US' : 'zh-CN')}>
                切换语言
            </button>
        </div>
    );
};

export default Language;
````

查看下面效果：

![2024-09-11 10.48.16](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/BlhyPF-2024-09-11%2010.48.16.gif)

> TIPS：
>
> 上图的控制台每个打印都有两次，原因是因为 `<React.StrictMode>` 导致的
>
> ![QQ_1726036582587](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/4tr4pC-QQ_1726036582587.png)
>
> 严格模式在生产环境是不生效的，如果你想在测试环境也不生效的话，直接不使用 `<React.StrictMode>` 包裹根组件。

从上面的效果不难看出，改变 theme 会导致 Language 组件渲染，改变 Language 会导致 Theme 重新渲染，但是实际上这两个组件都没任何关系，至于为什么会导致这样，后面专门写一篇 zustand 的源码解读去分析，我们先看看这种方式要怎么优化？

第一种方案就是将原先的解构改成单值返回：

```ts
// Theme.tsx
const { theme, setTheme } = useConfigureStore();

// Language.tsx
const { lang, setLang } = useConfigureStore();
```

替换成：

```ts
// Theme.tsx
const theme = useConfigureStore(state => state.theme);
const setTheme = useConfigureStore(state => state.setTheme);

// Language.tsx
const lang = useConfigureStore(state => state.lang);
const setLang = useConfigureStore(state => state.setLang);
```

优化之后就不会出现上面的问题了，如下图：

![2024-09-11 11.48.40](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/zkSDnd-2024-09-11%2011.48.40.gif)

其实上面的写法还是不够优雅，因为要写很多遍 `useConfigStore`，我们可以把单值返回改成一个对象，这样在多属性的时候代码就相对要简洁一点。

```ts
// Theme.tsx
const { theme, setTheme } = useConfigureStore(state => ({
    theme: state.theme,
    setTheme: state.setTheme,
}));

// Language.tsx
const { lang, setLang } = useConfigStore(state => ({
    lang: state.lang,
    setLang: state.setLang,
}));
```

这种写法仍然有个问题：任意属性改变之后都会返回一个新的对象，zustand 内部拿到返回值后与上次比较，发现每次都是一个新对象，然后就重新渲染。好在 zustand 提供了解决方案，对外暴露了一个 useShallow 方法，可以浅比较两个对象；我们把上面的对象改写一下，完整代码如下：

```tsx
import { useShallow } from 'zustand/react/shallow';
import useConfigureStore from '../store/useConfigureStore';

const Theme = () => {
    const { theme, setTheme } = useConfigureStore(
        useShallow(state => ({
            theme: state.theme,
            setTheme: state.setTheme,
        }))
    );
    console.log('theme render', theme);

    return (
        <div>
            <div>{theme}</div>
            <button
                onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}>
                切换主题
            </button>
        </div>
    );
};

export default Theme;
```

```tsx
import { useShallow } from 'zustand/react/shallow';
import useConfigStore from '../store/useConfigureStore';

const Language = () => {
    const { lang, setLang } = useConfigStore(
        useShallow(state => ({
            lang: state.lang,
            setLang: state.setLang,
        }))
    );
    console.log('lang render', lang);
    return (
        <div>
            <div>{lang}</div>
            <button
                onClick={() => setLang(lang === 'zh-CN' ? 'en-US' : 'zh-CN')}>
                切换语言
            </button>
        </div>
    );
};

export default Language;
```

优化之后看看效果：

![2024-09-11 14.26.52](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/0fsAOu-2024-09-11%2014.26.52.gif)

