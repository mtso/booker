import { RECEIVE_INCOMING } from '../actions/types'

const defaultState = {
  incoming: [],
  outgoing: [],
}

const trades = (state = defaultState, action) => {
  switch(action.type) {
    case RECEIVE_INCOMING:
      const incoming = action.trades
      return Object.assign({}, state, { incoming })

    default:
      return state
  }
}

export default trades
