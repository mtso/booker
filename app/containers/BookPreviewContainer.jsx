import BookPreview from '../components/BookPreview'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import { getBook, postTradeFromPreview, cancelTradeFromPreview } from '../actions'

const mapStateToProps = ({ books, user }) => ({
  user,
  ...books.preview,
})

const mapDispatchToProps = (dispatch, { match, history }) => ({
  componentDidMount: dispatch(getBook(match.params.id)),
  onTrade: (bookid) => () => dispatch(postTradeFromPreview(bookid))
    .then(() => dispatch(getBook(bookid)))
    .catch(console.warn),
  onCancel: (id) => (bookid) => () => dispatch(cancelTradeFromPreview(id))
    .then(() => dispatch(getBook(bookid)))
    .catch(console.warn),
  onUnauthedTrade: (bookid) => () => history.push('/signup', { from: '/book/'+bookid }),
})

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookPreview))
