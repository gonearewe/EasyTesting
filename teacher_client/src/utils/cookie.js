import Cookies from 'js-cookie'

const AddrKey = 'easy_testing_server_addr'

export function getServerAddr() {
  return Cookies.get(AddrKey)
}

export function setServerAddr(token) {
  return Cookies.set(AddrKey, token)
}

const TokenKey = 'easy_testing_token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
