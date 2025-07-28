import * as fs from 'fs';
import * as path from 'path';
import { ParsedMarkdown, FileInfo } from '../types';
import { categoryMapping, tagMapping } from '../config/api';

export class MarkdownParser {
    /**
     * 检查文件是否为空（内容为空或只包含空白字符）
     */
    async isEmptyFile(fileInfo: FileInfo): Promise<boolean> {
        try {
            const content = await fs.promises.readFile(fileInfo.filePath, 'utf-8');
            const trimmedContent = content.trim();
            return trimmedContent.length === 0;
        } catch (error) {
            console.error(`检查文件是否为空失败: ${fileInfo.filePath}`, error);
            return false; // 如果读取失败，不认为是空文件
        }
    }

    /**
     * 解析markdown文件
     */
    async parseFile(fileInfo: FileInfo, position?: number): Promise<ParsedMarkdown> {
        const content = await fs.promises.readFile(fileInfo.filePath, 'utf-8');

        // 检查内容是否为空
        if (content.trim().length === 0) {
            throw new Error('文件内容为空');
        }

        const title = this.extractTitle(content, fileInfo.filePath);
        const summary = this.generateSummary(content);
        const categoryID = this.getCategoryID(fileInfo.relativePath);
        const tags = this.getTags(fileInfo.relativePath);

        return {
            title,
            content,
            cover: 'https://images.unsplash.com/photo-1555066931-4365d14bab8c?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3',
            summary,
            categoryID,
            postType: 1,
            position: position ? position % 10000 : 0,
            status: 2,
            tags,
            relativePath: fileInfo.relativePath,
            createdAt: fileInfo.stats.birthtime
        };
    }

    /**
     * 提取标题
     */
    private extractTitle(content: string, filePath: string): string {
        // 尝试从内容中提取第一个标题
        const titleMatch = content.match(/^#\s+(.+)$/m);
        if (titleMatch) {
            return titleMatch[1].trim();
        }

        // 如果没有找到标题，使用文件名（去掉扩展名）
        const fileName = path.basename(filePath, '.md');
        return fileName;
    }

    /**
     * 生成摘要（前100个字符）
     */
    private generateSummary(content: string): string {
        // 移除markdown语法，只保留纯文本
        const plainText = content
            .replace(/#{1,6}\s+/g, '') // 移除标题标记
            .replace(/\*\*(.+?)\*\*/g, '$1') // 移除粗体标记
            .replace(/\*(.+?)\*/g, '$1') // 移除斜体标记
            .replace(/\[(.+?)\]\(.+?\)/g, '$1') // 移除链接，保留文本
            .replace(/```[\s\S]*?```/g, '') // 移除代码块
            .replace(/`(.+?)`/g, '$1') // 移除行内代码标记
            .replace(/\n+/g, ' ') // 将换行符替换为空格
            .trim();

        return plainText.length > 100 ? plainText.substring(0, 100) + '...' : plainText;
    }

    /**
     * 根据文件路径获取分类ID
     */
    private getCategoryID(relativePath: string): number {
        const pathParts = relativePath.split('/');

        // 尝试匹配完整路径
        for (const [pattern, categoryID] of Object.entries(categoryMapping)) {
            if (relativePath.startsWith(pattern)) {
                return categoryID;
            }
        }

        // 尝试匹配第一级目录
        const firstDir = pathParts[0];
        if (categoryMapping[firstDir]) {
            return categoryMapping[firstDir];
        }

        // 默认分类
        return 1;
    }

    /**
     * 根据文件路径获取标签
     */
    private getTags(relativePath: string): number[] {
        // 尝试匹配完整路径
        for (const [pattern, tags] of Object.entries(tagMapping)) {
            if (relativePath.startsWith(pattern)) {
                return tags;
            }
        }

        // 尝试匹配第一级目录
        const pathParts = relativePath.split('/');
        const firstDir = pathParts[0];
        if (tagMapping[firstDir]) {
            return tagMapping[firstDir];
        }

        // 默认标签
        return [];
    }

    /**
     * 批量解析文件
     */
    async parseFiles(fileInfos: FileInfo[]): Promise<ParsedMarkdown[]> {
        const results: ParsedMarkdown[] = [];

        for (const fileInfo of fileInfos) {
            try {
                const parsed = await this.parseFile(fileInfo);
                results.push(parsed);
            } catch (error) {
                console.error(`解析文件失败: ${fileInfo.filePath}`, error);
            }
        }

        // 按position排序
        return results.sort((a, b) => a.position - b.position);
    }
}