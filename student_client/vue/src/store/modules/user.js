import {login} from '@/api'
import {getToken, removeToken, setToken} from '@/utils/auth'
import jwt_decode from "jwt-decode"
import {parseTime} from "@/utils/time";

const getDefaultState = () => {
  return {
    student_id: '',
    name: '',
    class_id: 0,
    token: getToken(),
    exam_session_id: 0,
    exam_deadline: '',
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_STUDENT_ID: (state, student_id) => {
    state.student_id = student_id
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_CLASS_ID: (state, class_id) => {
    state.id = class_id
  },
  SET_EXAM_SESSION_ID: (state, exam_session_id) => {
    state.exam_session_id = exam_session_id
  },
  SET_EXAM_DEADLINE: (state, exam_deadline) => {
    state.exam_deadline = exam_deadline
  },
}

const actions = {
  // user login
  login({commit}, userInfo) {
    const {student_id, exam_id} = userInfo
    return new Promise((resolve, reject) => {
      login({student_id: student_id.trim(), exam_id: parseInt(exam_id.trim())}).then(data => {
        let token = data.token
        commit('SET_TOKEN', token)
        setToken(token)
        try {
          let decoded = jwt_decode(data.token)
          commit('SET_STUDENT_ID', decoded.student_id)
          commit('SET_NAME', decoded.name)
          commit('SET_CLASS_ID', decoded.class_id)
          commit('SET_EXAM_SESSION_ID', decoded.exam_session_id)
          commit('SET_EXAM_DEADLINE', parseTime(decoded.exam_deadline))
          resolve()
        } catch (error) {
          // return error in production env
          console.log('error from decoding token: ', error)
          reject(error)
        }
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({commit}) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      commit('SET_ROLES', [])
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

