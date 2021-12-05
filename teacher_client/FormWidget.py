from enum import Enum
from typing import *

from PyQt5.QtWidgets import *


class FormWidget(QWidget):
    def __init__(self, parent, rows: Iterable['FormRow']):
        super().__init__(parent)
        self.layout = QFormLayout(self)
        self.setLayout(self.layout)
        for row in rows:
            v = None
            if row.type_ == FormValueType.SINGLE_LINE:
                v = QLineEdit(row.default_val, self)
            elif row.type_ == FormValueType.RICH_TEXT:
                v = QLabel("TODO")
            elif row.type_ == FormValueType.COMBO_BOX:
                v = QComboBox(self)
                v.addItems(row.default_val)
            self.layout.addRow(row.key, v)


class FormRow:
    def __init__(self, key: str, default_val: Union[str, Iterable], type_: 'FormValueType'):
        self.key = key
        self.default_val = default_val
        self.type_ = type_


class FormValueType(Enum):
    SINGLE_LINE = 1
    RICH_TEXT = 2
    COMBO_BOX = 3
