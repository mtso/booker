import React from 'react'

const placeholder = 'https://placem.at/people?h=400&w=300'

const BookCell = ({ image_url, children }) => (
  <div>
    <img style={{width: '200px'}} src={image_url || placeholder} />
    { children }
  </div>
)

export default BookCell
