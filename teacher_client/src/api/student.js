import request from '@/utils/request'


export function getStudents(params) {
  return request()({
    url: '/students',
    method: 'get',
    params: params
  })
}

export function createStudents(students) {
  return request()({
    url: '/students',
    method: 'post',
    data: students
  })
}

export function updateStudent(student) {
  return request()({
    url: '/students',
    method: 'put',
    data: student
  })
}

export function deleteStudents(ids) {
  return request()({
    url: '/students',
    method: 'delete',
    params: {'ids': ids.join(',')}
  })
}
