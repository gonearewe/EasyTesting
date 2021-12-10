from typing import Dict, List
import pylightxl as xl
from common.dialogs import AlertDialog
from teacher_client import api
from teacher_client.form_widget import *
from teacher_client.table_widget import TableWidget


class StudentTableWidget(TableWidget):
    tab_name = "学生管理"

    def __init__(self, tab_widget):
        self.tab_widget = tab_widget

        super().__init__(parent=tab_widget,
                         queries=[("学号", "student_id"), ("姓名", "name"), ("班号", "class_id")],
                         columns=["id", "student_id", "name", "class_id"],
                         columns_text=["学号", "姓名", "班号", ("修改", "删除")],
                         is_readonly=False)
        self.updateTable(page_index=1)  # initial data

    def doExport(self, **kwargs):
        return api.getStudents(**kwargs)

    def onInsert(self):
        form_rows = []
        for i, (key, key_text) in enumerate((("student_id", "学号"), ("name", "姓名"), ("class_id", "班号"))):
            form_rows.append(FormRow(key, key_text, "", FormValueType.SINGLE_LINE))
        self.tab_widget.newTab(CreateFormWidget(self.tab_widget, form_rows, api.postStudents), "新建学生")

    def onImport(self, filepath: str):
        students = [[]]
        try:
            worksheet = xl.readxl(filepath).ws("Sheet1")
            students = [{"student_id": row[0], "name": row[1], "class_id": row[2]} for row in list(worksheet.rows)[1:]]
        except Exception as e:
            print(e)
            status.failure("无法导入数据")
            AlertDialog("无法导入数据", detail=str(e))
            return
        successful = api.postStudents(students)
        if not successful:
            status.failure("数据文件读取成功，但是网络出错或服务器拒绝了请求")
            AlertDialog("无法导入数据", detail="网络出错或服务器拒绝了请求")
        else:
            status.success("数据导入成功")

    def doModify(self, data: List):
        form_rows = []
        for i, (key, key_text) in enumerate((("id", "ID"), ("student_id", "学号"), ("name", "姓名"), ("class_id", "班号"))):
            form_rows.append(FormRow(key, key_text, data[i], FormValueType.SINGLE_LINE))
        self.tab_widget.newTab(ModifyFormWidget(self.tab_widget, form_rows, api.putStudents), f"修改学生 {data[1]}")

    def doDelete(self, ids: List[int]) -> bool:
        return api.delStudents(ids=",".join(map(str, ids)))

    def onGetDataNum(self, queries: Dict[str, str]):
        return api.getStudentsNum(**queries)

    def onGetData(self, queries: Dict[str, str], page_size: int, page_index: int):
        return api.getStudents(**queries, page_size=page_size, page_index=page_index)
