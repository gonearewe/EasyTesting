接下来介绍教师客户端的设计。

教师客户端的需求总的来说都能用图与表来实现。那么最方便的客户端方案莫过于基于 [vue-element-admin](https://panjiachen.github.io/vue-element-admin-site/zh/) 二次开发。
它提炼了典型的业务模型，提供了丰富的功能组件，使用了稳健的技术栈，适用于简单 Vue 应用的快速开发。
不过，vue-element-admin 是大而全的方案，实际开发中初始模板用的是 [vue-admin-template](https://github.com/PanJiaChen/vue-admin-template/blob/master/README-zh.md)。这是一个极简的 vue admin 管理后台。它只包含了 Element UI & axios & iconfont & permission control & lint，
这些搭建后台必要的东西。在需要的时候，我们再从 vue-element-admin 中复制需要的模块。

源码文件夹结构如下：

```
.
├── public              打包需要的静态文件
│   ├── favicon.ico       网页图标
│   └── index.html        单页 html
├── src                 真正的源代码
│   ├── App.vue         
│   ├── main.js
│   ├── permission.js     权限管理
│   ├── settings.js
│   ├── api               单纯的 JS API 请求库
│   │   ├── exam.js
│   │   ├── question.js
│   │   ├── student.js
│   │   ├── teacher.js
│   │   └── user.js
│   ├── assets            应用中引用的静态文件
│   │   ├── 404_images
│   │   │   ├── 404.png
│   │   │   └── 404_cloud.png
│   │   ├── dashboard
│   │   │   ├── exam.png
│   │   │   ├── student.png
│   │   │   └── teacher.png
│   │   └── logo.png
│   ├── components        通用组件
│   │   ├── Breadcrumb
│   │   │   └── index.vue
│   │   ├── GithubCorner
│   │   │   └── index.vue
│   │   ├── Hamburger
│   │   │   └── index.vue
│   │   ├── MarkdownEditor
│   │   │   ├── default-options.js
│   │   │   └── index.vue
│   │   ├── Pagination
│   │   │   └── index.vue
│   │   └── SvgIcon
│   │       └── index.vue
│   ├── directive         通用的 Vue 指令
│   │   ├── clipboard
│   │   │   ├── clipboard.js
│   │   │   └── index.js
│   │   ├── el-drag-dialog
│   │   │   ├── drag.js
│   │   │   └── index.js
│   │   ├── el-table
│   │   │   ├── adaptive.js
│   │   │   └── index.js
│   │   ├── permission
│   │   │   ├── index.js
│   │   │   └── permission.js
│   │   ├── sticky.js
│   │   └── waves
│   │       ├── index.js
│   │       ├── waves.css
│   │       └── waves.js
│   ├── icons             应用中引用的图标文件
│   │   ├── avatar.gif
│   │   ├── index.js
│   │   ├── svg
│   │   │   ├── analysis.svg
│   │   │   ├── bfq.svg
│   │   │   ├── cq.svg
│   │   │   ├── crq.svg
│   │   │   ├── dashboard.svg
│   │   │   ├── exam.svg
│   │   │   ├── example.svg
│   │   │   ├── eye-open.svg
│   │   │   ├── eye.svg
│   │   │   ├── form.svg
│   │   │   ├── link.svg
│   │   │   ├── maq.svg
│   │   │   ├── mcq.svg
│   │   │   ├── nested.svg
│   │   │   ├── password.svg
│   │   │   ├── profile.svg
│   │   │   ├── question.svg
│   │   │   ├── server.svg
│   │   │   ├── student.svg
│   │   │   ├── table.svg
│   │   │   ├── teacher.svg
│   │   │   ├── tfq.svg
│   │   │   ├── tree.svg
│   │   │   └── user.svg
│   │   └── svgo.yml
│   ├── layout            界面的整体布局
│   │   ├── components
│   │   │   ├── AppMain.vue 界面主体
│   │   │   ├── Navbar.vue  导航栏
│   │   │   ├── Sidebar     侧边栏
│   │   │   │   ├── FixiOSBug.js
│   │   │   │   ├── Item.vue
│   │   │   │   ├── Link.vue
│   │   │   │   ├── Logo.vue
│   │   │   │   ├── SidebarItem.vue
│   │   │   │   └── index.vue
│   │   │   └── index.js
│   │   ├── index.vue
│   │   └── mixin
│   │       └── ResizeHandler.js
│   ├── router            页面路由
│   │   └── index.js
│   ├── store             Vuex
│   │   ├── getters.js
│   │   ├── index.js
│   │   └── modules
│   │       ├── app.js
│   │       ├── permission.js
│   │       ├── settings.js
│   │       └── user.js
│   ├── styles            通用的样式文件
│   │   ├── btn.scss
│   │   ├── element-ui.scss
│   │   ├── element-variables.scss
│   │   ├── index.scss
│   │   ├── mixin.scss
│   │   ├── sidebar.scss
│   │   ├── transition.scss
│   │   └── variables.scss
│   ├── utils             辅助工具库
│   │   ├── Export2Excel.js
│   │   ├── ImportFromExcel.js
│   │   ├── cookie.js
│   │   ├── get-page-title.js
│   │   ├── index.js
│   │   ├── request.js
│   │   ├── scroll-to.js
│   │   ├── time.js
│   │   └── validate.js
│   └── views             各个页面的界面主体
│       ├── 404.vue
│       ├── dashboard       首页
│       │   └── index.vue
│       ├── exam            考试管理
│       │   ├── chart.vue
│       │   ├── detail.vue
│       │   ├── index.vue
│       │   └── resize.js
│       ├── login           登录界面
│       │   └── index.vue
│       ├── question        试题管理
│       │   ├── bfq.vue
│       │   ├── common.js
│       │   ├── cq.vue
│       │   ├── crq.vue
│       │   ├── maq.vue
│       │   ├── mcq.vue
│       │   └── tfq.vue
│       ├── student         学生管理
│       │   └── index.vue
│       └── teacher         教师管理
│           └── index.vue
├── LICENSE
├── babel.config.js
├── jest.config.js
├── jsconfig.json
├── package.json        项目依赖
├── postcss.config.js
└── vue.config.js
```

主要的开发工作在 views 目录下进行，这里的 Vue 文件是各个页面（*考试管理*、*试题管理*等）的界面主体（AppMain）。
Vue 文件的内容通常分为三个部分：类 html 的模板（template）、Javascript 编写的逻辑、样式。
因为各个页面的需求都很相似，所以各个 Vue 文件也很相似。

大多数页面的主体都是一个 [El-Table](https://element.eleme.io/#/zh-CN/component/table)。
我们在 El-Table 的左侧加一列多选框方便批量删除；第二列通常是数据的 id；表格的接下来几列展示对应数据的主要信息；
最后一列盛放*编辑*与*删除*按钮。
El-Table 的设计满含细节：启用斑马纹（stripe）以方便数行数，还禁止选中行以避免误导用户（批量删除要求选中 checkbox 而非选中行），
列宽度根据实际情况设置，单元格通过 [El-Tooltip](https://element.eleme.io/#/zh-CN/component/tooltip) 提供补充信息等。

基本上所有页面都有多个对话框（[El-Dialog](https://element.eleme.io/#/zh-CN/component/dialog)）。
每个 El-Dialog 都通过一个单独的布尔变量控制显示与否。El-Dialog 通常包含一个表单（或表格）和*取消*、*确定*按钮。

在这些页面被创建时（created() 方法），通常会通过 getList() 调用对应的 api 获取各自的数据，
数据随后被绑定为 El-Table 的 model。在*分页器*换页时与操作数据时也会进行这一流程。
在修改某条记录时，则会把要修改的数据拷贝到一个变量 temp 上，再弹出 El-Dialog 要求用户确认，
temp 将被绑定为 El-Dialog 的 model；通过拷贝变量 temp，可以避免 El-Dialog 表单中的修改影响到 El-Table 中的只读数据。

