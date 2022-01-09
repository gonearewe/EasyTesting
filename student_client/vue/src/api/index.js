import request from '@/utils/request'

export function login(params) {
  return request({
    url: '/student_auth',
    method: 'get',
    params: params,
  })
}

export function getMyQuestions() {
  return request({
    url: '/exams/my_questions',
    method: 'get',
  })
}






