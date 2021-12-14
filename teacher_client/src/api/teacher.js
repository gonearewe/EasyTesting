import request from '@/utils/request'


export function getTeachers(params) {
  return request({
    url: '/teachers',
    method: 'get',
    params: params
  })
}

export function createTeachers(teachers) {
  return request({
    url: '/teachers',
    method: 'post',
    data: teachers
  })
}

export function updateTeacher(teacher) {
  return request({
    url: '/teachers',
    method: 'put',
    data: teacher
  })
}

export function deleteTeachers(ids){
  return request({
    url: '/teachers',
    method: 'delete',
    params: ids.join(',')
  })
}
