import React from 'react'
import Searchbox from './Searchbox'
import BookTable from './BookTable'

const BookCreator = ({ books, onSearch }) => (
  <div>
    <h3>Add a Book</h3>
    <Searchbox onSearch={onSearch} />
    <BookTable books={books} />
  </div>
)

export default BookCreator
