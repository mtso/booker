import React, { Component } from 'react'
import BookCell from './BookCell'
import Masonry from 'masonry-layout'
import onImagesLoaded from 'imagesloaded'

const layoutGrid = (selector, i) => {
  const mountExists = !!document.querySelector(selector)
  const cellsExist = !!document.querySelector('.book-cell')
  if ((i < 20) && (!mountExists || !cellsExist)) {    
    setTimeout(() => layoutGrid(selector, ++i), 200)
    return
  }

  const mount = document.querySelector(selector)
  onImagesLoaded(mount, () => setTimeout(() => {
    const grid = new Masonry(mount, {
      itemSelector: '.book-cell',
      columnWidth: 200,
      gutter: 10,
    })
  }, 100))
}

class BookTable extends Component {
  componentDidMount() {
    const { isCreator } = this.props
    if (isCreator) {
      return
    }

    const { className } = this.props
    layoutGrid(className || '.book-table', 0)
  }
  
  render() {
    const { books, className, controls, details, isLinked } = this.props
    return (
      <div className={className || 'book-table'}>
        {books.map((book, i, books) => (
          <BookCell
            key={i}
            isLinked={isLinked}
            {...book}
          >
            { controls && controls(book, i, books) }
            { details && details(book, i, books)}
          </BookCell>
        ))}
      </div>
    )
  }
}

export default BookTable
