from typing import List, Callable, Tuple

from PyQt5 import QtCore
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *

from common.card_widget import CardWidget
from common.flow_layout import FlowLayout
from teacher_client.maq_tab import MaqTab
from teacher_client.mcq_tab import McqTab
from teacher_client.student_tab import StudentTab


class HomeWidget(QWidget):
    def __init__(self, tab_widget: QTabWidget):
        super().__init__()
        self.setContentsMargins(100, 40, 100, 40)
        layout = QVBoxLayout(self)
        layout.setSpacing(30)

        def add_header(text: str, path: str):
            hbox = QHBoxLayout(self)
            hbox.setAlignment(QtCore.Qt.AlignCenter)
            hbox.setSpacing(20)
            layout.addLayout(hbox)
            img = QLabel(self)
            img.setPixmap(QPixmap(path))
            hbox.addWidget(img)
            label = QLabel(text + "----------------------------")
            label.setSizePolicy(QSizePolicy.Maximum, QSizePolicy.Maximum)
            hbox.addWidget(label)

        def add_cards(cards: List[Tuple[str, str, Callable]]):
            flow = FlowLayout(self)
            layout.addLayout(flow)
            for card in cards:
                flow.addWidget(CardWidget(
                    parent=self, text=card[0], image=card[1],
                    on_click=lambda: tab_widget.addTab(card[2](tab_widget))))

        add_header("试题管理", "./img/question.svg")
        add_cards([("单选题", "./img/mcq.svg", McqTab), ("多选题", "./img/maq.svg", MaqTab)])

        # add_header("考试管理", "./img/exam.svg")
        add_header("用户管理", "./img/user.svg")
        add_cards([("学生", "./img/student.svg", StudentTab)])
