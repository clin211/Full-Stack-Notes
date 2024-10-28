interface IArrayQueue<T> {
    enqueue(item: T): void
    dequeue(): T
    getFront(): T
    isEmpty(): boolean
    isFull(): boolean
    getSize(): number
}

export default class ArrayQueue<T> implements IArrayQueue<T> {
    private data: T[]
    private capacity: number
    private front: number
    private rear: number
    private size: number

    constructor(capacity: number = 10) {
        this.data = new Array(capacity)
        this.capacity = capacity
        this.front = 0
        this.rear = -1
        this.size = 0
    }

    enqueue(item: T) {
        if (this.isFull()) {
            throw new Error('queue is full')
        }
        this.rear = (this.rear + 1) % this.capacity
        this.data[this.rear] = item
        this.size++
    }

    dequeue(): T {
        if (this.isEmpty()) {
            throw new Error('queue is empty')
        }
        const data = this.data[this.front]
        this.front = (this.front + 1) % this.capacity
        this.size--
        return data
    }

    getFront() {
        if (this.isEmpty()) {
            throw new Error('queue is empty')
        }
        return this.data[this.front]
    }

    isFull() {
        return this.size === this.capacity
    }

    isEmpty() {
        return this.size === 0
    }

    getSize() {
        return this.size
    }
}