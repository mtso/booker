import React from 'react'
import store from '../store'
import { Route } from 'react-router'
import { Redirect, withRouter } from 'react-router-dom'
import { connect } from 'react-redux'

const PrivateRoute = ({ isLoggedIn, render, component: Component, ...rest }) => {  
  if (!!render) {
    // Return Route with "render" rendered
    return (
      <Route {...rest} render={ isLoggedIn ? (
        render
      ) : (
        ({ location }) => (<Redirect to={{
            pathname: '/signup',
            state: { from: location },
          }} />
        )
      )}/>
    )
  } else {
    // Return Route with "Component" rendered
    return (
      <Route {...rest} render={(props) => (
        isLoggedIn ? (
          <Component {...props} />
        ) : (
          <Redirect to={{
            pathname: '/signup',
            state: { from: props.location },
          }} />
        )
      )}/>
    )
  }
}

const mapStateToProps = ({ user }) => ({
  isLoggedIn: !!user.username,
})

export default withRouter(connect(
  mapStateToProps,
  null
)(PrivateRoute))
