import sys
from typing import Dict, List

import qtmodern.styles
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import QApplication, QMainWindow, QDesktopWidget

from common.dialogs import AlertDialog
from teacher_client import api
from teacher_client.table_widget import TableWidget


class StudentTableWidget(TableWidget):
    tab_name = "学生管理"

    def __init__(self, tab_widget):
        self.tab_widget = tab_widget
        self.PAGE_SIZE = 20
        self.queries = {}
        super().__init__(parent=tab_widget,
                         queries=[("学号", "student_id"), ("姓名", "name"), ("班号", "class_id")],
                         columns=["学号", "姓名", "班号", ("修改", "删除")],
                         row_num=self.PAGE_SIZE,
                         is_readonly=True)
        self.onSearch({})  # initial data

    def onExport(self, filepath: str):
        students = api.getStudents(**self.queries)
        if students is None:
            AlertDialog("无法获取数据").exec_()
        else:
            self.exportData(filepath,
                            ((student["student_id"], student["name"], student["class_id"]) for student in students))

    def onDelete(self, ids: List[int]):
        api.delStudents(ids=ids)

    def onSearch(self, queries: Dict[str, str]):
        self.queries = queries
        self.updateTable(page_index=1)

    def updateTable(self, page_index):
        def handle_op(id: int, op: str):
            if op == "删除":
                self.onDelete([id])

        num = api.getStudentsNum(**self.queries)
        students = api.getStudents(**self.queries, page_size=self.PAGE_SIZE, page_index=page_index)
        if students is None or num is None:
            AlertDialog("无法获取数据").exec_()
        else:
            li = [[]] * len(students)
            for i, student in enumerate(students):
                li[i] = [student[k] for k in ("id", "student_id", "name", "class_id")]
            self.setData(page_num=-(num // -self.PAGE_SIZE), page_index=page_index, data=li,
                         op_callback=lambda id, op: handle_op(id, op))

    def onTurnToPage(self, page_index: int):
        self.updateTable(page_index)


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
