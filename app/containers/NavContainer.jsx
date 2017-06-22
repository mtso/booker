import { withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import Nav from '../components/Nav'

const mapStateToProps = ({ username }) => ({
  isLoggedIn: username !== null,
})

export default withRouter(connect(
  mapStateToProps,
  null
)(Nav))
