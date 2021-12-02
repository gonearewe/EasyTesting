from typing import Dict, List

from common.dialogs import AlertDialog
from teacher_client import api
from teacher_client.table_widget import TableWidget


class TeacherTableWidget(TableWidget):
    tab_name = "教师管理"

    def __init__(self, tab_widget):
        self.tab_widget = tab_widget
        self.PAGE_SIZE = 20
        self.queries = {}
        super().__init__(parent=tab_widget,
                         queries=[("工号", "teacher_id"), ("姓名", "name")],
                         columns=["工号", "姓名", "是管理员", ("修改", "删除")],
                         row_num=self.PAGE_SIZE,
                         is_readonly=False)
        self.onSearch({})  # initial data

    def onExport(self, filepath: str):
        teachers = api.getTeachers(**self.queries)
        if teachers is None:
            AlertDialog("无法获取数据").exec_()
        else:
            self.exportData(filepath,
                            ((teacher["teacher_id"], teacher["name"], teacher["is_admin"]) for teacher in teachers))

    def onDelete(self, ids: List[int]):
        api.delTeachers(ids=ids)

    def onSearch(self, queries: Dict[str, str]):
        self.queries = queries
        self.updateTable(page_index=1)

    def updateTable(self, page_index):
        def handle_op(id: int, op: str):
            if op == "删除":
                self.onDelete([id])

        num = api.getTeachersNum(**self.queries)
        teachers = api.getTeachers(**self.queries, page_size=self.PAGE_SIZE, page_index=page_index)
        if teachers is None or num is None:
            AlertDialog("无法获取数据").exec_()
        else:
            li = [None] * len(teachers)
            for i, teacher in enumerate(teachers):
                li[i] = [teacher[k] for k in ("id", "teacher_id", "name", "is_admin")]
            self.setData(page_num=-(num // -self.PAGE_SIZE), page_index=page_index, data=li,
                         op_callback=lambda id, op: handle_op(id,op))

    def onTurnToPage(self, page_index: int):
        self.updateTable(page_index)
