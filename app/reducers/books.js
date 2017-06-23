import { RECEIVE_BOOKS, RECEIVE_MYBOOKS, RECEIVE_BOOKSEARCH } from '../actions/types'

const defaultState = {
  all: [],
  mybooks: [],
  search: [],
}

const books = (state = defaultState, action) => {
  switch(action.type) {
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
