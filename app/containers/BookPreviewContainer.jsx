import BookPreview from '../components/BookPreview'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import { getBook } from '../actions/books'

const mapStateToProps = ({ books }) => ({
  ...books.preview,
})

const mapDispatchToProps = (dispatch, { match }) => ({
  componentDidMount: dispatch(getBook(match.params.id)),
})

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookPreview))
