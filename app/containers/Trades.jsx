import React from 'react'
import TradeTable from '../components/TradeTable'
import { connect } from 'react-redux'
import { NavLink, Redirect, Route, withRouter } from 'react-router-dom'
import { getIncoming, getOutgoing, acceptTrade, cancelTrade } from '../actions'

const mapStateToProps = ({ trades }) => ({ trades })

const mapDispatchToProps = (dispatch, { match }) => ({
  componentDidMount: dispatch(
    (match.url === '/trades/incoming')
      ? getIncoming()
      : getOutgoing()
  ).catch(console.warn),
  onAccept: (id) => () => dispatch(acceptTrade(id))
    .catch(console.warn),
  onCancel: (id) => () => dispatch(cancelTrade(id))
    .catch(console.warn),
})

const Trades = ({ trades, match, onAccept, onCancel }) => (
  <div className='page-container'>
    <div className='page-header'>
      <h2 className='content-title'>
        Trades
      </h2>
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
    </div>
    <div className='content-container'>
      <Route path={'/trades/incoming'} component={() => (
        <TradeTable
          trades={
            // TODO: there may be an opportunity to try using one
            // Route with match.params.type or match.url, etc.
            trades.incoming
          }
          controls={
            ({ id }) => (<button onClick={onAccept(id)}>Accept Trade</button>)
          }
          cell={'temp'}
        />
      )} />
      <Route path={'/trades/outgoing'} component={() => (
        <TradeTable
          trades={trades.outgoing}
          controls={
            ({ id }) => (<button onClick={onCancel(id)}>Cancel Trade</button>)
          }
          cell={'temp'}
        />
      )} />
    </div>
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(Trades))
