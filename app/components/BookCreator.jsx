import React from 'react'
import Searchbox from './Searchbox'
import BookTable from './BookTable'
import Modal from './Modal'

const BookCreator = ({ books, onSearch, onAddBook, history }) => (
  <Modal>
    <h3>Add a Book <button onClick={() => {
      if (history.length < 3) {
        history.push('/')
      } else {
        history.goBack()
      }
    }}>Close</button></h3>
    <Searchbox onSearch={onSearch} />
    <BookTable
      books={books}
      controls={(book) => (
        <button onClick={() => onAddBook(book)}>Add Book</button>
      )}
    />
  </Modal>
)

export default BookCreator
