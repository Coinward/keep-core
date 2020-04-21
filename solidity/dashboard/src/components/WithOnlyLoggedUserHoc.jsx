import React, { useContext } from "react"
import { Web3Context } from "../components/WithWeb3Context"

export const withOnlyLoggedUser = (WrapperedComponent) => (props) => {
  const { yourAddress } = useContext(Web3Context)

  return yourAddress ? <WrapperedComponent {...props} /> : null
}
