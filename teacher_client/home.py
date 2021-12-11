from typing import List, Callable, Tuple

from PyQt5 import QtCore
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *

from common.card_widget import CardWidget
from teacher_client.mcq_table_widget import McqTableWidget
from teacher_client.student_table_widget import StudentTableWidget
from teacher_client.teacher_table_widget import TeacherTableWidget


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
            hbox = QHBoxLayout(self)
            layout.addLayout(hbox)
            for card in cards:
                tab_class = card[2]
                hbox.addWidget(CardWidget(
                    parent=self, text=card[0], image=card[1],
                    on_click=lambda cls=tab_class: (tab_widget.addTab(cls(tab_widget), cls.tab_name),
                                                    tab_widget.setCurrentIndex(tab_widget.count()-1))))

        add_header("试题管理", "./img/question.svg")
        add_cards([("单选题", "./img/mcq.svg", McqTableWidget), ("多选题", "./img/maq.svg", McqTableWidget)])

        # add_header("考试管理", "./img/exam.svg")
        add_header("用户管理", "./img/user.svg")
        add_cards([("学生", "./img/student.svg", StudentTableWidget), ("教师", "./img/teacher.svg", TeacherTableWidget)])
