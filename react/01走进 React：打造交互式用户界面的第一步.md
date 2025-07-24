> 在当今快节奏的技术发展中，学习一门新技术的最佳方式往往是亲身实践。而对于前端开发者而言，React 已经成为了不可或缺的技能之一。

本文大纲:

- 开发环境
- 开发工具
- 什么是 react
- 创建第一个 React 应用
- 文件结构
- Hello World

React 以其简洁、高效的组件化开发方式和强大的生态系统吸引了越来越多的开发者投身其中。对于初学者而言，掌握 React 的基础知识和关键概念是迈向 React 开发者之路的第一步。本文旨在通过一步步的实践，帮助初学者快速入门 React。

将从最基础的搭建开发环境开始，了解 React 的基本概念，逐步引领读者创建第一个 React 应用。在此过程中，会深入探讨 JSX 语法的使用、组件化开发的重要性以及组件的概念与用法。除此之外，还将介绍事件处理、条件渲染、列表渲染和表单处理等方面的内容，让读者对 React 的开发方式和核心特性有更加清晰的认识。

## 开发环境

在踏入 React 的世界之前，我们需要先搭建好开发环境。这个过程将涉及到使用一些常见的工具和技术，包括 Node.js、pnpm 以及 Vite。让我们一起来了解如何搭建这个环境，为后续的React学习和开发做好准备。

### Node.js

> Node.js® is an open-source, cross-platform JavaScript runtime environment.

Node.js 是由 Ryan Dahl 于 2009 年创建的开源、跨平台的 JavaScript **运行时环境**。它的设计目标是构建高度可伸缩的网络应用程序，特别是那些需要处理大量并发连接的实时应用。Node.js 的核心组件是基于 Chrome V8 引擎，它执行 JavaScript 代码，并采用事件驱动、非阻塞 I/O 的模型。除了服务端外，Node.js 提供的形如 `fs`、`http`、`path` 等内置模块，同时又为创造各种工具提供了方便的能力。**前端**生态也因此变得庞大，虽然现在 Rust 开始侵入，但泛前端领域自身的 CLI 基本上都还是 Node.js 写的。

Node.js 不仅在服务器端应用广泛，还成功拓展到桌面端。从 NW.js 到 Electron.js，以及诸如 heX 的桌面端运行时，Node.js 为跨平台桌面应用提供了强大支持。Electron.js 尤为突出，许多知名软件如 Visual Studio Code、1Password、钉钉等都是基于它开发的。这种开发方式允许使用熟悉的 Web 技术，同时兼顾了 Node.js 在后端处理的能力，为桌面应用开发提供了灵活而创新的选择。

#### 通过 NVM 安装 Node.js

>如何在一台电脑上安装多个 Node.js 版本？

在公司开发项目的过程中，我们常常会面临一些使用老版本 Node.js 的项目，但同时也需要为新项目采用最新版本的 Node.js。这时候，我们面临一个严峻的问题：如何在同一台设备上安装多个不同版本的 Node.js？显然，不能因为开发新项目就忽略掉原有的老项目。如果我们强行升级老项目所使用的 Node.js 版本，可能会导致一些第三方库的依赖出现各种奇怪的问题。

  这时，Node Version Manager（NVM）就成为了解决这一问题的利器。通过使用 NVM，我们可以在同一设备上轻松地安装和管理多个不同版本的 Node.js。这使得我们能够灵活地切换 Node.js 版本，而不必担心新项目与老项目之间的冲突。NVM 的存在为我们提供了一种高效、可靠的方式，在不同项目之间自由切换 Node.js版本，确保项目的稳定性和兼容性。

  > NVM 的 github 地址：
  >
  > - [Mac](https://github.com/nvm-sh/nvm): https://github.com/nvm-sh/nvm
  > - [Windows](https://github.com/coreybutler/nvm-windows)：https://github.com/coreybutler/nvm-windows 是使用 Go 实现的。


  **NVM（Node Version Manager）** 是一个用于管理 Node.js 版本的工具，它允许您轻松地在同一台计算机上安装和切换不同版本的 Node.js。以下是使用 NVM 搭建 Node.js 环境的详细步骤：

1. 安装 NVM：

    打开终端（命令行工具），在终端中运行以下命令安装 NVM（在 Linux/macOS 上）：

    ```bash
    $ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
    ```

    或者使用 wget 安装：

    ```bash
    $ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
    ```

    在 Windows 上，可以通过下载并运行 [nvm-windows](https://github.com/coreybutler/nvm-windows) 安装程序来安装 NVM。

2. 重启终端：

    - 安装完成后，关闭当前终端并重新打开一个新的终端窗口，或者执行以下命令以使 NVM 起效：

      ```bash
      source ~/.bashrc  # 或者 source ~/.zshrc（根据使用的终端类型）
      ```

    - 查看 NVM 的所有命令：

      ```shell
      $ nvm -h
      ```

    - **查看 NVM 的版本**以查看是否安装成功：

      ```shell
      $ nvm -v # 安装成功后就会打印当前设备的安装的 nvm 版本
      # 0.39.7
      ```

    - 在终端中运行以下命令**安装最新版本**的 Node.js：

      ```shell
      $ nvm install node # 默认安装最新版本的 Node.js
      ```

    - **安装指定版本**的 Node.js：

      ```shell
      $ nvm install <version> # nvm install v20.11.0
      ```

    - 安装完成后，可以使用以下**命令切换 Node.js 版本**：

      ```shell
      $ nvm use <version> 
      ```

    - 使用以下命令**查看已安装**的 Node.js 版本：

      ```shell
      $ nvm ls
      ```

    - **查看当前正在使用**的版本：

      ```shell
      $ nvm current
      ```

    - 还可以**查看所有可用的 Node.js 版本**列表：

      ```shell
      $ nvm ls-remote
      ```

    - 如果想要**将某个版本设置为默认版本**，可以使用以下命令：

      ```shell
      $ nvm alias default <version>
      ```

  通过使用 NVM，可以方便地在不同的 Node.js 项目中使用不同的版本，同时保持环境的整洁和灵活性。这对于开发需要特定 Node.js 版本的项目是非常有用的。

#### 安装 nrm

  在开发中，项目或多或少的都会依赖一些第三方库，一定会使用 npm 或者其他包管理工具来安装第三方依赖包，但由于 npm 默认的下载仓储地址是 https://registry.npmjs.org/，这个镜像网站在国外，所以我们下载的时候可能会非常的慢。所以淘宝也做了一个 npm 的镜像网站（[npmmirror 镜像站](https://www.npmmirror.com/)）。

  所以要更改要将默认的镜像源改成淘宝源，命令如下：

  ```shell
  npm config set registry https://registry.npmmirror.com/
  ```

  我们天天跟程序打交道，这么几个命令也还可以将就记，但是后面的镜像源这么长，不好记，有没有什么方便的工具来实现快速切换呢？NRM 没错就是它，[nrm](https://www.npmjs.com/package/nrm)（NPM registry manager）是 npm 的镜像源管理工具，使用它可以快速切换 npm 源。

- 安装命令

  ```shell
  $ npm i -g nrm
  ```

  > 可以通过 `nrm -h` 查看所有命令。

- 查看所有镜像源

  ```shell
  $ nvm ls
  ```

  ![所有镜像源](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/WUFDcc-631de43e-8262-46f9-be7b-a73bd6b42718.png)

- 切换镜像

  ```shell
  $ nrm use <镜像名>
  ```

  比如：切换到 npm 源

  ![切换到 npm 源](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/wyta8y-7baef6ea-0421-4efe-b713-e0b7a6315505.png)

- 添加镜像源

  > 适用于企业内部定制的私有源，`<registry>` 表示源名称，`<url>` 表示源地址

  ```shell
  $ nrm add <registry> <url>
  ```

- 删除镜像源

  ```shell
  nrm del <镜像名>
  ```

- 测试镜像源的响应时间

  ```shell
  $ nrm test <镜像名>
  ```

#### 包管理工具 pnpm

  >使用 npm 时，依赖每次被不同的项目使用，都会重复安装一次。  而在使用 pnpm 时，依赖会被存储在内容可寻址的存储中，所以：
  >如果你用到了某依赖项的不同版本，只会将不同版本间有差异的文件添加到仓库。 例如，如果某个包有100个文件，而它的新版本只改变了其中1个文件。那么 pnpm update 时只会向存储中心额外添加1个新文件，而不会因为仅仅一个文件的改变复制整新版本包的内容。
  >所有文件都会存储在硬盘上的某一位置。 当软件包被被安装时，包里的文件会硬链接到这一位置，而不会占用额外的磁盘空间。 这允许你跨项目地共享同一版本的依赖。 —— 摘抄自官网

  [pnpm](https://pnpm.io/zh/) 是一个 Node.js 高效的、快速的、节省磁盘空间的包管理工具，它提供了与 npm 和 Yarn 相似的功能，但具有一些独特的特性。以下是关于 pnpm 的详细介绍：

  - 快速、高效的独立版本控制： 与传统的包管理工具不同，pnpm 采用一种称为“硬链接”或“符号链接”的机制，将项目的依赖安装到单个位置，并通过符号链接的方式引用它们。这使得相同的依赖在多个项目之间可以共享，减少了磁盘空间的占用，同时提高了安装速度。
  - 并发安装和缓存： pnpm 支持并发安装，允许同时安装多个包，从而提高安装速度。此外，它使用了智能的缓存策略，避免了重复下载相同版本的包，节省了网络带宽和存储空间。
  - 原子操作： pnpm 在执行安装和卸载操作时采用了原子操作的原则，这意味着只有当所有依赖都成功安装或卸载时，才会对项目的 node_modules 目录进行更改。这有助于确保项目的稳定性和一致性。
  - 支持 npm 生态系统： pnpm 兼容 npm 的 package.json 文件，可以直接使用 npm 发布的包，也可以通过 pnpm 命令进行管理。这使得迁移到 pnpm 对于已有的 npm 项目相对容易。
  - 版本锁定： pnpm 允许你锁定项目的依赖版本，确保不同开发者或构建环境之间使用相同的依赖版本，减少因版本不一致而导致的问题。
  - 自动清理： pnpm 自带一个自动清理功能，定期清理不再使用的依赖，以释放磁盘空间。
  - 支持 Workspaces： pnpm 支持使用 workspaces 特性，方便地管理多包存储库（monorepo）。

  总体而言，pnpm 通过采用一些创新性的策略，如硬链接和并发安装，提供了更快、更高效、更节省空间的包管理解决方案，特别适用于大型项目和 monorepos 的场景。

- 安装 pnpm

  > [官网](https://pnpm.io/zh/installation)提供了很多种安装姿势，这里就介绍最简单的安装方式吧！—— 使用 npm 安装

  ```shell
  $ npm install -g pnpm
  ```

- 查看安装的 pnpm 版本，可以使用如下命令：

  ```shell
  $ pnpm -v
  ```

## 开发工具
当涉及到开发工具的选择时，[Visual Studio Code](https://code.visualstudio.com/)（简称 vs code）和 [WebStorm](https://www.jetbrains.com/zh-cn/webstorm/) 是两个备受开发者青睐的编辑器。

[vs code](https://code.visualstudio.com/) 是一个轻量级的开源代码编辑器，由微软开发。它具有丰富的插件生态系统和智能代码补全功能，同时提供语法高亮、代码导航、调试器和 Git 集成等特性，方便开发者进行代码编辑和调试。

![截图于 vscode 官网](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/VrVNUX-b21b4ba3-9412-45cc-8982-2248a9dbbc8d.png)


[WebStorm](https://www.jetbrains.com/zh-cn/webstorm/) 是一款由JetBrains开发的专注于Web开发的IDE。它支持多种前端技术和框架，提供智能代码补全、错误纠正、实时代码分析等功能。

![截图于 webstorm 官网](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/aWR5L7-90a2d083-2ba1-42c4-bbfb-aeb9c9e66348.png)

总而言之，它两都差不多，想定制化高一点就使用 vs code，不想折腾就是 webstorm，用着顺手就好，下文内容都是使用 vs code IDE（我个人比较喜欢使用 vs code）。

## 初识 React

![官网](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/qlDVJB-adeb7f5e-a83c-4b49-a474-f21a8c3b85bb.png)

### 什么是 React？

[React](https://react.dev/) 是一个用于构建 Web 和原生用户界面的 JavaScript 库。开发者可以用 JavaScript 编写由单个部分组成的组件，然后将它们无缝地组合起来。React 的优势在于，它使你能够与不同的人、团队和组织共享和重用组件。它的基本概念包括 DOM、组件、虚拟 DOM 等。React 主要用于构建 UI。React 起源于 Facebook 的内部项目，在 2013 年 5 月开源。

### 创建第一个 React 应用

如果你是初次接触 React，我建议你使用 Create React App 来创建你的第一个 React 应用，这是最流行的尝试 React 功能的方式，也是构建新的单页客户端应用的最好方法。以下是一些步骤：
1. 首先，你需要在你的电脑上安装 Node.js 和 npm（如果没有，可以按照上面的步骤进行操作）。
2. 然后，打开命令行工具，执行下面的命令来创建一个新的 React 应用：
    ```sh
    $ npx create-react-app demo
    ```
    这里的 "my-app" 是你的应用的名字，你可以选择你喜欢的名字。
    > 这里**项目名不能使用驼峰的形式**(即不能包含大写字母)，不然会有以下报错：`name can no longer contain capital letters`。
    
    <img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/CPJtRe-1e1321af-a955-41b1-b2f9-b900b2bfab29.png" />

3. 创建完毕后，进入你的应用的目录：
    ```sh
    $ cd demo
    ```
4. 然后，启动你的应用：
    ```sh
    npm start
    ```
    现在，你可以在浏览器中打开 "http://localhost:3000" 来查看你的 React 应用了！

    ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/TsISwu-5385716e-e1be-453f-ad51-a0851d043e87.gif)
    
### 文件结构
接下来，我们看下安装好的 React 应用的文件结构。
```tree
.
├── README.md                  // 文档
├── package-lock.json          // npm 依赖的详细
├── package.json               // npm 依赖
├── public                     // 静态资源文件夹
│   ├── favicon.ico            // 网站 icon 图标
│   ├── index.html             // 模板
│   ├── logo192.png            // 192*192大小的react logo
│   ├── logo512.png            // 512*512大小的react logo
│   ├── manifest.json          // 移动桌面快捷方式配置文件
│   └── robots.txt             // 网站与爬虫间的协议
└── src                        // 源码文件夹
    ├── App.css
    ├── App.js                 // 根组件
    ├── App.test.js
    ├── index.css              // 全局样式
    ├── index.js               // 入口文件
    ├── logo.svg
    ├── reportWebVitals.js
    └── setupTests.js
```

### 创建你的 Hello World
为了快速了解 React，我们先改动一个文件看看效果，从动手尝试中去学习。
打开 `src/index.js` 文件，可以看到 render 的模版是 App.js，代码如下：
```js
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
```

> `<React.StrictMode>` 和`<App />` 是 JSX 语法（后面会讲什么是 JSX 语法）。
>
> root 是 index.html 模版里的元素，渲染出来的 App 组件放在此处。

我们来尝试修改 App.js 文件，修改后的代码如下所示：
```js
import './App.css';

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <p>hello world！！！</p>
            </header>
        </div>
    );
}

export default App;
```

可以看到上面代码中，我们将 header 中的内容换成了 `<p>hello world！！！</p>`，保存代码，浏览器会自动帮我们刷新，接着我们就能看到页面展示如下：


![hello world 效果](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/Z3TfH3-4e8013e8-84e4-4822-8020-5c775623677d.png)

到此，我们已经完成了使用 React 的第一步。