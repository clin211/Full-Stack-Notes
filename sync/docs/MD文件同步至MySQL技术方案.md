# MD文件同步至MySQL技术方案

## 一、项目概述

本方案旨在将Full-Stack-Notes项目中的所有Markdown文件内容同步到MySQL数据库的post表中，实现文档内容的结构化存储和管理。

## 二、技术栈

- **运行环境**: Node.js + TypeScript
- **数据库**: MySQL 8.x
- **数据库驱动**: mysql2
- **文件处理**: fs/promises + path
- **包管理**: pnpm

## 三、数据库配置

```yaml
mysql:
  addr: 127.0.0.1:11006
  username: root
  password: 123456
  database: miniblog_v2
```

## 四、目标表结构

使用现有的post表，关键字段映射：
- `post_id`: 基于文件路径生成的唯一ID
- `title`: 从MD文件名中提取
- `content`: MD文件的完整内容
- `user_id`: 固定值（如'system'）
- `category_id`: 基于文件夹结构映射
- `created_at`: 文件创建时间
- `updated_at`: 文件最后修改时间

category_id 的映射关系：
- algorithm -> 5
- design-patterns -> 5
- go -> 12
- go/基础 -> 12
- go/并发编程 -> 12
- go/Go Web -> 13
- go/go-zero -> 14
- go/goframe -> 13
- javascript -> 8
- micro-service -> 14
- Next.js -> 9
- projects/01 深入理解 Server-Sent Events (SSE)：实时数据流的简单解决方案.md -> 7
- react -> 9
- shell -> 30
- vim -> 30
- vscode -> 30
- web-security -> 6

## 五、核心功能模块

### 5.1 文件扫描模块

```typescript
import { promises as fs } from 'fs';
import * as path from 'path';

interface FileInfo {
  filePath: string;
  relativePath: string;
  stats: fs.Stats;
}

async function scanMarkdownFiles(rootDir: string): Promise<FileInfo[]> {
  const files: FileInfo[] = [];
  const ignorePatterns = [
    'node_modules',
    '.git',
    'dist',
    'build',
    'pnpm-lock.yaml',
    'package-lock.json'
  ];

  async function walkDir(dir: string): Promise<void> {
    const entries = await fs.readdir(dir, { withFileTypes: true });
    
    for (const entry of entries) {
      const fullPath = path.join(dir, entry.name);
      const relativePath = path.relative(rootDir, fullPath);
      
      // 跳过忽略的目录和文件
      if (ignorePatterns.some(pattern => relativePath.includes(pattern))) {
        continue;
      }
      
      if (entry.isDirectory()) {
        await walkDir(fullPath);
      } else if (entry.isFile() && path.extname(entry.name) === '.md') {
        const stats = await fs.stat(fullPath);
        files.push({
          filePath: fullPath,
          relativePath,
          stats
        });
      }
    }
  }
  
  await walkDir(rootDir);
  return files;
}
```

### 5.2 MD文件解析模块

```typescript
import { createHash } from 'crypto';

interface ParsedMarkdown {
  postId: string;
  title: string;
  content: string;
  summary: string;
  categoryId: number;
  relativePath: string;
  createdAt: Date;
  updatedAt: Date;
}

function generatePostId(relativePath: string): string {
  return createHash('md5').update(relativePath).digest('hex').substring(0, 32);
}

function extractTitle(content: string, fileName: string): string {
  // 尝试从内容中提取第一个标题
  const titleMatch = content.match(/^#\s+(.+)$/m);
  if (titleMatch) {
    return titleMatch[1].trim();
  }
  
  // 如果没有找到标题，使用文件名
  return path.basename(fileName, '.md');
}

function generateSummary(content: string): string {
  // 移除markdown语法，提取前500个字符作为摘要
  const plainText = content
    .replace(/#{1,6}\s+/g, '') // 移除标题标记
    .replace(/\*\*(.+?)\*\*/g, '$1') // 移除粗体
    .replace(/\*(.+?)\*/g, '$1') // 移除斜体
    .replace(/\[(.+?)\]\(.+?\)/g, '$1') // 移除链接
    .replace(/```[\s\S]*?```/g, '') // 移除代码块
    .replace(/`(.+?)`/g, '$1') // 移除行内代码
    .trim();
  
  return plainText.substring(0, 500);
}

function getCategoryId(relativePath: string): number {
  // 基于文件夹结构映射分类ID
  const categoryMap: Record<string, number> = {
    'Next.js': 1,
    'algorithm': 2,
    'css': 3,
    'design-patterns': 4,
    'frontend': 5,
    'go': 6,
    'javascript': 7,
    'react': 8,
    'shell': 9,
    'vim': 10,
    'web-security': 11,
    'micro-service': 12,
    'interview': 13,
    'projects': 14,
    'musings': 15
  };
  
  const firstDir = relativePath.split('/')[0];
  return categoryMap[firstDir] || 0;
}

async function parseMarkdownFile(fileInfo: FileInfo): Promise<ParsedMarkdown> {
  const content = await fs.readFile(fileInfo.filePath, 'utf-8');
  const title = extractTitle(content, fileInfo.filePath);
  const summary = generateSummary(content);
  const postId = generatePostId(fileInfo.relativePath);
  const categoryId = getCategoryId(fileInfo.relativePath);
  
  return {
    postId,
    title,
    content,
    summary,
    categoryId,
    relativePath: fileInfo.relativePath,
    createdAt: fileInfo.stats.birthtime,
    updatedAt: fileInfo.stats.mtime
  };
}
```

### 5.3 数据库操作模块

```typescript
import mysql from 'mysql2/promise';

class DatabaseManager {
  private connection: mysql.Connection;
  
  constructor() {
    this.connection = null;
  }
  
  async connect(): Promise<void> {
    this.connection = await mysql.createConnection({
      host: '127.0.0.1',
      port: 11006,
      user: 'root',
      password: '123456',
      database: 'miniblog_v2',
      charset: 'utf8mb4'
    });
  }
  
  async disconnect(): Promise<void> {
    if (this.connection) {
      await this.connection.end();
    }
  }
  
  async upsertPost(post: ParsedMarkdown): Promise<void> {
    const sql = `
      INSERT INTO post (
        post_id, title, content, summary, user_id, category_id,
        post_type, status, created_at, updated_at
      ) VALUES (?, ?, ?, ?, 'system', ?, 1, 2, ?, ?)
      ON DUPLICATE KEY UPDATE
        title = VALUES(title),
        content = VALUES(content),
        summary = VALUES(summary),
        category_id = VALUES(category_id),
        updated_at = VALUES(updated_at)
    `;
    
    await this.connection.execute(sql, [
      post.postId,
      post.title,
      post.content,
      post.summary,
      post.categoryId,
      post.createdAt,
      post.updatedAt
    ]);
  }
  
  async getExistingPosts(): Promise<Set<string>> {
    const [rows] = await this.connection.execute(
      'SELECT post_id FROM post WHERE user_id = "system"'
    );
    
    return new Set((rows as any[]).map(row => row.post_id));
  }
  
  async deletePost(postId: string): Promise<void> {
    await this.connection.execute(
      'UPDATE post SET deleted_at = NOW() WHERE post_id = ?',
      [postId]
    );
  }
}
```

### 5.4 主程序模块

```typescript
import * as path from 'path';

class MarkdownSyncService {
  private db: DatabaseManager;
  private rootDir: string;
  
  constructor(rootDir: string) {
    this.db = new DatabaseManager();
    this.rootDir = rootDir;
  }
  
  async sync(): Promise<void> {
    console.log('开始同步MD文件到数据库...');
    
    try {
      await this.db.connect();
      
      // 1. 扫描所有MD文件
      console.log('扫描MD文件...');
      const files = await scanMarkdownFiles(this.rootDir);
      console.log(`发现 ${files.length} 个MD文件`);
      
      // 2. 获取数据库中已存在的文章
      const existingPosts = await this.db.getExistingPosts();
      const currentPostIds = new Set<string>();
      
      // 3. 处理每个文件
      for (const file of files) {
        try {
          const parsed = await parseMarkdownFile(file);
          currentPostIds.add(parsed.postId);
          
          await this.db.upsertPost(parsed);
          console.log(`✓ 同步: ${parsed.relativePath}`);
        } catch (error) {
          console.error(`✗ 处理文件失败: ${file.relativePath}`, error);
        }
      }
      
      // 4. 删除不再存在的文件对应的记录
      for (const existingId of existingPosts) {
        if (!currentPostIds.has(existingId)) {
          await this.db.deletePost(existingId);
          console.log(`✓ 删除不存在的文件记录: ${existingId}`);
        }
      }
      
      console.log('同步完成!');
    } catch (error) {
      console.error('同步过程中发生错误:', error);
    } finally {
      await this.db.disconnect();
    }
  }
}

// 使用示例
async function main() {
  const projectRoot = '/Users/changlin/code/front-end/Full-Stack-Notes';
  const syncService = new MarkdownSyncService(projectRoot);
  await syncService.sync();
}

if (require.main === module) {
  main().catch(console.error);
}
```

## 六、项目结构
```tree
sync/
├── src/
│   ├── scanner.ts          # 文件扫描模块
│   ├── parser.ts           # MD解析模块
│   ├── database.ts         # 数据库操作模块
│   ├── sync-service.ts     # 主服务类
│   └── index.ts           # 入口文件
├── docs/
│   └── MD文件同步至MySQL技术方案.md
├── package.json
├── tsconfig.json
└── README.md
```

## 七、依赖包配置

```json
{
  "name": "markdown-sync",
  "version": "1.0.0",
  "scripts": {
    "build": "tsc",
    "start": "node dist/index.js",
    "dev": "ts-node src/index.ts",
    "sync": "npm run build && npm start"
  },
  "dependencies": {
    "mysql2": "^3.6.0"
  },
  "devDependencies": {
    "@types/node": "^20.0.0",
    "typescript": "^5.0.0",
    "ts-node": "^10.9.0"
  }
}
```

## 八、使用方法

1. 安装依赖：`pnpm install`
2. 编译项目：`pnpm run build`
3. 执行同步：`pnpm run sync`

## 九、注意事项

1. **数据安全**: 建议在生产环境中使用数据库事务
2. **性能优化**: 大量文件时可考虑批量插入
3. **错误处理**: 记录详细的错误日志便于排查
4. **增量同步**: 可基于文件修改时间实现增量同步
5. **分类映射**: 根据实际需要调整文件夹到分类ID的映射关系

## 十、扩展功能

- 支持定时同步（使用cron表达式）
- 支持文件变更监听（使用chokidar）
- 支持多种文档格式（rst、txt等）
- 支持内容全文检索索引
- 支持同步状态Web界面