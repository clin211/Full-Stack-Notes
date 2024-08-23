React 状态管理是指处理和更新 React 组件状态的过程。它涉及创建、更新和跨不同组件共享数据，以确保用户界面 (UI) 保持一致且响应迅速。

组件的状态是可以根据用户交互或其他因素而变化的动态方面。有效的状态管理对于构建能够有效处理用户交互和数据变化的复杂 React 应用程序至关重要。

## 什么是状态？

“状态” 是指描述某一时刻应用程序的全部数据，而状态管理则是指在应用中维护和更新状态的方法和工具。而根据使用范围和管理方式，状态又可以大致划分为三类。

- **局部状态**：这是指那些***仅在单个组件内使用和维护的状态***。在函数式组件中，使用 `useState` 或 `useReducer` 钩子来创建和管理这些状态。在类式组件中，则使用 `this.state` 和 `this.setState` 方法。局部状态一般用于控制组件的内部行为，如用户输入、表单控制等。
- **全局状态**：当***状态需要在多个组件间共享***时，我们使用全局状态。这些状态不属于某一个特定的组件，而是整个应用的共享资源。为了管理全局状态，React 提供了 Context API，但通常我们会选择更专业的库，如 React Redux、Zustand、Jotai 等，它们提供了更强大的功能和更好的开发体验。
- **服务器状态**：这种状态涉及到与服务器的交互。服务器状态是从服务器获取的数据，并在前端应用中展示和管理。例如，当你从 API 获取数据并在 UI 中展示时，就在处理服务器状态。管理服务器状态的流行库包括 Tanstack Query、SWR、Apollo Client 等，它们提供数据获取、缓存、更新等功能，简化了与后端数据交互的复杂性。

## 主流的状态管理库

截止今日(2024-08-22)各主流库的 [trends](https://npmtrends.com/jotai-vs-mobx-vs-react-query-vs-recoil-vs-redux-vs-swr-vs-valtio-vs-xstate-vs-zustand) 数据如下：

![QQ_1724295328908](assets/QQ_1724295328908.png)

| 名称                                                     | star数 | 最新版本 | 创建时间 | 最近更新时间 |
| -------------------------------------------------------- | ------ | -------- | -------- | ------------ |
| [Redux](https://github.com/reduxjs/redux)                | 60.7k  | v5.0.1   | 2012     | 2023-11-24   |
| [Zustand](https://github.com/pmndrs/zustand)             | 45.6k  | V4.5.5   | 2017     | 2024-08-16   |
| [MobX](https://github.com/mobxjs/mobx)                   | 27.4k  | v6.13.1  | 2016     | 2024-07-18   |
| [XState](https://github.com/statelyai/xstate)            | 26.8k  | v2.2.1   | 2017     | 2024-08-20   |
| [Recoil](https://github.com/facebookexperimental/Recoil) | 19.5k  | v0.7.7   | 2020     | 2023-04-12   |
| [Jotai](https://github.com/pmndrs/jotai)                 | 18.1k  | v2.9.3   | 2021     | 2024-08-18   |
| [Valtio](https://github.com/pmndrs/valtio)               | 8.8k   | v1.13.2  | 2021     | 2024-03-02   |
| [react-query](https://github.com/TanStack/query)         | 41.4k  | v5.52.0  | 2019     | 2024-08-21   |
| [swr](https://github.com/vercel/swr)                     | 30.1k  | v2.2.5   | 2019     | 2024-02-16   |

从上面的图中可以看到，可以看到 react-redux 无论 star 数还是流行度都遥遥领先，其次就是 Zustand。虽然 redux 作为老牌的状态管理库仍被用在大量项目中，但他的 Stars 上升排名仅仅只能排到第十位，而且其由于较高的上手难度以及较为繁琐的配置和模板代码，被很多人所诟病。下图就是2023年过去一年状态管理库 star 数的上升情况（https://risingstars.js.org/2023/en#section-statemanagement）:

![QQ_1724312823400](assets/QQ_1724312823400.png)

## React redux

[React Redux](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Freduxjs%2Freact-redux) 是 Redux 的官方 React UI 绑定层。它允许 React 组件从 Redux 存储中读取数据，并将操作调度到存储以更新状态。