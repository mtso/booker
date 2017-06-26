import React from 'react'
import TradeCell from './TradeCell'

const TradeTable = ({ cell: Cell, trades, className, controls, isIncoming }) => (
  <div className={className || 'trade-table'}>
    {trades.length > 0 ? (
      trades.map((trade, i, trades) => (
        <TradeCell
          key={i}
          {...trade}
          isIncoming={isIncoming}
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
