import axios from 'axios'
import {Message} from 'element-ui'
import store from '@/store'
import {getToken} from '@/utils/auth'

// create an axios instance
export const remoteService = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

export const localService = axios.create({
  baseURL: process.env.STATIC_URL, // url = base url + request url
  timeout: 5000 // request timeout
})

const services = [remoteService, localService]

// request interceptor
services.forEach(service => {
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
})

// response interceptor
services.forEach(service => {
  service.interceptors.response.use(
    /**
     * If you want to get http information such as headers or status
     * Please return  response => response
     */

    response => {
      console.log(response)
      if (response.status !== 200) {
        Message({
          message: '网络错误',
          type: 'error',
          showClose: true,
        })
        return Promise.reject(new Error('网络错误'))
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
})
