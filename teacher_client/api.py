from hashlib import sha256
from typing import List, Dict

from common import network


def login(teacher_id: str, password: str):
    return network.getAuth("/teacher_auth",
                           {"teacher_id": teacher_id, "password": sha256(password.encode("utf-8")).hexdigest()})


# student

def getStudents(**kwargs):
    return network.get("/students", params=kwargs)


def getStudentsNum(**kwargs):
    return network.get("/students/num", params=kwargs)


def postStudents(body: List[Dict]):
    return network.post("/students", body)


def putStudents(body: List[Dict]):
    return network.put("/students", body)


def delStudents(**kwargs):
    return network.delete("/students", params=kwargs)


# teacher

def getTeachers(**kwargs):
    return network.get("/teachers", params=kwargs)


def getTeachersNum(**kwargs):
    return network.get("/teachers/num", params=kwargs)


def postTeachers(body: List[Dict]):
    return network.post("/teachers", body)


def putTeachers(body: List[Dict]):
    return network.put("/teachers", body)


def delTeachers(**kwargs):
    return network.delete("/teachers", params=kwargs)


# mcq

def getMcqs(**kwargs):
    return network.get("/mcq", params=kwargs)


def getMcqsNum(**kwargs):
    return network.get("/mcq/num", params=kwargs)


def postMcqs(body: List[Dict]):
    return network.post("/mcq", body)


def putMcqs(body: List[Dict]):
    return network.put("/mcq", body)


def delMcqs(**kwargs):
    return network.delete("/mcq", params=kwargs)
