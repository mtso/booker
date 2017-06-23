import React from 'react'

const Modal = ({ children }) => (
  <div className='modal-container fade-in'>
    <div className='modal-content'>
      {children}
    </div>
  </div>
)

export default Modal
