import BookCreator from '../components/BookCreator'
import { connect } from 'react-redux'
import { searchBooks } from '../actions/books'

const mapStateToProps = ({ books }) => ({
  books: books.search,
})

const mapDispatchToProps = (dispatch) => ({
  onSearch: (term) => dispatch(searchBooks(term))
    .catch(console.warn),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BookCreator)
