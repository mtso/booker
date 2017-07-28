import React from 'react'
import { Route } from 'react-router'
import NavContainer from './NavContainer'
import BookTableContainer from './BookBrowser'

import SigninForm from './SigninForm'

import { Link, Redirect } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'
import BookBrowser from './BookBrowser'
import Trades from './Trades'
import AccountFormContainer from './AccountFormContainer'
import BookCreatorContainer from './BookCreatorContainer'
import BookPreviewContainer from './BookPreviewContainer'
import AboutPage from '../components/AboutPage'

const App = () => (
  <div className='container'>
    <NavContainer />
    <div className='page'>
      <Route path='/book/:id' component={BookPreviewContainer} />
      <Route exact path='/' component={BookBrowser} />
      <Route path='/mybooks' component={BookBrowser} />
      <Route path='/signup' component={SigninForm} />
      <Route path='/signin' component={SigninForm} />
      <Route path='/about' component={AboutPage} />
      <PrivateRoute path='/account' component={AccountFormContainer} />
      <PrivateRoute path='/new' component={BookCreatorContainer} />

      <PrivateRoute exact path='/trades' render={() => (<Redirect to='/trades/incoming' />)} />
      <PrivateRoute path='/trades/:type' component={Trades} />
    </div>
  </div>
)

export default App
