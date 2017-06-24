import React from 'react'
import TradeTable from '../components/TradeTable'
import { connect } from 'react-redux'
import { NavLink, Redirect, Route, withRouter } from 'react-router-dom'
import { getIncoming } from '../actions'

const mapStateToProps = ({ trades }) => ({ trades })

const mapDispatchToProps = (dispatch, { match }) => ({
  componentDidMount: dispatch(
    (match.url === '/trades/incoming')
      ? getIncoming()
      : ({ type: 'NOOP' })
  ).catch(console.warn),
})

const Trades = ({ trades, match }) => (
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
      <TradeTable trades={trades.incoming}/>
    )} />
    <Route path={'/trades/outgoing'} component={() => (<p>outgoing</p>)} />
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(Trades))
