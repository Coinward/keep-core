import React from 'react'

const Loadable = ({ text, isFetching, children }) => (
  isFetching ? <div>{text}</div> : children
)

Loadable.defaultProps = {
  text: 'Loading...',
  isFetching: true,
}

export const ClockIndicator = (props) => (<div className='indicator-clock'/>)

export const LoadingOverlay = React.memo(({ isFetching, classNames, children }) => {
  if (!isFetching) {
    return children
  }

  return (
    <div className={`loading-overlay-container ${classNames}`}>
      {children}
      <div className='loading-overlay'>
        Loading...
      </div>
    </div>
  )
})

export default Loadable
