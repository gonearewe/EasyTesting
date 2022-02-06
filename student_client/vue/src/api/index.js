import service from '@/utils/request'
import {getServerAddr} from "@/utils/cookie"

export function login(params) {
  return service(getServerAddr(), true)({
    url: '/student_auth',
    method: 'get',
    params: params,
  })
}

export function getMyQuestions() {
  return service(getServerAddr(), true)({
    url: '/exams/my_questions',
    method: 'get',
  })
}

export function runCode(body) {
  return service(process.env.STATIC_URL, true)({
    url: '/code',
    method: 'put',
    data: body
  })
}

export function submitMyAnswers(body) {
  return service(getServerAddr(), true)({
    url: '/exams/my_answers',
    method: 'put',
    data: body
  })
}

export function saveMyAnswerModels(body) {
  return service(getServerAddr(), true)({
    url: '/cache',
    method: 'put',
    data: JSON.stringify(body)
  })
}

export function loadMyAnswerModels() {
  return service(getServerAddr(), false)({
    url: '/cache',
    method: 'get',
  })
}


