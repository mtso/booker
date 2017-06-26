import React from 'react'
import Modal from './Modal'
import { STATUS_REQUESTED } from '../tradestatus'

const closeModal = (history) => () => {
  if (history.length < 3) {
    history.push('/')
  } else {
    history.goBack()
  }
}

const BookPreview = ({ user, history, id, title, image_url, owner, trade,
  onTrade, onCancel, onUnauthedTrade }) => {
  // Prevent render if data does not exist
  if (!owner) {
    return (
      <Modal>
        <button onClick={closeModal(history)}>Close</button>
      </Modal>
    )
  }

  const isLoggedIn = !!user.username
  const { username, display_name, city, state } = owner

  let location
  if (city !== '') {
    location = city
    if (state !== '') {
      location += ', ' + state
    }
  } else if (state !== '') {
    location = state
  }

  let tradeControl
  const { status } = trade
  if (status && status === STATUS_REQUESTED) {
    tradeControl = <button className='action-button' onClick={onCancel(trade.id)(id)}>Cancel Trade</button>
  } else if (user.id && owner.id === user.id) {
    tradeControl = null
  } else {
    tradeControl = <button className='action-button' onClick={
      isLoggedIn ? onTrade(id) : onUnauthedTrade(id)
    }>Request Trade</button>
  }

  return (
    <Modal>
      <div className='bookcreator-container'>
        <div className='bookcreator-header'>
          <button className='float-right' onClick={closeModal(history)}>Close</button>
        </div>

        <img src={image_url} />
        <div className='book-detail'>
          <h2>{ title }</h2>
          <h4>{ (display_name || username) && 'Owner: ' + (display_name || username) }</h4>
          <h4>{ location && 'Location: '+location }</h4>
          { tradeControl && tradeControl }
        </div>
      </div>
    </Modal>
  )
}

export default BookPreview
