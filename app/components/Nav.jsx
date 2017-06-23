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
  // {
  //   to: '/logout',
  //   displayText: 'Log Out',
  // }
  {
    to: '/',
    exact: true,
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
      <h1 className='typemark'>Booker</h1>
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
//    <NavLink to='/about' activeStyle={{fontWeight: 'bold'}}>About</NavLink>
//    <NavLink to='/signin' activeStyle={{fontWeight: 'bold'}}>Sign In</NavLink>

    // <Link to='/new'>Add Book</Link>
    // {!isLoggedIn && <AuthenticationForm />}

export default Nav
