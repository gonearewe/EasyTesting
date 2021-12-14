import {getInfo, login, logout} from '@/api/user'
import {getToken, removeToken, setToken} from '@/utils/auth'
import {resetRouter} from '@/router'
import jwt_decode from "jwt-decode"
import {sha256} from "js-sha256";

const getDefaultState = () => {
  return {
    token: getToken(),
    teacher_id: '',
    name: ''
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
  SET_TEACHER_ID: (state, teacher_id) => {
    state.teacher_id = teacher_id
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
}

const actions = {
  // user login
  login({commit}, userInfo) {
    const {teacher_id, password} = userInfo
    return new Promise((resolve, reject) => {
      login({teacher_id: teacher_id.trim(), password: sha256(password)}).then(data => {
        let token = data.token
        commit('SET_TOKEN', token)
        try {
          let decoded = jwt_decode(token)
          commit('SET_TEACHER_ID',decoded.teacher_id)
          commit('SET_NAME',decoded.name)
        } catch (error) {
          // return error in production env
          console.log( 'error from decoding token: ',error)
        }
        setToken(token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({commit, state}) {
    return new Promise((resolve, reject) => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
    })
  },

  // remove token
  resetToken({commit}) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
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

