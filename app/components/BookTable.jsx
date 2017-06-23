import React from 'react'
import BookCell from './BookCell'

const BookTable = ({ books, className, controls }) => (
  <div className={className || 'book-table'}>
    {books.map((book, i, books) => (
      <BookCell
        key={i}
        {...book}
      >
        {controls && controls(book, i, books)}
      </BookCell>
    ))}
  </div>
)

export default BookTable
