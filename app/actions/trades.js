import * as T from './types'
import request from 'superagent'

// export const getIncoming = () => console.log('hello~')

// export const getIncoming = (m) => {
//   console.log(m)
//   return (dispatch) => request
//   .get('/api/trades/incoming')
//   .then(({ body }) => body)
//   .then(({ ok, trades, message }) => {
//     if (ok) {
//       dispatch(receiveIncoming(trades))
//     } else {
//       throw new Error(message)
//     }
//   })
// }
export const getIncoming = () => (dispatch) => request
  .get('/api/trades/incoming')
  .then(({ body }) => body)
  .then(({ ok, trades, message }) => {
    if (ok) {
      dispatch(receiveIncoming(trades))
    } else {
      throw new Error(message)
    }
  })

export const receiveIncoming = (trades) => ({
  type: T.RECEIVE_INCOMING,
  trades,
})
