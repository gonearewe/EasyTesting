import request from '@/utils/request'


export function getExams(params) {
  return request()({
    url: '/exams',
    method: 'get',
    params: params
  })
}

export function getEndedExams(params) {
  return request()({
    url: '/exams/ended',
    method: 'get',
    params: params
  })
}

export function createExams(exams) {
  return request()({
    url: '/exams',
    method: 'post',
    data: exams
  })
}

export function updateExam(exam) {
  return request()({
    url: '/exams',
    method: 'put',
    data: exam
  })
}

export function deleteExams(ids) {
  return request()({
    url: '/exams',
    method: 'delete',
    params: {'ids': ids.join(',')}
  })
}

export function getExaminees(params) {
  return request()({
    url: '/exams/examinees',
    method: 'get',
    params: params
  })
}
