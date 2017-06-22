import React from 'react'
import { Route } from 'react-router'
import NavContainer from './NavContainer'
import BookTableContainer from './BookBrowser'

import BookCreator from './BookCreator'
import SigninForm from './SigninForm'

const App = () => (
  <div>
    <NavContainer />
    <Route path='/signup' component={SigninForm} />
  </div>
)
//    <Route exact path='/' render={() => <BookTableContainer />} />
//    <Route path='/new' render={() => (<BookCreator />)} />

export default App
