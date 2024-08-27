/**
 * 实现一个最大堆
 * initDataArray ---> 初始化数据
 * maxSize ---> 堆的最大容量
 */
class MaxHeap {
    constructor(initDataArray, maxSize = 9999) {
        let arr = initDataArray || [];

        // 获取初始化数据的总数
        this.currSize = arr.length;

        this.maxSize = maxSize
        //创建一个长度为arr长度的数组
        this.heap = new Array(arr.length);
        this.init()
    }

    // 初始化堆结构
    init () {
        // 将初始化数据插入到数组中
        for (let i = 0; i < this.currSize; i++) {
            this.heap[i] = arr[i];
        }
        //如果initDataArray非空，那么需要对数据进行堆化。若initDataArray为空，此步可以省略
        let currPos = Math.floor((this.currSize - 2) / 2);//获取最后一个分支节点的父节点索引
        while (currPos >= 0) {
            //局部自上向下下滑调整
            this.upToDown(currPos, this.currSize - 1);
            //调整下一个分支节点
            currPos--;
        }
    }

    // 得到栈结构数组
    getHeap () {
        return this.heap
    }
    // 由上至下堆化数据
    upToDown (start, m) {
        //父节点
        let parentIndex = start,
            //左子节点
            maxChildIndex = parentIndex * 2 + 1;

        while (maxChildIndex <= m) {
            // 如果右节点的值 > 左节点的值
            if (maxChildIndex < m && this.heap[maxChildIndex] < this.heap[maxChildIndex + 1]) {
                // 将交换索引换成右节点的索引
                maxChildIndex = maxChildIndex + 1;
            }
            // 如果父节点的值大于子节点值，那么不需要交换
            if (this.heap[parentIndex] >= this.heap[maxChildIndex]) {
                break;
            } else {
                // 否则, 交换父子节点的值
                let temp = this.heap[parentIndex];
                this.heap[parentIndex] = this.heap[maxChildIndex];
                this.heap[maxChildIndex] = temp;
                parentIndex = maxChildIndex;
                maxChildIndex = maxChildIndex * 2 + 1
            }
        }
    }

    /**
     * 插入一个数据 并返回插入是否成功
     */
    insert (data) {
        // 如果当前大小等于最大容量，返回插入失败
        if (this.currSize === this.maxSize) {
            return false
        }
        // 将值放入数组中的最后
        this.heap[this.currSize] = data;

        // 由下至上堆化数据
        this.downToUp(this.currSize);
        this.currSize++;
        return true;
    };

    // 由下至上堆化数据
    downToUp (start) {
        let childIndex = start;   //当前叶节点
        let parentIndex = Math.floor((childIndex - 1) / 2); //父节点 (k - 1) / 2

        while (childIndex > 0) {
            //如果大就不交换
            if (this.heap[parentIndex] >= this.heap[childIndex]) {
                break;
            } else {
                // 否则, 交换父子节点的值
                let temp = this.heap[parentIndex];
                this.heap[parentIndex] = this.heap[childIndex];
                this.heap[childIndex] = temp;
                childIndex = parentIndex;
                parentIndex = Math.floor((parentIndex - 1) / 2);
            }
        }
    }

    /**
     * 移除根元素，并返回根元素数据
     */
    removeRoot () {
        if (this.currSize <= 0) {
            return null;
        }
        let maxValue = this.heap[0];

        // 将数组的最后一个设置为根元素
        this.heap[0] = this.heap[this.currSize];
        this.currSize--;
        // 重新堆化数据
        this.upToDown(0, this.currSize - 1);
        return maxValue;
    };

}

// ---------------------------------------测试代码分割线--------------------------------
//初始化一个堆结构
const maxHeap = new MaxHeap();

//向堆结构中添加数据
const initArray = [1, 2, 3, 4, 5];
for (let i = 0; i < 5; i++) {
    maxHeap.insert(initArray[i])
}
// 得到堆数据
console.log(maxHeap.getHeap()) // [5, 4, 2, 1, 3]

console.log('max1', maxHeap.removeRoot()); // 得到5

console.log('max2', maxHeap.removeRoot()); // 得到4

console.log('max3', maxHeap.removeRoot()); // 得到3