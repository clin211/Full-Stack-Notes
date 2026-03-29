# 21 Next.js 16 + Prisma 构建 AI 应用完全指南

## 开篇引言

在上一篇文章《[Next.js 16 与 Prisma](https://...)》中，我们介绍了 Next.js 16 与 Prisma 的基础集成，包括如何使用 Prisma 进行数据库操作、如何利用 Next.js 16 的 Server Actions 简化 API 开发等内容。

继上次深入剖析之后，我们将探索如何使用 Next.js 16 和 Prisma 构建现代化的 AI 应用。随着大语言模型（LLM）技术的飞速发展，AI 应用已成为当今技术领域的热门话题。而 Next.js 16 与 Vercel AI SDK 的结合，为开发者提供了构建 AI 应用的完整解决方案。

接下来，我们进一步深入探讨 Next.js 16 + Prisma + AI 的应用开发，将探讨以下几个核心问题：

- AI 应用与传统的 Web 应用有何区别？
- Next.js 16 与 Vercel AI SDK 如何协同工作？
- 如何使用 Prisma 管理向量数据和对话历史？
- 如何构建一个完整的 RAG（检索增强生成）应用？

---

## 一、AI 应用开发概述

### 什么是 AI 应用

AI 应用是指集成了大语言模型（LLM）或其他人工智能能力的 Web 应用程序。与传统应用相比，AI 应用具备以下特点：

- **自然语言交互**：用户可以使用自然语言与应用进行对话式交互，无需学习复杂的界面操作
- **智能内容生成**：应用能够根据上下文生成文本、代码、图像等内容，而不仅仅是检索预存的数据
- **语义理解能力**：应用能够理解用户意图，而非简单的关键词匹配
- **上下文感知**：应用能够记住对话历史，提供连贯的交互体验

### AI 应用与传统应用的对比

为了更清晰地理解 AI 应用的特点，我们将其与传统 Web 应用进行对比：

| 特性 | 传统 Web 应用 | AI 应用 |
|------|--------------|---------|
| **用户输入** | 表单、按钮、点击 | 自然语言文本 |
| **数据处理** | 结构化数据（JSON/SQL） | 非结构化数据 + 向量 |
| **响应方式** | 预定义逻辑 | 动态生成内容 |
| **渲染方式** | CSR/SSR | SSR + 流式响应 |
| **API 调用** | REST/GraphQL | LLM API + 向量检索 |
| **状态管理** | 客户端/服务端状态 | 对话上下文 + 向量库 |

### Next.js 16 与 AI 开发的契合点

Next.js 16 的多项新特性使其成为构建 AI 应用的理想框架：

1. **Server Actions**：简化了 API 调用，无需手动创建路由处理器，可以直接从客户端调用服务端函数
2. **流式渲染**：支持 Server-Sent Events（SSE），适合 AI 的流式响应，实现打字机效果
3. **React Server Components**：减少客户端 JS，提升首屏性能，让 AI 应用加载更快
4. **增强的缓存机制**：减少 AI API 调用，降低成本
5. **边缘运行时支持**：低延迟的全球部署

### 常见 AI 应用类型

根据功能和复杂度，AI 应用可以分为以下几类：

| 应用类型 | 描述 | 典型场景 | 技术要点 | 复杂度 |
|---------|------|---------|---------|--------|
| **聊天机器人**（Chatbot） | 最基础的 AI 应用形态，实现人与 AI 的对话交互 | 客服助手、智能问答、语言学习 | 流式响应、对话历史管理 | 低 |
| **RAG 应用**（检索增强生成） | 结合向量检索和 LLM 生成，基于私有知识库回答问题 | 企业知识库、文档问答、技术支持 | 文档嵌入、向量检索、提示工程 | 中 |
| **AI Agent**（智能代理） | 具备工具调用能力的 AI，可以执行具体操作 | 个人助理、任务自动化、代码生成 | 工具调用、状态管理、错误处理 | 高 |
| **生成式 UI**（Generative UI） | AI 驱动的动态界面，根据上下文渲染不同组件 | 动态表单、自适应布局、个性化推荐 | 结构化输出、组件映射、流式 UI | 中高 |
| **图像/多模态应用** | 处理和生成图像、音频、视频等多模态内容 | 图片生成、图像分析、视频摘要 | 多模态模型、文件处理、媒体存储 | 中 |

---

## 二、项目初始化与环境配置

### 技术栈选择

在构建 AI 应用时，我们需要选择合适的技术栈：

| 技术 | 用途 | 推荐选择 |
|------|------|---------|
| **前端框架** | React 应用框架 | Next.js 16 |
| **ORM** | 数据库操作 | Prisma |
| **LLM SDK** | 大语言模型集成 | Vercel AI SDK |
| **向量数据库** | 语义搜索 | Pgvector / Pinecone |
| **数据库** | 关系型数据存储 | PostgreSQL |
| **缓存** | 响应缓存 | Redis / Vercel KV |
| **认证** | 用户管理 | NextAuth.js |

### 创建项目

按照本系列文章的惯例，在进入正题之前，我们先来准备一下相关的环境。下面是本文内容的搭建环境：

> - Node.js：v20.10.0
> - pnpm：10.4.1
> - OS：MacBook Pro
> - IDE：VS Code 1.99.3

使用命令创建一个新的项目：

```bash
# 创建 Next.js 项目
npx create-next-app@latest ai-app --typescript --tailwind --app

# 进入项目目录
cd ai-app

# 安装核心依赖
pnpm add @prisma/client @ai-sdk/openai ai @ai-sdk/react
pnpm add -D prisma

# 初始化 Prisma
npx prisma init
```

运行上述命令后，会出现如下配置选项：

![](./assets/project-setup.png)

> 上图可以看到，项目创建选项包括 TypeScript、Tailwind CSS、App Router 等配置。这些都是构建现代 AI 应用的推荐配置。

### 环境变量配置

创建 `.env.local` 文件，配置必要的环境变量：

```env
# 数据库连接
DATABASE_URL="postgresql://user:password@localhost:5432/ai_app"

# OpenAI API
OPENAI_API_KEY="sk-proj-..."

# Anthropic API（可选）
ANTHROPIC_API_KEY="sk-ant-..."

# 向量数据库（如果使用 Pinecone）
PINECONE_API_KEY="..."
PINECONE_ENVIRONMENT="..."

# Vercel KV（用于缓存，可选）
KV_URL="..."
KV_REST_API_URL="..."
KV_REST_API_TOKEN="..."
```

> 注意：不要将 `.env.local` 文件提交到版本控制系统。项目创建时会自动将其添加到 `.gitignore` 中。

### 项目目录结构

初始化完成后，我们的项目目录结构如下：

```text
ai-app/
├── prisma/
│   ├── schema.prisma          # Prisma 数据模型
│   └── migrations/            # 数据库迁移文件
├── src/
│   ├── app/
│   │   ├── api/               # API 路由
│   │   │   ├── chat/          # 聊天 API
│   │   │   ├── rag/           # RAG API
│   │   │   └── agent/         # Agent API
│   │   ├── chat/              # 聊天页面
│   │   ├── layout.tsx
│   │   └── page.tsx
│   ├── components/
│   │   ├── Chat.tsx           # 聊天组件
│   │   ├── MessageList.tsx    # 消息列表
│   │   └── InputBox.tsx       # 输入框
│   ├── lib/
│   │   ├── prisma.ts          # Prisma 客户端
│   │   ├── embeddings.ts      # 向量嵌入
│   │   ├── vector-search.ts   # 向量检索
│   │   └── chat-history.ts    # 对话历史
│   └── types/
│       └── index.ts           # TypeScript 类型
├── .env.local
├── next.config.ts
├── package.json
└── tsconfig.json
```

> 本文中的所有代码都能在 <https://github.com/clin211/next-awesome/tree/main/ai-app> 上找到！

---

## 三、数据库设计与 Prisma 配置

### Prisma Schema 设计

AI 应用的数据模型需要支持以下核心功能：

- 用户管理
- 对话历史存储
- 消息记录
- 文档存储（用于 RAG）
- 向量嵌入

以下是完整的 Prisma Schema：

```prisma
// prisma/schema.prisma

generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

// 用户模型
model User {
  id            String        @id @default(cuid())
  email         String        @unique
  name          String?
  conversations Conversation[]
  documents     Document[]
  createdAt     DateTime      @default(now())
  updatedAt     DateTime      @updatedAt
}

// 对话模型
model Conversation {
  id        String   @id @default(cuid())
  userId    String
  user      User     @relation(fields: [userId], references: [id], onDelete: Cascade)
  title     String?
  messages  Message[]
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt

  @@index([userId])
}

// 消息模型
model Message {
  id             String       @id @default(cuid())
  conversationId String
  conversation   Conversation @relation(fields: [conversationId], references: [id], onDelete: Cascade)
  role           String       // user | assistant | system
  content        String       @db.Text
  tokens         Int?
  metadata       Json?        // 存储额外的元数据，如工具调用等
  createdAt      DateTime     @default(now())

  @@index([conversationId])
}

// 文档模型（用于 RAG）
model Document {
  id        String   @id @default(cuid())
  userId    String
  user      User     @relation(fields: [userId], references: [id], onDelete: Cascade)
  title     String
  content   String   @db.Text
  embedding Unsupported("vector(1536)")?
  metadata  Json?    // 存储文档元数据
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt

  @@index([userId])
}

// Agent 工具调用记录
model ToolCall {
  id        String   @id @default(cuid())
  name      String   // 工具名称
  arguments Json     // 工具参数
  result    Json?    // 工具执行结果
  messageId String
  message   Message  @relation(fields: [messageId], references: [id], onDelete: Cascade)
  createdAt DateTime @default(now())

  @@index([messageId])
}
```

> 在上述 Schema 中，我们使用了 `Unsupported("vector(1536)")` 来声明 pgvector 的向量类型。这样 Prisma 不会验证该字段，但可以在原始 SQL 查询中使用。

### 向量扩展配置

为了在 PostgreSQL 中存储和查询向量数据，我们需要启用 `pgvector` 扩展：

```sql
-- 启用 Pgvector 扩展
CREATE EXTENSION IF NOT EXISTS vector;

-- 为 Document 表创建向量索引
-- 使用 IVFFlat 索引方法，适合大规模向量检索
CREATE INDEX document_embedding_idx
ON "Document"
USING ivfflat (embedding vector_cosine_ops)
WITH (lists = 100);

-- 如果数据量较小，也可以使用 HNSW 索引
-- CREATE INDEX document_embedding_hnsw_idx
-- ON "Document"
-- USING hnsw (embedding vector_cosine_ops);
```

> IVFFlat 和 HNSW 是两种常用的向量索引方法：
>
> - **IVFFlat**：适合大规模数据（百万级以上），构建速度快
> - **HNSW**：适合中小规模数据，查询精度更高

### 数据库迁移

```bash
# 创建迁移
npx prisma migrate dev --name init

# 生成 Prisma Client
npx prisma generate

# （可选）打开 Prisma Studio 查看数据
npx prisma studio
```

执行迁移命令后，Prisma 会自动创建数据库表并应用 pgvector 扩展。如果成功，你会看到如下输出：

![](./assets/prisma-migrate.png)

### Prisma 客户端配置

为了避免在开发环境中创建多个 Prisma Client 实例，我们需要进行如下配置：

```typescript
// src/lib/prisma.ts
import { PrismaClient } from '@prisma/client'

const globalForPrisma = globalThis as unknown as {
  prisma: PrismaClient | undefined
}

export const prisma =
  globalForPrisma.prisma ??
  new PrismaClient({
    log: process.env.NODE_ENV === 'development' ? ['query', 'error', 'warn'] : ['error'],
  })

if (process.env.NODE_ENV !== 'production') globalForPrisma.prisma = prisma
```

> 在开发环境中，由于热重载可能会创建多个 Prisma Client 实例，上述配置可以确保全局只有一个实例。

---

## 四、Vercel AI SDK 核心概念

### 什么是 Vercel AI SDK

Vercel AI SDK 是一个专为 Next.js 设计的 AI 应用开发工具包，提供了：

- **统一的 LLM 接口**：支持 OpenAI、Anthropic、Google Gemini 等多种模型
- **流式响应处理**：内置 Server-Sent Events 支持
- **工具调用**：声明式定义 AI 可使用的工具
- **生成式 UI**：基于 AI 输出动态渲染组件
- **类型安全**：完整的 TypeScript 支持

### 核心组件对比

Vercel AI SDK 提供了多个核心组件，每个组件都有其特定的用途：

| 组件 | 用途 | 返回类型 | 适用场景 |
|------|------|---------|---------|
| `streamText` | 流式文本生成 | `DataStreamResponse` | 聊天、实时响应 |
| `generateText` | 非流式文本生成 | `GenerateTextResult` | 批量处理、后台任务 |
| `useChat` | 聊天 UI Hook | `UseChatReturn` | 聊天界面 |
| `useCompletion` | 补全 UI Hook | `UseCompletionReturn` | 文本补全 |
| `generateObject` | 结构化输出 | `GenerateObjectResult` | 数据提取、分类 |
| `streamUI` | 生成式 UI | `UIMessage` | 动态界面 |

### 与 Next.js 16 的集成优势

#### 1. Server Actions 无缝集成

使用 Server Actions 可以极大地简化 AI API 的调用：

```typescript
// 使用 Server Actions 简化 API 调用
'use server'

import { openai } from '@ai-sdk/openai'
import { generateText } from 'ai'

export async function generateResponse(prompt: string) {
  const { text } = await generateText({
    model: openai('gpt-4o'),
    prompt,
  })
  return text
}
```

#### 2. RSC 友好

Vercel AI SDK 与 React Server Components 完美集成：

```typescript
// 在 Server Component 中直接调用
import { openai } from '@ai-sdk/openai'
import { generateText } from 'ai'

export default async function Page() {
  const { text } = await generateText({
    model: openai('gpt-4o'),
    prompt: '生成一段欢迎语',
  })

  return <div>{text}</div>
}
```

#### 3. 边缘运行时支持

在边缘运行时执行，可以降低延迟：

```typescript
// 在边缘运行时执行，降低延迟
export const runtime = 'edge'

import { openai } from '@ai-sdk/openai'
import { streamText } from 'ai'

export async function POST(req: Request) {
  const { messages } = await req.json()

  const result = streamText({
    model: openai('gpt-4o'),
    messages,
  })

  return result.toDataStreamResponse()
}
```

### 模型提供商配置

Vercel AI SDK 支持多种模型提供商，以下展示了如何配置 OpenAI 和 Anthropic：

```typescript
// src/lib/ai.ts
import { createOpenAI } from '@ai-sdk/openai'
import { createAnthropic } from '@ai-sdk/anthropic'

// OpenAI 配置
export const openai = createOpenAI({
  apiKey: process.env.OPENAI_API_KEY,
  compatibility: 'strict', // 启用严格模式
})

// Anthropic 配置
export const anthropic = createAnthropic({
  apiKey: process.env.ANTHROPIC_API_KEY,
})

// 模型配置
export const models = {
  gpt4o: openai('gpt-4o'),
  gpt4oMini: openai('gpt-4o-mini'),
  claudeOpus: anthropic('claude-3-opus-20240229'),
  claudeSonnet: anthropic('claude-3-5-sonnet-20241022'),
}
```

> `compatibility: 'strict'` 模式可以确保与 OpenAI API 的完全兼容，适合生产环境使用。

---

## 五、构建聊天机器人应用

### 基础聊天实现

#### API 路由实现

首先，我们创建一个基础的聊天 API：

```typescript
// src/app/api/chat/route.ts
import { openai } from '@ai-sdk/openai'
import { streamText } from 'ai'

export async function POST(req: Request) {
  const { messages } = await req.json()

  const result = streamText({
    model: openai('gpt-4o'),
    messages,
    temperature: 0.7,
    maxTokens: 1000,
  })

  return result.toDataStreamResponse()
}
```

#### 前端聊天组件

接下来，创建一个聊天界面组件：

```typescript
// src/components/Chat.tsx
'use client'

import { useChat } from 'ai/react'
import { useEffect, useRef } from 'react'

export function Chat() {
  const { messages, input, handleInputChange, handleSubmit, isLoading } = useChat({
    api: '/api/chat',
  })

  const messagesEndRef = useRef<HTMLDivElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  return (
    <div className="flex flex-col h-screen max-w-2xl mx-auto p-4">
      {/* 消息列表 */}
      <div className="flex-1 overflow-y-auto mb-4 space-y-4">
        {messages.map(m => (
          <div
            key={m.id}
            className={`flex ${m.role === 'user' ? 'justify-end' : 'justify-start'}`}
          >
            <div
              className={`max-w-[80%] rounded-lg px-4 py-2 ${
                m.role === 'user'
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-200 text-gray-800'
              }`}
            >
              <p className="whitespace-pre-wrap">{m.content}</p>
            </div>
          </div>
        ))}
        {isLoading && (
          <div className="flex justify-start">
            <div className="bg-gray-200 rounded-lg px-4 py-2">
              <div className="flex space-x-2">
                <div className="w-2 h-2 bg-gray-500 rounded-full animate-bounce" />
                <div className="w-2 h-2 bg-gray-500 rounded-full animate-bounce delay-100" />
                <div className="w-2 h-2 bg-gray-500 rounded-full animate-bounce delay-200" />
              </div>
            </div>
          </div>
        )}
        <div ref={messagesEndRef} />
      </div>

      {/* 输入框 */}
      <form onSubmit={handleSubmit} className="flex gap-2">
        <input
          value={input}
          onChange={handleInputChange}
          placeholder="输入你的消息..."
          className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          disabled={isLoading}
        />
        <button
          type="submit"
          disabled={isLoading || !input.trim()}
          className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300"
        >
          发送
        </button>
      </form>
    </div>
  )
}
```

效果如下：

![](./assets/chat-demo.png)

> `useChat` Hook 是 Vercel AI SDK 提供的便捷工具，它自动处理了流式响应、消息状态管理等复杂逻辑。

### 历史记录持久化

#### 保存对话到数据库

为了实现对话历史的持久化，我们需要创建相应的服务函数：

```typescript
// src/lib/chat-history.ts
import { prisma } from '@/lib/prisma'
import { CoreMessage } from 'ai'

export async function saveConversation(
  userId: string,
  messages: CoreMessage[],
  title?: string
) {
  // 使用第一条用户消息作为标题
  const conversationTitle =
    title || messages.find(m => m.role === 'user')?.content.slice(0, 50) + '...'

  const conversation = await prisma.conversation.create({
    data: {
      userId,
      title: conversationTitle,
      messages: {
        create: messages
          .filter(m => m.role !== 'system') // 不保存系统消息
          .map(m => ({
            role: m.role,
            content: m.content as string,
          })),
      },
    },
    include: {
      messages: true,
    },
  })

  return conversation
}

export async function getUserConversations(userId: string) {
  return prisma.conversation.findMany({
    where: { userId },
    include: {
      messages: {
        orderBy: { createdAt: 'asc' },
        take: 1, // 只获取第一条消息作为预览
      },
    },
    orderBy: { updatedAt: 'desc' },
  })
}

export async function getConversationMessages(conversationId: string) {
  return prisma.message.findMany({
    where: { conversationId },
    orderBy: { createdAt: 'asc' },
  })
}

export async function addMessageToConversation(
  conversationId: string,
  role: 'user' | 'assistant',
  content: string,
  tokens?: number
) {
  return prisma.message.create({
    data: {
      conversationId,
      role,
      content,
      tokens,
    },
  })
}
```

#### API 集成历史记录

然后，我们在 API 路由中集成历史记录功能：

```typescript
// src/app/api/chat/route.ts
import { openai } from '@ai-sdk/openai'
import { streamText } from 'ai'
import { saveConversation, addMessageToConversation } from '@/lib/chat-history'
import { getServerSession } from 'next-auth'

export async function POST(req: Request) {
  const session = await getServerSession()
  if (!session?.user) {
    return new Response('Unauthorized', { status: 401 })
  }

  const { messages, conversationId } = await req.json()

  // 保存用户消息
  const lastMessage = messages[messages.length - 1]
  if (conversationId && lastMessage.role === 'user') {
    await addMessageToConversation(conversationId, 'user', lastMessage.content)
  }

  const result = streamText({
    model: openai('gpt-4o'),
    messages,
    onFinish: async ({ text, usage }) => {
      // 保存助手消息
      if (conversationId) {
        await addMessageToConversation(conversationId, 'assistant', text, usage.totalTokens)
      } else {
        // 创建新对话
        await saveConversation(session.user.id, messages)
      }
    },
  })

  return result.toDataStreamResponse()
}
```

### 上下文管理

在多轮对话中，我们需要合理管理上下文，避免 Token 消耗过大：

```typescript
// src/lib/context.ts
import { CoreMessage } from 'ai'

const MAX_CONTEXT_MESSAGES = 10
const MAX_CONTEXT_TOKENS = 4000

export function trimMessages(messages: CoreMessage[], maxTokens: number = MAX_CONTEXT_TOKENS) {
  // 简单的消息截断策略
  // 保留系统消息和最近的 N 条消息
  const systemMessages = messages.filter(m => m.role === 'system')
  const conversationMessages = messages.filter(m => m.role !== 'system')

  const trimmedMessages = conversationMessages.slice(-MAX_CONTEXT_MESSAGES)

  return [...systemMessages, ...trimmedMessages]
}

export function estimateTokens(text: string): number {
  // 粗略估计：英文约 4 字符/token，中文约 2 字符/token
  const chineseChars = (text.match(/[\u4e00-\u9fa5]/g) || []).length
  const otherChars = text.length - chineseChars
  return Math.ceil(chineseChars / 2 + otherChars / 4)
}
```

> 上下文管理是 AI 应用的重要环节。合理的上下文管理可以在保证对话连贯性的同时，控制 Token 消耗。

---

## 六、实现 RAG（检索增强生成）

### RAG 架构说明

RAG（Retrieval-Augmented Generation，检索增强生成）是一种结合信息检索和文本生成的技术架构。

```
┌─────────────────────────────────────────────────────────────┐
│                        RAG 工作流程                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  用户查询                                                     │
│     │                                                       │
│     ▼                                                       │
│  ┌─────────┐                                               │
│  │ 向量化   │ ──> 查询向量                                   │
│  └─────────┘                                               │
│     │                                                       │
│     ▼                                                       │
│  ┌─────────────┐                                           │
│  │ 向量相似度检索 │ ──> 召回相关文档                           │
│  └─────────────┘                                           │
│     │                                                       │
│     ▼                                                       │
│  ┌─────────────────────────────────┐                       │
│  │ 提示词工程 + 检索到的文档上下文    │                       │
│  └─────────────────────────────────┘                       │
│     │                                                       │
│     ▼                                                       │
│  ┌─────────┐                                               │
│  │ LLM 生成 │ ──> 带引用的答案                               │
│  └─────────┘                                               │
│     │                                                       │
│     ▼                                                       │
│  返回答案 + 源文档引用                                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 文档嵌入实现

首先，我们需要实现文档的向量化功能：

```typescript
// src/lib/embeddings.ts
import { openai } from '@ai-sdk/openai'
import { embed } from 'ai'
import { prisma } from '@/lib/prisma'

// 文本分块策略
export function chunkText(text: string, chunkSize: number = 500, overlap: number = 50): string[] {
  const chunks: string[] = []
  let start = 0

  while (start < text.length) {
    const end = Math.min(start + chunkSize, text.length)
    chunks.push(text.slice(start, end))
    start = end - overlap
  }

  return chunks
}

// 生成文本嵌入
export async function embedText(text: string): Promise<number[]> {
  const { embedding } = await embed({
    model: openai.embedding('text-embedding-3-small'),
    value: text,
  })

  return embedding
}

// 保存文档到数据库
export async function saveDocument(
  userId: string,
  title: string,
  content: string,
  metadata?: Record<string, any>
) {
  // 分块处理
  const chunks = chunkText(content)

  // 批量创建文档（每个分块一条记录）
  const documents = await prisma.document.createMany({
    data: await Promise.all(
      chunks.map(async (chunk) => {
        const embedding = await embedText(chunk)
        return {
          userId,
          title: `${title} [Chunk ${chunks.indexOf(chunk) + 1}]`,
          content: chunk,
          embedding: JSON.stringify(embedding),
          metadata: {
            ...metadata,
            originalTitle: title,
            totalChunks: chunks.length,
          },
        }
      })
    ),
  })

  return documents
}

// 批量导入文档
export async function importDocuments(
  userId: string,
  documents: Array<{ title: string; content: string; metadata?: Record<string, any> }>
) {
  const results = []

  for (const doc of documents) {
    try {
      await saveDocument(userId, doc.title, doc.content, doc.metadata)
      results.push({ success: true, title: doc.title })
    } catch (error) {
      results.push({ success: false, title: doc.title, error })
    }
  }

  return results
}
```

### 向量相似度检索

接下来，实现向量相似度检索功能：

```typescript
// src/lib/vector-search.ts
import { prisma } from '@/lib/prisma'
import { embedText } from './embeddings'

export interface SearchResult {
  id: string
  title: string
  content: string
  similarity: number
  metadata?: Record<string, any>
}

// 使用余弦相似度搜索文档
export async function searchDocuments(
  userId: string,
  query: string,
  limit: number = 5,
  threshold: number = 0.7
): Promise<SearchResult[]> {
  // 生成查询向量
  const queryEmbedding = await embedText(query)

  // 使用原生 SQL 查询进行向量相似度搜索
  const results = await prisma.$queryRaw<Array<any>>`
    SELECT
      id,
      title,
      content,
      metadata,
      1 - (embedding <=> ${JSON.stringify(queryEmbedding)}::vector) as similarity
    FROM "Document"
    WHERE "userId" = ${userId}
      AND embedding IS NOT NULL
    ORDER BY embedding <=> ${JSON.stringify(queryEmbedding)}::vector
    LIMIT ${limit}
  `

  // 过滤低于阈值的结果
  return results
    .filter(r => r.similarity >= threshold)
    .map(r => ({
      id: r.id,
      title: r.title,
      content: r.content,
      similarity: parseFloat(r.similarity),
      metadata: r.metadata,
    }))
}

// 混合搜索（向量 + 关键词）
export async function hybridSearch(
  userId: string,
  query: string,
  limit: number = 5
): Promise<SearchResult[]> {
  // 向量搜索
  const vectorResults = await searchDocuments(userId, query, limit)

  // 关键词搜索（使用 PostgreSQL 的全文搜索）
  const keywordResults = await prisma.document.findMany({
    where: {
      userId,
      OR: [
        { title: { contains: query, mode: 'insensitive' } },
        { content: { contains: query, mode: 'insensitive' } },
      ],
    },
    take: limit,
  })

  // 合并并去重结果
  const combinedMap = new Map<string, SearchResult>()

  vectorResults.forEach(r => combinedMap.set(r.id, r))

  keywordResults.forEach(r => {
    if (!combinedMap.has(r.id)) {
      combinedMap.set(r.id, {
        id: r.id,
        title: r.title,
        content: r.content,
        similarity: 0.5, // 关键词匹配的基础相似度
        metadata: r.metadata as Record<string, any>,
      })
    }
  })

  return Array.from(combinedMap.values())
    .sort((a, b) => b.similarity - a.similarity)
    .slice(0, limit)
}
```

> `<=>` 操作符是 pgvector 提供的向量距离运算符，用于计算两个向量之间的余弦距离。`1 - distance` 就是余弦相似度。

### RAG 聊天实现

最后，我们实现完整的 RAG 聊天功能：

```typescript
// src/app/api/rag-chat/route.ts
import { openai } from '@ai-sdk/openai'
import { streamText, system } from 'ai'
import { getServerSession } from 'next-auth'
import { searchDocuments } from '@/lib/vector-search'

// 构建带引用的提示词
function buildRAGPrompt(context: string, query: string): string {
  return `请根据以下参考文档回答用户的问题。如果参考文档中没有相关信息，请明确说明。

参考文档：
${context}

用户问题：${query}

请提供准确、详细的回答，并在回答中注明引用的来源。`
}

export async function POST(req: Request) {
  const session = await getServerSession()
  if (!session?.user) {
    return new Response('Unauthorized', { status: 401 })
  }

  const { messages } = await req.json()
  const lastMessage = messages[messages.length - 1]

  // 检索相关文档
  const searchResults = await searchDocuments(
    session.user.id,
    lastMessage.content,
    5,
    0.6 // 相似度阈值
  )

  // 构建上下文
  const context = searchResults
    .map(r => `【${r.title}】\n${r.content}`)
    .join('\n\n')

  // 保存引用信息
  const citations = searchResults.map(r => ({
    id: r.id,
    title: r.title,
    similarity: r.similarity,
  }))

  const result = streamText({
    model: openai('gpt-4o'),
    system: system(`你是一个有帮助的 AI 助手，擅长基于提供的文档回答问题。
    回答时请遵循以下原则：
    1. 优先使用提供的文档信息
    2. 如果文档信息不足，明确说明
    3. 在回答中引用相关文档
    4. 保持回答准确、客观`),
    messages: [
      ...messages.slice(0, -1),
      {
        role: 'user',
        content: buildRAGPrompt(context, lastMessage.content),
      },
    ],
    experimental_telemetry: {
      isEnabled: true,
      functionId: 'rag-chat',
    },
  })

  // 在响应头中添加引用信息
  const response = await result.toDataStreamResponse()
  response.headers.set('X-Citations', JSON.stringify(citations))

  return response
}
```

### RAG 前端组件

最后，创建 RAG 聊天的前端组件：

```typescript
// src/components/RagChat.tsx
'use client'

import { useChat } from 'ai/react'
import { useState } from 'react'

interface Citation {
  id: string
  title: string
  similarity: number
}

export function RagChat() {
  const { messages, input, handleInputChange, handleSubmit, isLoading } = useChat({
    api: '/api/rag-chat',
    onResponse: (response) => {
      // 提取引用信息
      const citationsHeader = response.headers.get('X-Citations')
      if (citationsHeader) {
        setCitations(JSON.parse(citationsHeader))
      }
    },
  })

  const [citations, setCitations] = useState<Citation[]>([])

  return (
    <div className="flex flex-col h-screen max-w-4xl mx-auto p-4">
      {/* 消息列表 */}
      <div className="flex-1 overflow-y-auto mb-4 space-y-4">
        {messages.map(m => (
          <div
            key={m.id}
            className={`flex ${m.role === 'user' ? 'justify-end' : 'justify-start'}`}
          >
            <div
              className={`max-w-[80%] rounded-lg px-4 py-2 ${
                m.role === 'user'
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-100 text-gray-800'
              }`}
            >
              <p className="whitespace-pre-wrap">{m.content}</p>
            </div>
          </div>
        ))}

        {/* 引用文档显示 */}
        {citations.length > 0 && messages.length > 0 && (
          <div className="mt-4 p-4 bg-yellow-50 rounded-lg border border-yellow-200">
            <h3 className="font-semibold text-sm text-yellow-800 mb-2">参考文档</h3>
            <ul className="space-y-1">
              {citations.map((c, i) => (
                <li key={c.id} className="text-xs text-yellow-700">
                  <span className="font-medium">{i + 1}. {c.title}</span>
                  <span className="ml-2 text-yellow-600">
                    (相似度: {(c.similarity * 100).toFixed(1)}%)
                  </span>
                </li>
              ))}
            </ul>
          </div>
        )}

        {isLoading && (
          <div className="text-gray-500">正在搜索文档并生成回答...</div>
        )}
      </div>

      {/* 输入框 */}
      <form onSubmit={handleSubmit} className="flex gap-2">
        <input
          value={input}
          onChange={handleInputChange}
          placeholder="输入你的问题..."
          className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          disabled={isLoading}
        />
        <button
          type="submit"
          disabled={isLoading || !input.trim()}
          className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300"
        >
          提问
        </button>
      </form>
    </div>
  )
}
```

RAG 应用的效果如下：

![](./assets/rag-demo.png)

> RAG 应用可以大大提升 AI 回答的准确性和可信度，因为它基于真实的文档内容生成答案，而不是仅依赖模型的训练数据。

---

## 七、AI Agent 与工具调用

### 什么是 AI Agent

AI Agent 是一种具备工具调用能力的 AI 系统，它可以：

- **理解用户意图**：分析用户需求，决定执行哪些操作
- **调用外部工具**：执行数据库查询、API 调用、文件操作等
- **多步推理**：将复杂任务分解为多个步骤
- **状态管理**：维护对话上下文和执行状态

| 特性 | 普通聊天机器人 | AI Agent |
|------|--------------|----------|
| **能力范围** | 对话生成 | 对话 + 执行操作 |
| **工具支持** | 无 | 可调用多种工具 |
| **决策能力** | 被动响应 | 主动规划 |
| **应用场景** | 问答、对话 | 任务自动化 |

### 工具定义与配置

Vercel AI SDK 使用 Zod Schema 来定义工具的参数：

```typescript
// src/lib/agent-tools.ts
import { z } from 'zod'

// 天气查询工具
export const weatherTool = {
  description: '获取指定城市的当前天气信息',
  parameters: z.object({
    city: z.string().describe('城市名称，例如：北京、上海'),
  }),
  execute: async ({ city }: { city: string }) => {
    // 模拟天气 API 调用
    return {
      city,
      temperature: 25,
      description: '晴天',
      humidity: 60,
    }
  },
}

// 数据库查询工具
export const databaseQueryTool = {
  description: '查询产品数据库，根据关键词搜索产品',
  parameters: z.object({
    keyword: z.string().describe('搜索关键词'),
    category: z.string().optional().describe('产品分类（可选）'),
  }),
  execute: async ({ keyword, category }: { keyword: string; category?: string }) => {
    // 实际应用中这里会连接真实数据库
    return [
      { id: 1, name: 'iPhone 15', price: 5999, category: '手机' },
      { id: 2, name: 'MacBook Pro', price: 12999, category: '电脑' },
    ].filter(p => p.name.includes(keyword) && (!category || p.category === category))
  },
}

// 邮件发送工具
export const sendEmailTool = {
  description: '发送邮件给指定收件人',
  parameters: z.object({
    to: z.string().email().describe('收件人邮箱地址'),
    subject: z.string().describe('邮件主题'),
    body: z.string().describe('邮件正文'),
  }),
  execute: async ({ to, subject, body }: { to: string; subject: string; body: string }) => {
    // 实际应用中这里会调用真实的邮件服务
    console.log(`发送邮件到 ${to}: ${subject}`)
    return { success: true, message: '邮件已发送' }
  },
}

// 导出所有工具
export const agentTools = {
  weather: weatherTool,
  databaseQuery: databaseQueryTool,
  sendEmail: sendEmailTool,
}
```

> 工具的 `description` 非常重要，它告诉 AI 何时以及如何使用该工具。一个好的描述可以显著提升 Agent 的表现。

### Agent 实现

接下来，我们实现 Agent 的 API 路由：

```typescript
// src/app/api/agent/route.ts
import { openai } from '@ai-sdk/openai'
import { generateText } from 'ai'
import { agentTools } from '@/lib/agent-tools'

export async function POST(req: Request) {
  const { messages } = await req.json()

  const result = await generateText({
    model: openai('gpt-4o'),
    system: `你是一个智能助手，可以帮助用户完成各种任务。
    你可以使用以下工具：
    - weather: 查询天气信息
    - databaseQuery: 查询产品数据库
    - sendEmail: 发送邮件

    请根据用户需求选择合适的工具，并解释你的操作过程。`,
    messages,
    tools: agentTools,
    maxToolRoundtrips: 5, // 最多进行 5 轮工具调用
    experimental_telemetry: {
      isEnabled: true,
    },
  })

  return Response.json({
    text: result.text,
    toolCalls: result.toolCalls,
    usage: result.usage,
  })
}
```

### 前端 Agent 组件

最后，创建 Agent 的前端组件：

```typescript
// src/components/Agent.tsx
'use client'

import { useChat } from 'ai/react'

export function Agent() {
  const { messages, input, handleInputChange, handleSubmit, isLoading } = useChat({
    api: '/api/agent',
  })

  return (
    <div className="max-w-2xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">AI Agent 助手</h1>

      {/* 消息列表 */}
      <div className="space-y-4 mb-4">
        {messages.map(m => (
          <div
            key={m.id}
            className={`p-4 rounded-lg ${
              m.role === 'user' ? 'bg-blue-50 ml-8' : 'bg-gray-50 mr-8'
            }`}
          >
            <p className="font-semibold mb-2">{m.role === 'user' ? '用户' : '助手'}</p>
            <p className="whitespace-pre-wrap">{m.content}</p>

            {/* 显示工具调用 */}
            {m.toolCalls && m.toolCalls.length > 0 && (
              <div className="mt-2 p-2 bg-yellow-50 rounded text-sm">
                <p className="font-semibold">调用工具:</p>
                {m.toolCalls.map((call: any, i: number) => (
                  <div key={i} className="ml-2">
                    <span className="font-mono">{call.toolName}</span>
                    <span className="text-gray-600">
                      ({JSON.stringify(call.args)})
                    </span>
                  </div>
                ))}
              </div>
            )}
          </div>
        ))}

        {isLoading && (
          <div className="text-center text-gray-500">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500" />
            <p className="mt-2">正在处理...</p>
          </div>
        )}
      </div>

      {/* 输入框 */}
      <form onSubmit={handleSubmit} className="flex gap-2">
        <input
          value={input}
          onChange={handleInputChange}
          placeholder="告诉我你需要什么帮助..."
          className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          disabled={isLoading}
        />
        <button
          type="submit"
          disabled={isLoading || !input.trim()}
          className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300"
        >
          发送
        </button>
      </form>

      {/* 快捷操作 */}
      <div className="mt-4 flex flex-wrap gap-2">
        <button
          onClick={() => {
            const form = new FormData()
            form.set('message', '北京今天天气怎么样？')
            handleSubmit(new Event('submit'), { body: form })
          }}
          className="px-3 py-1 text-sm bg-gray-100 rounded hover:bg-gray-200"
        >
          查天气
        </button>
        <button
          onClick={() => {
            const form = new FormData()
            form.set('message', '搜索 iPhone 15 产品')
            handleSubmit(new Event('submit'), { body: form })
          }}
          className="px-3 py-1 text-sm bg-gray-100 rounded hover:bg-gray-200"
        >
          搜索产品
        </button>
        <button
          onClick={() => {
            const form = new FormData()
            form.set('message', '发送邮件给客户')
            handleSubmit(new Event('submit'), { body: form })
          }}
          className="px-3 py-1 text-sm bg-gray-100 rounded hover:bg-gray-200"
        >
          发邮件
        </button>
      </div>
    </div>
  )
}
```

Agent 应用的效果如下：

![](./assets/agent-demo.png)

> AI Agent 是 AI 应用的进阶形态。通过工具调用，AI 可以执行实际操作，而不仅仅是生成文本。

---

## 八、性能优化与成本控制

### 缓存策略

AI API 调用成本较高，合理的缓存策略可以大幅降低成本。

| 缓存类型 | 适用场景 | 存储位置 | 过期时间 |
|---------|---------|---------|---------|
| **响应缓存** | 相同查询的响应 | Redis / Vercel KV | 1-24 小时 |
| **嵌入缓存** | 文档向量嵌入 | 数据库 | 永久 |
| **提示词缓存** | 预定义的提示词 | 内存 / Redis | 永久 |
| **用户会话缓存** | 对话上下文 | Redis | 30 分钟 |

### Token 使用优化

#### 提示词压缩

```typescript
// 压缩提示词
export function compressPrompt(prompt: string, targetLength: number = 1000): string {
  if (prompt.length <= targetLength) {
    return prompt
  }

  const keepStart = Math.floor(targetLength * 0.4)
  const keepEnd = Math.floor(targetLength * 0.4)

  return (
    prompt.slice(0, keepStart) +
    `\n... [省略 ${prompt.length - keepStart - keepEnd} 字符] ...\n` +
    prompt.slice(-keepEnd)
  )
}

// 优化消息历史
export function optimizeMessages(
  messages: CoreMessage[],
  maxTokens: number = 4000
): CoreMessage[] {
  const optimized: CoreMessage[] = []
  let totalTokens = 0

  for (let i = messages.length - 1; i >= 0; i--) {
    const message = messages[i]
    const estimatedTokens = estimateTokens(message.content as string)

    if (totalTokens + estimatedTokens > maxTokens) {
      const remainingTokens = maxTokens - totalTokens
      if (remainingTokens > 100) {
        optimized.unshift({
          ...message,
          content: compressPrompt(message.content as string, remainingTokens),
        })
      }
      break
    }

    optimized.unshift(message)
    totalTokens += estimatedTokens
  }

  return optimized
}
```

### 速率限制

实现速率限制可以防止 API 滥用：

```typescript
// src/lib/rate-limit.ts
import { kv } from '@vercel/kv'

export async function rateLimit(identifier: string) {
  const key = `rate-limit:${identifier}`
  const limit = 10 // 每分钟 10 次
  const window = 60 // 60 秒

  const requests = await kv.lrange(key, 0, -1)
  const now = Date.now()
  const validRequests = requests
    .map(r => parseInt(r))
    .filter(t => t > now - window * 1000)

  if (validRequests.length >= limit) {
    throw new Error('请求过于频繁，请稍后再试')
  }

  await kv.lpush(key, now.toString())
  await kv.expire(key, window)
}
```

---

## 九、总结

### 核心要点回顾

#### 1. Next.js 16 的 AI 优势

- **Server Actions** 简化了 AI API 调用
- **流式渲染** 支持实时 AI 响应
- **RSC 优化** 减少客户端 JS
- **边缘运行时** 实现全球低延迟部署

#### 2. Prisma 的数据管理价值

- **类型安全** 的数据库操作
- **向量扩展** 支持语义搜索
- **迁移系统** 简化数据库演进
- **关系管理** 轻松处理关联数据

#### 3. Vercel AI SDK 的核心能力

- **统一接口** 支持多种 LLM 提供商
- **流式处理** 内置 SSE 支持
- **工具调用** 实现 AI Agent
- **生成式 UI** 动态创建界面

### 最佳实践总结

| 状态类型 | 推荐方案 | 理由 |
|---------|---------|------|
| 对话历史 | Prisma + 数据库 | 持久化存储 |
| 用户信息 | NextAuth.js Session | 安全认证 |
| 临时 UI 状态 | React useState | 组件内状态 |
| 缓存数据 | Vercel KV / Redis | 高性能读写 |

### 学习路径建议

```
阶段一：基础（1-2 周）
├── 熟悉 Next.js 16 App Router
├── 掌握 Prisma 基础操作
└── 实现简单聊天机器人

阶段二：进阶（2-3 周）
├── 学习 Vercel AI SDK
├── 实现向量检索和 RAG
└── 添加对话历史管理

阶段三：高级（3-4 周）
├── 构建 AI Agent
├── 实现生成式 UI
└── 性能优化和成本控制

阶段四：实战（4-6 周）
├── 完整项目开发
├── 部署和监控
└── 持续迭代优化
```

---

『参考资料』

- **Vercel AI SDK**: <https://sdk.vercel.ai/docs>
- **Prisma AI Guide**: <https://www.prisma.io/docs/guides/ai>
- **OpenAI API**: <https://platform.openai.com/docs>
- **Pgvector Documentation**: <https://github.com/pgvector/pgvector>
- **Next.js 16 Documentation**: <https://nextjs.org/docs>
- **RAG 论文**: Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks
