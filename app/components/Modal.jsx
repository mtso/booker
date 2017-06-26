import React from 'react'
import { withRouter } from 'react-router-dom'

const Modal = ({ children, isFade = true }) => (
  <div
    className={'modal-container' + (isFade ? ' fade-in' : '')}
  >
    <div className='modal-content'>
      {children}
    </div>
  </div>
)

export default Modal
