import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import rootReducer from '../reducers'

const initialState = (typeof window === 'undefined')
  ? undefined
  : window.__PRELOADED_STATE__

const store = createStore(
  rootReducer,
  initialState,
  applyMiddleware(thunk)
)

export default store
