import React from 'react'
import { Route } from 'react-router'
import NavContainer from './NavContainer'
import BookTableContainer from './BookBrowser'

import SigninForm from './SigninForm'

import { Link } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'
import BookBrowser from './BookBrowser'
import AccountFormContainer from './AccountFormContainer'
import BookCreatorContainer from './BookCreatorContainer'

const App = () => (
  <div>
    <NavContainer />
    <Route exact path='/' component={BookBrowser} />
    <Route path='/mybooks' component={BookBrowser} />
    <Route path='/signup' component={SigninForm} />
    <PrivateRoute path='/test' render={() => (<p>in test route</p>)} />
    <PrivateRoute path='/account' component={AccountFormContainer} />
    <PrivateRoute path='/new' component={BookCreatorContainer} />
  </div>
)

//    <Route exact path='/' render={() => <BookTableContainer />} />
//    <Route path='/new' render={() => (<BookCreator />)} />

export default App
