import sys

import qtmodern.styles
from PyQt5.QtCore import Qt
from PyQt5.QtWidgets import *

from student_client.maq import MultipleAnswerQuestionWidget as MAQWidget
from student_client.mcq import MultipleChoiceQuestionWidget as MCQWidget


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        tab_widget = QTabWidget()  # 在主窗口上创建一个QTabWidget控件
        self.setCentralWidget(tab_widget)
        tab_widget.addTab(MCQWidget([
            MCQWidget.Question(1, "下面说法正确的是", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
            MCQWidget.Question(2, "下面说法错误的是", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
        ]), "单项选择题")
        tab_widget.addTab(MAQWidget([
            MAQWidget.Question(1, "下面说法正确的有", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
            MAQWidget.Question(2, "下面说法错误的有", [(1, "66"), (2, "gh"), (3, "yu"), (4, "po")]),
        ]), "多项选择题")


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
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    window = MainWindow()
    window.show()
    sys.exit(app.exec_())
