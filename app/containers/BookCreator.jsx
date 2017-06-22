import React from 'react'
import Searchbox from '../components/Searchbox'
import BookTable from '../components/BookTable'

const BookCreator = () => (
  <div>
    Add a Book
    <Searchbox />
    <BookTable books={[{}]} />
  </div>
)

export default BookCreator
