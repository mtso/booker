import React from 'react'
import BookCell from './BookCell'

const BookTable = ({ books }) => (
  <div>
    {books.map((book, i) => <BookCell key={i} {...book} />)}
  </div>
)

export default BookTable
