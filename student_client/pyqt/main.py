import sys

import qtmodern.styles
from PyQt5.QtCore import Qt, QUrl
from PyQt5.QtGui import QFont
from PyQt5.QtWebChannel import QWebChannel
from PyQt5.QtWebEngineWidgets import QWebEngineView
from PyQt5.QtWidgets import *

import local_server
from student_client.pyqt.code_runner import CodeRunner
from student_client.pyqt.question import *
from student_client.pyqt.question_widget import QuestionWidget


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        tab_widget = QTabWidget()  # 在主窗口上创建一个QTabWidget控件
        self.setCentralWidget(tab_widget)
        tab_widget.addTab(QuestionWidget([
            MultipleChoiceQuestion(1, "下面说法正确的是", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
            MultipleChoiceQuestion(2, "下面说法错误的是", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
        ]), "单项选择题")
        tab_widget.addTab(QuestionWidget([
            MultipleAnswerQuestion(1, "下面说法正确的有", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
            MultipleAnswerQuestion(2, "下面说法错误的有", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
        ]), "多项选择题")
        tab_widget.addTab(QuestionWidget([
            BlankFillingQuestion(1, "这是一个__"),
            BlankFillingQuestion(2,
                                 """
  __是一个**核心**：
  ```py
      def setup(self):
          pass
  ```
  """)
        ]), "填空题")
        tab_widget.addTab(QuestionWidget([
            CodeReadingQuestion(1, "_ _ _ _ _", 5)
        ]), "代码阅读题")


class MenuWidget(QWidget):
    def __init__(self):
        super(MenuWidget, self).__init__()
        layout = QHBoxLayout()
        self.setLayout(layout)
        layout.setAlignment(Qt.AlignRight)
        layout.addWidget(QLCDNumber())
        layout.addWidget(QPushButton("Save"))
        layout.addWidget(QPushButton("Submit"))


if __name__ == "__main__":
    local_server.start()

    app = QApplication(sys.argv)
    app.setFont(QFont("YaHei", 25))
    qtmodern.styles.light(app)

    browser = QWebEngineView()
    browser.load(QUrl("http://localhost:2998"))
    browser.show()
    channel = QWebChannel()
    channel.registerObject('code_runner', CodeRunner())
    browser.page().setWebChannel(channel)

    # window = QMainWindow()
    # window.showMaximized()
    sys.exit(app.exec_())
