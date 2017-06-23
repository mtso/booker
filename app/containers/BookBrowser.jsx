import React from 'react'
import BookTable from '../components/BookTable'
import { connect } from 'react-redux'
import { getBooks, getMyBooks } from '../actions/books'
import { postTrade } from '../actions'
import { Link, NavLink, Route, withRouter } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'

const mapStateToProps = ({ username, books }) => (
  {
    isLoggedIn: !!username,
    books,
    username,
  }
)

const mapDispatchToProps = (dispatch, { match, isLoggedIn }) => ({
  componentDidMount: dispatch(match.path === '/'
      ? getBooks()
      : getMyBooks())
    .catch(console.warn),
  onTrade: (bookid) => () => dispatch(postTrade(bookid))
    .catch(console.warn),
})

const BookBrowser = ({ isLoggedIn, username, books, match, onTrade }) => (
  <div>
    <div className='tab-container'>
      <NavLink
        className='tab-link'
        activeClassName='active'
        exact
        to='/'
      >All Books</NavLink>
      {isLoggedIn &&
        <NavLink
          className='tab-link'
          activeClassName='active'
          to='/mybooks'
        >My Books</NavLink>}
      {isLoggedIn && 
        <Link
          to='/new'
        >Add a Book</Link>}
    </div>
    <Route exact path='/' component={() => (
      <BookTable
        books={books.all}
        controls={(book) => {
          if (book.username !== username) {
            return (
              <button onClick={
                onTrade(book.id)
              }>Request Trade</button>
            )
          }
        }}
      />
    )} />
    <PrivateRoute exact path='/mybooks' component={() => (
      <BookTable books={books.mybooks} />
    )} />
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookBrowser))
