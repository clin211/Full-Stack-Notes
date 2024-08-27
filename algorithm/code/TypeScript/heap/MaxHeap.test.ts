import MaxHeap from './MaxHeap'; // 确保路径正确
import 'jest';

describe('MaxHeap', () => {
    it('使用数组初始化', () => {
        const initArray = [5, 3, 8, 4, 2];
        const maxHeap = new MaxHeap<number>(initArray);
        expect(maxHeap.getHeap()).toEqual([8, 5, 4, 3, 2]);
    });

    it('插入一个新的元素', () => {
        const maxHeap = new MaxHeap<number>();
        maxHeap.insert(5);
        maxHeap.insert(10);
        expect(maxHeap.getHeap()).toEqual([10, 5]);
    });

    it('删除指定根节点', () => {
        const maxHeap = new MaxHeap<number>([5, 3, 8, 4, 2]);
        expect(maxHeap.removeRoot()).toBe(8);
        expect(maxHeap.getHeap()).toEqual([5, 4, 3, 2]);
    });

    it('在多次插入后，堆化', () => {
        const maxHeap = new MaxHeap<number>();
        const elements = [20, 14, 7, 85, 3, 68, 5];
        elements.forEach(element => maxHeap.insert(element));
        expect(maxHeap.getHeap()).toEqual([85, 20, 14, 7, 3, 68, 5]);
    });
});