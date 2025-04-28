# Shell 脚本高级特性

Shell 脚本不仅可以用于简单的自动化任务，还可以实现更复杂的编程概念。本文档将介绍 Shell 脚本的高级特性，包括函数的高级用法、模块化、封装和类似继承的概念实现等。

## 函数的高级用法

### 函数库

函数库是一种将常用函数集中存放在单独文件中的方法，可以被多个脚本引用，实现代码复用。

#### 创建函数库

```bash
#!/bin/bash
# 文件名: utils.sh

# 日志函数
log_info() {
    echo "[INFO] $(date +"%Y-%m-%d %H:%M:%S") - $1"
}

log_error() {
    echo "[ERROR] $(date +"%Y-%m-%d %H:%M:%S") - $1" >&2
}

log_warn() {
    echo "[WARN] $(date +"%Y-%m-%d %H:%M:%S") - $1"
}

# 检查命令是否存在
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# 安全地创建目录
safe_mkdir() {
    if [ ! -d "$1" ]; then
        mkdir -p "$1" && log_info "目录创建成功: $1" || log_error "无法创建目录: $1"
    else
        log_info "目录已存在: $1"
    fi
}
```

#### 使用函数库

```bash
#!/bin/bash
# 文件名: script.sh

# 引入函数库
source ./utils.sh
# 或者使用 . 命令
# . ./utils.sh

# 使用库中的函数
log_info "脚本开始执行"

if command_exists "docker"; then
    log_info "Docker 已安装"
else
    log_error "Docker 未安装，请先安装 Docker"
    exit 1
fi

safe_mkdir "./backup"
log_info "脚本执行完毕"
```

### 函数环境变量

#### 局部变量

使用 `local` 关键字可以创建函数内部的局部变量，避免全局命名空间污染：

```bash
#!/bin/bash

global_var="我是全局变量"

test_scope() {
    local local_var="我是局部变量"
    echo "函数内部 - 全局变量: $global_var"
    echo "函数内部 - 局部变量: $local_var"
}

test_scope
echo "函数外部 - 全局变量: $global_var"
echo "函数外部 - 局部变量: $local_var"  # 这将不会输出任何内容
```

#### 环境变量导出

可以使用 `export` 将函数导出，使其在子 Shell 中可用：

```bash
#!/bin/bash

# 导出函数
hello() {
    echo "Hello, $1!"
}
export -f hello

# 在子 Shell 中调用函数
bash -c 'hello World'
```

### 高级函数技巧

#### 函数默认参数

Shell 不直接支持默认参数，但可以模拟实现：

```bash
greet() {
    local name=${1:-"Guest"}  # 如果 $1 未设置或为空，则使用 "Guest"
    echo "Hello, $name!"
}

greet           # 输出: Hello, Guest!
greet "Alice"   # 输出: Hello, Alice!
```

#### 可变参数函数

```bash
sum_all() {
    local total=0
    for num in "$@"; do
        total=$((total + num))
    done
    echo "总和: $total"
}

sum_all 1 2 3 4 5  # 输出: 总和: 15
```

#### 函数返回复杂数据

Shell 函数的 `return` 只能返回 0-255 的数字，但可以通过 echo 输出或者全局变量来返回复杂数据：

```bash
# 方法 1: 通过 echo 返回
get_user_info() {
    echo "张三,25,男"
}

# 方法 2: 通过全局变量返回
parse_user_info() {
    local info="$1"
    IFS=',' read -r USER_NAME USER_AGE USER_GENDER <<< "$info"
}

# 使用方法 1 的结果
user_info=$(get_user_info)
echo "用户信息: $user_info"

# 使用方法 2 处理数据
parse_user_info "$user_info"
echo "姓名: $USER_NAME, 年龄: $USER_AGE, 性别: $USER_GENDER"
```

## 模块化与封装

### 模块化设计

良好的模块化设计可以使 Shell 脚本更易于维护和扩展。

#### 目录结构

```
project/
├── main.sh          # 主脚本
├── lib/             # 库目录
│   ├── logger.sh    # 日志模块
│   ├── config.sh    # 配置模块
│   └── utils.sh     # 工具函数
├── modules/         # 功能模块
│   ├── backup.sh    # 备份模块
│   └── cleanup.sh   # 清理模块
└── config/          # 配置目录
    └── settings.cfg # 配置文件
```

#### 模块引入机制

```bash
#!/bin/bash
# 文件名: main.sh

# 设置基础路径
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LIB_DIR="$BASE_DIR/lib"
MODULE_DIR="$BASE_DIR/modules"
CONFIG_DIR="$BASE_DIR/config"

# 引入库
source "$LIB_DIR/logger.sh"
source "$LIB_DIR/config.sh"
source "$LIB_DIR/utils.sh"

# 加载配置
load_config "$CONFIG_DIR/settings.cfg"

# 根据配置决定加载哪些模块
if [ "$ENABLE_BACKUP" = "true" ]; then
    source "$MODULE_DIR/backup.sh"
fi

if [ "$ENABLE_CLEANUP" = "true" ]; then
    source "$MODULE_DIR/cleanup.sh"
fi

# 主程序
log_info "开始执行主程序"

# 执行功能
if [ "$ENABLE_BACKUP" = "true" ]; then
    perform_backup
fi

if [ "$ENABLE_CLEANUP" = "true" ]; then
    perform_cleanup
fi

log_info "主程序执行完毕"
```

### 封装技术

#### 命名空间模拟

Shell 不直接支持命名空间，但可以通过前缀约定来模拟：

```bash
#!/bin/bash
# 文件名: math.sh

# 数学相关函数都使用 math_ 前缀
math_add() {
    echo $(($1 + $2))
}

math_subtract() {
    echo $(($1 - $2))
}

math_multiply() {
    echo $(($1 * $2))
}

math_divide() {
    if [ $2 -eq 0 ]; then
        echo "错误: 除数不能为零"
        return 1
    fi
    echo $(($1 / $2))
}
```

#### 私有函数

Shell 中没有真正的私有函数，但可以通过命名约定来表示某函数是"私有"的：

```bash
#!/bin/bash
# 文件名: backup.sh

# 公共函数
backup_start() {
    log_info "开始备份"
    _backup_check_prereqs
    _backup_create_dirs
    _backup_copy_files
    log_info "备份完成"
}

# 私有函数（以下划线开头）
_backup_check_prereqs() {
    log_info "检查先决条件"
    # 检查逻辑...
}

_backup_create_dirs() {
    log_info "创建备份目录"
    # 创建目录逻辑...
}

_backup_copy_files() {
    log_info "复制文件"
    # 复制文件逻辑...
}
```

## 类似继承的实现

Shell 脚本不是面向对象的语言，没有真正的类和继承机制，但可以模拟类似的行为。

### 基础"类"模式

```bash
#!/bin/bash
# 文件名: animal.sh

# 定义"基类"
Animal_create() {
    local name=$1
    
    # 属性
    eval "Animal_${name}_type=\"Animal\""
    eval "Animal_${name}_sound=\"\""
    
    # 方法
    eval "Animal_${name}_speak() {
        echo \"\$Animal_${name}_type says: \$Animal_${name}_sound\"
    }"
    
    eval "Animal_${name}_set_sound() {
        Animal_${name}_sound=\"\$1\"
    }"
}

# 使用方法:
# Animal_create "pet1"
# Animal_pet1_set_sound "generic sound"
# Animal_pet1_speak
```

### "继承"模式

```bash
#!/bin/bash
# 文件名: dog.sh

# 引入"基类"
source ./animal.sh

# 定义"子类"
Dog_create() {
    local name=$1
    
    # 首先创建"父类"实例
    Animal_create "$name"
    
    # 覆盖属性
    eval "Animal_${name}_type=\"Dog\""
    eval "Animal_${name}_sound=\"Woof\""
    
    # 添加新方法
    eval "Animal_${name}_fetch() {
        echo \"Dog ${name} is fetching the ball!\"
    }"
}

# 使用方法:
# Dog_create "rover"
# Animal_rover_speak  # 输出: Dog says: Woof
# Animal_rover_fetch  # 输出: Dog rover is fetching the ball!
```

### 示例应用

```bash
#!/bin/bash
# 文件名: pets.sh

source ./animal.sh
source ./dog.sh

# 创建动物实例
Animal_create "generic"
Animal_generic_set_sound "some noise"

# 创建狗实例
Dog_create "rover"

# 测试方法
echo "测试基本动物:"
Animal_generic_speak

echo "测试狗:"
Animal_rover_speak
Animal_rover_fetch
```

## 高级错误处理

### 陷阱设置

使用 `trap` 命令可以捕获信号并执行自定义操作：

```bash
#!/bin/bash

cleanup() {
    echo "执行清理操作..."
    # 删除临时文件等清理工作
    rm -f /tmp/tempfile-$$
    echo "清理完成"
}

# 设置陷阱，在脚本退出前执行清理
trap cleanup EXIT

# 在收到中断信号时执行特定操作
trap 'echo "接收到中断信号，正在退出..."; exit 1' INT

echo "创建临时文件..."
touch /tmp/tempfile-$$

echo "脚本继续运行..."
sleep 5
echo "脚本正常结束"
```

### 调试技术

```bash
#!/bin/bash

# 启用调试模式
debug_mode() {
    set -x  # 开启命令跟踪
}

# 关闭调试模式
normal_mode() {
    set +x  # 关闭命令跟踪
}

echo "正常执行"
debug_mode
echo "这行会显示执行细节"
normal_mode
echo "恢复正常执行"
```

## 高级文本处理

### 使用 awk 进行复杂文本处理

```bash
process_log() {
    local log_file=$1
    local threshold=$2
    
    awk -v threshold="$threshold" '
    BEGIN {
        print "开始分析日志文件..."
        count = 0
    }
    
    /ERROR/ {
        errors[count++] = $0
    }
    
    END {
        print "发现 " count " 个错误"
        if (count > threshold) {
            print "错误数超过阈值 " threshold
            for (i = 0; i < count; i++) {
                print errors[i]
            }
        }
    }' "$log_file"
}

# 使用示例
# process_log "application.log" 5
```

### 使用 sed 进行高级替换

```bash
batch_replace() {
    local file=$1
    local pattern=$2
    local replacement=$3
    
    sed -i.bak "s#$pattern#$replacement#g" "$file"
    if [ $? -eq 0 ]; then
        echo "替换成功，备份文件为 ${file}.bak"
    else
        echo "替换失败"
    fi
}

# 使用示例
# batch_replace "config.xml" "<version>1.0</version>" "<version>2.0</version>"
```

## 实战案例：监控系统框架

下面是一个模块化的系统监控框架，展示了如何使用高级 Shell 脚本技术构建复杂应用：

```bash
#!/bin/bash
# 文件名: monitor.sh

# 基础目录设置
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG_FILE="$BASE_DIR/config/monitor.cfg"
LOG_DIR="$BASE_DIR/logs"
MODULES_DIR="$BASE_DIR/modules"

# 引入工具函数
source "$BASE_DIR/lib/utils.sh"

# 初始化
init() {
    log_info "初始化监控系统"
    safe_mkdir "$LOG_DIR"
    
    # 加载配置
    if [ -f "$CONFIG_FILE" ]; then
        source "$CONFIG_FILE"
    else
        log_error "配置文件不存在: $CONFIG_FILE"
        exit 1
    fi
    
    # 加载启用的模块
    for module in $ENABLED_MODULES; do
        if [ -f "$MODULES_DIR/${module}.sh" ]; then
            log_info "加载模块: $module"
            source "$MODULES_DIR/${module}.sh"
        else
            log_warn "模块不存在: $module"
        fi
    done
}

# 运行监控
run() {
    log_info "开始运行监控"
    
    # 调用每个模块的检查函数
    for module in $ENABLED_MODULES; do
        if type -t "${module}_check" &>/dev/null; then
            log_info "执行 $module 检查"
            ${module}_check
        else
            log_warn "模块 $module 没有实现 check 函数"
        fi
    done
    
    log_info "监控运行完成"
}

# 主程序
init
run

log_info "监控系统退出"
```

模块示例 (cpu.sh):

```bash
#!/bin/bash
# 文件名: modules/cpu.sh

# CPU 模块配置
CPU_WARNING_THRESHOLD=${CPU_WARNING_THRESHOLD:-80}
CPU_CRITICAL_THRESHOLD=${CPU_CRITICAL_THRESHOLD:-95}

# CPU 检查函数
cpu_check() {
    log_info "检查 CPU 使用率"
    
    # 获取 CPU 使用率
    local cpu_usage=$(top -l 1 | grep "CPU usage" | awk '{print $3}' | cut -d% -f1)
    
    log_info "当前 CPU 使用率: ${cpu_usage}%"
    
    # 根据阈值判断
    if (( $(echo "$cpu_usage >= $CPU_CRITICAL_THRESHOLD" | bc -l) )); then
        log_error "CPU 使用率过高: ${cpu_usage}%"
        [ -n "$NOTIFY_CMD" ] && $NOTIFY_CMD "严重: CPU 使用率 ${cpu_usage}%"
    elif (( $(echo "$cpu_usage >= $CPU_WARNING_THRESHOLD" | bc -l) )); then
        log_warn "CPU 使用率偏高: ${cpu_usage}%"
        [ -n "$NOTIFY_CMD" ] && $NOTIFY_CMD "警告: CPU 使用率 ${cpu_usage}%"
    else
        log_info "CPU 使用率正常"
    fi
}
```

配置文件示例 (monitor.cfg):

```bash
# 监控系统配置

# 启用的模块
ENABLED_MODULES="cpu disk memory network"

# 通知命令
NOTIFY_CMD="./lib/notify.sh"

# CPU 模块配置
CPU_WARNING_THRESHOLD=75
CPU_CRITICAL_THRESHOLD=90

# 磁盘模块配置
DISK_WARNING_THRESHOLD=85
DISK_CRITICAL_THRESHOLD=95
DISK_MONITOR_PATHS="/ /home"

# 内存模块配置
MEMORY_WARNING_THRESHOLD=80
MEMORY_CRITICAL_THRESHOLD=90

# 网络模块配置
NETWORK_INTERFACES="eth0 wlan0"
```

## 练习任务

1. 创建一个文件操作函数库，包含复制、移动、删除、备份等常用文件操作，每个操作都包含错误处理和日志记录。

2. 实现一个配置管理系统，可以读取、解析、修改和保存配置文件，支持不同的配置文件格式。

3. 扩展监控系统框架，添加新的监控模块（如进程监控、服务监控等）。

4. 实现一个简单的命令行参数解析框架，支持短选项、长选项和参数值。

## 小结

本文档介绍了 Shell 脚本的高级特性，包括：

- 函数的高级用法和函数库
- 模块化和封装技术
- 模拟面向对象的"继承"
- 高级错误处理和调试
- 高级文本处理
- 实战案例

掌握这些高级特性，可以让你编写更加健壮、可维护和可扩展的 Shell 脚本，满足更复杂的自动化需求。

## 进阶资源

- 《Advanced Bash-Scripting Guide》- Shell 脚本高级编程的详细指南
- 《Linux Shell Scripting Cookbook》- 包含许多实用的高级脚本示例
- <https://github.com/dylanaraps/pure-bash-bible> - 纯 Bash 解决方案集合
