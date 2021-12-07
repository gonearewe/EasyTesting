from hashlib import sha256
from typing import List, Dict

from common import network


def login(teacher_id: str, password: str):
    return network.getAuth("/teacher_auth",
                           {"teacher_id": teacher_id, "password": sha256(password.encode("utf-8")).hexdigest()})


def getStudents(**kwargs):
    return network.get("/students", params=kwargs)


def getStudentsNum(**kwargs):
    return network.get("/students/num", params=kwargs)


def postStudents(**kwargs):
    return network.post("/students", params=kwargs)


def putStudents(**kwargs):
    return network.put("/students", params=kwargs)


def delStudents(**kwargs):
    return network.delete("/students", params=kwargs)


def getTeachers(**kwargs):
    return network.get("/teachers", params=kwargs)


def getTeachersNum(**kwargs):
    return network.get("/teachers/num", params=kwargs)


def postTeachers(body:List[Dict]):
    return network.post("/teachers", body)


def putTeachers(body:List[Dict]):
    return network.put("/teachers", body)


def delTeachers(**kwargs):
    return network.delete("/teachers", params=kwargs)
