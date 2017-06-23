import React from 'react'
import { Link } from 'react-router-dom'

const TradeCell = ({ id, book_id, user_id, status, book, user, children }) => (
  <div>
    <Link to={'/book/'+book_id}>
      <img style={{width: '200px'}} src={book.image_url} />
    </Link>
    <pre>id      {id}</pre>
    <pre>book_id {book_id}</pre>
    <pre>user_id {user_id}</pre>
    <pre>status  {status}</pre>
    <pre>book    {JSON.stringify(book)}</pre>
    <pre>user    {JSON.stringify(user)}</pre>

    { children }
  </div>
)

export default TradeCell
