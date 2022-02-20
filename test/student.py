import random

import mysql.connector
from locust import FastHttpUser, task, between, events


# refresh database on locust init
@events.init.add_listener
def on_locust_init(environment, **kwargs):
    mydb = mysql.connector.connect(
        host="localhost",
        user="root",
        password="12345"
    )
    cursor = mydb.cursor()
    cursor.execute(open("../server/sql/test.sql", encoding='utf-8').read(), multi=True)


class Student(FastHttpUser):
    # Every Student will wait for 100~200ms after each task completion
    wait_time = between(0.1, 0.2)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        student = random.choice(students)
        response = self.client.get("/student_auth", params={**student, "exam_id": 4})
        self.auth = {"Authorization": "Bearer " + response.json()["token"]}
        response = self.client.get("/exams/my_questions", headers=self.auth)
        questions = response.json()
        self.answers = self.build_answers(questions)

    @task
    def participate(self):
        self.client.put("/exams/my_answers", json=self.answers, headers=self.auth)

    def build_answers(self, questions):
        return {
            "mcq": [{"id": q["id"], "answer": random.randint(1, 4)} for q in questions["mcq"]],
            "maq": [{"id": q["id"], "answer": random.choice([[], [1, 4], [2, 3, 1], [4]])} for q in questions["maq"]],
            "bfq": [{"id": q["id"], "answer": ["[678]"] * q["blank_num"]} for q in questions["bfq"]],
            "tfq": [{"id": q["id"], "answer": random.choice([True, False])} for q in questions["tfq"]],
            "crq": [{"id": q["id"], "answer": ["{xxx}"] * q["blank_num"]} for q in questions["crq"]],
            "cq": [{"id": q["id"], "answer": long_str, "right": False} for q in questions["cq"]],
        }


long_str = "long_str" * 500

students = [
    {'student_id': a, 'name': b} for a, b in [
        (2020501880, '小明'),
        (2020501826, '小红'),
        (2020501827, '小亮'),
        (2020501828, '小张'),
        (2020501829, '小李'),
        (2020501830, '小陆'),
        (2020501700, '小甲'),
        (2020501701, '小丁'),
        (2020501702, '小吴'),
        (2020501703, '小唐'),
        (2020201733, '小高'),
        (2020201734, '小岛'),
        (2020201735, '小凯'),
        (2020501096, '小雅'),
        (2020501098, '小伞'),
        (2020501099, '小坡'),
        (2019501826, '小古'),
        (2019501827, '小六'),
        (2019501829, '小六'),
        (2019501844, '小齐'),
        (2019501848, '小拍'),
        (2019501849, '小含'),
        (2018201826, '小示'),
        (2018216381, '小吞'),
        (2018216382, '小真'),
        (2018216385, '小夏'),
        (2018216386, '小阿'),
        (2018216387, '小金'),
        (2017216387, '小贵'),
        (2016664026, '小韩'),
    ]]
