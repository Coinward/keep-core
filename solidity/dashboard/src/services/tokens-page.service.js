import { contractService } from './contracts.service'
import web3Utils from 'web3-utils'
import {
  TOKEN_STAKING_CONTRACT_NAME,
  TOKEN_GRANT_CONTRACT_NAME,
  KEEP_TOKEN_CONTRACT_NAME,
} from '../constants/constants'
import { sub, gt } from '../utils/arithmetics.utils'
import moment from 'moment'

export const fetchTokensPageData = async (web3Context) => {
  const { yourAddress } = web3Context

  const [
    keepTokenBalance,
    grantTokenBalance,
    minimumStake,
    operatorsAddresses,
    undelegationPeriod,
    initializationPeriod,
  ] = await Promise.all([
    contractService.makeCall(web3Context, KEEP_TOKEN_CONTRACT_NAME, 'balanceOf', yourAddress),
    contractService.makeCall(web3Context, TOKEN_GRANT_CONTRACT_NAME, 'balanceOf', yourAddress),
    contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'minimumStake'),
    contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'operatorsOf', yourAddress),
    contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'undelegationPeriod'),
    contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'initializationPeriod'),
  ])
  const operatorsAddressesSet = new Set(operatorsAddresses)
  const granteeOperators = await contractService.makeCall(web3Context, TOKEN_GRANT_CONTRACT_NAME, 'getGranteeOperators', yourAddress)
  const granteeOperatorsSet = new Set(granteeOperators)

  const [
    ownedDelegations,
    ownedUndelegations,
    tokenStakingBalance,
    pendingUndelegationBalance,
  ] = await getDelegations(operatorsAddressesSet, web3Context, initializationPeriod, undelegationPeriod)


  const [
    granteeDelegations,
    granteeUndelegations,
  ] = await getDelegations(granteeOperatorsSet, web3Context, initializationPeriod, undelegationPeriod, true)


  const delegations = [...ownedDelegations, ...granteeDelegations].sort((a, b) => sub(b.createdAt, a.createdAt))
  const undelegations = [...ownedUndelegations, ...granteeUndelegations].sort((a, b) => sub(b.undelegatedAt, a.undelegatedAt))

  return {
    delegations,
    undelegations,
    keepTokenBalance,
    grantTokenBalance,
    ownedTokensDelegationsBalance: tokenStakingBalance.toString(),
    ownedTokensUndelegationsBalance: pendingUndelegationBalance.toString(),
    minimumStake,
    initializationPeriod,
    undelegationPeriod,
  }
}

const getDelegations = async (
  operatorAddresses,
  web3Context,
  initializationPeriod,
  undelegationPeriod,
  isFromGrant = false,
) => {
  let tokenStakingBalance = web3Utils.toBN(0)
  let pendingUndelegationBalance = web3Utils.toBN(0)
  const delegations = []
  const undelegations = []

  for (const operatorAddress of operatorAddresses) {
    const {
      createdAt,
      undelegatedAt,
      amount,
    } = await contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'getDelegationInfo', operatorAddress)
    const beneficiary = await contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'magpieOf', operatorAddress)
    const authorizerAddress = await contractService.makeCall(web3Context, TOKEN_STAKING_CONTRACT_NAME, 'authorizerOf', operatorAddress)

    let grantId;
    if (isFromGrant) {
      let grantStakeDetails = await contractService.makeCall(web3Context, TOKEN_GRANT_CONTRACT_NAME, 'getGrantStakeDetails', operatorAddress)
      
      grantId = grantStakeDetails.grantId
      console.log("grantId", grantId)
    }

    const operatorData = {
      undelegatedAt,
      amount,
      beneficiary,
      operatorAddress,
      createdAt,
      authorizerAddress,
      isFromGrant,
      grantId,
    }
    const balance = web3Utils.toBN(amount)

    if (!balance.isZero() && operatorData.undelegatedAt === '0') {
      const initializationOverAt = moment.unix(createdAt).add(initializationPeriod, 'seconds')
      operatorData.isInInitializationPeriod = moment().isSameOrBefore(initializationOverAt)
      operatorData.initializationOverAt = initializationOverAt
      delegations.push(operatorData)
      if (!isFromGrant) {
        tokenStakingBalance = tokenStakingBalance.add(balance)
      }
    }
    if (operatorData.undelegatedAt !== '0' && gt(amount, 0)) {
      operatorData.undelegationCompleteAt = moment.unix(undelegatedAt).add(undelegationPeriod, 'seconds')
      operatorData.canRecoverStake = operatorData.undelegationCompleteAt.isBefore(moment())
      if (!isFromGrant) {
        pendingUndelegationBalance = pendingUndelegationBalance.add(balance)
      }
      undelegations.push(operatorData)
    }
  }

  return [
    delegations,
    undelegations,
    tokenStakingBalance,
    pendingUndelegationBalance,
  ]
}

const delegateStake = async (web3Context, data, onTransactionHashCallback) => {
  const {
    authorizerAddress,
    beneficiaryAddress,
    operatorAddress,
    stakeTokens,
    context,
    selectedGrant,
  } = data
  const amount = web3Utils.toBN(stakeTokens).mul(web3Utils.toBN(10).pow(web3Utils.toBN(18))).toString()
  const delegation = '0x' + Buffer.concat([
    Buffer.from(beneficiaryAddress.substr(2), 'hex'),
    Buffer.from(operatorAddress.substr(2), 'hex'),
    Buffer.from(authorizerAddress.substr(2), 'hex'),
  ]).toString('hex')

  const { token, stakingContract, grantContract, yourAddress } = web3Context

  if (context === 'owned') {
    await token.methods
      .approveAndCall(stakingContract.options.address, amount, delegation)
      .send({ from: yourAddress })
      .on('transactionHash', onTransactionHashCallback)
  } else if (context === 'granted') {
    await grantContract.methods.stake(selectedGrant.id, stakingContract.options.address, amount, delegation)
      .send({ from: yourAddress })
      .on('transactionHash', onTransactionHashCallback)
  }
}

export const tokensPageService = {
  fetchTokensPageData,
  delegateStake,
}
