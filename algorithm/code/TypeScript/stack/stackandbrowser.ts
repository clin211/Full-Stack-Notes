interface IStack<T> {
    push(item: T): void;
    pop(): T;
}

class Stack<T> implements IStack<T> {
    private data: T[] = [];
    private length: number = 0;

    push(item: T): void {

    }

    pop(): T {
        return null
    }
}