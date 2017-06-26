import { LOGIN, LOGOUT, RECEIVE_USER } from '../actions/types'

const defaultState = {
  username: null,
  city: null,
  state: null,
  display_name: null,
}

const user = (state = defaultState, action) => {
  switch(action.type) {
    case LOGIN:
      return Object.assign({}, state, action.user)

    case LOGOUT:
      return defaultState

    case RECEIVE_USER:
      return Object.assign({}, state, action.user)

    default:
      return state
  }
}

export default user
