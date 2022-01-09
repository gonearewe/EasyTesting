from abc import ABC, abstractmethod


class Question(ABC):
    @abstractmethod
    def __init__(self, id: int, stem: str):
        self.id = id
        self.stem = stem
        self.student_answer = None


class MultipleChoiceQuestion(Question):
    def __init__(self, id: int, stem: str, choices: (int, str)):
        super().__init__(id, stem)
        self.choices = choices


class MultipleAnswerQuestion(Question):
    def __init__(self, id: int, stem: str, choices: (int, str)):
        super().__init__(id, stem)
        self.choices = choices
        self.student_answer = set()


class BlankFillingQuestion(Question):
    def __init__(self, id: int, stem: str):
        super().__init__(id, stem)


class TrueFalseQuestion(Question):
    def __init__(self, id: int, stem: str):
        super().__init__(id, stem)


class CodeReadingQuestion(Question):
    def __init__(self, id: int, stem: str, blank_num: int):
        super().__init__(id, stem)
        self.student_answer = [""] * blank_num


class CodingQuestion(Question):
    def __init__(self, id: int, stem: str, template: str):
        super().__init__(id, stem)
        self.template = template
