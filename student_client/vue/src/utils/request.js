import axios from 'axios'
import {Message} from 'element-ui'
import store from '@/store'
import {getToken} from '@/utils/cookie'


export default function (server_addr, reportOnErr) {
  // create an axios instance
  let service = axios.create({
    baseURL: server_addr || process.env.VUE_APP_BASE_API, // url = base url + request url
    timeout: 5000 // request timeout
  })

  // request interceptor
  service.interceptors.request.use(
    config => {
      // do something before request is sent
      if (store.getters.token) {
        // let each request carry token
        config.headers['AUTHORIZATION'] = 'Bearer ' + getToken()
      }
      return config
    },
    error => {
      // do something with request error
      console.log(error) // for debug
      return Promise.reject(error)
    }
  )

  // response interceptor
  service.interceptors.response.use(
    response => {
      console.log(response)
      if (response.status !== 200) {
        if (reportOnErr) {
          Message({
            message: '网络错误',
            type: 'error',
            showClose: true,
          })
          return Promise.reject(new Error('网络错误'))
        } else {
          return Promise.reject(new Error("expected err"))
        }
      } else {
        return response.data
      }
    },
    error => {
      console.log('err' + error) // for debug
      Message({
        message: error.message,
        type: 'error',
        showClose: true,
      })
      return Promise.reject(error)
    }
  )

  return service
}
