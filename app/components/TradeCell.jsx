import React from 'react'
import { Link } from 'react-router-dom'

const renderLocation = (city, state) => {
  let location
  if (city !== '') {
    location = city
    if (state !== '') {
      location += ', ' + state
    }
  } else if (state !== '') {
    location = state
  }
  return location
}

const TradeCell = ({ id, book_id, user_id, status, book, user, children, isIncoming }) => (
  <div className='trade-cell'>
    <Link to={'/book/'+book_id}>
      <img style={{width: '200px'}} src={book.image_url} />
    </Link>
    <div className='trade-detail'>
      <h2>{ (!!isIncoming) ? (
        'Requested By'
      ) : (
        'Request To'
      )
      }</h2>
      <h4>{ (user.display_name || user.username) 
        ? 'Name: ' + (user.display_name || user.username) 
        : '' }</h4>
      <h4>{ !!renderLocation(user.city, user.state)
        ? 'Location: ' + renderLocation(user.city, user.state)
        : '' }</h4>
      { children }
    </div>
    
  </div>
)

export default TradeCell
