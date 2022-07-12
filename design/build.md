这里介绍如何从源代码构建系统。
源码的 build 文件夹中提供了 Windows 环境下用于自动构建的 bat 脚本，
进入目录双击运行即可。如果你需要在其他操作系统中手动完成构建，
或者想知道脚本的工作流程，请接着往下看。

## 服务端与教师客户端

教师客户端的单页 Web 应用就作为静态文件部署在服务端软件的同级目录下。
所以，直接访问服务端的 HTTP 根路径就会进入教师客户端，其他的大多数接口则是 RESTful 的。

教师客户端的 package.json 与 package-lock.json 记录着 Vue 应用的依赖。
通过 `npm run build` 即可自动下载依赖，完成构建。
构建得到的文件将位于 teacher_client/dist 下，它们是一个名为 static 的文件夹、一个 index.html 与一个 favicon.ico。

服务端软件的构建同样简单。
go.mod 与 go.sum 记录着 Go 软件的依赖。
只要在 server 目录下执行 `go build` 即可自动下载依赖，完成构建。
构建同时得到 EasyTesting（用于 Linux）和 EasyTesting.exe（用于 Windows）两个二进制文件。

新建一个文件夹存放它们的构建成果。注意，还需要拷贝一份服务端软件的配置文件 server-config.yaml。
为了方便用户部署服务端时顺便处理好数据库，我们还可拷贝一份 server 文件夹下的 sql 文件夹。
最后得到的文件夹结构将类似于：

```
├── EasyTesting          服务端主程序
├── server-config.yaml   服务端主程序的配置文件
├── favicon.ico          软件图标
├── index.html           教师端主页面
├── README.md            使用须知
├── LICENSE              软件开源许可证
├── sql                  MySQL 脚本
│   ├── setup.sql        创建数据库的脚本
│   └── test.sql         插入测试数据的脚本
└── static               教师端主页面的静态文件
    ├── css              样式文件
    │   ├── app.4c2aef82.css
    │   ├── chunk-019156c7.a8088982.css
    │   ├── chunk-07c1fcd4.04dea0eb.css
    │   └── ...
    ├── fonts            字体文件
    │   ├── element-icons.535877f5.woff
    │   └── element-icons.732389de.ttf
    ├── img              图片
    │   ├── 404.a57b6f31.png
    │   ├── 404_cloud.0f4bc32b.png
    │   ├── avatar.ecba1844.gif
    │   └── ...
    └── js  javascript  脚本文件
        ├── app.82de0e6e.js
        ├── app.9d60ec40.js
        ├── chunk-019156c7.b0ca593f.js
        └── ...
```

> README.md 和 LICENSE 都是从项目根目录拷贝来的，不影响软件的使用。

## 学生客户端

相对而言，学生客户端的构建更加复杂。

首先我们需要从官网下载一份 Python 解释器，它将作为学生客户端的编程题运行环境。
默认构建使用的是 [Python v3.7.9](https://www.python.org/downloads/release/python-379/)，
Windows 环境下载 `Windows x86-64 embeddable zip file`，它已包含所需的解释器与标准库，我们并不需要包含 pip 等工具的完整版。
要把它解压得到的文件夹放到 student_client/pyqt 下。
然后要在 student_client/vue 下运行 `npm run build` 构建好 Vue 应用。
其构建成果与 teacher_client 类似。
接下来，我们需要把生成的 static 文件夹与另两个文件也放到 student_client/pyqt 下。

我们使用 PyInstaller 打包 PyQt，这样一来学生用户就无需自行安装 Python，更不需要操心依赖库。
PyInstaller 使用 student_client/pyqt/main.spec 作为配置文件，
其内容含义可参考[官方文档](https://pyinstaller.readthedocs.io/en/v4.8/spec-files.html)。
我们主要修改两处地方，第一是创建 Analysis 对象的 datas 参数，我们提供包含四对元组的一个列表，
它表示将每个元组的第一个元素代表的文件或文件夹拷贝到打包后的文件夹内，
每个元组的第二个元素表示打包后的文件夹内的新文件或文件夹名。
这里，我们要求原样拷贝 favicon.ico、index.html 与 static 文件夹，
同时将 Python 解释器文件夹拷贝并重命名为 runner。
第二个修改的地方是创建 Analysis 对象的 pathex 参数，它告知 PyInstaller 从何处打包项目所需的依赖库。
PyQt 子项目使用 [pipenv](https://pipenv.pypa.io/en/latest/) 管理依赖。
pipenv 根据 student_client/pyqt/Pipfile 的要求将依赖下载到本地，依赖的本地路径即为 `pipenv --venv` 命令的输出。

```py
# main.spec
# ...
a = Analysis(['main.py'],
             # replace this with where your dependencies lay, i.e. your output for `pipenv --venv`
             pathex=['C:\\\\Users\\\\John Mactavish\\\\.virtualenvs\\\\pyqt-Jt_vAwyy'],
             # ...
             # data files copy, see documentation for details
             datas=[('favicon.ico','.'),('index.html','.'),('static','static'),('python-3.7.9-embed-amd64','runner')],
             # ...
             noarchive=False)
# ...
```

这两项修改做好后，在 student_client/pyqt 目录下运行 `pipenv run pyinstaller main.spec` 即可在 student_client/pyqt/dist 中得到
打包后的文件夹 main。