import { promises as fs } from 'fs';
import * as path from 'path';
import { FileInfo } from '../types';
import { ignorePatterns } from '../config/api';

export class FileScanner {
    private rootDir: string;

    constructor(rootDir: string) {
        this.rootDir = rootDir;
    }

    /**
     * 扫描指定目录下的所有Markdown文件
     */
    async scanMarkdownFiles(): Promise<FileInfo[]> {
        const files: FileInfo[] = [];
        await this.walkDirectory(this.rootDir, files);
        return files;
    }

    /**
     * 递归遍历目录
     */
    private async walkDirectory(dir: string, files: FileInfo[]): Promise<void> {
        try {
            const entries = await fs.readdir(dir, { withFileTypes: true });

            for (const entry of entries) {
                const fullPath = path.join(dir, entry.name);
                const relativePath = path.relative(this.rootDir, fullPath);

                // 跳过忽略的目录和文件
                if (this.shouldIgnore(relativePath)) {
                    continue;
                }

                if (entry.isDirectory()) {
                    await this.walkDirectory(fullPath, files);
                } else if (entry.isFile() && this.isMarkdownFile(entry.name)) {
                    const stats = await fs.stat(fullPath);
                    files.push({
                        filePath: fullPath,
                        relativePath,
                        stats
                    });
                }
            }
        } catch (error) {
            console.error(`扫描目录失败: ${dir}`, error);
        }
    }

    /**
     * 检查是否应该忽略该路径
     */
    private shouldIgnore(relativePath: string): boolean {
        return ignorePatterns.some(pattern => relativePath.includes(pattern));
    }

    /**
     * 检查是否为Markdown文件
     */
    private isMarkdownFile(fileName: string): boolean {
        return path.extname(fileName).toLowerCase() === '.md';
    }
}