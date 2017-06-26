import React from 'react'
import BookTable from '../components/BookTable'
import { connect } from 'react-redux'
import { getBooks, getMyBooks } from '../actions/books'
import { postTrade } from '../actions'
import { Link, NavLink, Route, withRouter } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'

const mapStateToProps = ({ user, books }) => (
  {
    isLoggedIn: !!user.username,
    books,
    user,
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


const BookBrowser = ({ isLoggedIn, user, books, match, onTrade, history }) => (
  <div className='page-container'>
    <div className='page-header'>
      <h2 className='content-title'>
        Booker
      </h2>
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
            className='action-button add-button'
          >Add a Book</Link>}
      </div>
    </div>
    <div className='content-container'>
      <Route exact path='/' render={() => (
        <BookTable
          isLinked={true}
          books={books.all}
          controls={(book) => {
            if (book.username !== user.username) {
              return (
                <div className='cell-button-container cell-topright'>
                  <button className='cell-button' onClick={
                    () => history.push('/book/'+book.id)
                  }>Tradeable</button>
                </div>
              )
            }
          }}
        />
      )} />
      <PrivateRoute exact path='/mybooks' render={() => (
        <BookTable
          books={books.mybooks}
          isLinked={true}
        />
      )} />
    </div>
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookBrowser))
