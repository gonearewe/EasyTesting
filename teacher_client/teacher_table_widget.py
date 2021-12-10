from typing import Dict, List
import pylightxl as xl
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

    def doExport(self, **kwargs):
        return api.getTeachers(**kwargs)

    def onInsert(self):
        form_rows = []
        for i, (key, key_text) in enumerate((("teacher_id", "工号"), ("name", "姓名"))):
            form_rows.append(FormRow(key, key_text, "", FormValueType.SINGLE_LINE))
        form_rows.append(FormRow("is_admin", "是管理员", ("是", "否"), FormValueType.COMBO_BOX))
        self.tab_widget.newTab(CreateFormWidget(self.tab_widget, form_rows, api.postTeachers), "新建教师")

    def onImport(self, filepath: str):
        teachers = [[]]
        try:
            worksheet = xl.readxl(filepath).ws("Sheet1")
            teachers = [{"teacher_id": row[0], "name": row[1], "is_admin": row[2]} for row in list(worksheet.rows)[1:]]
        except Exception as e:
            print(e)
            status.failure("无法导入数据")
            AlertDialog("无法导入数据", detail=str(e))
            return
        successful = api.postTeachers(teachers)
        if not successful:
            status.failure("数据文件读取成功，但是网络出错或服务器拒绝了请求")
            AlertDialog("无法导入数据", detail="网络出错或服务器拒绝了请求")
        else:
            status.success("数据导入成功")

    def doModify(self, data: List):
        form_rows = []
        for i, (key, key_text) in enumerate((("id", "ID"), ("teacher_id", "工号"), ("name", "姓名"))):
            form_rows.append(FormRow(key, key_text, data[i], FormValueType.SINGLE_LINE))
        form_rows.append(FormRow("is_admin", "是管理员", ("是", "否"), FormValueType.COMBO_BOX))
        self.tab_widget.newTab(ModifyFormWidget(self.tab_widget, form_rows, api.putTeachers), f"修改教师 {data[1]}")

    def doDelete(self, ids: List[int]) -> bool:
        return api.delTeachers(ids=",".join(map(str, ids)))

    def onGetDataNum(self, queries: Dict[str, str]):
        return api.getTeachersNum(**queries)

    def onGetData(self, queries: Dict[str, str], page_size: int, page_index: int):
        return api.getTeachers(**queries, page_size=page_size, page_index=page_index)
