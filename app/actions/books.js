import { RECEIVE_BOOK, RECEIVE_BOOKS, RECEIVE_MYBOOKS, RECEIVE_BOOKSEARCH, CLEAR_SEARCH } from './types'
import request from 'superagent'

// Async thunks

export const getBook = (id) => (dispatch) => request
  .get('/api/book/'+id)
  .then(({ body }) => body)
  .then(({ ok, book, message }) => {
    if (ok) {
      dispatch(receiveBook(book))
    } else {
      throw new Error(message)
    }
  })

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

const getFirst = (arr, cb) => {
  for (let i = 0; i < arr.length; i++) {
    if (cb(arr[i], i, arr)) {
      return arr[i]
    }
  }
  return null
}

// https://www.googleapis.com/books/v1/volumes?q=search+terms
export const searchBooks = (q) => (dispatch) => request
  .get('https://www.googleapis.com/books/v1/volumes')
  .query({ q })
  .then(({ body }) => body)
  .then(({ items }) => items
    .filter(({ volumeInfo }) => {
      const { imageLinks, industryIdentifiers } = volumeInfo
      return imageLinks && industryIdentifiers && industryIdentifiers.some(({ type }) => type === 'ISBN_13')
    })
    .map(({ volumeInfo }) => {
      const { title, industryIdentifiers, imageLinks } = volumeInfo
      const { identifier } = getFirst(industryIdentifiers, ({ type }) => type === 'ISBN_13')
      const { thumbnail } = imageLinks

      const { description, pageCount, categories } = volumeInfo

      const books = {
        title,
        isbn: identifier,
        image_url: thumbnail,
        description,
        pageCount,
        categories,
      }
      return books
    }))
  .then((books) => dispatch(receiveBookSearch(books)))

export const addBook = (book) => (dispatch) => request
  .post('/api/book')
  .send(book)
  .then(({ body }) => body)

// Sync actions

export const receiveBook = (book) => ({
  type: RECEIVE_BOOK,
  book,
})

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

