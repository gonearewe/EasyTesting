from PyQt5.QtWidgets import *

from teacher_client.home import HomeWidget


class TabWidget(QTabWidget):
    def __init__(self,parent):
        super().__init__(parent)
        self.addTab(HomeWidget(self), "主页")
        self.setTabsClosable(True)
        # to make tabs closeable, we have to connect signal and slot manually besides `setTabsClosable(True)`
        self.tabCloseRequested.connect(lambda i: self.removeTab(i))
        # by removing the close button, we can make all tabs closeable except the very first one
        self.tabBar().setTabButton(0, QTabBar.RightSide, None)

    def newTab(self,widget:QWidget,title:str):
        i = self.addTab(widget,title)
        self.setCurrentIndex(i)

    def closeCurrentTab(self):
        self.removeTab(self.currentIndex())

    def addStudentTab(self):
        pass

    def addTeacherTab(self):
        pass
