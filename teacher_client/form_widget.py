from enum import Enum
from typing import *

from PyQt5.QtWidgets import *

from common.dialogs import ConfirmDialog
from teacher_client import status


class FormWidget(QWidget):
    def __init__(self, parent, rows: List['FormRow']):
        super().__init__(parent)
        self.setContentsMargins(300, 100, 300, 100)
        self.layout = QFormLayout(self)
        self.setLayout(self.layout)
        self.widgets = []
        self.rows = rows
        for row in rows:
            v = None
            if row.type_ == FormValueType.SINGLE_LINE:
                v = QLineEdit(str(row.default_val), self)
                v.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
            elif row.type_ == FormValueType.RICH_TEXT:
                v = QWidget()
                l = QHBoxLayout(v)
                v.setLayout(l)
                text = """支持 Markdown 格式：\n\n**加粗** *倾斜* ***倾斜加粗***\n\n- 列表一号\n- 列表二号\n- 列表三号\n\n```py\n\n""" + \
                       """# 代码区域内 Markdown 语法会被屏蔽\ndef foo(self, *args, **kwargs):\n    pass\n```\n\n""" + \
                       """Markdown 不支持换行与缩进，请用一个空行分割段落\n\n""" + \
                       """更多格式参考 https://www.markdownguide.org/basic-syntax/ """
                edit = QTextEdit()
                edit.setPlainText(row.default_val if row.default_val else text)
                l.addWidget(edit)
                l.addWidget(QLabel("预览："))
                display = QTextBrowser()
                display.setMarkdown(edit.toPlainText())
                edit.textChanged.connect(
                    lambda ignored=None, display=display, edit=edit: display.setMarkdown(edit.toPlainText()))
                l.addWidget(display)
            elif row.type_ == FormValueType.COMBO_BOX:
                v = QComboBox(self)
                v.addItems(row.default_val)
                v.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
            self.widgets.append(v)
            self.layout.addRow(row.key_text, v)

    def get_form(self) -> Dict[str, str]:
        m = {}
        for i, w in enumerate(self.widgets):
            row = self.rows[i]
            if row.type_ == FormValueType.COMBO_BOX and w.currentText() in ("是", "否"):
                m[row.key] = w.currentText() == "是"
            else:
                if type(row.default_val) == int:
                    m[row.key] = int(w.text())
                else:
                    m[row.key] = w.text()
        return m


class FormRow:
    def __init__(self, key: str, key_text: str, default_val: Union[Any, Iterable], type_: 'FormValueType'):
        self.key = key
        self.key_text = key_text
        self.default_val = default_val
        self.type_ = type_


class FormValueType(Enum):
    SINGLE_LINE = 1
    RICH_TEXT = 2
    COMBO_BOX = 3


class CreateFormWidget(FormWidget):
    def __init__(self, tab_widget, rows: List['FormRow'], on_create):
        super().__init__(parent=tab_widget, rows=rows)
        btn = QPushButton("创建")
        btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        self.layout.addWidget(btn)

        def on_pressed():
            if ConfirmDialog("确认要创建吗？").exec_() != QMessageBox.Ok:
                return
            successful = on_create([self.get_form()])
            if successful:
                status.success("创建成功")
                tab_widget.closeCurrentTab()
            else:
                status.failure("创建失败")

        btn.pressed.connect(on_pressed)


class ModifyFormWidget(FormWidget):
    def __init__(self, tab_widget, rows: List['FormRow'], on_update):
        super().__init__(parent=tab_widget, rows=rows)
        self.widgets[0].setReadOnly(True)  # ID field is read-only
        btn = QPushButton("保存修改")
        btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        self.layout.addWidget(btn)

        def on_pressed():
            if ConfirmDialog("确认要保存吗？").exec_() != QMessageBox.Ok:
                return
            successful = on_update(self.get_form())
            if successful:
                status.success("修改已成功提交至数据库")
                tab_widget.closeCurrentTab()
            else:
                status.failure("保存失败")

        btn.pressed.connect(on_pressed)
