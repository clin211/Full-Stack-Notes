# 你必须知道的 45 个强大的 JavaScript 代码段

> 有关 JavaScript 中的数组、对象、函数、字符串和数字操作的有用代码片段！

今天，我总结了 45 个强大的 JavaScript 单行代码（代码片段），这些代码在编码时经常需要用到。我根据它们的行为将它们分为多个部分。

这些 JavaScript 单行代码包括数组、对象、函数、字符串和数字操作。我尽量让它简单一点，让我们开始吧！

## 数组操作

### 1. 展平数组

```js
const flat = (arr) => arr.flat(Infinity)
```

示例：

```js
const arr1 = [1, 2, [3, 4, [5, 6]], 7, [8, [9]]]
const arr2 = [[1], [2, [3, [4]]], [5, [6, [7]]]]
const arr3 = [
    [[1], [2]],
    [[3], [4]]
]

console.log(flat(arr1)) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
console.log(flat(arr2)) // [1, 2, 3, 4, 5, 6, 7]
console.log(flat(arr3)) // [1, 2, 3, 4]
```

### 2.从数组中删除重复的元素

```js
const unique = arr => […new Set(arr)];
```

示例：

```js
const unique = (arr) => [...new Set(arr)]

const arr1 = [1, 2, 2, 3, 4, 4, 5]
const arr2 = ['apple', 'banana', 'apple', 'orange']
const arr3 = [true, false, true, false]

console.log(unique(arr1)) // [1, 2, 3, 4, 5]
console.log(unique(arr2)) // ['apple', 'banana', 'orange']
console.log(unique(arr3)) // [true, false]
```

### 3.从数组中获取最后一个元素

```js
const last = (arr) => arr[arr.length - 1]
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = ['apple', 'banana', 'orange']
const arr3 = [true, false, true]

console.log(last(arr1)) // 5
console.log(last(arr2)) // 'orange'
console.log(last(arr3)) // true
```

### 4.打乱数组

```js
const shuffle = (arr) => arr.sort(() => Math.random() - 0.5)
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = ['apple', 'banana', 'orange']
const arr3 = [true, false, true]

console.log(shuffle(arr1)) // [3, 5, 2, 1, 4] (随机顺序)
console.log(shuffle(arr2)) // ['orange', 'apple', 'banana'] (随机顺序)
console.log(shuffle(arr3)) // [false, true, true] (随机顺序)
```

### 5.数组总和

```js
const sum = (arr) => arr.reduce((a, b) => a + b, 0)
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [10, 20, 30]
const arr3 = [5, 5, 5]

console.log(sum(arr1)) // 15
console.log(sum(arr2)) // 60
console.log(sum(arr3)) // 15
```

### 6.获取平均值

```js
const avg = (arr) => arr.reduce((a, b) => a + b, 0) / arr.length
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [10, 20, 30]
const arr3 = [5, 5, 5]

console.log(avg(arr1)) // 3
console.log(avg(arr2)) // 20
console.log(avg(arr3)) // 5
```

### 7.获取数组中最大值

```js
const max = (arr) => Math.max(...arr)
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [10, 20, 30]
const arr3 = [5, 5, 5]

console.log(max(arr1)) // 5
console.log(max(arr2)) // 30
console.log(max(arr3)) // 5
```

### 8.数组中的最小数量

```js
const min = (arr) => Math.min(...arr)
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [10, 20, 30]
const arr3 = [5, 5, 5]

console.log(min(arr1)) // 1
console.log(min(arr2)) // 10
console.log(min(arr3)) // 5
```

### 9.数组交集

```js
const intersect = (a, b) => a.filter((v) => b.includes(v))
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [4, 5, 6, 7, 8]
const arr3 = [1, 2, 3]
const arr4 = [3, 4, 5]

console.log(intersect(arr1, arr2)) // [4, 5]
console.log(intersect(arr1, arr3)) // [1, 2, 3]
console.log(intersect(arr3, arr4)) // [3]
```

### 10.对比数组

```js
const diff = (a, b) => a.filter((v) => !b.includes(v))
```

示例：

```js
const arr1 = [1, 2, 3, 4, 5]
const arr2 = [4, 5, 6, 7, 8]
const arr3 = [1, 2, 3]
const arr4 = [3, 4, 5]

console.log(diff(arr1, arr2)) // [1, 2, 3]
console.log(diff(arr1, arr3)) // [4, 5]
console.log(diff(arr3, arr4)) // [1, 2]
```

## 对象操作

### 11. 合并对象

```js
const merge = (a, b) => ({ …a, …b });
```

示例：

```js
const obj1 = { a: 1, b: 2 }
const obj2 = { c: 3, d: 4 }
const obj3 = { e: 5, f: 6 }
const obj4 = { g: 7, h: 8 }

console.log(merge(obj1, obj2)) // { a: 1, b: 2, c: 3, d: 4 }
console.log(merge(obj1, obj3)) // { a: 1, b: 2, e: 5, f: 6 }
console.log(merge(obj3, obj4)) // { e: 5, f: 6, g: 7, h: 8 }
```

### 12. 浅拷贝对象

```js
const clone = obj => ({ …obj });
```

示例：

```js
const obj1 = { a: 1, b: 2 }
const obj2 = { c: 3, d: 4, e: { f: 5 } }

const clonedObj1 = clone(obj1)
const clonedObj2 = clone(obj2)

console.log(clonedObj1) // { a: 1, b: 2 }
console.log(clonedObj2) // { c: 3, d: 4, e: { f: 5 } }

// 修改原始对象
obj1.a = 10
obj2.e.f = 20

console.log(clonedObj1) // { a: 1, b: 2 }
console.log(clonedObj2) // { c: 3, d: 4, e: { f: 20 } }
```

### 13. 深拷贝对象

```js
const deepClone = (obj) => JSON.parse(JSON.stringify(obj))
```

示例：

```js
const obj1 = { a: 1, b: 2 }
const obj2 = { c: 3, d: 4, e: { f: 5 } }

const clonedObj1 = deepClone(obj1)
const clonedObj2 = deepClone(obj2)

console.log(clonedObj1) // { a: 1, b: 2 }
console.log(clonedObj2) // { c: 3, d: 4, e: { f: 5 } }

// 修改原始对象
obj1.a = 10
obj2.e.f = 20

console.log(clonedObj1) // { a: 1, b: 2 }
console.log(clonedObj2) // { c: 3, d: 4, e: { f: 5 } }
```

### 14.将对象的 key 组成一个数组

```js
const keys = (obj) => Object.keys(obj)
```

示例：

```js
const obj1 = { a: 1, b: 2, c: 3 }
const obj2 = { name: 'John', age: 30, city: 'New York' }

console.log(keys(obj1)) // ['a', 'b', 'c']
console.log(keys(obj2)) // ['name', 'age', 'city']
```

### 15.将对象的 Value 组成一个数组

```js
const values = (obj) => Object.values(obj)
```

示例：

```js
const obj1 = { a: 1, b: 2, c: 3 }
const obj2 = { name: 'John', age: 30, city: 'New York' }

console.log(values(obj1)) // [1, 2, 3]
console.log(values(obj2)) // ['John', 30, 'New York']
```

### 16.将对象的 key value 组成一个数组

```js
const entries = (obj) => Object.entries(obj)
```

示例：

```js
const obj1 = { a: 1, b: 2, c: 3 }
const obj2 = { name: 'John', age: 30, city: 'New York' }

console.log(entries(obj1)) // [['a', 1], ['b', 2], ['c', 3]]
console.log(entries(obj2)) // [['name', 'John'], ['age', 30], ['city', 'New York']]
```

### 17.将键值对列表转成一个对象

```js
const fromEntries = (arr) => Object.fromEntries(arr)
```

示例：

```js
const arr1 = [
    ['a', 1],
    ['b', 2],
    ['c', 3]
]
const arr2 = [
    ['name', 'John'],
    ['age', 30],
    ['city', 'New York']
]

console.log(fromEntries(arr1)) // { a: 1, b: 2, c: 3 }
console.log(fromEntries(arr2)) // { name: 'John', age: 30, city: 'New York' }
```

### 18. 检查对象是否为空

```js
const isEmpty = (obj) => Object.keys(obj).length === 0
```

示例：

```js
const obj1 = { a: 1, b: 2 }
const obj2 = {}
const obj3 = { name: 'John' }

console.log(isEmpty(obj1)) // false
console.log(isEmpty(obj2)) // true
console.log(isEmpty(obj3)) // false
```

### 19.重命名对象的 key

```js
const renameKey = (obj, oldKey, newKey) => ({ …obj, [newKey]: obj[oldKey], [oldKey]: undefined });
```

示例：

```js
const obj1 = { a: 1, b: 2 }
const obj2 = { name: 'John', age: 30 }

console.log(renameKey(obj1, 'a', 'x')) // { b: 2, x: 1, a: undefined }
console.log(renameKey(obj2, 'name', 'username')) // { age: 30, username: 'John', name: undefined }
```

### 20.过滤对象的 key

```js
const filterKeys = (obj, keys) => keys.reduce((o, k) => (obj.hasOwnProperty(k) ? { ...o, [k]: obj[k] } : o), {})
```

示例：

```js
const obj1 = { a: 1, b: 2, c: 3, d: 4 }
const obj2 = { name: 'John', age: 30, city: 'New York' }

const filteredObj1 = filterKeys(obj1, ['a', 'c', 'e'])
const filteredObj2 = filterKeys(obj2, ['name', 'city', 'country'])

console.log(filteredObj1) // { a: 1, c: 3 }
console.log(filteredObj2) // { name: 'John', city: 'New York' }
```

## 函数操作

### 21.柯里化函数

```js
const curry = fn => (…args) => args.length >= fn.length ? fn(…args) : curry(fn.bind(null, …args));
```

示例：

```js
const add = (a, b, c) => a + b + c
const curriedAdd = curry(add)

console.log(curriedAdd(1)(2)(3)) // 6
console.log(curriedAdd(1, 2)(3)) // 6
console.log(curriedAdd(1, 2, 3)) // 6
```

### 22.带记忆功能

```js
const memoize = fn => { const cache = {}; return (…args) => cache[args] = cache[args] || fn(…args); };
```

示例：

```js
const fibonacci = (n) => {
    if (n < 2) return n
    return fibonacci(n - 1) + fibonacci(n - 2)
}

const memoizedFibonacci = memoize(fibonacci)

console.log(memoizedFibonacci(10)) // 55
console.log(memoizedFibonacci(10)) // 55 (从缓存中取值)
console.log(memoizedFibonacci(20)) // 6765
console.log(memoizedFibonacci(20)) // 6765 (从缓存中取值)
```

### 23.类似于 Unix 中的管道符 (|), 将多个函数组合起来，形成一个新的函数

```js
const pipe = (…fns) => x => fns.reduce((v, f) => f(v), x);
```

示例：

```js
const double = (x) => x * 2
const addOne = (x) => x + 1
const square = (x) => x * x

const pipeline = pipe(double, addOne, square)

console.log(pipeline(2)) // 25
```

### 24. 将多个函数组合成一个新的函数

依次执行这些函数，从右到左，并将每个函数的输出作为下一个函数的输入。

```js
const compose = (…fns) => x => fns.reduceRight((v, f) => f(v), x);
```

示例：

```js
const double = (x) => x * 2
const addOne = (x) => x + 1
const square = (x) => x * x

const composed = compose(square, addOne, double)

console.log(composed(2)) // 25 ==> 等价于 square(addOne(double(2)))
console.log(compose(5)) // 121 ==> 等价于 square(addOne(double(5)))
console.log(compose(10)) // 441 ==> 等价于 square(addOne(double(10)))
```

### 25.节流

```js
const throttle = (fn, wait) => {
    let lastTime = 0;
    return (…args) => {
        const now = new Date().getTime();
        if (now — lastTime >= wait) {
            lastTime = now; fn(…args);
        }
    };
};
```

示例：

```js
const throttledFn = throttle((...args) => {
    console.log(...args)
}, 1000)

throttledFn('Hello') // 输出 'Hello'
throttledFn('World') // 不输出

setTimeout(() => {
    throttledFn('Hello') // 输出 'Hello'
    throttledFn('World') // 不输出
}, 1500)

setTimeout(() => {
    throttledFn('Hello') // 输出 'Hello'
    throttledFn('World') // 不输出
}, 2500)
```

### 26.防抖

```js
const debounce = (fn, delay) => {
    let timeoutId;
    return (…args) => {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => fn(…args), delay);
    };
};
```

示例：

```js
const debounce = (fn, delay) => {
    let timeoutId
    return (...args) => {
        clearTimeout(timeoutId)
        timeoutId = setTimeout(() => fn(...args), delay)
    }
}

const debouncedFn = debounce((...args) => {
    console.log(...args)
}, 500)

debouncedFn('Hello')
debouncedFn('World')

setTimeout(() => {
    debouncedFn('Hello')
}, 1000)

// 输出顺序为：
// "World"
// "Hello"
```

### 27.Once

```js
const once = fn => {
    let called = false;
    return (…args) => !called && (called = true, fn(…args));
};
```

示例：

```js
const onceFn = once((...args) => {
    console.log(...args)
})

onceFn('Hello') // 输出 'Hello'
onceFn('World') // 不输出
onceFn('Hello1') // 不输出
```

## 字符串操作

### 28.字符串反转

```js
const reverse = (str) => str.split('').reverse().join('')
```

示例：

```js
console.log(reverse('hello')) // olleh
console.log(reverse('abcdefg')) // gfedcba
console.log(reverse('')) // (空字符串)
```

### 29.字符串转成大写

```js
const capitalize = (str) => str.charAt(0).toUpperCase() + str.slice(1).toLowerCase()
```

示例：

```js
console.log(capitalize('hello')) // Hello
console.log(capitalize('world')) // World
console.log(capitalize('')) // (空字符串)
```

### 30.检查回文

```js
const isPalindrome = (str) => str === str.split('').reverse().join('')
```

示例：

```js
console.log(isPalindrome('radar')) // true
console.log(isPalindrome('hello')) // false
console.log(isPalindrome('')) // true
console.log(isPalindrome('a')) // true
```

### 31.重复字符串

```js
const repeat = (str, n) => str.repeat(n)
```

示例：

```js
console.log(repeat('hello', 3)) // 'hellohellohello'
console.log(repeat('abc', 2)) // 'abcabc'
console.log(repeat('', 5)) // ''
console.log(repeat('a', 4)) // 'aaaa'
```

### 32.统计字符串中元音字母（a、o、e、i、u）的数量

```js
const countVowels = (str) => str.match(/[aeiou]/gi)?.length || 0
```

示例：

```js
console.log(countVowels('hello')) // 2
console.log(countVowels('world')) // 1
console.log(countVowels('')) // 0
console.log(countVowels('aeiou')) // 5
```

### 33.删除空格

```js
const trim = (str) => str.trim()
```

示例：

```js
console.log(trim('   hello   ')) // 'hello'
console.log(trim('   world')) // 'world'
console.log(trim('hello   ')) // 'hello'
console.log(trim('')) // ''
```

### 34.将字符串转换为短横线连接的形式（kebab case）

```js
const kebabCase = (str) => str.replace(/\s+/g, '-').toLowerCase()
```

示例：

```js
console.log(kebabCase('hello world')) // 'hello-world'
console.log(kebabCase('Hello World')) // 'hello-world'
console.log(kebabCase('multiple   spaces')) // 'multiple-spaces'
console.log(kebabCase('')) // ''
```

### 35.将字符串转换为下划线连接的形式（snake case）

```js
const snakeCase = (str) => str.replace(/\s+/g, '_').toLowerCase()
```

示例：

```js
console.log(snakeCase('hello world')) // 'hello_world'
console.log(snakeCase('Hello World')) // 'hello_world'
console.log(snakeCase('multiple   spaces')) // 'multiple_spaces'
console.log(snakeCase('')) // ''
```

## 数字操作

### 36.检查偶数

```js
const isEven = (num) => num % 2 === 0
```

示例：

```js
console.log(isEven(10)) // true
console.log(isEven(11)) // false
console.log(isEven(0)) // true
console.log(isEven(-10)) // true
console.log(isEven(-11)) // false
```

### 37.检查奇数

```js
const isOdd = (num) => num % 2 !== 0
```

示例：

```js
console.log(isOdd(10)) // false
console.log(isOdd(11)) // true
console.log(isOdd(0)) // false
console.log(isOdd(-10)) // false
console.log(isOdd(-11)) // true
```

### 38.随机数

```js
const random = (min, max) => Math.random() * (max — min) + min;
```

示例：

```js
console.log(random(1, 10)) // 生成一个介于 1 和 10 之间的随机数
console.log(random(0, 100)) // 生成一个介于 0 和 100 之间的随机数
console.log(random(-10, 10)) // 生成一个介于 -10 和 10 之间的随机数
```

### 39.随机整数

```js
const randomInt = (min, max) => Math.floor(Math.random() * (max — min + 1)) + min;
```

示例：

```js
console.log(randomInt(1, 10)) // 生成一个介于 1 和 10 之间的随机整数
console.log(randomInt(0, 100)) // 生成一个介于 0 和 100 之间的随机整数
console.log(randomInt(-10, 10)) // 生成一个介于 -10 和 10 之间的随机整数
```

### 40.阶乘

```js
const factorial = n => n <= 1 ? 1 : n * factorial(n — 1);
```

示例：

```js
console.log(factorial(5)) // 120
console.log(factorial(3)) // 6
console.log(factorial(0)) // 1
console.log(factorial(1)) // 1
```

### 41.斐波那契

```js
const fib = n => n <= 1 ? n : fib(n — 1) + fib(n — 2);
```

示例：

```js
console.log(fib(10)) // 55
console.log(fib(5)) // 5
console.log(fib(0)) // 0
console.log(fib(1)) // 1
```

### 42.计算两个数的最大公约数（GCD）

```js
const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b))
```

示例：

```js
console.log(gcd(48, 18)) // 6
console.log(gcd(101, 103)) // 1
console.log(gcd(12, 15)) // 3
console.log(gcd(24, 30)) // 6
```

### 43.这个函数的功能是：计算两个数的最小公倍数（LCM）

```js
const lcm = (a, b) => Math.abs(a * b) / gcd(a, b)
```

示例：

```js
console.log(lcm(4, 6)) // 12
console.log(lcm(10, 15)) // 30
console.log(lcm(12, 18)) // 36
console.log(lcm(24, 30)) // 60
```

### 44.转二进制

```js
const toBinary = (num) => num.toString(2)
```

示例：

```js
console.log(toBinary(10)) // '1010'
console.log(toBinary(20)) // '10100'
console.log(toBinary(30)) // '11110'
console.log(toBinary(40)) // '101000'
```

### 45.转十六进制

```js
const toHex = (num) => num.toString(16)
```

示例：

```js
console.log(toHex(10)) // 'a'
console.log(toHex(20)) // '14'
console.log(toHex(30)) // '1e'
console.log(toHex(40)) // '28'
```
