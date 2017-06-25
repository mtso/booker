import React from 'react'
import { Link } from 'react-router-dom'

const placeholder = 'https://placem.at/people?h=400&w=300'

const BookCell = ({ id, image_url, children, status, onTrade }) => (
  <div className='book-cell'>
    <Link to={'/book/'+id}>
      <img style={{width: '200px'}} src={image_url || placeholder} />
    </Link>
    { children }
  </div>
)

export default BookCell
