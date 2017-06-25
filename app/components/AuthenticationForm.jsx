import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import Modal from './Modal'

const defaultAction = (e) => {
  e.preventDefault()
}

const AuthenticationForm = ({ onSubmit, isSignup, location }) => (
  <Modal isFade={
    (location.pathname === '/signin' || location.pathname === '/signup')
      ? false : true
  }>
    <div className='auth-container'>
      <Link to='/'>
        <img className='logo' src='/static/img/book-logo.svg' />
      </Link>

      { isSignup ? (
        <Link className='float-right' to='/signin'>Or Sign In.</Link>
      ) : (
        <Link className='float-right' to='/signup'>Or Sign Up.</Link>
      ) }
      <form onSubmit={onSubmit || defaultAction}>
        <input
          type='text'
          name='username'
          placeholder='Username'
        />
        <input
          type='password'
          name='password'
          placeholder='Password'
        />
        { isSignup && (
          <input
            type='password'
            name='password-repeat'
            placeholder='Repeat Password'
          />
        ) }
        <input type='submit' value={
          isSignup ? 'Sign Up' : 'Sign In'
        } />
      </form>
    </div>
  </Modal>
)

export default AuthenticationForm
