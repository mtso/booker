import React from 'react'

const AccountForm = ({ city, state }) => (
  <div>
    <h3>Edit Profile</h3>

    <form>
      <label htmlFor='city'>City</label>
      <input type='text' id='city' placeholder='City' value={city || ''} />

      <label htmlFor='state'>State</label>
      <input type='text' id='state' placeholder='State' value={state || ''} />
    </form>

    <form>
      <label htmlFor='password'>New Password</label>
      <input type='password' id='password' placeholder='New Password' />
    </form>
  </div>
)

export default AccountForm
