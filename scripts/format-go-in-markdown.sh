#!/bin/bash
# 格式化 Markdown 文件中的 Go 代码块

# 设置颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 计数器
total_files=0
total_blocks=0
total_formatted=0
total_fixed_indent=0
total_skipped=0

# 查找所有包含 Go 代码块的 Markdown 文件
echo -e "${YELLOW}正在扫描 Markdown 文件中的 Go 代码块...${NC}"

# 使用 grep 查找包含 ```go 的文件
files=$(grep -rl '^\```go$' --include="*.md" . 2>/dev/null)

if [ -z "$files" ]; then
    echo -e "${YELLOW}未找到包含 Go 代码块的 Markdown 文件${NC}"
    exit 0
fi

# 创建临时文件
temp_file=$(mktemp)
temp_code=$(mktemp)
trap "rm -f $temp_file $temp_code" EXIT

# 辅助函数：检查代码块是否包含空格缩进
has_spaces_indent() {
    local content="$1"
    # 使用 sed 检测是否有4个空格或更多空格的缩进
    echo "$content" | sed -n 's/^$/@/p' | grep -q '@'
    # 检查是否有行以4个空格开头（不是tab）
    echo "$content" | grep -qE '^    '
    return $?
}

# 处理每个文件
while IFS= read -r file; do
    ((total_files++))
    echo -e "${GREEN}处理文件: $file${NC}"

    # 读取文件内容
    content=$(cat "$file")

    # 使用 awk 处理代码块
    # 状态: 0=在代码块外, 1=在go代码块内
    in_go_block=0
    block_content=""
    block_start=0
    line_num=0
    file_changed=0
    has_syntax_error=0

    # 清空临时文件
    > "$temp_file"

    while IFS= read -r line; do
        ((line_num++))

        # 检测 ```go 开始标记
        if [[ "$line" =~ ^\`\`\`go$ ]]; then
            in_go_block=1
            block_start=$line_num
            block_content=""
            echo "$line" >> "$temp_file"
            continue
        fi

        # 检测 ``` 结束标记
        if [[ "$line" =~ ^\`\`\`$ ]] && [ $in_go_block -eq 1 ]; then
            in_go_block=0
            ((total_blocks++))
            has_syntax_error=0

            # 将代码块写入临时文件用于测试
            echo "$block_content" > "$temp_code"

            # 尝试格式化代码块
            if gofmt -l "$temp_code" >/dev/null 2>&1; then
                # 代码可以格式化
                formatted=$(gofmt "$temp_code" 2>/dev/null)

                if [ "$block_content" != "$formatted" ]; then
                    ((total_formatted++))
                    file_changed=1
                    echo -e "  ${GREEN}✓ 格式化第 $block_start 行的代码块${NC}"
                fi
                # 写入格式化后的代码
                echo "$formatted" >> "$temp_file"
            else
                # 代码有语法错误，尝试修复缩进
                # 检查是否使用了空格缩进（检测以4个空格开头但不是tab+空格的行）
                if echo "$block_content" | grep -qE '^    '; then
                    # 将空格缩进转换为tab缩进
                    # 处理不同层级的缩进
                    fixed=$(echo "$block_content" | sed -E 's/^    /\t/g; s/^\t    /\t\t/g; s/^\t\t    /\t\t\t/g; s/^\t\t\t    /\t\t\t\t/g; s/^\t\t\t\t    /\t\t\t\t\t/g')

                    if [ "$block_content" != "$fixed" ]; then
                        ((total_fixed_indent++))
                        file_changed=1
                        echo -e "  ${BLUE}⊙ 修复第 $block_start 行的代码块缩进${NC}"
                    fi
                    echo "$fixed" >> "$temp_file"
                else
                    # 保持原样
                    ((total_skipped++))
                    echo -e "  ${YELLOW}○ 跳过第 $block_start 行的代码块（语法错误）${NC}"
                    echo "$block_content" >> "$temp_file"
                fi
            fi

            echo "$line" >> "$temp_file"
            block_content=""
            continue
        fi

        if [ $in_go_block -eq 1 ]; then
            # 在 go 代码块内，累积内容
            if [ -z "$block_content" ]; then
                block_content="$line"
            else
                block_content="$block_content"$'\n'"$line"
            fi
        else
            # 在代码块外，直接输出
            echo "$line" >> "$temp_file"
        fi
    done < <(printf '%s\n' "$content")

    # 如果文件有变化，则更新
    if [ $file_changed -eq 1 ]; then
        cp "$temp_file" "$file"
        echo -e "  ${GREEN}✓ 已更新${NC}"
    fi

done <<< "$files"

# 输出统计
echo ""
echo -e "${GREEN}========== 格式化完成 ==========${NC}"
echo -e "扫描文件: $total_files"
echo -e "Go 代码块: $total_blocks"
echo -e "已格式化: $total_formatted"
echo -e "已修复缩进: $total_fixed_indent"
echo -e "已跳过: $total_skipped"
echo -e "${GREEN}================================${NC}"
