import sys

import qtmodern.styles
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import *

from teacher_client.home import HomeWidget
from teacher_client.login import LoginDialog


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        tab_widget = QTabWidget()  # 在主窗口上创建一个QTabWidget控件
        self.setCentralWidget(tab_widget)
        tab_widget.addTab(HomeWidget(tab_widget), "主页")
        tab_widget.setTabsClosable(True)
        # to make tabs closeable, we have to connect signal and slot manually besides `setTabsClosable(True)`
        tab_widget.tabCloseRequested.connect(lambda i: tab_widget.removeTab(i))
        # by removing the close button, we can make all tabs closeable except the very first one
        tab_widget.tabBar().setTabButton(0, QTabBar.RightSide, None)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    app.setFont(QFont("YaHei", 25))
    window = MainWindow()
    if LoginDialog().exec_() == QDialog.Accepted:
        window.showMaximized()
    sys.exit(app.exec_())
