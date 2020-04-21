import React, { useEffect, useContext } from "react"
import { useFetchData } from "../hooks/useFetchData"
import { operatorService } from "../services/token-staking.service"
import {
  displayAmount,
  isSameEthAddress,
  isEmptyObj,
  formatDate,
} from "../utils/general.utils"
import { LoadingOverlay } from "./Loadable"
import { Web3Context } from "./WithWeb3Context"
import StatusBadge, { BADGE_STATUS } from "./StatusBadge"
import { DataTable, Column } from "./DataTable"
import { PENDING_STATUS } from "../constants/constants"
import moment from "moment"

const initialData = { pendinUndelegations: [] }

const PendingUndelegation = ({ latestUnstakeEvent }) => {
  const { stakingContract, yourAddress } = useContext(Web3Context)
  const [state, setData] = useFetchData(
    operatorService.fetchPendingUndelegation,
    initialData
  )
  const { isFetching, data } = state
  const { undelegationCompletedAt, undelegationPeriod } = data

  useEffect(() => {
    const {
      returnValues: { operator, undelegatedAt },
    } = latestUnstakeEvent
    if (!isEmptyObj(latestUnstakeEvent)) {
      if (!isSameEthAddress(yourAddress, operator)) {
        return
      }
      const undelegationCompletedAt = moment
        .unix(undelegatedAt)
        .add(undelegationPeriod, "seconds")
      stakingContract.methods
        .getDelegationInfo(operator)
        .call()
        .then((data) => {
          const { amount } = data
          setData({
            ...state.data,
            undelegationCompletedAt,
            pendingUnstakeBalance: amount,
            undelegationStatus: "PENDING",
          })
        })
    }
  }, [latestUnstakeEvent.transactionHash])

  return (
    <LoadingOverlay isFetching={isFetching}>
      <section id="pending-undelegation" className="tile">
        <h3 className="text-grey-60">Token Undelegation</h3>
        <DataTable data={[data]} itemFieldId="undelegationComplete">
          <Column
            header="amount"
            field="pendingUnstakeBalance"
            renderContent={({ pendingUnstakeBalance }) =>
              pendingUnstakeBalance && `${displayAmount(pendingUnstakeBalance)}`
            }
          />
          <Column
            header="status"
            field="undelegationStatus"
            renderContent={({ undelegationStatus }) => {
              const statusText =
                undelegationStatus === PENDING_STATUS
                  ? `${undelegationStatus.toLowerCase()}, ${undelegationCompletedAt.fromNow(
                      true
                    )}`
                  : undelegationStatus

              return (
                undelegationStatus && (
                  <StatusBadge
                    className="self-start"
                    status={BADGE_STATUS[undelegationStatus]}
                    text={statusText.toLowerCase()}
                  />
                )
              )
            }}
          />
          <Column
            header="estimate"
            field="undelegationComplete"
            renderContent={({ undelegationCompletedAt }) =>
              undelegationCompletedAt
                ? formatDate(undelegationCompletedAt)
                : "-"
            }
          />
          <Column
            header="undelegation period"
            field="undelegationPeriod"
            renderContent={({ undelegationPeriod }) => {
              const undelegationPeriodRelativeTime = moment()
                .add(undelegationPeriod, "seconds")
                .fromNow(true)
              return undelegationPeriodRelativeTime
            }}
          />
        </DataTable>
      </section>
    </LoadingOverlay>
  )
}

export default PendingUndelegation
