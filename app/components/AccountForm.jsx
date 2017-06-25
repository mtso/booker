import React, { Component } from 'react'

class AccountForm extends Component {
  constructor(props) {
    super(props)

    const { display_name = '', city = '', state = '' } = this.props

    const pass1 = '', pass2 = ''

    this.state = {
      display_name,
      city,
      state,
      pass1,
      pass2,
    }
  }

  onChange = (field) => (e) => {
    e.preventDefault()

    let next = {}
    next[field] = e.target.value

    this.setState(next)
  }

  onSubmitPassword = (e) => {
    e.preventDefault()
    const { pass1, pass2 } = this.state
    if (pass1 === pass2) {
      this.setState({
        pass1: '',
        pass2: '',
      }, () => this.props.onSavePassword(pass1))
    }
  }

  render() {
    const {
      onSaveProfile,
      noop,

    } = this.props
    return (
      <div>
        <h3>Edit Profile</h3>

        <form className='account-form' onSubmit={onSaveProfile || noop}>
          <label htmlFor='display-name'>Profile Name</label>
          <input onChange={this.onChange('display_name')} type='text' id='display-name' name='display-name' placeholder='Profile Name' value={this.state.display_name} />

          <label htmlFor='city'>City</label>
          <input onChange={this.onChange('city')} type='text' id='city' name='city' placeholder='City' value={this.state.city} />

          <label htmlFor='state'>State</label>
          <input onChange={this.onChange('state')} type='text' id='state' name='state' placeholder='State' value={this.state.state} />

          <input type='submit' value='Save Profile' />
        </form>

        <form className='account-form' onSubmit={this.onSubmitPassword || noop}>
          <label htmlFor='password'>New Password</label>
          <input value={this.state.pass1} onChange={this.onChange('pass1')} type='password' id='password' placeholder='New Password' />
          <input value={this.state.pass2} onChange={this.onChange('pass2')} type='password' id='password-repeat' placeholder='Repeat Password' />

          <input type='submit' value='Change Password' />
        </form>
      </div>
    )
  }
}

export default AccountForm
