import React from 'react'
import AuthenticationForm from './AuthenticationForm'
import { Link } from 'react-router-dom'

const Nav = ({ isLoggedIn }) => (
  <nav>
    <Link to='/'>
      <h1 className='typemark'>Booker</h1>
    </Link>
    <Link to='/new'>Add Book</Link>
    {!isLoggedIn && <AuthenticationForm />}
  </nav>
)

export default Nav
