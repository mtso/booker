import { RECEIVE_BOOKS } from './types'
import request from 'superagent'

export const getBooks = () => (dispatch) => request
  .get('/api/books')
  .then(({ body }) => body)
  .then(({ ok, data, message }) => {
    if (ok) {
      dispatch(receiveBooks(data))
    } else {
      throw new Error(message)
    }
  })

export const receiveBooks = (books) => ({
  type: RECEIVE_BOOKS,
  books,
})
