import React from 'react'
import { Route } from 'react-router'
import Nav from '../components/Nav'
import BookTableContainer from './BookTableContainer'

import BookCreator from '../components/BookCreator'

const App = () => (
  <div>
    <Nav />
    <Route exact path='/' render={() => <BookTableContainer />} />
    <Route path='/new' render={() => (<BookCreator />)} />
  </div>
)

export default App
