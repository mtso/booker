import React, { Component } from 'react'
import BookCell from './BookCell'
import Masonry from 'masonry-layout'

// class BookTable extends Component {
//   componentDidMount(prevProps, prevState) {
//     const mount = document.querySelector('.book-table')
//     console.log(mount)
//     const grid = new Masonry(mount, {
//       itemSelector: '.book-cell',
//       columnWidth: 200,
//       gutter: 10,
//     })
//   }
//   render() {
//     const { books, className, controls } = this.props
//     return (
//       <div className={className || 'book-table'}>
//         {books.map((book, i, books) => (
//           <BookCell
//             key={i}
//             {...book}
//           >
//             {controls && controls(book, i, books)}
//           </BookCell>
//         ))}
//       </div>
//     )
//   }
// }

const BookTable = ({ books, className, controls }) => (
  <div className={className || 'book-table'}>
    {books.map((book, i, books) => (
      <BookCell
        key={i}
        {...book}
      >
        {controls && controls(book, i, books)}
      </BookCell>
    ))}
  </div>
)

export default BookTable
