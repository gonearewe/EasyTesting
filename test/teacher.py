import random

from locust import FastHttpUser, task


class Teacher(FastHttpUser):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        teacher = random.choice(teachers)
        response = self.client.get("/teacher_auth", params=teacher)
        self.auth = {"Authorization": "Bearer " + response.json()["token"]}
        self.mcq = {
            "id": random.randint(1, 30),
            "stem": "stem" * 50,
            "choices": [
                "2" * 50,
                "3" * 50,
                "4" * 50,
                "5" * 50
            ],
            "right_answer": random.randint(1, 4)
        }

    @task
    def hello(self):
        self.client.get("/hello", headers=self.auth)

    @task
    def get_students(self):
        self.client.get("/students", params={"name": "小", "page_size": 50, "page_index": 1}, headers=self.auth)

    @task
    def put_mcq(self):
        self.client.put("/mcq", json=self.mcq, headers=self.auth)


teachers = [
    {'teacher_id': a, 'password': b} for a, b in [
        ('0', 'e66ac56a12d26003451e18a29215995ff52c26441b28c399bb6ce45e9b81fad8'),
        ('2010301800', '8fcd1836eba80875018ae5bb1f5e0754ccd28ff157379ca1a29db8f6d4450726'),
        ('2012550921', '721396f6e3a63c490ba3cbba3069506e6fa56d86cef7b142a5ee267230e96fca'),
    ]]
