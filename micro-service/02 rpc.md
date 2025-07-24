RPC（Remote Procedure Call，远程过程调用）是一种允许程序调用另一台计算机上的函数或方法的技术，其核心思想是：**让开发者能够像调用本地函数一样调用远程服务**，从而隐藏底层网络通信的复杂性。

## RPC 的背景

随着计算机网络的发展，传统的单体应用架构逐渐向分布式系统演进。不同的服务或组件可能运行在不同的服务器、数据中心甚至云端，如何让它们高效、透明地进行通信成为一个重要问题。RPC（远程过程调用）正是为了解决这个问题而诞生的。

在早期，进程间通信（IPC，Inter-Process Communication）主要依赖于本地机制，如共享内存、管道（pipe）、消息队列等，但这些方法仅限于单机环境。当系统扩展到多台机器时，就需要通过网络通信来完成数据交互，而 RPC 让远程调用像本地调用一样透明，使开发者无需关注底层网络细节。

## RPC 的特点

- **透明性**：调用远程方法的方式与本地方法几乎一致，降低开发成本
- **跨语言支持**：支持不同语言间的调用（如 Go 调用 Java 服务）
- **高效通信**：可以使用二进制协议（如 Protobuf）提高传输效率
- **自动生成代码**：如 gRPC 可以通过 `.proto` 文件生成客户端和服务器端代码

在分布式系统中，RPC 的价值体现在以下方面：

- **简化开发**：开发者无需关注网络协议、序列化、错误重试等底层细节，只需定义接口并调用，即可实现跨进程通信。

- **提升系统模块化**：通过接口定义（如 Protocol Buffers 或 Thrift IDL），服务间的依赖关系清晰可见，便于维护和迭代。

- **性能优势**：相比基于 HTTP 的 REST API，RPC 通常采用二进制协议（如 gRPC 的 Protocol Buffers），传输效率更高，延迟更低。

- **微服务的基石**：在微服务架构中，服务拆分后必然面临通信问题。RPC 通过标准化、高性能的通信机制，成为微服务落地的首选方案。

## 如何进行远程过程调用？

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/85ClAl-image-20250302225340780.png)

> 上图源自 Bruce Jay Nelson 于 1984 年发表的经典论文《[Implementing Remote Procedure Calls](https://web.eecs.umich.edu/~mosharaf/Readings/RPC.pdf)》，该论文首次系统化定义了远程过程调用的核心设计范式，为现代分布式系统通信奠定了理论基础。尽管当前主流的 RPC 框架（如 gRPC、Dubbo 等）在协议实现、序列化方式和传输层优化等方面存在差异，但其核心设计思想——**通过抽象网络通信为本地方法调用，降低分布式系统开发复杂性**——仍与 Nelson 提出的原始模型一脉相承。

由上面 RPC 的架构图我们可以看出主要分为三大部分：

- 调用方机器（Caller Machine）
- 网络（Network）
- 被调用方机器（Callee Machine）

在进行远程过程调用时，涉及到五个角色：

- 用户(User)：发起调用的客户端代码。
- 用户存根(User-stud)：客户端代理，负责将本地调用转换为网络请求。核心作用是参数打包（序列化）、生成调用包（Call Packet）。
- RPC通信包(RPCRuntime)：管理底层网络通信，如连接池、超时重试、管理网络监听和响应返回。
- 服务器存根(Server-stub)：服务端骨架，接收请求并路由到实际服务。核心作用是请求解包（反序列化）、调用服务端方法。
- 服务器(Server)：实际执行业务逻辑的代码（如查询数据库）。

调用流程分步详解

1. 用户发起调用（User → User-stub）
   - 用户调用本地接口，实际由 User-stub 接管。
   - 关键设计：用户感知不到远程调用，接口与本地代码一致。
2. 参数打包与生成调用包（User-stub → RPC Runtime）
   - 打包（Pack）：将方法名、参数序列化为 Call Packet（如 Protobuf 编码）。
   - 依赖接口（Interface）：通过预定义的接口描述（如 IDL）确保数据格式一致。
3. 网络传输（RPC Runtime → Network）
   - RPC 运行时将 Call Packet 发送到网络层，准备传输至被调用方机器。

4. 跨机器传输（Network → Callee Machine）

   - 调用包通过底层协议（如 TCP）传输到目标机器。

   - 服务发现：依赖注册中心或 DNS 定位目标服务地址（图中未显式标注）。

5. 接收与解包（Server-stub）

   - 服务端 RPC 运行时接收 Call Packet，交给 Server-stub 解包（反序列化）。

   - 还原请求：提取方法名、参数，准备调用服务端代码。

6. 调用服务端方法（Server-stub → Server）

   - 反射或路由：根据方法名调用对应的服务实现。

   - 业务执行（Work）：服务端执行业务逻辑（如查询数据库）。

7. 生成结果包（Server-stub → RPC Runtime）
   - 打包（Pack）：将结果或异常序列化为 Result Packet。

8. 结果回传（Network → Caller Machine）
   - Result Packet 通过网络返回调用方机器。

9. 解包与返回用户（User-stub → User）

   - 调用方 RPC 运行时接收结果，User-stub 解包后返回给用户。

   - 用户感知：结果如同本地方法返回，全程无感知远程交互。

## RPC 与 REST API 的对比

### **🔹 RPC（远程过程调用）**

- 允许调用远程服务器上的方法，就像调用本地方法一样。
- 典型实现：**gRPC、JSON-RPC、Thrift、Dubbo**。
- 使用二进制协议（如 **Protobuf**）或文本协议（如 **JSON**）。
- **强调方法调用和执行**。

### **🔹 REST API（基于 HTTP）**

- 通过 **HTTP 请求** 在客户端和服务器之间传输数据。
- 以 **资源（Resource）** 为核心，使用 **CRUD（增删改查）** 操作。
- 典型实现：**基于 JSON 的 RESTful API**。
- **强调资源的访问和操作**。

| 维度               | RPC                                    | REST API                                             |
| ------------------ | -------------------------------------- | ---------------------------------------------------- |
| **通信协议**       | TCP、HTTP/2（gRPC），HTTP（JSON-RPC）  | 仅支持 HTTP                                          |
| **数据格式**       | Protobuf（gRPC）、JSON（JSON-RPC）     | JSON / XML                                           |
| **性能**           | 更快，二进制协议占用带宽小             | 速度较慢，HTTP 头部冗余较大                          |
| **方法调用方式**   | 直接调用远程方法，如 `SayHello(name)`  | 通过 **HTTP 动词**（GET、POST、PUT、DELETE）访问资源 |
| **客户端代码生成** | gRPC 支持自动生成客户端 SDK            | 需手写 API 请求代码                                  |
| **流式数据支持**   | gRPC 支持双向流式通信                  | 需要 WebSocket 实现流式数据                          |
| **可读性**         | Protobuf 可读性较差，JSON-RPC 可读性好 | 采用 JSON，可读性强                                  |
| **安全性**         | 需要额外实现身份认证（TLS、OAuth）     | 依赖 HTTP 认证（JWT、OAuth）                         |
| **适用场景**       | 内部微服务通信、高性能场景             | Web 应用、RESTful 架构                               |

### 🔹传输协议

- **RPC**（如 gRPC）支持 **HTTP/2**，默认使用 **TCP 连接**，可实现 **长连接、流式传输**，性能更优。
- **REST API** 依赖 **HTTP/1.1** 或 **HTTP/2**，每次请求都需要**携带完整的 HTTP 头部信息**，开销较大。

### 🔹数据格式

- **RPC（gRPC）** 主要使用 **Protocol Buffers（Protobuf）**，比 JSON **更紧凑、解析速度更快**。
- **REST API** 主要使用 **JSON**，可读性强但**数据冗余较大**。

 **示例：相同的数据，Protobuf 和 JSON 的对比**

```proto
message User {
  string name = 1;
  int32 age = 2;
}
```

 对应 JSON 版本：

```json
{
  "name": "Alice",
  "age": 25
}
```

Protobuf 序列化后的数据更紧凑，占用 **更少带宽**，适合大规模高并发场景。

### 🔹性能

- **RPC**（gRPC）由于使用 **二进制格式（Protobuf）+ HTTP/2**，比基于 **JSON + HTTP/1.1** 的 REST API **快 5-10 倍**。
- **REST API** 需要携带额外的 HTTP 头部信息（如 `Content-Type`），数据量更大，解析速度较慢。

**性能对比（百万次请求）：**

| 方案                        | 吞吐量（requests/sec） | 数据大小 |
| --------------------------- | ---------------------- | -------- |
| **gRPC（Protobuf）**        | 100,000+               | 小       |
| **JSON-RPC（TCP）**         | 80,000+                | 小       |
| **REST API（JSON + HTTP）** | 10,000-50,000          | 大       |

### 🔹客户端 SDK

- **RPC（gRPC）**：支持 **自动生成客户端 SDK**，调用更方便。
- **REST API**：通常需要手写 API 请求代码，开发效率较低。

#### **gRPC 客户端调用**

```go
client.SayHello(context.Background(), &pb.HelloRequest{Name: "Alice"})
```

#### **REST API 客户端调用**

```javascript
fetch("https://api.example.com/hello", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name: "Alice" })
})
.then(response => response.json())
.then(data => console.log(data.message));
```

REST API 需要**手动管理请求 URL、Headers、Body**，而 gRPC 直接调用方法，更简洁。

### 🔹流式数据

- **gRPC** 支持 **双向流式通信（Server Streaming、Client Streaming、Bidirectional Streaming）**，适用于**实时聊天、视频推流**等场景。
- **REST API** 需要 **WebSocket** 额外实现流式通信，复杂度更高。

### 🔹适用场景

| 适用场景                | 选择 RPC                 | 选择 REST API         |
| ----------------------- | ------------------------ | --------------------- |
| **微服务架构**          | ✅ gRPC 更高效            | ❌ REST API 有性能瓶颈 |
| **Web 应用**            | ❌ 需要额外实现 REST 兼容 | ✅ REST API 直接可用   |
| **移动端 / 低带宽环境** | ✅ Protobuf 省流量        | ❌ JSON 传输数据大     |
| **实时数据流**          | ✅ gRPC Streaming         | ❌ 需要 WebSocket/SSE  |

### **🔹何时选择 RPC？何时选择 REST API？**

#### **✅ 选择 RPC（如 gRPC）**

- 需要高性能、低延迟的通信（如 **微服务**）。
- 需要**双向流式通信**（如 **聊天应用、视频推流**）。
- 需要跨语言通信，并且希望 **自动生成 SDK**。
- 适用于 **企业级内部服务**，而非公开 API。

#### **✅ 选择 REST API**

- 构建 **Web 应用、前后端分离** 场景（如 React/Vue + 后端）。
- 需要对外暴露 **公共 API**，让第三方易于调用。
- 关注 **可读性、调试性**，希望直接用浏览器访问。

## 使用 Go 语言标准库实现 rpc

以下是使用 Go 语言标准库 `net/rpc` 实现 RPC 的完整示例，包含 **TCP 协议**和 **HTTP 协议**两种实现方式：

### 🔹基于 TCP 的标准 RPC 实现

目录结构为：

```tree
.
├── go.mod
├── tcp-rpc
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
```

我们将实现一个简单的 **计算服务（Calculator）**，提供加法方法 `Multiply`，让客户端可以远程调用。

#### **服务端代码**

```go
// /tcp-rpc/server/main.go
package main

import (
 "net"
 "net/rpc"
)

// 定义服务结构体
type MathService struct{}

// 实现服务方法（需满足RPC方法签名：func (t *T) MethodName(argType T1, replyType *T2) error）
func (s *MathService) Multiply(args *int, reply *int) error {
 *reply = *args * 10
 return nil
}

func main() {
 // 注册RPC服务
 rpc.Register(new(MathService))

 // 监听TCP端口
 listener, err := net.Listen("tcp", ":1234")
 if err != nil {
  panic(err)
 }
 defer listener.Close()

 // 处理连接
 for {
  conn, err := listener.Accept()
  if err != nil {
   continue
  }
  go rpc.ServeConn(conn) // 为每个连接启动RPC服务
 }
}
```

#### **客户端代码**

```go
// /tcp-rpc/client/main.go
package main

import (
 "fmt"
 "net/rpc"
)

func main() {
 // 连接RPC服务端
 client, err := rpc.Dial("tcp", "localhost:1234") // 端口必须跟服务端保持一致
 if err != nil {
  panic(err)
 }
 defer client.Close()

 // 调用远程方法
 var reply int
 err = client.Call("MathService.Multiply", 5, &reply)
 if err != nil {
  panic(err)
 }

 fmt.Println("Result:", reply) // 输出：Result: 50
}
```

在终端中分别执行 `go run server/main.go` 和 `go run client/main.go` 两个命令，可以就可以看到 client 终端显示`Result: 50` 的结果:

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/tcxqxT-image-20250302233110724.png)

### 🔹基于 HTTP 的 JSON-RPC 实现

仍然是根据上面 **计算服务** 的示例实现！目录结果如下：

```tree
.
├── go.mod
├── json-rpc
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
```

#### **服务端代码**

```go
package main

import (
 "net"
 "net/rpc"
 "net/rpc/jsonrpc"
)

type MathService struct{}

func (s *MathService) Multiply(args *int, reply *int) error {
 *reply = *args * 10
 return nil
}

func main() {
 rpc.Register(new(MathService))

 // 注册HTTP处理器
 rpc.HandleHTTP()

 // 启动HTTP服务
 l, e := net.Listen("tcp", ":9092")
 if e != nil {
  panic(e)
 }

 for {
  conn, _ := l.Accept()

  // 使用 JSON RPC 处理连接
  rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
 }
}
```

#### **客户端代码**

```go
package main

import (
 "fmt"
 "net/rpc/jsonrpc"
)

func main() {
 client, err := jsonrpc.Dial("tcp", "localhost:9092")
 if err != nil {
  panic(err)
 }
 defer client.Close()

 var reply int
 err = client.Call("MathService.Multiply", 5, &reply)
 if err != nil {
  panic(err)
 }

 fmt.Println("Result:", reply) // 输出：Result: 50
}

```

在对应的目录下面，执行 `go run server/main.go` 和 `go run client/main.go` 两个命令，如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/8jedup-image-20250302235136222.png)

> 关键规则说明
>
> 1. **方法签名要求**：
>
>    ```go
>    func (t *T) MethodName(argType T1, replyType *T2) error
>    ```
>
>    - 方法必须为**导出类型**（首字母大写）
>    - 第二个参数必须为指针类型（用于接收返回值）
>    - 必须返回 `error` 类型
>
> 2. **通信协议**：
>    - `net/rpc` 默认使用 **Gob 编码**（二进制协议）
>    - `net/rpc/jsonrpc` 使用 **JSON 编码**（可跨语言调用）
>
> 3. **服务注册**：
>
>    ```go
>    rpc.Register(service)      // 注册服务实例
>    rpc.HandleHTTP()           // HTTP协议专用注册
>    ```

通过标准库可快速实现 RPC 基础功能，但若需**服务发现**、**负载均衡**等高级特性，建议使用更完善的框架（如 [gRPC-Go](https://github.com/grpc/grpc-go) 或 [rpcx](https://github.com/smallnest/rpcx)）。

> 上面所有代码都可以在 https://github.com/clin211/grpc 中找到!


## 思考

大家如果有兴趣的话，可以试试其他语言的调用，比如 Java、Python、Rust，甚至 Node.js 的实现！
