export const apiConfig = {
    baseUrl: 'http://localhost:5555',
    endpoints: {
        login: '/v1/system/auth/login',
        createPost: '/v1/system/posts',
        health: '/health'
    },
    timeout: 10000,
    retryAttempts: 3
};

export const authConfig = {
    username: 'uploadfile',
    password: 'kKOcmxVf99l9OFz'
};

// 分类映射
export const categoryMapping: Record<string, number> = {
    'algorithm': 7,
    'go': 2,
    'go/基础': 35,
    'go/并发编程': 116,
    'go/Go Web/Gin': 36,
    'go/go-zero': 37,
    'go/goframe': 38,
    'javascript': 25,
    'micro-service': 22,
    'Next.js': 30,
    'react': 29,
    'shell': 86,
    'vim': 84,
    'vscode': 84,
    'web-security': 101,
    'frontend': 33,
    'interview': 96
};

// 标签映射
export const tagMapping: Record<string, number[]> = {
    'algorithm/00数据结构与算法': [38, 39],
    'algorithm/01复杂度分析': [38, 40],
    'algorithm/02堆、栈与队列': [38],
    'algorithm/02数组': [38],
    'algorithm/03链表(单链表)': [38],
    'algorithm/04链表(双向链表)': [38],
    'algorithm/05链表(循环链表)': [38],
    'algorithm/06堆栈': [38],
    'algorithm/07队列': [38],
    'frontent/01前端工程化的最佳实践(如何写出高质量的前端代码)': [101],
    'go/并发编程/01初识并发编程': [22, 100],
    'go/并发编程/02Go并发编程的秘密武器：Goroutine深度剖析': [22, 100],
    'go/并发编程/03Go语言并发原语之Mutex深度剖析': [22, 100],
    'go/基础': [22],
    'go/Go Web': [22, 23],
    'go/go-zero': [12, 24],
    'go/goframe': [12, 25],
    'micro-service': [26, 62],
    'Next.js': [2, 1],
    'react': [2],
    'shell': [51],
    'vim': [49],
    'vscode': [48],
    'frontend': [101],
    'javascript': [5],
    'projects/01 深入理解 Server-Sent Events (SSE)：实时数据流的简单解决方案': [96],
    'web-security': [39, 45, 46, 47],
    'interview': [48, 49]
};

export const ignorePatterns = [
    'node_modules',
    '.git',
    'dist',
    'build',
    'pnpm-lock.yaml',
    'package-lock.json',
    '.vscode',
    '.DS_Store',
    'css',
    'design-patterns',
    'musings（随笔）',
    'plan',
    'frontend',
    'projects/02 账号与设备绑定',
    'README',
    'sync' // 排除技术方案文档
];