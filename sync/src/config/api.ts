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
    'algorithm': 5,
    'design-patterns': 5,
    'go': 12,
    'go/基础': 12,
    'go/并发编程': 12,
    'go/Go Web': 13,
    'go/go-zero': 14,
    'go/goframe': 13,
    'javascript': 8,
    'micro-service': 14,
    'Next.js': 9,
    'react': 9,
    'shell': 30,
    'vim': 30,
    'vscode': 30,
    'web-security': 6,
    'frontend': 1,
    'interview': 32
};

// 标签映射
export const tagMapping: Record<string, number[]> = {
    'algorithm': [41],
    'go/并发编程': [12, 48],
    'go/基础': [12],
    'go/Go Web': [12, 13, 14, 20],
    'go/go-zero': [12, 13, 15],
    'go/goframe': [12, 13],
    'micro-service': [12, 19, 44],
    'Next.js': [6, 8],
    'react': [6, 8, 44, 47],
    'shell': [46, 47],
    'vim': [46, 47],
    'vscode': [46, 47],
    'frontend': [1],
    'javascript': [3],
    'projects': [44],
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
    'sync' // 排除技术方案文档
];