import React from 'react'
import AuthenticationForm from './AuthenticationForm'
import { NavLink, Link, withRouter } from 'react-router-dom'

const unAuthenticatedLinks = [
  {
    to: '/signup',
    activeClassName: 'active',
    displayText: 'Sign Up',
  },
  {
    to: '/about',
    activeClassName: 'active',
    displayText: 'About',    
  },
]

const authenticatedLinks = [
  {
    to: '/',
    isActive: (_, { pathname }) => pathname === '/' || pathname === '/mybooks',
    activeClassName: 'active',
    displayText: 'Books',
  },
  {
    to: '/trades',
    activeClassName: 'active',
    displayText: 'Trades',
  },
  {
    to: '/account',
    activeClassName: 'active',
    displayText: 'Account',
  },
]

const Nav = ({ isLoggedIn, onLogout }) => (
  <div className='nav'>
    <Link to='/'>
      <img className='logo' src='/static/img/book-logo.svg' />
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
      isLoggedIn && <button onClick={onLogout}>Log Out</button>
    }
  </div>
)

export default Nav
