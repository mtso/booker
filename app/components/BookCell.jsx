import React from 'react'
import { Link } from 'react-router-dom'

const placeholder = 'https://placem.at/people?h=400&w=300'

const BookCell = ({ id, image_url, children, status, onTrade, isLinked }) => (
  <div className='book-cell'>
    { isLinked ? (
        <Link to={'/book/'+id}>
          <img style={{width: '200px'}} src={image_url || placeholder} />
        </Link>
      ) : (
        <img style={{width: '200px'}} src={image_url || placeholder} />
      ) }
    { children }
  </div>
)

export default BookCell
