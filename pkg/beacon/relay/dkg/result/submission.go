package result

import (
	"fmt"

	relayChain "github.com/keep-network/keep-core/pkg/beacon/relay/chain"
	"github.com/keep-network/keep-core/pkg/beacon/relay/event"
	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/operator"
)

// SubmittingMember represents a member submitting a DKG result to the
// blockchain along with signatures received from other group members supporting
// the result.
type SubmittingMember struct {
	// Represents the member's position for submission.
	index group.MemberIndex
}

// NewSubmittingMember creates a member to execute submitting the DKG result hash.
func NewSubmittingMember(
	memberIndex group.MemberIndex,
) *SubmittingMember {
	return &SubmittingMember{
		index: memberIndex,
	}
}

// SubmitDKGResult sends a result, which contains the group public key and
// signatures, to the chain.
//
// It checks if the result has already been published to the blockchain by
// checking if a group with the given public key is already registered. If not,
// it determines if the current member is eligible to submit a result.
// If allowed, it submits the result to the chain.
//
// A user's turn to publish is determined based on the user's index and block
// step.
//
// If a result is submitted by another member and it's accepted by the chain,
// the current member finishes the phase immediately, without submitting
// their own result.
//
// It returns the on-chain block height of the moment when the result was
// successfully submitted on chain by the member. In case of failure or result
// already submitted by another member it returns `0`.
//
// See Phase 14 of the protocol specification.
func (sm *SubmittingMember) SubmitDKGResult(
	result *relayChain.DKGResult,
	signatures map[group.MemberIndex]operator.Signature,
	chainRelay relayChain.Interface,
	blockCounter chain.BlockCounter,
	startBlockHeight uint64,
) error {
	config, err := chainRelay.GetConfig()
	if err != nil {
		return fmt.Errorf(
			"could not fetch chain's config: [%v]",
			err,
		)
	}

	onSubmittedResultChan := make(chan *event.DKGResultSubmission)

	subscription, err := chainRelay.OnDKGResultSubmitted(
		func(event *event.DKGResultSubmission) {
			onSubmittedResultChan <- event
		},
	)
	if err != nil {
		close(onSubmittedResultChan)
		return fmt.Errorf(
			"could not watch for DKG result publications: [%v]",
			err,
		)
	}

	returnWithError := func(err error) error {
		subscription.Unsubscribe()
		close(onSubmittedResultChan)
		return err
	}

	alreadySubmitted, err := chainRelay.IsGroupRegistered(result.GroupPublicKey)
	if err != nil {
		return returnWithError(
			fmt.Errorf(
				"could not check if the result is already submitted: [%v]",
				err,
			),
		)
	}

	// Someone who was ahead of us in the queue submitted the result. Giving up.
	if alreadySubmitted {
		return returnWithError(nil)
	}

	// Wait until the current member is eligible to submit the result.
	eligibleToSubmitWaiter, err := sm.waitForSubmissionEligibility(
		blockCounter,
		startBlockHeight,
		config.ResultPublicationBlockStep,
	)
	if err != nil {
		return returnWithError(
			fmt.Errorf("wait for eligibility failure: [%v]", err),
		)
	}

	for {
		select {
		case <-eligibleToSubmitWaiter:
			// Member becomes eligible to submit the result.
			errorChannel := make(chan error)
			defer close(errorChannel)

			subscription.Unsubscribe()
			close(onSubmittedResultChan)

			logger.Infof("[member:%v] submitting DKG result", sm.index)
			chainRelay.SubmitDKGResult(
				sm.index,
				result,
				signatures,
			).
				OnComplete(func(
					dkgResultPublishedEvent *event.DKGResultSubmission,
					err error,
				) {
					errorChannel <- err
				})
			return <-errorChannel
		case <-onSubmittedResultChan:
			logger.Infof(
				"[member:%v] leaving; DKG result was submitted by other member",
				sm.index,
			)
			// A result has been submitted by other member. Leave without
			// publishing the result.
			return returnWithError(nil)
		}
	}
}

// waitForSubmissionEligibility waits until the current member is eligible to
// submit a result to the blockchain. First member is eligible to submit straight
// away, each following member is eligible after pre-defined block step.
func (sm *SubmittingMember) waitForSubmissionEligibility(
	blockCounter chain.BlockCounter,
	startBlockHeight uint64,
	blockStep uint64,
) (<-chan uint64, error) {
	// T_init + (member_index - 1) * T_step
	blockWaitTime := (uint64(sm.index) - 1) * blockStep

	eligibleBlockHeight := startBlockHeight + blockWaitTime
	logger.Infof(
		"[member:%v] waiting for block [%v] to submit",
		sm.index,
		eligibleBlockHeight,
	)

	waiter, err := blockCounter.BlockHeightWaiter(eligibleBlockHeight)
	if err != nil {
		return nil, fmt.Errorf("block height waiter failure [%v]", err)
	}

	return waiter, err
}
