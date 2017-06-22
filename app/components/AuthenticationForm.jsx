import React, { Component } from 'react'

const defaultAction = (e) => {
  e.preventDefault()
}

const AuthenticationForm = ({ buttonTitle, onSubmit }) => (
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
    <input type='submit' value={buttonTitle || 'Submit'} />
  </form>
)

export default AuthenticationForm
