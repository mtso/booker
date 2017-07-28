import React from 'react'

const AboutPage = () => (
  <div className='page-container'>
    <div className='page-header'>
      <h2 className='content-title'>
        About Booker
      </h2>
    </div>
    <div className='content-container'>
      <p>
      Booker is a book collecting and trading network.
      </p>
      <p>
      Booker lets you search for books that you own and add it to your own
      collection. You can also browse the books from other users and request to
      trade for them. The Trade tab shows your incoming and outgoing book trades.
      You may accept incoming trades or cancel outgoing trades. Furthermore,
      view book details directly from the Browse or Trade tab by clicking on the
      book cover.
      </p>
      <p>
      Designed and built by <a href='https://github.com/mtso'>Matthew Tso</a>
      with Go, React, Redux, and Postgres (<a href='https://github.com/mtso/booker'>View Source</a>).
      </p>
    </div>
  </div>
)

export default AboutPage
