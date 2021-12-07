from typing import Dict, List

from common.dialogs import AlertDialog
from teacher_client import api
from teacher_client.form_widget import *
from teacher_client.table_widget import TableWidget


class TeacherTableWidget(TableWidget):
    tab_name = "教师管理"

    def __init__(self, tab_widget):
        self.tab_widget = tab_widget

        super().__init__(parent=tab_widget,
                         queries=[("工号", "teacher_id"), ("姓名", "name")],
                         columns=["id", "teacher_id", "name", "is_admin"],
                         columns_text=["工号", "姓名", "是管理员", ("修改", "删除")],
                         is_readonly=False)
        self.updateTable(page_index=1)  # initial data

    def onExport(self, filepath: str):
        teachers = api.getTeachers(**self.queries)
        if teachers is None:
            AlertDialog("无法获取数据").exec_()
        else:
            self.exportData(filepath,
                            ((teacher["teacher_id"], teacher["name"], teacher["is_admin"]) for teacher in teachers))

    def doModify(self, data: List) -> bool:
        form_rows = []
        for i, (key,key_text) in enumerate((("id","ID"),("teacher_id", "工号"),("name", "姓名"))):
            form_rows.append(FormRow(key,key_text, data[i], FormValueType.SINGLE_LINE))
        form_rows.append(FormRow("is_admin", "是管理员", ("是", "否"), FormValueType.COMBO_BOX))
        self.tab_widget.addTab(ModifyFormWidget(self.tab_widget, form_rows, api.putTeachers),f"修改教师 {data[1]}")

    def doDelete(self, ids: List[int]) -> bool:
        return api.delTeachers(ids=ids)

    def onGetDataNum(self, queries: Dict[str, str]):
        return api.getTeachersNum(**queries)

    def onGetData(self, queries: Dict[str, str], page_size: int, page_index: int):
        return api.getTeachers(**queries, page_size=page_size, page_index=page_index)
