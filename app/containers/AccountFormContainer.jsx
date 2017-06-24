import AccountForm from '../components/AccountForm'
import { connect } from 'react-redux'
import { saveProfile } from '../actions'

const mapStateToProps = ({ user }) => ({ ...user })

const mapDispatchToProps = (dispatch) => ({
  onSaveProfile: (e) => {
    e.preventDefault()

    const els = e.target.elements
    dispatch(saveProfile({
      display_name: els['display-name'].value,
      city: els['city'].value,
      state: els['state'].value,
    }))
  },
  onSavePassword: (password) => dispatch(saveProfile({ password })),
  noop: () => _,
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AccountForm)
