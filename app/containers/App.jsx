import React from 'react'
import { Route } from 'react-router'
import NavContainer from './NavContainer'
import BookTableContainer from './BookBrowser'

import BookCreator from './BookCreator'
import SigninForm from './SigninForm'

import { Link } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'

const App = () => (
  <div>
    <Link to='/test'>Test</Link>
    <NavContainer />
    <Route path='/signup' component={SigninForm} />
    <PrivateRoute path='/test' render={() => (<p>in test route</p>)} />
  </div>
)

//    <Route exact path='/' render={() => <BookTableContainer />} />
//    <Route path='/new' render={() => (<BookCreator />)} />

export default App
