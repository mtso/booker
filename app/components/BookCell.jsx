import React from 'react'

const placeholder = 'https://placem.at/people?h=400&w=300'

const BookCell = ({ image_url }) => (
  <div>
    <img style={{width: '200px'}} src={image_url || placeholder} />
  </div>
)

export default BookCell
