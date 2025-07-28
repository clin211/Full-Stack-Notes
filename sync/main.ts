import * as fs from "fs"

const ignoreDir = [".vscode", ".git", ".idea", ".DS_Store", "plan", "node_modules", "dist", "build", "public", "static", "uploads", "logs", "cache", "temp"]

const isMarkdownFile = (path: string) => path.endsWith(".md")

/**
 * 读取指定路径下的文件和文件夹。
 * @param {string} path - 需要读取的目录路径。
 * @returns {Array<{name: string, path: string, isDir: boolean}>} 文件和文件夹的详细信息数组。
 */
const readFile = (path: string) => {
    const files = fs.readdirSync(path)
    return files.map((file: string) => {
        const filePath = path + "/" + file
        const isDir = fs.statSync(filePath).isDirectory()
        return {
            name: file,
            path: filePath,
            isDir,
        }
    })
}