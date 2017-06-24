import React from 'react'
import TradeCell from './TradeCell'

const TradeTable = ({ cell: Cell, trades, className, controls }) => (
  <div className={className || 'trade-table'}>
    {trades.length > 0 ? (
      trades.map((trade, i, trades) => (
        <TradeCell
          key={i}
          {...trade}
        >
          {controls && controls(trade, i, trades)}
        </TradeCell>
      ))
    ) : (
      <div>No trades.</div>
    )}
  </div>
)

export default TradeTable
