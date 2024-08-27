class MaxHeap<T> {
    private heap: T[];
    private currSize: number;
    private maxSize: number;

    constructor(initDataArray: T[] = [], maxSize: number = 9999) {
        // 获取初始化数据的总数
        this.currSize = initDataArray.length;
        this.maxSize = maxSize;

        // 创建一个长度为初始化数据长度的数组
        this.heap = new Array(this.currSize);
        this.init(initDataArray);
    }

    // 初始化堆结构
    private init(initDataArray: T[]): void {
        // 将初始化数据插入到数组中
        for (let i = 0; i < this.currSize; i++) {
            this.heap[i] = initDataArray[i];
        }

        // 获取最后一个分支节点的父节点索引
        let currPos = Math.floor((this.currSize - 2) / 2);
        while (currPos >= 0) {
            // 局部自上向下下滑调整
            this.upToDown(currPos, this.currSize - 1);
            // 调整下一个分支节点
            currPos--;
        }
    }

    // 得到堆结构数组
    public getHeap(): T[] {
        return this.heap;
    }

    // 由上至下堆化
    private upToDown(start: number, m: number) {
        // 父节点
        let parentIndex = start;

        // 左子节点
        let maxChildIndex = parentIndex * 2 + 1;
        while (maxChildIndex <= m) {
            // 右节点的值 > 左节点的值
            if (maxChildIndex < m && this.heap[maxChildIndex] < this.heap[maxChildIndex + 1]) {
                // 将交换索引换成右节点的索引
                maxChildIndex++;
            }

            // 如果父节点的值大于子节点值，那么不需要交换
            if (this.heap[parentIndex] >= this.heap[maxChildIndex]) {
                break;
            } else {
                // 交换父子节点的值
                let temp = this.heap[parentIndex];
                this.heap[parentIndex] = this.heap[maxChildIndex];
                this.heap[maxChildIndex] = temp;
                parentIndex = maxChildIndex;
                maxChildIndex = maxChildIndex * 2 + 1;
            }
        }
    }

    // 插入一个数据，返回插入是否成功
    public insert(data: T) {
        // 如果当前大小等于最大容量，返回插入失败
        if (this.currSize === this.maxSize) {
            return false;
        }
        // 将值放入数组中的最后
        this.heap[this.currSize] = data;

        // 由下至上堆化
        this.downToUp(this.currSize);
        this.currSize++;
        return true;
    }

    // 由上至下堆化
    private downToUp(start: number) {
        let childIndex = start; // 当前叶节点的索引
        let parentIndex = Math.floor((childIndex - 1) / 2); // 当前叶节点的父节点的索引 (k-1)/2
        while (childIndex > 0) {
            // 如果大就不交换
            if (this.heap[parentIndex] >= this.heap[childIndex]) {
                break;
            } else {
                // 交换父子节点的值
                let temp = this.heap[parentIndex];
                this.heap[parentIndex] = this.heap[childIndex];
                this.heap[childIndex] = temp;
                childIndex = parentIndex;
                parentIndex = Math.floor((parentIndex - 1) / 2);
            }
        }
    }

    // 移除根节点，并返回根元素数据
    public removeRoot(): T | null {
        // 如果当前数据大小等于0，返回null
        if (this.currSize <= 0) {
            return null;
        }
        // 默认将数组的第一个设置为根元素
        let maxValue = this.heap[0];

        // 将数组的最后一个设置为根元素
        this.heap[0] = this.heap[this.currSize - 1];
        this.currSize--;

        // 重新堆化数据
        this.upToDown(0, this.currSize - 1);
        return maxValue;
    }
}

export default MaxHeap;
