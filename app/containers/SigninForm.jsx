import { connect } from 'react-redux'
import AuthenticationForm from '../components/AuthenticationForm'
import { submitLogin } from '../actions/session'
import { withRouter } from 'react-router-dom'

const mapStateToProps = () => ({
  buttonTitle: 'Sign In',
})

const mapDispatchToProps = (dispatch, { history }) => ({
  onSubmit: (e) => {
    e.preventDefault()
    const username = e.target.elements['username'].value
    const password = e.target.elements['password'].value
    dispatch(submitLogin(username, password))
      .then(() => history.push('/'))
      .catch((err) => console.warn(err))
  },
})

export default withRouter(connect(
  mapStateToProps,
  mapDispatchToProps
)(AuthenticationForm))
