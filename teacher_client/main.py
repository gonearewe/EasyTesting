import sys

import qtmodern.styles
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import *

from teacher_client import status
from teacher_client.home import HomeWidget
from teacher_client.login import LoginDialog
from teacher_client.tab_widget import TabWidget


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        status._STATUS_BAR = self.statusBar()
        tab_widget = TabWidget(self)  # 在主窗口上创建一个QTabWidget控件
        self.setCentralWidget(tab_widget)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    app.setFont(QFont("YaHei", 25))
    window = MainWindow()
    if LoginDialog().exec_() == QDialog.Accepted:
        window.showMaximized()
    sys.exit(app.exec_())
