import React from 'react'
import { Link } from 'react-router-dom'

const TradeCell = ({ book, user, book_id, children }) => (
  <div>
    <Link to={'/book/'+book_id}>
      <img src={book.image_url} />
    </Link>
    { children }
  </div>
)

export default TradeCell
