import { combineReducers } from 'redux'

import user from './user'
import books from './books'
import trades from './trades'

export default combineReducers({
  user,
  books,
  trades,
})
