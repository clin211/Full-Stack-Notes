![截自官网 https://tanstack.com/query/latest](assets/QQ_1726303102679.png)

[React-Query](https://tanstack.com/query/latest) 是一个用于在 web 应用中获取、缓存、同步和更新服务器状态的库。它简化了数据获取过程，使开发者能够专注于业务逻辑，而无需处理繁琐的状态管理。它自动管理请求状态，包括加载、错误处理和数据缓存，极大提高了开发效率。内置的缓存机制不仅减少了网络请求，还提升了应用性能和用户体验。还支持复杂用例，如分页和实时数据获取。另外，它能与现代框架（如 React 和 Vue）及其他状态管理库（如 Redux 和 Zustand）无缝集成，增强了灵活性。

React-Query 通过使用查询键来标识不同接口返回的数据，而查询函数就是我们请求后端接口的函数。React-Query 中的查询是对异步数据源的声明性依赖，它与唯一键绑定。查询可以与任何基于 Promise 的方法一起使用（包括 GET 和 POST 方法）来从服务器获取数据。

说明其重要性和用途
讨论为什么选择 TanStack Query 而不是其他数据获取库
## 1. 什么是 TanStack Query？
定义和背景
主要功能和特点
与 React Query 的关系
## 2. 安装和设置
安装步骤
npm 或 yarn 命令
基本配置
提供 QueryClient
在应用中使用 QueryClientProvider
## 3. 基本用法
创建第一个查询
使用 useQuery 钩子
示例代码
创建变更请求
使用 useMutation 钩子
示例代码
## 4. 查询状态管理
查询状态的不同状态（加载、成功、错误）
如何处理状态和错误
使用 isLoading 和 isError 属性
## 5. 数据缓存与更新
数据缓存机制
如何手动更新缓存
使用 queryClient.invalidateQueries 方法
## 6. 处理分页和无限滚动
实现分页查询
实现无限滚动
示例代码
## 7. 订阅与实时数据
如何使用 WebSocket 或其他实时数据源
示例代码
## 8. 结合其他状态管理库
在 Redux 或 Zustand 中使用 TanStack Query
示例代码
## 9. 性能优化
使用 staleTime 和 cacheTime 来优化性能
如何避免不必要的请求
## 10. 常见问题
TanStack Query 的常见误区
常见错误及其解决方法