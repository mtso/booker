import React from 'react'
import Modal from './Modal'

const BookPreview = ({ history, id, title, image_url, display_name }) => (
  <Modal>
    <button onClick={() => history.goBack()}>Close</button>
    <br />
    <img src={image_url} />
    <br />
    { title }
    <br />
    { display_name }
  </Modal>
)

export default BookPreview
