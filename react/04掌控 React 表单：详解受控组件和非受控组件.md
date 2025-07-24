在开发过程中，经常涉及到用户输入的表单处理；表单可以分为两种类型：受控表单（Controlled Components）和非受控表单（Uncontrolled Components）。这两种表单在处理用户输入和状态管理时有着不同的方式。例如 input 元素会根据输入内容自动变化，而不是从组件中去获取，这种不受 React 控制的表单元素我们定义为非受控组件。

我们先来创建一个 React 项目，用于下面相关概念的示例演示！下面就使用 vite 和 pnpm 来创建一个 React 应用。
```sh
$ pnpm create vite@latest controlled-vs-uncontrolled -- --template react
```

由于下面的示例会用到一些样式，我们就来集成一下 tailwindcss
- 安装 tailwindcss
```sh
$ pnpm add -D tailwindcss postcss autoprefixer
```
- 初始化配置文件
  ```sh
  $ npx tailwindcss init -p
  ```
  命令执行完成之后，就会在项目的根目录有一个 tailwind.config.js 文件和 postcss.config.js 文件，然后修改 content，完整内容如下：
  ```js
  /** @type {import('tailwindcss').Config} */
  export default {
    content: [
      "./index.html",
      "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
      extend: {},
    },
    plugins: [],
  }
  ```
- 将 index.css 中的内容替换为下面的内容：
  ```css
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  ```
  - base: 这个指定表示导入 tailwindcss 的基本样式，里面会包含一些预设样式，主要目的是重置浏览器的样式，重置了浏览器样式之后，可以保证所有的浏览器中的外观是一致的。
  - components：组件样式，默认情况下没有任何的组件样式，后期我们可以在配置文件里面自定义我们的组件样式，以及使用第三方插件添加一些组件样式，这一条指令是为了让自定义组件样式以及其他第三方组件样式能够生效。
  - utilities：这个指令就是导入实用的原子类。
  
- 可以把原来的 App.css 文件删除，因为我们使用 tailwindcss 就不需要脚手架自动生成的那些 css 文件了，当然也需要把 App.jsx 中引入的 App.css 删掉。不删掉可能会对我们后面的演示内容有一些影响。
- 把App.jsx 中的DOM 内容替换，完整内容如下：
  ```jsx
  import { useState } from 'react'
  
  function App () {
    const [count, setCount] = useState(0)
  
    return (
      <div>
        <h1 className="text-3xl font-bold underline">
          Hello world! {count}
        </h1>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          onClick={() => setCount(count + 1)}
        >
          Click me
        </button>
      </div>
    )
  }
  
  export default App
  ```
- 在终端中执行 `pnpm dev`，在浏览器中访问 `localhost:5173`，效果如下：
<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/bvW2nN-40593189-d8c4-47b1-ad1d-214a4cf018f1.png" />



## 一、基础
前面的环境准备好了，接下来就进入正题！
### 什么是受控组件
受控组件就是由 React 来管理表单元素的值，同时表单元素的变化需要实时映射到 React 的 state 中，这个类似于双向数据绑定。不同的表单元素，React 控制方式是不一样的，如 `input` 用 `value` 来控制，`checkbox` 用 `checked` 来判断是否选中等。

### 什么是非受控组件
非受控组件（Uncontrolled Components）是指表单数据直接由 DOM 元素来管理，而不是通过 React 组件的 state 来管理。换句话说，非受控组件依赖于传统的 HTML 表单处理方式，使用 ref 来直接访问 DOM 元素并获取其值。

## 二、受控组件
### 定义和特征
- 值由 React 状态控制：表单元素的值由组件的状态控制和管理。
- 即时更新：每次用户输入时，都会触发状态更新，从而更新表单元素的值。
- 单一数据源：表单元素的值和组件的状态保持一致，确保数据源的唯一性。
### 受控组件的优劣
#### 优点
- 更好的数据控制和验证
- 更容易实现复杂的交互逻辑

#### 缺点
- 代码量较大
- 需要更多的状态管理

### 示例
下面就来下一个注册页面，其中包含文本框、单选框、复选框和下拉框的受控组件的应用，代码如下：
```jsx
import { useState } from 'react';

const hobby = ['Sports', 'Cooking', 'Traveling'];

const RegistrationForm = () => {
    const [form, setForm] = useState({
        nickname: '',
        age: '',
        password: '',
        sex: 'male', // 默认为male
        hobby: [],
        city: 'beijing',
    });

    const handleInputChange = (e) => {
        const { name, value, type, checked } = e.target;

        const newFormData = { ...form };
        const hobby = [...newFormData.hobby];
        const isCheckbox = type === 'checkbox';

        if (isCheckbox && !checked && hobby.includes(value)) {
            hobby.splice(hobby.indexOf(value), 1)
        } else hobby.push(value)
        newFormData[name] = isCheckbox ? hobby : value;

        setForm(newFormData);
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Form submitted:', form);
        // 在这里你可以添加提交表单的逻辑
    };

    return (
        <form onSubmit={handleSubmit} className="max-w-sm mx-auto">
            <div className="mb-6 w-full px-4">
                <label htmlFor="nickname" className="block mb-2">Nickname:</label>
                <input
                    type="text"
                    id="nickname"
                    name="nickname"
                    value={form.nickname}
                    onChange={handleInputChange}
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="age" className="block mb-2">Age:</label>
                <input
                    type="number"
                    id="age"
                    name="age"
                    value={form.age}
                    onChange={handleInputChange}
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="password" className="block mb-2">Password:</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    value={form.password}
                    onChange={handleInputChange}
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label className="block mb-2">Sex:</label>
                <label className="inline-flex items-center mr-4">
                    <input
                        type="radio"
                        name="sex"
                        value="male"
                        checked={form.sex === 'male'}
                        onChange={handleInputChange}
                    />
                    <span className="ml-2">Male</span>
                </label>
                <label className="inline-flex items-center">
                    <input
                        type="radio"
                        name="sex"
                        value="female"
                        checked={form.sex === 'female'}
                        onChange={handleInputChange}
                    />
                    <span className="ml-2">Female</span>
                </label>
            </div>

            <div className="mb-6 w-full px-4">
                <fieldset>
                    <legend className="block mb-2">Hobby:</legend>
                    {hobby.map(item => (<label className="inline-flex items-center mr-4" key={item}>
                        <input
                            type="checkbox"
                            name='hobby'
                            value={item}
                            checked={form.hobby.includes(item)}
                            onChange={handleInputChange}
                        />
                        <span className="ml-2">{item}</span>
                    </label>))}

                </fieldset>
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="city" className="block mb-2">City:</label>
                <select
                    id="city"
                    name="city"
                    value={form.city}
                    onChange={handleInputChange}
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                >
                    <option value="beijing">Beijing</option>
                    <option value="shanghai">Shanghai</option>
                    <option value="guangzhou">Guangzhou</option>
                    <option value="sichuan">Sichuan</option>
                </select>
            </div>

            <button
                type="submit"
                className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-700"
            >
                Register
            </button>
        </form>
    );
};

export default RegistrationForm;
```
效果如下：
<img src="https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/iDWx4Z-00535ab5-618e-44f3-8270-8596a7c9936f.png" />


上面的例子通过 `useState` 钩子定义和初始化表单的状态对象 form，其中包含字段：`nickname`、`age`、`password`、`sex`、`hobby`、`city`。
- 昵称、年龄、密码：这些文本和数字输入框的值通过 `value` 属性与状态绑定，通过 `onChange` 事件处理函数更新状态。
- 性别：单选按钮组的选中状态通过 `checked` 属性与状态绑定，通过 `onChange` 事件处理函数更新状态。
- 兴趣爱好：复选框的选中状态通过 `checked` 属性与状态绑定，通过 `onChange` 事件处理函数更新状态。函数会根据复选框的选中与否更新状态中的 `hobby` 数组。
- 城市：下拉选择框的值通过 `value` 属性与状态绑定，通过 `onChange` 事件处理函数更新状态。

## 三、非受控组件
### 特点
- 表单数据不受 React 控制：非受控组件中的表单元素值并不保存在 React 的 state 中，而是直接从 DOM 元素中获取。
- 使用 ref 访问 DOM：通过 `ref` 可以直接引用 DOM 元素并读取其值。
### 优点
- 更少的代码量：由于不需要维护 state，非受控组件的代码量较少。
- 状态管理简单：无需通过事件处理函数来更新 state，状态管理相对简单。
### 缺点
- 数据同步较难：因为表单数据不保存在state中，数据同步和验证相对较难。
- 处理复杂交互逻辑不便：在处理复杂的表单交互逻辑时，非受控组件较为不便。

### 示例
把上面的受控组件改成非受控组件，完整代码如下：
```json
const hobby = ['Sports', 'Cooking', 'Traveling'];

const RegistrationFormUncontrolled = () => {
    // 处理表单提交
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Form submitted:', e);
        const nickname = e.target.nickname.value;
        const age = e.target.age.value;
        const sex = e.target.sex.value;
        const password = e.target.password.value;
        const hobby = e.target.hobby.value;
        const city = e.target.city.value;
        console.log('🚀 ~ handleSubmit ~ formData:', nickname, age, sex, password, hobby, city)
        // 在这里你可以添加提交表单的逻辑
    };

    return (
        <form onSubmit={handleSubmit} className="max-w-sm mx-auto">
            <div className="mb-6 w-full px-4">
                <label htmlFor="nickname" className="block mb-2">Nickname:</label>
                <input
                    type="text"
                    id="nickname"
                    name="nickname"
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="age" className="block mb-2">Age:</label>
                <input
                    type="number"
                    id="age"
                    name="age"
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="password" className="block mb-2">Password:</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                />
            </div>

            <div className="mb-6 w-full px-4">
                <label className="block mb-2">Sex:</label>
                <label className="inline-flex items-center mr-4">
                    <input
                        type="radio"
                        name="sex"
                        value="male"
                    />
                    <span className="ml-2">Male</span>
                </label>
                <label className="inline-flex items-center">
                    <input
                        type="radio"
                        name="sex"
                        value="female"
                    />
                    <span className="ml-2">Female</span>
                </label>
            </div>

            <div className="mb-6 w-full px-4">
                <fieldset>
                    <legend className="block mb-2">Hobby:</legend>
                    {hobby.map(item => (<label className="inline-flex items-center mr-4" key={item}>
                        <input
                            type="checkbox"
                            name='hobby'
                            value={item}
                        />
                        <span className="ml-2">{item}</span>
                    </label>))}

                </fieldset>
            </div>

            <div className="mb-6 w-full px-4">
                <label htmlFor="city" className="block mb-2">City:</label>
                <select
                    id="city"
                    name="city"
                    className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 ps-4 py-[10px] text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2"
                >
                    <option value="beijing">Beijing</option>
                    <option value="shanghai">Shanghai</option>
                    <option value="guangzhou">Guangzhou</option>
                    <option value="sichuan">Sichuan</option>
                </select>
            </div>

            <button
                type="submit"
                className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-700"
            >
                Register
            </button>
        </form>
    );
};

export default RegistrationFormUncontrolled;
```
效果也是一样的：

![QQ_1721729411124](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/FElCnP-668c460b-1cdd-4df0-bd28-252a1e4a6192.png)

不过获取数据的方式不太一样，使用非受控表单获取数据的几种方法：
- 使用 ref 获取表单数据
  ```jsx
  import React, { useRef } from 'react';
  
  function UncontrolledForm() {
    const nameRef = useRef(null);
    const emailRef = useRef(null);
  
    const handleSubmit = (e) => {
      e.preventDefault();
      const name = nameRef.current.value;
      const email = emailRef.current.value;
      alert(`Name: ${name}, Email: ${email}`);
    };
  
    return (
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            Name:
            <input type="text" ref={nameRef} />
          </label>
        </div>
        <div>
          <label>
            Email:
            <input type="email" ref={emailRef} />
          </label>
        </div>
        <button type="submit">Submit</button>
      </form>
    );
  }
  
  export default UncontrolledForm;
  ```
- 直接访问 event.target 获取表单数据
  ```jsx
  const handleSubmit = (e) => {
      e.preventDefault();
      const name = e.target.name.value;
      const email = e.target.email.value;
      alert(`Name: ${name}, Email: ${email}`);
  };
  ```
- 使用 FormData 获取表单数据
  ```jsx
  const handleSubmit = (e) => {
      e.preventDefault();
      const formData = new FormData(e.target);
      const name = formData.get('name');
      const email = formData.get('email');
      alert(`Name: ${name}, Email: ${email}`);
  };
  ```
## 四、受控组件与非受控组件的对比
### 使用场景
#### 受控组件适用场景
1. **表单处理**
    - *实时验证*：可以在用户输入时立即进行验证（如字符长度、格式）；例如，验证电子邮件格式、密码强度等。
    - *动态表单*：根据用户输入动态调整表单内容，如显示/隐藏额外的输入字段。  例如，选择某个选项后显示额外的详细信息输入框。
    - *多步骤表单*：管理多步骤表单的状态和步骤之间的数据传递。例如，分步注册表单，逐步收集用户信息。
2. **统一状态管理**
    - *状态同步*：输入状态与组件状态保持同步，确保数据的一致性。例如，相同的输入数据在多个组件间共享。
    - *可预测性*：状态变化可控，便于追踪和调试。例如，调试用户输入导致的状态变化问题。
3. **复杂交互**
    - *动态内容更新*：根据用户输入实时更新显示内容。例如，搜索框实时显示搜索结果。
    - *用户反馈*：根据输入情况给用户即时反馈，如错误提示、成功消息等。例如，输入有效时显示绿色勾选图标，无效时显示红色叉号。
4. **一致的用户体验**
    - 统一样式：通过状态管理和样式绑定，确保所有输入元素的样式一致。例如，所有输入框获得焦点时边框颜色一致。
    - 行为一致：通过状态控制确保所有输入元素的交互行为一致。例如，所有输入框在失去焦点时进行验证。
5. **数据绑定**
    - *API数据绑定*：将用户输入的数据与后端API数据进行绑定和同步。例如，用户输入地址信息时实时查询和显示地址建议。
    - *本地存储同步*：将用户输入的数据与本地存储（如localStorage）进行同步，保持数据持久性。例如，表单数据自动保存到本地存储，防止页面刷新导致的数据丢失。
      
#### 非受控组件适用场景
1. **简单表单**；对于简单表单，不需要复杂状态管理的场景，可以使用非受控组件。
    - *快速实现*：适用于简单的表单或输入元素，不需要实时验证或动态更新。例如，单个输入框或简单的留言表单。
    - *低频交互*：适用于用户交互频率较低的表单。例如，静态内容提交或一次性表单提交。

2. **第三方库集成**；在使用某些第三方库时，非受控组件可能更适合，因为这些库可能会直接操作DOM节点。

    - *文件上传*：使用非受控组件处理文件上传，因为文件输入通常直接与DOM节点交互。例如，使用`<input type="file">`进行文件选择。
    - *绘图或富文本编辑*：第三方绘图库或富文本编辑器可能需要直接访问DOM节点。例如，集成Quill富文本编辑器或Canvas绘图库。

3. **性能优化**；在某些情况下，使用非受控组件可以减少不必要的重渲染，提高性能。
    - *避免过多状态更新*：对于高频输入场景，非受控组件可以减少状态更新次数和组件重渲染。例如，实时输入大文本或快速输入场景。
    - *减少渲染开销*：避免由于状态变更导致的频繁渲染，提升性能。例如，处理较大数据或复杂表单时。

4. **简化代码**，对于一些场景，使用非受控组件可以简化代码，降低复杂度。

    - *减少状态管理*：不需要维护复杂的状态管理逻辑，表单数据直接从DOM中获取。例如，简单的搜索框或单个输入框。
    - *简洁代码*：代码更加简洁明了，减少不必要的状态和事件处理。例如，快速创建简单的输入表单。

## 五、最佳实践
### 选择合适的组件类型
在项目中选择使用受控组件还是非受控组件可以根据以下几个关键因素进行考虑：
| 关键因素         | 受控组件适用场景                                             | 非受控组件适用场景                                           |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **表单复杂性**   | 复杂表单：需要处理许多交互和验证逻辑<br>实时验证：需要实时验证用户输入 | 简单表单：输入字段较少且不需要复杂的状态管理<br>低频交互：不需要实时反馈和验证 |
| **性能考虑**     | 高频交互：可能导致过多状态更新，需优化<br>可预测性：状态变化可控，便于调试 | 避免过多状态更新：减少不必要的状态更新<br>减少渲染开销：直接操作DOM节点 |
| **第三方库集成** | 受限：某些第三方库可能不适合与受控组件配合使用               | 灵活性：适合需要直接操作DOM节点的第三方库，如文件上传、富文本编辑器 |
| **数据初始化**   | 预填充数据：组件挂载时预填充数据，并实时更新状态             | 延迟初始化：组件挂载后从外部源获取初始数据并填充表单         |
| **代码简洁性**   | 全面控制：提供对表单和输入状态的全面控制，但需编写较多状态管理和事件处理代码 | 简洁实现：适用于不需要复杂状态管理的场景，代码实现简洁明了   |

### 何时以及如何混合使用受控和非受控组件
混合使用受控和非受控组件可以在特定场景下带来最佳的用户体验和性能优化。

### 何时混合使用：
- 性能优化：在高频交互的场景下，某些输入字段使用非受控组件以减少不必要的状态更新，而其他需要严格控制和验证的字段使用受控组件。
- 复杂表单：对于包含多个输入字段的复杂表单，部分字段需要实时验证和反馈（使用受控组件），而其他字段可以直接与DOM交互（使用非受控组件）。
- 第三方库集成：当使用需要直接操作DOM的第三方库（如文件上传、富文本编辑器）时，可以将这些部分作为非受控组件，而将其他部分作为受控组件。

#### 混合使用
- 高频输入和实时验证
  ```jsx
  import { useState, useRef } from 'react';
  
  function MixedForm () {
      const [name, setName] = useState('');
      const passwordRef = useRef(null);
  
      const handleNameChange = (e) => {
          setName(e.target.value);
      };
  
      const handleSubmit = (e) => {
          e.preventDefault();
          const email = passwordRef.current.value;
          alert(`Name: ${name}, Email: ${email}`);
      };
  
      return (
          <div className="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
              <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                  <img className="mx-auto h-10 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company" />
                  <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Sign in to your account</h2>
              </div>
  
              <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                  <form className="space-y-6" action="#" method="POST" onSubmit={handleSubmit}>
                      <div>
                          <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">Email address</label>
                          <div className="mt-2">
                              {/* 受控组件 */}
                              <input id="email" name="email" type="email" autoComplete="email" required className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" onChange={handleNameChange} />
                          </div>
                      </div>
  
                      <div>
                          <div className="flex items-center justify-between">
                              <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">Password</label>
                              <div className="text-sm">
                                  <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">Forgot password?</a>
                              </div>
                          </div>
                          <div className="mt-2">
                              {/* 非受控组件 */}
                              <input ref={passwordRef} id="password" name="password" type="password" autoComplete="current-password" required className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                          </div>
                      </div>
  
                      <div>
                          <button type="submit" className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Sign in</button>
                      </div>
                  </form>
  
                  <p className="mt-10 text-center text-sm text-gray-500">
                      Not a member?
                      <a href="#" className="font-semibold leading-6 text-indigo-600 hover:text-indigo-500">Start a 14 day free trial</a>
                  </p>
              </div>
          </div>
      );
  }
  
  export default MixedForm;
  ```
- 文件上传和其他输入
  ```jsx
  import { useState, useRef } from 'react';
  
  function FileUploadForm () {
      const [username, setUsername] = useState('');
      const fileInputRef = useRef(null);
  
      const handleUsernameChange = (e) => {
          setUsername(e.target.value);
      };
  
      const handleSubmit = (e) => {
          e.preventDefault();
          const file = fileInputRef.current.files[0];
          alert(`Username: ${username}, File: ${file ? file.name : 'No file selected'}`);
      };
  
      return (
          <form onSubmit={handleSubmit}>
              {/* 受控组件 */}
              <div className="w-full px-4 md:w-1/2 lg:w-1/3">
                  <div className="mb-12">
                      <label htmlFor="" className="mb-[10px] block text-base font-medium text-dark dark:text-white">
                          受控组件
                      </label>
                      <input type="text" placeholder="entry your username" value={username} onChange={handleUsernameChange} className="w-full bg-transparent rounded-md border border-stroke dark:border-dark-3 py-[10px] px-5 text-dark-6 outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-gray-2 disabled:border-gray-2" />
                  </div>
              </div>
  
              {/* 非受控组件 */}
              <div className="w-full px-4 md:w-1/2 lg:w-1/3">
                  <div className="mb-12">
                      <label className="mb-[10px] block text-base font-medium text-dark dark:text-white">
                          非受控组件
                      </label>
                      <div className="relative">
                          <label htmlFor="file" className="flex min-h-[175px] w-full cursor-pointer items-center justify-center rounded-md border border-dashed border-primary p-6">
                              <div>
                                  <input type="file" name="file" id="file" ref={fileInputRef} className="sr-only" />
                                  <span className="mx-auto mb-3 flex h-[50px] w-[50px] items-center justify-center rounded-full border border-stroke dark:border-dark-3 bg-white dark:bg-dark-2">
                                      <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                                          <path fillRule="evenodd" clipRule="evenodd" d="M2.5013 11.666C2.96154 11.666 3.33464 12.0391 3.33464 12.4993V15.8327C3.33464 16.0537 3.42243 16.2657 3.57871 16.4219C3.73499 16.5782 3.94695 16.666 4.16797 16.666H15.8346C16.0556 16.666 16.2676 16.5782 16.4239 16.4219C16.5802 16.2657 16.668 16.0537 16.668 15.8327V12.4993C16.668 12.0391 17.0411 11.666 17.5013 11.666C17.9615 11.666 18.3346 12.0391 18.3346 12.4993V15.8327C18.3346 16.4957 18.0712 17.1316 17.6024 17.6004C17.1336 18.0693 16.4977 18.3327 15.8346 18.3327H4.16797C3.50493 18.3327 2.86904 18.0693 2.4002 17.6004C1.93136 17.1316 1.66797 16.4957 1.66797 15.8327V12.4993C1.66797 12.0391 2.04106 11.666 2.5013 11.666Z" fill="#3056D3"></path>
                                          <path fillRule="evenodd" clipRule="evenodd" d="M9.41074 1.91009C9.73618 1.58466 10.2638 1.58466 10.5893 1.91009L14.7559 6.07676C15.0814 6.4022 15.0814 6.92984 14.7559 7.25527C14.4305 7.58071 13.9028 7.58071 13.5774 7.25527L10 3.67786L6.42259 7.25527C6.09715 7.58071 5.56951 7.58071 5.24408 7.25527C4.91864 6.92984 4.91864 6.4022 5.24408 6.07676L9.41074 1.91009Z" fill="#3056D3"></path>
                                          <path fillRule="evenodd" clipRule="evenodd" d="M10.0013 1.66602C10.4615 1.66602 10.8346 2.03911 10.8346 2.49935V12.4994C10.8346 12.9596 10.4615 13.3327 10.0013 13.3327C9.54106 13.3327 9.16797 12.9596 9.16797 12.4994V2.49935C9.16797 2.03911 9.54106 1.66602 10.0013 1.66602Z" fill="#3056D3"></path>
                                      </svg>
                                  </span>
                                  <span className="text-base text-body-color dark:text-dark-6">
                                      Drag &amp; drop or
                                      <span className="text-primary underline"> browse </span>
                                  </span>
                              </div>
                          </label>
                      </div>
                  </div>
              </div>
  
              <button type="submit" className="border-dark dark:border-dark-2 border rounded-md inline-flex items-center justify-center py-3 px-7 text-center text-base font-medium text-dark dark:text-white hover:bg-gray-4 dark:hover:bg-dark-3 disabled:bg-gray-3 disabled:border-gray-3 disabled:text-dark-5">Submit</button>
          </form>
      );
  }
  
  export default FileUploadForm;
  ```
### 状态管理
在使用受控组件时，高效的状态管理可以提升应用的性能和代码的可读性。以下是一些最佳实践和策略，可以帮助你在React应用中更高效地管理状态：
1. 使用React Hooks

React Hooks（例如 useState 和 useReducer）是管理组件状态的核心工具。

- useState

  适用于简单的状态管理，例如单个表单输入字段。

  ```jsx
  import React, { useState } from 'react';
  
  function SimpleForm() {
    const [name, setName] = useState('');
  
    const handleNameChange = (e) => {
      setName(e.target.value);
    };
  
    return (
      <form>
        <input type="text" value={name} onChange={handleNameChange} />
      </form>
    );
  }
  ```
  
- useReducer

  适用于更复杂的状态管理，例如包含多个字段的表单或需要处理多个状态更新逻辑。

  ```jsx
  import React, { useReducer } from 'react';
  
  const initialState = { name: '', email: '' };
  
  function reducer(state, action) {
    switch (action.type) {
      case 'SET_NAME':
        return { ...state, name: action.payload };
      case 'SET_EMAIL':
        return { ...state, email: action.payload };
      default:
        return state;
    }
  }
  
  function ComplexForm() {
    const [state, dispatch] = useReducer(reducer, initialState);
  
    const handleNameChange = (e) => {
      dispatch({ type: 'SET_NAME', payload: e.target.value });
    };
  
    const handleEmailChange = (e) => {
      dispatch({ type: 'SET_EMAIL', payload: e.target.value });
    };
  
    return (
      <form>
        <input type="text" value={state.name} onChange={handleNameChange} />
        <input type="email" value={state.email} onChange={handleEmailChange} />
      </form>
    );
  }
  ```
  
2. 将状态提升

    将状态提升到最近的公共祖先组件中，以便多个子组件共享和同步状态。

    ```jsx
    import React, { useState } from 'react';
    
    function ParentComponent() {
      const [formData, setFormData] = useState({ name: '', email: '' });
    
      const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({ ...prevData, [name]: value }));
      };
    
      return (
        <>
          <ChildComponent1 formData={formData} handleChange={handleChange} />
          <ChildComponent2 formData={formData} handleChange={handleChange} />
        </>
      );
    }
    
    function ChildComponent1({ formData, handleChange }) {
      return (
        <input type="text" name="name" value={formData.name} onChange={handleChange} />
      );
    }
    
    function ChildComponent2({ formData, handleChange }) {
      return (
        <input type="email" name="email" value={formData.email} onChange={handleChange} />
      );
    }
    ```
    
3. 使用自定义Hooks
  
    自定义Hooks可以封装和重用复杂的状态逻辑。

    ```jsx
    import React, { useState } from 'react';
    
    function useForm(initialState) {
      const [formData, setFormData] = useState(initialState);
    
      const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({ ...prevData, [name]: value }));
      };
    
      return [formData, handleChange];
    }
    
    function MyForm() {
      const [formData, handleChange] = useForm({ name: '', email: '' });
    
      return (
        <form>
          <input type="text" name="name" value={formData.name} onChange={handleChange} />
          <input type="email" name="email" value={formData.email} onChange={handleChange} />
        </form>
      );
    }
    ```
    
4. 使用状态管理库

对于大型应用，可以考虑使用状态管理库，如 Redux 或 Zustand。

### 社区的 React form 方案
- [react-hook-form](https://github.com/react-hook-form/react-hook-form)
- [formik](https://github.com/jaredpalmer/formik)
- [react-jsonschema-form](https://github.com/rjsf-team/react-jsonschema-form)
- [@tanstack/react-form](https://tanstack.com/form/latest/docs/overview)
- [@formily/react](https://formilyjs.org/zh-CN/guide)
- [redux-form](https://redux-form.com/8.3.0/)