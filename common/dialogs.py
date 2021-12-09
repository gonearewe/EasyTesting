from typing import List, Iterator

from PyQt5 import QtCore
from PyQt5.QtWidgets import QMessageBox, QVBoxLayout, QTableWidget, QTableWidgetItem


class AlertDialog(QMessageBox):
    def __init__(self, text: str, detail: str = None):
        super().__init__()
        self.setWindowTitle("错误")
        self.setIcon(QMessageBox.Critical)
        self.setText(text)
        if detail is not None:
            self.setDetailedText(detail)
        self.setStandardButtons(QMessageBox.Ok)


class ConfirmDialog(QMessageBox):
    def __init__(self, text: str, detail: str = None):
        super().__init__()
        self.setWindowTitle("确认操作")
        self.setIcon(QMessageBox.Information)
        self.setText(text)
        if detail is not None:
            self.setDetailedText(detail)
        self.setStandardButtons(QMessageBox.Ok | QMessageBox.Cancel)


# BUG
class ConfirmDialogWithTable(QMessageBox):
    def __init__(self, text: str, header: List[str], data: Iterator[List]):
        super().__init__()
        self.setWindowTitle("确认操作")
        self.setIcon(QMessageBox.Information)
        self.setText(text)
        self.setStandardButtons(QMessageBox.Ok | QMessageBox.Cancel)

        layout = QVBoxLayout(self)
        self.setLayout(layout)
        table = QTableWidget(self)
        table.setHorizontalHeaderLabels(header)
        for i, row in enumerate(data):
            for j, cell in enumerate(row):
                item = QTableWidgetItem(str(cell))
                item.setTextAlignment(QtCore.Qt.AlignCenter)
                table.setItem(i, j, item)
        layout.addWidget(table)
