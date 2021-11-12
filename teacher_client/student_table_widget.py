import sys
from typing import List

import qtmodern.styles
import xlsxwriter
from PyQt5.QtGui import QFont
from PyQt5.QtWidgets import QApplication, QMainWindow, QDesktopWidget

from common.alert_dialog import AlertDialog
from teacher_client import api
from teacher_client.table_widget import TableWidget


class StudentTableWidget(TableWidget):
    def __init__(self, parent):
        self.PAGE_SIZE = 20
        self.queries = {}
        super().__init__(parent, ["学号", "姓名", "班号", ("修改", "删除")], self.PAGE_SIZE, True)

    def onExport(self, filepath: str):
        students = api.getStudents(**self.queries)
        if students is None:
            AlertDialog("无法获取数据").exec_()
        else:
            self.exportTemplate(filepath)
            try:
                workbook = xlsxwriter.Workbook(filepath)
                worksheet = workbook.get_worksheet_by_name("Sheet1")
                for i, student in enumerate(students):
                    for j, column in enumerate(("student_id", "name", "class_id")):
                        worksheet.write(i + 1, j, student[column])
                workbook.close()
            except Exception as e:
                AlertDialog("无法导出文件", detail=str(e))

    def onSearch(self, queries: List[str]):
        self.queries = {}
        for i, field in ("student_id", "name", "class_id"):
            query = queries[i].strip()
            if query != "":
                self.queries[field] = query
        self.updateTable(page_index=1)

    def updateTable(self, page_index):
        num = api.getStudentsNum(**self.queries)
        students = api.getStudents(**self.queries, page_size=self.PAGE_SIZE, page_index=page_index)
        if students is None or num is None:
            AlertDialog("无法获取数据").exec_()
        else:
            li = [None] * len(students)
            for i, student in enumerate(students):
                li[i] = [student[k] for k in ("id", "student_id", "name", "class_id")]
            self.setData(page_num=-(num["num"] // -self.PAGE_SIZE), page_index=page_index, data=li,
                         op_callback=lambda id, op: print(f"{id}, {op}"))

            # self.setData(1, 1,
            #              [[3, 2018300000, "小明", "08060000"],
            #               [4, 2018300001, "小红", "08060000"]],
            #              lambda id, op: print(f"{id}, {op}"))

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
