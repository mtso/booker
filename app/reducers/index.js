import { combineReducers } from 'redux'

import username from './username'
import books from './books'

export default combineReducers({
  username,
  books,
})
