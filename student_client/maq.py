import random
from typing import List

from PyQt5.QtWidgets import *


class MultipleAnswerQuestionWidget(QWidget):
    class Question:
        def __init__(self, id: int, stem: str, choices: (int, str)):
            self.id = id
            self.stem = stem
            self.choices = choices
            self.student_answer = set()

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
            btn_group.setExclusive(False)

            def update_answer(btn):
                i = btn_group.id(btn)
                if btn_group.checkedId() not in question.student_answer:
                    print(f"{i} clicked")
                    question.student_answer.add(i)
                else:
                    question.student_answer.remove(i)

            btn_group.buttonToggled.connect(update_answer)
            for choice in question.choices:
                btn = QCheckBox(text=choice[1])
                vbox.addWidget(btn)
                btn_group.addButton(btn, id=choice[0])
