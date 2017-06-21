import React from 'react'
import BookCell from './BookCell'

const BookTable = ({ books }) => (
  <div>
    {books.map((book) => <BookCell {...book} />)}
  </div>
)

export default BookTable
