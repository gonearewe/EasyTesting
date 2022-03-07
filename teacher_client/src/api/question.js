import request from '@/utils/request'


export function getQuestions(type, params) {
  return request()({
    url: '/' + type,
    method: 'get',
    params: params
  })
}

export function createQuestions(type, questions) {
  return request()({
    url: '/' + type,
    method: 'post',
    data: questions
  })
}

export function updateQuestion(type, question) {
  return request()({
    url: '/' + type,
    method: 'put',
    data: question
  })
}

export function deleteQuestions(type, ids) {
  return request()({
    url: '/' + type,
    method: 'delete',
    params: {'ids': ids.join(',')}
  })
}
