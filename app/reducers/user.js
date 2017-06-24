import { LOGIN, LOGOUT } from '../actions/types'

const defaultState = {
  username: null,
  city: null,
  state: null,
  display_name: null,
}

const user = (state = defaultState, action) => {
  switch(action.type) {
    case LOGIN:
      console.log(action.user)
      return Object.assign({}, state, action.user)

    case LOGOUT:
      return defaultState

    default:
      return state
  }
}

export default user
