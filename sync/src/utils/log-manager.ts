import * as fs from 'fs';
import * as path from 'path';

export interface SyncLogEntry {
    timestamp: string;
    filePath: string;
    error: string;
    errorType: 'PARSE_ERROR' | 'API_ERROR' | 'NETWORK_ERROR' | 'UNKNOWN_ERROR';
    retryCount?: number; // 新增重试次数
    lastRetryTime?: string; // 新增最后重试时间
}

export interface SyncLogSummary {
    syncStartTime: string;
    syncEndTime: string;
    totalFiles: number;
    successCount: number;
    failureCount: number;
    errors: SyncLogEntry[];
    isRetrySession?: boolean; // 新增是否为重试会话
    originalLogFile?: string; // 新增原始日志文件路径
}

export class LogManager {
    private logsDir: string;
    private currentLogFile!: string;
    private syncLog!: SyncLogSummary;

    constructor(rootDir: string) {
        this.logsDir = path.join(rootDir, 'sync', 'logs');
        this.ensureLogsDirectory();
        this.initializeSyncLog();
    }

    /**
     * 确保日志目录存在
     */
    private ensureLogsDirectory(): void {
        if (!fs.existsSync(this.logsDir)) {
            fs.mkdirSync(this.logsDir, { recursive: true });
        }
    }

    /**
     * 初始化同步日志
     */
    private initializeSyncLog(isRetry: boolean = false, originalLogFile?: string): void {
        const timestamp = new Date().toISOString().replace(/[:.]/g, '-');
        const prefix = isRetry ? 'retry' : 'sync';
        this.currentLogFile = path.join(this.logsDir, `${prefix}-${timestamp}.json`);

        this.syncLog = {
            syncStartTime: new Date().toISOString(),
            syncEndTime: '',
            totalFiles: 0,
            successCount: 0,
            failureCount: 0,
            errors: [],
            isRetrySession: isRetry,
            originalLogFile
        };
    }

    /**
     * 开始重试会话
     */
    startRetrySession(originalLogFile: string, totalFiles: number): void {
        this.initializeSyncLog(true, originalLogFile);
        this.syncLog.totalFiles = totalFiles;
        this.syncLog.syncStartTime = new Date().toISOString();
        console.log(`🔄 开始重试会话，原始日志: ${path.basename(originalLogFile)}`);
        console.log(`📝 重试日志文件: ${path.basename(this.currentLogFile)}`);
    }

    /**
     * 开始同步记录
     */
    startSync(totalFiles: number): void {
        this.syncLog.totalFiles = totalFiles;
        this.syncLog.syncStartTime = new Date().toISOString();
        console.log(`📝 开始同步记录，日志文件: ${path.basename(this.currentLogFile)}`);
    }

    /**
     * 记录成功
     */
    logSuccess(filePath: string): void {
        this.syncLog.successCount++;
    }

    /**
     * 记录错误
     */
    logError(filePath: string, error: string, errorType: SyncLogEntry['errorType'] = 'UNKNOWN_ERROR', retryCount: number = 0): void {
        const logEntry: SyncLogEntry = {
            timestamp: new Date().toISOString(),
            filePath,
            error,
            errorType,
            retryCount,
            lastRetryTime: retryCount > 0 ? new Date().toISOString() : undefined
        };

        this.syncLog.errors.push(logEntry);
        this.syncLog.failureCount++;

        const retryInfo = retryCount > 0 ? ` (重试第${retryCount}次)` : '';
        console.error(`❌ [${errorType}] ${filePath}: ${error}${retryInfo}`);
    }

    /**
     * 结束同步并保存日志
     */
    async endSync(): Promise<string> {
        this.syncLog.syncEndTime = new Date().toISOString();

        try {
            await fs.promises.writeFile(
                this.currentLogFile,
                JSON.stringify(this.syncLog, null, 2),
                'utf-8'
            );

            const summary = this.generateSummary();
            console.log(summary);

            return this.currentLogFile;
        } catch (error) {
            console.error('保存日志文件失败:', error);
            throw error;
        }
    }

    /**
     * 生成同步摘要
     */
    private generateSummary(): string {
        const duration = new Date(this.syncLog.syncEndTime).getTime() -
            new Date(this.syncLog.syncStartTime).getTime();
        const durationStr = `${Math.round(duration / 1000)}秒`;

        const sessionType = this.syncLog.isRetrySession ? '重试' : '同步';
        let summary = `\n📊 ${sessionType}完成摘要:\n`;
        summary += `⏱️  耗时: ${durationStr}\n`;
        summary += `📁 总文件数: ${this.syncLog.totalFiles}\n`;
        summary += `✅ 成功: ${this.syncLog.successCount}\n`;
        summary += `❌ 失败: ${this.syncLog.failureCount}\n`;
        summary += `📝 日志文件: ${this.currentLogFile}\n`;

        if (this.syncLog.isRetrySession && this.syncLog.originalLogFile) {
            summary += `🔗 原始日志: ${this.syncLog.originalLogFile}\n`;
        }

        if (this.syncLog.failureCount > 0) {
            summary += `\n❌ 失败详情:\n`;
            this.syncLog.errors.forEach((error, index) => {
                const retryInfo = error.retryCount ? ` (已重试${error.retryCount}次)` : '';
                summary += `${index + 1}. [${error.errorType}] ${error.filePath}${retryInfo}\n`;
                summary += `   错误: ${error.error}\n`;
            });

            if (!this.syncLog.isRetrySession) {
                summary += `\n💡 提示: 可以使用重试功能重新同步失败的文件\n`;
            }
        }

        return summary;
    }

    /**
     * 从日志文件读取失败记录
     */
    async readFailedEntriesFromLog(logFilePath: string): Promise<SyncLogEntry[]> {
        try {
            const logContent = await fs.promises.readFile(logFilePath, 'utf-8');
            const logData: SyncLogSummary = JSON.parse(logContent);
            return logData.errors || [];
        } catch (error) {
            console.error(`读取日志文件失败: ${logFilePath}`, error);
            return [];
        }
    }

    /**
     * 获取最近的日志文件列表
     */
    async getRecentLogs(limit: number = 10): Promise<string[]> {
        try {
            const files = await fs.promises.readdir(this.logsDir);
            const logFiles = files
                .filter(file => (file.startsWith('sync-') || file.startsWith('retry-')) && file.endsWith('.json'))
                .sort((a, b) => b.localeCompare(a)) // 按时间倒序
                .slice(0, limit);

            return logFiles.map(file => path.join(this.logsDir, file));
        } catch (error) {
            console.error('读取日志目录失败:', error);
            return [];
        }
    }

    /**
     * 获取有失败记录的日志文件
     */
    async getLogsWithFailures(limit: number = 20): Promise<Array<{ file: string, failureCount: number, syncTime: string }>> {
        try {
            const logFiles = await this.getRecentLogs(limit);
            const logsWithFailures = [];

            for (const logFile of logFiles) {
                try {
                    const logContent = await fs.promises.readFile(logFile, 'utf-8');
                    const logData: SyncLogSummary = JSON.parse(logContent);

                    if (logData.failureCount > 0) {
                        logsWithFailures.push({
                            file: logFile,
                            failureCount: logData.failureCount,
                            syncTime: logData.syncStartTime
                        });
                    }
                } catch (error) {
                    console.warn(`跳过无效日志文件: ${logFile}`);
                }
            }

            return logsWithFailures.sort((a, b) => b.syncTime.localeCompare(a.syncTime));
        } catch (error) {
            console.error('获取失败日志列表失败:', error);
            return [];
        }
    }

    /**
     * 清理旧日志文件（保留最近N个）
     */
    async cleanupOldLogs(keepCount: number = 50): Promise<void> {
        try {
            const files = await fs.promises.readdir(this.logsDir);
            const logFiles = files
                .filter(file => (file.startsWith('sync-') || file.startsWith('retry-')) && file.endsWith('.json'))
                .sort((a, b) => b.localeCompare(a));

            if (logFiles.length > keepCount) {
                const filesToDelete = logFiles.slice(keepCount);
                for (const file of filesToDelete) {
                    await fs.promises.unlink(path.join(this.logsDir, file));
                }
                console.log(`🗑️  清理了 ${filesToDelete.length} 个旧日志文件`);
            }
        } catch (error) {
            console.error('清理日志文件失败:', error);
        }
    }
}