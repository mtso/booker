import React from 'react'
import { connect } from 'react-redux'
import AuthenticationForm from '../components/AuthenticationForm'
import { submitLogin, submitSignup } from '../actions'
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
    
    if (location.pathname === '/signup') {
      const password2 = e.target.elements['password-repeat'].value
      if (password2 !== password) {
        console.warn('mismatching passwords')
        return
      }
      console.log(username, password)
      dispatch(submitSignup(username, password))
        .catch(console.warn)
    } else {
      dispatch(submitLogin(username, password))
        .catch(console.warn)
    }    
  },
})

const RedirectingSignin = ({ isLoggedIn, history, location, ...rest }) => (
  isLoggedIn ? (
    <Redirect to={location.state && location.state.from || { pathname: '/' }} />
  ) : (
    <AuthenticationForm
      {
        ...{
          history,
          location,
          isLoggedIn,
        }
      }
      {...rest}
      isSignup={location.pathname === '/signup'}
    />
  )
)

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(RedirectingSignin))
