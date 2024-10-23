import LinkedListStack from './linkedliststack';

// 定义一个简单的类型用于测试
type TestDataType = number;

describe('LinkedListStack', () => {
    let stack: LinkedListStack<TestDataType>;

    beforeEach(() => {
        stack = new LinkedListStack<TestDataType>();
    });

    test('should be empty on initialization', () => {
        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
    });

    test('should push elements onto the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        expect(stack.isEmpty()).toBe(false);
        expect(stack.size()).toBe(3);
        expect(stack.top()).toBe(3);
    });

    test('should pop elements from the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        expect(stack.pop()).toBe(3);
        expect(stack.size()).toBe(2);
        expect(stack.top()).toBe(2);
    });

    test('should return null when popping from an empty stack', () => {
        expect(stack.pop()).toBeNull();
    });

    test('should return the top element without removing it', () => {
        stack.push(1);
        expect(stack.top()).toBe(1);
        expect(stack.size()).toBe(1);
    });

    test('should clear the stack', () => {
        stack.push(1);
        stack.push(2);
        stack.clear();

        expect(stack.isEmpty()).toBe(true);
        expect(stack.size()).toBe(0);
    });

    test('should print the stack elements', () => {
        const consoleSpy = jest.spyOn(console, 'log');
        stack.push(1);
        stack.push(2);
        stack.push(3);

        stack.print();

        expect(consoleSpy).toHaveBeenCalledWith(3);
        expect(consoleSpy).toHaveBeenCalledWith(2);
        expect(consoleSpy).toHaveBeenCalledWith(1);

        consoleSpy.mockRestore(); // 清理 spy
    });
});