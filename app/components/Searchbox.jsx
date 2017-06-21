import React from 'react'

const Searchbox = ({ onSubmit, ...props }) => (
  <form onSubmit={onSubmit || ((e) => e.preventDefault())}>
    <input
      type='search'
      placeholder='Search Title'
      {...props}
    />
    <button
      type='submit'
    >
      Search
    </button>
  </form>
)

export default Searchbox
