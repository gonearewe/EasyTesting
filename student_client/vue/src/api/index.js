import {localService, remoteService} from '@/utils/request'
import axios from "axios";
import store from "@/store";
import {getToken} from "@/utils/auth";
import {Message} from "element-ui";

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

export function saveMyAnswerModels(body) {
  return remoteService({
    url: '/cache',
    method: 'put',
    data: JSON.stringify(body)
  })
}

export function loadMyAnswerModels() {
  let service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  })
  service.interceptors.request.use(
    config => {
      if (store.getters.token) {
        config.headers['AUTHORIZATION'] = 'Bearer ' + getToken()
      }
      return config
    }
  )
  service.interceptors.response.use(
    response => {
      console.log(response)
      if (response.status === 200) {
        return response.data
      } else {
        Promise.reject(new Error("expected err"))
      }
    }
  )
  return service({
    url: '/cache',
    method: 'get',
  })
}


