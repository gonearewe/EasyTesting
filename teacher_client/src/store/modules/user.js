import {login} from '@/api/user'
import {getToken, removeToken, setToken} from '@/utils/auth'
import {resetRouter} from '@/router'
import jwt_decode from "jwt-decode"
import {sha256} from "js-sha256"

const getDefaultState = () => {
  return {
    token: getToken(),
    teacher_id: '',
    name: '',
    id: '',
    roles: []
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
  SET_ID: (state, id) => {
    state.id = id
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  }
}

const actions = {
  // user login
  login({commit}, userInfo) {
    const {teacher_id, password} = userInfo
    return new Promise((resolve, reject) => {
      login({teacher_id: teacher_id.trim(), password: sha256(password)}).then(data => {
        let token = data.token
        commit('SET_TOKEN', token)
        setToken(token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({commit, state}) {
    return new Promise((resolve, reject) => {
      try {
        let decoded = jwt_decode(state.token)
        commit('SET_TEACHER_ID', decoded.teacher_id)
        commit('SET_NAME', decoded.name)
        commit('SET_ID', decoded.id)
        commit('SET_ROLES', decoded.is_admin ? ['admin'] : ['teacher'])
        resolve({roles: decoded.is_admin ? ['admin'] : ['teacher']})
      } catch (error) {
        // return error in production env
        console.log('error from decoding token: ', error)
        reject(error)
      }
    })
  },

  // user logout
  logout({commit, state}) {
    return new Promise((resolve, reject) => {
      removeToken() // must remove  token  first
      resetRouter()
      commit('RESET_STATE')
      commit('SET_ROLES', [])
      resolve()
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

