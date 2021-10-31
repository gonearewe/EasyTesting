import random
from typing import List

from PyQt5.QtWidgets import *


class MultipleChoiceQuestionWidget(QWidget):
    class Question:
        def __init__(self, id: int, stem: str, choices: (int, str)):
            self.id = id
            self.stem = stem
            self.choices = choices
            self.student_answer = None

    def __init__(self, questions: List[Question]):
        super().__init__()
        self.questions = questions

        for question in questions:
            random.shuffle(question.choices)
        random.shuffle(questions)

        layout = QVBoxLayout(self)
        for question in questions:
            widget = QWidget()
            layout.addWidget(widget)
            vbox = QVBoxLayout(widget)
            text = QTextBrowser()
            vbox.addWidget(text)
            text.setText(question.stem)
            btn_group = QButtonGroup(vbox)
            btn_group.buttonClicked.connect(lambda btn: setattr(question, "student_answer", btn_group.id(btn)))
            for choice in question.choices:
                btn = QRadioButton(text=choice[1])
                vbox.addWidget(btn)
                btn_group.addButton(btn, id=choice[0])
