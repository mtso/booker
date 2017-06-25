import React from 'react'

const Modal = ({ children, isFade = true }) => (
  <div className={'modal-container' + (isFade ? ' fade-in' : '')}>
    <div className='modal-content'>
      {children}
    </div>
  </div>
)

export default Modal
