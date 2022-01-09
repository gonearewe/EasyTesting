import random
from typing import List

from PyQt5.QtWidgets import *
from pyqode.python.backend import server
from pyqode.python.widgets import PyCodeEdit

from student_client.pyqt.question import *


class QuestionWidget(QWidget):
    def __init__(self, questions: List[Question]):
        super().__init__()
        self.questions = questions

        if isinstance(questions[0], (MultipleChoiceQuestion, MultipleAnswerQuestion)):
            for question in questions:
                random.shuffle(question.choices)
        random.shuffle(questions)

        layout = QVBoxLayout(self)
        layout.setContentsMargins(500, 0, 500, 0)
        for index, question in enumerate(questions):
            widget = QWidget()
            layout.addWidget(widget)
            vbox = QVBoxLayout(widget)

            # question number
            vbox.addWidget(QLabel(str(index + 1)))

            # stem of question
            text = QTextEdit()
            text.setReadOnly(True)
            text.setMarkdown(question.stem)
            vbox.addWidget(text)

            if isinstance(question, MultipleChoiceQuestion):
                # choices of question
                btn_group = QButtonGroup(vbox)
                btn_group.buttonClicked.connect(lambda btn: setattr(question, "student_answer", btn_group.id(btn)))
                for choice in question.choices:
                    btn = QRadioButton(text=choice[1])
                    vbox.addWidget(btn)
                    btn_group.addButton(btn, id=choice[0])
            elif isinstance(question, MultipleAnswerQuestion):
                # choices of question

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
            elif isinstance(question, BlankFillingQuestion):
                # blank of question
                blank = QLineEdit()
                blank.editingFinished.connect(lambda: setattr(question, "student_answer", blank.text()))
                vbox.addWidget(blank)
            elif isinstance(question, TrueFalseQuestion):
                # (only 2) choices of question
                btn_group = QButtonGroup(vbox)
                btn_group.buttonClicked.connect(lambda btn: setattr(question, "student_answer", btn_group.id(btn) == 1))
                btn_false = QRadioButton(text="False")
                vbox.addWidget(btn_false)
                btn_group.addButton(btn_false, id=0)
                btn_true = QRadioButton(text="True")
                vbox.addWidget(btn_true)
                btn_group.addButton(btn_true, id=1)
            elif isinstance(question, CodeReadingQuestion):
                # blanks of question
                blanks = QGridLayout()
                vbox.addLayout(blanks)
                for i in range(len(question.student_answer)):
                    blanks.addWidget(QLabel(str(i + 1)), i // 2, i % 2 * 2)
                    blank = QLineEdit()
                    blank.editingFinished.connect(lambda: setattr(question, f"student_answer[{i}]", blank.text()))
                    blanks.addWidget(blank, i // 2, i % 2 * 2 + 1)
            elif isinstance(question, CodingQuestion):
                # code editor
                editor = PyCodeEdit(server_script=server.__file__)
                # show the PyCodeEdit module in the editor
                editor.setPlainText(question.template)
                editor.textChanged.connect(lambda: setattr(question, "student_answer", editor.text()))
