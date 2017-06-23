import { RECEIVE_BOOKS, RECEIVE_MYBOOKS } from '../actions/types'

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

    default:
      return state
  }
}

export default books
