import sys

import qtmodern.styles
from PyQt5.QtWidgets import *


class LoginDialog(QDialog):
    def __init__(self):
        super().__init__()
        w, h = 640, 480
        self.setFixedSize(w, h)

        widget = QWidget(self)
        widget.setFixedSize(250, 200)
        widget.move((w - 250) // 2, (h - 200) // 2)
        form = QFormLayout()
        widget.setLayout(form)
        l = QLineEdit()
        l.setPlaceholderText("如：192.168.0.2:1234")
        form.addRow(QLabel("服务器地址"), l)
        form.addRow(QLabel("考试号"), QLineEdit())
        form.addRow(QLabel("学号"), QLineEdit())
        btn = QPushButton("登录")
        btn.clicked.connect(lambda: self.login())
        form.addWidget(btn)

    def login(self) -> None:
        # TODO
        self.accept()


class AlertDialog(QDialog):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("登录失败！")

        buttonBox = QDialogButtonBox(QDialogButtonBox.Ok)
        buttonBox.accepted.connect(self.accept)

        layout = QVBoxLayout(self)
        layout.addWidget(QLabel("Something happened, is that OK?"))
        layout.addWidget(buttonBox)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    while True:
        if LoginDialog().exec_() == QDialog.Rejected:
            AlertDialog().exec_()
            continue

    # window = QMainWindow()
    # window.show()
    # sys.exit(app.exec_())
