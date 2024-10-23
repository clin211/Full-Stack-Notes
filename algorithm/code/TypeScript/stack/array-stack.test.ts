import ArrayStack from './arraystack'; // 确保路径正确

describe('ArrayStack', () => {
    let stack: ArrayStack<number>;

    beforeEach(() => {
        stack = new ArrayStack<number>();
    });

    test('should initialize empty stack', () => {
        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
    });

    test('should push items onto the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        expect(stack.isEmpty()).toBe(false);
        expect(stack.size()).toBe(3);
        expect(stack.top()).toBe(3);
    });

    test('should pop items from the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        expect(stack.pop()).toBe(3);
        expect(stack.size()).toBe(2);
        expect(stack.top()).toBe(2);
    });

    test('should return undefined when popping from an empty stack', () => {
        expect(stack.pop()).toBeUndefined();
    });

    test('should return the top item without removing it', () => {
        stack.push(1);
        stack.push(2);

        expect(stack.top()).toBe(2);
        expect(stack.size()).toBe(2); // Size should remain the same
    });

    test('should clear the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.clear();

        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
        expect(stack.top()).toBeUndefined();
    });

    test('should handle multiple push and pop operations', () => {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        expect(stack.pop()).toBe(3);
        expect(stack.pop()).toBe(2);
        expect(stack.pop()).toBe(1);
        expect(stack.isEmpty()).toBe(true);
    });
});