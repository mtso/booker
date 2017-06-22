import React from 'react'
import BookTable from '../components/BookTable'
import { connect } from 'react-redux'

const mapStateToProps = (state) => (
  {
    books: [{}, {}, {}],
  }
)

export default connect(
  mapStateToProps,
  null
)(BookTable)
