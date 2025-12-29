# Vim 学习指南

## 基础知识

### Vim 的历史与理念

Vim（Vi IMproved）是 Vi 编辑器的增强版本，由 Bram Moolenaar 于 1991 年发布。Vim 的设计理念是提高文本编辑效率，让编辑过程如同思考般流畅。Vim 强调通过键盘完成所有操作，避免鼠标的使用，从而提高编辑速度。

### Vim 的模式

Vim 有几种主要模式，理解这些模式是掌握 Vim 的基础：

- **普通模式（Normal Mode）**：默认模式，用于导航和操作文本。在此模式下输入的按键被解释为命令而非文本。
- **插入模式（Insert Mode）**：用于输入文本，类似于其他编辑器的常规编辑模式。
- **可视模式（Visual Mode）**：用于选择文本块，可以是字符级、行级或块级选择。
- **命令模式（Command Mode）**：用于执行保存、退出、搜索、替换等命令。
- **Ex 模式（Ex Mode）**：用于执行多行命令。
- **替换模式（Replace Mode）**：覆盖已有文本的模式。

理解这些模式的切换和各自特点是学习 Vim 的第一步。

### 安装 Vim

#### Linux

大多数 Linux 发行版默认已安装 Vim 或 Vi。如需安装：

- Ubuntu/Debian: `sudo apt install vim`
- CentOS/RHEL: `sudo yum install vim`
- Arch Linux: `sudo pacman -S vim`

#### macOS

macOS 自带 Vim，也可以通过 Homebrew 安装：

```bash
brew install vim
```

#### Windows

从 [Vim 官网](https://www.vim.org/download.php) 下载安装程序，或通过 Chocolatey：

```bash
choco install vim
```

### 开始使用 Vim

打开终端，输入 `vim 文件名` 即可打开或创建文件。如果文件不存在，会在保存时创建。

首次打开 Vim 会进入普通模式。记住：在普通模式下输入的任何字符都被视为命令，而非文本输入。

## 基本操作

### 模式切换

- `i` - 进入插入模式（在光标位置插入）
- `I` - 进入插入模式（在行首插入）
- `a` - 进入插入模式（在光标后插入）
- `A` - 进入插入模式（在行尾插入）
- `o` - 进入插入模式（在当前行下方新建一行）
- `O` - 进入插入模式（在当前行上方新建一行）
- `Esc` - 从任何模式返回普通模式（可按多次确保返回）
- `:` - 从普通模式进入命令模式
- `v` - 从普通模式进入字符可视模式
- `V` - 从普通模式进入行可视模式
- `Ctrl+v` - 从普通模式进入块可视模式
- `R` - 从普通模式进入替换模式（覆盖已有文本）

**提示**：养成使用 `Esc` 回到普通模式的习惯，这是 Vim 操作的基础。也可以使用 `Ctrl+[` 代替 `Esc`。

### 退出 Vim

- `:q` - 退出（如有未保存更改则无法退出）
- `:q!` - 强制退出（丢弃未保存更改）
- `:wq` - 保存并退出
- `ZZ` - 保存并退出（快捷方式，普通模式下使用）
- `:w` - 保存但不退出
- `:w filename` - 另存为其他文件名
- `:wq!` - 强制保存并退出（当文件为只读但你有权限覆盖时）
- `:x` 或 `:exit` - 保存（如有更改）并退出
- `:qa` - 退出所有打开的文件
- `:qa!` - 强制退出所有打开的文件

**提示**：初学者常见问题是无法退出 Vim。记住先按 `Esc` 回到普通模式，然后输入 `:q!`。

### 帮助系统

Vim 有强大的内置帮助系统：

- `:help` 或 `:h` - 打开帮助文档
- `:help {主题}` - 查询特定主题，如 `:help insert`
- `:help 'option'` - 查询特定选项，如 `:help 'number'`
- `Ctrl+]` - 在帮助中跳转到光标下的主题
- `Ctrl+o` - 跳回上一个位置
- `:helpgrep {pattern}` - 在所有帮助文档中搜索
- `:h tutor` - 获取关于教程的帮助

**提示**：不确定命令时，尝试使用 `:help` 查找。例如想知道 `dd` 做什么，可以用 `:help dd`。

### 保存文件

- `:w` - 保存当前文件
- `:w filename` - 将当前内容保存到指定文件
- `:saveas filename` - 另存为新文件并切换到新文件
- `:w >> filename` - 将当前内容追加到指定文件
- `:w !sudo tee %` - 用 sudo 保存当前文件（需要权限时）

**提示**：% 是 Vim 中表示当前文件名的特殊符号。

## 移动光标（普通模式）

### 基本移动

- `h` - 左移一个字符
- `j` - 下移一行
- `k` - 上移一行
- `l` - 右移一个字符
- `5j` - 向下移动 5 行（数字前缀可与任何移动命令结合使用）
- `gj` - 在自动换行的行中向下移动（视觉行而非实际行）
- `gk` - 在自动换行的行中向上移动（视觉行而非实际行）

**提示**：使用 `hjkl` 而非箭头键可以提高效率，因为无需移动手指离开主键盘区域。

### 单词级移动

- `w` - 移动到下一个单词开头（标点视为单独的单词）
- `W` - 移动到下一个大单词开头（忽略标点）
- `b` - 移动到上一个单词开头（标点视为单独的单词）
- `B` - 移动到上一个大单词开头（忽略标点）
- `e` - 移动到当前单词结尾（标点视为单独的单词）
- `E` - 移动到当前大单词结尾（忽略标点）
- `ge` - 移动到上一个单词结尾
- `gE` - 移动到上一个大单词结尾

**提示**：`W`、`B`、`E` 等大写命令将标点符号视为单词的一部分，而非单独的单词。

### 行级移动

- `0` - 移动到行首（包括空白字符）
- `^` - 移动到行首非空字符
- `$` - 移动到行尾
- `g_` - 移动到行尾非空字符
- `|` - 移动到行的第 n 列（如 `5|` 移动到第 5 列）
- `f{char}` - 向前查找字符并移动到该字符（如 `fa` 查找下一个 'a'）
- `F{char}` - 向后查找字符并移动到该字符
- `t{char}` - 向前查找字符并移动到该字符前一个位置
- `T{char}` - 向后查找字符并移动到该字符后一个位置
- `;` - 重复上次 f、F、t 或 T 命令
- `,` - 反方向重复上次 f、F、t 或 T 命令

**提示**：`f` 和 `t` 的区别在于 `f` 会将光标移动到目标字符上，而 `t` 会移动到字符前。

### 屏幕移动

- `H` - 移动到屏幕顶部（Home）
- `M` - 移动到屏幕中间（Middle）
- `L` - 移动到屏幕底部（Low）
- `zt` - 重绘屏幕，使当前行位于屏幕顶部
- `zz` - 重绘屏幕，使当前行位于屏幕中央
- `zb` - 重绘屏幕，使当前行位于屏幕底部
- `Ctrl+e` - 向下滚动一行（光标位置不变）
- `Ctrl+y` - 向上滚动一行（光标位置不变）
- `Ctrl+d` - 向下滚动半页
- `Ctrl+u` - 向上滚动半页
- `Ctrl+f` - 向下滚动一页（Page Down）
- `Ctrl+b` - 向上滚动一页（Page Up）

**提示**：`zz` 是查看上下文很有用的命令，它使当前行居中显示。

### 文档移动

- `gg` - 移动到文档开头（第一行）
- `G` - 移动到文档结尾（最后一行）
- `:{number}` 或 `{number}G` - 移动到指定行号（如 `:42` 或 `42G` 跳到第 42 行）
- `%` - 跳转到匹配的括号（光标需位于括号上）
- `{` - 跳转到上一个空行
- `}` - 跳转到下一个空行
- `(` - 跳转到上一个句子
- `)` - 跳转到下一个句子
- `[[` - 跳转到上一个章节或函数开头
- `]]` - 跳转到下一个章节或函数开头
- `[]` - 跳转到上一个章节或函数结尾
- `][` - 跳转到下一个章节或函数结尾
- `'.` - 跳转到上次编辑的位置
- `g;` - 跳转到上次编辑的位置（可重复使用回溯历史）
- `g,` - 跳转到下一个编辑位置（与 `g;` 相反）
- `Ctrl+o` - 跳回到跳转列表中的上一个位置
- `Ctrl+i` 或 `Tab` - 跳转到跳转列表中的下一个位置

**提示**：Vim 会记住你的跳转历史，`Ctrl+o` 和 `Ctrl+i` 可以在这些位置之间快速移动。

### 搜索跳转

- `/{pattern}` - 向下搜索模式
- `?{pattern}` - 向上搜索模式
- `n` - 重复上次搜索（同方向）
- `N` - 反向重复上次搜索
- `*` - 向下搜索光标下的单词
- `#` - 向上搜索光标下的单词
- `gd` - 跳转到局部变量定义处
- `gD` - 跳转到全局变量定义处
- `/` 加上 `↑` 或 `↓` - 可以查看搜索历史

**提示**：搜索支持正则表达式。使用 `:set hlsearch` 可高亮显示所有匹配项，`:noh` 可以临时关闭高亮。

## 编辑操作

### 插入文本

- `i` - 在光标前插入
- `I` - 在行首非空字符处插入
- `a` - 在光标后插入
- `A` - 在行尾插入
- `o` - 在当前行下方新建一行并插入
- `O` - 在当前行上方新建一行并插入
- `gi` - 在上次离开插入模式的位置继续插入
- `gI` - 在行首插入
- `Esc` 或 `Ctrl+[` - 退出插入模式，返回普通模式

**提示**：在插入模式中，可以使用 `Ctrl+o` 临时进入普通模式执行一个命令，然后自动返回插入模式。

### 删除文本

- `x` - 删除光标处字符
- `X` - 删除光标前字符
- `r{char}` - 替换光标处字符为指定字符
- `R` - 进入替换模式（覆盖已有文本）
- `dd` - 删除整行（并复制到寄存器）
- `D` - 删除光标到行尾（等同于 `d$`）
- `d0` - 删除光标到行首
- `dw` - 删除从光标到下一个单词开头
- `de` - 删除从光标到当前单词结尾
- `d$` - 删除从光标到行尾
- `d^` - 删除从光标到行首非空字符
- `dG` - 删除从当前行到文档末尾
- `dgg` - 删除从当前行到文档开头
- `diw` - 删除整个单词（不含周围空格）
- `daw` - 删除整个单词（含周围空格）
- `di"` - 删除双引号内的内容
- `da"` - 删除双引号及其内容
- `di(` - 删除括号内的内容
- `da(` - 删除括号及其内容
- `dit` - 删除 XML/HTML 标签内的内容
- `dat` - 删除 XML/HTML 标签及其内容
- `5dd` - 删除 5 行（数字可变）
- `d/pattern` - 删除从光标到 pattern 匹配处

**提示**：`d` 是删除操作符，可以与各种移动命令组合。格式为 `d{motion}`，其中 motion 是任何移动命令。

### 复制粘贴

- `yy` 或 `Y` - 复制整行
- `y$` - 复制从光标到行尾
- `y^` - 复制从光标到行首非空字符
- `yw` - 复制从光标到单词结尾
- `yiw` - 复制整个单词（不含周围空格）
- `yaw` - 复制整个单词（含周围空格）
- `yi"` - 复制双引号内的内容
- `ya"` - 复制双引号及其内容
- `5yy` - 复制 5 行（数字可变）
- `p` - 在光标后粘贴
- `P` - 在光标前粘贴
- `gp` - 在光标后粘贴并将光标移至粘贴内容后
- `gP` - 在光标前粘贴并将光标移至粘贴内容后
- `:reg` - 显示所有寄存器内容
- `"ay` - 复制到 a 寄存器
- `"ap` - 从 a 寄存器粘贴
- `"+y` - 复制到系统剪贴板（如果支持）
- `"+p` - 从系统剪贴板粘贴（如果支持）

**提示**：`y` 是复制操作符，可以与各种移动命令组合，格式为 `y{motion}`。Vim 使用寄存器存储复制的内容，默认使用无名寄存器。

### 剪切操作

在 Vim 中，删除（`d`）操作实际上是剪切，所以：

- `dd` 后跟 `p` - 剪切并粘贴一行（移动一行）
- `dwP` - 剪切一个单词并在原位置前粘贴（可用于交换单词）

### 撤销与重做

- `u` - 撤销上一次更改
- `Ctrl+r` - 重做（撤销的反操作）
- `U` - 撤销对当前行的所有更改
- `5u` - 撤销 5 次操作
- `:earlier 5m` - 回到 5 分钟前的状态
- `:later 5m` - 前进 5 分钟
- `:earlier 2f` - 回到 2 次文件保存前的状态
- `:undolist` - 显示撤销树
- `g-` - 回到上一个编辑状态
- `g+` - 前进到下一个编辑状态

**提示**：Vim 提供了强大的撤销树功能，允许回到文件的任何历史状态。

### 文本对象

Vim 的文本对象允许你操作文本结构：

- `iw` - 内部单词（inside word）
- `aw` - 整个单词（a word）含周围空格
- `is` - 内部句子（inside sentence）
- `as` - 整个句子（a sentence）
- `ip` - 内部段落（inside paragraph）
- `ap` - 整个段落（a paragraph）
- `i"` - 内部双引号内容
- `a"` - 包含双引号的内容
- `i'` - 内部单引号内容
- `a'` - 包含单引号的内容
- `i)` 或 `ib` - 内部圆括号内容
- `a)` 或 `ab` - 包含圆括号的内容
- `i]` - 内部方括号内容
- `a]` - 包含方括号的内容
- `i}` 或 `iB` - 内部花括号内容
- `a}` 或 `aB` - 包含花括号的内容
- `it` - 内部 HTML/XML 标签内容
- `at` - 包含 HTML/XML 标签的内容

**提示**：文本对象通常与操作符组合使用，如 `di"`（删除双引号内内容）、`ya]`（复制包含方括号的内容）。

### 缩进操作

- `>>` - 向右缩进当前行
- `<<` - 向左缩进当前行
- `>G` - 从当前行到文档末尾向右缩进
- `<G` - 从当前行到文档末尾向左缩进
- `>ip` - 向右缩进当前段落
- `<ip` - 向左缩进当前段落
- `=G` - 从当前行到文档末尾自动缩进
- `=ip` - 自动缩进当前段落
- `=%` - 自动缩进当前代码块（光标需在括号上）
- `gg=G` - 自动缩进整个文档

**提示**：自动缩进功能在编辑代码时特别有用。

### 大小写转换

- `~` - 切换光标所在字符的大小写
- `g~w` - 切换单词的大小写
- `g~$` - 切换从光标到行尾的大小写
- `gUw` - 将单词转换为大写
- `guw` - 将单词转换为小写
- `gUU` 或 `gUgU` - 将整行转换为大写
- `guu` 或 `gugu` - 将整行转换为小写
- `g~~` - 切换整行大小写
- `Vu` - 将整行转换为小写（V 模式下）
- `VU` - 将整行转换为大写（V 模式下）

**提示**：`gU` 和 `gu` 是操作符，可以与任何移动命令组合。

### 连接行

- `J` - 将下一行连接到当前行，保留一个空格
- `gJ` - 将下一行连接到当前行，不保留空格
- `5J` - 连接 5 行（数字可变）

### 替换操作

- `r{char}` - 替换光标所在字符
- `gr{char}` - 虚拟替换模式（不影响布局）
- `R` - 进入替换模式（覆盖已有文本）
- `gR` - 进入虚拟替换模式

### 重复操作

- `.` - 重复上一次更改
- `@:` - 重复上一次命令模式命令
- `@@` - 重复上一次宏操作

**提示**：`.` 命令是 Vim 中最有用的命令之一，它能重复几乎所有更改操作。

## 高级操作

### 查找替换

#### 基本查找

- `/pattern` - 向下查找模式
- `?pattern` - 向上查找模式
- `n` - 继续查找下一个匹配（同方向）
- `N` - 反向查找下一个匹配
- `*` - 向下查找光标下的单词
- `#` - 向上查找光标下的单词
- `g*` - 向下查找光标下的单词（部分匹配）
- `g#` - 向上查找光标下的单词（部分匹配）
- `:noh` - 暂时关闭搜索高亮
- `/{pattern}/e` - 查找并将光标放在匹配结尾

#### 基本替换

- `:%s/old/new/g` - 全文替换
- `:%s/old/new/gc` - 全文替换并确认每次替换
- `:s/old/new/g` - 当前行替换
- `:5,12s/old/new/g` - 第 5 行至第 12 行替换
- `:.,$s/old/new/g` - 当前行至文件末尾替换
- `:%s/old/new/gi` - 全文替换（忽略大小写）
- `:%s/\<word\>/new/g` - 全文替换整个单词
- `:%s/old/&/g` - & 代表匹配的文本

#### 高级替换

- `:%s/\(\w\+\)\s\+\(\w\+\)/\2 \1/g` - 交换每行中的前两个单词
- `:%s//new/g` - 使用上次搜索模式进行替换
- `:%s/\v(pattern)/\1/g` - 使用 very magic 模式（简化正则表达式）
- `:%s/\n//g` - 删除所有换行符（合并所有行）
- `g/pattern/d` - 删除包含 pattern 的所有行
- `v/pattern/d` - 删除不包含 pattern 的所有行
- `:g/^$/d` - 删除所有空行
- `:g/pattern/y A` - 将所有匹配行复制到寄存器 a
- `:5,20s/^/#/` - 给第 5 行到第 20 行添加注释符号

**提示**：Vim 的替换功能支持强大的正则表达式。`\(\)` 用于分组，`\1`,`\2` 引用分组内容。

### 可视模式操作

#### 可视模式类型

- `v` - 字符可视模式（选择字符）
- `V` - 行可视模式（选择整行）
- `Ctrl+v` - 块可视模式（选择矩形块）
- `gv` - 重新选择上次的可视区域
- `o` - 在可视模式中切换光标位置（开始/结束）

#### 可视模式下的操作

- 选中后 `d` - 删除选中内容
- 选中后 `y` - 复制选中内容
- 选中后 `c` - 删除选中内容并进入插入模式
- 选中后 `>` - 增加缩进
- 选中后 `<` - 减少缩进
- 选中后 `~` - 切换大小写
- 选中后 `u` - 转换为小写
- 选中后 `U` - 转换为大写
- 选中后 `:normal 命令` - 对每行执行普通模式命令
- 选中后 `!command` - 通过外部命令过滤选中内容
- 选中后 `J` - 合并选中的行

#### 块操作（块可视模式）

- `Ctrl+v` 选择后 `I` - 在块前插入文本
- `Ctrl+v` 选择后 `A` - 在块后插入文本
- `Ctrl+v` 选择后 `r` - 替换块中的每个字符
- `Ctrl+v` 选择后 `c` - 删除块并插入文本
- `Ctrl+v` 选择后 `d` - 删除块
- `Ctrl+v` 选择后 `o` - 移动到块的另一角

**提示**：块可视模式是 Vim 独特的功能，非常适合编辑表格数据或多行代码。

### 多文件操作

#### 缓冲区管理

- `:ls` 或 `:buffers` - 列出所有缓冲区
- `:b {number}` 或 `:b {name}` - 切换到指定缓冲区
- `:bnext` 或 `:bn` - 切换到下一个缓冲区
- `:bprevious` 或 `:bp` - 切换到上一个缓冲区
- `:bfirst` 或 `:bf` - 切换到第一个缓冲区
- `:blast` 或 `:bl` - 切换到最后一个缓冲区
- `:badd {file}` - 添加文件到缓冲区列表
- `:bdelete` 或 `:bd` - 删除当前缓冲区
- `:bufdo {command}` - 对所有缓冲区执行命令
- `:e {file}` - 编辑另一个文件
- `:e!` - 重新加载当前文件（放弃修改）
- `:e #` - 编辑上一个缓冲区文件
- `:files` - 显示缓冲区列表

#### 窗口管理

- `:split` 或 `:sp` - 水平分割窗口
- `:vsplit` 或 `:vs` - 垂直分割窗口
- `:split {file}` - 水平分割窗口并打开文件
- `:vsplit {file}` - 垂直分割窗口并打开文件
- `:new` - 创建新的水平分割窗口
- `:vnew` - 创建新的垂直分割窗口
- `Ctrl+w s` - 水平分割当前窗口
- `Ctrl+w v` - 垂直分割当前窗口
- `Ctrl+w w` - 在窗口间循环切换
- `Ctrl+w h/j/k/l` - 移动到左/下/上/右窗口
- `Ctrl+w H/J/K/L` - 移动窗口到左/下/上/右侧
- `Ctrl+w =` - 使所有窗口等宽等高
- `Ctrl+w _` - 最大化当前窗口高度
- `Ctrl+w |` - 最大化当前窗口宽度
- `Ctrl+w +` - 增加窗口高度
- `Ctrl+w -` - 减少窗口高度
- `Ctrl+w >` - 增加窗口宽度
- `Ctrl+w <` - 减少窗口宽度
- `{height}Ctrl+w _` - 设置窗口高度
- `{width}Ctrl+w |` - 设置窗口宽度
- `Ctrl+w o` - 关闭其他所有窗口
- `Ctrl+w c` - 关闭当前窗口
- `Ctrl+w q` - 退出当前窗口
- `:windo {command}` - 对所有窗口执行命令

#### 标签页管理

- `:tabnew` 或 `:tabedit {file}` - 在新标签页中打开文件
- `:tabclose` 或 `:tabc` - 关闭当前标签页
- `:tabonly` 或 `:tabo` - 关闭所有其他标签页
- `:tabnext` 或 `:tabn` 或 `gt` - 切换到下一个标签页
- `:tabprevious` 或 `:tabp` 或 `gT` - 切换到上一个标签页
- `{number}gt` - 切换到指定编号的标签页
- `:tabmove {position}` - 移动当前标签页
- `:tabs` - 列出所有标签页和其中的窗口
- `:tabfirst` 或 `:tabf` - 切换到第一个标签页
- `:tablast` 或 `:tabl` - 切换到最后一个标签页
- `:tabdo {command}` - 对所有标签页执行命令

**提示**：使用 `:set hidden` 可以允许在有未保存改动的情况下切换缓冲区。

### 标记与跳转

#### 标记操作

- `m{a-zA-Z}` - 设置标记，小写字母为文件局部标记，大写字母为全局标记
- `` `{mark} `` - 跳转到标记的精确位置（行和列）
- `'{mark}` - 跳转到标记所在行的第一个非空白字符
- `` `. `` - 跳转到上次修改的位置
- `` `^ `` - 跳转到上次插入的位置
- `` `[ `` 或 `` `] `` - 跳转到上次修改或复制的文本的开始/结束位置
- `` `< `` 或 `` `> `` - 跳转到上次可视选择的开始/结束位置
- `:marks` - 列出所有标记
- `:delmarks {marks}` - 删除指定标记
- `:delmarks!` - 删除所有小写字母标记

#### 自动标记

- `''` - 跳转到上次跳转前的位置
- ``` ` - 跳转到上次跳转前的位置（精确位置）
- `'.` - 跳转到上次修改的行
- `'^` - 跳转到上次插入模式退出的位置
- `'[` 或 `']` - 跳转到上次修改或复制的文本的开始/结束行
- `'<` 或 `'>` - 跳转到上次可视选择的开始/结束行

#### 跳转列表

- `:jumps` - 显示跳转列表
- `Ctrl+o` - 跳回到较旧的位置
- `Ctrl+i` 或 `Tab` - 跳到较新的位置

#### 变更列表

- `:changes` - 显示变更列表
- `g;` - 跳到上一个变更位置
- `g,` - 跳到下一个变更位置

**提示**：标记非常适合在大文件中快速导航。全局标记（大写字母）可以在不同文件间跳转。

### 宏录制

- `q{a-z}` - 开始录制宏到寄存器 {a-z}
- `q` - 停止录制
- `@{a-z}` - 执行寄存器中的宏
- `@@` - 重复执行上次的宏
- `{count}@{a-z}` - 执行指定宏多次
- `:reg {a-z}` - 查看寄存器中的宏内容
- `:let @a='commands'` - 手动编辑宏内容

**提示**：宏是自动化重复任务的强大工具。录制时考虑光标的初始位置和最终位置，确保宏可以连续多次执行。

### 折叠

- `zf{motion}` - 创建折叠
- `zd` - 删除光标下的折叠
- `zD` - 递归删除光标下的折叠
- `zo` - 打开光标下的折叠
- `zO` - 递归打开光标下的折叠
- `zc` - 关闭光标下的折叠
- `zC` - 递归关闭光标下的折叠
- `za` - 切换光标下的折叠状态
- `zA` - 递归切换光标下折叠状态
- `zr` - 减少折叠级别（打开一级折叠）
- `zR` - 打开所有折叠
- `zm` - 增加折叠级别（关闭一级折叠）
- `zM` - 关闭所有折叠
- `zn` - 禁用折叠
- `zN` - 启用折叠
- `zi` - 切换折叠功能
- `:set foldmethod=manual|indent|syntax|marker|expr|diff` - 设置折叠方式

**提示**：在 `.vimrc` 中配置 `set foldmethod=indent` 可以根据缩进自动创建代码折叠。

### 外部命令集成

- `:!{command}` - 执行外部命令
- `:r !{command}` - 读取命令输出并插入到当前位置
- `!!{command}` - 用命令输出替换当前行
- `:{range}!{command}` - 将指定范围的文本通过命令过滤
- `v{motion}:!{command}` - 对选中文本执行命令并替换
- `:shell` 或 `:sh` - 临时进入 shell（退出后返回 Vim）
- `:read !{command}` - 插入命令输出到光标下方
- `:write !{command}` - 将当前文件内容作为命令的输入
- `Ctrl+z` - 挂起 Vim（在 Unix-like 系统中，使用 `fg` 返回）

**提示**：利用外部命令可以实现复杂的文本处理，如排序（`:!sort`）、格式化（`:!fmt`）等。

## 配置 Vim

### .vimrc 文件简介

`.vimrc` 是 Vim 的配置文件，位于用户主目录下：

- Linux/macOS: `~/.vimrc`
- Windows: `$HOME\_vimrc` 或 `$VIM\_vimrc`

通过编辑这个文件，你可以定制 Vim 的行为、外观和功能。若文件不存在，可以创建一个。

### 基本配置示例

下面是一个基础 `.vimrc` 配置示例：

```vim
" 基本设置
set nocompatible              " 不兼容 Vi 模式（启用 Vim 特性）
filetype plugin indent on     " 启用文件类型检测和相关插件
syntax on                     " 语法高亮

" 界面设置
set number                    " 显示行号
set relativenumber            " 显示相对行号
set ruler                     " 显示光标位置
set showcmd                   " 在状态栏显示命令
set showmode                  " 显示当前模式
set laststatus=2              " 总是显示状态栏
set wildmenu                  " 命令行补全菜单
set cursorline                " 高亮当前行
set scrolloff=3               " 光标上下保留 3 行可见
set sidescrolloff=5           " 光标左右保留 5 列可见

" 缩进设置
set autoindent                " 自动缩进
set smartindent               " 智能缩进
set tabstop=4                 " Tab 宽度为 4 空格
set softtabstop=4             " 编辑时 Tab 的宽度
set shiftwidth=4              " 自动缩进宽度
set expandtab                 " 将 Tab 转换为空格
set smarttab                  " 智能处理 Tab

" 搜索设置
set hlsearch                  " 高亮搜索结果
set incsearch                 " 增量搜索（边输入边搜索）
set ignorecase                " 搜索时忽略大小写
set smartcase                 " 如果搜索包含大写字母，不忽略大小写

" 编辑体验
set backspace=indent,eol,start " 允许退格键删除换行等
set whichwrap+=<,>,h,l        " 允许光标在行首尾时跨行移动
set wrap                      " 自动换行
set linebreak                 " 不在单词中间断行
set showmatch                 " 显示匹配的括号
set matchtime=2               " 匹配括号高亮时间（0.2 秒）
set hidden                    " 允许在有未保存的更改时切换缓冲区

" 文件处理
set autoread                  " 当文件在外部被修改时，自动重新读取
set nobackup                  " 不创建备份文件
set noswapfile                " 不创建交换文件
set fileformats=unix,dos,mac  " 文件格式优先级

" 性能相关
set lazyredraw                " 执行宏时不重绘屏幕
set ttyfast                   " 指示快速终端连接
set timeout timeoutlen=1000 ttimeoutlen=10 " 按键响应超时设置

" 编码设置
set encoding=utf-8            " 内部编码
set fileencoding=utf-8        " 写入文件时使用的编码
set fileencodings=utf-8,gb2312,gb18030,gbk,ucs-bom,cp936,latin1 " 文件编码探测列表

" 实用映射
" 使用空格作为 leader 键
let mapleader = " "

" jk 映射为 Esc
inoremap jk <Esc>

" 快速保存和退出
nnoremap <leader>w :w<CR>
nnoremap <leader>q :q<CR>
nnoremap <leader>wq :wq<CR>

" 在窗口间快速移动
nnoremap <C-h> <C-w>h
nnoremap <C-j> <C-w>j
nnoremap <C-k> <C-w>k
nnoremap <C-l> <C-w>l

" 重新加载 vimrc
nnoremap <leader>sv :source $MYVIMRC<CR>

" 清除搜索高亮
nnoremap <leader><space> :nohlsearch<CR>
```

### 常用设置详解

#### 基本行为设置

- `set nocompatible` - 使用 Vim 的增强功能，而非 Vi 兼容模式
- `filetype plugin indent on` - 启用针对不同文件类型的插件和缩进
- `syntax on` - 启用语法高亮
- `set backspace=indent,eol,start` - 使退格键可以删除缩进、行尾和插入点之前的字符

#### 显示与界面

- `set number` - 显示行号
- `set relativenumber` - 显示相对行号（当前行为 0，其他行显示与当前行的相对距离）
- `set ruler` - 在右下角显示光标位置信息
- `set laststatus=2` - 总是显示状态栏
- `set statusline=...` - 自定义状态栏格式
- `set wildmenu` - 在命令模式下，Tab 补全显示可能的匹配项菜单
- `set list` - 显示制表符和行尾等不可见字符
- `set listchars=tab:>-,trail:-` - 自定义不可见字符的显示方式
- `set background=dark` - 使用暗色背景（可选 light）
- `colorscheme name` - 设置配色方案

#### 缩进与格式化

- `set tabstop=4` - 设置 Tab 显示宽度
- `set shiftwidth=4` - 设置自动缩进的宽度
- `set expandtab` - 使用空格替代制表符
- `set autoindent` - 新行继承前一行的缩进
- `set smartindent` - 智能缩进，适用于编程
- `set textwidth=80` - 设置自动换行宽度
- `set formatoptions=...` - 控制文本格式化行为

#### 搜索

- `set hlsearch` - 高亮搜索结果
- `set incsearch` - 输入搜索模式时同时高亮部分匹配
- `set ignorecase` - 搜索时忽略大小写
- `set smartcase` - 如果搜索模式包含大写字母，则不忽略大小写

#### 性能与其他

- `set hidden` - 允许切换缓冲区时有未保存的更改
- `set history=1000` - 命令历史记录数量
- `set undofile` - 持久化撤销历史
- `set undodir=~/.vim/undodir` - 撤销文件存储目录
- `set noswapfile` - 不使用交换文件
- `set nobackup` - 不创建备份文件
- `set mouse=a` - 启用鼠标支持（所有模式）

### 键盘映射

Vim 允许自定义键盘映射，主要分为以下几种：

- `map` - 在所有模式下映射
- `nmap` - 在普通模式下映射
- `imap` - 在插入模式下映射
- `vmap` - 在可视模式下映射
- `cmap` - 在命令行模式下映射

为避免递归映射引起的问题，推荐使用非递归版本：

- `noremap` - 非递归映射（所有模式）
- `nnoremap` - 普通模式非递归映射
- `inoremap` - 插入模式非递归映射
- `vnoremap` - 可视模式非递归映射
- `cnoremap` - 命令行模式非递归映射

映射示例：

```vim
" 设置 leader 键（默认为反斜杠）
let mapleader = " "  " 使用空格作为 leader 键

" 普通模式下的映射
nnoremap <leader>s :w<CR>  " 使用 空格+s 保存文件
nnoremap <F5> :!python %<CR>  " F5 运行当前 Python 文件

" 插入模式下的映射
inoremap jk <Esc>  " 使用 jk 退出插入模式
inoremap <C-d> <Del>  " Ctrl+d 删除右侧字符

" 可视模式下的映射
vnoremap < <gv  " 缩进后保持选中状态
vnoremap > >gv  " 缩进后保持选中状态
```

### 自动命令

使用 `autocmd` 可以在特定事件发生时自动执行命令：

```vim
" 创建自动命令组，避免重复定义
augroup MyAutoCommands
    " 清除组内所有自动命令
    autocmd!

    " 编辑 Python 文件时设置缩进为 4 空格
    autocmd FileType python setlocal tabstop=4 shiftwidth=4 expandtab

    " 编辑 HTML/JS/CSS 文件时设置缩进为 2 空格
    autocmd FileType html,javascript,css setlocal tabstop=2 shiftwidth=2 expandtab

    " 打开文件时自动定位到上次编辑位置
    autocmd BufReadPost *
        \ if line("'\"") > 0 && line("'\"") <= line("$") |
        \   exe "normal! g`\"" |
        \ endif

    " 保存时自动删除行尾空白
    autocmd BufWritePre * %s/\s\+$//e

    " 保存时自动创建目录（如果不存在）
    autocmd BufWritePre * :call s:MkNonExDir(expand('<afile>'), +expand('<abuf>'))
augroup END

" 创建目录的函数
function! s:MkNonExDir(file, buf)
    if empty(getbufvar(a:buf, '&buftype')) && a:file!~#'\v^\w+\:\/'
        let dir=fnamemodify(a:file, ':h')
        if !isdirectory(dir)
            call mkdir(dir, 'p')
        endif
    endif
endfunction
```

### 函数定义

Vim 脚本允许定义函数：

```vim
" 定义一个简单的函数
function! MyFunction()
    echo "Hello, Vim!"
endfunction

" 带参数的函数
function! Greet(name)
    echo "Hello, " . a:name . "!"
endfunction

" 调用函数
call MyFunction()
call Greet("World")

" 映射到函数
nnoremap <leader>g :call Greet("Vim")<CR>
```

### 插件管理

#### 插件管理器

现代 Vim 配置通常使用插件管理器来管理插件。最常用的插件管理器有：

1. **Vim-plug**：轻量级、简洁且快速的插件管理器

```vim
" 安装 vim-plug（如果尚未安装）
if empty(glob('~/.vim/autoload/plug.vim'))
  silent !curl -fLo ~/.vim/autoload/plug.vim --create-dirs
    \ https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
  autocmd VimEnter * PlugInstall --sync | source $MYVIMRC
endif

" 使用 vim-plug 管理插件
call plug#begin('~/.vim/plugged')

" 插件列表
Plug 'tpope/vim-surround'                " 处理成对的引号、括号等
Plug 'preservim/nerdtree'                " 文件浏览器
Plug 'ctrlpvim/ctrlp.vim'                " 模糊文件查找
Plug 'vim-airline/vim-airline'           " 状态栏美化
Plug 'vim-airline/vim-airline-themes'    " 状态栏主题
Plug 'junegunn/fzf', { 'do': { -> fzf#install() } } " 模糊查找器
Plug 'junegunn/fzf.vim'                  " fzf 的 Vim 插件
Plug 'dense-analysis/ale'                " 异步语法检查
Plug 'tpope/vim-commentary'              " 快速注释
Plug 'jiangmiao/auto-pairs'              " 自动补全括号
Plug 'morhetz/gruvbox'                   " 配色方案
Plug 'neoclide/coc.nvim', {'branch': 'release'} " 代码补全框架
Plug 'sheerun/vim-polyglot'              " 语法高亮集合

" 根据开发语言增加特定插件
" Python
Plug 'vim-python/python-syntax', { 'for': 'python' }
Plug 'Vimjas/vim-python-pep8-indent', { 'for': 'python' }

" JavaScript
Plug 'pangloss/vim-javascript', { 'for': 'javascript' }
Plug 'maxmellon/vim-jsx-pretty', { 'for': ['javascript', 'jsx'] }

" 添加更多插件...

call plug#end()
```

安装插件：在 Vim 中运行 `:PlugInstall`
更新插件：`:PlugUpdate`
删除未使用的插件：`:PlugClean`

1. **其他插件管理器**

- **Vundle**：较老但仍很受欢迎的插件管理器
- **Pathogen**：简单的运行时路径管理器
- **Dein.vim**：性能优先的插件管理器
- **minpac**：利用 Vim 8 本地包支持的极简管理器

### 常用插件配置

以下是一些常用插件的基本配置：

#### NERDTree（文件浏览器）

```vim
" NERDTree 配置
" 映射切换 NERDTree 的快捷键
nnoremap <leader>n :NERDTreeToggle<CR>
" 显示隐藏文件
let NERDTreeShowHidden = 1
" 忽略特定文件
let NERDTreeIgnore = ['^\.git$', '\.pyc$', '\.DS_Store$']
" 关闭 NERDTree 帮助文本
let NERDTreeMinimalUI = 1
" 当 NERDTree 是最后一个窗口时自动关闭 Vim
autocmd BufEnter * if (winnr("$") == 1 && exists("b:NERDTree") && b:NERDTree.isTabTree()) | q | endif
" 当打开 Vim 且没有指定文件时自动打开 NERDTree
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif
```

#### fzf.vim（模糊查找）

```vim
" fzf.vim 配置
" 映射文件查找
nnoremap <leader>f :Files<CR>
" 映射缓冲区查找
nnoremap <leader>b :Buffers<CR>
" 映射文本内容查找
nnoremap <leader>r :Rg<CR>
" 使用 Ctrl-P 打开文件搜索
nnoremap <C-p> :Files<CR>
" 历史命令搜索
nnoremap <leader>h :History:<CR>
" 使用 Esc 键在 fzf 搜索界面中取消
let g:fzf_layout = { 'window': { 'width': 0.9, 'height': 0.6 } }
```

#### vim-airline（状态栏）

```vim
" vim-airline 配置
let g:airline_powerline_fonts = 1                " 使用 Powerline 字体
let g:airline#extensions#tabline#enabled = 1     " 显示缓冲区标签栏
let g:airline#extensions#tabline#formatter = 'unique_tail' " 文件名格式
let g:airline_theme = 'gruvbox'                  " 设置主题
let g:airline#extensions#branch#enabled = 1      " 显示 Git 分支
let g:airline#extensions#hunks#enabled = 1       " 显示 Git 改动
let g:airline#extensions#ale#enabled = 1         " 集成 ALE
```

#### ALE（语法检查）

```vim
" ALE 配置
let g:ale_sign_error = '✘'                    " 错误标志
let g:ale_sign_warning = '⚠'                  " 警告标志
let g:ale_lint_on_text_changed = 'never'      " 不是在文本改变时检查
let g:ale_lint_on_enter = 0                   " 不是在进入缓冲区时检查
let g:ale_lint_on_save = 1                    " 保存时检查
let g:ale_fix_on_save = 1                     " 保存时自动修复
" 设置特定语言的检查器和修复器
let g:ale_linters = {
\   'python': ['flake8', 'pylint'],
\   'javascript': ['eslint'],
\}
let g:ale_fixers = {
\   '*': ['remove_trailing_lines', 'trim_whitespace'],
\   'python': ['autopep8', 'isort'],
\   'javascript': ['prettier', 'eslint'],
\}
" 快捷键导航错误
nmap <silent> <C-k> <Plug>(ale_previous_wrap)
nmap <silent> <C-j> <Plug>(ale_next_wrap)
```

#### vim-surround（处理包围符号）

```vim
" vim-surround 无需额外配置，默认映射已够用
" 一些基本操作：
" cs"' - 将 "hello" 改为 'hello'
" ds" - 将 "hello" 改为 hello
" ysiw] - 将 hello 改为 [hello]
" yss) - 将整行包围在括号中
```

#### coc.nvim（代码补全）

```vim
" coc.nvim 配置
" 安装扩展
let g:coc_global_extensions = [
  \ 'coc-json',
  \ 'coc-python',
  \ 'coc-tsserver',
  \ 'coc-html',
  \ 'coc-css',
  \ 'coc-prettier',
  \ 'coc-eslint',
  \ ]

" TextEdit 可能会失败如果没有设置这个选项
set hidden

" 更短的更新时间提供更好的用户体验
set updatetime=300

" 不要传递消息到 |ins-completion-menu|
set shortmess+=c

" 使用 Tab 进行补全导航
inoremap <silent><expr> <TAB>
      \ pumvisible() ? "\<C-n>" :
      \ <SID>check_back_space() ? "\<TAB>" :
      \ coc#refresh()
inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"

function! s:check_back_space() abort
  let col = col('.') - 1
  return !col || getline('.')[col - 1]  =~# '\s'
endfunction

" 使用 <c-space> 触发补全
inoremap <silent><expr> <c-space> coc#refresh()

" 使用 <cr> 确认补全
inoremap <silent><expr> <cr> pumvisible() ? coc#_select_confirm()
                              \: "\<C-g>u\<CR>\<c-r>=coc#on_enter()\<CR>"

" 使用 `[g` 和 `]g` 导航诊断
nmap <silent> [g <Plug>(coc-diagnostic-prev)
nmap <silent> ]g <Plug>(coc-diagnostic-next)

" 代码导航
nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)

" 显示文档
nnoremap <silent> K :call <SID>show_documentation()<CR>

function! s:show_documentation()
  if (index(['vim','help'], &filetype) >= 0)
    execute 'h '.expand('<cword>')
  elseif (coc#rpc#ready())
    call CocActionAsync('doHover')
  else
    execute '!' . &keywordprg . " " . expand('<cword>')
  endif
endfunction

" 符号重命名
nmap <leader>rn <Plug>(coc-rename)

" 格式化选中代码
xmap <leader>f  <Plug>(coc-format-selected)
nmap <leader>f  <Plug>(coc-format-selected)
```

### 高级 Vim 配置策略

为了保持 `.vimrc` 文件的可读性和组织性，可以将配置拆分为多个文件：

```text
~/.vim/
  ├── vimrc             # 主配置文件，可以软链接到 ~/.vimrc
  ├── plugins.vim       # 插件定义和管理
  ├── mappings.vim      # 按键映射
  ├── settings.vim      # 基本设置
  ├── autocommands.vim  # 自动命令
  ├── functions.vim     # 自定义函数
  └── plugin/           # 插件特定配置
      ├── nerdtree.vim
      ├── fzf.vim
      ├── coc.vim
      └── ...
```

在主 `.vimrc` 文件中：

```vim
" 加载拆分的配置文件
source ~/.vim/settings.vim
source ~/.vim/plugins.vim
source ~/.vim/mappings.vim
source ~/.vim/autocommands.vim
source ~/.vim/functions.vim
```

这种方法使配置更模块化，更易于维护。

## 插件推荐

Vim 可以通过插件极大地增强功能。推荐使用 Vim-plug 等插件管理器：

### 必备插件

1. **NERDTree** - 文件树浏览器，提供直观的文件系统导航
2. **fzf.vim** - 模糊查找器，快速搜索文件、缓冲区、标记等
3. **vim-airline** - 状态栏美化，提供丰富的状态信息
4. **vim-surround** - 快速操作括号、引号等包围符号
5. **ALE** - 异步语法检查与自动修复

### 编码增强

1. **coc.nvim** - 代码补全框架，支持 LSP
2. **vim-commentary** - 快速注释/取消注释代码
3. **auto-pairs** - 自动补全括号、引号等
4. **vim-gitgutter** - 显示 Git diff 标记
5. **vim-fugitive** - Git 集成
6. **vim-polyglot** - 语法高亮包集合，支持数百种语言

### 外观美化

1. **gruvbox** - 流行的配色方案
2. **vim-devicons** - 添加文件类型图标
3. **indentLine** - 显示缩进线

### 导航与移动

1. **vim-easymotion** - 快速光标移动
2. **vim-sneak** - 两字符搜索导航
3. **vim-tmux-navigator** - 在 Vim 和 Tmux 分屏间无缝导航

### 特定语言支持

1. **vim-go** - Go 语言支持
2. **vim-python-pep8-indent** - Python PEP8 缩进支持
3. **rust.vim** - Rust 语言支持

## 学习路径

学习 Vim 可能看起来有些困难，但有了合适的学习路径，你可以循序渐进地掌握它。以下是推荐的学习步骤：

### 第一阶段：基础技能（1-2 周）

1. **完成 Vim 自带教程**：在终端输入 `vimtutor`，这是官方的交互式教程，大约需要 30 分钟。每天重复一次，直到熟悉基本操作。

2. **练习基本移动和编辑**：
    - 掌握 hjkl 移动
    - 学习基本编辑命令：i, a, x, dd, yy, p
    - 学习基本保存退出：:w, :q, :wq
    - 强制自己不使用箭头键，只用 hjkl

3. **建立日常使用习惯**：
    - 每天至少用 Vim 编辑文本 30 分钟
    - 创建一个专门用来练习的文本文件
    - 开始使用 Vim 编辑配置文件或简单的代码

4. **了解 Vim 模式**：
    - 普通模式、插入模式、可视模式、命令模式的基本切换
    - 养成频繁按 Esc 返回普通模式的习惯

### 第二阶段：提高效率（2-4 周）

1. **学习高效导航**：
    - 单词级移动：w, e, b
    - 行内导航：0, ^, $, f, t
    - 段落移动：{, }
    - 屏幕导航：H, M, L, Ctrl+f, Ctrl+b
    - 文件导航：gg, G, :{number}

2. **掌握文本编辑技巧**：
    - 文本对象：i", a(, it 等
    - 各种删除命令：diw, ci", da(
    - 重复上一次操作：.
    - 撤销和重做：u, Ctrl+r

3. **命令组合**：
    - 理解"动词+名词"结构（如 d3w 删除三个单词）
    - 学习常用组合如 ci", dip, yap
    - 创建一个你常用命令的备忘录

4. **开始使用寄存器和标记**：
    - 复制到命名寄存器："ayy
    - 从寄存器粘贴："ap
    - 设置和跳转到标记：ma, 'a

### 第三阶段：定制和高级功能（1-2 月）

1. **创建自己的 .vimrc 文件**：
    - 设置基本选项如 number, hlsearch, incsearch
    - 定义简单的映射简化常用操作
    - 参考其他人的配置，但确保理解每一行

2. **学习使用插件**：
    - 安装插件管理器如 Vim-plug
    - 从少量必要插件开始（如 NERDTree, fzf.vim）
    - 逐步添加提高效率的插件
    - 花时间学习每个插件的用法

3. **学习高级搜索和替换**：
    - 正则表达式搜索
    - 全局替换：:%s/old/new/g
    - 范围替换
    - 确认替换：:%s/old/new/gc

4. **掌握窗口和缓冲区管理**：
    - 分割窗口：:split, :vsplit
    - 窗口导航：Ctrl+w hjkl
    - 缓冲区操作：:ls, :b{number}

### 第四阶段：专业技能（持续学习）

1. **学习宏的使用**：
    - 录制宏：qa...q
    - 执行宏：@a
    - 在多行上应用宏

2. **探索高级功能**：
    - 折叠
    - 多文件编辑
    - 跳转列表
    - 自动命令
    - Vim 脚本基础

3. **特定领域优化**：
    - 针对你使用的编程语言配置 Vim
    - 设置特定文件类型的缩进、语法等
    - 集成构建系统、调试器

4. **持续学习和改进**：
    - 订阅 Vim 技巧博客或关注 Vim 社区
    - 找到优秀的 Vim 用户并学习他们的工作流
    - 定期审视自己的 Vim 使用习惯并优化
    - 贡献到 Vim 社区或分享你的技巧

### 实用学习方法

1. **每天学习 1-3 个新命令**：
    - 选择一个命令并刻意练习多次
    - 在实际工作中应用这个命令
    - 直到它成为肌肉记忆

2. **创建备忘录卡片**：
    - 将重要命令写在卡片上
    - 放在显眼的地方随时查阅
    - 定期复习

3. **设置学习目标**：
    - "本周我要熟练使用文本对象"
    - "下周我要熟悉宏的录制和使用"

4. **找一个 Vim 学习伙伴**：
    - 互相分享技巧
    - 共同解决问题
    - 定期交流进展

5. **解决实际问题**：
    - 用 Vim 完成日常编辑工作
    - 遇到问题时查找 Vim 解决方案
    - 记录解决过程

6. **坚持使用**：
    - 最重要的是持续使用 Vim
    - 度过初期的不适应期
    - 当习惯成自然后，效率会显著提高

## 有用的资源

### 官方资源

- **Vim 官方文档**: 在 Vim 中输入 `:help`
- **Vim 官方网站**: <https://www.vim.org/>
- **Vim 邮件列表**: <https://www.vim.org/maillist.php>

### 交互式学习工具

- **Vim Adventures**: <https://vim-adventures.com/> - 通过游戏学习 Vim
- **OpenVim**: <https://www.openvim.com/> - 交互式 Vim 教程
- **Vim Genius**: <http://www.vimgenius.com/> - 记忆 Vim 命令
- **PacVim**: <https://github.com/jmoon018/PacVim> - 基于终端的游戏

### 速查表和参考

- **Vim Cheat Sheet**: <https://vim.rtorr.com/> - 全面的速查表
- **Vim Awesome**: <https://vimawesome.com/> - Vim 插件目录
- **Vim Casts**: <http://vimcasts.org/> - 屏幕录像教程

### 书籍

- **Practical Vim** (Drew Neil): 非常推荐的实用技巧书
- **Modern Vim** (Drew Neil): Practical Vim 的续集，针对 Neovim 和 Vim 8
- **Learning the vi and Vim Editors** (O'Reilly): 全面的入门书籍
- **Vim 实用技巧** (Drew Neil 著，杨源译): Practical Vim 的中文版

### 博客和论坛

- **r/vim**: <https://www.reddit.com/r/vim/> - Reddit Vim 社区
- **Vim Tips Wiki**: <https://vim.fandom.com/wiki/Vim_Tips_Wiki>
- **UsevVim**: <https://medium.com/usevim> - Vim 技巧博客

### 配置示例

- **vim-sensible**: <https://github.com/tpope/vim-sensible> - 理性的默认设置
- **spf13-vim**: <https://github.com/spf13/spf13-vim> - 流行的 Vim 发行版
- **ThinkVim**: <https://github.com/hardcoreplayers/ThinkVim> - 现代化 Vim 配置

### 视频课程

- **Vim Masterclass**: <https://www.udemy.com/course/vim-commands-cheat-sheet/>
- **Mastering Vim**: <https://thoughtbot.com/upcase/mastering-vim>

## 常见问题

### 如果卡在 Vim 中怎么办？

最常见的解决方法是按 `Esc`（可能需要多按几次）然后输入 `:q!` 强制退出。

如果你发现自己处于一个未知模式：

- 如果底部显示 `-- INSERT --`，你在插入模式，按 `Esc` 返回普通模式
- 如果底部显示 `-- VISUAL --` 或 `-- VISUAL LINE --`，你在可视模式，按 `Esc` 返回普通模式
- 如果底部显示 `-- REPLACE --`，你在替换模式，按 `Esc` 返回普通模式
- 如果底部显示 `Recording @x`，你正在录制宏，按 `q` 停止录制

如果你不小心按了 `Ctrl+s` 导致终端冻结，可以按 `Ctrl+q` 恢复。

### 如何在不同的模式间切换？

随时按 `Esc` 可返回普通模式，这是切换到其他模式的起点。

从普通模式切换到其他模式：

- 按 `i`, `a`, `o` 等进入插入模式
- 按 `v`, `V`, `Ctrl+v` 进入可视模式
- 按 `:` 进入命令模式
- 按 `R` 进入替换模式

记住普通模式是 Vim 的"家"，经常返回普通模式是良好的习惯。

### 如何剪切/删除多行？

在普通模式下，`dd` 删除（并剪切）当前行，`5dd` 将删除（并剪切）5行。

其他方式：

- `dj` 删除当前行和下一行
- `d}` 删除到段落结尾
- `dG` 删除到文件结尾
- 在可视模式下，选择多行然后按 `d` 删除选中内容

### 如何复制文本到系统剪贴板？

如果 Vim 支持系统剪贴板（可以用 `:echo has('clipboard')` 检查，返回 1 表示支持），可以：

- `"+y` 复制到系统剪贴板
- `"+p` 从系统剪贴板粘贴

如果不支持系统剪贴板，可以：

1. 配置 Vim 支持系统剪贴板（安装带有 clipboard 特性的 Vim）
2. 使用终端多选（如 macOS 终端按住 Option 键选择）
3. 使用 `:w /tmp/temp.txt` 将内容写入临时文件，然后在其他应用打开

### 我编辑了文件却不能保存，提示"只读"怎么办？

如果文件是只读的，你可能会看到 `E45: 'readonly' option is set` 错误。

解决方法：

1. 如果你有权限覆盖文件，可以使用 `:w!` 强制保存
2. 或者使用 `:w new_filename` 保存为新文件
3. 对于系统文件，可能需要超级用户权限：`:w !sudo tee % > /dev/null`

### 如何搜索和替换特定文本？

基本搜索：

- `/pattern` 向下搜索
- `?pattern` 向上搜索
- `n` 下一个匹配
- `N` 上一个匹配

基本替换：

- `:%s/old/new/g` 替换整个文件中的所有匹配
- `:%s/old/new/gc` 替换整个文件中的所有匹配，但每次都询问
- `:10,20s/old/new/g` = 只替换第 10-20 行之间的匹配

### 如何处理中文等非 ASCII 字符？

确保 Vim 正确设置编码：

```vim
set encoding=utf-8
set fileencoding=utf-8
set fileencodings=utf-8,gb2312,gb18030,gbk,ucs-bom,cp936,latin1
```

如果遇到显示问题，检查终端是否支持 UTF-8 编码，字体是否支持中文字符。

### 如何配置语法高亮？

启用语法高亮：

```vim
syntax on
```

设置配色方案：

```vim
colorscheme desert  " 或其他你喜欢的配色方案
```

针对特定文件类型的设置：

```vim
autocmd FileType python set syntax=python
```

### 插件安装后不生效怎么办？

检查以下几点：

1. 插件管理器是否正确安装
2. 插件是否正确安装（如运行 `:PlugStatus` 查看状态）
3. 是否有冲突的键映射或设置
4. 是否启用了插件（有些插件需要手动启用）
5. Vim 版本是否符合插件要求

可以查看插件的错误信息：`:messages` 或者 `:echo errmsg`

### Vim 启动很慢，如何优化？

1. 使用 `vim --startuptime startup.log` 来分析启动时间
2. 检查 `.vimrc` 中耗时的设置
3. 延迟加载不常用插件
4. 删除或禁用不必要的插件
5. 使用更高效的插件管理器（如 vim-plug）
6. 清理 `~/.vim` 目录中的无用文件

### 如何记忆这么多 Vim 命令？

1. 不要试图一次记住所有命令
2. 理解 Vim 的命令语法结构（动词+名词）
3. 先掌握最常用的 20% 命令，它们能完成 80% 的工作
4. 创建个人的命令备忘录
5. 定期练习不常用但有用的命令
6. 理解命令的逻辑（如 `i` 是 insert，`a` 是 append）

### 我应该使用 Vim 还是 Neovim？

- **Vim** 是原始的、稳定的选择，几乎可在任何系统上运行
- **Neovim** 是 Vim 的现代化分支，有更好的默认设置、异步支持和更活跃的社区开发

对初学者来说，两者差异不大。Neovim 可能对新用户更友好，因为它有更好的默认设置，但学习曲线基本相同。

大多数 Vim 技能和配置可以直接应用到 Neovim，所以可以从一个开始，之后随时切换。

## 结语

学习 Vim 是一个持续的过程。一开始可能会感到困难，但随着时间的推移，编辑效率会大大提高。关键是持之以恒，并且实际应用到日常工作中。

不要期望一夜之间掌握所有内容。就像学习一种乐器一样，精通 Vim 需要时间和练习，但回报是巨大的。当你发现自己"思考"并通过 Vim 命令"说话"时，你就会体验到真正的编辑效率。

记住：在 Vim 中，你思考的不再是"如何输入文本"，而是"如何编辑文本"。这种思维方式的转变是掌握 Vim 的关键。

祝你在 Vim 的学习之旅中取得成功！
