import React from 'react'
import TradeTable from '../components/TradeTable'
import { connect } from 'react-redux'
import { NavLink, Redirect, Route, withRouter } from 'react-router-dom'
import { getIncoming, getOutgoing, acceptTrade } from '../actions'

const mapStateToProps = ({ trades }) => ({ trades })

const mapDispatchToProps = (dispatch, { match }) => ({
  onAccept: (id) => () => dispatch(acceptTrade(id))
    .catch(console.warn),
  componentDidMount: dispatch(
    (match.url === '/trades/incoming')
      ? getIncoming()
      : getOutgoing()
  ).catch(console.warn),
})

const Trades = ({ trades, match, onAccept }) => (
  <div>
    <div className='tab-container'>
      <NavLink
        className='tab-link'
        activeClassName='active'
        to={'/trades/incoming'}
      >Incoming</NavLink>
      <NavLink
        className='tab-link'
        activeClassName='active'
        exact
        to={'/trades/outgoing'}
      >Outgoing</NavLink>
    </div>
    <Route path={'/trades/incoming'} component={() => (
      <TradeTable
        trades={
          // TODO: there may be an opportunity to try using one
          // Route with match.params.type or match.url, etc.
          trades.incoming
        } controls={
          ({ id }) => (<button onClick={onAccept(id)}>Accept Trade</button>)
        } cell={'temp'}
      />
    )} />
    <Route path={'/trades/outgoing'} component={() => (
      <TradeTable trades={trades.outgoing} cell={'temp'} />
    )} />
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(Trades))
