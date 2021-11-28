import sys

import qtmodern.styles
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import *

from teacher_client.home import HomeWidget


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        tab_widget = QTabWidget()  # 在主窗口上创建一个QTabWidget控件
        self.setCentralWidget(tab_widget)
        tab_widget.addTab(HomeWidget(self), "主页")
        tab_widget.setTabsClosable(True)
        tab_widget.tabBar().setTabButton(0, QTabBar.RightSide, None)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    app.setFont(QFont("YaHei", 25))
    window = MainWindow()
    # if LoginDialog().exec_() == QDialog.Accepted:
    window.showMaximized()
    sys.exit(app.exec_())
