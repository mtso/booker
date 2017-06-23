import React from 'react'
import Modal from './Modal'

const BookPreview = ({ history, id, title, image_url, username }) => (
  <Modal>
    <button onClick={() => history.goBack()}>Close</button>
    <br />
    <img src={image_url} />
    <br />
    { id }
    <br />
    { title }
    <br />
    { username }
  </Modal>
)

export default BookPreview
