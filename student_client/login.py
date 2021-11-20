from PyQt5.QtWidgets import *

from common import network
from student_client import api


class LoginDialog(QDialog):
    def __init__(self):
        super().__init__()
        w, h = 640, 480
        self.setFixedSize(w, h)

        self.server_addr = ""
        self.exam_id = ""
        self.student_id = ""

        widget = QWidget(self)
        widget.move((w - 500) // 2, (h - 200) // 2)
        form = QFormLayout()
        widget.setLayout(form)
        l = QLineEdit()
        l.textChanged.connect(lambda: setattr(self, 'server_addr', l.text()))
        l.setPlaceholderText("如：192.168.0.2:1234")
        l.setMinimumWidth(350)
        form.addRow(QLabel("服务器地址"), l)
        l2 = QLineEdit()
        l2.textChanged.connect(lambda: setattr(self, 'exam_id', l2.text()))
        form.addRow(QLabel("考试号"), l2)
        l3 = QLineEdit()
        l3.textChanged.connect(lambda: setattr(self, "student_id", l3.text()))
        form.addRow(QLabel("学号"), l3)
        btn = QPushButton("登录")
        btn.setFixedSize(200, 50)
        btn.clicked.connect(lambda: self.login())
        form.addWidget(btn)

    def login(self) -> None:
        err_msg = None
        ok = network.setServerAddr(self.server_addr)
        if not ok:
            err_msg = "服务器地址错误"
        elif api.login(self.exam_id, self.student_id) is False:
            err_msg = "工号或密码错误"
        else:
            self.accept()
        if err_msg is not None:
            msg = QMessageBox()
            msg.setWindowTitle("Error")
            msg.setIcon(QMessageBox.Critical)
            msg.setText("Error")
            msg.setInformativeText(err_msg)
            msg.exec_()
