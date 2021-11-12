from PyQt5.QtWidgets import QMessageBox


class AlertDialog(QMessageBox):
    def __init__(self, text: str, detail: str = None):
        super().__init__()
        self.setWindowTitle("错误")
        self.setIcon(QMessageBox.Critical)
        self.setText(text)
        if detail is not None:
            self.setDetailedText(detail)
        self.setStandardButtons(QMessageBox.Ok)
