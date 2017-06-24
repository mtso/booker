import React, { Component } from 'react'

const defaultAction = (e) => {
  e.preventDefault()
}

const AuthenticationForm = ({ buttonTitle, onSubmit, isSignup }) => (
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
    <input type='submit' value={buttonTitle || 'Submit'} />
  </form>
)

export default AuthenticationForm
