import React from 'react'
import AuthenticationForm from './AuthenticationForm'
import { Link } from 'react-router-dom'

const Nav = ({ isLoggedIn }) => (
  <nav>
    <h1 className='typemark'>Booker</h1>
    {!isLoggedIn && <AuthenticationForm />}
  </nav>
)

export default Nav
