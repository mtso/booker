import React from 'react'
import BookCell from './BookCell'

const BookTable = ({ books, className }) => (
  <div className={className || 'book-table'}>
    {books.map((book, i) => <BookCell key={i} {...book} />)}
  </div>
)

export default BookTable
