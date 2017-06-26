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
      <div className='fixed-content'>
        <h2 className='content-title'>
          Account
        </h2>
        
        <h3>Edit Profile</h3>
        <form className='account-form' onSubmit={onSaveProfile || noop}>
          <input
            onChange={this.onChange('display_name')}
            type='text'
            id='display-name'
            name='display-name'
            placeholder='Profile Name'
            value={this.state.display_name || ''}
          />

          <input
            onChange={this.onChange('city')}
            type='text'
            id='city'
            name='city'
            placeholder='City'
            value={this.state.city || ''}
          />

          <input
            onChange={this.onChange('state')}
            type='text'
            id='state'
            name='state'
            placeholder='State'
            value={this.state.state || ''}
          />

          <input className='primary-button' type='submit' value='Save Profile' />
        </form>

        <h3>Change Password</h3>
        <form className='account-form' onSubmit={this.onSubmitPassword || noop}>
          <input value={this.state.pass1} onChange={this.onChange('pass1')} type='password' id='password' placeholder='New Password' />
          <input value={this.state.pass2} onChange={this.onChange('pass2')} type='password' id='password-repeat' placeholder='Repeat Password' />

          <input className='primary-button' type='submit' value='Change Password' />
        </form>
      </div>
    )
  }
}

export default AccountForm
