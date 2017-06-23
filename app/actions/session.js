import * as T from './types'
import request from 'superagent'

export const submitLogin = (username, password) => (dispatch) => request
  .post('/auth/login')
  .send({
    username,
    password,
  })
  .then(({ body }) => body)
  .then(({ ok }) => {
    if (ok) {
      dispatch(login(username))
    } else {
      throw new Error("invalid password")
    }
  })

export const submitLogout = () => (dispatch) => request
  .post('/auth/logout')
  .then(({ body }) => body)
  .then(({ ok }) => {
    if (ok) {
      dispatch(logout())
    } else {
      throw new Error("unable to logout")
    }
  })

export const login = (username) => ({
  type: T.LOGIN,
  username,
})

export const logout = () => ({type: T.LOGOUT})
