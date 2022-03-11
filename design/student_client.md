学生客户端的技术选型则有点变化。为了防作弊和整合编程题调试功能，学生客户端必须为 PC 端软件，
而不能向教师客户端一样采用浏览器平台。鉴于编程题为 Python 编程，PC 端实现最好也采用 Python 的技术栈。
因而基础框架选择的是 PyQt5，本来计划整个客户端都使用 PyQt 构建，但是中途发觉 Python 语言实在是不适合描述 UI。
Qt 还有一个 DSL 叫 QML，更适合描述 UI，但是这样一来只会让设计变得更加复杂。
最终方案是只使用 PyQt 来作为 PC 端平台，然后在其上用 [Qt WebEngine](https://doc.qt.io/qt-5/qtwebengine-index.html) 搭建浏览器平台，
这样就能使用已有的 [vue-admin-template](https://github.com/PanJiaChen/vue-admin-template/blob/master/README-zh.md) 构建主体应用。
而 Python 编程题的在线调试功能不能一并交给 Vue 应用，因为我们没有找到能在浏览器运行的 Javascript 实现的 Python 解释器，
所以只好让 Python 平台程序另外启动一个本地服务端，以处理调试请求。
我们没有采用 [Qt WebChannel](https://doc.qt.io/qt-5/qtwebchannel-index.html) 的通信方案，
是因为它没有本地 Client/Server 的方案常用。
而为了避免直接向用户暴露 Python 源文件，更为了方便没有安装 Python 解释器的用户能直接使用客户端，
项目最终使用 PyInstaller 来作打包。

学生客户端的源码文件夹结构如下：

```
.
├── pyqt                Python 实现的程序平台
│   ├── Pipfile           pipenv 的依赖文件
│   ├── Pipfile.lock      自动生成的 pipenv 依赖文件
│   ├── code_runner.py    用于编程题调试的 Python Code Runner
│   ├── config.py         配置文件
│   ├── local_server.py   用于编程题调试的本地服务器
│   ├── main.py           程序入口
│   └── main.spec         PyInstaller 的配置文件
└── vue                 Vue 部分
    ├── LICENSE
    ├── babel.config.js
    ├── jest.config.js
    ├── jsconfig.json
    ├── package.json      Vue 的项目依赖
    ├── postcss.config.js
    ├── public            打包需要的静态文件
    │   ├── favicon.ico     网页图标
    │   └── index.html      单页 html
    ├── src               真正的源代码
    │   ├── App.vue
    │   ├── main.js
    │   ├── permission.js   权限管理 
    │   ├── settings.js
    │   ├── api             单纯的 JS API 请求库
    │   │   └── index.js
    │   ├── assets          应用中引用的静态文件
    │   │   ├── 404_images
    │   │   │   ├── 404.png
    │   │   │   └── 404_cloud.png
    │   │   ├── logo.png
    │   │   └── tip_zoom.png
    │   ├── components      通用组件
    │   │   ├── BackToTop
    │   │   │   └── index.vue
    │   │   └── SvgIcon
    │   │       └── index.vue
    │   ├── icons           应用中引用的图标文件
    │   │   ├── avatar.gif
    │   │   ├── index.js
    │   │   ├── svg
    │   │   │   ├── exam.svg
    │   │   │   ├── id.svg
    │   │   │   ├── server.svg
    │   │   │   └── user.svg
    │   │   └── svgo.yml
    │   ├── layout          界面的整体布局
    │   │   ├── components
    │   │   │   ├── AppMain.vue
    │   │   │   └── index.js
    │   │   ├── index.vue
    │   │   └── mixin
    │   │       └── ResizeHandler.js
    │   ├── router          页面路由
    │   │   └── index.js
    │   ├── store           Vuex
    │   │   ├── getters.js
    │   │   ├── index.js
    │   │   └── modules
    │   │       ├── settings.js
    │   │       └── user.js
    │   ├── styles          通用的样式文件
    │   │   ├── btn.scss
    │   │   ├── element-ui.scss
    │   │   ├── element-variables.scss
    │   │   ├── index.scss
    │   │   ├── mixin.scss
    │   │   ├── sidebar.scss
    │   │   ├── transition.scss
    │   │   └── variables.scss
    │   ├── utils           辅助工具库
    │   │   ├── cookie.js
    │   │   ├── get-page-title.js
    │   │   ├── index.js
    │   │   ├── random.js
    │   │   ├── request.js
    │   │   ├── scroll-to.js
    │   │   ├── time.js
    │   │   └── validate.js
    │   └── views           各个页面的界面主体
    │       ├── 404.vue
    │       ├── home          主界面
    │       │   └── index.vue
    │       └── login         登录界面
    │           └── index.vue
    └── vue.config.js
```

学生客户端的 Vue 应用相对教师端更简单，只有登录界面与答题的主界面。
主界面使用一个 [El-Tabs](https://element.eleme.io/#/zh-CN/component/tabs) 承载所有的题目，
不同的题型被放在不同的 Tab。主界面在被创建时获取题目至 questions 变量，
而学生的作答对应的 model 为 answers 变量。应用每隔几分钟就将作答保存到服务端，这包括两个 API 调用：
PUT exams/my_answers 会将作答提交至数据库，而 PUT cache 会将 answers 变量直接序列化成字符串保存到服务端的内存 Cache 中去；
后者是方便学生中途退出程序重进时，通过 GET cache 快速还原 answers 变量的状态。
设计的*检查进度*功能是通过一个按钮触发的，之所以不设计成一个常驻的实时展示进度的组件，是因为 answers 变量结构较复杂，
不适合通过 model 监听。

Python 部分主要是启动 Qt WebEngine 和本地服务端。
本地服务端选型是 [Flask](https://flask.palletsprojects.com/en/2.0.x/)，在 http://localhost:2998 监听，
处理 PUT code 请求与静态文件请求。Flask 在单独的进程运行，与 PyQt 并行执行。
在处理 PUT code 请求时，程序使用 PyInstaller 打包好的单独的 Python 解释器执行用户代码，并按题目要求处理好输入输出与异常。
值得一提的是，编程题的对错就是在此时判定的，客户端判分必然是脆弱的、不安全的，但这也是受限于实际情境中服务器不足的性能的结果，
在重构时也许会考虑作出改变。


