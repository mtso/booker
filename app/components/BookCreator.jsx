import React from 'react'
import Searchbox from './Searchbox'
import BookTable from './BookTable'
import Modal from './Modal'

const closeModal = (history) => () => {
  if (history.length < 3) {
    history.push('/')
  } else {
    history.goBack()
  }
}

const BookCreator = ({ books, onSearch, onAddBook, history, onClose }) => (
  <Modal>
    <div className='bookcreator-container'>
      <div className='bookcreator-header'>
        <h2 className='content-title'>Add a Book</h2>
        <button
          className='float-right'
          onClick={closeModal(history)}
        >Close</button>
        <Searchbox onSearch={onSearch} />
      </div>

      <BookTable
        isCreator={true}
        books={books}
        controls={(book) => (
          <div className='cell-detail' key={book.title}>
            <h2>{book.title}</h2>
            { book.pageCount && (
              <h4>{book.pageCount} pages</h4>
            ) }
            { book.categories && book.categories.map(
              (cat) => (<h4 key={cat}>{cat}</h4>)
            ) }
            <button
              className='action-button'
              onClick={() => onAddBook(book)}
            >Add Book</button>
          </div>
        )}
      />
    </div>
  </Modal>
)

export default BookCreator
