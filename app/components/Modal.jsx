import React from 'react'
import { withRouter } from 'react-router-dom'

// class Modal extends React.Component {
//   componentDidMount() {
//     this.refs.container.addEventListener('click', () => {
//       const { history, location } = this.props
//       if (location.state) {
//         history.push(location.state.from)
//       } else {
//         history.push('/')
//       }
//     })
//   }
//   render() {
//     let { history, location, children, isFade } = this.props
//     if (isFade === undefined) {
//       isFade = true
//     }
//     return (
//       <div
//         ref='container'
//         className={'modal-container' + (isFade ? ' fade-in' : '')}
//       >
//         <div className='modal-content'>
//           {children}
//         </div>
//       </div>
//     )
//   }
// }

const Modal = ({ children, isFade = true }) => (
  <div
    className={'modal-container' + (isFade ? ' fade-in' : '')}
  >
    <div className='modal-content'>
      {children}
    </div>
  </div>
)

export default Modal
