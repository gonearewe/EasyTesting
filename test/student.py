# Run this with `pipenv run locust -f .\student.py -H http://localhost:9345`,
# and use Web UI at `http://localhost:8089/`
import random

from locust import FastHttpUser, task, between


class Student(FastHttpUser):
    # Every Student will wait for 500~1000ms after each task completion
    wait_time = between(0.5, 1)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.student = random.choice(students)
        response = self.client.get("/student_auth", params={**self.student, "exam_id": 4})
        self.auth = {"Authorization": "Bearer " + response.json()["token"]}
        response = self.client.get("/exams/my_questions", headers=self.auth)
        self.questions = response.json()
        self.validateQuestions(self.questions)
        self.answers = self.build_answers(self.questions)

    def validateQuestions(self, questions):
        for q in ["mcq", "maq", "bfq", "tfq", "crq", "cq"]:
            li = [e["id"] for e in questions[q]]
            # validate that there's no duplicated question issued
            assert len(li) == len(set(li))

    @task(10)
    def participate(self):
        self.client.put("/exams/my_answers", json=self.answers, headers=self.auth)
        self.client.put("/cache", data=bytes(long_str, 'utf-8'), headers=self.auth)

    @task
    def restart(self):
        self.client.get("/student_auth", params={**self.student, "exam_id": 4})
        response = self.client.get("/exams/my_questions", headers=self.auth)
        assert self.questions == response.json()
        self.client.get("/cache", headers=self.auth)

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
        (2020501880, '??????'),
        (2020501826, '??????'),
        (2020501827, '??????'),
        (2020501828, '??????'),
        (2020501829, '??????'),
        (2020501830, '??????'),
        (2020501700, '??????'),
        (2020501701, '??????'),
        (2020501702, '??????'),
        (2020501703, '??????'),
        (2020201733, '??????'),
        (2020201734, '??????'),
        (2020201735, '??????'),
        (2020501096, '??????'),
        (2020501098, '??????'),
        (2020501099, '??????'),
        (2019501826, '??????'),
        (2019501827, '??????'),
        (2019501829, '??????'),
        (2019501844, '??????'),
        (2019501848, '??????'),
        (2019501849, '??????'),
        (2018201826, '??????'),
        (2018216381, '??????'),
        (2018216382, '??????'),
        (2018216385, '??????'),
        (2018216386, '??????'),
        (2018216387, '??????'),
        (2017216387, '??????'),
        (2016664026, '??????'),
    ]]
