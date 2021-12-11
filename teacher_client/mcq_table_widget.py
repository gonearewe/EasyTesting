from typing import Dict, List
import pylightxl as xl
from common.dialogs import AlertDialog
from teacher_client import api
from teacher_client.form_widget import *
from teacher_client.table_widget import TableWidget


class McqTableWidget(TableWidget):
    tab_name = "单选题管理"

    def __init__(self, tab_widget):
        self.tab_widget = tab_widget

        super().__init__(parent=tab_widget,
                         queries=[("出题人", "publisher_teacher_id")],
                         columns=["id", "publisher_teacher_id", "stem", "choice_1", "choice_2", "choice_3",
                                  "choice_4", "right_answer"],
                         columns_text=["出题人", "题干", "选项 1", "选项 2", "选项 3", "选项 4", "正确答案", ("修改", "删除")],
                         is_readonly=False)
        self.updateTable(page_index=1)  # initial data

    def doExport(self, **kwargs):
        return api.getMcqs(**kwargs)

    def onInsert(self):
        form_rows = []
        for i, (key, key_text) in enumerate((("stem", "题干"), ("choice_1", "选项 1"), ("choice_2", "选项 2"),
                                             ("choice_3", "选项 3"), ("choice_4", "选项 4"))):
            form_rows.append(FormRow(key, key_text, "", FormValueType.RICH_TEXT))
        form_rows.append(FormRow("right_answer", "正确答案", [1, 2, 3, 4], FormValueType.COMBO_BOX))
        self.tab_widget.newTab(CreateFormWidget(self.tab_widget, form_rows, api.postMcqs), "新建单选题")

    def onImport(self, filepath: str):
        mcqs = [[]]
        try:
            worksheet = xl.readxl(filepath).ws("Sheet1")
            mcqs = [{"stem": row[0], "choice_1": row[1], "choice_2": row[2], "choice_3": row[3], "choice_4": row[4]}
                    for row in list(worksheet.rows)[1:]]
        except Exception as e:
            print(e)
            status.failure("无法导入数据")
            AlertDialog("无法导入数据", detail=str(e))
            return
        successful = api.postMcqs(mcqs)
        if not successful:
            status.failure("数据文件读取成功，但是网络出错或服务器拒绝了请求")
            AlertDialog("无法导入数据", detail="网络出错或服务器拒绝了请求")
        else:
            status.success("数据导入成功")

    def doModify(self, data: List):
        form_rows = []
        for i, (key, key_text) in enumerate((("id", "ID"), ("stem", "题干"), ("choice_1", "选项 1"),
                                             ("choice_2", "选项 2"), ("choice_3", "选项 3"), ("choice_4", "选项 4"))):
            form_rows.append(FormRow(key, key_text, data[i], FormValueType.RICH_TEXT))
        form_rows.append(FormRow("right_answer", "正确答案", [1, 2, 3, 4], FormValueType.COMBO_BOX))
        self.tab_widget.newTab(ModifyFormWidget(self.tab_widget, form_rows, api.putMcqs), f"修改单选题 {data[0]}")

    def doDelete(self, ids: List[int]) -> bool:
        return api.delMcqs(ids=",".join(map(str, ids)))

    def onGetDataNum(self, queries: Dict[str, str]):
        return api.getMcqsNum(**queries)

    def onGetData(self, queries: Dict[str, str], page_size: int, page_index: int):
        return api.getMcqs(**queries, page_size=page_size, page_index=page_index)
