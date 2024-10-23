// 定义 IBrowserHistory 接口
interface IBrowserHistory {
    visit(url: string): void;
    back(): string;
    forward(): string;
    currentUrl(): string;
}

// ListNode 类表示双向链表中的一个节点
class ListNode<T> {
    url: string;
    prev: ListNode<T> | null;
    next: ListNode<T> | null;

    constructor(url: string) {
        this.url = url;
        this.prev = null;
        this.next = null;
    }
}

// BrowserHistory 类实现 IBrowserHistory 接口
export default class BrowserHistory implements IBrowserHistory {
    private current: ListNode<string> | null;

    constructor() {
        this.current = null;
    }

    // 访问新 URL
    visit(url: string) {
        const newNode = new ListNode<string>(url);
        if (this.current) {
            newNode.prev = this.current;
            this.current.next = newNode;
        }
        this.current = newNode;
    }

    // 后退到前一个 URL
    back() {
        if (!this.current || !this.current.prev) {
            return "没有之前的 URL";
        }
        this.current = this.current.prev;
        return this.current.url;
    }

    // 前进到下一个 URL
    forward() {
        if (!this.current || !this.current.next) {
            return "没有下一个 URL";
        }
        this.current = this.current.next;
        return this.current.url;
    }

    // 返回当前的 URL
    currentUrl() {
        if (!this.current) {
            return "没有当前 URL";
        }
        return this.current.url;
    }
}
