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
      request
        .get('/api/user/'+username)
        .then(({ body }) => body)
        .then(({ ok, message, ...user }) => {
          if (ok) {
            dispatch(login(user))
          } else {
            throw new Error(message)
          }
        })
        .catch(console.warn)
    } else {
      throw new Error("invalid password")
    }
  })

export const submitSignup = (username, password) => (dispatch) => request
  .post('/auth/signup')
  .send({
    username,
    password,
  })
  .then(({ body }) => body)
  .then(({ ok, ...rest }) => {
    console.log(ok, rest)
    if (ok) {
      return request
        .get('/api/user/'+username)
        .then(({ body }) => body)
        .then(({ ok, message, ...user }) => {
          console.log(ok, message, user)
          if (ok) {
            return dispatch(login(user))
          } else {
            throw new Error(message)
          }
        })
    } else {
      throw new Error("invalid password")
    }
  })

export const saveProfile = (profile) => (dispatch) => request
  .post('/api/user')
  .send({ ...profile })
  .then(({ body }) => body)
  .then(({ ok, ...rest }) => {
    console.log(rest)
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

export const login = (user) => ({
  type: T.LOGIN,
  user,
})

export const logout = () => ({type: T.LOGOUT})

export const receiveUser = (user) => ({
  type: T.RECEIVE_USER,
  user,
})
