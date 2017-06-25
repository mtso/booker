import AccountForm from '../components/AccountForm'
import { connect } from 'react-redux'
import { saveProfile, receiveUser } from '../actions'

const mapStateToProps = ({ user }) => ({ ...user })

const mapDispatchToProps = (dispatch) => ({
  onSaveProfile: (e) => {
    e.preventDefault()

    const els = e.target.elements
    const profile = {
      display_name: els['display-name'].value,
      city: els['city'].value,
      state: els['state'].value,
    }

    dispatch(saveProfile(profile))
      .then(() => dispatch(receiveUser(profile)))
  },
  onSavePassword: (password) => dispatch(saveProfile({ password })),
  noop: () => _,
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AccountForm)
