import { ParsedMarkdown, ApiResponse, LoginRequest, LoginResponse } from '../types';
import { apiConfig, authConfig } from '../config/api';

export class ApiClient {
    private baseUrl: string;
    private token: string | null = null;
    private tokenExpiry: number | null = null;

    constructor() {
        this.baseUrl = apiConfig.baseUrl;
    }

    /**
     * 登录获取token
     */
    async login(): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}${apiConfig.endpoints.login}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: authConfig.username,
                    password: authConfig.password
                })
            });

            if (!response.ok) {
                console.error(`登录失败: HTTP ${response.status}`);
                return false;
            }

            const data = await response.json() as LoginResponse;
            this.token = data.token;

            // 设置token过期时间（如果API返回了过期时间）
            if (data.expiresIn) {
                this.tokenExpiry = Date.now() + (data.expiresIn * 1000);
            }

            console.log('登录成功，获取到token');
            return true;
        } catch (error) {
            console.error('登录请求失败:', error);
            return false;
        }
    }

    /**
     * 检查token是否有效
     */
    private isTokenValid(): boolean {
        if (!this.token) {
            return false;
        }

        // 如果设置了过期时间，检查是否过期
        if (this.tokenExpiry && Date.now() >= this.tokenExpiry) {
            this.token = null;
            this.tokenExpiry = null;
            return false;
        }

        return true;
    }

    /**
     * 确保已登录（如果需要的话进行登录）
     */
    private async ensureAuthenticated(): Promise<boolean> {
        if (this.isTokenValid()) {
            return true;
        }

        console.log('需要重新登录...');
        return await this.login();
    }

    /**
     * 获取认证头
     */
    private getAuthHeaders(): Record<string, string> {
        const headers: Record<string, string> = {
            'Content-Type': 'application/json'
        };

        if (this.token) {
            headers['Authorization'] = `Bearer ${this.token}`;
        }

        return headers;
    }

    /**
     * 创建文章
     */
    async createPost(post: ParsedMarkdown): Promise<ApiResponse> {
        // 确保已登录
        const isAuthenticated = await this.ensureAuthenticated();
        if (!isAuthenticated) {
            return {
                success: false,
                message: '登录失败，无法创建文章'
            };
        }

        try {
            const response = await fetch(`${this.baseUrl}${apiConfig.endpoints.createPost}`, {
                method: 'POST',
                headers: this.getAuthHeaders(),
                body: JSON.stringify({
                    title: post.title,
                    content: post.content,
                    cover: post.cover,
                    summary: post.summary,
                    categoryID: post.categoryID,
                    postType: post.postType,
                    position: post.position,
                    status: post.status,
                    tags: post.tags
                })
            });

            const data = await response.json() as any;

            if (!response.ok) {
                // 如果是401错误，可能是token过期，尝试重新登录
                if (response.status === 401) {
                    console.log('Token可能已过期，尝试重新登录...');
                    this.token = null;
                    this.tokenExpiry = null;

                    const reAuthenticated = await this.ensureAuthenticated();
                    if (reAuthenticated) {
                        // 重新尝试请求
                        return await this.createPost(post);
                    }
                }

                return {
                    success: false,
                    message: (data as any)?.message || `HTTP ${response.status}: ${response.statusText}`
                };
            }

            return {
                success: true,
                data
            };
        } catch (error) {
            return {
                success: false,
                message: `网络请求失败: ${error}`
            };
        }
    }

    /**
     * 检查API服务是否可用
     */
    async checkHealth(): Promise<boolean> {
        try {
            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 5000);

            const response = await fetch(`${this.baseUrl}${apiConfig.endpoints.health}`, {
                method: 'GET',
                signal: controller.signal
            });

            clearTimeout(timeoutId);
            return response.ok;
        } catch {
            return false;
        }
    }

    /**
     * 批量创建文章
     */
    async batchCreatePosts(posts: ParsedMarkdown[], batchSize: number = 10): Promise<ApiResponse[]> {
        // 确保已登录
        const isAuthenticated = await this.ensureAuthenticated();
        if (!isAuthenticated) {
            return [{
                success: false,
                message: '登录失败，无法批量创建文章'
            }];
        }

        const results: ApiResponse[] = [];

        for (let i = 0; i < posts.length; i += batchSize) {
            const batch = posts.slice(i, i + batchSize);
            console.log(`处理批次 ${Math.floor(i / batchSize) + 1}/${Math.ceil(posts.length / batchSize)}，包含 ${batch.length} 个文件`);

            const batchPromises = batch.map(post => this.createPost(post));
            const batchResults = await Promise.all(batchPromises);
            results.push(...batchResults);

            // 添加延迟避免API限流
            if (i + batchSize < posts.length) {
                console.log('等待1秒后处理下一批次...');
                await new Promise(resolve => setTimeout(resolve, 1000));
            }
        }

        return results;
    }

    /**
     * 手动清除token（用于测试或强制重新登录）
     */
    clearToken(): void {
        this.token = null;
        this.tokenExpiry = null;
    }
}