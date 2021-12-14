import request from '@/utils/request'

export function login(params) {
  return request({
    url: '/teacher_auth',
    method: 'get',
    params: params,
  })
}




