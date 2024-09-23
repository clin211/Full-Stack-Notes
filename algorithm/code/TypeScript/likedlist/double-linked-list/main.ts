interface List<T> {
    insertToHead(value: T): void
    findByValue(value: T): any
    findByIndex(index: number): any
    insertToIndex(value: T, index: number): void
    remove(value: T): boolean
    insertToHead(value: T): void
    insertToTail(value: T): void
    toString(): string
}

interface ListNode<T> {
    data: T
    prev: ListNode<T> | null
    next: ListNode<T> | null
}

class LinkedListNode<T> implements ListNode<T> {
    data: T
    prev: ListNode<T> | null = null
    next: ListNode<T> | null = null

    constructor(props: ListNode<T>) {
        this.data = props.data;
        this.prev = props.prev;
        this.next = props.next;
    }
}

class LinkedList<T> implements List<T> {
    constructor() { }
    findByIndex(index: number) {

    }
    findByValue(value: T) {

    }
    insertToHead(value: T): void
    insertToHead(value: T): void
    insertToHead(value: unknown): void {

    }

    insertToTail(value: T): void {

    }
    insertToIndex(value: T, index: number): void {

    }
    remove(value: T): boolean {
        return true
    }
    toString(): string {
        return ''
    }
}