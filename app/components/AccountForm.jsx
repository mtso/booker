import React from 'react'

const AccountForm = ({ city, state, display_name }) => (
  <div>
    <h3>Edit Profile</h3>

    <form className='account-form'>
      <label htmlFor='display-name'>Profile Name</label>
      <input type='text' id='display-name' placeholder='Profile Name' value={display_name || ''} />

      <label htmlFor='city'>City</label>
      <input type='text' id='city' placeholder='City' value={city || ''} />

      <label htmlFor='state'>State</label>
      <input type='text' id='state' placeholder='State' value={state || ''} />

      <input type='submit' value='Save Profile' />
    </form>

    <form className='account-form'>
      <label htmlFor='password'>New Password</label>
      <input type='password' id='password' placeholder='New Password' />

      <input type='submit' value='Change Password' />
    </form>
  </div>
)

export default AccountForm
