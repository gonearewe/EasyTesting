from typing import List

from PyQt5.QtWidgets import *


class TableWidget(QWidget):
    def __init__(self, parent, columns: List, row_num: int):
        super().__init__(parent)
        self.ops = columns[-1]
        self.row_num = row_num
        layout = QVBoxLayout(self)

        hbox = QHBoxLayout(self)
        layout.addLayout(hbox)
        self.edits = [QLineEdit() for _ in columns]
        for i, column in enumerate(columns[:-1]):
            self.edits[i].setPlaceholderText(column)
            hbox.addWidget(self.edits[i])
        self.searchBtn = QPushButton("搜索")
        hbox.addWidget(self.searchBtn)

        self.table = QTableWidget(self)
        layout.addWidget(self.table)
        self.table.setColumnCount(len(columns) + 1)
        headers = columns[:-1].copy()
        headers.insert(0, "序号")
        headers.append("操作")
        self.table.setHorizontalHeaderLabels(headers)
        self.table.setRowCount(row_num)

    # data: [id, data for each column]
    def setData(self, page_num: int, page_index: int, data, op_callback):
        self.table.clearContents()
        for i, row in enumerate(data):
            # row index, starts with 1
            self.table.setItem(row=i, column=0,
                               item=QTableWidgetItem(text=str((page_index - 1) * self.row_num + i + 1)))
            # data columns
            for j, column in enumerate(data[i][1:]):
                self.table.setItem(row=i, column=j + 1, item=QTableWidgetItem(text=str(column)))
            # operation column
            w = QWidget(self.table)
            box = QHBoxLayout()
            w.setLayout(box)
            for op in self.ops:
                btn = QPushButton(w)
                btn.setText(op)
                btn.pressed.connect(lambda: op_callback(data[i][0], op))
                box.addWidget(btn)
            self.table.setCellWidget(row=i, column=len(data[i]), widget=w)

    def setOnSearch(self, callback):
        self.searchBtn.pressed.connect(lambda: callback([edit.text() for edit in self.edits]))
