# CLAUDE.md

## 仓库概述

这是一个**文档仓库**，包含中文全栈开发学习资料，原内容首发于「长林啊」微信公众号。仓库包含 194+ 个按技术主题组织的 Markdown 文件，以及多个子目录中的配套代码示例。

**核心要点**：这主要是一个笔记/文档仓库，而非传统的软件项目。大多数更改是向 Markdown 文件添加新的教育内容。

## 目录结构

```text
Full-Stack-Notes/
├── Next.js/              # Next.js 框架相关文章 (17 篇)
├── react/                # React 开发相关文章 (13 篇)
├── go/                   # Go 语言教程
│   ├── 基础/             # Go 基础 (17 篇)
│   ├── 并发编程/         # 并发编程 (3 篇)
│   ├── Go Web/Gin/       # Gin 框架
│   ├── goframe/          # GoFrame 框架 (4 篇)
│   └── go-zero/          # go-zero 微服务框架 (规划 27 篇)
├── algorithm/code/       # 算法实现 (TypeScript + Go)
├── design-patterns/code/ # 设计模式实现
├── javascript/           # JavaScript 教程
├── css/                  # CSS 方法论
├── micro-service/        # 微服务架构 (5 篇)
├── web-security/         # Web 安全主题
├── frontend/             # 前端最佳实践
├── interview/            # 面试准备
├── shell/                # Shell 脚本指南
├── vim/                  # Vim 编辑器指南
├── vscode/               # VSCode 配置
├── sync/                 # 同步工具 (Node.js 工具)
└── plan/                 # 规划文档
```

## 开发命令

### 算法代码 (`/algorithm/code/`)

```bash
# 运行测试
cd algorithm/code
pnpm test              # 或 npm test

# TypeScript 编译使用 Jest 的 ts-jest
```

### 设计模式代码 (如 `/design-patterns/code/strategy/ts/`)

```bash
cd design-patterns/code/strategy/ts
pnpm test              # 运行测试
pnpm test:watch        # 监听模式
```

### Markdown 代码规范检查

仓库使用 `markdownlint`，自定义规则位于 `.markdownlint.json`：

- 行长度扩展至 2000 字符（代码块除外）
- ATX 标题样式（`#` 前缀）
- 星号强调/粗体样式
- 反引号围栏代码块

```bash
# 如果本地安装了 markdownlint-cli：
npx markdownlint "**/*.md"
```

## 代码架构

### 多语言方案

- **Go**：用于后端示例、算法和 Web 框架教程（Gin、GoFrame、go-zero）
- **TypeScript/JavaScript**：用于前端（React、Next.js）和算法实现
- **内容语言**：中文，技术术语保留英文

### 算法实现模式

位于 `/algorithm/code/`：

- TypeScript 实现位于 `TypeScript/` 目录
- Go 实现并存
- 基于 Jest 的测试框架
- ES 模块，ESNext 目标
- 启用严格类型检查

### 设计模式实现模式

每个模式都有独立的子目录，包含：

- 独立的 `package.json`，使用 pnpm
- Jest + ts-jest 测试
- TypeScript 严格模式

### 同步工具 (`/sync/`)

- 用于仓库同步的 Node.js 工具
- 使用 `ts-node` 执行 TypeScript
- 依赖：`node-fetch`、`form-data`、`typescript`

## 内容组织

1. **分层学习路径**：主题从基础到高级组织
2. **混合内容格式**：理论 Markdown 文件 + 实用代码示例
3. **活跃规划**：`go-zero/plan/` 目录显示该系列规划了 27 篇文章
4. **微信公众号首发**：内容从微信公众号流转到此仓库

## 重要说明

- **包管理器**：使用 `pnpm`（从 `pnpm-lock.yaml` 和 packageManager 配置可见）
- **Git 忽略**：包含 Go、Node.js、Python 和 Rust 的全面 `.gitignore`
- **无 CI/CD**：这是文档仓库，没有自动化流水线
- **许可**：内容仅供个人学习；商业用途需获得许可
- **署名**：转载时请注明出处「长林啊」微信公众号

## 常见任务

在此仓库中工作时，通常会：

1. 向主题目录添加新的 Markdown 教程文件
2. 在 `/algorithm/code/` 或 `/design-patterns/code/` 中创建配套代码示例
3. 运行算法/设计模式实现的测试
4. 确保 Markdown 文件符合 `.markdownlint.json` 规则
5. 维护中文内容，技术术语使用英文
