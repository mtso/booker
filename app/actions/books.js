import { RECEIVE_BOOKS, RECEIVE_MYBOOKS } from './types'
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

export const getMyBooks = () => (dispatch) => request
  .get('/api/books/mybooks')
  .then(({ body }) => body)
  .then(({ ok, data, message }) => {
    if (ok) {
      dispatch(receiveMyBooks(data))
    } else {
      throw new Error(message)
    }
  })

export const receiveBooks = (books) => ({
  type: RECEIVE_BOOKS,
  books,
})

export const receiveMyBooks = (books) => ({
  type: RECEIVE_MYBOOKS,
  books,
})
