# Shell 脚本基础

Shell 脚本是一种将多个命令组织在一起的文本文件，它可以被 Shell 程序（如 bash、zsh 等）解释和执行。通过学习 Shell 脚本，你可以自动化日常任务，提高工作效率。

## 创建你的第一个 Shell 脚本

### 脚本文件创建

1. 使用文本编辑器创建一个新文件：

   ```bash
   touch hello.sh
   ```

2. 用编辑器打开文件：

   ```bash
   nano hello.sh   # 或使用其他编辑器如 vim、VSCode 等
   ```

3. 添加以下内容：

   ```bash
   #!/bin/bash
   
   echo "Hello, World!"
   echo "这是我的第一个 Shell 脚本"
   ```

4. 保存文件并退出编辑器。

### 脚本解释

- 第一行 `#!/bin/bash` 被称为 "shebang" 行，它告诉系统使用哪个 Shell 解释器来执行脚本。
- 后续行是要执行的命令，与你在命令行中输入的完全相同。

### 执行脚本

在执行脚本之前，需要给文件添加执行权限：

```bash
chmod +x hello.sh
```

然后，可以通过以下方式执行脚本：

```bash
./hello.sh
```

输出应该是：

```
Hello, World!
这是我的第一个 Shell 脚本
```

## 变量

Shell 脚本中的变量用于存储数据，可以在脚本中多次引用。

### 变量定义

```bash
NAME="张三"
AGE=25
```

注意：等号两边**不能有空格**。

### 变量使用

使用变量时，需要在变量名前加上 `$` 符号：

```bash
echo "我的名字是 $NAME"
echo "我的年龄是 $AGE 岁"
```

也可以使用花括号来明确变量名的边界：

```bash
echo "你好，${NAME}！"
```

### 命令结果存储

可以将命令的执行结果存储到变量中：

```bash
# 使用反引号
CURRENT_DIR=`pwd`

# 或使用 $() 语法（推荐）
CURRENT_DATE=$(date +%Y-%m-%d)

echo "当前目录是: $CURRENT_DIR"
echo "今天的日期是: $CURRENT_DATE"
```

### 只读变量

使用 `readonly` 命令可以将变量设为只读：

```bash
readonly PI=3.14159
PI=3.14  # 这会导致错误，因为 PI 是只读的
```

### 删除变量

使用 `unset` 命令可以删除变量：

```bash
NAME="李四"
echo $NAME
unset NAME
echo $NAME  # 不会输出任何内容
```

## 脚本参数

Shell 脚本可以接收命令行参数，这些参数可以在脚本中使用。

### 位置参数

```bash
#!/bin/bash

echo "脚本名称: $0"
echo "第一个参数: $1"
echo "第二个参数: $2"
echo "所有参数: $@"
echo "参数个数: $#"
```

如果你执行 `./params.sh hello world`，输出将是：

```
脚本名称: ./params.sh
第一个参数: hello
第二个参数: world
所有参数: hello world
参数个数: 2
```

### 特殊参数

- `$0` - 脚本名称
- `$1` 到 `$9` - 第 1 个到第 9 个参数
- `${10}` 及以上 - 第 10 个及以上的参数（需要使用花括号）
- `$@` - 所有参数（作为独立的字符串）
- `$*` - 所有参数（作为单个字符串）
- `$#` - 参数个数
- `$$` - 当前 Shell 的进程 ID
- `$?` - 最后一个命令的退出状态码

## 条件判断

Shell 脚本中可以使用条件语句来实现逻辑判断。

### if 语句

基本语法：

```bash
if [ 条件 ]
then
    # 如果条件为真，执行这里的命令
fi
```

if-else 语法：

```bash
if [ 条件 ]
then
    # 如果条件为真，执行这里的命令
else
    # 如果条件为假，执行这里的命令
fi
```

if-elif-else 语法：

```bash
if [ 条件 1 ]
then
    # 如果条件 1 为真，执行这里的命令
elif [ 条件 2 ]
then
    # 如果条件 1 为假且条件 2 为真，执行这里的命令
else
    # 如果所有条件都为假，执行这里的命令
fi
```

### 条件测试

文件测试：

```bash
if [ -f file.txt ]  # 如果文件存在且是普通文件
then
    echo "file.txt 存在"
fi

if [ -d dir ]  # 如果目录存在
then
    echo "dir 目录存在"
fi
```

常见的文件测试操作符：

- `-e file` : 如果文件存在，返回真
- `-f file` : 如果文件存在且是普通文件，返回真
- `-d file` : 如果文件存在且是目录，返回真
- `-r file` : 如果文件存在且可读，返回真
- `-w file` : 如果文件存在且可写，返回真
- `-x file` : 如果文件存在且可执行，返回真
- `-s file` : 如果文件存在且大小不为 0，返回真

字符串比较：

```bash
if [ "$NAME" = "张三" ]  # 字符串相等
then
    echo "你好，张三"
fi

if [ -z "$VAR" ]  # 如果字符串长度为 0
then
    echo "VAR 为空"
fi
```

常见的字符串比较操作符：

- `=` : 相等
- `!=` : 不相等
- `-z string` : 如果字符串长度为 0，返回真
- `-n string` : 如果字符串长度不为 0，返回真

数值比较：

```bash
if [ "$AGE" -eq 25 ]  # 数值相等
then
    echo "年龄是 25 岁"
fi

if [ "$COUNT" -gt 10 ]  # 大于
then
    echo "COUNT 大于 10"
fi
```

常见的数值比较操作符：

- `-eq` : 相等
- `-ne` : 不相等
- `-gt` : 大于
- `-lt` : 小于
- `-ge` : 大于等于
- `-le` : 小于等于

逻辑操作符：

```bash
if [ "$AGE" -ge 18 ] && [ "$AGE" -le 60 ]  # 逻辑与
then
    echo "成年人"
fi

if [ "$DAY" = "星期六" ] || [ "$DAY" = "星期日" ]  # 逻辑或
then
    echo "周末"
fi

if [ ! -f file.txt ]  # 逻辑非
then
    echo "file.txt 不存在"
fi
```

### case 语句

`case` 语句用于多条件分支判断，类似于其他语言中的 switch 语句：

```bash
#!/bin/bash

echo "请输入一个水果名称:"
read FRUIT

case $FRUIT in
    "苹果")
        echo "苹果的价格是 5 元/斤"
        ;;
    "香蕉")
        echo "香蕉的价格是 3 元/斤"
        ;;
    "桔子")
        echo "桔子的价格是 4 元/斤"
        ;;
    *)
        echo "没有 $FRUIT 的价格信息"
        ;;
esac
```

## 循环

Shell 脚本中提供了几种循环结构。

### for 循环

```bash
#!/bin/bash

# 遍历列表
for FRUIT in 苹果 香蕉 桔子
do
    echo "我喜欢吃 $FRUIT"
done

# 遍历数字范围
for NUM in {1..5}
do
    echo "数字: $NUM"
done

# 使用 seq 生成序列
for NUM in $(seq 1 2 10)  # 从 1 到 10，步长为 2
do
    echo "数字: $NUM"
done

# C 语言风格的 for 循环
for ((i=1; i<=5; i++))
do
    echo "计数: $i"
done
```

### while 循环

```bash
#!/bin/bash

# 基本 while 循环
COUNT=1
while [ $COUNT -le 5 ]
do
    echo "COUNT: $COUNT"
    COUNT=$((COUNT+1))
done

# 无限循环
# while true
# do
#     echo "按 Ctrl+C 退出"
#     sleep 1
# done
```

### until 循环

`until` 循环与 `while` 相反，它在条件为假时执行循环体：

```bash
#!/bin/bash

COUNT=1
until [ $COUNT -gt 5 ]
do
    echo "COUNT: $COUNT"
    COUNT=$((COUNT+1))
done
```

### 循环控制

- `break` - 跳出当前循环
- `continue` - 跳过当前迭代，继续下一次迭代

```bash
#!/bin/bash

for NUM in {1..10}
do
    if [ $NUM -eq 5 ]
    then
        echo "跳过 5"
        continue
    fi
    
    if [ $NUM -eq 8 ]
    then
        echo "在 8 处终止循环"
        break
    fi
    
    echo "数字: $NUM"
done
```

## 函数

函数可以让你将一组相关的命令打包在一起，以便重复使用。

### 定义函数

```bash
#!/bin/bash

# 函数定义
hello() {
    echo "Hello, World!"
}

greet() {
    echo "你好, $1!"
}

add() {
    return $(($1 + $2))
}

# 函数调用
hello

greet "张三"

add 5 3
echo "5 + 3 = $?"
```

### 函数参数

函数内部可以使用 `$1`, `$2` 等访问传入的参数，就像脚本参数一样。

### 函数返回值

函数可以使用 `return` 语句返回一个数值（0-255 之间的整数），表示函数的退出状态。要获取函数的计算结果，通常使用命令替换。

```bash
sum() {
    echo $(($1 + $2))  # 通过 echo 输出结果
}

result=$(sum 10 20)
echo "10 + 20 = $result"
```

## 实例分析：watch-chrome.sh

让我们分析之前创建的 `watch-chrome.sh` 脚本，理解其工作原理：

```bash
#!/bin/bash

# Chrome 目录
DATA_DIR="$HOME/Library/Application Support/Google/Chrome"
CACHE_DIR="$HOME/Library/Caches/Google/Chrome"
APP_DIR="/Applications/Google Chrome.app/Contents"

# 检查 fswatch 是否安装
if ! command -v fswatch > /dev/null; then
    echo "此脚本需要 fswatch，请使用 Homebrew 安装："
    echo "brew install fswatch"
    exit 1
fi

# 创建时间戳函数
timestamp() {
  date +"%Y-%m-%d %H:%M:%S"
}

# 处理文件变化事件的函数
handle_change() {
    local file=$1
    local event=$2
    local dir_type=$3
    
    echo "$(timestamp) [$dir_type] $file: $event"
}

# 监控目录的函数
monitor_directory() {
    local dir=$1
    local dir_type=$2
    
    echo "开始监控 $dir_type 目录: $dir"
    
    fswatch -r "$dir" | while read file event; do
        handle_change "$file" "$event" "$dir_type"
    done &
}

echo "=== Chrome 目录监控 ==="
echo "按 Ctrl+C 停止监控"
echo

# 开始监控每个目录
monitor_directory "$DATA_DIR" "数据"
monitor_directory "$CACHE_DIR" "缓存"
monitor_directory "$APP_DIR" "应用"

# 保持脚本运行
wait
```

这个脚本的主要组成部分：

1. **变量定义**：定义了要监控的 Chrome 目录
2. **依赖检查**：检查 fswatch 命令是否可用
3. **函数定义**：
   - `timestamp()` - 生成时间戳
   - `handle_change()` - 处理文件变化事件
   - `monitor_directory()` - 设置目录监控
4. **主程序流程**：
   - 显示欢迎信息
   - 启动三个监控进程（在后台运行）
   - 使用 `wait` 命令保持脚本运行

脚本使用 `fswatch` 命令监控目录变化，并为每个变化的文件输出一条带有时间戳的日志信息。

## 练习任务

1. 创建一个脚本，接收文件名作为参数，检查文件是否存在，如果存在则显示其内容，否则创建该文件。

2. 编写一个脚本，计算 1 到用户输入的数字之间的所有整数的和。

3. 编写一个备份脚本，将指定目录的内容复制到一个以当前日期命名的备份目录中。

4. 修改 watch-chrome.sh 脚本，添加以下功能：
   - 将监控日志保存到文件
   - 添加一个选项可以指定要监控的特定文件类型

## 小结

本文档介绍了 Shell 脚本的基础知识，包括：

- 脚本创建和执行
- 变量使用
- 条件判断
- 循环结构
- 函数定义与调用

通过这些基础知识，你应该能够理解和修改简单的 Shell 脚本，并开始编写自己的脚本来自动化日常任务。

随着你对 Shell 脚本的深入学习，你将能够创建更复杂和功能更强大的脚本，进一步提高工作效率。

## 进阶资源

- 《Linux 命令行与 Shell 脚本编程大全》- 深入学习 Shell 脚本的全面资源
- <https://www.shellscript.sh/> - Shell 脚本教程网站
- `man bash` - Bash 手册中关于脚本编程的部分
