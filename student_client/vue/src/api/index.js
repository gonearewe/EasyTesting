import {localService, remoteService} from '@/utils/request'

export function login(params) {
  return remoteService({
    url: '/student_auth',
    method: 'get',
    params: params,
  })
}

export function getMyQuestions() {
  return remoteService({
    url: '/exams/my_questions',
    method: 'get',
  })
}

export function runCode(body) {
  return localService({
    url: '/code',
    method: 'put',
    data: body
  })
}

export function submitMyAnswers(body) {
  return remoteService({
    url: '/exams/my_answers',
    method: 'put',
    data: body
  })
}

export function saveMyAnswersLocal(body, key) {
  return localService({
    url: '/my_answers',
    method: 'put',
    data: {...body, key: key}
  })
}



