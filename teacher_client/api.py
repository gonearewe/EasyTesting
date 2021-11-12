from common import network


def getStudents(**kwargs):
    return network.get("/students", params=kwargs)


def getStudentsNum(**kwargs):
    return network.get("/students/num", params=kwargs)
