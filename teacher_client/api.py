from hashlib import sha256

from common import network


def login(teacher_id: str, password: str):
    return network.getAuth("/teacher_auth",
                           {"teacher_id": teacher_id, "password": sha256(password.encode("utf-8")).hexdigest()})


def getStudents(**kwargs):
    return network.get("/students", params=kwargs)


def getStudentsNum(**kwargs):
    return network.get("/students/num", params=kwargs)


def getTeachers(**kwargs):
    return network.get("/teachers", params=kwargs)


def getTeachersNum(**kwargs):
    return network.get("/teachers/num", params=kwargs)
