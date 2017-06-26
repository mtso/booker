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
      <form
        className='searchbox-container'
        onSubmit={this.onSearch || ((e) => e.preventDefault())}
      >
        <input
          className='searchbox'
          type='search'
          placeholder='Search Title'
          value={this.state.value}
          onChange={this.updateSearch}
        />
        <input
          type='submit'
          value='Search'
        />
      </form>
    )
  }
}

export default Searchbox
