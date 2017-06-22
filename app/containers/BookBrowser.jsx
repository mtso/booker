import React from 'react'
import BookTable from '../components/BookTable'
import { connect } from 'react-redux'
import { getBooks } from '../actions/books'

const mapStateToProps = (state) => (
  {
    books: state.books.books, // [{}, {}, {}],
  }
)

const mapDispatchToProps = (dispatch) => ({
  componentDidMount: dispatch(getBooks()),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BookTable)
