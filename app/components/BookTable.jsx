import React, { Component } from 'react'
import BookCell from './BookCell'
import Masonry from 'masonry-layout'

const layoutGrid = (i) => {
  if (!document.querySelector('.book-cell') && (i < 10)) {
    setTimeout(() => layoutGrid(++i), 100)
  }
  const mount = document.querySelector('.book-table')
  const grid = new Masonry(mount, {
    itemSelector: '.book-cell',
    columnWidth: 200,
    gutter: 10,
  })
}

class BookTable extends Component {
  componentDidMount() {
    const { isCreator } = this.props
    if (isCreator) {
      return
    }
    layoutGrid(0)
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
