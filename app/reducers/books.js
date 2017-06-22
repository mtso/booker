import { RECEIVE_BOOKS } from '../actions/types'

const defaultState = {
  books: [],
  mybooks: [],
}

const books = (state = defaultState, action) => {
  switch(action.type) {
    case RECEIVE_BOOKS:
      const newBooks = action.books
      const newState = Object.assign({}, state, {
        books: newBooks,
      })
      return newState
    default:
      return state
  }
}

export default books
