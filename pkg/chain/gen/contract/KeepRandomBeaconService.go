// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ethereumabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ipfs/go-log"

	"github.com/keep-network/keep-common/pkg/chain/ethereum/ethutil"
	"github.com/keep-network/keep-common/pkg/subscription"
	"github.com/keep-network/keep-core/pkg/chain/gen/abi"
)

// Create a package-level logger for this contract. The logger exists at
// package level so that the logger is registered at startup and can be
// included or excluded from logging at startup by name.
var krbsLogger = log.Logger("keep-contract-KeepRandomBeaconService")

type KeepRandomBeaconService struct {
	contract          *abi.KeepRandomBeaconServiceImplV1
	contractAddress   common.Address
	contractABI       *ethereumabi.ABI
	caller            bind.ContractCaller
	transactor        bind.ContractTransactor
	callerOptions     *bind.CallOpts
	transactorOptions *bind.TransactOpts
	errorResolver     *ethutil.ErrorResolver

	transactionMutex *sync.Mutex
}

func NewKeepRandomBeaconService(
	contractAddress common.Address,
	accountKey *keystore.Key,
	backend bind.ContractBackend,
	transactionMutex *sync.Mutex,
) (*KeepRandomBeaconService, error) {
	callerOptions := &bind.CallOpts{
		From: contractAddress,
	}
	transactorOptions := bind.NewKeyedTransactor(
		accountKey.PrivateKey,
	)

	randomBeaconContract, err := abi.NewKeepRandomBeaconServiceImplV1(
		contractAddress,
		backend,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to instantiate contract at address: %s [%v]",
			contractAddress.String(),
			err,
		)
	}

	contractABI, err := ethereumabi.JSON(strings.NewReader(abi.KeepRandomBeaconServiceImplV1ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ABI: [%v]", err)
	}

	return &KeepRandomBeaconService{
		contract:          randomBeaconContract,
		contractAddress:   contractAddress,
		contractABI:       &contractABI,
		caller:            backend,
		transactor:        backend,
		callerOptions:     callerOptions,
		transactorOptions: transactorOptions,
		errorResolver:     ethutil.NewErrorResolver(backend, &contractABI, &contractAddress),
		transactionMutex:  transactionMutex,
	}, nil
}

// ----- Non-const Methods ------

// Transaction submission.
func (krbs *KeepRandomBeaconService) EntryCreated(
	requestId *big.Int,
	entry []uint8,
	submitter common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction entryCreated",
		"params: ",
		fmt.Sprint(
			requestId,
			entry,
			submitter,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.EntryCreated(
		transactorOptions,
		requestId,
		entry,
		submitter,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"entryCreated",
			requestId,
			entry,
			submitter,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction entryCreated with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallEntryCreated(
	requestId *big.Int,
	entry []uint8,
	submitter common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"entryCreated",
		&result,
		requestId,
		entry,
		submitter,
	)

	return err
}

func (krbs *KeepRandomBeaconService) EntryCreatedGasEstimate(
	requestId *big.Int,
	entry []uint8,
	submitter common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"entryCreated",
		krbs.contractABI,
		krbs.transactor,
		requestId,
		entry,
		submitter,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) FundDkgFeePool(
	value *big.Int,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction fundDkgFeePool",
		"value: ", value,
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	transactorOptions.Value = value

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.FundDkgFeePool(
		transactorOptions,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			value,
			"fundDkgFeePool",
		)
	}

	krbsLogger.Debugf(
		"submitted transaction fundDkgFeePool with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallFundDkgFeePool(
	value *big.Int,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, value,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"fundDkgFeePool",
		&result,
	)

	return err
}

func (krbs *KeepRandomBeaconService) FundDkgFeePoolGasEstimate() (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"fundDkgFeePool",
		krbs.contractABI,
		krbs.transactor,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) RemoveOperatorContract(
	operatorContract common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction removeOperatorContract",
		"params: ",
		fmt.Sprint(
			operatorContract,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.RemoveOperatorContract(
		transactorOptions,
		operatorContract,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"removeOperatorContract",
			operatorContract,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction removeOperatorContract with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallRemoveOperatorContract(
	operatorContract common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"removeOperatorContract",
		&result,
		operatorContract,
	)

	return err
}

func (krbs *KeepRandomBeaconService) RemoveOperatorContractGasEstimate(
	operatorContract common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"removeOperatorContract",
		krbs.contractABI,
		krbs.transactor,
		operatorContract,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) RenounceOwnership(

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction renounceOwnership",
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.RenounceOwnership(
		transactorOptions,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"renounceOwnership",
		)
	}

	krbsLogger.Debugf(
		"submitted transaction renounceOwnership with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallRenounceOwnership(
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"renounceOwnership",
		&result,
	)

	return err
}

func (krbs *KeepRandomBeaconService) RenounceOwnershipGasEstimate() (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"renounceOwnership",
		krbs.contractABI,
		krbs.transactor,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) RequestRelayEntry(
	value *big.Int,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction requestRelayEntry",
		"value: ", value,
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	transactorOptions.Value = value

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.RequestRelayEntry(
		transactorOptions,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			value,
			"requestRelayEntry",
		)
	}

	krbsLogger.Debugf(
		"submitted transaction requestRelayEntry with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallRequestRelayEntry(
	value *big.Int,
	blockNumber *big.Int,
) (*big.Int, error) {
	var result *big.Int

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, value,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"requestRelayEntry",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) RequestRelayEntryGasEstimate() (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"requestRelayEntry",
		krbs.contractABI,
		krbs.transactor,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) ExecuteCallback(
	requestId *big.Int,
	entry *big.Int,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction executeCallback",
		"params: ",
		fmt.Sprint(
			requestId,
			entry,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.ExecuteCallback(
		transactorOptions,
		requestId,
		entry,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"executeCallback",
			requestId,
			entry,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction executeCallback with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallExecuteCallback(
	requestId *big.Int,
	entry *big.Int,
	blockNumber *big.Int,
) (common.Address, error) {
	var result common.Address

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"executeCallback",
		&result,
		requestId,
		entry,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) ExecuteCallbackGasEstimate(
	requestId *big.Int,
	entry *big.Int,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"executeCallback",
		krbs.contractABI,
		krbs.transactor,
		requestId,
		entry,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) FinishWithdrawal(
	payee common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction finishWithdrawal",
		"params: ",
		fmt.Sprint(
			payee,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.FinishWithdrawal(
		transactorOptions,
		payee,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"finishWithdrawal",
			payee,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction finishWithdrawal with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallFinishWithdrawal(
	payee common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"finishWithdrawal",
		&result,
		payee,
	)

	return err
}

func (krbs *KeepRandomBeaconService) FinishWithdrawalGasEstimate(
	payee common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"finishWithdrawal",
		krbs.contractABI,
		krbs.transactor,
		payee,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) Initialize(
	dkgContributionMargin *big.Int,
	withdrawalDelay *big.Int,
	registry common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction initialize",
		"params: ",
		fmt.Sprint(
			dkgContributionMargin,
			withdrawalDelay,
			registry,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.Initialize(
		transactorOptions,
		dkgContributionMargin,
		withdrawalDelay,
		registry,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"initialize",
			dkgContributionMargin,
			withdrawalDelay,
			registry,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction initialize with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallInitialize(
	dkgContributionMargin *big.Int,
	withdrawalDelay *big.Int,
	registry common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"initialize",
		&result,
		dkgContributionMargin,
		withdrawalDelay,
		registry,
	)

	return err
}

func (krbs *KeepRandomBeaconService) InitializeGasEstimate(
	dkgContributionMargin *big.Int,
	withdrawalDelay *big.Int,
	registry common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"initialize",
		krbs.contractABI,
		krbs.transactor,
		dkgContributionMargin,
		withdrawalDelay,
		registry,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) Initialize0(
	sender common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction initialize0",
		"params: ",
		fmt.Sprint(
			sender,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.Initialize0(
		transactorOptions,
		sender,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"initialize0",
			sender,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction initialize0 with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallInitialize0(
	sender common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"initialize0",
		&result,
		sender,
	)

	return err
}

func (krbs *KeepRandomBeaconService) Initialize0GasEstimate(
	sender common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"initialize0",
		krbs.contractABI,
		krbs.transactor,
		sender,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) RequestRelayEntry0(
	callbackContract common.Address,
	callbackGas *big.Int,
	value *big.Int,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction requestRelayEntry0",
		"params: ",
		fmt.Sprint(
			callbackContract,
			callbackGas,
		),
		"value: ", value,
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	transactorOptions.Value = value

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.RequestRelayEntry0(
		transactorOptions,
		callbackContract,
		callbackGas,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			value,
			"requestRelayEntry0",
			callbackContract,
			callbackGas,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction requestRelayEntry0 with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallRequestRelayEntry0(
	callbackContract common.Address,
	callbackGas *big.Int,
	value *big.Int,
	blockNumber *big.Int,
) (*big.Int, error) {
	var result *big.Int

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, value,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"requestRelayEntry0",
		&result,
		callbackContract,
		callbackGas,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) RequestRelayEntry0GasEstimate(
	callbackContract common.Address,
	callbackGas *big.Int,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"requestRelayEntry0",
		krbs.contractABI,
		krbs.transactor,
		callbackContract,
		callbackGas,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) TransferOwnership(
	newOwner common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction transferOwnership",
		"params: ",
		fmt.Sprint(
			newOwner,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.TransferOwnership(
		transactorOptions,
		newOwner,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"transferOwnership",
			newOwner,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction transferOwnership with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallTransferOwnership(
	newOwner common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"transferOwnership",
		&result,
		newOwner,
	)

	return err
}

func (krbs *KeepRandomBeaconService) TransferOwnershipGasEstimate(
	newOwner common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"transferOwnership",
		krbs.contractABI,
		krbs.transactor,
		newOwner,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) AddOperatorContract(
	operatorContract common.Address,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction addOperatorContract",
		"params: ",
		fmt.Sprint(
			operatorContract,
		),
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.AddOperatorContract(
		transactorOptions,
		operatorContract,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"addOperatorContract",
			operatorContract,
		)
	}

	krbsLogger.Debugf(
		"submitted transaction addOperatorContract with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallAddOperatorContract(
	operatorContract common.Address,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"addOperatorContract",
		&result,
		operatorContract,
	)

	return err
}

func (krbs *KeepRandomBeaconService) AddOperatorContractGasEstimate(
	operatorContract common.Address,
) (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"addOperatorContract",
		krbs.contractABI,
		krbs.transactor,
		operatorContract,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) FundRequestSubsidyFeePool(
	value *big.Int,

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction fundRequestSubsidyFeePool",
		"value: ", value,
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	transactorOptions.Value = value

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.FundRequestSubsidyFeePool(
		transactorOptions,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			value,
			"fundRequestSubsidyFeePool",
		)
	}

	krbsLogger.Debugf(
		"submitted transaction fundRequestSubsidyFeePool with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallFundRequestSubsidyFeePool(
	value *big.Int,
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, value,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"fundRequestSubsidyFeePool",
		&result,
	)

	return err
}

func (krbs *KeepRandomBeaconService) FundRequestSubsidyFeePoolGasEstimate() (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"fundRequestSubsidyFeePool",
		krbs.contractABI,
		krbs.transactor,
	)

	return result, err
}

// Transaction submission.
func (krbs *KeepRandomBeaconService) InitiateWithdrawal(

	transactionOptions ...ethutil.TransactionOptions,
) (*types.Transaction, error) {
	krbsLogger.Debug(
		"submitting transaction initiateWithdrawal",
	)

	krbs.transactionMutex.Lock()
	defer krbs.transactionMutex.Unlock()

	// create a copy
	transactorOptions := new(bind.TransactOpts)
	*transactorOptions = *krbs.transactorOptions

	if len(transactionOptions) > 1 {
		return nil, fmt.Errorf(
			"could not process multiple transaction options sets",
		)
	} else if len(transactionOptions) > 0 {
		transactionOptions[0].Apply(transactorOptions)
	}

	transaction, err := krbs.contract.InitiateWithdrawal(
		transactorOptions,
	)

	if err != nil {
		return transaction, krbs.errorResolver.ResolveError(
			err,
			krbs.transactorOptions.From,
			nil,
			"initiateWithdrawal",
		)
	}

	krbsLogger.Debugf(
		"submitted transaction initiateWithdrawal with id: [%v]",
		transaction.Hash().Hex(),
	)

	return transaction, err
}

// Non-mutating call, not a transaction submission.
func (krbs *KeepRandomBeaconService) CallInitiateWithdrawal(
	blockNumber *big.Int,
) error {
	var result interface{} = nil

	err := ethutil.CallAtBlock(
		krbs.transactorOptions.From,
		blockNumber, nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"initiateWithdrawal",
		&result,
	)

	return err
}

func (krbs *KeepRandomBeaconService) InitiateWithdrawalGasEstimate() (uint64, error) {
	var result uint64

	result, err := ethutil.EstimateGas(
		krbs.callerOptions.From,
		krbs.contractAddress,
		"initiateWithdrawal",
		krbs.contractABI,
		krbs.transactor,
	)

	return result, err
}

// ----- Const Methods ------

type entryFeeBreakdown struct {
	EntryVerificationFee *big.Int
	DkgContributionFee   *big.Int
	GroupProfitFee       *big.Int
	GasPriceCeiling      *big.Int
}

func (krbs *KeepRandomBeaconService) EntryFeeBreakdown() (entryFeeBreakdown, error) {
	var result entryFeeBreakdown
	result, err := krbs.contract.EntryFeeBreakdown(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"entryFeeBreakdown",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) EntryFeeBreakdownAtBlock(
	blockNumber *big.Int,
) (entryFeeBreakdown, error) {
	var result entryFeeBreakdown

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"entryFeeBreakdown",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) EntryFeeEstimate(
	callbackGas *big.Int,
) (*big.Int, error) {
	var result *big.Int
	result, err := krbs.contract.EntryFeeEstimate(
		krbs.callerOptions,
		callbackGas,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"entryFeeEstimate",
			callbackGas,
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) EntryFeeEstimateAtBlock(
	callbackGas *big.Int,
	blockNumber *big.Int,
) (*big.Int, error) {
	var result *big.Int

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"entryFeeEstimate",
		&result,
		callbackGas,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) IsOwner() (bool, error) {
	var result bool
	result, err := krbs.contract.IsOwner(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"isOwner",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) IsOwnerAtBlock(
	blockNumber *big.Int,
) (bool, error) {
	var result bool

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"isOwner",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) Owner() (common.Address, error) {
	var result common.Address
	result, err := krbs.contract.Owner(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"owner",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) OwnerAtBlock(
	blockNumber *big.Int,
) (common.Address, error) {
	var result common.Address

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"owner",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) BaseCallbackGas() (*big.Int, error) {
	var result *big.Int
	result, err := krbs.contract.BaseCallbackGas(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"baseCallbackGas",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) BaseCallbackGasAtBlock(
	blockNumber *big.Int,
) (*big.Int, error) {
	var result *big.Int

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"baseCallbackGas",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) Initialized() (bool, error) {
	var result bool
	result, err := krbs.contract.Initialized(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"initialized",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) InitializedAtBlock(
	blockNumber *big.Int,
) (bool, error) {
	var result bool

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"initialized",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) PreviousEntry() ([]uint8, error) {
	var result []uint8
	result, err := krbs.contract.PreviousEntry(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"previousEntry",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) PreviousEntryAtBlock(
	blockNumber *big.Int,
) ([]uint8, error) {
	var result []uint8

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"previousEntry",
		&result,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) SelectOperatorContract(
	seed *big.Int,
) (common.Address, error) {
	var result common.Address
	result, err := krbs.contract.SelectOperatorContract(
		krbs.callerOptions,
		seed,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"selectOperatorContract",
			seed,
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) SelectOperatorContractAtBlock(
	seed *big.Int,
	blockNumber *big.Int,
) (common.Address, error) {
	var result common.Address

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"selectOperatorContract",
		&result,
		seed,
	)

	return result, err
}

func (krbs *KeepRandomBeaconService) Version() (string, error) {
	var result string
	result, err := krbs.contract.Version(
		krbs.callerOptions,
	)

	if err != nil {
		return result, krbs.errorResolver.ResolveError(
			err,
			krbs.callerOptions.From,
			nil,
			"version",
		)
	}

	return result, err
}

func (krbs *KeepRandomBeaconService) VersionAtBlock(
	blockNumber *big.Int,
) (string, error) {
	var result string

	err := ethutil.CallAtBlock(
		krbs.callerOptions.From,
		blockNumber,
		nil,
		krbs.contractABI,
		krbs.caller,
		krbs.errorResolver,
		krbs.contractAddress,
		"version",
		&result,
	)

	return result, err
}

// ------ Events -------

type keepRandomBeaconServiceRelayEntryRequestedFunc func(
	RequestId *big.Int,
	blockNumber uint64,
)

func (krbs *KeepRandomBeaconService) WatchRelayEntryRequested(
	success keepRandomBeaconServiceRelayEntryRequestedFunc,
	fail func(err error) error,
) (subscription.EventSubscription, error) {
	errorChan := make(chan error)
	unsubscribeChan := make(chan struct{})

	// Delay which must be preserved before a new resubscription attempt.
	// There is no sense to resubscribe immediately after the fail of current
	// subscription because the publisher must have some time to recover.
	retryDelay := 5 * time.Second

	watch := func() {
		failCallback := func(err error) error {
			fail(err)
			errorChan <- err // trigger resubscription signal
			return err
		}

		subscription, err := krbs.subscribeRelayEntryRequested(
			success,
			failCallback,
		)
		if err != nil {
			errorChan <- err // trigger resubscription signal
			return
		}

		// wait for unsubscription signal
		<-unsubscribeChan
		subscription.Unsubscribe()
	}

	// trigger the resubscriber goroutine
	go func() {
		go watch() // trigger first subscription

		for {
			select {
			case <-errorChan:
				krbsLogger.Warning(
					"subscription to event RelayEntryRequested terminated with error; " +
						"resubscription attempt will be performed after the retry delay",
				)
				time.Sleep(retryDelay)
				go watch()
			case <-unsubscribeChan:
				// shutdown the resubscriber goroutine on unsubscribe signal
				return
			}
		}
	}()

	// closing the unsubscribeChan will trigger a unsubscribe signal and
	// run unsubscription for all subscription instances
	unsubscribeCallback := func() {
		close(unsubscribeChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}

func (krbs *KeepRandomBeaconService) subscribeRelayEntryRequested(
	success keepRandomBeaconServiceRelayEntryRequestedFunc,
	fail func(err error) error,
) (subscription.EventSubscription, error) {
	eventChan := make(chan *abi.KeepRandomBeaconServiceImplV1RelayEntryRequested)
	eventSubscription, err := krbs.contract.WatchRelayEntryRequested(
		nil,
		eventChan,
	)
	if err != nil {
		close(eventChan)
		return eventSubscription, fmt.Errorf(
			"error creating watch for RelayEntryRequested events: [%v]",
			err,
		)
	}

	var subscriptionMutex = &sync.Mutex{}

	go func() {
		for {
			select {
			case event, subscribed := <-eventChan:
				subscriptionMutex.Lock()
				// if eventChan has been closed, it means we have unsubscribed
				if !subscribed {
					subscriptionMutex.Unlock()
					return
				}
				success(
					event.RequestId,
					event.Raw.BlockNumber,
				)
				subscriptionMutex.Unlock()
			case ee := <-eventSubscription.Err():
				fail(ee)
				return
			}
		}
	}()

	unsubscribeCallback := func() {
		subscriptionMutex.Lock()
		defer subscriptionMutex.Unlock()

		eventSubscription.Unsubscribe()
		close(eventChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}

type keepRandomBeaconServiceOwnershipTransferredFunc func(
	PreviousOwner common.Address,
	NewOwner common.Address,
	blockNumber uint64,
)

func (krbs *KeepRandomBeaconService) WatchOwnershipTransferred(
	success keepRandomBeaconServiceOwnershipTransferredFunc,
	fail func(err error) error,
	previousOwnerFilter []common.Address,
	newOwnerFilter []common.Address,
) (subscription.EventSubscription, error) {
	errorChan := make(chan error)
	unsubscribeChan := make(chan struct{})

	// Delay which must be preserved before a new resubscription attempt.
	// There is no sense to resubscribe immediately after the fail of current
	// subscription because the publisher must have some time to recover.
	retryDelay := 5 * time.Second

	watch := func() {
		failCallback := func(err error) error {
			fail(err)
			errorChan <- err // trigger resubscription signal
			return err
		}

		subscription, err := krbs.subscribeOwnershipTransferred(
			success,
			failCallback,
			previousOwnerFilter,
			newOwnerFilter,
		)
		if err != nil {
			errorChan <- err // trigger resubscription signal
			return
		}

		// wait for unsubscription signal
		<-unsubscribeChan
		subscription.Unsubscribe()
	}

	// trigger the resubscriber goroutine
	go func() {
		go watch() // trigger first subscription

		for {
			select {
			case <-errorChan:
				krbsLogger.Warning(
					"subscription to event OwnershipTransferred terminated with error; " +
						"resubscription attempt will be performed after the retry delay",
				)
				time.Sleep(retryDelay)
				go watch()
			case <-unsubscribeChan:
				// shutdown the resubscriber goroutine on unsubscribe signal
				return
			}
		}
	}()

	// closing the unsubscribeChan will trigger a unsubscribe signal and
	// run unsubscription for all subscription instances
	unsubscribeCallback := func() {
		close(unsubscribeChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}

func (krbs *KeepRandomBeaconService) subscribeOwnershipTransferred(
	success keepRandomBeaconServiceOwnershipTransferredFunc,
	fail func(err error) error,
	previousOwnerFilter []common.Address,
	newOwnerFilter []common.Address,
) (subscription.EventSubscription, error) {
	eventChan := make(chan *abi.KeepRandomBeaconServiceImplV1OwnershipTransferred)
	eventSubscription, err := krbs.contract.WatchOwnershipTransferred(
		nil,
		eventChan,
		previousOwnerFilter,
		newOwnerFilter,
	)
	if err != nil {
		close(eventChan)
		return eventSubscription, fmt.Errorf(
			"error creating watch for OwnershipTransferred events: [%v]",
			err,
		)
	}

	var subscriptionMutex = &sync.Mutex{}

	go func() {
		for {
			select {
			case event, subscribed := <-eventChan:
				subscriptionMutex.Lock()
				// if eventChan has been closed, it means we have unsubscribed
				if !subscribed {
					subscriptionMutex.Unlock()
					return
				}
				success(
					event.PreviousOwner,
					event.NewOwner,
					event.Raw.BlockNumber,
				)
				subscriptionMutex.Unlock()
			case ee := <-eventSubscription.Err():
				fail(ee)
				return
			}
		}
	}()

	unsubscribeCallback := func() {
		subscriptionMutex.Lock()
		defer subscriptionMutex.Unlock()

		eventSubscription.Unsubscribe()
		close(eventChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}

type keepRandomBeaconServiceRelayEntryGeneratedFunc func(
	RequestId *big.Int,
	Entry *big.Int,
	blockNumber uint64,
)

func (krbs *KeepRandomBeaconService) WatchRelayEntryGenerated(
	success keepRandomBeaconServiceRelayEntryGeneratedFunc,
	fail func(err error) error,
) (subscription.EventSubscription, error) {
	errorChan := make(chan error)
	unsubscribeChan := make(chan struct{})

	// Delay which must be preserved before a new resubscription attempt.
	// There is no sense to resubscribe immediately after the fail of current
	// subscription because the publisher must have some time to recover.
	retryDelay := 5 * time.Second

	watch := func() {
		failCallback := func(err error) error {
			fail(err)
			errorChan <- err // trigger resubscription signal
			return err
		}

		subscription, err := krbs.subscribeRelayEntryGenerated(
			success,
			failCallback,
		)
		if err != nil {
			errorChan <- err // trigger resubscription signal
			return
		}

		// wait for unsubscription signal
		<-unsubscribeChan
		subscription.Unsubscribe()
	}

	// trigger the resubscriber goroutine
	go func() {
		go watch() // trigger first subscription

		for {
			select {
			case <-errorChan:
				krbsLogger.Warning(
					"subscription to event RelayEntryGenerated terminated with error; " +
						"resubscription attempt will be performed after the retry delay",
				)
				time.Sleep(retryDelay)
				go watch()
			case <-unsubscribeChan:
				// shutdown the resubscriber goroutine on unsubscribe signal
				return
			}
		}
	}()

	// closing the unsubscribeChan will trigger a unsubscribe signal and
	// run unsubscription for all subscription instances
	unsubscribeCallback := func() {
		close(unsubscribeChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}

func (krbs *KeepRandomBeaconService) subscribeRelayEntryGenerated(
	success keepRandomBeaconServiceRelayEntryGeneratedFunc,
	fail func(err error) error,
) (subscription.EventSubscription, error) {
	eventChan := make(chan *abi.KeepRandomBeaconServiceImplV1RelayEntryGenerated)
	eventSubscription, err := krbs.contract.WatchRelayEntryGenerated(
		nil,
		eventChan,
	)
	if err != nil {
		close(eventChan)
		return eventSubscription, fmt.Errorf(
			"error creating watch for RelayEntryGenerated events: [%v]",
			err,
		)
	}

	var subscriptionMutex = &sync.Mutex{}

	go func() {
		for {
			select {
			case event, subscribed := <-eventChan:
				subscriptionMutex.Lock()
				// if eventChan has been closed, it means we have unsubscribed
				if !subscribed {
					subscriptionMutex.Unlock()
					return
				}
				success(
					event.RequestId,
					event.Entry,
					event.Raw.BlockNumber,
				)
				subscriptionMutex.Unlock()
			case ee := <-eventSubscription.Err():
				fail(ee)
				return
			}
		}
	}()

	unsubscribeCallback := func() {
		subscriptionMutex.Lock()
		defer subscriptionMutex.Unlock()

		eventSubscription.Unsubscribe()
		close(eventChan)
	}

	return subscription.NewEventSubscription(unsubscribeCallback), nil
}
