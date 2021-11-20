from common import network


def login(exam_id: str, student_id: str):
    return network.getAuth("/student_auth", {"exam_id": exam_id, "student_id": student_id})
