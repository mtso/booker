import React from 'react'
import Searchbox from './Searchbox'
import BookTable from './BookTable'

const BookCreator = ({ books }) => (
  <div>
    <h3>Add a Book</h3>
    <Searchbox />
    <BookTable books={books} />
  </div>
)

export default BookCreator
