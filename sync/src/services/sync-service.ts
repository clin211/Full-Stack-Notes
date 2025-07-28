import { FileScanner } from '../scanner/file-scanner';
import { MarkdownParser } from '../parser/markdown-parser';
import { ApiClient } from '../http/api-client';
import { SyncResult, FileInfo } from '../types';
import { LogManager, SyncLogEntry } from '../utils/log-manager';
import * as path from 'path';

export class MarkdownSyncService {
    private fileScanner: FileScanner;
    private markdownParser: MarkdownParser;
    private apiClient: ApiClient;
    private logManager: LogManager;
    private rootDir: string;

    constructor(rootDir: string) {
        this.rootDir = rootDir;
        this.fileScanner = new FileScanner(rootDir);
        this.markdownParser = new MarkdownParser();
        this.apiClient = new ApiClient();
        this.logManager = new LogManager(rootDir);
    }

    /**
     * 过滤空文件
     */
    private async filterEmptyFiles(files: FileInfo[]): Promise<FileInfo[]> {
        const nonEmptyFiles: FileInfo[] = [];
        let emptyFileCount = 0;

        for (const file of files) {
            const isEmpty = await this.markdownParser.isEmptyFile(file);
            if (!isEmpty) {
                nonEmptyFiles.push(file);
            } else {
                emptyFileCount++;
                console.log(`⏭️  跳过空文件: ${file.relativePath}`);
            }
        }

        if (emptyFileCount > 0) {
            console.log(`📝 已跳过 ${emptyFileCount} 个空文件`);
        }

        return nonEmptyFiles;
    }

    /**
     * 执行同步操作
     */
    async sync(): Promise<SyncResult> {
        const result: SyncResult = {
            success: false,
            processedFiles: 0,
            errors: []
        };

        console.log('🚀 开始同步MD文件到服务端...');

        try {
            // 1. 检查API服务是否可用
            console.log('🔍 检查API服务状态...');
            const isHealthy = await this.apiClient.checkHealth();
            if (!isHealthy) {
                console.warn('⚠️  警告: API服务可能不可用，但继续尝试同步');
            }

            // 2. 扫描所有MD文件
            console.log('📂 扫描MD文件...');
            const allFiles = await this.fileScanner.scanMarkdownFiles();
            console.log(`📄 发现 ${allFiles.length} 个MD文件`);

            // 3. 过滤空文件
            const files = await this.filterEmptyFiles(allFiles);
            console.log(`📄 有效文件 ${files.length} 个`);

            // 4. 初始化日志记录
            this.logManager.startSync(files.length);

            // 5. 按创建时间排序文件
            const sortedFiles = this.sortFilesByCreationTime(files);

            // 5. 处理每个文件
            for (let i = 0; i < sortedFiles.length; i++) {
                const file = sortedFiles[i];
                try {
                    const position = i + 1;
                    const parsed = await this.markdownParser.parseFile(file, position);

                    const response = await this.apiClient.createPost(parsed);

                    if (response.success) {
                        console.log(`✅ 同步成功: ${parsed.relativePath}`);
                        this.logManager.logSuccess(parsed.relativePath);
                        result.processedFiles++;
                    } else {
                        const errorMsg = `同步失败: ${response.message}`;
                        console.error(`❌ ${errorMsg}: ${parsed.relativePath}`);
                        this.logManager.logError(parsed.relativePath, errorMsg, 'API_ERROR');
                        result.errors.push(`${parsed.relativePath}: ${errorMsg}`);
                    }
                } catch (error) {
                    const errorMsg = `处理文件失败: ${error}`;
                    console.error(`❌ ${errorMsg}: ${file.relativePath}`);

                    // 根据错误类型分类
                    let errorType: 'PARSE_ERROR' | 'API_ERROR' | 'NETWORK_ERROR' | 'UNKNOWN_ERROR' = 'UNKNOWN_ERROR';
                    if (error instanceof Error) {
                        if (error.message.includes('ENOENT') || error.message.includes('parse')) {
                            errorType = 'PARSE_ERROR';
                        } else if (error.message.includes('fetch') || error.message.includes('network')) {
                            errorType = 'NETWORK_ERROR';
                        } else if (error.message.includes('API') || error.message.includes('HTTP')) {
                            errorType = 'API_ERROR';
                        }
                    }

                    this.logManager.logError(file.relativePath, errorMsg, errorType);
                    result.errors.push(`${file.relativePath}: ${errorMsg}`);
                }

                // 添加延迟避免API请求过于频繁
                if (i < sortedFiles.length - 1) {
                    await this.delay(100);
                }
            }

            result.success = true;

            // 6. 结束日志记录
            const logFile = await this.logManager.endSync();

            // 7. 清理旧日志文件
            await this.logManager.cleanupOldLogs(50);

        } catch (error) {
            const errorMsg = `同步过程中发生错误: ${error}`;
            console.error(errorMsg);
            this.logManager.logError('SYSTEM', errorMsg, 'UNKNOWN_ERROR');
            result.errors.push(errorMsg);

            // 即使出错也要保存日志
            await this.logManager.endSync();
        }

        return result;
    }

    /**
     * 批量同步（分批处理）
     */
    async batchSync(batchSize: number = 10): Promise<SyncResult> {
        const result: SyncResult = {
            success: false,
            processedFiles: 0,
            errors: []
        };

        console.log(`🚀 开始批量同步MD文件，批次大小: ${batchSize}`);

        try {
            const files = await this.fileScanner.scanMarkdownFiles();
            const sortedFiles = this.sortFilesByCreationTime(files);

            // 初始化日志记录
            this.logManager.startSync(files.length);

            // 分批处理
            for (let i = 0; i < sortedFiles.length; i += batchSize) {
                const batch = sortedFiles.slice(i, i + batchSize);
                console.log(`📦 处理批次 ${Math.floor(i / batchSize) + 1}/${Math.ceil(sortedFiles.length / batchSize)}`);

                const batchPromises = batch.map(async (file, index) => {
                    try {
                        const position = i + index + 1;
                        const parsed = await this.markdownParser.parseFile(file, position);
                        const response = await this.apiClient.createPost(parsed);

                        if (response.success) {
                            console.log(`✅ 批次同步成功: ${parsed.relativePath}`);
                            this.logManager.logSuccess(parsed.relativePath);
                            return { success: true };
                        } else {
                            const errorMsg = `批次同步失败: ${response.message}`;
                            console.error(`❌ ${errorMsg}: ${parsed.relativePath}`);
                            this.logManager.logError(parsed.relativePath, errorMsg, 'API_ERROR');
                            return { success: false, error: `${parsed.relativePath}: ${errorMsg}` };
                        }
                    } catch (error) {
                        const errorMsg = `批次处理文件失败: ${error}`;
                        console.error(`❌ ${errorMsg}: ${file.relativePath}`);
                        this.logManager.logError(file.relativePath, errorMsg, 'UNKNOWN_ERROR');
                        return { success: false, error: `${file.relativePath}: ${errorMsg}` };
                    }
                });

                const batchResults = await Promise.all(batchPromises);

                // 统计结果
                batchResults.forEach(res => {
                    if (res.success) {
                        result.processedFiles++;
                    } else if (res.error) {
                        result.errors.push(res.error);
                    }
                });

                // 批次间延迟
                if (i + batchSize < sortedFiles.length) {
                    await this.delay(500);
                }
            }

            result.success = true;

            // 结束日志记录
            await this.logManager.endSync();

            // 清理旧日志文件
            await this.logManager.cleanupOldLogs(50);

        } catch (error) {
            const errorMsg = `批量同步过程中发生错误: ${error}`;
            console.error(errorMsg);
            this.logManager.logError('SYSTEM', errorMsg, 'UNKNOWN_ERROR');
            result.errors.push(errorMsg);

            // 即使出错也要保存日志
            await this.logManager.endSync();
        }

        return result;
    }

    /**
     * 获取最近的同步日志
     */
    async getRecentLogs(limit: number = 10): Promise<string[]> {
        return this.logManager.getRecentLogs(limit);
    }

    /**
     * 按创建时间排序文件
     */
    private sortFilesByCreationTime(files: FileInfo[]): FileInfo[] {
        return files.sort((a, b) => {
            return a.stats.birthtime.getTime() - b.stats.birthtime.getTime();
        });
    }

    /**
     * 延迟函数
     */
    private delay(ms: number): Promise<void> {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    /**
     * 重试失败的文件
     */
    async retryFailedFiles(logFilePath: string, maxRetries: number = 3): Promise<SyncResult> {
        const result: SyncResult = {
            success: false,
            processedFiles: 0,
            errors: []
        };

        console.log(`🔄 开始重试失败的文件，日志文件: ${path.basename(logFilePath)}`);

        try {
            // 1. 读取失败记录
            const allFailedEntries = await this.logManager.readFailedEntriesFromLog(logFilePath);

            // 2. 过滤掉空文件相关的失败记录
            const failedEntries = [];
            let skippedEmptyFiles = 0;

            for (const entry of allFailedEntries) {
                try {
                    const fileInfo = await this.getFileInfo(entry.filePath);
                    // 添加null检查
                    if (!fileInfo) {
                        console.log(`⏭️  跳过不存在的文件: ${entry.filePath}`);
                        continue;
                    }

                    const isEmpty = await this.markdownParser.isEmptyFile(fileInfo);

                    if (!isEmpty) {
                        failedEntries.push(entry);
                    } else {
                        skippedEmptyFiles++;
                        console.log(`⏭️  跳过空文件重试: ${entry.filePath}`);
                    }
                } catch (error) {
                    // 如果文件不存在或无法读取，也跳过
                    console.log(`⏭️  跳过不存在的文件: ${entry.filePath}`);
                }
            }

            if (skippedEmptyFiles > 0) {
                console.log(`📝 已跳过 ${skippedEmptyFiles} 个空文件的重试`);
            }

            if (failedEntries.length === 0) {
                console.log('✅ 没有找到需要重试的有效文件记录');
                result.success = true;
                return result;
            }

            console.log(`📄 发现 ${failedEntries.length} 个需要重试的有效文件`);

            // 3. 开始重试会话
            this.logManager.startRetrySession(logFilePath, failedEntries.length);

            // 3. 检查API服务
            const isHealthy = await this.apiClient.checkHealth();
            if (!isHealthy) {
                console.warn('⚠️  警告: API服务可能不可用，但继续尝试重试');
            }

            // 4. 重试每个失败的文件
            for (const failedEntry of failedEntries) {
                const currentRetryCount = (failedEntry.retryCount || 0) + 1;

                if (currentRetryCount > maxRetries) {
                    console.warn(`⏭️  跳过文件 ${failedEntry.filePath}，已达到最大重试次数 ${maxRetries}`);
                    this.logManager.logError(
                        failedEntry.filePath,
                        `已达到最大重试次数 ${maxRetries}`,
                        'UNKNOWN_ERROR',
                        currentRetryCount
                    );
                    result.errors.push(`${failedEntry.filePath}: 已达到最大重试次数`);
                    continue;
                }

                try {
                    console.log(`🔄 重试文件 ${failedEntry.filePath} (第${currentRetryCount}次重试)`);

                    // 重新扫描文件信息
                    const fileInfo = await this.getFileInfo(failedEntry.filePath);
                    if (!fileInfo) {
                        this.logManager.logError(
                            failedEntry.filePath,
                            '文件不存在或无法访问',
                            'PARSE_ERROR',
                            currentRetryCount
                        );
                        result.errors.push(`${failedEntry.filePath}: 文件不存在`);
                        continue;
                    }

                    // 解析文件
                    const parsed = await this.markdownParser.parseFile(fileInfo);

                    // 发送API请求
                    const response = await this.apiClient.createPost(parsed);

                    if (response.success) {
                        console.log(`✅ 重试成功: ${parsed.relativePath}`);
                        this.logManager.logSuccess(parsed.relativePath);
                        result.processedFiles++;
                    } else {
                        const errorMsg = `重试失败: ${response.message}`;
                        console.error(`❌ ${errorMsg}: ${parsed.relativePath}`);
                        this.logManager.logError(
                            parsed.relativePath,
                            errorMsg,
                            'API_ERROR',
                            currentRetryCount
                        );
                        result.errors.push(`${parsed.relativePath}: ${errorMsg}`);
                    }
                } catch (error) {
                    const errorMsg = `重试处理失败: ${error}`;
                    console.error(`❌ ${errorMsg}: ${failedEntry.filePath}`);

                    let errorType: SyncLogEntry['errorType'] = 'UNKNOWN_ERROR';
                    if (error instanceof Error) {
                        if (error.message.includes('ENOENT') || error.message.includes('parse')) {
                            errorType = 'PARSE_ERROR';
                        } else if (error.message.includes('fetch') || error.message.includes('network')) {
                            errorType = 'NETWORK_ERROR';
                        } else if (error.message.includes('API') || error.message.includes('HTTP')) {
                            errorType = 'API_ERROR';
                        }
                    }

                    this.logManager.logError(
                        failedEntry.filePath,
                        errorMsg,
                        errorType,
                        currentRetryCount
                    );
                    result.errors.push(`${failedEntry.filePath}: ${errorMsg}`);
                }

                // 添加延迟
                await this.delay(200);
            }

            result.success = true;

            // 5. 结束重试会话
            const retryLogFile = await this.logManager.endSync();
            result.logFile = retryLogFile;

        } catch (error) {
            const errorMsg = `重试过程中发生错误: ${error}`;
            console.error(errorMsg);
            result.errors.push(errorMsg);

            // 即使出错也要保存日志
            await this.logManager.endSync();
        }

        return result;
    }

    /**
     * 获取文件信息
     */
    private async getFileInfo(relativePath: string): Promise<FileInfo | null> {
        try {
            const fullPath = path.join(this.rootDir, relativePath);
            const stats = await import('fs').then(fs => fs.promises.stat(fullPath));

            return {
                filePath: fullPath,
                relativePath,
                stats
            };
        } catch (error) {
            return null;
        }
    }

    /**
     * 列出可重试的日志文件
     */
    async listRetryableLogs(): Promise<Array<{ file: string, failureCount: number, syncTime: string }>> {
        return this.logManager.getLogsWithFailures();
    }

    /**
     * 交互式重试选择
     */
    async interactiveRetry(): Promise<SyncResult> {
        console.log('🔍 查找可重试的日志文件...');

        const retryableLogs = await this.listRetryableLogs();

        if (retryableLogs.length === 0) {
            console.log('✅ 没有找到包含失败记录的日志文件');
            return {
                success: true,
                processedFiles: 0,
                errors: []
            };
        }

        console.log('\n📋 可重试的日志文件:');
        retryableLogs.forEach((log, index) => {
            const fileName = path.basename(log.file);
            const syncTime = new Date(log.syncTime).toLocaleString('zh-CN');
            console.log(`${index + 1}. ${fileName} - ${log.failureCount}个失败 (${syncTime})`);
        });

        // 这里可以添加用户输入选择逻辑
        // 为了演示，我们重试最新的一个
        const selectedLog = retryableLogs[0];
        console.log(`\n🎯 自动选择最新的日志文件进行重试: ${path.basename(selectedLog.file)}`);

        return this.retryFailedFiles(selectedLog.file);
    }
}