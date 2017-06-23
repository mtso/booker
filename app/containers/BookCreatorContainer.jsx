import BookCreator from '../components/BookCreator'
import { connect } from 'react-redux'
import { searchBooks, addBook } from '../actions/books'
import { withRouter } from 'react-router-dom'

const mapStateToProps = ({ books }) => ({
  books: books.search,
})

const mapDispatchToProps = (dispatch, { history }) => ({
  onSearch: (term) => dispatch(searchBooks(term))
    .catch(console.warn),
  onAddBook: (book) => dispatch(addBook(book))
    .then(({ ok, message }) => {
      if (ok) {
        history.goBack()
      } else {
        throw new Error(message)
      }
    })
    .catch(console.warn),
})

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookCreator))
