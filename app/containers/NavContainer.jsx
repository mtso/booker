import { withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import Nav from '../components/Nav'
import { submitLogout } from '../actions/session'

const mapStateToProps = ({ user }) => ({
  isLoggedIn: !!user.username,
})

const mapDispatchToProps = (dispatch, { history }) => ({
  onLogout: () => dispatch(submitLogout())
    .then(() => history.push('/'))
    .catch((err) => console.warn(err)),
})

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(Nav))
