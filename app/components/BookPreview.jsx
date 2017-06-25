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
  const { display_name, city, state } = owner

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
    tradeControl = <button onClick={onCancel(trade.id)(id)}>Cancel Trade</button>
  } else if (user.id && owner.id === user.id) {
    tradeControl = null
  } else {
    console.log(isLoggedIn)
    tradeControl = <button onClick={
      isLoggedIn ? onTrade(id) : onUnauthedTrade(id)
    }>Request Trade</button>
  }

  return (
    <Modal>
      <button onClick={closeModal(history)}>Close</button>
      <br />
      <img src={image_url} />
      <br />
      { title }
      <br />
      Owner: { display_name }
      <br />
      { location && 'Location: '+location }
      <br />
      { tradeControl && tradeControl }
    </Modal>
  )
}

export default BookPreview
