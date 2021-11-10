import sys
from typing import List

import qtmodern.styles
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import QApplication, QMainWindow, QDesktopWidget

from teacher_client.table_widget import TableWidget


class StudentTableWidget(TableWidget):
    def __init__(self, parent):
        super().__init__(parent, ["学号", "姓名", "班号", ("修改", "删除")], 20, True)
        self.setData(1, 1,
                     [[3, 2018300000, "小明", "08060000"],
                      [4, 2018300001, "小红", "08060000"]],
                     lambda id, op: print(f"{id}, {op}"))

    def onExport(self, filepath: str):
        print(filepath)

    def onSearch(self, queries: List[str]):
        print(queries)

    def onTurnToPage(self, page_index: int):
        print(page_index)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    app.setFont(QFont("YaHei", 25))
    w = QMainWindow()
    s = StudentTableWidget(w)
    s.frameGeometry().moveCenter(QDesktopWidget().availableGeometry().center())
    w.setCentralWidget(s)
    w.showMaximized()
    sys.exit(app.exec_())
