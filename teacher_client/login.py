import sys

import qtmodern.styles
from PyQt5.QtWidgets import *

from teacher_client.api import login


class LoginDialog(QDialog):
    def __init__(self):
        super().__init__()
        w, h = 640, 480
        self.setFixedSize(w, h)

        self.server_addr = ''
        self.teacher_id = ''
        self.password = ''

        widget = QWidget(self)
        widget.setFixedSize(250, 200)
        widget.move((w - 250) // 2, (h - 200) // 2)
        form = QFormLayout()
        widget.setLayout(form)
        l = QLineEdit()
        l.textChanged.connect(lambda: setattr(self, 'server_addr', l.text()))
        l.setPlaceholderText("如：192.168.0.2:1234")
        form.addRow(QLabel("服务器地址"), l)
        l2 = QLineEdit()
        l2.textChanged.connect(lambda: setattr(self, 'teacher_id', l.text()))
        form.addRow(QLabel("工号"), l2)
        l3 = QLineEdit()
        l3.textChanged.connect(lambda: setattr(self, 'password', l.text()))
        form.addRow(QLabel("密码"), l3)
        btn = QPushButton("登录")
        btn.clicked.connect(lambda: self.login())
        form.addWidget(btn)

    def login(self) -> None:
        ok, err_msg = login.login(self.server_addr, self.teacher_id, self.password)
        if ok:
            self.accept()
        else:
            msg = QMessageBox()
            msg.setWindowTitle("Error")
            msg.setIcon(QMessageBox.Critical)
            msg.setText("Error")
            msg.setInformativeText(err_msg)
            msg.exec_()


if __name__ == "__main__":
    app = QApplication(sys.argv)
    qtmodern.styles.light(app)
    if LoginDialog().exec_() == QDialog.Accepted:
        print('succ')

    # window = QMainWindow()
    # window.show()
    # sys.exit(app.exec_())
