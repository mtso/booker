import * as T from './types'
import handleError from '../utils/handleError'
import request from 'superagent'

export const submitLogin = (username, password) => (dispatch) => request
  .post('/auth/login')
  .send({
    username,
    password,
  })
  .then((resp) => {
    const { ok } = resp.body
    if (ok) {
      dispatch(login(username))
    } else {
      throw new Error("invalid password")
    }
  })

// export const submitLogout = () => (dispatch) => request
//   .post('/auth/logout')
//   .then((data) => {
//     data = JSON.parse(data)
//     if (data.ok) {
      
//     }
//   })

export const login = (username) => ({
  type: T.LOGIN,
  username,
})

export const logout = () => ({type: T.LOGOUT})
