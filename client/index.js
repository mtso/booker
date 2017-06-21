import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import App from '../app'
import { Provider } from 'react-redux'
import { createStore } from 'redux'

const store = createStore((state = {}) => state)

render(
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>,
  document.querySelector("#app")
)
