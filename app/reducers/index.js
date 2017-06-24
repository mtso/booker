import { combineReducers } from 'redux'

import username from './username'
import books from './books'
import trades from './trades'

export default combineReducers({
  username,
  books,
  trades,
})
