---
theme: fancy
---
## 语法

### 区分大小写

无论是变量名、函数名还是操作符，都是区分大小写的，比如：test 和 Test 是两个完全不同的变量，类似 typeof 不能作为函数名，因为他是一个关键字；typeof 也是一个完全有效的函数名。

### 标识符

所谓标识符，就是变量、函数、属性或者函数参数的名称；其规范：

*   第一个字符必须是一个字母、下划线或者美元符号。
*   剩下的字符可以是字母(可以是 ASCII 编码或者 Unicode 编码)、下划线、美元符号或者数字。

按照惯例，ECMAScript 标识符使用驼峰大小写形式，也就是第一个单词的首字母小写，后面每个单词的首字母大写（不是强制性要求，但是建议这么写）。

### 注释：

**单行注释和块注释**

*   单行注释：以两个斜杠字符开头。
*   块注释：以一个斜杠和一个星号开头，以它们反向组合结尾。

```javascript
// 单行注释

/* 块注释 */
```

### 严格模式：

在 ECMAScript 5 增加了严格模式的概念。严格模式是一种不同的 JavaScript 解析和执行模型，ECMAScript 3 的一些不规范写法在这种模式下不会被处理，对于不安全的活动将抛出错误，要对整个脚本启动严格模式，在脚本开头加上这一行：`use strict`。

虽然看起来像个没有赋值的字符串，但它其实是一个预处理指令。任何支持 JavaScript 引擎看到它都会切换到严格模式。选择这种语法形式的目的是不破坏 ECMAScript 3 语法。

也可以单独自定一个函数在严格模式下执行，只要把这个预处理指令放到函数体头部即可：

```javascript
function doSomething() {
    'use strict';
}
```

> 严格模式会影响 JavaScript 的执行，所有现代浏览器都支持严格模式。

## 语句

ECMAScript 中的语句**以分号结尾，省略分号意味着有解析器确定语句在哪里结尾**。如下面的例子：

```javascript
let sum = a + b // 没有加分号也有效，但不推荐
let diff = a - b; // 加分号有效，推荐
```

即使语句末尾的分号不是必须的，也应该加上；**分号有助于防止省略造成的问题**，比如一个可以避免输入内容不完整；此外，**加分号也便于开发者通过删除空行来压缩代码**（如果没有结尾的分号，只能删除空行，则会导致语法出错）；**加分号也有助于在某些情况下提升性能**，因此解析器会尝试在合适的位置不上分号以纠正语法错误。
多条语句可以合并一个代码块中；代码块由一个左花括号标识开始，一个右花括号标识结束：

```javascript
if (test) {
    test = false;
    console.log(test);
}
```

**if 之类的控制语句只在执行多条语句时要求必须有代码块**。不过，最佳实践是始终在控制语句中使用代码块，即使要执行的只有一个语句，如下：

```javascript
if (test) console.log(test);

if (test) {
    console.log(test);
}
```

### 保留字与关键字

ECMA-262 描述了一组保留的关键字，这些关键字有特殊用途：

*   比如标识**控制语句的开始与结束或者执行特定的操作**。
*   **保留字不能标识符和属性名**。

## 变量

ECMAScript 变量是松散类型的，意思是变量可以用于保存任何类型的数据；每个变量只不过是一个用户保存任意值的命名占位符。目前有三个关键字可以生命变量：`var`、`let`、`const`。

> `var` 在 ECMAScript 的所有版本中都可以使用，而 `const` 和 `let` 只能在 ECMAScript 6 及更高的版本中使用。

### var 关键字

用于定义变量，关键字 `var` 后跟变量名，可用于保存任何类型的值。

```javascript
var message;
```

*   var 声明作用域
    *   **局部变量**：比如在一函数内定义一个变量，就意味着该变量在函数退出时销毁该变量。
        ```javascript
        function sayHi() {
            var test = 'hi!';
        }
        sayHi();
        console.log(test); //  test is not defined
        ```

    *   **全局变量**：比如在函数内部定义变量时省略 var 操作符即可创建一个全局变量；在严格模式下，给未声明的标量赋值，则会抛出`ReferenceError`错误。

        ```javascript
        function sayHi() {
            test = 'global variable';
        }
        
        sayHi();
        console.log(test); // global variable
        ```

如果需要定义多个变量，可以在一条语句中用逗号隔开每个变量。

```javascript
var message = 'hi',
    found = false,
    age = 20;
```

在严格模式下，不能定义名为 eval 和 arguments 的变量，否则会导致语句错误。

> **使用 var 关键字声明的变量会自动提升到函数作用域顶部**

```javascript
function foo() {
    console.log(age);
    var age = 26;
}

foo() // undefined
// 之所以打印出undefined，没有报错，是因为ECMAScript运行时把它看成等价如下代码：
function foo() {
    var age;
    console.log(age);
    age = 20;
}
```

这就是所谓的变量提升，也就是把**所有变量声明都拉到函数作用域的顶部**。还可以使用 `var` 声明同一个变量也没有问题：

```javascript
function foo() {
    var age = 20;
    var age = 21;
    var age = 22;
    var age = 23;
    console.log(age); // 23
}
```

### let 声明

和 `var` 的作用差不多，但有着非常重要的区别，最明显的区别时，`let` 声明的范围是块级作用域，而 `var` 声明的范围是函数作用域；如下两段代码：

```javascript
// a.js
if (true) {
    var name = 'matt';
    console.log(name); // matt
}
console.log(name); // matt

// b.js
if (true) {
    let name = 'matt';
    console.log(name); // matt
}
console.log(name) // ReferenceError: name a is undefined
```

> 之所以 b.js 中 `if` 外的打印会报错，是因为 `name` 变量的作用域被限制在 `if` 以内了，变量不能在外部使用；块作用域是函数作用域的子集，因此适用于 `var` 的作用域同样也适用于 `let`。

```javascript
var name = 'foo';
var name = 'bar';

let name = 'foo';
let name = 'bar'; // Uncaught SyntaxError: Identifier 'name' has already been declared
```

> var 对同一标识符在同一作用域中重新定义不会报错，而 let 则会报错。

```javascript
var age = 20;
let age = 20; // Uncaught SyntaxError: Identifier 'age' has already been declared
```

> 对声明冗余报错不会因混用 `let` 和 `var` 而受影响；这两关键字声明的并不是不同类型的变量，他们只是指出变量在相关作用域如何存在。

| 区别                      |                             var                             |                             let                              |
| ------------------------- | :---------------------------------------------------------: | :----------------------------------------------------------: |
| 暂时性死区                |               变量提升，没有暂时性死区的概念                | 在变量声明的作用域前使用该变量或者作用域之外使用该变量都会报错，故而不能在此之前以任何方式来引用未声明的变量 |
| 全局声明                  |          var 声明的变量会会成为 window 对象的属性           |          `let` 声明的则不会成为 `window` 对象的属性          |
| 条件声明                  | JavaScript 引擎会自动将多余的声明在作用域顶部合并为一个声明 |            条件块中 `let` 声明的作用域仅限于该块             |
| `for` 循环中的 `let` 声明 |              定义的迭代变量会渗透到循环体外部               |          定义的迭代变量的作用域仅在 `for` 循环内部           |

### const 声明

> `const` 的行为与 `let` 基本相同，唯一一个重要的区别就是用它声明变量的时候必须初始化值，且尝试修改 `const` 声明的变量会导致运行出错。

特点：

*   **声明时必须初始化值，且不能重新赋值给变量**。
*   **不允许重复声明**。
*   **声明的变量的作用域也是块**。
*   `const` 不能用在 `for` 循环中来声明迭代变量（因为迭代变量会自增）；如果只想用 `const` 声明每次迭代只是创建一个新变量，则不会有问题，比如用在 `for-in`、`for-of` 中。

```javascript
<script>
    const age = 26;
    age = 36; // Uncaught TypeError: Assignment to constant variable.

		// for
    for (const i = 0; i < 10; i++) { // Uncaught TypeError: Assignment to constant variable.
        console.log(i)
    }

    // for in
    for (const key in {name: 'Forest', age: 21}) {
        console.log(key); // name age
    }

    // for of
    for (const item of [1, 1, 2, 3, 5, 8, 13, 21, 34]) {
        console.log(item); // 1, 1, 2, 3, 5, 8, 13, 21, 34
    }
</script>
```

### 声明风格及最佳实践

> ECMAScript 6 增加的 `let` 和 `const` 解决了怪异行为的 `var` 所造成的各种问题，从客观上为这门语言更精确地声明作用域和语义提供了过呢个好的支持，这也有助于代码提升质量的最佳实践也逐渐显现。

#### 实践建议

*   尽可能的不使用 `var`；`const` 和 `let` 声明的变量有助于代码质量的提升，因为有了明确的作用域、声明位置，以及不变的值。
*   `const` 优先，`let` 次之；使用从 `const` 声明可以让浏览器运行时强制保持变量不变，也可以然静态代码分析工具提前发现不合法的赋值操作，只在提前知道未来会修改变量时，在使用 `let`。

## 基本数据类型

> 在 ECMAScript 中的数据类型有：`Undefined`、`Null`、`Boolean`、`Number`、`String`、`Array`、`Symbol`、`Object`。

### 判断数据类型

#### typeof 操作符

> 因为 ECMAScript 的类型系统时松散的，所以需要一种给手段来确定任意变量的数据类型；typeof 就是为此而生；返回下列字符串的相关说明：
>
> *   **`undefined`** : 表示*未定义*。
> *   **`boolean`** : 表示值为*布尔值*。
> *   **`string`** ：表示值为*字符串*。
> *   **`number`** ：表示值为*数值*。
> *   **`object`** ：表示值为*对象（而不是函数）或者 null*。
> *   **`function`** ：表示值为*函数*。
> *   **`symbol`** ：表示值为*符号*。

```javascript
<script>
    const name = 'Forest';
    const age = 25;
    const haveGirlFriend = false;
    const skills = ['html5', 'css3', 'javascript', 'git', 'vue', 'react', 'node', '微信小程序', '微信云开发', '微信公众号开发', 'go', 'mongodb', 'mysql', 'redis', 'docker', 'jenkins'];
    const hobby = {
        ball: ['篮球', '乒乓球'],
        read: {name: 'javascript高级程序设计', version: '第四版'}
    }
    const func = () => hobby['read']
    const symbolVal = Symbol(name);
    const deposit = null;

    console.log(typeof name);           // string
    console.log(typeof age);            // number
    console.log(typeof haveGirlFriend); // boolean
    console.log(typeof skills);         // object  数组也是一个特殊的object
    console.log(typeof hobby);          // object
    console.log(typeof deposit);        // object  因为特殊值null被认为是一个对空对象的引用，所以返回值为null
    console.log(typeof func);           // function
    console.log(typeof symbolVal);      // symbol
    console.log(typeof girlFriend);     // undefined
</script>
```

> 在 JavaScript 中判断类型的方法可不止`typeof`， 还可以使用`instanceof`、`Object.prototype.toString`。

#### instance 判断值的类型

instanceof 的方法通过 new 一个对象，这个新对象就是它原型链继承上面的对象了，通过 instanceof 我们能判断这个对象是否是之前那个构造函数生成的对象，这样就基本可以判断出这个新对象的数据类型。

```javascript
let Car = function () {}
let benz = new Car();
console.log(benz instanceof Car);       // true
let car = new String('Mercedes Benz');
console.log(car instanceof String);     // true
let str = 'Covid-19';
console.log(str instanceof String);     // false
```

> Tips:
>
> *   `instanceof` 可以准确地判断复杂引用数据类型，但是不能正确判断基础数据类型。
> *   而 `typeof` 也存在弊端，它虽然可以判断基础数据类型（`null` 除外），但是引用数据类型中，除了 `function` 类型以外，其他的也无法判断。

#### Object.prototype.toString()

> `toString()` 是 `Object` 的原型方法，调用该方法，可以统一返回格式为`[object Xxx]` 的字符串，其中 `Xxx` 就是对象的类型。对于 `Object` 对象，直接调用 `toString()` 就能返回 `[object Object]`；而对于其他对象，则需要通过 `call` 来调用，才能返回正确的类型信息。

```javascript
console.log(Object.prototype.toString({}));                  // [object Object]
console.log(Object.prototype.toString.call({}));             // [object Object]
console.log(Object.prototype.toString.call(1));              // [object Number]
console.log(Object.prototype.toString.call('1'));            // [object String]
console.log(Object.prototype.toString.call(true));           // [object Boolean]
console.log(Object.prototype.toString.call(function () {})); // [object Function]
console.log(Object.prototype.toString.call(null));           // [object Null]
console.log(Object.prototype.toString.call(undefined));      // [object Undefined]
console.log(Object.prototype.toString.call(/123/g));         // [object RegExp]
console.log(Object.prototype.toString.call(new Date()));     // [object Date]
console.log(Object.prototype.toString.call([]));             // [object Array]
console.log(Object.prototype.toString.call(document));       // [object HTMLDocument]
console.log(Object.prototype.toString.call(window));         // [object Window]
```

> `Object.prototype.toString.call()` 可以很好地判断引用类型，甚至可以把 `document` 和 `window` 都区分开来。

但是在写判断条件的时候一定要注意，使用这个方法最后**返回统一字符串格式为 "\[object Xxx]" ，而这里字符串里面的 "Xxx" ，第一个首字母要大写（注意：使用 typeof 返回的是小写）**。

### 基础类型

#### Undefined

`Undefined` 类型是一个特殊的值，当声明变量后没有赋初始值时，变量的默认值就是 `undefined`。

```js
let nickname; // 只声明变量未初始化值
console.log(nickname == undefined); // true
```

> 包含 `undefined` 值的变量跟未定义变量是有区别的：
>
> *   一个未定义的遍历那个是会直接报错的。
> *   对未声明的变量，只能执行一个有用的操作，就是对它调用 `typeof`，在对未初始化的变量调用 `typeof` 时，返回的结果还是 `undefined`。
>     ```js
>     let nickname;
>     // let age;
>     
>     console.log('nickname 变量的类型：', typeof nickname); // nickname 变量的类型： undefined
>     console.log('未声明变量 age 的类型：', typeof age)      // 未声明变量 age 的类型： undefined
>     ```
>     最佳实践：建议在声明变量的同时进行初始化，当 `typeof` 返回 `undefined` 时，你就会知道那是因为给定的变量尚为声明而不是生命力但未初始化。
> *   `undefined` 是一个假值。

#### Null

`Null`同样是一个特殊值，逻辑上讲，`null` 值表示一个空对象指针，这也是给 `typeof` 传一个 `null` 会返回 `object` 的原因。

```js
let nickname = null;
console.log('nickname 的类型：', typeof nickname); // nickname 的类型： object
```

> 在定义将来要保存对象的变量时，建议使用 `null` 来初始化，不要使用其他值。

`undefined` 值是由 `null` 值派生而来的，因此 ECMA-262 将它们定义为表面上相等，如下面的例
子所示：

```js
console.log(null == undefined); // true
```

> 用等于操作符（`==`）比较 `null` 和 `undefined` 始终返回 `true`。但要注意，这个操作符会为了比较而转换它的操作数。

#### Boolean

Boolean（布尔值）类型是 ECMAScript 中使用最频繁的类型之一，有两个字面值：`true` 和 `false`。
这两个布尔值不同于数值，**因此 `true` 不等于 1，`false` 不等于 0**。

```js
const found = true;
const lost = false;
```

> **布尔值字面量 `true` 和 `false` 是区分大小写的**，*因此 `True` 和 `False`（及其他大小混写形式）是有效的标识符，但不是布尔值*。

虽然布尔值只有两个，但所有其他 ECMAScript 类型的值都有相应布尔值的等价形式。要将一个其它类型的值转换为布尔值，可以调用特定的 `Boolean()`转型函数：

```js
const message = 'hello world!';
const messageAsBoolean = Boolean(message); // 字符串 message 会被转换为布尔值并保存在变量 messageAsBoolean 中
```

`Boolean()` 转型函数可以在任意类型的数据上调用，而且始终返回一个布尔值。什么值能转换为 `true`
或 `false` 的规则取决于数据类型和实际的值。转换规则如下：

| 数据类型  | 转换为true的值         | 转换为false的值 |
| --------- | ---------------------- | --------------- |
| Boolean   | true                   | false           |
| String    | 非空字符串             | ""              |
| Number    | 非零数值（包括无穷值） | 0、NaN          |
| Object    | 任意对象               | null            |
| Undefined | N/A（不存在）          | undefined       |

*   `Boolean` 类型：只有 `true` 会转换为 `true`，而 `false` 会转换为 `false`。
    ```javascript
    console.log(Boolean(true));  // 输出: true
    console.log(Boolean(false)); // 输出: false
    ```
*   `String` 类型：任何非空字符串都会转换为 `true`，而空字符串""会转换为 `false`。
    ```javascript
    console.log(Boolean("Hello"));  // 输出: true
    console.log(Boolean(""));       // 输出: false
    ```
*   `Number`类型：任何非零数值（包括无穷值）会转换为 `true`，而0和 `NaN` 会转换为 `false`。
    ```javascript
    console.log(Boolean(123));      // 输出: true
    console.log(Boolean(-456));     // 输出: true
    console.log(Boolean(Infinity)); // 输出: true
    console.log(Boolean(0));        // 输出: false
    console.log(Boolean(NaN));      // 输出: false
    ```
*   `Object` 类型：任何对象都会转换为 `true`，而 `null` 会转换为 `false`。
    ```javascript
    console.log(Boolean({name: "Tom"}));  // 输出: true
    console.log(Boolean(null));           // 输出: false
    ```
*   Undefined类型：undefined会转换为false。
    ```javascript
    console.log(Boolean(undefined));  // 输出: false
    ```

#### Number

Number 类型是一种基于 [IEEE 754 标准的双精度 64 位二进制格式的值](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Number#number_%E7%BC%96%E7%A0%81)。它**能够存储 $2^{-1074}$（`Number.MIN_VALUE`）和 $2^{1024}$（`Number.MAX_VALUE`）之间的正浮点数，以及 $-2^{-1074}$ 和 $-2^{1024}$ 之间的负浮点数**，但是它**仅能安全地存储在 -($2^{53} − 1$)（`Number.MIN_SAFE_INTEGER`）到 $2^{53} − 1$（`Number.MAX_SAFE_INTEGER`）范围内的整数**。超出这个范围，JavaScript 将不能安全地表示整数；相反，它们将由双精度浮点近似表示。你可以使用 `Number.isSafeInteger()` 检查一个数是否在安全的整数范围内。

±(2-1074 \~ 21024) 范围之外的值会自动转换：

*   大于 `Number.MAX_VALUE` 的正值被转换为 `+Infinity`。
*   小于 `Number.MIN_VALUE` 的正值被转换为 `+0`。
*   小于 `-Number.MAX_VALUE` 的负值被转换为 `-Infinity`。
*   大于 `-Number.MIN_VALUE` 的负值被转换为 `-0`。

> `+Infinity` 和 `-Infinity` 行为类似于数学上的无穷大，但是有一些细微的区别；

**最基本的数值字面量格式是十进制整数**，如下：

```js
const intNum = 55; // 整数
```

整数也可以用**八进制**（*以 8 为基数*）或**十六进制**（*以 16 为基数*）字面量表示。*对于八进制字面量，第一个数字必须是零（0），然后是相应的八进制数字（数值 0\~7）*。**如果字面量中包含的数字超出了应有的范围，就会忽略前缀的零，后面的数字序列会被当成十进制数**，如下所示：

```js
const octalNum1 = 070;   // 八进制的 56 ==> 0*8^0 + 7*8^1 + 0*8^3 ==> 0 + 56 + 0
const octalNum2 = 079;   // 无效的八进制值，直接当十进制处理，结果为 79
const octalNum3 = 08;    // 无效的八进制值，直接当十进制处理，结果为 8
```

> 八进制字面量在严格模式下是无效的，会导致 JavaScript 引擎抛出语法错误。**ECMAScript 2015(ES6) 中的八进制值通过前缀 `0o` 来表示；严格模式下，前缀 0 会被视为语法错误，如果要表示八进制值，应该使用前缀 `0o`。**

要创建**十六进制字面量，必须让真正的数值前缀 0x（区分大小写），然后是十六进制数字（0~~9 以
及 A~~F）**。十六进制数字中的字母大小写均可。下面是几个例子：

```js
const hexNum1 = 0xA;  // 十六进制 10 ==> A=10 B=11 C=12 D=13 E=14 F=15
const hexNum2 = 0x1f; // 十六进制 31 ==> 15*16^0 + 1*16^1 ==> 15 + 16 ==> 31
```

> 由于 JavaScript 保存数值的方式，实际中可能存在正零（`+0`）和负零（`-0`）。**正零和
> 负零在所有情况下都被认为是等同的**。

##### 浮点值

要定义浮点值，数值中必须包含小数点，而且小数点后面必须至少有一个数字。虽然小数点前面不
是必须有整数，但推荐加上。

```js
const floatNum1 = 1.1; 
const floatNum2 = 0.1; 
const floatNum3 = .1; // 有效，但不推荐
```

因为**存储浮点值使用的内存空间是存储整数值的两倍**，所以 ECMAScript 总是想方设法把值转换为
整数。在小数点后面没有数字的情况下，数值就会变成整数。类似地，如果数值本身就是整数，只是小
数点后面跟着 0（如 1.0），那它也会被转换为整数。

```js
const floatNum1 = 1.;   // 小数点后面没有数字，当成整数 1 处理
const floatNum2 = 10.0; // 小数点后面是零，当成整数 10 处理
```

> 为什么存储浮点值使用的内存空间是存储整数值的两倍？
>
> JavaScript 的 Number 类型是一个[双精度 64 位二进制格式 IEEE 754 值](https://zh.wikipedia.org/wiki/%E9%9B%99%E7%B2%BE%E5%BA%A6%E6%B5%AE%E9%BB%9E%E6%95%B8)，类似于 Java 或者 C# 中的 double。这意味着它可以表示小数值，但是存储的数字的大小和精度有一些限制。简而言之，I**EEE 754 双精度浮点数使用 64 位来表示 3 个部分**：
>
> *   **1 位用于表示符号（sign）（正数或者负数）**
> *   **11 位用于表示指数（exponent）（-1022 到 1023）**
> *   **52 位用于表示尾数（mantissa）（表示 0 和 1 之间的数值）**
>
> 尾数（也称为有效数）是表示实际值（有效数字）的数值部分。指数是尾数应乘以的 2 的幂次。将其视为科学计数法：
>
> $$
> Number=(−1)^{sign}⋅(1+mantissa)⋅2^{exponent}
> $$

**对于非常大或非常小的数值，浮点值可以用科学记数法来表示**。*科学记数法用于表示一个应该乘以
10 的给**定次幂**的数值。ECMAScript 中科学记数法的格式要求是一个数值（整数或浮点数）后跟**一个大写或小写的字母 e**，再加上一个要乘的 10 的多少次幂*。

```js
const floatNum = 3.125e7; // 等于 31250000 ==> 3.125*10^7 ==> 以 3.125 作为系数，乘以 10 的 7 次幂
```

科学记数法也可以用于表示非常小的数值；例如 0.000 000 000 000 000 03：

```js
const floatNum = 3e-17; // ==> 3*10^-17 ==> 以 3 为系数，乘以 10 的 -17 次幂
```

**浮点值的精确度最高可达 *17* 位小数，但在算术计算中远不如整数精确**。
例如，0.1 加 0.2 得到的不是 0.3，而是 0.300 000 000 000 000 04。由于这种微小的舍入错误，导致很难测试特定的浮点值。

> 之所以存在这种舍入错误，是因为使用了 **IEEE 754** 数值，这种错误并非 ECMAScript
> 所独有。其他使用相同格式的语言也有这个问题。

##### 值的范围

> 由于**内存的限制**，ECMAScript 并**不支持表示这个世界上的所有数值**；**如果某个计算得到的数值结果超出了 JavaScript 可以表示的范围，那么这个数值会被自动转换为一个特殊的 Infinity（无穷）值**。任何无法表示的负数以 `-Infinity`（负无穷大）表示，任何无法表示的正数以 `Infinity`（正无穷大）表示。

*   ECMAScript 可以表示的**最小数值保存在 `Number.MIN_VALUE`** 中，这个值在多数浏览器中是 `5e-324`。
*   可以表示的**最大数值保存在 `Number.MAX_VALUE`** 中，这个值在多数浏览器中是 `1.797693134862315 7e+308`
*   **如果计算返回正 `Infinity` 或负 `Infinity`，则该值将不能再进一步用于任何计算**。这是因为 `Infinity` 没有可用于计算的数值表示形式。要确定一个值是不是有限大（即介于 JavaScript 能表示的最小值和最大值之间），可以使用 `isFinite()` 函数，如下所示：
    ```js
    let result = Number.MAX_VALUE + Number.MAX_VALUE; 
    console.log(isFinite(result)); // false
    ```

只有在 $-2^{53}+1$ 到 $2^{53}-1$ 范围内（闭区间）的整数才能在不丢失精度的情况下被表示（可通过 `Number.MIN_SAFE_INTEGER` 和 `Number.MAX_SAFE_INTEGER` 获得），因为尾数只能容纳 53 位（包括前导 1）。

```js
const biggestInt = Number.MAX_SAFE_INTEGER; // (2**53 - 1) ==> 9007199254740991
const smallestInt = Number.MIN_SAFE_INTEGER; // -(2**53 - 1) ==> -9007199254740991
```

当解析已序列化为 JSON 的数据时，**如果 JSON 解析器将它们强制为 Number 类型，超出此范围的整数值可能会被损坏。一个可能的变通办法是使用 String。较大的数值可以使用 BigInt 类型表示**。

##### NaN

NaN(Not a Number)表示**本来要返回数值的操作失败了**（而不是抛出错误）。比如，用 0 除任意数值在其他语言中通常都会导致错误，从而中止代码执行。但在 ECMAScript 中，0、+0 或0 相除会返回 NaN：

```js
console.log(0/0);   // NaN
console.log(-0/+0); // NaN
```

如果**分子是非 0 值，分母是有符号 0 或无符号 0，则会返回 `Infinity` 或 `-Infinity`**：

```js
console.log(5/0);   // Infinity 
console.log(5/-0);  // -Infinity
```

有符号和无符号是什么意思？

> "有符号"和"无符号"是用来指代**数值数据类型是否可以表示正数和负数**的术语。
>
> *   有符号（Signed）：**一个有符号的整数可以表示正数、负数以及零**。在这种情况下，它的一部分位（通常是最高位，也被称为"符号位"）被用来表示这个数是正数还是负数。例如，如果我们有一个8位的有符号整数，那么它可以表示的范围就是 -128 到 127 。
> *   无符号（Unsigned）：**一个无符号的整数只能表示非负整数（也就是，零和正数）**。因为没有位被用来表示符号，所以无符号的整数可以表示一个更大的范围。例如，一个 8 位的无符号整数可以表示的范围是 0 到 255。

NaN 的几个特性：

*   **任何涉及 `NaN` 的操作始终返回 `NaN`**
*   **`NaN` 不等于包括 `NaN` 在内的任何值，`NaN` 与 `NaN` 之间相等操作始终是 `false`**
    ```js
    console.log(NaN == NaN); // false
    ```
*   \*\*判断一个数是否是 `NaN`，可以通过 `isNaN()`，\*\*接收一个参数，可以是任意数据类型，然后判断
    这个参数是否“不是数值”。
    > **把一个值传给 isNaN()后，该函数会尝试把它转换为数值**。某些非数值的
    > 值可以直接转换成数值，如字符串"10"或布尔值。**任何不能转换为数值的值都会导致这个函数返回
    > true**。例如：
    >
    > ```js
    > console.log(isNaN(NaN));     // true 
    > console.log(isNaN(10));      // false，10 是数值
    > console.log(isNaN("10"));    // false，可以转换为数值 10 
    > console.log(isNaN("blue"));  // true，不可以转换为数值
    > console.log(isNaN(true));    // false，可以转换为数值 1
    > ```

##### 数值转换

有 3 个函数可以将非数值转换为数值：`Number()`、`parseInt()`和 `parseFloat()`，`Number()` 是转型函数，可用于任何数据类型。`parseInt()`和 `parseFloat()` 主要用于将字符串转换为数值。

`Number()` 函数转换规则如下：

* 布尔值， `true` 转换为 1， `false` 转换为 0
    ```js
    console.log(Number(true));  // 输出: 1
    console.log(Number(false)); // 输出: 0
    ```
* 数值就直接返回
* `null` 返回 0
    ```js
    console.log(Number(null));  // 输出: 0
    ```
* `undefined` 返回 `NaN`
    ```js
    console.log(Number(undefined));  // 输出: NaN
    ```
* 字符串，应用以下规则。
    * 如果字符串包含数值字符，包括数值字符前面带加、减号的情况，则转换为一个十进制数值。
        因此，`Number("1")`返回 `1`，`Number("123")`返回 `123`，`Number("011")`返回 `11`（忽略前面
        的零）。
    ```js
      console.log(Number("1"));     // 输出: 1
      console.log(Number("123"));   // 输出: 123
      console.log(Number("011"));   // 输出: 11（忽略前导0
    ```
    * 如果字符串包含有效的浮点值格式如"1.1"，则会转换为相应的浮点值（同样，忽略前面的零）。
        ```js
        console.log(Number("1.1"));   // 输出: 1.1
        ```

    * 如果字符串包含有效的十六进制格式如"0xf"，则会转换为与该十六进制值对应的十进制整
        数值。
        ```js
        console.log(Number("0xf"));   // 输出: 15
        ```

    * 如果是空字符串（不包含字符），则返回 0。
        ```js
        console.log(Number(""));      // 输出: 0
        ```

    * 如果字符串包含除上述情况之外的其他字符，则返回 NaN。
        ```js
        console.log(Number("hello")); // 输出: NaN
        ```

    * 前导和尾随的空格/换行符会被忽略。
        ```js
        console.log(Number(" 123 ")); // 输出: 123
        ```

    * 前导的数字 0 不会导致该数值成为八进制字面量（或在严格模式下被拒绝）。
        ```js
        console.log(Number("0123"));  // 输出: 123
        ```

    * `+` 和 `- `允许出现在字符串的开头以指示其符号。（在实际代码中，它们“看起来像”文字的一部分，但实际上是独立的一元运算符）然而，该标志只能出现一次，并且后面不能跟空格。
        ```js
        console.log(Number("-123"));  // 输出: -123
        ```

    * `Infinity` 和 `-Infinity` 被当作是字面量。在实际代码中，它们是全局变量。
        ```js
        console.log(Number("Infinity"));   // 输出: Infinity
        console.log(Number("-Infinity"));  // 输出: -Infinity
        ```

    * 空字符串或仅包含空格的字符串转换为 0。
        ```js
        console.log(Number(" "));  // 输出: 0
        ```

    * 不允许使用数字分隔符。
        ```js
        console.log(Number("1_000"));  // 输出: NaN
        ```

    * 对象首先通过按顺序调用它们的 `[@@toPrimitive]()`（使用 "number" 提示）、`valueOf()` 和 `toString()` 方法将其转换为原始值。然后将得到的原始值转换为数字。
        ```js
        let obj = {
            valueOf: function() {
                return 3;
            }
        };
        console.log(Number(obj));  // 输出: 3
        ```

**`parseInt()` 函数更专注于字符串是否包含数值模式**；字符串最前面的空格会被忽略，从第一个非空格字符开始转换\*\*；如果第一个字符不是数值字符、加号或减号，`parseInt()`立即返回 NaN\*\*；这意味着**空字符串也会返回 `NaN`（这一点跟 `Number()`不一样，它返回 0）**。*如果第一个字符是数值字符、加号或减号，则继续依次检测每个字符，直到字符串末尾，或碰到非数值字符*。比如：

```js
const str = "1234blue";
console.log(parseInt(str));  // 1234

const str1 = "22.5";
console.log(parseInt(str1)); // 22
```

`parseInt()` 函数也能识别不同的整数格式（十进制、八进制、十六进制）。换句话说，如果字符串以`0x`开头，就会被解释为十六进制整数:

```js
const num1 = parseInt("");         // NaN
const Num2 = parseInt("0xA");      // 10 解释为 16 进制
const num3 = parseInt("70");       // 70 解释为 10 进制
const num4 = parseInt(null);       // NaN
const num5 = parseInt(undefined);  // NaN
const num6 = parseInt(true);       // NaN
const num7 = parseInt(false);      // NaN
```

不同的数值格式很容易混淆，因此 `parseInt()` 也接收第二个参数，用于指定底数（进制数）。如
果知道要解析的值是十六进制，那么可以传入 16 作为第二个参数。

```js
const num = parsetInt("0xAF", 16); // 175 ==> 15*16^0 + 10*16^1 ==> 15+160
```

其实，如果第二个参数传入了 16，那么字符串前面的 `0x` 可以省略：

```js
const num1 = parseInt("AF", 16); // 175

const num2 = parseInt("AF"); // NaN --> 区别于上一个语句，是因为传入了进制数作为解析目标
```

第二个参数可以极大扩展转换后获得的结果类型：

```js
const num1 = parseInt("10", 2);   // 2，按二进制解析
const num2 = parseInt("10", 8);   // 8，按八进制解析
const num3 = parseInt("10", 10);  // 10，按十进制解析
const num4 = parseInt("10", 16);  // 16，按十六进制解析
```

`parseFloat()` 函数的工作方式跟 `parseInt()` 函数类似，都是从位置 0 开始检测每个字符。同样，
它也是**解析到字符串末尾或者解析到一个无效的浮点数值字符为止**。这意味着**第一次出现的小数点是有
效的，但第二次出现的小数点就无效了**，此时字符串的剩余字符都会被忽略。

```js
console.log(parseFloat("22.34.5")); // 22.34
```

`parseFloat()` 函数的另一个不同之处在于:

* **它始终忽略字符串开头的零**。
* **十六进制数值始终会返回 0**。
* **`parseFloat()` 只解析十进制值**，
* **如果字符串表示整数（没有小数点或者小数点后面只有一个零），则 parseFloat()返回整数**。

```js
const num1 = parseFloat("1234blue"); // 1234，按整数解析
const num2 = parseFloat("0xA");      // 0 
const num3 = parseFloat("22.5");     // 22.5 
const num4 = parseFloat("22.34.5");  // 22.34 
const num5 = parseFloat("0908.5");   // 908.5 
const num6 = parseFloat("3.125e7");  // 31250000
```

#### String

String（字符串）数据类型表示零或多个 16 位 Unicode 字符序列。字符串可以使用双引号（"）、
单引号（'）或反引号（\`）标示；**ECMAScript 语法中表示字符串的引号没有区别**。不过要注意的是，**以某种引号作为字符串开头，必须仍然以该种引号作为字符串结尾**。

```js
const string1 = "A string primitive";
const string2 = 'Also a string primitive';
const string3 = `Yet another string primitive`;
```

也可以通过 String() 构造函数创建为字符串对象：

```js
const string4 = new String("A String object");
```

##### 字符字面量

| 字面量 | 含义                                                         |
| ------ | ------------------------------------------------------------ |
| \n     | 换行                                                         |
| \t     | 制表                                                         |
| \b     | 退格                                                         |
| \r     | 回车                                                         |
| \f     | 换页                                                         |
| \\\\   | 反斜杠                                                       |
| \\'    | 单引号，在字符串以单引号表示时使用，例如：'He said, 'hey.''  |
| \\"    | 双引号，在字符串以双引号标示时使用                           |
| \\\`   | 反引号（`），在字符串以反引号标示时使用，例如`He said, \`hey.\`\` |
| \xnn   | 以十六进制编码 nn 表示的字符（其中 n 是十六进制数字 0\~F），例如\x41 等于"A" |
| \unnnn | 以十六进制编码 nnnn 表示的 Unicode 字符（其中 n 是十六进制数字 0\~F），例如\u03a3 等于希腊字符"Σ" |

> 如果字符串中包含双字节字符，那么 `length` 属性返回的值可能不是准确的字符数。
>
> 在JavaScript中，字符串的 `length` 属性计算的是字符串中的单位字符数，而不是实际的字符数。这是因为JavaScript内部使用的是 UTF-16 编码，对于大部分常用字符（包括英文字符、数字和标点符号），每个字符占用一个单位长度，所以 `length` 属性的返回值和实际字符数是一致的。
>
> 但是，对于一些字符，如表情符号🙂，某些特殊符号，或者一些非拉丁语系的字符（如中文、日文、韩文等），它们在UTF-16编码中占用两个单位长度，这就导致了 `length` 属性返回的值可能大于实际的字符数。
> 例如：
>
> ```js
> var str = "hello🙂";
> console.log(str.length); // 输出: 7，但实际字符数是6
> ```
>
> 解决这个问题的一个方法是使用 `Array.from()` 函数，它可以正确地将字符串转换成字符数组，然后通过获取数组长度来获取实际的字符数：
>
> ```js
> var str = "hello🙂";
> console.log(Array.from(str).length); // 输出: 6，这是正确的字符数
> ```
>
> 这样就可以得到字符串中实际的字符数，而不是单位字符数。

##### 字符串的特点

**ECMAScript 中的字符串是不可变的（immutable）**，意思是**一旦创建，它们的值就不能变了**。**要修改某个变量中的字符串值，必须先销毁原始的字符串**，然后将包含新值的另一个字符串保存到该变量，如
下所示：

```js
let lang = "Java";
lang = lang + "Script"; // 将 lang 原来的值（Java）和拼接值（Script）组合后，重新赋值给 `lang` 变量
```

##### 转换为字符串

* 字符串按原样返回。只是简单地返回自身的一个副本。
    ```js
    console.log(String("hello"));    // 输出: "hello"
    ```
* undefined 转换成 "undefined"。
    ```js
    console.log(String(undefined));  // 输出: "undefined"
    ```
*   null 转换成 "null"。
    ```js
    console.log(String(null));  // 输出: "null"
    ```
* true 转换成 "true"；false 转换成 "false"。
    ```js
    console.log(String(true));  // 输出: "true"
    console.log(String(false));  // 输出: "false"
    ```
* 使用与 toString(10) 相同的算法转换数字。
    ```js
    console.log(String(123));  // 输出: "123"
    console.log(String(0));    // 输出: "0"
    console.log(String(-123)); // 输出: "-123"
    ```
* 使用与 toString(10) 相同的算法转换 BigInt。
    ```js
    console.log(String(BigInt(123)));  // 输出: "123"
    ```
* Symbol 抛出 TypeError。
    ```js
    console.log(String(Symbol("hello")));  // 抛出 TypeError: Cannot convert a Symbol value to a string
    ```
* 对于对象，首先，通过依次调用其 [@@toPrimitive]()（hint 为 `"string"`）、`toString()` 和 `valueOf()` 方法将其转换为原始值。然后将生成的原始值转换为一个字符串。

##### 模板字面量

模板字面量是 ECMAScript 6 新增的定义字符串的能力，模板字面量是**用反引号（\`）分隔的字面量，允许多行字符串、带嵌入表达式**的字符串插值和一种叫带标签的模板的特殊结构。模板字面量**有时被非正式地叫作模板字符串**，因为它们最常被用作字符串插值（通过替换占位符来创建字符串）

```js
let myMultiLineString = 'first line\nsecond line'; 
let myMultiLineTemplateLiteral = `first line 
second line`; 
console.log(myMultiLineString); 
// first line 
// second line" 

console.log(myMultiLineTemplateLiteral); 
// first line
// second line 

console.log(myMultiLineString === myMultiLinetemplateLiteral); // true
```

模板字面量用反引号（\`）括起来，而不是双引号（"）或单引号（'）。 除了普通字符串外，模板字面量还可以包含占位符——一种由**美元符号和大括号分隔的嵌入式表达式**：`${expression}`

* **若要转义模板字面量中的反引号（\`），需在反引号之前加一个反斜杠（\）**。
    ```js
    console.log(`\`` === "`");       // true
    ```
* **美元符号 `$` 也可以被转义**，来阻止插值。
    ```js
    console.log(`\${1}` === "${1}"); // true
    ```
* **模板字面量会保持反引号内部的空格**，因此在使用时要格外注意。格式正确的模板字符串看起来可能会缩进不当：
    ```js
    // 这个模板字面量在换行符之后有 25 个空格符
    let myTemplateLiteral = `first line 
     second line`; 
    console.log(myTemplateLiteral.length);          // 47 
    
    // 这个模板字面量以一个换行符开头
    let secondTemplateLiteral = ` 
    first line 
    second line`; 
    console.log(secondTemplateLiteral[0] === '\n'); // true 
    
    // 这个模板字面量没有意料之外的字符
    let thirdTemplateLiteral = `first line 
    second line`; 
    console.log(thirdTemplateLiteral); 
    // first line 
    // second line
    ```

##### 字符串插值

**模板字面量最常用的一个特性是支持字符串插值**，也就是可以在一个连续定义中**插入一个或多个
值**。技术上讲，模板字面量不是字符串，而是一种特殊的 JavaScript 句法表达式，只不过求值后得到的
是字符串。模板字面量在定义时立即求值并转换为字符串实例，任何插入的变量也会从它们最接近的作
用域中取值。

字符串插值通过在 `${}` 中使用一个 JavaScript 表达式实现：

```js
let value = 5; 
let exponent = 'second'; 

// 以前，字符串插值是这样实现的：
let interpolatedString = value + ' to the ' + exponent + ' power is ' + (value * value); 
 
// 现在，可以用模板字面量这样实现：
let interpolatedTemplateLiteral = `${ value } to the ${ exponent } power is ${ value * value }`; 
console.log(interpolatedString);           // 5 to the second power is 25 
console.log(interpolatedTemplateLiteral);  // 5 to the second power is 25
```

> 所有插入的值都会使用 `toString()` 强制转型为字符串，而且任何 JavaScript 表达式都可以用于插
> 值。

* 嵌套的模板字符串无须转义
    ```js
    console.log(`Hello, ${ `World` }!`); // Hello, World!
    ```
* 将表达式转换为字符串时会调用 toString()
    ```js
    let foo = { toString: () => 'World' }; 
    console.log(`Hello, ${ foo }!`); // Hello, World!
    ```
* 在插值表达式中可以调用函数和方法
    ```js
    function capitalize(word) { 
     return `${ word[0].toUpperCase() }${ word.slice(1) }`; 
    } 
    console.log(`${ capitalize('hello') }, ${ capitalize('world') }!`); // Hello, World!
    ```
* 模板也可以插入自己之前的值
    ```js
    let value = ''; 
    function append() { 
       value = `${value}abc`; 
       console.log(value);
    } 
    append(); // abc 
    append(); // abcabc 
    append(); // abcabcabc
    ```

##### 模板字面量标签函数

模板字面量也支持定义**标签函数**（tag function），而通过标签函数可以自定义插值行为。**标签函数
会接收被插值记号分隔后的模板和对每个表达式求值的结果**。标签函数接收到的参数依次是原始字符串数组和对每个表达式求值的结果

```js
const person = "Mike";
const age = 28;

function myTag(strings, personExp, ageExp) {
    const str0 = strings[0]; // "That "
    const str1 = strings[1]; // " is a "
    const str2 = strings[2]; // "."

    const ageStr = ageExp > 99 ? "centenarian" : "youngster";

    // 我们甚至可以返回使用模板字面量构建的字符串
    return `${str0}${personExp}${str1}${ageStr}${str2}`;
}

const output = myTag`That ${person} is a ${age}.`;

console.log(output); // That Mike is a youngster.

```

```js
let a = 6; 
let b = 9; 
function simpleTag1(strings, aValExpression, bValExpression, sumExpression) { 
   console.log(strings);         // ["", " + ", " = ", ""]
   console.log(aValExpression);  // 6
   console.log(bValExpression);  // 9
   console.log(sumExpression);   // 15 ==> 6+9
   return 'foobar'; 
} 

function simpleTag2(strings, ...expressions) { 
   console.log(strings); 
   for(const expression of expressions) { 
     console.log(expression);  // 与 simpleTag 中打印一样
   } 
   return 'foobar'; 
}

let untaggedResult = `${ a } + ${ b } = ${ a + b }`; 
let taggedResult1 = simpleTag1`${ a } + ${ b } = ${ a + b }`; 
let taggedResult2 = simpleTag2`${ a } + ${ b } = ${ a + b }`; 
console.log(untaggedResult); // "6 + 9 = 15" 
console.log(taggedResult1); // "foobar"
console.log(taggedResult2); // "foobar"
```

标签不必是普通的标识符，你可以使用任何[优先级](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Operator_precedence#%e6%b1%87%e6%80%bb%e8%a1%a8)大于 16 的表达式，包括[属性访问](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/Property_accessors)、函数调用、[new 表达式](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/new)，甚至其他带标签的模板字面量。

```js
console.log`Hello`; // [ 'Hello' ]
console.log.bind(1, 2)`Hello`; // 2 [ 'Hello' ]

function recursive(strings, ...values) {
    console.log(strings, values);
    return recursive;
}
recursive`Hello``World`;
// [ 'Hello' ] []
// [ 'World' ] []

```

虽然语法从技术上允许这么做，但不带标签的模板字面量是字符串，并且在链式调用时会抛出 `TypeError`

```js
console.log`Hello`;
console.log(`Hello``World`); // TypeError: "Hello" is not a function
```

##### 原始字符串

使用模板字面量也可以直接获取原始的模板字面量内容（如换行符或 Unicode 字符），而不是被转
换后的字符表示。为此，可以使用默认的 `String.raw` 标签函数：

```js
function tag(strings) {
  console.log(strings.raw[0]);
}

tag`string text line 1 \n string text line 2`;
// logs "string text line 1 \n string text line 2" ,
// including the two characters '\' and 'n'
```

使用 `String.raw()` 方法创建原始字符串和使用默认模板函数和字符串连接创建是一样的

```js
const str = String.raw`Hi\n${2+3}!`;
console.log(str);                      // "Hi\\n5!"
console.log(str.length);               // 6
console.log(str.split('').join(','));  // "H,i,\\,n,5,!"
```

```js
// Unicode 示例
// \u00A9 是版权符号
console.log(`\u00A9`);                               // © 
console.log(String.raw`\u00A9`);                     // \u00A9 

// 换行符示例
console.log(`first line\nsecond line`); 
// first line 
// second line 
console.log(String.raw`first line\nsecond line`);    // "first line\nsecond line" 

// 对实际的换行符来说是不行的
// 它们不会被转换成转义序列的形式
console.log(`first line
second line`); 
// first line 
// second line 

console.log(String.raw`first line 
second line`); 
// first line 
// second line
```

也可以通过标签函数的第一个参数，即字符串数组的.raw 属性取得每个字符串的原始内容：

```js
function printRaw(strings) { 
   console.log('Actual characters:');     // Actual characters:
   for (const string of strings) { 
       console.log(string); 
   } 
   // © 
   //（换行符）
   
   console.log('Escaped characters;');   // Escaped characters: 
   for (const rawString of strings.raw) { 
       console.log(rawString); 
   } 
   // \u00A9 
   // \n
} 
printRaw`\u00A9${ 'and' }\n`; 
```

##### 带标签的模板字面量及转义序列

* `\` 后跟 0 以外的任何十进制数字，或 `\0` 后跟一个十进制数字，例如 `\9` 和 `\07`（这是一种已弃用的语法）
* `\x` 后跟两位以下十六进制数字，例如 `\xz`
* `\u` 后不跟 {，并且后跟四个以下十六进制数字，例如 \uz
* `\u{}` 包含无效的 Unicode 码点——包含一个非十六进制数字，或者它的值大于 10FFFF，例如 `\u{110000}` 和 `\u{z}`

> `\` 后面跟着其他字符，虽然它们可能没有用，因为没有转义，但它们不是语法错误。

#### Symbol

[Symbol](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Symbol)（符号）是 ECMAScript 6 新增的数据类型；`Symbol()` 函数会返回 `symbol` 类型的值，该类型具有静态属性和静态方法。它的静态属性会暴露几个内建的成员对象；它的静态方法会暴露全局的 `symbol` 注册，且类似于内建对象类，但作为构造函数来说它并不完整，因为它不支持语法：`new Symbol()`。

> **每个从 `Symbol()` 返回的 `symbol` 值都是唯一的**。一个 `symbol` 值能作为对象属性的标识符；这是该数据类型仅有的目的。

##### 基本用法

使用 `Symbol()` 函数初始化。因为符号本身是原始类型，所以 `typeof` 操作符对符号返回 `symbol`。

```js
const mySymbol = Symbol();
console.log(typeof mySymbol); // symbol

const symbol1 = Symbol();
const symbol2 = Symbol(42);
const symbol3 = Symbol('foo');

console.log(typeof symbol1);                    // 输出: "symbol"
console.log(symbol2 === 42);                    // 输出: false
console.log(symbol3.toString());                // 输出: "Symbol(foo)"
console.log(Symbol('foo') === Symbol('foo'));   // 输出: false
```

**`Symbol()` 函数不能与 `new` 关键字一起作为构造函数使用**，使用 `new` 运算符的语句将抛出 `TypeError` 错误

```js
const sym = new Symbol(); // Uncaught TypeError: Symbol is not a constructor
```

##### 全局共享的 Symbol

如果**运行时的不同部分需要共享和重用符号实例**，要创建跨文件可用的 symbol，甚至跨域（每个都有它自己的全局作用域），**使用 `Symbol.for()` 方法和 `Symbol.keyFor()` 方法从全局的 `symbol` 注册表设置和取得 `symbol`**:

```js
const fooGlobalSymbol = Symbol.for('foo'); 
console.log(typeof fooGlobalSymbol); // symbol
```

**`Symbol.for()` 对每个字符串键都执行幂等操作**

* 第一次使用某个字符串调用时，它会检查全局运行时注册表，发现不存在对应的`Symbol`，就会生成一个新 `Symbol` 实例并添加到注册表中。
* 使用相同字符串的调用同样会检查注册表，发现存在与该字符串对应的 `Symbol`，然后就会返回该 `Symbol` 实例。
    ```js
    const fooGlobalSymbol = Symbol.for('foo');               // 创建新符号
    const otherFooGlobalSymbol = Symbol.for('foo');          // 并不是重新创建，而是重用已有Symbol
    console.log(fooGlobalSymbol === otherFooGlobalSymbol); // true
    ```
* 即使采用相同的 `Symbol` 描述，在全局注册表中定义的 `Symbol` 跟使用 `Symbol()` 定义的符号也并不等同
    ```js
    let localSymbol = Symbol('foo'); 
    let globalSymbol = Symbol.for('foo'); 
    console.log(localSymbol === globalSymbol); // false
    ```
* **全局注册表中的 `Symbol` 必须使用字符串键来创建**，因此作为参数传给 `Symbol.for()` 的任何值都会被转换为字符串；此外，注册表中使用的键同时也会被用作 `Symbol` 描述
    ```js
    const emptyGlobalSymbol = Symbol.for(); 
    console.log(emptyGlobalSymbol); // Symbol(undefined)
    ```
* **还可以使用 `Symbol.keyFor()` 来查询全局注册表**，这个方法接收符号，返回该全局符号对应的字
    符串键。如果查询的不是全局符号，则返回 `undefined`。
    ```js
    // 创建全局符号
    let s = Symbol.for('foo'); 
    console.log(Symbol.keyFor(s)); // foo 
    // 创建普通符号
    let s2 = Symbol('bar'); 
    console.log(Symbol.keyFor(s2)); // undefined 
    ```
    如果传给 `Symbol.keyFor()` 的不是 Symbol，则该方法抛出 `TypeError`：
    ```js
    Symbol.keyFor(123); // TypeError: 123 is not a symbol
    ```

##### 使用 Symbol 作为属性

**凡是可以使用字符串或数值作为属性的地方，都可以使用 Symbol**。这就包括了对象字面量属性和
`Object.defineProperty()`/`Object.defineProperties()`定义的属性。对象字面量只能在计算属
性语法中使用 Symbol 作为属性。

```js
let s1 = Symbol('foo'), 
    s2 = Symbol('bar'), 
    s3 = Symbol('baz'), 
    s4 = Symbol('qux'); 
let o = { 
   [s1]: 'foo val',
}; 

// 这样也可以：o[s1] = 'foo val'; 
console.log(o);                   // {Symbol(foo): foo val}
Object.defineProperty(o, s2, {value: 'bar val'}); 
console.log(o);                   // {Symbol(foo): foo val, Symbol(bar): bar val}

Object.defineProperties(o, { 
   [s3]: {value: 'baz val'}, 
   [s4]: {value: 'qux val'} 
}); 
console.log(o); 
// {Symbol(foo): foo val, Symbol(bar): bar val, 
// Symbol(baz): baz val, Symbol(qux): qux val}
```

* `Object.getOwnProperty.Symbols()` 返回对象实例的 Symbol 属性数组
* `Object.getOwnProperty.Descriptors()`会返回同时包含常规和 Symbol 属性描述符的对象
* `Reflect.ownKeys()` 会返回两种类型的键

```js
let s1 = Symbol('foo'), 
    s2 = Symbol('bar'); 
let o = { 
   [s1]: 'foo val', 
   [s2]: 'bar val', 
   baz: 'baz val', 
   qux: 'qux val' 
}; 
console.log(Object.getOwnPropertySymbols(o));     // [Symbol(foo), Symbol(bar)] 
console.log(Object.getOwnPropertyNames(o));       // ["baz", "qux"] 
console.log(Object.getOwnPropertyDescriptors(o)); // {baz: {...}, qux: {...}, Symbol(foo): {...}, Symbol(bar): {...}} 
console.log(Reflect.ownKeys(o));                  // ["baz", "qux", Symbol(foo), Symbol(bar)]
```

##### Symbol 类型转换

* 将一个 `symbol` 值转换为一个 `number` 值时，会抛出一个 `TypeError` 错误 (e.g. +sym or sym | 0).
* 使用宽松相等时，`Object(sym) == sym returns true`.
* 一个 symbol 值隐式地创建一个新的 `string` 类型的属性名。例如，Symbol("foo") + "bar" 将抛出一个 TypeError (can't convert symbol to string).
* "safer" String(sym) conversion 的作用会像 `symbol` 类型调用 `Symbol.prototype.toString()` 一样，但是注意 `new String(sym)` 将抛出异常。

```js
// 创建一个 symbol 值
const sym = Symbol("foo");

// 将一个 symbol 值转换为一个 number 值时，会抛出一个 TypeError 错误
console.log(+sym);                      // 抛出 TypeError: Cannot convert a Symbol value to a number
console.log(sym | 0);                   // 抛出 TypeError: Cannot convert a Symbol value to a number

// 使用宽松相等时，Object(sym) == sym returns true
console.log(Object(sym) == sym);        // 输出: true

// 一个 symbol 值不能隐式地创建一个新的 string 类型的属性名
console.log(Symbol("foo") + "bar");    // 抛出 TypeError: Cannot convert a Symbol value to a string

// "safer" String(sym) conversion 的作用会像 symbol 类型调用 Symbol.prototype.toString() 一样
console.log(String(sym));              // 输出: "Symbol(foo)"

// 但是注意 new String(sym) 将抛出异常
console.log(new String(sym));          // 抛出 TypeError: Cannot convert a Symbol value to a string
```

> Tips
>
> *   当使用 `JSON.stringify()` 时，以 `symbol` 值作为键的属性会被完全忽略
>     ```js
>     console.log(JSON.stringify({ [Symbol("foo")]: "foo" })); // '{}'
>     console.log(JSON.stringify({ [Symbol("foo").toString()]: "foo" }));  // '{"Symbol(foo)":"foo"}'
>     ```
> *   Symbols 在 `for...in` 迭代中不可枚举。另外，`Object.getOwnPropertyNames()` 不会返回 `symbol` 对象的属性，但是你能使用 `Object.getOwnPropertySymbols()` 得到它们。
>     ```js
>     const obj = {};
>     
>     obj[Symbol("a")] = "a";
>     obj[Symbol.for("b")] = "b";
>     obj["c"] = "c";
>     obj.d = "d";
>     
>     for (var i in obj) {
>         console.log(i); // logs "c" and "d"
>     }
>     ```

#### Object

在 JavaScript 中，几乎所有的对象都是 Object 的实例，ECMAScript 中的对象其实就是一组数据和功能的集合。对象通过 `new` 操作符后跟对象类型的名称来创建。

```js
const o = new Object();
```

ECMAScript 中的 `Object` 也是派生其他对象的基类。`Object` 类型的所有属性和方法在派生的对象上同样存在。
每个 Object 实例都有如下属性和方法。

* `constructor`：用于创建当前对象的函数。
* `hasOwnProperty(propertyName)`：**用于判断当前对象实例（不是原型）上是否存在给定的属性**。要**检查的属性名必须是字符串**（如 o.hasOwnProperty("name")）或 Symbol。
* `isPrototypeOf(object)`：用于判断**当前对象是否为另一个对象的原型**。
* `propertyIsEnumerable(propertyName)`：用于判断给定的属性是否可以使用 `for-in` 语句枚举。与 `hasOwnProperty()` 一样，**属性名必须是字符串**。
* `toLocaleString()`：返回对象的字符串表示，该**字符串反映对象所在的本地化执行环境**。
* `toString()`：返回**对象的字符串**表示。
* `valueOf()`：返回**对象对应的字符串、数值或布尔值表示**。通常与 `toString()` 的返回值相同。

##### 对象强制转换

* 对象则按原样返回。
    ```js
    const obj = {name: "obj"};
    console.log(Object(obj) === obj);  // 输出: true
    ```
* undefined 和 null 则抛出 TypeError。
    ```js
    console.log(Object(undefined));   // 抛出 TypeError: Cannot convert undefined or null to object
    console.log(Object(null));        // 抛出 TypeError: Cannot convert undefined or null to object
    ```
*   Number、String、Boolean、Symbol、BigInt 等基本类型被封装成其对应的基本类型对象。
    ```js
    // Number
    const num = 123;
    console.log(Object(num));      // Number {123}
    
    // String
    const str = "hello";
    console.log(Object(str));      // String {'hello'}
    
    // Boolean
    const bool = true;
    console.log(Object(bool));     // Boolean {true}
    
    // Symbol
    const sym = Symbol("foo");
    console.log(Object(sym));      // Symbol {Symbol(foo), description: 'foo'}
    
    // BigInt
    const bigInt = BigInt(123);
    console.log(Object(bigInt));   // BigInt {123n}
    ```

## 操作符

ECMA-262 描述了一组可用于操作数据值的操作符，包括数学操作符（如加、减）、位操作符、关
系操作符和相等操作符等

### 一元操作符

只操作一个值的操作符叫做一元操作符

#### 递增、递减操作符

照搬 C 语言的，所以有前缀版和后缀版。顾名思义，**前缀版就是位于要操作的变量前头**，**后缀版就是位于要操作的变量后头**。前缀递增操作符会给数值加 1，把两个加号（++）放到变量前头：

```js
let age = 20;

// 前缀递增操作符把 age 的值变成了 21（给之前的值 20 加 1）
++age; // ==> age = age + 1;
```

前缀递减操作符也类似，只不过是从一个数值减 1。使用前缀递减操作符，只要把两个减号（--）放到变量前头:

```js
let age = 20;
--age; // ==> age = age - 1;
```

**无论使用前缀递增还是前缀递减操作符，变量的值都会在语句被求值之前改变**

```js
let age = 29; 
const anotherAge = --age + 2; // ==> 递减操作先发生，所以 age 的值先变成 28，然后再加 2，结果是 30
console.log(age);             // 28 
console.log(anotherAge);      // 30
```

**前缀递增和递减在语句中的优先级是相等的，因此会从左到右依次求值**。比如：

```js
let num1 = 2; 
let num2 = 20; 
let num3 = --num1 + num2; // ==> num1 先减 1 后的结果再加 num2 ==> 1+20 ==> 21
let num4 = num1 + num2;   // ==> 因为上一步的计算中，num1 已经被改变为 1，num2 没有改变，所以 num1 + num2 = 1 + 20 ==> 21
console.log(num3);        // 21 
console.log(num4);        // 21
```

递增和递减的后缀版语法一样（分别是++和--），只不过要放在变量后面。后缀版与前缀版的主要
区别在于，**后缀版递增和递减在语句被求值后才发生**

* 当操作是唯一时，结果就没什么变化，比如：
    ```js
    let age = 18;
    ++age; // age++
    console.log(age); // 19
    ```
* 当存在混合操作时，差异就比较明显，比如：
    ```js
    let num1 = 2; 
    let num2 = 20; 
    let num3 = num1-- + num2; // 因为递减是整条语句求值后才发生，所以计算步骤为：先计算 num1 + num2 的和，然后再执行 num1-- ==> 2 + 20 ==> 22 然后 num1 - 1 ==> num1=1, num3=22
    let num4 = num1 + num2;   // 因为上一条语句执行后 num1 的值为 1 ==> 1 + 20 ==> 21
    console.log(num3); // 22 
    console.log(num4); // 21
    ```

前置加加、后置加加、前置减减、后置减减这 4 个操作符可以作用于任何值，意思是不限于整数——字符串、布尔值、浮点值，甚至对象都可以。递增和递减操作符遵循如下规则。

* 对于字符串，如果是有效的数值形式，则转换为数值再应用改变。变量类型从字符串变成数值。
    ```js
    let strNum = "2";
    console.log(++strNum);        // 3
    console.log(typeof strNum);   // number
    ```
* 对于字符串，如果不是有效的数值形式，则将变量的值设置为 `NaN` 。变量类型从字符串变成数值。
    ```js
    let str = "hello";
    console.log(++str);        // NaN
    console.log(typeof str);   // number
    ```
* 对于布尔值，如果是 `false`，则转换为 0 再应用改变。变量类型从布尔值变成数值。
    ```js
    let boolFalse = false;
    console.log(++boolFalse);       // 1
    console.log(typeof boolFalse);  // number
    ```
* 对于布尔值，如果是 `true`，则转换为 1 再应用改变。变量类型从布尔值变成数值。
    ```js
    let boolTrue = true;
    console.log(++boolTrue);        // 2
    console.log(typeof boolTrue);   // number
    ```
* 对于浮点值，加 1 或减 1。
    ```js
    let floatNum = 1.1;
    console.log(++floatNum);   // 输出: 2.1
    ```
* 如果是对象，则调用其 `valueOf()` 方法取得可以操作的值。对得到的值应用上述规则。如果是 `NaN`，则调用 `toString()` 并再次应用其他规则。变量类型从对象变成数值。
    ```js
    let obj = {
      valueOf: function() {
        return 2;
      }
    };
    console.log(++obj);       // 3
    console.log(typeof obj);  // number
    ```

#### 一元加和减

它们在 ECMAScript 中跟在高中数学中的用途一样。**一元加由一个加号（+）表示，放在变量前头，对数值没有任何影响**

```js
let num = 25; 
num = +num; 
console.log(num); // 25
```

**如果将一元加应用到非数值，则会执行与使用 `Number()` 转型函数一样的类型转换**：布尔值 `false` 和 `true` 转换为 0 和 1，字符串根据特殊规则进行解析，对象会调用它们的 `valueOf()` 和/或 `toString()`方法以得到可以转换的值。

```js
let s1 = "01"; 
s1 = +s1; 
console.log(s1); // 值变成数值 1 

let s2 = "1.1";
s2 = +s2; 
console.log(s2); // 值变成数值 1.1 

let s3 = "z"; 
s3 = +s3; 
console.log(s3); // 值变成 NaN

let b = false; 
b = +b; 
console.log(b)   // 值变成数值 0 

let f = 1.1; 
f = +f; 
console.log(f);  // 不变，还是 1.1 

let o = { 
   valueOf() { 
     return -1; 
   } 
}; 
o = +o; 
console.log(o);  // 值变成数值 -1
```

一元减由一个减号（-）表示，放在变量前头，主要用于把数值变成负值，如把 1 转换为 -1。

```js
let num = 25; 
num = -num; 
console.log(num); // -25
```

对数值使用一元减会将其变成相应的负值。在应用到非数值时，一元减会遵循与一元加同样的规则，先对它们进行转换，然后再取负值：

```js
let s1 = "01"; 
s1 = -s1; 
console.log(s1);  // 值变成数值 -1 

let s2 = "1.1"; 
s2 = -s2; 
console.log(s2);  // 值变成数值 -1.1 

let s3 = "z"; 
s3 = -s3; 
console.log(s3);  // 值变成 NaN 

let b = false; 
b = -b; 
console.log(b);   // 值变成数值 0 

let f = 1.1; 
f = -f; 
console.log(f);   // 变成 -1.1 

let o = { 
   valueOf() { 
     return -1; 
   } 
}; 
o = -o; 
console.log(o);   // 值变成数值 1
```

### 位操作符

。ECMAScript
中的所有数值都以 IEEE 754 64 位格式存储，但位操作并不直接应用到 64 位表示，而是先把值转换为
32 位整数，再进行位操作，之后再把结果转换为 64 位。

> 对开发者而言，就好像只有 32 位整数一样，因为 64 位整数存储格式是不可见的。

**有符号整数使用 32 位的前 31 位表示整数值；第 32 位表示数值的符号**。如 0 表示正，1 表示负。这一位称为符号位（sign bit），它的值决定了数值其余部分的格式。正值以真正的二进制格式存储，即 31位中的每一位都代表 2 的幂。第一位（称为第 0 位）表示 20，第二位表示 21，依此类推。如果一个位是空的，则以0填充，相当于忽略不计。

> **默认情况下，ECMAScript 中的所有整数都表示为有符号数**。不过，确实存在无符号整数。**对无符号整数来说，第 32 位不表示符号，因为只有正值**。无符号整数比有符号整数的范围更大，因为符号位被用来表示数值了。

比如，数值18的二进制格式为 `00000000000000000000000000010010`，或更精简的 `10010`。后者是用到的 5 个有效位，决定了实际的值:

<img src="https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/97d44c20fc8d43aaa77798ac58c73457~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=748&#x26;h=370&#x26;s=27657&#x26;e=png&#x26;b=ffffff" alt="进制图" height="160">

负值以一种称为二补数（或补码）的二进制编码存储。一个数值的二补数通过如下 3 个步骤计算得到：

1.  确定绝对值的二进制表示（如，对于18，先确定 18 的二进制表示）；
2.  找到数值的一补数（或反码），换句话说，就是每个 0 都变成 1，每个 1 都变成 0；
3.  给结果加 1。
    基于上述步骤确定 -18 的二进制表示，首先从 18 的二进制表示开始：

```js
0000 0000 0000 0000 0000 0000 0001 0010 
```

然后，计算一补数，即反转每一位的二进制值：

```js
1111 1111 1111 1111 1111 1111 1110 1101 
```

最后，给一补数加 1：

```js
1111 1111 1111 1111 1111 1111 1110 1101 
---------------------------------------
1111 1111 1111 1111 1111 1111 1110 1110 
```

那么，-18 的二进制表示就是 `11111111111111111111111111101110`。要注意的是，在处理有符号整数
时，我们无法访问第 31 位。
ECMAScript 会帮我们记录这些信息。在把负值输出为一个二进制字符串时，我们会得到一个前面加了减号的绝对值，如下所示：

```js
let num = -18; 
console.log(num.toString(2)); // "-10010"
```

在将 -18 转换为二进制字符串时，结果得到 `-10010`。转换过程会求得二补数，然后再以更符合逻辑的形式表示出来。

> 如果将位操作符应用到非数值，那么首先会使用 `Number()` 函数将该值转换为数值（这个过程是自动的），然后再应用位操作；最终结果是数值。

#### 按位非

按位非操作符用波浪符（`~`）表示，它的作用是返回数值的一补数。按位非是 ECMAScript 中为数不多的几个二进制数学操作符之一。

```js
let num1 = 25;     // 二进制 00000000000000000000000000011001 
let num2 = ~num1;  // 二进制 11111111111111111111111111100110 
console.log(num2); // -26
```

按位非操作符作用到了数值 25，得到的结果是 -26。由此可以看出，**按位非的最终效果是对数值取反并减 1**，就像执行如下操作的结果一样：

```js
let num1 = 25; 
let num2 = -num1 - 1; 
console.log(num2); // "-26"
```

#### 按位与

按位与操作符用和号（`&`）表示，有两个操作数。本质上，**按位与就是将两个数的每一个位对齐，然后基于真值表中的规则，对每一位执行相应的与操作**。

| 第一位数值的位 | 第二位数值的位 | 结果 |
| -------------- | -------------- | ---- |
| 1              | 1              | 1    |
| 1              | 0              | 0    |
| 0              | 1              | 0    |
| 0              | 0              | 0    |

*按位与操作在两个位都是 1 时返回 1，在任何一位是 0 时返回 0*；也就是**遇0则0**。

对数值 25 和 3 求与操作，如下所示：

```js
const result = 25 & 3;
console.log('result:', result); // 1
```

25 和 3 的按位与操作的结果是 1。为什么呢？看下面的二进制计算过程：

```js
 25 = 0000 0000 0000 0000 0000 0000 0001 1001 
 3  = 0000 0000 0000 0000 0000 0000 0000 0011 
--------------------------------------------- 
AND = 0000 0000 0000 0000 0000 0000 0000 0001 
```

如上所示，25 和 3 的二进制表示中，只有第 0 位上的两个数都是 1。于是结果数值的所有其他位都会以 0 填充，因此结果就是 1。

#### 按位或

按位或操作符用管道符（`|`）表示，同样有两个操作数。按位或遵循如下真值表：

| 第一位数值的位 | 第二位数值的位 | 结果 |
| -------------- | -------------- | ---- |
| 1              | 1              | 1    |
| 1              | 0              | 1    |
| 0              | 1              | 1    |
| 0              | 0              | 0    |

按位或操作在至少一位是 1 时返回 1，两位都是 0 时返回 0；也就是**遇1则1**。

25 和 3 执行按位或，代码如下所示：

```js
let result = 25 | 3; 
console.log(result); // 27
```

可见 25 和 3 的按位或操作的结果是 27：

```js
 25 = 0000 0000 0000 0000 0000 0000 0001 1001 
 3  = 0000 0000 0000 0000 0000 0000 0000 0011 
--------------------------------------------- 
 OR = 0000 0000 0000 0000 0000 0000 0001 1011 
```

在参与计算的两个数中，有 4 位都是 1，因此它们直接对应到结果上。二进制码 `11011` 等于 27。

#### 按位异或

按位异或用脱字符（^）表示，同样有两个操作数。下面是按位异或的真值表：

| 第一位数值的位 | 第二位数值的位 | 结果 |
| -------------- | -------------- | ---- |
| 1              | 1              | 0    |
| 1              | 0              | 1    |
| 0              | 1              | 1    |
| 0              | 0              | 1    |

按位异或与按位或的区别是，它只在一位上是 1 的时候返回 1（两位都是 1 或 0，则返回 0）；也就是**相同为1否反之为0**。

对数值 25 和 3 执行按位异或操作：

```js
let result = 25 ^ 3; 
console.log(result); // 26 
```

可见，25 和 3 的按位异或操作结果为 26，如下所示：

```js
 25 = 0000 0000 0000 0000 0000 0000 0001 1001 
 3  = 0000 0000 0000 0000 0000 0000 0000 0011 
--------------------------------------------- 
XOR = 0000 0000 0000 0000 0000 0000 0001 1010 
```

两个数在 4 位上都是 1，但两个数的第 0 位都是 1，因此那一位在结果中就变成了 0。其余位上的 1
在另一个数上没有对应的 1，因此会直接传递到结果中。二进制码 `11010` 等于 26。（注意，这比对同样
两个值执行按位或操作得到的结果小 1。）

#### 左移

左移操作符用两个小于号（`<<`）表示，会按照指定的位数将数值的所有位向左移动。

比如，数值 2（二进制 `10`）向左移 5 位，就会得到 64（二进制 `1000000`），如下所示：

```js
let oldValue = 2;             // 等于二进制 10 
let newValue = oldValue << 5; // 等于二进制 1000000，即十进制 64
```

*在移位后，数值右端会空出 5 位。左移会以 0 填充这些空位，让结果是完整的 32 位数值*。

![进制左移](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f1ff596073f4436c96da9ba0608f518a~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1672\&h=524\&s=79727\&e=png\&b=ffffff)
**左移会保留它所操作数值的符号**。比如，如果2 左移 5 位，将得到64，而不是正 64。

#### 有符号右移

有符号右移由两个大于号（`>>`）表示，会将数值的所有 32 位都向右移，同时保留符号（正或负）。**有符号右移实际上是左移的逆运算**。

比如，如果将 64 右移 5 位，那就是 2：

```js
let oldValue = 64;            // 等于二进制 1000000 
let newValue = oldValue >> 5; // 等于二进制 10，即十进制 2
console.log(newValue);        // 2
```

移位后就会出现空位。不过，右移后空位会出现在左侧，且在符号位之后。ECMAScript 会用符号位的值来填充这些空位，以得到完整的数值。

![image](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/8e0fbc08ecf44f35a4f8210b7c7c2fb1~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1624\&h=526\&s=68411\&e=png\&b=ffffff)

#### 无符号右移

无符号右移用 3 个大于号表示（`>>>`），会将数值的所有 32 位都向右移。对于正数，无符号右移与有符号右移结果相同。

仍然以前面有符号右移的例子为例，64 向右移动 5 位，会变成 2：

```js
let oldValue = 64;              // 等于二进制 00000000000000000000000001000000 
let newValue = oldValue >>> 5;  // 等于二进制 00000000000000000000000000000010
console.log(newValue);          // 等于十进制 2
```

对于负数，有时候差异会非常大。与有符号右移不同，无符号右移会给空位补 0，而不管符号位是什么。对正数来说，这跟有符号右移效果相同。但对负数来说，结果就差太多了。**无符号右移操作符将负数的二进制表示当成正数的二进制表示来处理**。

```js
let oldValue = -64;             // 等于二进制 11111111111111111111111111000000 
let newValue = oldValue >>> 5;  // 等于二进制 00000111111111111111111111111110
console.log(newValue);          // 等于十进制 134217726
```

### 布尔操作符

布尔操作符一共有 3 个：**逻辑非**、**逻辑与**和**逻辑或**。

#### 逻辑非

逻辑非操作符由一个叹号（`!`）表示，可应用给 ECMAScript 中的任何值。这个操作符始终返回布尔值，无论应用到的是什么数据类型。逻辑非操作符首先将操作数转换为布尔值，然后再对其取反。换句话说，逻辑非操作符会遵循如下规则。

*   如果操作数是对象，则返回 `false`。
    ```js
    const obj = {};
    console.log(!obj); // false
    ```
*   如果操作数是空字符串，则返回 `true`。
    ```js
    const emptyStr = "";
    console.log(!emptyStr); // true
    ```
*   如果操作数是非空字符串，则返回 `false`。
    ```js
    const str = "this is string";
    console.log(!str); // false
    ```
*   如果操作数是数值 0，则返回 `true`。
    ```js
    const zero = 0;
    console.log(!zero); // true
    ```
*   如果操作数是非 0 数值（包括 `Infinity`），则返回 `false`。
    ```js
    const one = 1;
    const infin = Infinity;
    console.log(!one, !infin); // false false
    ```
*   如果操作数是 `null`，则返回 `true`。
    ```js
    const test = null;
    console.log(!test); // true
    ```
*   如果操作数是 `NaN`，则返回 `true`。
    ```js
    const num = Number("not number"); // NaN
    console.log(!num);
    ```
*   如果操作数是 `undefined`，则返回 `true`。

    ```js
    const unde = undefined;
    console.log(!unde); // true
    ```

    **逻辑非操作符也可以用于把任意值转换为布尔值**。*同时使用两个叹号（`!!`），相当于调用了转型函
    数 `Boolean()`*。无论操作数是什么类型，第一个叹号总会返回布尔值。第二个叹号对该布尔值取反，
    从而给出变量真正对应的布尔值。

```js
console.log(!!"blue"); // true 
console.log(!!0);      // false 
console.log(!!NaN);    // false 
console.log(!!"");     // false 
console.log(!!12345);  // true
```

#### 逻辑与

逻辑与操作符由两个和号（`&&`）表示

```js
let result = true && false;
console.log(result); // false
```

逻辑与操作符遵循如下真值表：

| 第一个操作数 | 第二个操作数 | 结果  |
| ------------ | ------------ | ----- |
| true         | true         | true  |
| true         | false        | false |
| false        | true         | false |
| false        | false        | false |

**逻辑与操作符可用于任何类型的操作数，不限于布尔值**。如果有操作数不是布尔值，则逻辑与并不一定会返回布尔值，而是遵循如下规则。

*   如果第一个操作数是对象，则返回第二个操作数。
    ```js
    const obj = {};
    const num = 1314;
    const result = obj && num;
    console.log(result); // 1314
    ```
*   如果第二个操作数是对象，则只有第一个操作数求值为 `true` 才会返回该对象。
    ```js
    const isGreaterThanZero = Math.PI > 0;
    const result = isGreaterThanZero && {};
    console.log(result); // {}
    ```
*   如果两个操作数都是对象，则返回第二个操作数。
    ```js
    const user = {name: 'forest'};
    const works = {desc: 'developer'};
    const result = user && works;
    console.log(result); // {desc: 'developer'}
    ```
*   如果有一个操作数是 `null`，则返回 `null`。
    ```js
    const obj = {};
    const templete = null;
    const result = obj && templete;
    console.log(result); // null
    ```
*   如果有一个操作数是 `NaN`，则返回 `NaN`。
    ```js
    const obj = {};
    const result = {} && NaN;
    console.log(result); // NaN
    ```
*   如果有一个操作数是 `undefined`，则返回 `undefined`。
    ```js
    const str = "";
    let un;
    const result = str && un;
    console.log(result); // undefined
    ```

**逻辑与操作符是一种短路操作符**，意思就是*如果第一个操作数决定了结果，那么永远不会对第二个
操作数求值*。

```js
let found = true; 
let result = found && someUndeclaredVariable; // 这里会出错，因为 someUndeclaredVariable 没有事先声明
console.log(result); // 不会执行这一行
```

如果把 `found` 的值改为 `false`，尽管 someUndeclaredVariable 没有声明也不会报错，因为它不会执行到逻辑与后面去；具体代码如下：

```js
let found = false; 
let result = found && someUndeclaredVariable; // 不会报错
console.log(result);                          // 会执行到这行代码
```

#### 逻辑或

逻辑或操作符由两个管道符（`||`）表示，比如：

```js
let result = true || false;
console.log(result); // true
```

逻辑或操作符遵循如下真值表：

| 第一个操作数 | 第二个操作数 | 结果  |
| ------------ | ------------ | ----- |
| true         | true         | true  |
| true         | false        | true  |
| false        | true         | true  |
| false        | false        | false |

与逻辑与类似，如果有一个操作数不是布尔值，那么逻辑或操作符也不一定返回布尔值。它遵循如下规则。
如果第一个操作数是对象，则返回第一个操作数。

*   如果第一个操作数求值为 `false`，则返回第二个操作数。
    ```js
    const obj = {};
    const str = "this is a string";
    const result = obj || str;
    console.log(result); // {}
    ```
*   如果两个操作数都是对象，则返回第一个操作数。
    ```js
    const user = {name: 'forest'};
    const works = {desc: 'developer'};
    const result = user || works;
    console.log(result); // {name: 'forest'}
    ```
*   如果两个操作数都是 `null`，则返回 `null`。
    ```js
    const val1 = null;
    const val2 = null;
    const result = val1 || val2;
    console.log(result); // null
    ```
*   如果两个操作数都是 `NaN`，则返回 `NaN`。
    ```js
    console.log(NaN || NaN); // NaN
    ```
*   如果两个操作数都是 `undefined`，则返回 `undefined`。
    ```js
    let name, str;
    const result = name || str;
    console.log(result); // undefined
    ```

### 乘性操作符
ECMAScript 定义了 3 个乘性操作符：乘法、除法和取模。如果乘性操作符有不是数值的操作数，则该操作数会在后台被使用 `Number()` 转型函数转换为数值。这意味着空字符串会被当成 0，而布尔值 `true` 会被当成 1。
#### 乘法操作符
乘法操作符由一个星号（`*`）表示，可以用于计算两个数值的乘积。
```js
const result = 45 * 2;
```
乘法操作符在处理特殊值时也有一些特殊的行为:
- 如果操作数都是数值，则执行常规的乘法运算，即两个正值相乘是正值，两个负值相乘也是正 值，正负符号不同的值相乘得到负值。如果 ECMAScript 不能表示乘积，则返回 `Infinity` 或 `-Infinity`。
  ```js
  // 两个正值相乘得到正值
  const result1 = 2 * 3;                // result1 = 6
  
  // 两个负值相乘得到正值
  const result2 = -4 * -5;              // result2 = 20
  
  // 正负符号不同的值相乘得到负值
  const result3 = -6 * 7;               // result3 = -42
  
  // 乘积超出 ECMAScript 能表示的范围，返回 Infinity
  const result4 = Number.MAX_VALUE * 2; // result4 = Infinity
  ```
- 如果有任一操作数是 `NaN`，则返回 `NaN`。
  ```js
  const result = NaN * 10; // result = NaN
  ```
- 如果是 `Infinity` 乘以 0，则返回 `NaN`。 
  ```js
  const result = Infinity * 0; // result = NaN
  ```
- 如果是 `Infinity` 乘以非 0 的有限数值，则根据第二个操作数的符号返回 `Infinity` 或 `-Infinity`。 
  ```js
  // Infinity 乘以非 0 的有限数值，根据符号返回 Infinity 或 -Infinity
  const result1 = Infinity * 10; // result1 = Infinity
  const result2 = Infinity * -5; // result2 = -Infinity
  ```
- 如果是 `Infinity` 乘以 `Infinity`，则返回 `Infinity`。 
  ```js
  const result = Infinity * Infinity; // result = Infinity
  ```
- 如果有不是数值的操作数，则先在后台用 `Number()` 将其转换为数值，然后再应用上述规则。

#### 除法操作符
除法操作符由一个斜杠（`/`）表示，用于计算第一个操作数除以第二个操作数的商
```js
const result = 44 / 22;
```
除法操作符针对特殊值也有一些特殊的行为:
- 如果操作数都是数值，则执行常规的除法运算，即两个正值相除是正值，两个负值相除也是正 值，符号不同的值相除得到负值。如果 ECMAScript 不能表示商，则返回 `Infinity` 或 `-Infinity`。 
  ```js
  // 两个正值相除得到正值
  const positiveDivision = 10 / 2; // positiveDivision = 5
  
  // 两个负值相除得到正值
  const negativeDivision = -20 / -4; // negativeDivision = 5
  
  // 符号不同的值相除得到负值
  const mixedSignsDivision = -30 / 6; // mixedSignsDivision = -5
  
  // 商超出 ECMAScript 能表示的范围，返回 Infinity
  const largeDivision = Number.MAX_VALUE / 0.5; // largeDivision = Infinity
  ```
- 如果有任一操作数是 `NaN`，则返回 `NaN`。
  ```js
  const nanDivision = NaN / 10; // nanDivision = NaN
  ```
- 如果是 `Infinity` 除以 `Infinity`，则返回 `NaN`。
  ```js
  const infinityDivision = Infinity / Infinity; // infinityDivision = NaN
  ```
- 如果是 `0` 除以 `0`，则返回 `NaN`。
  ```js
  const zeroDivision = 0 / 0; // zeroDivision = NaN
  ```
- 如果是非 `0` 的有限值除以 `0`，则根据第一个操作数的符号返回 `Infinity` 或 `-Infinity`。 
  ```js
  const nonZeroDivision1 = 50 / 0; // nonZeroDivision1 = Infinity
  const nonZeroDivision2 = -70 / 0; // nonZeroDivision2 = -Infinity
  ```
- 如果是 `Infinity` 除以任何数值，则根据第二个操作数的符号返回 `Infinity` 或 `-Infinity`。
  ```js
  const infinityDivision1 = Infinity / 100; // infinityDivision1 = Infinity
  const infinityDivision2 = Infinity / -5; // infinityDivision2 = -Infinity
  ```
- 如果有不是数值的操作数，则先在后台用 `Number()` 函数将其转换为数值，然后再应用上述规则。

#### 取模操作符
取模（余数）操作符由一个百分比符号（`%`）表示
```js
let result = 26 % 5; // 等于 1 
```
与其他乘性操作符一样，取模操作符对特殊值也有一些特殊的行为。
- 如果操作数是数值，则执行常规除法运算，返回余数。
  ```js
  const normalModulus = 10 % 3; // normalModulus = 1
  ```
- 如果被除数是无限值，除数是有限值，则返回 `NaN`。
  ```js
  const infinityModulus = Infinity % 5; // infinityModulus = NaN
  ```
- 如果被除数是有限值，除数是 0，则返回 `NaN`。
  ```js
  const zeroModulus = 8 % 0; // zeroModulus = NaN
  ```
- 如果是 `Infinity` 除以 `Infinity`，则返回 `NaN`。
  ```js
  const infinityInfinityModulus = Infinity % Infinity; // infinityInfinityModulus = NaN
  ```
- 如果**被除数是有限值，除数是无限值，则返回被除数**。
  ```js
  const finiteInfinityModulus = 15 % Infinity; // finiteInfinityModulus = 15
  ```
- 如果被除数是 0，除数不是 0，则返回 0。
  ```js
  const zeroModulusValue = 0 % 7; // zeroModulusValue = 0
  ```
- 如果有不是数值的操作数，则先在后台用 `Number()` 函数将其转换为数值，然后再应用上述规则。

### 指数操作符
ECMAScript 7 新增了指数操作符，`Math.pow()` 现在有了自己的操作符`**`，结果是一样的：
```js
console.log(Math.pow(3, 2);    // 9  
console.log(3 ** 2);           // 9 
console.log(Math.pow(16, 0.5); // 4 
console.log(16** 0.5);         // 4
```

指数操作符也有自己的指数赋值操作符 `**=`
```js
let squared = 3;
squared **= 2; 
console.log(squared); // 9
let sqrt = 16; 
sqrt **= 0.5; 
console.log(sqrt);    // 4
```

### 加性操作符
加性操作符，即加法和减法操作符，一般都是编程语言中最简单的操作符。不过，在 ECMAScript中，这两个操作符拥有一些特殊的行为。

#### 加法操作符
加法操作符（+）用于求两个数的和，比如：
```js
let result = 1 + 2;
```
如果两个操作数都是数值，加法操作符执行加法运算并根据如下规则返回结果：
- 如果有任一操作数是 `NaN`，则返回 `NaN`；
  ```js
  console.log(1 + NaN); // NaN
  ```
- 如果是 `Infinity` 加 `Infinity`，则返回 `Infinity`；
  ```js
  console.log(Infinity + Infinity); // Infinity
  ```
- 如果是 `-Infinity` 加 `-Infinity`，则返回 `-Infinity`；
  ```js
  console.log(-Infinity + -Infinity); // -Infinity
  ```
- 如果是 `Infinity` 加 `-Infinity`，则返回 `NaN`；
  ```js
  console.log(Infinity + -Infinity); // NaN
  ```
- 如果是 `+0` 加 `+0`，则返回 `+0`；
  ```js
  console.log(+0 + +0); // +0
  ```
- 如果是 `-0` 加 `+0`，则返回 `+0`；
  ```js
  console.log(-0 + +0); // +0
  ```
- 如果是 `-0` 加 `-0`，则返回 `-0`。
  ```js
  console.log(-0 + -0); // -0
  ```

不过，如果有一个**操作数是字符串**，则要应用如下规则：
- *如果两个操作数都是字符串，则将第二个字符串拼接到第一个字符串后面*。
  ```js
  console.log('Java' + 'Script'); // JavaScript
  ```
- *如果只有一个操作数是字符串，则将另一个操作数转换为字符串，再将两个字符串拼接在一起*。
  ```js
  console.log('string' + 12);        // string12
  console.log('string' + true);      // stringtrue
  console.log(NaN + 'string');       // NaNstring
  console.log(null + 'string');      // nullstring
  console.log(undefined + 'string'); // undefinedstring
  console.log({} + 'string');        // [object Object]string
  // console.log(Symbol() + 'string');  // Uncaught TypeError: Cannot convert a Symbol value to a string
  console.log(Symbol().toString() + 'string'); // Symbol()string
  ```
  > 如果有任一操作数是对象、数值或布尔值，则调用它们的 `toString()` 方法以获取字符串，然后再应用前面的关于字符串的规则。对于 `undefined` 和 `null`，则调用 `String()` 函数，分别获取 `"undefined"` 和 `"null"`。
  
### 减法操作符

减法操作符（-）也是使用很频繁的一种操作符，比如：
```js
let result = 2 - 1;
```
与加法操作符一样，减法操作符也有一组规则用于处理 ECMAScript 中不同类型之间的转换。
- 如果两个操作数都是数值，则执行数学减法运算并返回结果。
  ```js
  console.log(24 - 5); // 19 
  ```
- 如果有任一操作数是 `NaN`，则返回 `NaN`。
  ```js
  console.log(NaN - 5); // NaN
  console.log(5 - NaN); // NaN
  ```
- 如果是 `Infinity` 减 `Infinity`，则返回 `NaN`。
  ```js
  console.log(Infinity - Infinity); // NaN
  ```
- 如果是 `-Infinity` 减 `-Infinity`，则返回 `NaN`。
  ```js
  console.log(-Infinity - -Infinity); // NaN
  ```
- 如果是 `Infinity` 减 `-Infinity`，则返回 `Infinity`。
  ```js
  console.log(Infinity - -Infinity); // Infinity
  ```
- 如果是 `-Infinity` 减 `Infinity`，则返回 `-Infinity`。
  ```js
  console.log(-Infinity - Infinity); // -Infinity
  ```
- 如果是 `+0` 减 `+0`，则返回 `+0`。
  ```js
  console.log(+0 - +0); // +0
  ```
- 如果是 `+0` 减 `-0`，则返回 `-0`。
  ```js
  console.log(+0 - -0); // -0
  ```
- 如果是 `-0` 减 `-0`，则返回 `+0`。
  ```js
  console.log(-0 - -0); // +0
  ```
- 如果有任一操作数是字符串、布尔值、`null` 或 `undefined`，则先在后台使用 `Number()` 将其转换为数值，然后再根据前面的规则执行数学运算。如果转换结果是 `NaN`，则减法计算的结果是 `NaN`。
  ```js
  console.log('string' - 10); // NaN
  console.log(true - 2);      // -1 ==> Number(true) - 2 ==> 1 -2
  console.log(null - 1);      // -1 ==> Number(null) - 1 ==> 0 - -1
  console.log(undefined - 2); // NaN
  console.log(Symbol() - 2);  // Uncaught TypeError: Cannot convert a Symbol value to a number
  console.log({} - 1);        // NaN
  ```
- 如果有任一操作数是对象，则调用其 `valueOf()` 方法取得表示它的数值。如果该值是 `NaN`，则减法计算的结果是 `NaN`。如果对象没有 `valueOf()` 方法，则调用其 `toString()` 方法，然后再将得到的字符串转换为数值。

### 关系操作符
关系操作符执行比较两个值的操作，包括小于、大于、小于等于、大于等于。
```js
const lessThan = 3 < 5;
const greaterThan = 5 > 3;
const isGreaterOrEqual = 5 >= 3;
const isLessOrEqual = 5 <= 3;
```
与 ECMAScript 中的其他操作符一样，在将它们应用到不同数据类型时也会发生类型转换和其他行为。
> 在大多数比较的场景中，如果一个值不小于另一个值，那就一定大于或等于它。但在**比较 NaN 时，无论是小于还是大于等于，比较的结果都会返回 false**。
- 如果操作数都是数值，则执行数值比较。
  ```js
  const numericComparison = 5 > 3; // numericComparison = true
  ```
- 如果操作数都是字符串，则**逐个比较字符串中对应字符的 [ASCII](https://zh.wikipedia.org/wiki/ASCII) 编码**。
  ```js
  const stringComparison = 'hello' < 'world'; // stringComparison = true
  ```
- 如果有任一操作数是数值，则将另一个操作数转换为数值，执行数值比较。
  ```js
  const numericStringComparison = 10 < '20'; // numericStringComparison = true
  ```
- 如果有任一操作数是对象，则调用其 `valueOf()` 方法，取得结果后再根据前面的规则执行比较。如果没有 `valueOf()` 操作符，则调用 `toString()` 方法，取得结果后再根据前面的规则执行比较。
  ```js
  // 一个操作数是对象，调用 valueOf() 方法后再比较
  const objectComparison = 10 >= { valueOf: () => 5 };
  console.log('object comparison:', objectComparison); // object comparison: true
  
  // 一个操作数是对象，调用 toString() 方法后再比较
  const objectToStringComparison = '100' >= { toString: () => '50' }; // 调用 toString 方法后，返回字符串 50，然后字符串 100 和字符串 50 进行逐字比较，1 的ASCII码的十进制数为 49，5 的ASCII码的十进制数为 53，所以 49 不大于等于 53，故而结果为 false
  console.log('object to string comparison:', objectToStringComparison); // false
  ```
- **如果有任一操作数是布尔值，则将其转换为数值再执行比较**。
  ```js
  const booleanComparison = true > false; // booleanComparison = true
  ```
  
### 相等操作符
在比较字符串、数值和布尔值是否相等时，过程都很直观。但是在比较两个对象是否相等时，情形就比较复杂了。ECMAScript 中的相等和不相等操作符，原本在比较之前会执行类型转换，但很快就有人质疑这种转换是否应该发生。

#### 等于和不等于

ECMAScript 中的等于操作符用两个等于号（==）表示，如果操作数相等，则会返回 true。不等于操作符用叹号和等于号（!=）表示，如果两个操作数不相等，则会返回 true。这两个操作符都会先进行类型转换（通常称为强制类型转换）再确定操作数是否相等。

在转换操作数的类型时，相等和不相等操作符遵循如下规则。
- 如果任一操作数是布尔值，则将其转换为数值再比较是否相等。false 转换为 0，true 转换为 1。
  ```js
  console.log(false == 1); // false ==> Number(false) ==> 0
  console.log(true == 1);  // true  ==> Number(true)  ==> 1
  ```
- 如果一个操作数是字符串，另一个操作数是数值，则尝试将字符串转换为数值，再比较是否相等。
  ```js
  console.log("abc" == 2); // false ==> Number('a') ==> NaN
  ```
- 如果一个操作数是对象，另一个操作数不是，则调用对象的 `valueOf()` 方法取得其原始值，再根据前面的规则进行比较。
  ```js
  const obj = { valueOf: () => 'abc' };
  console.log(obj == 32); // false ==> obj 调用 valueOf 函数后，返回字符串 abc，然后 abc 再转成数值时的结果为 NaN，故而 NaN 等于 32 为 false
  ```

在进行比较时，这两个操作符会遵循如下规则。
- `null` 和 `undefined` 相等。
  ```js
  console.log(null == undefined); // true ==> undefined 也是 null 的派生类型
  ```
- `null` 和 `undefined` 不能转换为其他类型的值再进行比较。
  ```js
  console.log(null == 0);      // false
  console.log(undefined == 0); // false
  ```
- 如果有任一操作数是 `NaN`，则相等操作符返回 `false`，不相等操作符返回 `true`。记住：即使两个操作数都是 `NaN`，相等操作符也返回 `false`，因为按照规则，`NaN` 不等于 `NaN`。
  ```js
  console.log(NaN == 1);   // false
  console.log(NaN == NaN); // false
  ```
- 如果两个操作数都是对象，则比较它们是不是同一个对象。如果两个操作数都指向同一个对象，则相等操作符返回 `true`。否则，两者不相等。
  ```js
  const obj1 = { key: 'value' };
  const obj2 = obj1;
  const obj3 = { key: 'value' };
  
  console.log(obj1 == obj2); // true
  console.log(obj1 == obj3); // false
  ```


一些特殊情况的结果：
| 表达式            | 结果  |
| ----------------- | ----- |
| null == undefined | true  |
| "NaN" == NaN      | false |
| 5 == NaN          | false |
| NaN == NaN        | false |
| NaN != NaN        | true  |
| false == 0        | true  |
| true == 1         | true  |
| true == 2         | false |
| undefined == 0    | false |
| null == 0         | false |
| "5" == 5          | true  |

#### 全等和不全等
全等操作符由 3 个等于号（`===`）表示，只有**两个操作数在不转换的前提下相等才返回 `true`**
```js
const res  = 55 == '55';    // true  ==> 类型隐式转换后相等
const result = 55 === '55'; // false ==> 因为类型不同
```
不全等操作符用一个叹号和两个等于号（!==）表示，只有**两个操作数在不转换的前提下不相等才
返回 true**。
```js
let res = ("55" != 55);     // false ==> 转换后相等
let result = ("55" !== 55); // true  ==> 不相等，因为数据类型不同
```
> 另外，虽然 `null == undefined` 是 `true`（因为这两个值类似），但 `null === undefined` 是
> `false`，因为它们**不是相同的数据类型**。

### 条件操作符
**条件（三元）运算符是 JavaScript 唯一使用三个操作数的运算符**：一个条件后跟一个问号（`?`），如果条件为真值，则执行冒号（`:`）前的表达式；若条件为假值，则执行最后的表达式。该运算符经常当作 `if...else` 语句的简捷形式来使用
```js
condition ? exprIfTrue : exprIfFalse
```
- `condition` 计算结果用作条件的表达式。
- `exprIfTrue` 如果 `condition` 的计算结果为真值（等于或可以转换为 `true` 的值），则执行该表达式。
- `exprIfFalse` 如果 `condition` 为假值（等于或可以转换为 `false` 的值）时执行的表达式。
> 除了 `false`，可能的假值表达式还有：`null`、`NaN`、`0`、空字符串（`""`）和 `undefined`。如果 `condition` 是其中任何一个，那么条件表达式的结果就是 `exprIfFalse` 表达式执行的结果。

```js
const age = 26;
const beverage = age >= 21 ? "Beer" : "Juice";
console.log(beverage); // "Beer"
```

### 赋值操作符
简单赋值用等于号（`=`）表示，将右手边的值赋给左手边的变量，如下所示：
```js
let num = 10;
```
复合赋值使用乘性、加性或位操作符后跟等于号（`=`）表示。这些赋值操作符是类似如下常见赋值
操作的简写形式：
```js
let num = 10;
num = num + 10;
```
以上代码的第二行可以通过复合赋值来完成：
```js
let num = 10; 
num += 10;
```
每个数学操作符以及其他一些操作符都有对应的复合赋值操作符：
- 乘后赋值（`*=`）
- 除后赋值（`/=`）
- 取模后赋值（`%=`）
- 加后赋值（`+=`）
- 减后赋值（`-=`）
- 左移后赋值（`<<=`）
- 右移后赋值（`>>=`）
- 无符号右移后赋值（`>>>=`）

### 逗号操作符
**逗号操作符可以用来在一条语句中执行多个操作**，如下所示：
```js
let num1 = 1, num2 = 2, num3 = 3; 
```
*在一条语句中同时声明多个变量是逗号操作符最常用的场景*。不过，**也可以使用逗号操作符来辅助赋值**。在赋值时使用逗号操作符分隔值，最终会返回表达式中最后一个值：
```js
let num = (5, 1, 4, 8, 0); 
console.log(num); // num 的值为 0，因为 0 是表达式中最后一项
```
逗号操作符的这种使用场景并不多见，但这种行为的确存在。

## 语句

### if
当指定条件为真，if 语句会执行一段语句。如果条件为假，则执行另一段语句。
```js
let i = 32;
if (i > 25) 
   console.log("Greater than 25.");           // 只有一行代码的语句
else { 
   console.log("Less than or equal to 25.");  // 一个语句块
}
```

连续使用多个 if 语句：
```js
let i = 32;
if (i > 25) { 
   console.log("Greater than 25."); 
} else if (i < 0) { 
   console.log("Less than 0."); 
} else { 
   console.log("Between 0 and 25, inclusive."); 
}
```

### do-while
`do...while` 语句创建一个执行指定语句的循环，直到 `condition` 值为 `false`。在执行 `statement` 后检测 `condition`，所以指定的 `statement` 至少执行一次。

#### 语法
```js
do
   statement
while (condition);
```
- `statement` 执行至少一次的语句，并在每次 `condition` 值为真时重新执行。想执行多行语句，可使用 `block` 语句（{ ... }）包裹这些语句。

- `condition` 循环中每次都会计算的表达式。如果 `condition` 值为真， `statement` 会再次执行。当 `condition` 值为假，则跳到 `do...while` 之后的语句。

```js
let i = 0;

do { 
   i += 2; 
} while (i < 10);
```

### while
`while` 语句可以在某个条件表达式为真的前提下，循环执行指定的一段代码，直到那个表达式不为真时结束循环
```js
while (condition){
  statement
}
```

- `condition` 条件表达式，在每次循环前被求值。如果求值为真，`statement` 就会被执行。如果求值为假，则跳出 `while` 循环执行后面的语句。

- `statement` 只要条件表达式求值为真，该语句就会一直被执行。要在循环中执行多条语句，可以使用块语句（`{ ... }`）包住多条语句。注意：使用 `break` 语句在 `condition` 计算结果为真之前停止循环。

```js
let n = 0;
let x = 0;

while (n < 3) {
  n++;
  x += n;
}
```
在每次循环中，n 都会自增 1，然后再把 n 加到 x 上。因此，在每轮循环结束后，x 和 n 的值分别是：

- 第一轮后：n = 1，x = 1
- 第二轮后：n = 2，x = 3
- 第三轮后：n = 3，x = 6
当完成第三轮循环后，条件表达式 n< 3 不再为真，因此循环终止。

### for
`for` 语句用于创建一个循环，它包含了三个可选的表达式，这三个表达式被包围在圆括号之中，使用分号分隔，后跟一个用于在循环中执行的语句（通常是一个块语句）。

#### 语法
```js
for (initialization; condition; afterthought)
  statement
```
- `initialization` （可选）在循环开始前初始化的表达式（包含赋值表达式）或者变量声明。通常用于初始化计数器变量。该表达式可以选择使用 `var` 或 `let` 关键字声明新的变量，使用 `var` 声明的变量不是该循环的局部变量，而是与 `for` 循环处在同样的作用域中。用 `let` 声明的变量是语句的局部变量。

该表达式的结果会被丢弃。

- `condition` (可选) 每次循环迭代之前要判定的表达式。如果该表达式的判定结果为真，`statement` 将被执行。如果判定结果为假，那么执行流程将退出循环，并转到 `for` 结构后面的第一条语句。

这个条件测试是可选的。如果省略，该条件总是计算为真。

- `afterthought` (可选) 每次循环迭代结束时执行的表达式。执行时机是在下一次判定 `condition` 之前。通常被用于更新或者递增计数器变量。

- `statement` 只要条件的判定结果为真就会被执行的语句。你可以使用块语句来执行多个语句。如果没有任何语句要执行，请使用一个空语句（;）。

```js
let count = 10; 
for (let i = 0; i < count; i++) { 
   console.log(i); 
}
```

### for...in
`for...in` 语句迭代一个对象的所有可枚举字符串属性（除 `Symbol` 以外），包括继承的可枚举属性。

#### 语法
```js
for (variable in object)
    statement
```

- `variable` 在每次迭代时接收一个字符串属性名。它可以通过使用 `const`、`let` 或 `var` 进行声明，也可以是一个赋值目标（例如，先前声明的变量、对象属性或解构赋值模式）。使用 `var` 声明的变量不会局限于循环内部，即它们与 `for...in` 循环所在的作用域相同。
- `object` 被迭代非符号可枚举属性的对象。
- `statement` 每次迭代后执行的语句。可以引用 `variable`。可以使用块语句执行多个语句。

```js
const obj = { a: 1, b: 2, c: 3 };

for (const prop in obj) {
  console.log(`obj.${prop} = ${obj[prop]}`);
}

// 输出：
// "obj.a = 1"
// "obj.b = 2"
// "obj.c = 3"
```
### for...of
`for-of` 语句是一种严格的迭代语句，用于遍历可迭代对象的元素，语法如下：
```js
for (property of expression) statement 
```
下面是示例：
```js
for (const el of [2,4,6,8]) { 
   document.write(el); 
}
```

### 标签语句
标签语句用于给语句加标签，语法如下：
```js
label: statement 
```
下面是一个例子：
```js
start: for (let i = 0; i < count; i++) { 
 console.log(i); 
}
```

### break 和 continue
**`break` 和 `continue` 语句为执行循环代码提供了更严格的控制手段**。其中，**语句`break` 语句用于立即退出循环，强制执行循环后的下一条语句**。**而 `continue` 语句也用于立即退出循环，但会再次从循环顶部开始执行**。

```js
let num = 0;
for (let i = 1; i < 10; i++) { 
   if (i % 5 == 0) { 
       break;
   } 
   num++; 
} 
console.log(num); // 4
```

`break` 和 `continue` 都可以与标签语句一起使用，返回代码中特定的位置

```js
let num = 0; 
outermost: 
for (let i = 0; i < 10; i++) { 
   for (let j = 0; j < 10; j++) { 
       if (i == 5 && j == 5) { 
           break outermost; 
       } 
       num++; 
   } 
} 
console.log(num); // 55
```

### switch
`switch` 语句会对一个表达式求值，并将表达式的值与一系列 `case` 子句进行匹配，一旦遇到与表达式值相匹配的第一个 `case` 子句后，将执行该子句后面的语句，直到遇到 `break` 语句为止。若没有 `case` 子句与表达式的值匹配，如果没有任何 `case` 子句与表达式的值匹配，则会跳转至 `switch` 语句的 `default` 子句执行。

```js
const foo = 0;
switch (foo) {
  case -1:
    console.log("负 1");
    break;
  case 0: // foo 的值匹配这个条件；执行从这里开始
    console.log(0);
  // 忘记了 break！执行穿透
  case 1: // 'case 0:' 中没有 break 语句，所以这个 case 也会执行
    console.log(1);
    break; // 遇到 break，不会继续到 'case 2:'
  case 2:
    console.log(2);
    break;
  default:
    console.log("default");
}
// 输出 0 和 1
```