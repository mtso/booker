import { LOGOUT, RECEIVE_BOOK, RECEIVE_BOOKS, RECEIVE_MYBOOKS, RECEIVE_BOOKSEARCH } from '../actions/types'

const defaultState = {
  all: [],
  mybooks: [],
  search: [],
  preview: {},
}

const books = (state = defaultState, action) => {
  switch(action.type) {
    case LOGOUT:
      return Object.assign({}, state, {
        search: [],
      })

    case RECEIVE_BOOK:
      Object.freeze(state)
      return Object.assign({}, state, {
        preview: action.book,
      })

    case RECEIVE_BOOKS:
      const all = action.books
      return Object.assign({}, state, {
        all,
      })

    case RECEIVE_MYBOOKS:
      const mybooks = action.books
      return Object.assign({}, state, {
        mybooks,
      })

    case RECEIVE_BOOKSEARCH:
      const { search } = action
      return Object.assign({}, state, {
        search,
      })

    default:
      return state
  }
}

export default books
