import React, { useContext } from 'react'
import DelegatedTokens from '../components/DelegatedTokens'
import PendingUndelegation from '../components/PendingUndelegation'
import SlashedTokens from '../components/SlashedTokens'
import { useSubscribeToContractEvent } from '../hooks/useSubscribeToContractEvent'
import { TOKEN_STAKING_CONTRACT_NAME } from '../constants/constants'
import { Web3Context } from '../components/WithWeb3Context'

const OperatorPage = (props) => {
  const { yourAddress } = useContext(Web3Context)
  const { latestEvent } =
    useSubscribeToContractEvent(TOKEN_STAKING_CONTRACT_NAME, 'InitiatedUnstake', { filter: { operator: yourAddress } })

  return (
    <>
      <h3>My Token Operations</h3>
      <DelegatedTokens latestUnstakeEvent={latestEvent} />
      <PendingUndelegation latestUnstakeEvent={latestEvent} />
      {/* TODO fetching slashed info form the contract */}
      {/* <SlashedTokens /> */}
    </>

  )
}

export default OperatorPage
