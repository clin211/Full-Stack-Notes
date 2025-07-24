前面的文章介绍了 `Script` 组件和 `Link` 组件，今天我们来看看 `Font` 组件。在 Web 开发中，字体是一个常常被忽视但至关重要的元素。它不仅影响网站的可读性和用户体验，还体现了品牌的个性和风格。选择合适的字体，可以提高文本的可读性，帮助用户更好地理解内容，并传达品牌的独特气质。良好的字体选择不仅可以改善网站的美观度，还可以吸引用户在网站上停留更长时间。另外，选择可缩放且适应不同屏幕的字体，也是确保网站在各种设备上保持良好可读性的关键。同时，使用轻量级 Web 字体还可以加快页面加载速度，提高网站的性能。

## 传统字体
在Web设计中，字体是实现视觉表达和用户体验的关键，但与印刷设计不同，Web设计不受物理尺寸或颜色数量的限制，但带宽成本是一个不容忽视的因素。在 Web 上实现丰富的排版体验时，字体的使用尤为突出。传统的 Web 字体要求每个样式（如常规、粗体、斜体）都需要单独的字体文件，这可能导致页面加载时间增加和用户体验下降。例如，仅包含常规和粗体样式以及它们对应的斜体，字体数据可能达到好几百 kb 或更多的字体数据。

传统字体的使用很方便，通过 `@font-face` 指定一个自定义的字体，字体文件可以本地文件或者远程文件，如果提供了 `local()` 函数，从用户本地查找指定的字体名称，并且找到了一个匹配项，本地字体就会被使用。否则，字体就会使用 `url()` 函数下载的资源。然后使用 `font-family` 来标识“系列”，比如：
```css
@font-face {
    font-family: "Open Sans";
    src:
        url("/fonts/OpenSans-Regular-webfont.woff2") format("woff2"),
        url("/fonts/OpenSans-Regular-webfont.woff") format("woff");
}
```

具体的使用方法如下：
```html
<html lang="en">

<head>
    <title>Web Font Sample</title>
    <style type="text/css" media="screen, print">
        @font-face {
            font-family: "Bitstream Vera Serif Bold";
            src: url("https://mdn.github.io/css-examples/web-fonts/VeraSeBd.ttf");
        }

        body {
            font-family: "Bitstream Vera Serif Bold", serif;
        }
    </style>
</head>

<body>
    This is Bitstream Vera Serif Bold.
</body>

</html>
```
效果如下：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/ZetVRx-90860e94-1d90-407c-895d-ad7540ed5fe6-3293774.png)

## 可变字体

![来自iconfont的阿里妈妈灵动体](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/YPLnPJ-18ca4dcc-5ad1-4126-ac6e-9a7313bec7be-3293774.gif)


在可变字体出现之前，每个不同的字体风格都需要单独的文件，这限制了设计师在 Web 上使用丰富字体的能力。可变字体允许在一个单一的文件中包含多个字体风格（如上图），从而减少了文件大小和请求数量，提升了页面加载速度和性能。想了解更多的可变字体相关的内容，可以阅读《[Introducing OpenType Variable Fonts](https://medium.com/variable-fonts/https-medium-com-tiro-introducing-opentype-variable-fonts-12ba6cd2369)》；当然也可以看 [大漠_w3cpluscom](https://juejin.cn/user/1908407916041614/course) 大佬在掘进上的专栏《[现代 CSS](https://s.juejin.cn/ds/iP43u3L7/)》的 Web 上的可变字体。 


在 Web 上使用可变字体可以分以下几个步骤。

- 步骤一：获取可用的可变字体。
- 步骤二：在 CSS 中集成可变字体。
- 步骤三：找出可变字体的轴和范围值。
- 步骤四：设置可变字体样式。
- 步骤五：降级处理。

在 Google Fonts 可以查看到很多可变字体，筛选条件也很简单（选中左侧的 Variable），如下图：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/DDLgNX-bfdffed3-f7b0-4012-80eb-3a80f9acce6d-3293774.png)
每个字体的右上角也可以区分哪些字体是可变字体：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/YJUAFI-d7306a61-a331-4f88-b15e-ce24054c69f8-3293774.png)
选中一个可变字体之后，进去后可以调整可变字体的版本，然后点击右上角的 “Get font”按钮。如下图：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/eprfu3-5ab83dc0-41cb-4dc7-852e-1d0236870e30-3293774.png)

可以直接获取代码或者下载字体文件：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/EiNvTV-e5b7fdb5-e4ae-4c27-87ac-b27898ba92d8-3293774.png)

如果选择以远程字体文件引入的方式，则点击 “Get Embed code”后，就可以选择对应的平台的引入代码，如下图：
![QQ_1739785067108](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/VzRPpB-c0e2b57b-4ebc-43ad-af3a-7419aa41db2a-3293774.png)

整体示例如下：
```html
<html lang="en">

<head>
    <title>Web Font Sample</title>
    <link href="https://fonts.googleapis.com/css2?family=Oswald:wght@200;300;400;469;500;600;681;700&display=swap"
        rel="stylesheet">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto:ital,wdth,wght@0,94,280;1,94,280&display=swap"
        rel="stylesheet">
    <style type="text/css" media="screen, print">
        @font-face {
            font-family: "Bitstream Vera Serif Bold";
            src: url("https://mdn.github.io/css-examples/web-fonts/VeraSeBd.ttf");
        }

        body {
            font-family: "Bitstream Vera Serif Bold", serif;
        }

        .custom-font {
            font-family: 'Oswald';
            font-weight: 681;
            color: red;
        }

        .roboto-font {
            font-family: "Roboto", serif;
            font-optical-sizing: auto;
            font-weight: 400;
            font-style: normal;
            font-variation-settings:
                "wdth" 494;
        }
    </style>
</head>

<body>
    This is Bitstream Vera Serif Bold.
    <p class="custom-font">
        This is Oswald.
    </p>
    <p class="roboto-font">Whereas recognition of the inherent dignity
    </p>
</body>

</html>
```
效果如下图：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/KgosrM-5856d675-0b82-45f9-89dd-dba97b8f1885-3293774.png)

有了上面的理论基础后，按照前面的惯例，在进入正题之前，我们先来准备一下相关的环境，创建一个项目便于演示后面的内容。

使用命令 `npx create-next-app@latest --use-pnpm` 创建一个新的项目；具体的项目配置选项如下：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/H4B0Ru-34f3ecd4-4317-492a-9246-63e67182829d-3293774.png)

在 VS Code 打开以后，项目的依赖如下图：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/Zh770s-eda9c11c-3c33-45f8-9b04-0e87e84f6ab5-3293774.png)

我们下面就来看看 Next.js 做了哪些优化？在 Next.js 中怎么用可变字体？

## Font 组件
如果将上面的传统字体的示例代码迁移到 Next.js 项目中，在 Next.js 中的引入方式有：
- 通过 `<link href="https://fonts.googleapis.com/css2?family=Roboto:ital,wdth,wght@0,88,280;1,88,280&display=swap" rel="stylesheet">`。

- 通过 css 文件的方式引入，比如： `import "./global.css";`，在 `global.css` 文件中使用 `@import url()` 引入。
- 通过内置组件 `next/font` 包的 `localFont` 加载或者 Google Font 系列字体来加载。
  

可变字体在 Next.js 中也比较方便，Next.js 内置了 `next/font` 组件，帮助我们更好的管理和使用字体，而且 Next.js 官方还做了不少优化，其中就包括布局偏移的问题。`next/font` 具体又分为 `next/font/google` 和 `next/font/local`，分别对应使用 [Google 字体](https://fonts.google.com/?vfaxis=wdth)和使用本地字体。

### Google 字体
```jsx
// 加载 Google Font 系列的字体
import { Inter, Roboto_Mono } from 'next/font/google'

export const inter = Inter({
    subsets: ['latin'],
    display: 'swap',
})

export const roboto_mono = Roboto_Mono({
    subsets: ['latin'],
    display: 'swap',
})
```
> 上面这段字体加载的代码中，需要注意的是：如果字体是多单词，使用下划线 `_` 连接，比如 `Roboto Mono`，导入的时候写成 `Roboto_Mono`。

最终渲染后的效果如下：
![QQ_1739788489179](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/g2BnEw-b8318790-1e83-46ad-85e7-f1de7d168e18-3293774.png)
Next.js 推荐使用可变字体来获得最佳的性能和灵活性。如果不能使用可变字体，你需要声明 `weight`（字重，是指字体的粗细程度):
```jsx
import { Roboto } from 'next/font/google'

const roboto = Roboto({
    weight: '400',
    subsets: ['latin'],
    display: 'swap',
})

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en" className={roboto.className} >
            <body>{children} </body>
        </html>
    )
}
```
上面这段代码中，需要注意 `subsets` 这个属性，谷歌的字体是可以指定子集（[subset](https://fonts.google.com/knowledge/glossary/subsetting)）的，目前 Next.js 官方网站上也没有具体指出哪些字体支持那些子集，可以通过 next.js 源码搜字体查看，在 [packages/font/src/google/index.ts](https://github.com/vercel/next.js/blob/canary/packages/font/src/google/index.ts) 中；还可以随便写一个，然后通过 TS 的类型推导校正，如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/UhRdHQ-edb26982-e97f-46e0-a304-664644bba38a-3293774.png)

### 本地字体
- 多字体的加载方式
  ```jsx
  import localFont from 'next/font/local'；
  
  const roboto = localFont({
      src: [
          {
              path: './Roboto-Regular.woff2',
              weight: '400',
              style: 'normal',
          },
          {
              path: './Roboto-Italic.woff2',
              weight: '400',
              style: 'italic',
          },
          {
              path: './Roboto-Bold.woff2',
              weight: '700',
              style: 'normal',
          },
          {
              path: './Roboto-BoldItalic.woff2',
              weight: '700',
              style: 'italic',
          },
      ],
  })
  
  export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
      return (
          <html lang="en" className={roboto.className}>
              <body>{children}</body>
          </html>
      )
  }
  ```
- 单字体加载方式
  ```jsx
  import localFont from 'next/font/local'
  
  const manrope = localFont({
      src: './fonts/Manrope.woff2',
      display: 'swap',
  })
  
  export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
      return (
          <html lang="en" className={manrope.className}>
              <body>{children}</body>
          </html>
      )
  }
  ```

### 字体函数的相关参数
上面介绍了 Google 字体和本地字体的相关用法和注意点，接下来看看字体函数的相关参数。


| 属性               | font/google | font/local | 类型             | 必传 |
| ------------------ | ----------- | ---------- | ---------------- | ---- |
| src                | ❌           | ✅          | 字符串或对象数组 | 是   |
| weight             | ✅           | ✅          | 字符串或数组     | 可选 |
| style              | ✅           | ✅          | 字符串或数组     | -    |
| subsets            | ✅           | ❌          | 字符串数组       | -    |
| axes               | ✅           | ❌          | 字符串数组       | -    |
| display            | ✅           | ✅          | 字符串           | -    |
| preload            | ✅           | ✅          | 布尔             | -    |
| fallback           | ✅           | ✅          | 字符串数组       | -    |
| adjustFontFallback | ✅           | ✅          | 布尔或者字符串   | -    |
| variable           | ✅           | ✅          | 字符串           | -    |
| declarations       | ❌           | ✅          | 对象数组         | -    |
- **src**

  这个属性只有在 `next/font/local` 中必传，它可以是一个字符串，也可以是对象数组，其类型为：`Array<{path: string, weight?: string, style?: string}>`。

  如果是字符串的话，路径地址相当于字体加载函数调用的位置；比如：比如上面带代码 `app/page.tsx` 中使用 `src:'./fonts/Manrope.woff2'` 调用字体加载函数，`Manrope.woff2` 就要放置在 `app/fonts/` 下。

- **weight**

  字体字重，类似于 [`font-weight`](https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-weight)。如果是可变字体，则非必传，如果不是可变字体，则必传。它的值可以是字符串或字符串数组：
  - 单个字符串： `weight: '400'`、`weight: '100 900'`（表示可变字体取值范围为  100 ~ 900）。
  - 字符串数组：`weight: ['100', '400', '900']`（表示不可变字体的三个字重值）。
  
- **style**
  
  字体风格，类似于 [`font-style`](https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-style)。一共有三个可选值：`italic`、`oblique`、`normal`；默认值是：`normal`。
  
  如果使用 `next/font/google` 的非可变字体，也可以传入一组样式值，如 `style: ['italic','normal']`。
  
  它的值同样可以是字符串或字符串数组：
  - 单个字符串：`style: 'italic'` 或者 `style: 'oblique'` 或者 `style: 'normal'`。
  - 字符串数组：`style: ['italic', 'normal']`。
  
- **subsets**

  这个属性在上面也有使用过，字体子集（[font subsets](https://fonts.google.com/knowledge/glossary/subsetting)）可以通过一个字符串数组来定义，数组中包含每个子集的名称。通过子集指定的字体将在 `preload` 选项为 `true` 时（默认值）在 HTML 头部注入一个预加载属性的 `link` 标签。如果 `preload` 为 `true`，但不指定子集会有警告。有些字体只有一个默认子集，比如 `latin`，也需要手动制定。
  
  > `preload` 选项默认为 `true`，如果你不想预加载字体，可以将其设置为 `false`。
  
  使用方式也比较简单：
  ```jsx
  import { Roboto } from 'next/font/google'
  
  const roboto = Roboto({
      weight: '400',
      subsets: ['latin'],
      display: 'swap',
  })
  ```
  
- **axes**
  
  可变字体（Variable Fonts）通过内置的多种变形轴（axis 的复数形式 axes），允许用户动态调整字体的视觉特性。以字宽（Width）、字重（Weight）、倾斜（Slant）等为代表的变形轴，本质上是对字体形态的数字化控制参数，用户可通过调节这些轴值，使单一字体文件呈现出传统多款独立字体的效果。
  
  比如，[Inter 字体](https://fonts.google.com/variablefonts?vfquery=inter#font-families)，不仅支持基础的「字重轴」（调节笔画粗细）和「字宽轴」（调节字符横向比例），还包含「倾斜轴」（Slant，模拟斜体效果）或「光学尺寸轴」（Optical Size，适配不同显示环境）等参数，从而通过多维度调整实现更灵活的排版表现。这种技术将传统静态字体的离散样式转化为连续可变的动态系统，显著提升了设计效率与视觉一致性。
  
  
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/5Xgf9N-a283126b-42ad-48e7-a522-5157541ca813-3293774.png)

  上图的 `ital` 是斜体（italic）、`opsz` 是光学尺寸轴（Optical Size）、`wght` 是字重（wight）。
  
  `axes` 的值是一个字符串数组形式，比如 `axes: ['slnt', 'wght']`，你可以在 [Google 可变字体](https://fonts.google.com/variablefonts#font-families)页面查询字体的 Axes 有哪些(如上图)。默认情况下，只有 `weight` 轴会被留下（只留一个的目的是减少字体文件大小），如果需要其他的轴就需要单独声明。

- **display**

  字体显示策略，类似于（[font-display](https://developer.mozilla.org/zh-CN/docs/Web/CSS/@font-face/font-display)）；用于控制页面加载时 Web Fonts 引起的布局抖动和偏移的优化策略。它的可选值有 `'auto'`、`'block'`、`'swap'`、`'fallback'` 或 `'optional'`， 在 Next.js 的 `next/font` 组件中的默认值是 `swap`，在 Web CSS 中 `font-display` 的默认值是 `auto`。
  
  - `font-display: auto`：使用浏览器的预设值，一般是 `block`。
  - `font-display: block` 会使文本在网页字体加载期间隐藏（最长3秒），超时后自动用备用字体显示，待网页字体加载完成后再切换回原字体。
  - `font-display: swap` 会立即用备用字体显示文本（无加载等待），待网页字体加载完成后自动切换。优势是内容即时可见，但需确保备用字体与目标字体形态相近，以减少切换时的布局抖动或偏移的风险。
  - `font-display: fallback`：先显示空白（大约 100ms），然后切换为备用字体，时间大概是 3s，3s 内能加载完字体，就使用字体，3s 内加载不完，后续接着使用备用字体。
  - `font-display: optional`：先显示空白（大约 100ms），100ms 内能加载完就用，加载不完就直接使用备用字体。

  在 Next.js 的 `next/font` 组件中的默认值是 `swap`。
  
- **preload**

  用于指定是否预加载该字体，默认值为 `true`（启用预加载）。
  
- **fallback**

  当字体加载失败时使用的备用字体列表；该值为一个字符串数组：`fallback: ['system-ui', 'arial']`，没有默认值。
  
- **adjustFontFallback**

  对于 `next/font/google` 一个布尔值，用于设置是否启用自动备用字体以减少累积布局偏移（Cumulative Layout Shift），默认值为 `true`（启用）。

  对于 `next/font/local` 一个字符串或布尔值 `false`，用于设置是否启用自动备用字体以减少累积布局偏移。可选值为 `'Arial'`、`'Times New Roman'` 或 `false`，默认值为 `'Arial'`。
  
  示例：
  - `adjustFontFallback: false`：适用于 `next/font/google`（禁用自动备用字体优化）
  - `adjustFontFallback: 'Times New Roman'`：适用于 next/font/local（手动指定备用字体为 Times New Roman）

- variable

  一个字符串值，用于定义当通过 CSS 变量方式应用样式时所使用的 CSS 变量名称。在 Web CSS 中开发者自定义属性标记设定值，然后使用 CSS 的 `var()` 函数来获取值，比如：`.p {--primary-color: #f4500; background-color: var(--primary-color);}`，这里就不详细介绍这块的内容，如果不熟悉可以阅读 MDN 的《[使用 CSS 自定义属性（变量）](https://developer.mozilla.org/zh-CN/docs/Web/CSS/CSS_cascading_variables/Using_CSS_custom_properties)》。

  在 Next.js 中要怎么声明和使用一个 CSS 变量呢？就要借助 `variable` 属性。使用方式如下：
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/CJjLCy-bd2a2450-1cbc-4d8b-931f-c9943c203b75-3293774.png)

  在浏览器中解析后的效果如下图：
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/y4KbEy-35fb3e98-3229-4d9c-8ee8-831df8c5c9ba-3293774.png)

  上面只是声明了，如果我们要在 CSS 中使用还需要借助 `var()` 函数：
  ```css
  .title {
    font-family: var(--font-geist-sans);
    font-weight: 300;
    font-style: italic;
  }
  ```
  效果如下：
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/X8ZmJt-9d6c22b4-4759-4a81-b707-e7880f896e11-3293774.png)

- **declarations**

  进一步自定义 `@font-face` 内容的生成。
  
  ![QQ_1739870131503](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/TMjFZW-cafbe2df-f814-48b1-ad26-a75b681a052e-3293774.png)

  除了这些常用的属性外，还有不少我们开发中不是很常见的属性，比如 `font-variant-ligatures`、`font-variant-caps`、`font-variant-east-asian`、`font-variant-alternates`、`font-variant-numeric` 和 `font-variant-position` 等等；而 `declarations` 就是给你用来在 `next/font/local` 时自定义 `@font-face` 的生成，使用示例如下：
  ```jsx
  import localFont from 'next/font/local';
  
  const manrope = localFont({
      src: "./fonts/Manrope-VariableFont_wght.ttf",
      display: 'swap',
      variable: '--manrope',
      fallback: ['ui-sans-serif', 'system-ui'],
      adjustFontFallback: false,
      declarations: [
          { prop: 'font-optical-sizing', value: 'auto' },
          { prop: 'ascent-override', value: '90%' },
          { prop: '-webkit-font-smoothing', value: 'antialiased' },
          { prop: '-moz-osx-font-smoothing', value: 'grayscale' }
      ]
  });
  ```
## 使用样式
在 Next.js 中，可以通过 `className`、`style`、`CSS 变量` 三种方式应用字体的样式。

CSS 变量在前面已经有过演示了，下面就演示下 `className` 和 `style` 的使用。直接在 `className` 上使用已加载字体的类名和样式，如下：
```jsx
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  display: "swap",
  subsets: ["latin"],
  weight: ['100', '200', '400', '500', '600', '700', '800'],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// ...

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistMono.className} ${geistSans.className} ${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
```
## 最佳实践
### 多种字体的使用
使用多种字体有两种方法，一种是导出字体，通过 className 来使用；一种是通过 CSS 变量的形式。
- 导出字体
  
  在 `next/font/google` 中导出字体，然后添加配置：
  ```jsx
  import { Geist, Geist_Mono } from "next/font/google";
  
  const geistSans = Geist({
    variable: "--font-geist-sans",
    display: "swap",
    subsets: ["latin"],
    weight: ['100', '200', '400', '500', '600', '700', '800'],
  });
  
  const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
  });
  ```
  在需要的时候导入并使用：
  ```jsx
  export default function RootLayout({
    children,
  }: Readonly<{
    children: React.ReactNode;
  }>) {
    return (
      <html lang="en">
        <body
          className={`${geistMono.className} ${geistSans.className} antialiased`}
        >
          {children}
        </body>
      </html>
    );
  }
  ```
- 通过 CSS 变量使用
  
  前面已经使用过这种方式，这里也简单看一下：
  ```jsx
  import { Geist, Geist_Mono } from 'next/font/google'
  
  const geistSans = Geist({
    variable: "--font-geist-sans",
    display: "swap",
    subsets: ["latin"],
    weight: ['100', '200', '400', '500', '600', '700', '800'],
  });
  
  const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
  });
  
  export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }) {
    return (
      <html lang="en" className={`${geistSans.variable} ${geistMono.variable}`}>
        <body>
          <h1>My App</h1>
          <div>{children}</div>
        </body>
      </html>
    )
  }
  ```
  在浏览器渲染后的效果如下：
  
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/nPmuZp-681810ee-2210-49f9-9c29-6e287afe0351-3293775.png)
  
  从图中可以看到，声明了 `--font-geist-mono` 和 `--font-geist-sans` 两个 CSS 变量，当需要使用的地方时就可以通过 `var()` 的形式引用：
  ```css
  html {
      font-family: var(--font-geist-sans);
  }
  .title{
      font-family: var(--font-geist-mono);
  }
  ```

### 跟 tailwindcss 搭配一起使用
其实这块本质跟 CSS 变量的引用是一样的，只不过结合 tailwindcss 时，需要在 `tailwind.config.ts` 中添加配置！具体的配置如下：
- 通过 `variable` 声明 CSS 变量
  ```jsx
  import { Geist, Geist_Mono } from 'next/font/google'
  import './globals.css'  // 引入 css
  
  const geistSans = Geist({
    variable: "--font-geist-sans",
    display: "swap",
    subsets: ["latin"],
    weight: ['100', '200', '400', '500', '600', '700', '800'],
  });
  
  const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
  });
  
  export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }) {
    return (
      <html lang="en" className={`${geistSans.variable} ${geistMono.variable}`}>
        <body>
          <h1>My App</h1>
          <div>{children}</div>
        </body>
      </html>
    )
  }
  ```
- globals CSS 文件
  ```css
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  ```
- 将 CSS 变量添加到 `tailwind.config.ts` 配置中
  ```ts
  fontFamily: {
    sans: ["var(--font-geist-sans)"],
    mono: ["var(--font-geist-mono)"],
  },
  ```
  具体配置如下图：
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/ep2jn6-467c0de7-66d3-401a-90a7-da0fb585d5d4-3293775.png)
- 在页面中使用
  > 在 [tailwindcss 中自定义字体](https://tailwindcss.com/docs/font-family)，需要通过 `font-` 为前缀来使用！
  ```jsx
  <li className='font-sans'>Save and see your changes instantly.</li>
  ```
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/1EXafX-46c4466d-f4f7-467b-9d47-61e463ad4059-3293775.png)
  
  浏览器的效果：
  ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/dAYvNV-fe5a3e07-4e8d-42e2-b042-48f3ea4f50cd-3293775.png)

### iconfont 字体使用
这里就引用[阿里妈妈方圆体](https://www.iconfont.cn/fonts/detail?spm=a313x.fonts_index.i1.d9df05512.dd413a81lyTTMF&cnid=pOvFIr086ADR)字体，如下图：

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/aqx1Ti-424ad35a-3cde-4b5e-88bd-e1971e3e7216-3293775.png)
> 因为篇幅的问题，就不一步步演示字体的下载，下载字体需要登录！

下载字体后，将字体放入项目中并在 globals.css 中引入，然后创建 `alimama/page.tsx` 文件。 如下图：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/W6y8jH-e9cabcf1-63f0-4ffe-bf4c-e788d55b806f-3293775.png)

引入字体的 CSS 代码：
```css
@font-face {
    font-family: "alimama";
    font-weight: 400;
    src: url("../fonts/AlimamaFangYuanTiVF-Thin.woff2") format("woff2"),
        url("../fonts/AlimamaFangYuanTiVF-Thin.woff") format("woff");
}
```
接着就在在 `tailwind.config.ts` 文件中配置阿里妈妈字体：
```ts
alimama: ['alimama'],
```

![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/I0pRB7-71b121bb-cb40-4a4d-b891-3480e5076599-3293775.png)

下面是在 `app/alimama/page.tsx` 中的使用：
```jsx
const page = () => {
    return (
        <div className="font-alimama text-3xl">这是阿里妈妈字体</div>
    )
}

export default page
```
效果如下：
![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/8Ofzay-fdd7aef7-6abf-4945-ae99-0cc2c99f85f8-3293775.png)

关于 iconfont 图标的使用，这里推荐一篇掘金社区的文章 [nextjs封装svgIcon，使用iconfont](https://juejin.cn/post/7326414660487512116)!
> 上面引入阿里妈妈字体也可以通过 `localFont` 函数加载，因为篇幅太长，这里就不详细说明了，我在演示项目中完善了这一块的内容，可以通过 [https://github.com/clin211/next-awesome/commit/813b6c809ccd203051d0f6577f856acbb2c92218](https://github.com/clin211/next-awesome/commit/813b6c809ccd203051d0f6577f856acbb2c92218) 查看！
> ![](https://static-hub.oss-cn-chengdu.aliyuncs.com/notes-assets/cD1zAW-1acdd8c2-6720-4e6f-a241-0d0c288b6e0a-3293775.png)


## 总结
本文对比介绍了传统字体和可变字体的区别，接着引出 Next.js 中 Font 组件；Next.js 的 Font 组件通过**自动预加载**、**消除布局偏移**和**可变字体支持**，显著提升网页性能与开发效率，同时赋予设计更大灵活性。其*提供 Google 字体与本地托管双模式加载*方案，结合 subsets 子集、axes 可变轴等配置等，构建出高效字体加载体系。实践中推荐采用 CSS 变量全局管理，优先使用可变字体减少文件体积，配合子集化加载进一步优化性能。通过变量绑定无缝集成 Tailwind 等主流框架，并支持混合字体方案灵活应对多样化设计场景。开发中需注意本地字体路径、非可变字体的必填参数（如 weight/style），以及生产环境下预加载等等细节。


**「文章推荐」**
- [What is My Font?](https://medium.com/@paowens59/what-is-my-font-f775c72fe67a)：https://medium.com/@paowens59/what-is-my-font-f775c72fe67a
- [Introducing OpenType Variable Fonts](https://medium.com/variable-fonts/https-medium-com-tiro-introducing-opentype-variable-fonts-12ba6cd2369)：https://medium.com/variable-fonts/https-medium-com-tiro-introducing-opentype-variable-fonts-12ba6cd2369
- [如何在前端项目中使用出色的开源字体](https://juejin.cn/post/7381784676720050215?searchId=20250217154234D97334B97099A4765691)：https://juejin.cn/post/7381784676720050215?searchId=20250217154234D97334B97099A4765691
- [字体文件 OTF 压缩](https://juejin.cn/post/7397592619809308691?searchId=20250217154234D97334B97099A4765691)：https://juejin.cn/post/7397592619809308691?searchId=20250217154234D97334B97099A4765691
- [nextjs封装svgIcon，使用iconfont](https://juejin.cn/post/7326414660487512116)：https://juejin.cn/post/7326414660487512116

**「参考资料」**
- [@font-face](https://developer.mozilla.org/zh-CN/docs/Web/CSS/@font-face)：https://developer.mozilla.org/zh-CN/docs/Web/CSS/@font-face
- [subset](https://fonts.google.com/knowledge/glossary/subsetting)：https://fonts.google.com/knowledge/glossary/subsetting
- [MDN font-weight](https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-weight)：https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-weight
- [MDN font-style](https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-style)：https://developer.mozilla.org/zh-CN/docs/Web/CSS/font-style
- [tailwindcss 中自定义字体](https://tailwindcss.com/docs/font-family)：https://tailwindcss.com/docs/font-family
- [字字珠玑](https://github.com/fan2/FontType/blob/master/README.md)：https://github.com/fan2/FontType/blob/master/README.md