import abc
from typing import List, Tuple, Generator, Dict

import xlsxwriter
from PyQt5 import QtCore, Qt
from PyQt5.QtWidgets import *

from common.alert_dialog import AlertDialog


class TableWidget(QWidget):
    def __init__(self, parent, queries: List[Tuple[str, str]], columns: List, row_num: int, insertable: bool):
        super().__init__(parent)
        self.setContentsMargins(300,40,300,40)
        # self.setMaximumWidth(720)
        self.columns = columns[:-1]
        self.ops = columns[-1]
        self.row_num = row_num
        layout = QVBoxLayout(self)
        layout.setSpacing(30)

        # entries manipulation
        hbox0 = QHBoxLayout(self)
        hbox0.setSpacing(20)
        hbox0.setAlignment(QtCore.Qt.AlignLeft)
        layout.addLayout(hbox0)
        export_btn = QPushButton("导出")
        export_btn.setToolTip("导出所有数据至 Excel 文件")
        export_btn.pressed.connect(
            lambda: self.onExport(QFileDialog.getSaveFileName(self, filter="excel file (*.xlsx)")[0])
        )
        export_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)  # set policy in case it expands
        hbox0.addWidget(export_btn)
        if insertable:
            export_template_btn = QPushButton("导出模板")
            export_template_btn.setToolTip("生成 Excel 文件模板，仅包含列属性，而不包含数据")
            export_template_btn.pressed.connect(
                lambda: self.exportTemplate(QFileDialog.getSaveFileName(self, filter="excel file (*.xlsx)")[0])
            )
            export_template_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)  # set policy in case it expands
            hbox0.addWidget(export_template_btn)

            insert_btn = QPushButton("新建")
            insert_btn.setToolTip("新建一个条目")
            insert_btn.pressed.connect(
                lambda: self.onInsert()
            )
            insert_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)  # set policy in case it expands
            hbox0.addWidget(insert_btn)

            import_btn = QPushButton("导入")
            import_btn.setToolTip("批量新建条目，从指定格式的 Excel 文件导入数据，建议文件基于导出的模板添加数据")
            import_btn.pressed.connect(
                lambda: self.onImport(QFileDialog.getOpenFileName(self, filter="excel file (*.xlsx)")[0])
            )
            import_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)  # set policy in case it expands
            hbox0.addWidget(import_btn)

        # search bar
        hbox1 = QHBoxLayout(self)
        hbox1.setSpacing(20)
        hbox1.setAlignment(QtCore.Qt.AlignLeft)
        layout.addLayout(hbox1)
        self.edits = [QLineEdit() for _ in queries]
        for i, (text, query) in enumerate(queries):
            self.edits[i].setPlaceholderText(text)
            self.edits[i].setObjectName(query)
            self.edits[i].setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
            hbox1.addWidget(self.edits[i])
        self.searchBtn = QPushButton("搜索")
        self.searchBtn.pressed.connect(
            lambda: self.onSearch({edit.objectName(): edit.text().strip() for edit in self.edits}))
        self.searchBtn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        hbox1.addWidget(self.searchBtn)

        # table for data
        self.table = QTableWidget(self)
        layout.addWidget(self.table)
        self.table.setColumnCount(len(columns) + 1)
        headers = columns[:-1].copy()
        headers.insert(0, "序号")
        headers.append("操作")
        self.table.setHorizontalHeaderLabels(headers)
        # self.table.horizontalHeader().setSizePolicy(QSizePolicy.Preferred)
        self.table.setRowCount(row_num)
        self.table.resizeColumnsToContents()
        self.table.setEditTriggers(Qt.QAbstractItemView.NoEditTriggers)  # not directly editable

        # navigation bar
        hbox2 = QHBoxLayout(self)
        hbox2.setSpacing(20)
        hbox2.setAlignment(QtCore.Qt.AlignCenter)
        layout.addLayout(hbox2)
        self.prev_page_btn = QPushButton("<")
        self.prev_page_btn.setDisabled(True)
        self.prev_page_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        hbox2.addWidget(self.prev_page_btn)
        self._page_num = 0
        self._page_index = 0
        self.page_label = QLabel("0/0")
        self.page_label.setAlignment(QtCore.Qt.AlignCenter)
        hbox2.addWidget(self.page_label)
        self.next_page_btn = QPushButton(">")
        self.next_page_btn.setDisabled(True)
        self.next_page_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        hbox2.addWidget(self.next_page_btn)
        self.page_edit = QLineEdit()
        self.page_edit.setPlaceholderText("页码")
        self.page_edit.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        hbox2.addWidget(self.page_edit)
        self.page_btn = QPushButton("跳转")
        self.page_btn.setDisabled(True)
        self.page_edit.textChanged.connect(
            lambda text: self.page_btn.setDisabled(not (text.isdigit() and 1 <= int(text) <= self._page_num))
        )
        self.page_btn.pressed.connect(lambda: self.onTurnToPage(int(self.page_edit.text())))
        self.page_btn.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
        hbox2.addWidget(self.page_btn)

    @property
    def page_num(self):
        return self._page_num

    @page_num.setter
    def page_num(self, value):
        self.page_label.setText(f"{self.page_index}/{value}")
        self._page_num = value
        self.handle_page_btn()

    @property
    def page_index(self):
        return self._page_index

    @page_index.setter
    def page_index(self, value):
        self.page_label.setText(f"{value}/{self.page_num}")
        self._page_index = value
        self.handle_page_btn()

    def handle_page_btn(self):
        if 1 < self.page_index < self.page_num:
            self.prev_page_btn.setDisabled(False)
            self.next_page_btn.setDisabled(False)
        if self.page_index <= 1:
            self.prev_page_btn.setDisabled(True)
        if self.page_index >= self.page_num:
            self.next_page_btn.setDisabled(True)

    # data: [id, data for each column]
    def setData(self, page_num: int, page_index: int, data: List[List], op_callback):
        self.page_num = page_num
        self.page_index = page_index
        self.table.clearContents()
        for i, row in enumerate(data):
            # row index, starts with 1
            self.table.setItem(i, 0, QTableWidgetItem(str((page_index - 1) * self.row_num + i + 1)))
            # data columns
            for j, column in enumerate(data[i][1:]):
                self.table.setItem(i, j + 1, QTableWidgetItem(str(column)))
            # operation column
            w = QWidget(self.table)
            box = QHBoxLayout()
            w.setLayout(box)
            for op in self.ops:
                btn = QPushButton(w)
                btn.setText(op)
                # NOTICE: `ignored` is needed to receive pyqt slot parameter, in this case, a boolean
                btn.pressed.connect(lambda ignored, id=data[i][0], op=op: op_callback(id, op))
                box.addWidget(btn)
            self.table.setCellWidget(i, len(data[i]), w)
        self.table.resizeColumnsToContents()
        self.table.resizeRowsToContents()

    @abc.abstractmethod
    def onExport(self, filepath: str):
        pass

    def exportTemplate(self, filepath: str):
        try:
            workbook = xlsxwriter.Workbook(filepath)
            worksheet = workbook.add_worksheet()
            for i, column in enumerate(self.columns):
                worksheet.write(0, i, column)
            workbook.close()
        except Exception as e:
            print(e)
            AlertDialog("无法导出文件", detail=str(e))

    def exportData(self, filepath: str, data: Generator):
        self.exportTemplate(filepath)
        try:
            workbook = xlsxwriter.Workbook(filepath)
            worksheet = workbook.get_worksheet_by_name("Sheet1")
            for i, elem in enumerate(data):
                for j, content in enumerate(elem):
                    worksheet.write(i + 1, j, content)
            workbook.close()
        except Exception as e:
            print(e)
            AlertDialog("无法导出文件", detail=str(e))

    @abc.abstractmethod
    def onInsert(self):
        pass

    @abc.abstractmethod
    def onImport(self, filepath: str):
        pass

    @abc.abstractmethod
    def onSearch(self, queries: Dict[str, str]):
        pass

    @abc.abstractmethod
    def onTurnToPage(self, page_index: int):
        pass
