import AccountForm from '../components/AccountForm'
import { connect } from 'react-redux'

const mapStateToProps = ({ user }) => ({ ...user })

export default connect(
  mapStateToProps
)(AccountForm)
