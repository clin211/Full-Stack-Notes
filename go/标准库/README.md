# Go 标准库

> 按开发中的**使用频率**由高到低排序，每周学习 2 个包，总计 40 周（约 10 个月）。

## 学习节奏

- **频率**：每周学习 2 个包
- **单包时间**：约 2-3 个工作日
- **产出**：每个包一篇 Markdown 笔记，文件命名格式 `XX-包名.md`

## 第一阶段：每天必用（第 1-10 周）

> 写任何 Go 项目都会用到的包。掌握这 20 个包，日常开发基本无障碍。

### 第 1 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 01 | `fmt` | Print/Scan 系列、格式化动词（%v %T %+v %#v）、宽度精度、Fprint/Sprint |
| 周三-周四 | 02 | `errors` | errors.New、fmt.Errorf（%w）、errors.Is/As/Unwrap、错误包装最佳实践 |
| 周五 | | | 复习：整理 Go 错误处理范式，写一个带自定义错误类型的完整示例 |

### 第 2 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 03 | `strings` | Builder、Contains/Split/Join/Replace/Trim、EqualFold、Reader |
| 周三-周四 | 04 | `os` | 文件操作（Open/Create/Rename/Remove）、环境变量、Exit、Stdin/Stdout/Stderr |
| 周五 | | | 复习：用 fmt + strings + os 写一个日志解析小工具 |

### 第 3 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 05 | `encoding/json` | Marshal/Unmarshal、Encoder/Decoder（流式）、结构体 Tag、omitempty、自定义 JSON |
| 周三-周四 | 06 | `context` | Background/TODO、WithCancel/WithTimeout/WithValue、传播与取消、超时控制模式 |
| 周五 | | | 复习：写一个带上下文超时控制的 JSON API 客户端 |

### 第 4 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 07 | `time` | Time 类型、Duration、Parse/Format（Go 的参考时间 2006-01-02）、Timer/Ticker、Sleep |
| 周三-周四 | 08 | `log/slog` | 结构化日志、Info/Debug/Error、Handler（Text/JSON）、With/WithGroup、Context 传值 |
| 周五 | | | 复习：给之前的工具统一加上 slog 结构化日志和时间统计 |

### 第 5 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 09 | `io` | Reader/Writer 接口、Copy/CopyN/ReadAll、Pipe、MultiWriter/TeeReader |
| 周三-周四 | 10 | `net/http` | Handler/HandlerFunc、ListenAndServe、请求路由、ResponseWriter、ServeMux |
| 周五 | | | 复习：用 io + net/http 实现一个文件上传下载服务 |

### 第 6 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 11 | `strconv` | Atoi/Itoa、ParseFloat/ParseBool、FormatInt/FormatFloat、Append 系列 |
| 周三-周四 | 12 | `sort` | Ints/Strings/Float64s 排序、Slice/SliceStable、自定义排序（Interface）、Search 二分查找 |
| 周五 | | | 复习：写一个数据转换 + 排序的综合练习（如 CSV 数据排序输出） |

### 第 7 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 13 | `bufio` | Scanner/Reader/Writer、按行读取、缓冲写入、带缓冲的读写性能对比 |
| 周三-周四 | 14 | `bytes` | Buffer、Reader、Compare/Equal/Contains、Fields/Split/Join |
| 周五 | | | 复习：对比 bufio vs bytes 在不同场景下的适用性 |

### 第 8 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 15 | `path/filepath` | Join/Base/Dir/Ext、Walk/WalkDir、Glob、Abs/Rel、Clean |
| 周三-周四 | 16 | `sync` | Mutex/RWMutex、WaitGroup、Once、Cond、Map、Pool |
| 周五 | | | 复习：用 filepath.Walk + sync.WaitGroup 实现并发文件搜索 |

### 第 9 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 17 | `testing` | Test/Benchmark/Fuzz 函数、T/B/F 类型、t.Run 子测试、-v/-run/-bench 标志 |
| 周三-周四 | 18 | `net/http/httptest` | NewServer/NewRecorder、模拟 HTTP 请求/响应、测试 Handler |
| 周五 | | | 复习：给第 5 周写的 HTTP 服务补全单元测试 |

### 第 10 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 19 | `net/url` | Parse、Query 编解码、PathEscape、URL 构建 |
| 周三-周四 | 20 | `regexp` | Compile/MustCompile、Match/Find/Replace、子匹配、正则表达式语法要点 |
| 周五 | | | 复习：写一个 URL 解析 + 正则参数提取的路由工具 |

---

## 第二阶段：每周必用（第 11-20 周）

> Web 开发、并发编程中高频使用的包。掌握后能独立完成完整的 Go 项目。

### 第 11 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 21 | `html/template` | 模板语法、变量/管道/条件/循环、自定义函数、模板嵌套、XSS 防护 |
| 周三-周四 | 22 | `net/http`（客户端） | Get/Post/PostForm、Client/Transport、超时配置、请求体处理 |
| 周五 | | | 复习：写一个带重试和超时的 HTTP 客户端封装 |

### 第 12 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 23 | `sync/atomic` | Int32/Int64/Pointer/Bool、Add/Load/Store/Swap/CompareAndSwap、Value |
| 周三-周四 | 24 | `sync/errgroup` | Group、WithContext、SetLimit、并发任务编排与错误收集 |
| 周五 | | | 复习：对比 atomic、mutex、channel 三种并发安全方案的适用场景 |

### 第 13 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 25 | `channel` 实战 | 管道模式、fan-out/fan-in、select 多路复用、关闭与 range（非标准库，但需巩固） |
| 周三-周四 | 26 | `os/signal` | Notify/NotifyContext、捕获 SIGINT/SIGTERM、优雅关闭模式 |
| 周五 | | | 复习：实现一个支持优雅关闭的 HTTP 服务（channel + signal） |

### 第 14 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 27 | `embed` | //go:embed 指令、嵌入字符串/字节/FS、前端静态资源打包 |
| 周三-周四 | 28 | `io/fs` | FS 接口、ReadDir/ReadFile、FileInfo、embed.FS 的关系 |
| 周五 | | | 复习：用 embed + io/fs 把静态网页嵌入 Go 二进制并实现文件服务 |

### 第 15 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 29 | `math/rand` | 随机数生成、Intn/Float64/Shuffle、NewSource 种子控制、Go 1.22+ 自动种子 |
| 周三-周四 | 30 | `math` | 常量（Pi/E/MaxInt 等）、Abs/Min/Max/Ceil/Floor/Sqrt/Pow、浮点数比较 |
| 周五 | | | 复习：实现一个带权重的随机负载均衡器 |

### 第 16 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 31 | `encoding/base64` | StdEncoding/URLEncoding、EncodeToString/DecodeString、流式编解码 |
| 周三-周四 | 32 | `encoding/csv` | Reader/Writer、Comma/LazyQuotes、ReadAll、手动逐行处理 |
| 周五 | | | 复习：实现 CSV 文件导入导出的 HTTP 接口（含 base64 图片字段） |

### 第 17 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 33 | `encoding/xml` | Marshal/Unmarshal、结构体 Tag、XML 与 JSON 的差异 |
| 周三-周四 | 34 | `log` | 基础日志、Printf/Println、Fatal/Panic、SetOutput/Flags、与 slog 的关系 |
| 周五 | | | 复习：对比 Go 日志演进：log → log/slog，整理迁移指南 |

### 第 18 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 35 | `net` | Dial/Listen、TCP/UDP 编程、Resolver、IP 类型 |
| 周三-周四 | 36 | `os/exec` | Command/Run/Start/Output/CombinedOutput、管道、环境变量、工作目录 |
| 周五 | | | 复习：写一个 TCP 端口扫描小工具 |

### 第 19 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 37 | `flag` | String/Int/Bool/Var、Parse、子命令、Usage 自定义 |
| 周三-周四 | 38 | `runtime` | GOMAXPROCS/NumCPU/GC、Caller/Stack、MemStats、调试信息 |
| 周五 | | | 复习：给工具加上命令行参数 + runtime 版本/内存信息输出 |

### 第 20 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 39 | `runtime/pprof` | CPU/内存/Goroutine Profile、pprof 工具使用、go tool pprof 分析 |
| 周三-周四 | 40 | `net/http/pprof` | HTTP Profile 端点、线上性能分析、火焰图 |
| 周五 | | | 复习：给之前的 HTTP 服务加上 pprof 端点，做一次完整的性能分析 |

---

## 第三阶段：按需使用（第 21-30 周）

> 特定场景下经常用到，面试和进阶开发中会涉及。

### 第 21 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 41 | `reflect` | Type/Value、Kind、NumField/Field、NumMethod/Method、Set/Convert |
| 周三-周四 | 42 | `reflect`（实战） | 结构体遍历、通用 JSON/YAML 解析、ORM Tag 读取、注解模式 |
| 周五 | | | 复习：用 reflect 写一个通用的 struct-to-map 转换函数 |

### 第 22 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 43 | `mime` | TypeByExtension、ExtensionsByType、ParseMediaType |
| 周三-周四 | 44 | `net/http/cookiejar` | CookieJar 接口、自动 Cookie 管理、PublicSuffixList |
| 周五 | | | 复习：写一个带 Cookie 自动管理的 HTTP 客户端（模拟登录流程） |

### 第 23 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 45 | `crypto/sha256` + `crypto/md5` | Sum/.New、流式哈希、HMAC、密码哈希与安全建议 |
| 周三-周四 | 46 | `hash` | Hash 接口、Hash32/Hash64、自定义哈希实现 |
| 周五 | | | 复习：实现一个文件完整性校验工具（支持 md5/sha256） |

### 第 24 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 47 | `crypto/rand` | 安全随机数、Int/Prime/Read、与 math/rand 的区别 |
| 周三-周四 | 48 | `crypto/hmac` | HMAC 签名与验证、密钥管理 |
| 周五 | | | 复习：实现一个简单的 API 请求签名验证中间件 |

### 第 25 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 49 | `crypto/aes` | CBC/GCM 模式、加密解密实战、IV/Nonce 管理 |
| 周三-周四 | 50 | `crypto/rsa` + `crypto/ecdsa` | 密钥生成/加密/签名、PEM 编解码 |
| 周五 | | | 复习：实现一个文件加密解密工具（AES-GCM + RSA 密钥交换） |

### 第 26 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 51 | `crypto/tls` | TLS 配置、证书加载、自定义 Server/Client |
| 周三-周四 | 52 | `compress/gzip` | Reader/Writer、压缩与解压、与 HTTP 配合 |
| 周五 | | | 复习：写一个 HTTPS 服务器（自签名证书）+ gzip 压缩中间件 |

### 第 27 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 53 | `compress/zip` | OpenReader/Create、读写 ZIP、文件权限、目录结构 |
| 周三-周四 | 54 | `encoding/binary` | BigEndian/LittleEndian、Read/Write、定长编码、协议解析 |
| 周五 | | | 复习：用 binary 解析一个简单的二进制协议（如 BMP 文件头） |

### 第 28 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 55 | `encoding/gob` | Go 专有序列化、Register、Encode/Decode、与 JSON 的对比 |
| 周三-周四 | 56 | `encoding/hex` | EncodeToString/DecodeString、Dump、流式编解码 |
| 周五 | | | 复习：对比 JSON vs Gob 序列化性能和体积 |

### 第 29 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 57 | `testing/fstest` | MapFS、测试文件系统代码 |
| 周三-周四 | 58 | `testing/synctest`（Go 1.25+） | 协程同步测试、时间感知测试（如果版本支持） |
| 周五 | | | 复习：用 fstest 改造之前文件相关代码的测试 |

### 第 30 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 59 | `runtime/debug` | SetGCPercent/SetTraceback、PrintStack、FreeOSMemory |
| 周三-周四 | 60 | `syscall` | 基本系统调用、信号常量、Errno、与 os 包的关系 |
| 周五 | | | 复习：用 runtime/debug 实现一个运行时自诊断接口 |

---

## 第四阶段：进阶深入（第 31-40 周）

> 深入 Go 底层机制和专业场景，提升对语言本身的理解。

### 第 31 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 61 | `unsafe` | Pointer 转换规则、Sizeof/Offsetof、字符串与字节切片零拷贝转换 |
| 周三-周四 | 62 | `text/template` | 模板语法、管道、自定义函数、模板嵌套（比 html/template 更通用） |
| 周五 | | | 复习：用 text/template 写一个代码生成器 |

### 第 32 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 63 | `unicode` | Is/IsLetter/IsDigit、类别判断、Rune 类型 |
| 周三-周四 | 64 | `unicode/utf8` | EncodeRune/DecodeRune/RuneLen/Valid、中文字符处理 |
| 周五 | | | 复习：写一个支持中文的安全字符串截断函数 |

### 第 33 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 65 | `text/tabwriter` | 对齐输出、制表符宽度计算 |
| 周三-周四 | 66 | `text/scanner` | 词法分析、Scanner 定制、自定义 Token |
| 周五 | | | 复习：写一个终端表格输出工具，支持中文对齐 |

### 第 34 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 67 | `container/heap` | Interface 定义、自定义优先队列、Top-K 问题 |
| 周三-周四 | 68 | `hash/crc32` | ChecksumIEEE/Checksum/Update、校验和计算 |
| 周五 | | | 复习：用 heap 实现一个 Top-K 热点统计 + crc32 校验 |

### 第 35 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 69 | `runtime/trace` | 执行跟踪、Start/Stop、go tool trace 分析 |
| 周三-周四 | 70 | `runtime/metrics` | 运行时指标采集、Read、与 expvar 的关系 |
| 周五 | | | 复习：对之前的并发程序做 trace 分析，找出调度瓶颈 |

### 第 36 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 71 | `expvar` | 暴露运行时变量、Int/Map/Float、自定义指标、/debug/vars |
| 周三-周四 | 72 | `log/slog`（进阶） | 自定义 Handler、日志级别控制、性能优化、结构化日志最佳实践 |
| 周五 | | | 复习：实现一个自定义 slog Handler，支持按级别写入不同文件 |

### 第 37 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 73 | `os`（进阶） | FileMode/FileInfo、User/home 目录、进程信号、文件锁 |
| 周三-周四 | 74 | `debug/buildinfo` | 读取构建信息、Go 版本、依赖版本、-ldflags 注入信息 |
| 周五 | | | 复习：给之前的工具加上版本信息注入和文件锁 |

### 第 38 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 75 | `archive/tar` | Reader/Writer、Tar 文件读写、文件元数据 |
| 周三-周四 | 76 | `net/mail` | 解析邮件地址、邮件头、Address 类型 |
| 周五 | | | 复习：用 tar 写一个简单的项目打包工具 |

### 第 39 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 77 | `net/rpc`（了解） | RPC 基础、Register/Call、与 gRPC 的对比（已不推荐新项目使用） |
| 周三-周四 | 78 | `net/textproto` | MIME 头读写、文本协议解析（了解即可） |
| 周五 | | | 复习：对比 Go 标准库 RPC vs gRPC，整理选型建议 |

### 第 40 周

| 日程 | 序号 | 包名 | 重点内容 |
| --- | --- | --- | --- |
| 周一-周二 | 79 | `plugin` | 插件加载（仅 Linux/macOS）、Open/Lookup、插件模式与局限 |
| 周三-周四 | 80 | `go/build` | 解析 Go 源码包信息、构建约束、Context |
| 周五 | | | 总复习：对 40 周内容做整体回顾，查漏补缺 |

---

## 进度追踪

### 第一阶段：每天必用（第 1-10 周）

- [ ] 01-fmt
- [ ] 02-errors
- [ ] 03-strings
- [ ] 04-os
- [ ] 05-encoding/json
- [ ] 06-context
- [ ] 07-time
- [ ] 08-log/slog
- [ ] 09-io
- [ ] 10-net/http
- [ ] 11-strconv
- [ ] 12-sort
- [ ] 13-bufio
- [ ] 14-bytes
- [ ] 15-path/filepath
- [ ] 16-sync
- [ ] 17-testing
- [ ] 18-net/http/httptest
- [ ] 19-net/url
- [ ] 20-regexp

### 第二阶段：每周必用（第 11-20 周）

- [ ] 21-html/template
- [ ] 22-net/http（客户端）
- [ ] 23-sync/atomic
- [ ] 24-sync/errgroup
- [ ] 25-channel 实战
- [ ] 26-os/signal
- [ ] 27-embed
- [ ] 28-io/fs
- [ ] 29-math/rand
- [ ] 30-math
- [ ] 31-encoding/base64
- [ ] 32-encoding/csv
- [ ] 33-encoding/xml
- [ ] 34-log
- [ ] 35-net
- [ ] 36-os/exec
- [ ] 37-flag
- [ ] 38-runtime
- [ ] 39-runtime/pprof
- [ ] 40-net/http/pprof

### 第三阶段：按需使用（第 21-30 周）

- [ ] 41-reflect
- [ ] 42-reflect（实战）
- [ ] 43-mime
- [ ] 44-net/http/cookiejar
- [ ] 45-crypto/sha256+md5
- [ ] 46-hash
- [ ] 47-crypto/rand
- [ ] 48-crypto/hmac
- [ ] 49-crypto/aes
- [ ] 50-crypto/rsa+ecdsa
- [ ] 51-crypto/tls
- [ ] 52-compress/gzip
- [ ] 53-compress/zip
- [ ] 54-encoding/binary
- [ ] 55-encoding/gob
- [ ] 56-encoding/hex
- [ ] 57-testing/fstest
- [ ] 58-testing/synctest
- [ ] 59-runtime/debug
- [ ] 60-syscall

### 第四阶段：进阶深入（第 31-40 周）

- [ ] 61-unsafe
- [ ] 62-text/template
- [ ] 63-unicode
- [ ] 64-unicode/utf8
- [ ] 65-text/tabwriter
- [ ] 66-text/scanner
- [ ] 67-container/heap
- [ ] 68-hash/crc32
- [ ] 69-runtime/trace
- [ ] 70-runtime/metrics
- [ ] 71-expvar
- [ ] 72-log/slog（进阶）
- [ ] 73-os（进阶）
- [ ] 74-debug/buildinfo
- [ ] 75-archive/tar
- [ ] 76-net/mail
- [ ] 77-net/rpc
- [ ] 78-net/textproto
- [ ] 79-plugin
- [ ] 80-go/build

---

## 参考资源

- [Go 标准库官方文档](https://pkg.go.dev/std)
- [Go 标准库源码](https://github.com/golang/go/tree/master/src)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
