import React from 'react'
import Modal from './Modal'

const BookPreview = ({ history, id, title, image_url, owner }) => {
  if (!owner) {
    return (
      <Modal>
        <button onClick={() => {
          if (history.length < 3) {
            history.push('/')
          } else {
            history.goBack()
          }
        }}>Close</button>
      </Modal>
    )
  }
  const { display_name } = owner
  return (
    <Modal>
      <button onClick={() => {
        if (history.length < 3) {
          history.push('/')
        } else {
          history.goBack()
        }
      }}>Close</button>
      <br />
      <img src={image_url} />
      <br />
      { title }
      <br />
      { display_name }
    </Modal>
  )
}

export default BookPreview
