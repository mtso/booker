import { RECEIVE_BOOKS, RECEIVE_MYBOOKS, RECEIVE_BOOKSEARCH } from './types'
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

// https://www.googleapis.com/books/v1/volumes?q=search+terms
export const searchBooks = (q) => (dispatch) => request
  .get('https://www.googleapis.com/books/v1/volumes')
  .query({ q })
  .then(({ body }) => body)
  .then(({ items }) => {
    return items.map(({ volumeInfo }) => {
      const { title, industryIdentifiers, imageLinks } = volumeInfo

      // const isbn = industryIdentifiers.reduce((isbn, id) => {
      //   if (isbn.type === 'ISBN_13') {
      //     return id.identifier
      //   }
      //   return isbn
      // }, '')
      const id = industryIdentifiers.filter(({ type }) => type === 'ISBN_13')[0]
      const { identifier } = id 
      const { thumbnail } = imageLinks

      const books = {
        title,
        isbn: identifier,
        image_url: thumbnail,
      }
      return books
    })
  })
  .then((books) => dispatch(receiveBookSearch(books)))

export const receiveBookSearch = (search) => ({
  type: RECEIVE_BOOKSEARCH,
  search,
})

export const receiveBooks = (books) => ({
  type: RECEIVE_BOOKS,
  books,
})

export const receiveMyBooks = (books) => ({
  type: RECEIVE_MYBOOKS,
  books,
})
