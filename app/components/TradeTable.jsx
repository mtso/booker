import React from 'react'
import TradeCell from './TradeCell'

const TradeTable = ({ cell: Cell, trades, className, controls }) => (
  <div className={className || 'trade-table'}>
    {trades.map((trade, i, trades) => (
      <TradeCell
        key={i}
        {...trade}
      >
        {controls && controls(trade, i, trades)}
      </TradeCell>
    ))}
  </div>
)

export default TradeTable
