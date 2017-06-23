import React from 'react'
import BookTable from '../components/BookTable'
import { connect } from 'react-redux'
import { getBooks, getMyBooks } from '../actions/books'
import { Link, NavLink, Route, withRouter } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'

const mapStateToProps = ({ username, books }) => (
  {
    isLoggedIn: !!username,
    books,
  }
)

const mapDispatchToProps = (dispatch, { match, isLoggedIn }) => ({
  componentDidMount: dispatch(match.path === '/'
      ? getBooks()
      : getMyBooks())
    .catch(console.warn)
})

const BookBrowser = ({ isLoggedIn, books, match }) => (
  <div>
    <div>
      <NavLink
        activeStyle={{fontWeight: 'bold'}}
        exact
        to='/'
      >All Books</NavLink>
      {isLoggedIn &&
        <NavLink
          activeStyle={{fontWeight: 'bold'}}
          to='/mybooks'
        >My Books</NavLink>}
      {isLoggedIn && <Link to='/new'>Add a Book</Link>}
    </div>
    <Route exact path='/' component={() => (<BookTable books={books.all} />)} />
    <PrivateRoute exact path='/mybooks' component={() => (<BookTable books={books.mybooks} />)} />
  </div>
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(BookBrowser))
