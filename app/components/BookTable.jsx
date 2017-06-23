import React from 'react'
import BookCell from './BookCell'

const BookTable = ({ books, className, controls }) => (
  <div className={className || 'book-table'}>
    {books.map((book, i, j, k) => (
      <BookCell
        key={i}
        {...book}
      >
        {controls && controls(book, i, j, k)}
      </BookCell>
    ))}
  </div>
)

export default BookTable
