import React, { Component } from 'react'

class Searchbox extends Component {
  state = {
    value: '',
  }

  onSearch = (e) => {
    e.preventDefault()
    this.props.onSearch(this.state.value)
  }

  updateSearch = (e) => {
    e.preventDefault()
    const { value } = e.target
    this.setState({
      value,
    })
  }

  render() {
    return (
      <form onSubmit={this.onSearch || ((e) => e.preventDefault())}>
        <input
          type='search'
          placeholder='Search Title'
          value={this.state.value}
          onChange={this.updateSearch}
        />
        <button
          type='submit'
        >
          Search
        </button>
      </form>
    )
  }
}

// const Searchbox = ({ onSubmit, ...props }) => (
//   <form onSubmit={onSubmit || ((e) => e.preventDefault())}>
//     <input
//       type='search'
//       placeholder='Search Title'
//       {...props}
//     />
//     <button
//       type='submit'
//     >
//       Search
//     </button>
//   </form>
// )

export default Searchbox
