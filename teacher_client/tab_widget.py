from PyQt5.QtWidgets import *

from teacher_client.home import HomeWidget


class TabWidget(QTabWidget):
    def __init__(self):
        self.addTab(HomeWidget(self))
        self.tabBar().setTabButton(0, QTabBar.RightSide, None)

    def addStudentTab(self):
        pass

    def addTeacherTab(self):
        pass
