import BookCreator from '../components/BookCreator'
import { connect } from 'react-redux'

const mapStateToProps = ({ books }) => ({
  books: books.search,
})

export default connect(
  mapStateToProps
)(BookCreator)
