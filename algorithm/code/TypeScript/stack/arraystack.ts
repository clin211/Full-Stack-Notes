export default class ArrayStack<T> {
    private data: T[];
    private count: number;

    constructor() {
        this.data = [];
        this.count = -1;
    }

    push(item: T) {
        this.data.push(item);
        this.count++;
    }

    pop(): T | undefined {
        if (this.count < 0) {
            return undefined;
        }
        const value = this.data.pop();
        this.count--;
        return value;
    }

    top(): T | undefined {
        if (this.count < 0) {
            return undefined;
        }
        return this.data[this.count];
    }

    isEmpty(): boolean {
        return this.count < 0;
    }

    size(): number {
        return this.count + 1;
    }

    clear() {
        this.data = [];
        this.count = -1;
    }

    print() {
        console.log(this.data);
    }
}
