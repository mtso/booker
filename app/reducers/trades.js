import { RECEIVE_INCOMING, RECEIVE_OUTGOING } from '../actions/types'

const defaultState = {
  incoming: [],
  outgoing: [],
}

const trades = (state = defaultState, action) => {
  switch(action.type) {
    case RECEIVE_INCOMING:
      const incoming = action.trades
      return Object.assign({}, state, { incoming })

    case RECEIVE_OUTGOING:
      const outgoing = action.trades
      return Object.assign({}, state, { outgoing })

    default:
      return state
  }
}

export default trades
