import axios from 'axios'
import {Message} from 'element-ui'
import store from '@/store'
import {getServerAddr, getToken} from '@/utils/cookie'

export default function () {
  // create an axios instance
  let service = axios.create({
    baseURL: getServerAddr(), // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
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
      console.log('request config: ')
      console.log(config)
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
    /**
     * If you want to get http information such as headers or status
     * Please return  response => response
     */

    response => {
      console.log('response: ')
      console.log(response)
      if (response.status !== 200) {
        Message({
          message: 'Error',
          type: 'error',
          showClose: true,
          duration: 5 * 1000
        })
        return Promise.reject(new Error('Error'))
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
        duration: 5 * 1000
      })
      return Promise.reject(error)
    }
  )

  return service
}
