import React from "react"
import Loader from "react-loader-spinner"

const LoadingIndicator = ({ loading }) => {
  return (
    loading && (
      <Loader type="ThreeDots" color="#000000" height="50" width="50" />
    )
  )
}

export default LoadingIndicator
