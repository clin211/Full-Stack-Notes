class LinkedNode<T> {
    data: T
    next: LinkedNode<T> | null
    constructor(v: T, next: LinkedNode<T> | null) {
        this.data = v
        this.next = next
    }
}

interface LinkedListStack<T> {
    push(v: T): void
    pop(): T | null
    top(): T | null
    isEmpty(): boolean
    size(): number
    clear(): void
    print(): void
}

class LinkedListStack<T> implements LinkedListStack<T> {
    private head: LinkedNode<T> | null;
    private length: number;

    constructor() {
        // 初始化
        this.head = null;
        this.length = 0;
    }

    // 入栈
    push(v: T): void {
        const node = new LinkedNode(v, this.head);
        this.head = node;
        this.length++;
    }

    // 出栈
    pop(): T | null {
        if (this.isEmpty()) {
            return null;
        }
        const value = this.head?.data || null;
        this.head = this.head?.next || null;
        this.length--;
        return value;
    }

    // 获取栈顶元素
    top(): T | null {
        if (this.isEmpty()) {
            return null;
        }
        return this.head?.data || null;
    }

    // 判断栈是否为空
    isEmpty(): boolean {
        return this.length === 0;
    }

    // 获取栈长度
    size(): number {
        return this.length;
    }

    // 清空栈
    clear() {
        this.head = null;
        this.length = 0;
    }

    // 打印栈
    print() {
        let current = this.head;
        while (current) {
            console.log(current.data);
            current = current.next;
        }
    }
}

export default LinkedListStack