import ArrayQueue from './array-queue';

describe('ArrayQueue', () => {
    it('should create an empty queue', () => {
        const queue = new ArrayQueue<number>();
        expect(queue.getSize()).toBe(0);
        expect(queue.isEmpty()).toBe(true);
    });

    it('should enqueue and dequeue elements', () => {
        const queue = new ArrayQueue<number>();
        queue.enqueue(1);
        queue.enqueue(2);
        queue.enqueue(3);
        expect(queue.getSize()).toBe(3);
        expect(queue.getFront()).toBe(1);
        expect(queue.dequeue()).toBe(1);
        expect(queue.getSize()).toBe(2);
        expect(queue.getFront()).toBe(2);
    });

    it('should throw an error when dequeuing from an empty queue', () => {
        const queue = new ArrayQueue<number>();
        expect(() => queue.dequeue()).toThrow('queue is empty');
    });

    it('should throw an error when getting the front of an empty queue', () => {
        const queue = new ArrayQueue<number>();
        expect(() => queue.getFront()).toThrow(new Error('queue is empty'));
    });

    it('should throw an error when enqueuing to a full queue', () => {
        const queue = new ArrayQueue<number>(2);
        queue.enqueue(1);
        queue.enqueue(2);
        expect(() => queue.enqueue(3)).toThrow('queue is full');
    });

    it('should correctly check if the queue is full', () => {
        const queue = new ArrayQueue<number>(2);
        queue.enqueue(1);
        expect(queue.isFull()).toBe(false);
        queue.enqueue(2);
        expect(queue.isFull()).toBe(true);
    });
});