export interface FileInfo {
    filePath: string;
    relativePath: string;
    stats: import('fs').Stats;
}

export interface ParsedMarkdown {
    title: string;
    content: string;
    cover: string;
    summary: string;
    categoryID: number;
    postType: number;
    position: number;
    status: number;
    tags: number[];
    relativePath: string;
    createdAt: Date;
}

export interface ApiResponse {
    success: boolean;
    message?: string;
    data?: any;
}

export interface SyncResult {
    success: boolean;
    processedFiles: number;
    errors: string[];
    logFile?: string; // 新增日志文件路径
}

// 添加登录相关类型
export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    expiresIn?: number;
}

export interface AuthConfig {
    username: string;
    password: string;
}