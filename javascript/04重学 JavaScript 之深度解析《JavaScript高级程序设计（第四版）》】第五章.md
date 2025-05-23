本文内容：
- 理解对象
- 基本 JavaScript 数据类型
- 引用值与原始值的包装类型

引用值（或者对象）是某个特定引用类型的实例。对象被认为是某个特定引用类型的实例。新对象通过使用 new 操作符后跟一个构造函数（`constructor`）来创建。构造函数就是用来创建新对象的函数，比如下面这行代码：

```js
const now = new Date();
```
这行代码创建了引用类型 `Date` 的一个新实例，并将它保存在变量 `now` 中。`Date()` 在这里就是构造函数，它负责创建一个只有默认属性和方法的简单对象。

## Date

创建一个 JavaScript Date 实例，该实例呈现时间中的某个时刻。`Date` 类型将日期保存为自协调世界时（[UTC，Universal Time Coordinated](https://zh.wikipedia.org/wiki/%E5%8D%8F%E8%B0%83%E4%B8%96%E7%95%8C%E6%97%B6)）时间 1970 年 1 月 1 日午夜（零时）至今所经过的毫秒数。使用这种存储格式，`Date` 类型可以精确表示 1970 年 1 月 1 日之前及之后 285616 年的日期。

要创建日期对象，就使用 new 操作符来调用 Date 构造函数：
```js
const now = new Date();
```

`Date()` 构造函数有四种基本形式：
> - `year` 表示**年份**的整数值。0 到 99 会被映射至 1900 年至 1999 年，其他值代表实际年份。
> - `monthIndex` 表示**月份**的整数值，从 0（1 月）到 11（12 月）。
> - `day` 表示一个月中的**第几天**的整数值，从 1 开始；默认值为 1。
> - `hours` 表示一天中的**小时数**的整数值 (24 小时制)；默认值为 0（午夜）。
> - `minutes` 表示一个完整时间中的**分钟**部分的整数值；默认值为 0。
> - `seconds` 表示一个完整时间中的**秒**部分的整数值；默认值为 0。
> - `milliseconds` 表示一个玩这个时间的**毫秒**部分的整数值；默认值为 0。 

- **不传递任何参数**：创建一个代表当前日期和时间的 Date 对象。
  ```js
     // new Date();
     let now = new Date();
  ```
- **传递一个表示日期的字符串**：创建一个对应于该字符串描述的日期和时间的 Date 对象。字符串的格式应该可以被 `Date.parse()` 方法接受。
  ```js
     // new Date(dateString);
     let someDate = new Date("July 7, 2024 07:07:07");
  ```
- **传递年、月和日**：创建一个对应于指定日期和时间的 Date 对象。年份是四位数字，月份是 0-11（0代表一月），日期是1-31。
  ```js
     // new Date(value);
     let someDate2 = new Date(2024, 6, 7);
  ```
- **传递年、月、日、小时、分钟、秒和毫秒**：创建一个对应于指定日期和时间的Date对象。年份是四位数字，月份是0-11（0代表一月），日期是1-31，小时是0-23，分钟和秒都是0-59，毫秒是0-999。
  ```js
    // new Date(year, monthIndex [, day [, hours [, minutes [, seconds [, milliseconds]]]]]);
     let someDate3 = new Date(2024, 6, 7, 7, 7, 7, 7);
  ```
  这些都是创建 Date 对象的方式，每种方式都根据所给参数的不同，会返回代表不同日期和时间的 Date 对象。

> Tips：
> - 在不给 Date 构造函数传参数的情况下，创建的对象将保存当前日期和时间
> - 要给予其他日期和时间创建日期对象时，必须传入其毫秒表示（UNIX 纪元 1970 年 1 月 1 日午夜之后的毫秒数），可以使用 `Date.parse()` 和 `Date.UTC()` 两个辅助方法

如果传给 `Date.parse()` 的字符串并不表示日期，则该方法会返回 `NaN`。如果直接把表示日期的字符串传给 `Date` 构造函数，那么 `Date` 会在后台调用 `Date.parse()`。换句话说，下面这两是等价的：
```js
const someDate = new Date("May 23, 2024"); // ==> const someDate = new Date(Date.parse("May 23, 2024"));
```
得到的日期对象相同。

`Date.UTC()` 方法也返回日期的毫秒表示，但使用的是跟 `Date.parse()` 不同的信息来生成这个值。传给 `Date.UTC()` 的参数是年、零起点月数（1 月是 0，2 月是 1，以此类推）、日（1~31）、时（0~23）、分、秒和毫秒。这些参数中，只有前两个（年和月）是必需的。如果不提供日，那么默认为 1 日。其他参数的默认值都是 0。下面是使用 `Date.UTC()` 的两个例子：
```js
// GMT 时间 2024 年 4 月 1 日零点
const y2k = new Date(Date.UTC(2024, 3));
console.log(y2k); // Mon Apr 01 2024 08:00:00 GMT+0800 (中国标准时间)

// GMT 时间 2024 年 4 月 5 日下午 5 点 55 分 55 秒
const allFives = new Date(Date.UTC(2024, 3, 5, 17, 55, 55));
console.log(allFives); // Sat Apr 06 2024 01:55:55 GMT+0800 (中国标准时间)
```

这个例子创建了两个日期 。第一个日期是 2024 年 4 月 1 日零点（GMT），2024 代表年，3 代表月（1 月）。因为没有其他参数（日取 1，其他取 0），所以结果就是该月第 1 天零点。第二个日期表示 2024 年 4 月 5 日下午 5 点 55 分 55 秒（GMT）。虽然日期里面涉及的都是 4，但月数必须用 3，因为月数是零起点的。小时也必须是 17，因为这里采用的是 24 小时制，即取值范围是 0~23。其他参数就都很直观了。

与 `Date.parse()` 一样，`Date.UTC()` 也会被 `Date` 构造函数隐式调用，但有一个区别：这种情况下创建的是本地日期，不是 GMT 日期。不过 `Date` 构造函数跟 `Date.UTC()` 接收的参数是一样的。因此，如果第一个参数是数值，则构造函数假设它是日期中的年，第二个参数就是月，以此类推。上面的例子也可以这样写：
```js
// GMT 时间 2024 年 4 月 1 日零点
const y2k = new Date(2024, 3);
console.log(y2k); // Mon Apr 01 2024 00:00:00 GMT+0800 (中国标准时间)

// GMT 时间 2024 年 4 月 5 日下午 5 点 55 分 55 秒
const allFives = new Date(2024, 3, 5, 17, 55, 55);
console.log(allFives); // Fri Apr 05 2024 17:55:55 GMT+0800 (中国标准时间)
```
以上代码创建了与前面例子中相同的两个日期，但这次的两个日期是（**由于系统设置决定的**）本地时区的日期。

ECMAScript 还提供了 `Date.now()` 方法，返回表示方法执行时日期和时间的毫秒数。这个方法可以方便地用在代码分析中：
```js
// 起始时间
let start = Date.now();

// 调用函数
doSomething(); 

// 结束时间
let stop = Date.now(), 
    result = stop - start;
```

### 继承的方法

与其他类型一样，Date 类型重写了 `toLocaleString()`、`toString()` 和 `valueOf()` 方法。但与其他类型不同，重写后这些方法的返回值不一样。Date 类型的 `toLocaleString()` 方法返回与浏览器运行的本地环境一致的日期和时间。这通常意味着格式中包含针对时间的 AM（上午）或 PM（下午），但不包含时区信息（具体格式可能因浏览器而不同）。`toString()` 方法通常返回带时区信息的日期和时间，而时间也是以 24 小时制（0~23）表示的。

下面给出了 `toLocaleString()` 和 `toString()` 返回的 2024 年 4 月 22 日零点的示例（地区为"en-US"的 PST，即 Pacific Standard Time，太平洋标准时间）：
```js
let date = new Date(Date.UTC(2024, 3, 22));  // 创建一个日期对象，表示2024年4月22日零点（UTC）

// 使用 toLocaleString() 方法，会返回一个根据当前地区设置的日期和时间字符串
let dateLocaleString = date.toLocaleString("en-US", { timeZone: "PST" });
console.log(dateLocaleString);  // 输出："4/21/2024, 5:00:00 PM"，因为PST比UTC早8小时

// 使用 toString() 方法，会返回一个包含日期、时间和时区的字符串
let dateString = date.toString();
console.log(dateString);  // 输出："Sun Apr 22 2024 00:00:00 GMT+0000 (UTC)"
```
在这个示例中，我们首先创建了一个 `Date` 对象，表示 2024 年 4 月 22 日零点（UTC）。然后我们使用 `toLocaleString()` 和 `toString()` 方法将这个Date对象转换为字符串。注意，`toLocaleString()` 方法的输出取决于当前的地区设置。在这个示例中，我们将地区设置为`"en-US"`，并设置时区为 `"PST"`。因此，输出的日期和时间是根据太平洋标准时间来表示的。而 `toString()` 方法返回的字符串总是使用UTC（协调世界时）来表示日期和时间。

**Date 类型的 `valueOf()` 方法根本就不返回字符串，这个方法被重写后返回的是日期的毫秒表示**。因此，操作符（如小于号和大于号）可以直接使用它返回的值。比如下面的例子：
```js
let date1 = new Date(2024, 3, 22);  // 创建一个日期对象，表示2024年4月22日
let date2 = new Date(2025, 3, 22);  // 创建另一个日期对象，表示2025年4月22日

console.log(date1.valueOf());  // 输出：1713715200000，这是date1的毫秒表示
console.log(date2.valueOf());  // 输出：1745251200000，这是date2的毫秒表示

console.log(date1 < date2);     // true
console.log(date1 > date2);     // false
console.log(date1 == date2);    // false
console.log(date1 === date2);   // false
```

### 日期格式化
Date 类型有几个专门用于格式化日期的方法，它们都会返回字符串：
- toDateString()显示日期中的周几、月、日、年（格式特定于实现）；
  ```js
  let date = new Date(2024, 3, 22, 13, 24, 0);  // 创建一个日期对象，表示2024年4月22日13:24:00
  console.log(date.toDateString());  // 输出："Mon Apr 22 2024"
  ```
- toTimeString()显示日期中的时、分、秒和时区（格式特定于实现）；
  ```js
  let date = new Date(2024, 3, 22, 13, 24, 0);  // 创建一个日期对象，表示2024年4月22日13:24:00
  console.log(date.toTimeString());  // 输出："13:24:00 GMT+0800 (China Standard Time)"
  ```
- toLocaleDateString()显示日期中的周几、月、日、年（格式特定于实现和地区）；
  ```js
  let date = new Date(2024, 3, 22, 13, 24, 0);  // 创建一个日期对象，表示2024年4月22日13:24:00
  console.log(date.toLocaleDateString());  // 输出："4/22/2024"
  ```
- toLocaleTimeString()显示日期中的时、分、秒（格式特定于实现和地区）；
  ```js
  let date = new Date(2024, 3, 22, 13, 24, 0);  // 创建一个日期对象，表示2024年4月22日13:24:00
  console.log(date.toLocaleTimeString());  // 输出："1:24:00 PM"
  ```
- toUTCString()显示完整的 UTC 日期（格式特定于实现）。
  ```js
  let date = new Date(2024, 3, 22, 13, 24, 0);  // 创建一个日期对象，表示2024年4月22日13:24:00
  console.log(date.toUTCString());  // 输出："Mon, 22 Apr 2024 05:24:00 GMT"
  ```
  这些方法的输出与 `toLocaleString()` 和 `toString()` 一样，会因浏览器而异。因此不能用于在用户界面上一致地显示日期。
> 还有一个方法叫 `toGMTString()`，这个方法跟 `toUTCString()` 是一样的，目的是为了向后兼容。不过，规范建议新代码使用 `toUTCString()`。**这些方法的输出可能会因浏览器和操作系统的不同而有所不同**。这是因为**这些方法的实现取决于浏览器和操作系统的国际化设置**。因此，如果你需要在用户界面上一致地显示日期，可能需要使用一些专门的日期格式化库，比如：[day.js](https://day.js.org/)。

### 日期/时间组件方法
Date 类型剩下的方法（见下表）直接涉及取得或设置日期值的特定部分。注意表中“UTC 日期”，指的是没有时区偏移（将日期转换为 GMT）时的日期。

| 方法                             | 说明                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| getTime()                        | 返回日期的毫秒表示；与 valueOf() 相同                        |
| setTime(milliseconds)            | 设置日期的毫秒表示，从而修改整个日期                         |
| getFullYear()                    | 返回 4 位数年（即 2019 而不是 19）                           |
| getUTCFullYear()                 | 返回 UTC 日期的 4 位数年                                     |
| setFullYear(year)                | 设置日期的年（year 必须是 4 位数）                           |
| setUTCFullYear(year)             | 设置 UTC 日期的年（year 必须是 4 位数）                      |
| getMonth()                       | 返回日期的月（0 表示 1 月，11 表示 12 月）                   |
| getUTCMonth()                    | 返回 UTC 日期的月（0 表示 1 月，11 表示 12 月）              |
| setMonth(month)                  | 设置日期的月（month 为大于 0 的数值，大于 11 加年）          |
| setUTCMonth(month)               | 设置 UTC 日期的月（month 为大于 0 的数值，大于 11 加年）     |
| getDate()                        | 返回日期中的日（1~31）                                       |
| getUTCDate()                     | 返回 UTC 日期中的日（1~31）                                  |
| setDate(date)                    | 设置日期中的日（如果 date 大于该月天数，则加月）             |
| setUTCDate(date)                 | 设置 UTC 日期中的日（如果 date 大于该月天数，则加月）        |
| getDay()                         | 返回日期中表示周几的数值（0 表示周日，6 表示周六）           |
| getUTCDay()                      | 返回 UTC 日期中表示周几的数值（0 表示周日，6 表示周六）      |
| getHours()                       | 返回日期中的时（0~23）                                       |
| getUTCHours()                    | 返回 UTC 日期中的时（0~23）                                  |
| setHours(hours)                  | 设置日期中的时（如果 hours 大于 23，则加日）                 |
| setUTCHours(hours)               | 设置 UTC 日期中的时（如果 hours 大于 23，则加日）            |
| getMinutes()                     | 返回日期中的分（0~59）                                       |
| getUTCMinutes()                  | 返回 UTC 日期中的分（0~59）                                  |
| setMinutes(minutes)              | 设置日期中的分（如果 minutes 大于 59，则加时）               |
| setUTCMinutes(minutes)           | 设置 UTC 日期中的分（如果 minutes 大于 59，则加时）          |
| getSeconds()                     | 返回日期中的秒（0~59）                                       |
| getUTCSeconds()                  | 返回 UTC 日期中的秒（0~59）                                  |
| setSeconds(seconds)              | 设置日期中的秒（如果 seconds 大于 59，则加分）               |
| setUTCSeconds(seconds)           | 设置 UTC 日期中的秒（如果 seconds 大于 59，则加分）          |
| getMilliseconds()                | 返回日期中的毫秒                                             |
| getUTCMilliseconds()             | 返回 UTC 日期中的毫秒                                        |
| setMilliseconds(milliseconds)    | 设置日期中的毫秒                                             |
| setUTCMilliseconds(milliseconds) | 设置 UTC 日期中的毫秒                                        |
| getTimezoneOffset()              | 返回以分钟计的 UTC 与本地时区的偏移量（如美国 EST 即“东部标准时间”返回 300，进入夏令时的地区可能有所差异） |

## RegExp
**正则表达式是用于匹配字符串中字符组合的模式。** 在 JavaScript 中，正则表达式也是对象。这些模式被用于 RegExp 的 `exec` 和 `test` 方法，以及 String 的 `match`、`matchAll`、`replace`、`search` 和 `split` 方法。
```js
let expression = /pattern/flags;
```

这个正则表达式的 pattern（模式）可以是任何简单或复杂的正则表达式，包括字符类、限定符、分组、向前查找和反向引用。每个正则表达式可以带零个或多个 flags（标记），用于控制正则表达式的行为。下面给出了表示匹配模式的标记。
| 标记符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| g      | 全局模式，表示查找字符串的全部内容，而不是找到第一个匹配的内容就结束 |
| i      | 不区分大小写，表示在查找匹配时忽略 pattern 和字符串的大小写  |
| m      | 多行模式，表示查找到一行文本末尾时会继续查找                 |
| y      | 粘附模式，表示只查找从 lastIndex 开始及之后的字符串          |
| u      | Unicode 模式，启用 Unicode 匹配                              |
| s      | dotAll 模式，表示元字符.匹配任何字符（包括\n 或\r）          |

使用不同模式和标记可以创建出各种正则表达式，比如：
```js
// 匹配字符串中的所有"at" 
let pattern1 = /at/g; 

// 匹配第一个"bat"或"cat"，忽略大小写
let pattern2 = /[bc]at/i; 

// 匹配所有以"at"结尾的三字符组合，忽略大小写
let pattern3 = /.at/gi;
```

与其他语言中的正则表达式类似，所有元字符在模式中也必须转义，包括 `( [ { \ ^ $ | ) ] } ? * + .`；元字符在正则表达式中都有一种或多种特殊功能，所以要匹配上面这些字符本身，就必须使用反斜杠来转义。下面是几个例子：
```js
let pattern1 = /[bc]at/i; 
// 匹配第一个"[bc]at"，忽略大小写
let pattern2 = /\[bc\]at/i; 

// 匹配所有以"at"结尾的三字符组合，忽略大小写
let pattern3 = /.at/gi; 

// 匹配所有".at"，忽略大小写
let pattern4 = /\.at/gi;
```
这里的 pattern1 匹配"bat"或"cat"，不区分大小写。要直接匹配"[bc]at"，左右中括号都必须像 pattern2 中那样使用反斜杠转义。在 pattern3 中，点号表示"at"前面的任意字符都可以匹配。如果想匹配".at"，那么要像 pattern4 中那样对点号进行转义。

**正则表达式也可以使用 RegExp 构造函数来创建，它接收两个参数：模式字符串和（可选的）标记字符串**。任何使用字面量定义的正则表达式也可以通过构造函数来创建，比如：
```js
// 匹配第一个"bat"或"cat"，忽略大小写
let pattern1 = /[bc]at/i; 

// 跟 pattern1 一样，只不过是用构造函数创建的
let pattern2 = new RegExp("[bc]at", "i");
```
这里的 pattern1 和 pattern2 是等效的正则表达式。注意，RegExp 构造函数的两个参数都是字符串。因为 RegExp 的模式参数是字符串，所以在某些情况下需要二次转义。所有元字符都必须二次转义，包括转义字符序列，如\n（\转义后的字符串是\\，在正则表达式字符串中则要写成\\\\）。下表展示了几个正则表达式的字面量形式，以及使用 RegExp 构造函数创建时对应的模式字符串：
| 字面量模式       | 对应的字符串          |
| ---------------- | --------------------- |
| /\[bc\]at/       | "\\[bc\\]at"          |
| /\.at/           | "\\.at"               |
| /name\/age/      | "name\\/age"          |
| /\d.\d{1,2}/     | "\\d.\\d{1,2}"        |
| /\w\\hello\\123/ | "\\w\\\\hello\\\\123" |

此外，使用 RegExp 也可以基于已有的正则表达式实例，并可选择性地修改它们的标记：
```js
const re1 = /cat/g; 
console.log(re1); // "/cat/g" 

const re2 = new RegExp(re1); 
console.log(re2); // "/cat/g" 

const re3 = new RegExp(re1, "i"); 
console.log(re3); // "/cat/i"
```

## 原始值包装类型
在JavaScript中，有一些特殊的类型叫做引用类型，包括 `Boolean`（布尔值）、`Number`（数字）和 `String`（字符串）。这些类型都有一些特殊的行为，与我们在其他情况下使用的原始值相对应。每当我们需要使用原始值的方法或属性时，JavaScript会在后台为我们创建一个对应的“包装对象”。这样，我们就可以使用这些方法来操作原始值了。
```js
let s1 = "some text"; 
let s2 = s1.substring(2);
```
在这里，s1 是一个包含字符串的变量，它是一个原始值。第二行紧接着在 s1 上调用了 substring()方法，并把结果保存在 s2 中。原始值本身不是对象，因此逻辑上不应该有方法。而实际上这个例子又确实按照预期运行了。这是因为后台进行了很多处理，从而实现了上述操作。具体来说，当访问 s1 时，是以读模式访问的，也就是要从内存中读取变量保存的值。在以读模式访问字符串值的任何时候，后台都会执行以下 3 步：
1. 创建一个 String 类型的实例；
2. 调用实例上的特定方法；
3. 销毁实例

可以把这 3 步想象成执行了如下 3 行 ECMAScript 代码：
```js
let s1 = new String('some text');
const s2 = s1.substring(2);
s1 = null;
```
这种行为可以让原始值拥有对象的行为。对布尔值和数值而言，以上 3 步也会在后台发生，只不过使用的是 `Boolean` 和 `Number` 包装类型而已。

**引用类型与原始值包装类型的主要区别在于对象的生命周期。在通过 new 实例化引用类型后，得到的实例会在离开作用域时被销毁，而自动创建的原始值包装对象则只存在于访问它的那行代码执行期间。这意味着不能在运行时给原始值添加属性和方法。** 比如下面的例子：
```js
let s1 = "some text"; 
s1.color = "red"; 
console.log(s1.color); // undefined
```
这里的第二行代码尝试给字符串 s1 添加了一个 color 属性。接着访问 color 属性时，它却不见了。原因就是第二行代码运行时会临时创建一个 `String` 对象，而当第三行代码执行时，这个对象已经被销毁了。

可以显式地使用 `Boolean`、`Number` 和 `String` 构造函数创建原始值包装对象。不过应该在确实必要时再这么做，否则容易让开发者疑惑，分不清它们到底是原始值还是引用值。在原始值包装类型的实例上调用 `typeof` 会返回 `"object"`，所有原始值包装对象都会转换为布尔值 `true`。

另外，`Object` 构造函数作为一个工厂方法，能够根据传入值的类型返回相应原始值包装类型的实例。比如：
```js
let obj = new Object("some text"); 
console.log(obj instanceof String); // true 
```
如果传给 `Object` 的是字符串，则会创建一个 `String` 的实例。如果是数值，则会创建 `Number` 的实例。布尔值则会得到 `Boolean` 的实例。
注意，使用 `new` 调用原始值包装类型的构造函数，与调用同名的转型函数并不一样。例如：
```js
let value = "25"; 
let number = Number(value); // 转型函数
console.log(typeof number); // "number" 
let obj = new Number(value); // 构造函数
console.log(typeof obj); // "object" 
```
在这个例子中，变量 `number` 中保存的是一个值为 25 的原始数值，而变量 obj 中保存的是一个
Number 的实例。

虽然不推荐显式创建原始值包装类型的实例，但它们对于操作原始值的功能是很重要的。每个原始值包装类型都有相应的一套方法来方便数据操作。

### Boolean
Boolean 是对应布尔值的引用类型。要创建一个 Boolean 对象，就使用 Boolean 构造函数并传入 `true` 或 `false`；示例如下：
```js
const booleanObject = new Boolean(true);
```

Boolean 的实例会重写 `valueOf()` 方法，返回一个原始值 `true` 或 `false`。`toString()` 方法被调用时也会被覆盖，返回字符串 `"true"` 或 `"false"`。不过，Boolean 对象在 ECMAScript 中用得很少。不仅如此，它们还容易引起误会，尤其是在布尔表达式中使用 Boolean 对象时，比如：
```js
let falseObject = new Boolean(false); 
let result = falseObject && true; 
console.log(result); // true 

let falseValue = false; 
result = falseValue && true; 
console.log(result); // false
```
首先，`let falseObject = new Boolean(false);` 创建了一个Boolean包装对象，其值为 `false`。然后，`let result = falseObject && true;` 这行代码*判断 falseObject（**一个对象**）和 `true`之间的逻辑与关系*。在JavaScript中，对象（包括包装对象）在布尔上下文中被视为真值，因此`falseObject && true` 的结果是true。

在第二部分，`let falseValue = false;` 创建了一个原始的布尔值 `false`，然后 `result = falseValue && true;` 这行代码**判断 falseValue（**一个原始的布尔值**）和 `true` 之间的逻辑与关系**。因为 `falseValue` 的值是 `fals`e，所以 `falseValue && true` 的结果是false。所以虽然 `falseObject` 和 `falseValue` 看起来有相同的值，但由于它们是不同类型（一个是对象，一个是原始值），在逻辑与操作中表现出了不同的结果。

原始值和引用值（如Boolean对象）在JavaScript中是有很大区别的。以下是一些主要的区别：
- 存储方式：这个概念是在底层实现的，我们在JavaScript代码中看不到直接的示例。但是要知道原始值存储在栈中，引用值存储在堆中，并由栈中的一个指针指向它。
  ```js
  let a = 1;
  let b = a;
  b = 2;
  console.log(a); // 输出 1, a的值没有改变
  
  let obj1 = new Boolean(false);
  let obj2 = obj1;
  obj2.valueOf = function(){ return true; };
  console.log(obj1.valueOf()); // 输出 true, obj1的值被改变了
  ```
- 比较方式：
  ```js
  let a = 1;
  let b = 1;
  console.log(a === b); // 输出 true
  
  let obj1 = new Boolean(false);
  let obj2 = new Boolean(false);
  console.log(obj1 === obj2); // 输出 false
  ```
- 变动性：
  ```js
  let a = 1;
  a.prop = true;
  console.log(a.prop); // 输出 undefined, 不能给原始值添加属性
  
  let obj1 = new Boolean(false);
  obj1.prop = true;
  console.log(obj1.prop); // 输出 true, 可以给引用值添加属性
  ```
- 方法和属性：
  ```js
  let a = "hello";
  console.log(a.toUpperCase()); // 输出 "HELLO", JavaScript自动创建了一个临时的String对象
  
  let obj1 = new Boolean(false);
  console.log(obj1.valueOf()); // 输出 false, 引用值可以直接调用方法
  ```
  
### Number
Number 是对应数值的引用类型。要创建一个 Number 对象，就使用 Number 构造函数并传入一个数值，如下例所示：
```js
const numberObject = new Number(10);
```

与 Boolean 类型一样，Number 类型重写了 `valueOf()`、`toLocaleString()` 和 `toString()` 方法。`valueOf()` 方法返回 Number 对象表示的原始数值，另外两个方法返回数值字符串。`toString()` 方法可选地接收一个表示基数的参数，并返回相应基数形式的数值字符串，如下所示：
```js
// 创建一个原始数值
let num = 10; 

// 使用toString()方法将数值转换为十进制字符串表示
console.log(num.toString()); // 输出 "10" 

// 使用toString()方法并传入2作为参数，将数值转换为二进制字符串表示
console.log(num.toString(2)); // 输出 "1010" 

// 使用toString()方法并传入8作为参数，将数值转换为八进制字符串表示
console.log(num.toString(8)); // 输出 "12" 

// 使用toString()方法并传入10作为参数，将数值转换为十进制字符串表示
console.log(num.toString(10)); // 输出 "10" 

// 使用toString()方法并传入16作为参数，将数值转换为十六进制字符串表示
console.log(num.toString(16)); // 输出 "a" 
```
除了继承的方法，Number 类型还提供了几个用于将数值格式化为字符串的方法。
- `toFixed()` 方法返回包含指定小数点位数的数值字符串，如：
  ```js
  let num = 10; 
  console.log(num.toFixed(2)); // "10.00"
  ```
  这里的 `toFixed()` 方法接收了参数 2，表示返回的数值字符串要包含两位小数。结果返回值为"10.00"，小数位填充了 0。如果数值本身的小数位超过了参数指定的位数，则四舍五入到最接近的小数位：
  ```js
  let num = 10.005; 
  console.log(num.toFixed(2)); // "10.01"
  ```
  `toFixed()` 自动舍入的特点可以用于处理货币。不过要注意的是，多个浮点数值的数学计算不一定得到精确的结果。比如，`0.1 + 0.2 = 0.30000000000000004`。
  
- 用于格式化数值的方法是 toExponential()，返回以科学记数法（也称为指数记数法）表示的数值字符串。如：
  ```js
  let num = 10; 
  console.log(num.toExponential(1)); // "1.0e+1"
  ```
  这段代码的输出为 `"1.0e+1"`。一般来说，这么小的数不用表示为科学记数法形式。如果想得到数值最适当的形式，那么可以使用 `toPrecision()`。`toPrecision()` 方法会根据情况返回最合理的输出结果，可能是固定长度，也可能是科学记数法形式。这个方法接收一个参数，表示结果中数字的总位数（不包含指数）。来看几个例子：
  ```js
  // 创建一个原始数值
  let num = 99; 
  
  // 使用toPrecision()方法并传入1作为参数，将数值转换为长度为1的字符串表示
  // 由于99的长度超过了1，因此结果以科学记数法表示
  console.log(num.toPrecision(1)); // 输出 "1e+2" 
  
  // 使用toPrecision()方法并传入2作为参数，将数值转换为长度为2的字符串表示
  console.log(num.toPrecision(2)); // 输出 "99" 
  
  // 使用toPrecision()方法并传入3作为参数，将数值转换为长度为3的字符串表示
  // 由于99的长度小于3，因此在末尾添加了一个零以满足长度要求
  console.log(num.toPrecision(3)); // 输出 "99.0"
  ```
  > toPrecision()方法可以表示带 1~21 个小数位的数值。某些浏览器可能支持更大的范围，但这是通常被支持的范围
  

与 Boolean 对象类似，Number 对象也为数值提供了重要能力。但是，考虑到两者存在同样的潜在问题，因此并不建议直接实例化 Number 对象。在处理原始数值和引用数值时，`typeof` 和 `instacnceof` 操作符会返回不同的结果，如下所示：
```js
// 创建一个数值对象
let numberObject = new Number(10); 

// 创建一个原始数值
let numberValue = 10; 

// 使用typeof运算符检查数值对象的类型，结果是"object"
console.log(typeof numberObject); // 输出 "object" 

// 使用typeof运算符检查原始数值的类型，结果是"number"
console.log(typeof numberValue); // 输出 "number" 

// 使用instanceof运算符检查数值对象是否是Number的实例，结果是true
console.log(numberObject instanceof Number); // 输出 true 

// 使用instanceof运算符检查原始数值是否是Number的实例，结果是false
// 因为原始数值不是对象，所以不能是Number的实例
console.log(numberValue instanceof Number); // 输出 false
```
原始数值在调用 `typeof` 时始终返回 `"number"`，而 Number 对象则返回 `"object"`。类似地，Number 对象是 Number 类型的实例，而原始数值不是。

#### isInteger()方法与安全整数
ES6 新增了 `Number.isInteger()` 方法，用于辨别一个数值是否保存为整数。有时候，小数位的 0 可能会让人误以为数值是一个浮点值：
```js
console.log(Number.isInteger(1)); // true 
console.log(Number.isInteger(1.00)); // true 
console.log(Number.isInteger(1.01)); // false 
```
IEEE 754 数值格式有一个特殊的数值范围，在这个范围内二进制值可以表示一个整数值。这个数值范围从 Number.MIN_SAFE_INTEGER（-2^53 + 1）到 `Number.MAX_SAFE_INTEGER（2^53 - 1`）。对超出这个范围的数值，即使尝试保存为整数，IEEE 754 编码格式也意味着二进制值可能会表示一个完全不同的数值。为了鉴别整数是否在这个范围内，可以使用 `Number.isSafeInteger()` 方法：
```js
// 检查 -2^53 是否是安全整数，结果是false，因为-2^53超出了安全整数的范围
console.log(Number.isSafeInteger(-1 * (2 ** 53))); // 输出 false 

// 检查 -2^53 + 1 是否是安全整数，结果是true，因为-2^53 + 1在安全整数的范围内
console.log(Number.isSafeInteger(-1 * (2 ** 53) + 1)); // 输出 true 

// 检查 2^53 是否是安全整数，结果是false，因为2^53超出了安全整数的范围
console.log(Number.isSafeInteger(2 ** 53)); // 输出 false 

// 检查 2^53 - 1 是否是安全整数，结果是true，因为2^53 - 1在安全整数的范围内
console.log(Number.isSafeInteger((2 ** 53) - 1)); // 输出 true
```

### String

String 对象用于表示和操作字符序列。String 是对应字符串的引用类型。要创建一个 String 对象，使用 String 构造函数并传入一个数值，如下：
```js
const stringObject = new String('hello string');
```

String 对象的方法可以在所有字符串原始值上调用。3个继承的方法 `valueOf()`、`toLocaleString()` 和 `toString()` 都返回对象的原始字符串值。每个 String 对象都有一个 `length` 属性，表示字符串中字符的数量。比如：
```js
const stringValue = 'this is text';
console.log(stringValue.length);  // 12
```
> 注意，即使字符串中包含双字节字符（而不是单字节的 ASCII 字符），也仍然会按单字符来计数。

#### JavaScript 字符
JavaScript 字符串由 [UTF-16 位码](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String#utf-16_%E5%AD%97%E7%AC%A6%E3%80%81unicode_%E7%A0%81%E4%BD%8D%E5%92%8C%E5%AD%97%E7%B4%A0%E7%B0%87)元组成；对多数字来说，每 16 位码元对应一个字符。换句话说，字符串的 `length` 属性表示字符串包含多少 16 位码元:
```js
const letter = 'abcdefghijklmnopqrstuvwxyz';
console.log('length:', letter.length); // 26
```

- `charAt()` 方法返回给定索引位置的字符，由传给方法的整数参数指定。**传入的参数从零开始，且会被转换为整数(`undefined` 会被转换为 0)；传入的参数如果大于字符串的长度，则返回一个空字符串**。
  ```js
  const message = "abcde";
  console.log(message.charAt(2)); // c
  
  const anyString = "Brave new world";
  
  // 没有提供索引，使用 0 作为默认值
  console.log(`在索引 0 处的字符为 '${anyString.charAt()}'`);      // 在索引 0 处的字符为 'B'
  console.log(`在索引 0 处的字符为 '${anyString.charAt(0)}'`);     // 在索引 0 处的字符为 'B'
  console.log(`在索引 1 处的字符为 '${anyString.charAt(1)}'`);     // 在索引 1 处的字符为 'r'
  console.log(`在索引 2 处的字符为 '${anyString.charAt(2)}'`);     // 在索引 2 处的字符为 'a'
  console.log(`在索引 3 处的字符为 '${anyString.charAt(3)}'`);     // 在索引 3 处的字符为 'v'
  console.log(`在索引 4 处的字符为 '${anyString.charAt(4)}'`);     // 在索引 4 处的字符为 'e'
  console.log(`在索引 999 处的字符为 '${anyString.charAt(999)}'`); // 在索引 999 处的字符为 ''
  ```
  `charAt()` 可能会返回孤项代理，这些代理项不是有效的 Unicode 字符。
  ```js
  const str = "𠮷𠮾";
  console.log(str.charAt(0)); // "\ud842"，这不是有效的 Unicode 字符
  console.log(str.charAt(1)); // "\udfb7"，这不是有效的 Unicode 字符
  ```
  要获取给定索引处的完整 Unicode 码位，请使用按 Unicode 码位拆分的索引方法，例如 `String.prototype.codePointAt()` 和将字符串展开为 Unicode 码位数组。
  ```js
  const str = "𠮷𠮾";
  console.log(String.fromCodePoint(str.codePointAt(0))); // "𠮷"
  console.log([...str][0]); // "𠮷"
  ```
  > JavaScript 字符串使用了两种 Unicode 编码混合的策略：[UCS-2](https://zh.wikipedia.org/wiki/%E9%80%9A%E7%94%A8%E5%AD%97%E7%AC%A6%E9%9B%86) 和 [UTF-16](https://zh.wikipedia.org/wiki/UTF-16)。对于可以采用 16 位编码的字符（U+0000~U+FFFF），这两种编码实际上是一样的。要深入了解关于字符编码的内容，推荐 Joel Spolsky 写的博客文章：[“The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)”](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)。
  > 另一个有用的资源是 Mathias Bynens 的博文：[“JavaScript’s Internal Character Encoding: UCS-2 or UTF-16?”](https://mathiasbynens.be/notes/javascript-encoding)。

- `charCodeAt()` 方法可以查看指定码元的字符编码，传入的参数值从零开始，其值将被转换为整数（undefined 被转换为 0），**如果参数值超出了 0 到 str.length - 1 的范围 `NAN`，传入的参数值介于 0 和 65535 之间。**
  ```js
  let message = "abcde"; 
  
  // Unicode "Latin small letter C"的编码是 U+0063 
  console.log(message.charCodeAt(2)); // 99 ==> c 对应的 unicode 编码是 99 
  
  console.log(99 === 0x63); // true ==>  十进制 99 等于十六进制 63
  ```
  
  `charCodeAt()` 可能会返回单独代理项，它们不是有效的 Unicode 字符。
  ```js
  const str = "𠮷𠮾";
  console.log(str.charCodeAt(0)); // 55362 或 d842，不是有效的 Unicode 字符
  console.log(str.charCodeAt(1)); // 57271 或 dfb7，不是有效的 Unicode 字符
  ```
  
  要获取给定索引处的完整 Unicode 码位，可以使用 `String.prototype.codePointAt()` 方法。
  ```js
  const str = "𠮷𠮾";
  console.log(str.codePointAt(0)); // 134071
  ```
  
- `String.fromCharCode()` 静态方法返回由指定的 UTF-16 码元序列创建的字符串。
  ```js
  console.log(String.fromCharCode(189, 43, 190, 61)); // "½+¾="
  String.fromCharCode(65, 66, 67); // 返回 "ABC"
  String.fromCharCode(0x2014);     // 返回 "—"
  String.fromCharCode(0x12014);    // 也返回 "—"；数字 1 被截断并忽略
  String.fromCharCode(8212);       // 也返回 "—"；8212 是 0x2014 的十进制表示
  String.fromCharCode(0xd83c, 0xdf03); // 码位 U+1F303 "Night with
  String.fromCharCode(55356, 57091);   // Stars" == "\uD83C\uDF03"
  String.fromCharCode(0xd834, 0xdf06, 0x61, 0xd834, 0xdf07); // "\uD834\uDF06a\uD834\uDF07"
  ```
  
  对于 U+0000~U+FFFF 范围内的字符，`length`、`charAt()`、`charCodeAt()` 和 `fromCharCode()` 返回的结果都跟预期是一样的。这是因为在这个范围内，**每个字符都是用 16 位表示的，而这几个方法也都基于 16 位码元完成操作**。只要字符编码大小与码元大小一一对应，这些方法就能如期工作。
  
  这个对应关系在扩展到 Unicode 增补字符平面时就不成立了。问题很简单，即 **16 位只能唯一表示 65 536 个字符**。这对于大多数语言字符集是足够了，在 Unicode 中称为基本多语言平面（BMP）。为了表示更多的字符，Unicode 采用了一个策略，即每个字符使用另外 16 位去选择一个增补平面。这种每个字符使用两个 16 位码元的策略称为代理对。
  
  在涉及增补平面的字符时，前面讨论的字符串方法就会出问题。比如，下面的例子中使用了一个笑脸表情符号，也就是一个使用代理对编码的字符：
  ```js
  let message = "ab☺de"; 
  
  console.log(message.length);         // 6
  console.log(message.charAt(1));      // b
  console.log(message.charAt(2));      // <?>
  console.log(message.charAt(3));      // <?> 
  console.log(message.charAt(4));      // d 
  console.log(message.charCodeAt(1));  // 98 
  console.log(message.charCodeAt(2));  // 55357 
  console.log(message.charCodeAt(3));  // 56842 
  console.log(message.charCodeAt(4));  // 100 
  console.log(String.fromCodePoint(0x1F60A)); // ☺
  console.log(String.fromCharCode(97, 98, 55357, 56842, 100, 101)); // ab☺de
  ```
  
  #### normalize() 方法
  String 的 `normalize()` 方法返回该字符串的 Unicode 标准化形式。有的字符既可以通过一个 BMP 字符表示，也可以通过一个代理对表示。比如：
  ```js
  // U+00C5：上面带圆圈的大写拉丁字母 A 
  console.log(String.fromCharCode(0x00C5)); // Å 
  // U+212B：长度单位“埃”
  console.log(String.fromCharCode(0x212B)); // Å 
  // U+004：大写拉丁字母 A 
  // U+030A：上面加个圆圈
  console.log(String.fromCharCode(0x0041, 0x030A)); // Å
  ```
  比较操作符不在乎字符看起来是什么样的，因此这 3 个字符互不相等。
  ```js
  let a1 = String.fromCharCode(0x00C5), 
      a2 = String.fromCharCode(0x212B), 
      a3 = String.fromCharCode(0x0041, 0x030A); 
      
  console.log(a1, a2, a3); // Å, Å, Å 
  console.log(a1 === a2);  // false 
  console.log(a1 === a3);  // false 
  console.log(a2 === a3);  // false
  ```
  为解决这个问题，Unicode提供了 4种规范化形式，可以将类似上面的字符规范化为一致的格式，无论底层字符的代码是什么。这 4种规范化形式是：
  - NFD（Normalization Form D）：规范分解
  - NFC（Normalization Form C）：规范分解，然后进行规范组合。
  - NFKD（Normalization Form KD）：兼容分解
  - NFKC（Normalization Form KC）：兼容分解，然后进行规范组合
  
  可以使用 `normalize()` 方法对字符串应用上述规范化形式，使用时需要传入表示哪种形式的字符串："NFD"、"NFC"、"NFKD"或"NFKC"。
  ```js
  let a1 = String.fromCharCode(0x00C5),
      a2 = String.fromCharCode(0x212B), 
      a3 = String.fromCharCode(0x0041, 0x030A); 
      
  // U+00C5 是对 0+212B 进行 NFC/NFKC 规范化之后的结果
  console.log(a1 === a1.normalize("NFD"));  // false 
  console.log(a1 === a1.normalize("NFC"));  // true 
  console.log(a1 === a1.normalize("NFKD")); // false 
  console.log(a1 === a1.normalize("NFKC")); // true 
  
  // U+212B 是未规范化的
  console.log(a2 === a2.normalize("NFD"));  // false 
  console.log(a2 === a2.normalize("NFC"));  // false 
  console.log(a2 === a2.normalize("NFKD")); // false 
  console.log(a2 === a2.normalize("NFKC")); // false 
  
  // U+0041/U+030A 是对 0+212B 进行 NFD/NFKD 规范化之后的结果
  console.log(a3 === a3.normalize("NFD"));  // true 
  console.log(a3 === a3.normalize("NFC"));  // false 
  console.log(a3 === a3.normalize("NFKD")); // true 
  console.log(a3 === a3.normalize("NFKC")); // false
  ```
  选择同一种规范化形式可以让比较操作符返回正确的结果：
  ```js
  let a1 = String.fromCharCode(0x00C5), 
      a2 = String.fromCharCode(0x212B), 
      a3 = String.fromCharCode(0x0041, 0x030A); 
  console.log(a1.normalize("NFD") === a2.normalize("NFD"));   // true 
  console.log(a2.normalize("NFKC") === a3.normalize("NFKC")); // true 
  console.log(a1.normalize("NFC") === a3.normalize("NFC"));   // true
  ```
  
  #### 字符串操作方法
  - `concat()` 方法将字符串参数连接到调用的字符串，并返回一个新的字符串。
    ```js
    const hello = "Hello, ";
    console.log(hello.concat("Kevin", ". Have a nice day.")); // Hello, Kevin. Have a nice day.
    
    const greetList = ["Hello", " ", "Venkat", "!"];
    console.log("".concat(...greetList)); // "Hello Venkat!"
    console.log("".concat({}));   // "[object Object]"
    console.log("".concat([]));   // ""
    console.log("".concat(null)); // "null"
    console.log("".concat(true)); // "true"
    console.log("".concat(4, 5)); // "45"
    ```
    > `concat()` 方法与加号/字符串连接运算符（+，+=）非常相似，不同之处在于 `concat()` 直接将其参数强制转换为字符串进行连接，而加号运算符首先将其操作数强制转换为原始值，然后再进行连接。
  - 从字符串中提取子字符串的方法：`slice()`、`substr()`和 `substring()`。
  
  3个方法都返回调用它们的字符串的一个子字符串，而且都接收**一或两个参数**。**第一个参数表示子字符串开始的位置，第二个参数表示子字符串结束的位置**。对 `slice()` 和 `substring()` 而言，**第二个参数是提取结束的位置**（即该位置之前的字符会被提取出来）。对 `substr()` 而言，**第二个参数表示返回的子字符串数量**。任何情况下，**省略第二个参数都意味着提取到字符串末尾**。与 `concat()` 方法一样，`slice()`、`substr()` 和 `substring()` 也**不会修改调用原字符串**，而只会返回提取到的原始新字符串值。
  ```js
  let stringValue = "hello world"; 
  console.log(stringValue.slice(3));       // "lo world" 
  console.log(stringValue.substring(3));   // "lo world" 
  console.log(stringValue.substr(3));      // "lo world" 
  console.log(stringValue.slice(3, 7));    // "lo w" 
  console.log(stringValue.substring(3,7)); // "lo w" 
  console.log(stringValue.substr(3, 7));   // "lo worl"
  ```
  
  - `slice()` 方法**将所有负值参数都当成字符串长度加上负参数值**。
  - `substr()` 方法将**第一个负参数值当成字符串长度加上该值，将第二个负参数值转换为 0**。
  - `substring()` 方法会**将所有负参数值都转换为 0**。
    ```js
    let stringValue = "hello world"; 
    console.log(stringValue.slice(-3));        // "rld" 
    console.log(stringValue.substring(-3));    // "hello world" 
    console.log(stringValue.substr(-3));       // "rld" 
    console.log(stringValue.slice(3, -4));     // "lo w" 
    console.log(stringValue.substring(3, -4)); // "hel" 
    console.log(stringValue.substr(3, -4));    // ""
    ```
    
#### 字符串位置方法
查找指定字符串的位置有两个方法，分别是 `indexOf()` 和 `lastIndexOf()`；这两个方法从字符串中搜索传入的字符串，并返回位置，**如果没有找到则返回 -1**。

1. `indexOf()` 方法在字符串中搜索指定子字符串，并**返回其第一次出现的位置索引**。它可以**接受一个可选的参数指定搜索的起始位置**，*如果找到了指定的子字符串，则返回的位置索引大于或等于指定的数字，如果没有找到，则返回 -1*。
```js
const paragraph = "I think Ruth's dog is cuter than your dog!";

const searchTerm = 'dog';
const indexOfFirst = paragraph.indexOf(searchTerm);

console.log(`The index of the first "${searchTerm}" is ${indexOfFirst}`); // The index of the first "dog" is 15
console.log(`The index of the second "${searchTerm}" is ${paragraph.indexOf(searchTerm,indexOfFirst + 1)}`); // The index of the second "dog" is 38
```
- **搜索空字符串会产生奇怪的结果。如果没有第二个参数，或者第二个参数的值小于调用字符串的长度，则返回值与第二个参数的值相同**：
  ```js
  const str = "hello world";
  console.log(str.indexOf(""));     // 0
  console.log(str.indexOf("", 0));  // 0
  console.log(str.indexOf("", 3));  // 3
  console.log(str.indexOf("", 8));  // 8
  ```
- **如果第二个参数的值大于或等于字符串的长度，则返回值是字符串的长度**
  ```js
  const str = "hello world";
  console.log(str.indexOf("", 11));  // 11
  console.log(str.indexOf("", 13));  // 11
  console.log(str.indexOf("", 20));  // 11
  ```
- `indexOf()` 是**区分大小写**的：
  ```js
  console.log("Blue Whale".indexOf("blue")); // -1
  ```

2. `lastIndexOf()` 方法**搜索该字符串并返回指定子字符串最后一次出现的索引**。它可以**接受一个可选的起始位置参数，并返回指定子字符串在小于或等于指定数字的索引中的最后一次出现的位置，否则返回 -1**。
```js
const paragraph = "I think Ruth's dog is cuter than your dog!";
const searchTerm = 'dog';

console.log(`Index of the last ${searchTerm} is ${paragraph.lastIndexOf(searchTerm)}`); // Index of the last "dog" is 38
```

- 字符串的索引从 0 开始：**字符串第一个字符的索引为 0，字符串最后一个字符的索引为字符串长度减 1**。
  ```js
  const str = "canal";
  console.log(str.lastIndexOf("a"));     // 3
  console.log(str.lastIndexOf("a", 2));  // 1
  console.log(str.lastIndexOf("a", 0));  // -1
  console.log(str.lastIndexOf("x"));     // -1
  console.log(str.lastIndexOf("c", -5)); // 0
  console.log(str.lastIndexOf("c", 0));  // 0
  console.log(str.lastIndexOf(""));      // 5
  console.log(str.lastIndexOf("", 2));   // 2
  ```
- **`lastIndexOf()` 区分大小写**
  ```js
  console.log("change".lastIndexOf("Ge")); // -1
  ```

> **示例巩固**：
>
> 使用 `indexOf()` 和 `lastIndexOf()`，以下示例使用 `indexOf()` 和 `lastIndexOf()` 在字符串 "Brave, Brave New World" 中查找值。
> ```js
> const str = "Brave, Brave New World";
> 
> console.log(str.indexOf("Brave"));     // 0
> console.log(str.lastIndexOf("Brave")); // 7
> ```

#### 字符串包含方法
ECMAScript 6 增加了 3 个用于判断字符串中是否包含另一个字符串的方法：`startsWith()`、`endsWith()` 和 `includes()`。这些方法都会从字符串中搜索传入的字符串，并**返回一个表示是否包含的布尔值**。

的区别在于:
- `startsWith()` 检查开始于索引 0 的匹配项，
- `endsWith()` 检查开始于索引(`string.length - substring.length`)的匹配项
- `includes()` 检查整个字符串
```js
let message = "foobarbaz"; 
console.log(message.startsWith("foo")); // true 
console.log(message.startsWith("bar")); // false 
console.log(message.endsWith("baz"));   // true 
console.log(message.endsWith("bar"));   // false 
console.log(message.includes("bar"));   // true 
console.log(message.includes("qux"));   // false
```

`startsWith()` 和 `includes()` 方法**接收可选的第二个参数，表示开始搜索的位置**。如果传入**第二个参数，则意味着从指定位置向字符串末尾进行搜索，忽略该位置之前的所有字符**。
```js
let message = "foobarbaz";

console.log(message.startsWith("foo"));    // true 
console.log(message.startsWith("foo", 1)); // false 
console.log(message.includes("bar"));      // true 
console.log(message.includes("bar", 4));   // false
```

`endsWith()` 方法接收可选的第二个参数，表示应该**当作字符串末尾的位置**。如果不提供这个参数，那么**默认就是字符串长度**。
```js
let message = "foobarbaz";

console.log(message.endsWith("bar"));    // false 
console.log(message.endsWith("bar", 6)); // true
```

#### trim() 方法
`trim()` 方法会从字符串的两端移除空白字符，并返回一个新的字符串，而不会修改原始字符串。
```js
let stringValue = " hello world ";
let trimmedStringValue = stringValue.trim();

console.log(stringValue);        // " hello world "
console.log(trimmedStringValue); // "hello world"
```
由于 `trim()` 返回的是字符串的副本，因此**原始字符串不受影响，即原本的前、后空格符都会保留**。另外，`trimStart()` 和 `trimEnd()` 方法分别用于从字符串开始和末尾清理空格符。

#### repeat() 方法
`repeat()` 方法**接收一个整数参数，表示要将字符串复制多少次**，然后返回拼接所有副本后的结果。**如果传入的这个参数为负值，或者超过了字符串的最大长度，将抛出错误；如果是一个小数，则向下取整。**
```js
const str = "abc";
console.log(str.repeat(-1));  // 报错：RangeError: Invalid count value: -1
console.log(str.repeat(0));   // '' ---> 空字符串
console.log(str.repeat(2));   // 'abcabc'
console.log(str.repeat(3.2)); // 'abcabcabc'
console.log(str.repeat(3.7)); // 'abcabcabc'
console.log(str.repeat(1/0)); // 报错：RangeError: Invalid count value: Infinity
```

#### padStart()和 padEnd()方法
`padStart()` 和 `padEnd()` 方法会复制字符串，如果小于指定长度，则**在相应一边填充字符，直至满足长度条件**。这两个方法的**第一个参数是长度，第二个参数是可选的填充字符串，默认为空格**

```js
let stringValue = "foo"; 

console.log(stringValue.padStart(6));      // " foo" 
console.log(stringValue.padStart(9, ".")); // "......foo" 
console.log(stringValue.padEnd(6));        // "foo " 
console.log(stringValue.padEnd(9, "."));   // "foo......"
```

可选的第二个参数并不限于一个字符。**如果提供了多个字符的字符串，则会将其拼接并截断以匹配
指定长度**。此外，**如果长度小于或等于字符串长度，则会返回原始字符串**。

```js
let stringValue = "foo";

console.log(stringValue.padStart(8, "bar")); // "barbafoo" 
console.log(stringValue.padStart(2));        // "foo" 
console.log(stringValue.padEnd(8, "bar"));   // "foobarba" 
console.log(stringValue.padEnd(2));          // "foo"
```
##### 字符串迭代与解构
字符串的原型上暴露了一个 `@@iterator` 方法，表示可以迭代字符串的每个字符。
```js
let message = "abc";
let stringIterator = message[Symbol.iterator](); 
console.log(stringIterator.next()); // {value: "a", done: false} 
console.log(stringIterator.next()); // {value: "b", done: false} 
console.log(stringIterator.next()); // {value: "c", done: false} 
console.log(stringIterator.next()); // {value: undefined, done: true}
```
在 `for...of` 循环中可以通过这个迭代器按序访问每个字符
```js
for (const c of "abcde") { 
   console.log(c); 
} 
// a 
// b 
// c 
// d 
// e
```
有了这个迭代器之后，字符串就可以通过解构操作符来解构了。比如，可以更方便地把字符串分割为字符数组
```js
let message = "abcde"; 
console.log([...message]); // ["a", "b", "c", "d", "e"]
```

#### 字符串大小写转换
下一组方法涉及大小写转换，包括 4 个方法：`toLowerCase()`、`toLocaleLowerCase()`、`toUpperCase()` 和 `toLocaleUpperCase()`；`toLocaleLowerCase()` 和 `toLocaleUpperCase()` 方法旨在基于特定地区实现。
```js
let stringValue = "hello world"; 
console.log(stringValue.toLocaleUpperCase()); // "HELLO WORLD" 
console.log(stringValue.toUpperCase());       // "HELLO WORLD" 
console.log(stringValue.toLocaleLowerCase()); // "hello world" 
console.log(stringValue.toLowerCase());       // "hello world"
```
`toLowerCase()` 和 `toLocaleLowerCase()` 都返回 hello world，而 `toUpperCase()` 和 `toLocaleUpperCase()` 都返回 HELLO WORLD。

> 通常，如果不知道代码涉及什么语言，则最好使用**地区特定的转换方法**。

#### 字符串模式匹配方法
在字符串中实现模式匹配设计了几个方法。

1. 第一个就是 `match()` 方法，这个方法本质上跟 RegExp 对象的 `exec()` 方法相同。`match()` 方法**接收一个参数，可以是一个正则表达式字符串，也可以是一个 RegExp 对象**。
    ```js
    let text = "cat, bat, sat, fat"; 
    let pattern = /.at/; 
    
    // 等价于 pattern.exec(text) 
    let matches = text.match(pattern); 
    console.log(matches.index);         // 0 
    console.log(matches[0]);            // "cat" 
    console.log(pattern.lastIndex);     // 0
    ```
    
    `match()` 方法返回的数组与 RegExp 对象的 `exec()`方法返回的数组是一样的：
    - 第一个元素是与整个模式匹配的字符串，
    - 其余元素则是与表达式中的捕获组匹配的字符串（如果有的话）。
    
2. `search()`方法唯一的参数与 `match()` 方法一样：**正则表达式字符串或 RegExp 对象**。这个方法返回模式第一个匹配的位置索引，**如果没找到则返回 -1**。search()始终从字符串开头向后匹配模式。
    ```js
    let text = "cat, bat, sat, fat"; 
    let pos = text.search(/at/); 
    console.log(pos); // 1
    ```
    这里，`search(/at/)` 返回 1，即`"at"`的第一个字符在字符串中的位置。
    
3. `replace()` 方法。这个方法**接收两个参数，第一个参数可以是一个 RegExp 对象或一个字符串（这个字符串不会转换为正则表达式），第二个参数可以是一个字符串或一个函数**。
    > **如果第一个参数是字符串，那么只会替换第一个子字符串**。要想**替换所有子字符串，第一个参数必须为正则表达式并且带全局标记**。
    > ```js
    > let text = "cat, bat, sat, fat"; 
    > let result = text.replace("at", "ond"); 
    > console.log(result); // "cond, bat, sat, fat" 
    > 
    > result = text.replace(/at/g, "ond"); 
    > console.log(result); // "cond, bond, sond, fond"
    > ```
    
    第二个参数是字符串的情况下，有几个特殊的字符序列，可以用来插入正则表达式操作的值。
    | 字符序列                                                    | 替换文本                                                     |
    | ----------------------------------------------------------- | ------------------------------------------------------------ |
    | $$                                                          | $                                                            |
    | $&                                                          | 匹配整个模式的子字符串。与 `RegExp.lastMatch` 相同           |
    | $'                                                          | 匹配的子字符串之前的字符串。与 `RegExp.rightContext` 相同    |
    | $`|匹配的子字符串之后的字符串。与 `RegExp.leftContext` 相同 |                                                              |
    | $n                                                          | 匹配第 n 个捕获组的字符串，其中 n 是 0~9。比如，$1 是匹配第一个捕获组的字符串，$2 是匹配第二个捕获组的字符串，以此类推。如果没有捕获组，则值为空字符串 |
    | $nn                                                         | 匹配第 nn 个捕获组字符串，其中 nn 是 01~99。比如，$01 是匹配第一个捕获组的字符串，$02 是匹配第二个捕获组的字符串，以此类推。如果没有捕获组，则值为空字符串 |
    
#### localeCompare()方法
`localeCompare()`方法比较两个字符串，返回如下 3 个值中的一个。
- 如果**按照字母表顺序，字符串应该排在字符串参数前头，则返回负值**。（通常是-1，具体还要看与实际值相关的实现。）
- **如果字符串与字符串参数相等，则返回 0**。
- **如果按照字母表顺序，字符串应该排在字符串参数后头，则返回正值**。（通常是 1，具体还要看与实际值相关的实现。）

```js
let stringValue = "yellow";

console.log(stringValue.localeCompare("brick"));  // 1 
console.log(stringValue.localeCompare("yellow")); // 0 
console.log(stringValue.localeCompare("zoo"));    // -1
```
在这里，字符串 `"yellow"` 与 3 个不同的值进行了比较：`"brick"`、`"yellow"` 和 `"zoo"`。`"brick"` 按字母表顺序应该排在 `"yellow"` 前头，因此 `localeCompare()` 返回 1。`"yellow"` 等于 `"yellow"`，因此 `"localeCompare()"` 返回 0。最后，`"zoo"` 在 `"yellow"` 后面，因此 `localeCompare()` 返回-1。

#### HTML 方法
早期的浏览器开发商认为使用 JavaScript 动态生成 HTML 标签是一个需求。因此，早期浏览器扩展了规范，增加了辅助生成 HTML 标签的方法。下表总结了这些 HTML 方法。不过，这些方法基本上已经没有人使用了，因为结果通常不是语义化的标记。

| 方法             | 输出                                |
| ---------------- | ----------------------------------- |
| anchor(name)     | \<a name="name">string\</a>         |
| big()            | \<big>string\</big>                 |
| bold()           | \<b>string\</b>                     |
| fixed()          | \<tt>string\</tt>                   |
| fontcolor(color) | \<font color="color">string\</font> |
| fontsize(size)   | \<font size="size">string\</font>   |
| italics()        | \<i>string\</i>                     |
| link(url)        | \<a href="url">string\</a>          |
| small()          | \<small>string\</small>             |
| strike()         | \<strike>string\</strike>           |
| sub()            | \<sub>string\</sub>                 |
| sup()            | \<sup>string\</sup>                 |

## 单例内置对象
ECMA-262 对内置对象的定义是“任何由 ECMAScript 实现提供、与宿主环境无关，并在 ECMAScript 程序开始执行时就存在的对象”。这就意味着，开发者不用显式地实例化内置对象，因为它们已经实例化好了。

### Global
Global 对象是 ECMAScript 中最特别的对象，因为代码不会显式地访问它。ECMA-262 规定 Global 对象为一种兜底对象，它所针对的是不属于任何对象的属性和方法。事实上，不存在全局变量或全局函数这种东西。在全局作用域中定义的变量和函数都会变成 Global 对象的属性。包括 `isNaN()`、`isFinite()`、`parseInt()` 和 `parseFloat()`，实际上都是 Global 对象的方法

#### URL 编码方法
`encodeURI()` 和 `encodeURIComponent()` 方法用于编码统一资源标识符（URI），以便传给浏览器。有效的 URI 不能包含某些字符，比如空格。使用 URI 编码方法来编码 URI 可以让浏览器能够理解它们，同时又以特殊的 UTF-8 编码替换掉所有无效字符。

`ecnodeURI()` 方法用于对整个 URI 进行编码，比如 "www.wrox.com/illegal value.js"。而 `encodeURIComponent()` 方法用于编码 URI 中单独的组件，比如前面 URL 中的"illegal value.js"。

两个方法的区别：
- `encodeURI()` 不会编码属于 URL 组件的特殊字符，比如`:`（冒号）、`\`（斜杠）、`?`（问号）、`#`（井号）
- `encodeURIComponent()` 会编码它发现的所有非标准字符

```js
let uri = "http://www.wrox.com/illegal value.js#start"; 

// "http://www.wrox.com/illegal%20value.js#start" 
console.log(encodeURI(uri)); 

// "http%3A%2F%2Fwww.wrox.com%2Fillegal%20value.js%23start" 
console.log(encodeURIComponent(uri));
```
这里使用 `encodeURI()` 编码后，除空格被替换为 `%20` 之外，没有任何变化。而 `encodeURIComponent()` 方法将所有非字母字符都**替换成了相应的编码形式**。

与 `encodeURI()` 和 `encodeURIComponent()` 相对的是 `decodeURI()` 和 `decodeURIComponent()`。`decodeURI()` 只对使用 `encodeURI()`编码过的字符解码。`decodeURIComponent()` 解码所有被 `encodeURIComponent()` 编码的字符，基本上就是解码所有特殊值。
```js
let uri = "http%3A%2F%2Fwww.wrox.com%2Fillegal%20value.js%23start"; 

console.log(decodeURI(uri));         // http%3A%2F%2Fwww.wrox.com%2Fillegal value.js%23start
console.log(decodeURIComponent(uri));// http:// www.wrox.com/illegal value.js#start 
```

#### eval() 方法
eval() 函数会将传入的字符串当做 JavaScript 代码进行执行。如果 `eval()` 方法返回值为空，则返回 `undefined`。
```js
console.log(eval('2 + 2')); // 4
console.log(eval(new String('2 + 2'))); // 2 + 2
console.log(eval('2 + 2') === eval('4')); // true
console.log(eval('2 + 2') === eval(new String('2 + 2'))); // false
```
> `eval()` 是一个很**危险的函数，尽可能不要使用它**，它使用与调用者相同的权限执行代码。如果你用 eval() 运行的字符串代码被恶意方（不怀好意的人）修改，你最终可能会在你的网页/扩展程序的权限下，在用户计算机上运行恶意代码。更重要的是，第三方代码可以看到某一个 eval() 被调用时的作用域，这也有可能导致一些不同方式的攻击。

####  Global 对象属性
Global 对象有很多属性，其中一些前面已经提到过了。像 `undefined`、`NaN` 和 `Infinity` 等特殊值都是 Global 对象的属性。此外，所有原生引用类型构造函数，比如 `Object` 和 `Function`，也都是 Global 对象的属性。

| 属性           | 说明                      |
| -------------- | ------------------------- |
| undefined      | 特殊值 undefined          |
| NaN            | 特殊值 NaN                |
| Infinity       | 特殊值 Infinity           |
| Object         | Object 的构造函数         |
| Array          | Array 的构造函数          |
| Function       | Function 的构造函数       |
| Boolean        | Boolean 的构造函数        |
| String         | String 的构造函数         |
| Number         | Number 的构造函数         |
| Date           | Date 的构造函数           |
| RegExp         | RegExp 的构造函数         |
| Symbol         | Symbol 的伪构造函数       |
| Error          | Error 的构造函数          |
| EvalError      | EvalError 的构造函数      |
| RangeError     | RangeError 的构造函数     |
| ReferenceError | ReferenceError 的构造函数 |
| SyntaxError    | SyntaxError 的构造函数    |
| TypeError      | TypeError 的构造函数      |
| URIError       | URIError 的构造函数       |

#### window 对象
虽然 ECMA-262 没有规定直接访问 Global 对象的方式，但浏览器将 `window` 对象实现为 Global对象的代理。因此，所有全局作用域中声明的变量和函数都变成了 window 的属性。
```js
var color = "red"; 
function sayColor() { 
   console.log(window.color); 
} 
window.sayColor(); // "red"
```
这里定义了一个名为 `color` 的全局变量和一个名为 `sayColor()` 的全局函数。在 `sayColor()` 内部，通过 `window.color` 访问了 `color` 变量，说明全局变量变成了 `window` 的属性。接着，又通过 `window` 对象直接调用了 `window.sayColor()` 函数，从而输出字符串。

另一种获取 Global 对象的方式是使用如下的代码：
```js
let global = function() { 
   return this; 
}();
```
这段代码创建一个立即调用的函数表达式，返回了 `this` 的值。如前所述，当一个函数在没有明确（通过成为某个对象的方法，或者通过 `call()/apply()`）指定 this 值的情况下执行时，`this` 值等于 Global 对象。

### Math
Math 是一个内置对象，它拥有一些数学常数属性和数学函数方法。Math 不是一个函数对象。

#### Math 对象属性
Math 对象有一些属性，主要用于保存数学中的一些特殊值。
| 属性         | 说明                  |
| ------------ | --------------------- |
| Math.E       | 自然对数的基数 e 的值 |
| Math.LN10    | 10 为底的自然对数     |
| Math.LN2     | 2 为底的自然对数      |
| Math.LOG2E   | 以 2 为底 e 的对数    |
| Math.LOG10E  | 以 10 为底 e 的对数   |
| Math.PI      | π 的值                |
| Math.SQRT1_2 | 1/2 的平方根          |
| Math.SQRT2   | 2 的平方根            |

####  min()和 max()方法
min()和 max()方法用于确定一组数值中的最小值和最大值。这两个方法都接收任意多个参数；如下：
```js
let max = Math.max(3, 54, 32, 16);
console.log(max); // 54 

let min = Math.min(3, 54, 32, 16); 
console.log(min); // 3
```
要知道数组中的最大值和最小值，可以像下面这样使用扩展操作符：
```js
let values = [1, 2, 3, 4, 5, 6, 7, 8]; 
let max = Math.max(...values);
console.log(max); // 8
```

#### 舍入方法
用于把小数值舍入为整数的 4 个方法：`Math.ceil()`、`Math.floor()`、`Math.round()` 和 `Math.fround()`。
- `Math.ceil()` 方法始终向上舍入为最接近的整数。
  ```js
  console.log(Math.ceil(25.9)); // 26 
  console.log(Math.ceil(25.5)); // 26 
  console.log(Math.ceil(25.1)); // 26
  ```
- `Math.floor()` 方法始终向下舍入为最接近的整数。
  ```js
  console.log(Math.round(25.9)); // 26 
  console.log(Math.round(25.5)); // 26 
  console.log(Math.round(25.1)); // 25
  ```
- `Math.round()` 方法执行四舍五入。
  ```js
  console.log(Math.fround(0.4));  // 0.4000000059604645 
  console.log(Math.fround(0.5));  // 0.5 
  console.log(Math.fround(25.9)); // 25.899999618530273
  ```
- `Math.fround()` 方法返回数值最接近的单精度（32 位）浮点值表示。
  ```js
  console.log(Math.floor(25.9)); // 25 
  console.log(Math.floor(25.5)); // 25 
  console.log(Math.floor(25.1)); // 25
  ```

#### random()方法
`Math.random()` 静态方法返回一个大于等于 0 且小于 1 的伪随机浮点数，并在该范围内近似均匀分布，然后你可以缩放到所需的范围。其实现将选择随机数生成算法的初始种子；它不能由用户选择或重置。
> `Math.random()` 不能提供密码学安全的随机数。请不要使用它们来处理与安全相关的事情。请使用 Web Crypto API 代替，更具体地来说是 [`window.crypto.getRandomValues()`](https://developer.mozilla.org/zh-CN/docs/Web/API/Crypto/getRandomValues) 方法。

可以基于如下公式使用 `Math.random()` 从一组整数中随机选择一个数：
```txt
number = Math.floor(Math.random() * total_number_of_choices + first_possible_value)
```
这里使用了 `Math.floor()` 方法，因为 `Math.random()` 始终返回小数，即便乘以一个数再加上一个数也是小数。因此，如果想从 1~10 范围内随机选择一个数，代码就是这样的：
```js
let num = Math.floor(Math.random() * 10 + 1);
```
这样就有 10 个可能的值（1~10），其中最小的值是 1。如果想选择一个 2~10 范围内的值，则代码就要写成这样：
```js
let num = Math.floor(Math.random() * 9 + 2);
```

#### 其他方法
Math 对象还有很多涉及各种简单或高阶数运算的方法。
| 方法                | 说明                             |
| ------------------- | -------------------------------- |
| Math.abs(x)         | 返回 x 的绝对值                  |
| Math.exp(x)         | 返回 Math.E 的 x 次幂            |
| Math.expm1(x)       | 等于 Math.exp(x) - 1             |
| Math.log(x)         | 返回 x 的自然对数                |
| Math.log1p(x)       | 等于 1 + Math.log(x)             |
| Math.pow(x, power)  | 返回 x 的 power 次幂             |
| Math.hypot(...nums) | 返回 nums 中每个数平方和的平方根 |
| Math.clz32(x)       | 返回 32 位整数 x 的前置零的数量  |
| Math.sign(x)        | 返回表示 x 符号的 1、0、-0 或-1  |
| Math.trunc(x)       | 返回 x 的整数部分，删除所有小数  |
| Math.sqrt(x)        | 返回 x 的平方根                  |
| Math.cbrt(x)        | 返回 x 的立方根                  |
| Math.acos(x)        | 返回 x 的反余弦                  |
| Math.acosh(x)       | 返回 x 的反双曲余弦              |
| Math.asin(x)        | 返回 x 的反正弦                  |
| Math.asinh(x)       | 返回 x 的反双曲正弦              |
| Math.atan(x)        | 返回 x 的反正切                  |
| Math.atanh(x)       | 返回 x 的反双曲正切              |
| Math.atan2(y, x)    | 返回 y/x 的反正切                |
| Math.cos(x)         | 返回 x 的余弦                    |
| Math.sin(x)         | 返回 x 的正弦                    |
| Math.tan(x)         | 返回 x 的正切                    |