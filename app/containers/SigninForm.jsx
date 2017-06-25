import React from 'react'
import { connect } from 'react-redux'
import AuthenticationForm from '../components/AuthenticationForm'
import { submitLogin } from '../actions/session'
import { Redirect, withRouter } from 'react-router-dom'

const mapStateToProps = ({ user }) => ({
  buttonTitle: 'Sign In',
  isLoggedIn: !!user.username,
})

const mapDispatchToProps = (dispatch, { history, location }) => ({
  onSubmit: (e) => {
    e.preventDefault()
    const username = e.target.elements['username'].value
    const password = e.target.elements['password'].value
    dispatch(submitLogin(username, password))
      .catch((err) => console.warn(err))
  },
})

const RedirectingSignin = ({ isLoggedIn, history, location, ...rest }) => (
  isLoggedIn ? (
    <Redirect to={location.state && location.state.from || { pathname: '/' }} />
  ) : (
    <AuthenticationForm {...rest} />
  )
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(RedirectingSignin))
