import React from 'react'
import AuthenticationForm from './AuthenticationForm'
import { NavLink, Link, withRouter } from 'react-router-dom'

const unAuthenticatedLinks = [
  {
    to: '/signup',
    activeStyle: {fontWeight: 'bold'},
    displayText: 'Sign Up',
  },
  {
    to: '/about',
    activeStyle: {fontWeight: 'bold'},
    displayText: 'About',    
  },
]

const authenticatedLinks = [
  {
    to: '/',
    isActive: (_, { pathname }) => pathname === '/' || pathname === '/mybooks',
    activeStyle: {fontWeight: 'bold'},
    displayText: 'Books',
  },
  {
    to: '/account',
    activeStyle: {fontWeight: 'bold'},
    displayText: 'Account',
  },
]

const Nav = ({ isLoggedIn, onLogout }) => (
  <nav>
    <Link to='/'>
      <img className='logo' src='/img/book-logo.svg' />
    </Link>
    {
      isLoggedIn
      ? authenticatedLinks.map(({ displayText, ...props }, i) => (
        <NavLink {...props} key={i}>{displayText}</NavLink>
      ))
      : unAuthenticatedLinks.map(({ displayText, ...props }, i) => (
        <NavLink {...props} key={i}>{displayText}</NavLink>
      ))
    }
    {
      isLoggedIn && <button onClick={onLogout}>Log Out</button>// <Link onClick={onLogout}>Log Out</Link>
    }
  </nav>
)

export default Nav
