import { MarkdownSyncService } from './services/sync-service';
import * as path from 'path';

const rootDir = path.resolve(__dirname, '../..');
const syncService = new MarkdownSyncService(rootDir);

async function main() {
  const args = process.argv.slice(2);
  const command = args[0];

  try {
    switch (command) {
      case 'sync':
        console.log('🚀 开始全量同步...');
        await syncService.sync();
        break;

      case 'batch':
        const batchSize = parseInt(args[1]) || 10;
        console.log(`🚀 开始批量同步，批次大小: ${batchSize}`);
        await syncService.batchSync(batchSize);
        break;

      case 'retry':
        if (args[1]) {
          // 重试指定日志文件
          const logFile = args[1];
          const maxRetries = parseInt(args[2]) || 3;
          console.log(`🔄 重试指定日志文件: ${logFile}`);
          await syncService.retryFailedFiles(logFile, maxRetries);
        } else {
          // 交互式重试
          console.log('🔄 开始交互式重试...');
          await syncService.interactiveRetry();
        }
        break;

      case 'list-logs':
        console.log('📋 列出可重试的日志文件...');
        const logs = await syncService.listRetryableLogs();
        if (logs.length === 0) {
          console.log('✅ 没有找到包含失败记录的日志文件');
        } else {
          logs.forEach((log, index) => {
            const fileName = path.basename(log.file);
            const syncTime = new Date(log.syncTime).toLocaleString('zh-CN');
            console.log(`${index + 1}. ${fileName} - ${log.failureCount}个失败 (${syncTime})`);
          });
        }
        break;

      default:
        console.log('📖 使用说明:');
        console.log('  pnpm dev sync           - 全量同步');
        console.log('  pnpm dev batch [size]   - 批量同步 (默认批次大小: 10)');
        console.log('  pnpm dev retry          - 交互式重试最新失败的文件');
        console.log('  pnpm dev retry <logfile> [maxRetries] - 重试指定日志文件');
        console.log('  pnpm dev list-logs      - 列出可重试的日志文件');
        break;
    }
  } catch (error) {
    console.error('❌ 执行失败:', error);
    process.exit(1);
  }
}

main();