// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DelayedVetoableMetaData contains all meta data concerning the DelayedVetoable contract.
var DelayedVetoableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractSuperchainConfig\",\"name\":\"superchainConfig_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyDelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwardingEarly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetUnitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"DelayActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Forwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Vetoed\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatingDelay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"}],\"name\":\"queuedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"queuedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superchainConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"superchainConfig_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vetoer_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516109fb3803806109fb83398101604081905261002f91610078565b600180546001600160a01b039384166001600160a01b031991821617909155600080549290931691161790556100b2565b6001600160a01b038116811461007557600080fd5b50565b6000806040838503121561008b57600080fd5b825161009681610060565b60208401519092506100a781610060565b809150509250929050565b61093a806100c16000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80636a42b8f81161005b5780636a42b8f81461012b578063b912de5d14610133578063d4b8399214610146578063d8bff4401461014e57610088565b80632750a0bc1461009257806335e80ab3146100ad57806354fd4d50146100da5780635c39fcc114610123575b610090610156565b005b61009a610477565b6040519081526020015b60405180910390f35b6100b5610496565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100a4565b6101166040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100a491906107af565b6100b56104ba565b61009a6104c9565b61009a610141366004610822565b6104d7565b6100b56104fe565b6100b5610522565b361580156101645750600254155b1561028357610171610531565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580156101df57506101af6105c5565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610241576101ec610531565b6040517f295a81c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015233602482015260440160405180910390fd5b610249610635565b60028190556040519081527febf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a89060200160405180910390a1565b6000803660405161029592919061083b565b604051809103902090506102a7610531565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480156102ed5750600081815260036020526040902054155b156103535760025460000361030557610305816106c9565b6000818152600360205260408082204290555182917f87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13916103489190369061084b565b60405180910390a250565b61035b6105c5565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480156103a2575060008181526003602052604090205415155b156103ea576000818152600360205260408082208290555182917fbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040916103489190369061084b565b6000818152600360205260408120549003610407576101ec610531565b600254600082815260036020526040902054429161042491610898565b101561045c576040517f43dc986d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260036020526040812055610474816106c9565b50565b60003361048b57610486610635565b905090565b610493610156565b90565b60003361048b575060015473ffffffffffffffffffffffffffffffffffffffff1690565b60003361048b57610486610531565b60003361048b575060025490565b6000336104f1575060009081526003602052604090205490565b6104f9610156565b919050565b60003361048b575060005473ffffffffffffffffffffffffffffffffffffffff1690565b60003361048b576104866105c5565b600154604080517f5c39fcc1000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c39fcc19160048083019260209291908290030181865afa1580156105a1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048691906108d7565b600154604080517fd8bff440000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163d8bff4409160048083019260209291908290030181865afa1580156105a1573d6000803e3d6000fd5b600154604080517f6a42b8f8000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691636a42b8f89160048083019260209291908290030181865afa1580156106a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104869190610914565b807f4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e12296000366040516106fc92919061084b565b60405180910390a260008061072660005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1660003660405161074d92919061083b565b6000604051808303816000865af19150503d806000811461078a576040519150601f19603f3d011682016040523d82523d6000602084013e61078f565b606091505b5090925090508115156001036107a757805160208201f35b805160208201fd5b600060208083528351808285015260005b818110156107dc578581018301518582016040015282016107c0565b818111156107ee576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561083457600080fd5b5035919050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b600082198211156108d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b6000602082840312156108e957600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461090d57600080fd5b9392505050565b60006020828403121561092657600080fd5b505191905056fea164736f6c634300080f000a",
}

// DelayedVetoableABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedVetoableMetaData.ABI instead.
var DelayedVetoableABI = DelayedVetoableMetaData.ABI

// DelayedVetoableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelayedVetoableMetaData.Bin instead.
var DelayedVetoableBin = DelayedVetoableMetaData.Bin

// DeployDelayedVetoable deploys a new Ethereum contract, binding an instance of DelayedVetoable to it.
func DeployDelayedVetoable(auth *bind.TransactOpts, backend bind.ContractBackend, superchainConfig_ common.Address, target_ common.Address) (common.Address, *types.Transaction, *DelayedVetoable, error) {
	parsed, err := DelayedVetoableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelayedVetoableBin), backend, superchainConfig_, target_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelayedVetoable{DelayedVetoableCaller: DelayedVetoableCaller{contract: contract}, DelayedVetoableTransactor: DelayedVetoableTransactor{contract: contract}, DelayedVetoableFilterer: DelayedVetoableFilterer{contract: contract}}, nil
}

// DelayedVetoable is an auto generated Go binding around an Ethereum contract.
type DelayedVetoable struct {
	DelayedVetoableCaller     // Read-only binding to the contract
	DelayedVetoableTransactor // Write-only binding to the contract
	DelayedVetoableFilterer   // Log filterer for contract events
}

// DelayedVetoableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedVetoableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedVetoableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedVetoableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedVetoableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedVetoableSession struct {
	Contract     *DelayedVetoable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelayedVetoableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedVetoableCallerSession struct {
	Contract *DelayedVetoableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DelayedVetoableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedVetoableTransactorSession struct {
	Contract     *DelayedVetoableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DelayedVetoableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedVetoableRaw struct {
	Contract *DelayedVetoable // Generic contract binding to access the raw methods on
}

// DelayedVetoableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedVetoableCallerRaw struct {
	Contract *DelayedVetoableCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedVetoableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedVetoableTransactorRaw struct {
	Contract *DelayedVetoableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedVetoable creates a new instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoable(address common.Address, backend bind.ContractBackend) (*DelayedVetoable, error) {
	contract, err := bindDelayedVetoable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoable{DelayedVetoableCaller: DelayedVetoableCaller{contract: contract}, DelayedVetoableTransactor: DelayedVetoableTransactor{contract: contract}, DelayedVetoableFilterer: DelayedVetoableFilterer{contract: contract}}, nil
}

// NewDelayedVetoableCaller creates a new read-only instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableCaller(address common.Address, caller bind.ContractCaller) (*DelayedVetoableCaller, error) {
	contract, err := bindDelayedVetoable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableCaller{contract: contract}, nil
}

// NewDelayedVetoableTransactor creates a new write-only instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedVetoableTransactor, error) {
	contract, err := bindDelayedVetoable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableTransactor{contract: contract}, nil
}

// NewDelayedVetoableFilterer creates a new log filterer instance of DelayedVetoable, bound to a specific deployed contract.
func NewDelayedVetoableFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedVetoableFilterer, error) {
	contract, err := bindDelayedVetoable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableFilterer{contract: contract}, nil
}

// bindDelayedVetoable binds a generic wrapper to an already deployed contract.
func bindDelayedVetoable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelayedVetoableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedVetoable *DelayedVetoableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedVetoable.Contract.DelayedVetoableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedVetoable *DelayedVetoableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.DelayedVetoableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedVetoable *DelayedVetoableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.DelayedVetoableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedVetoable *DelayedVetoableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedVetoable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedVetoable *DelayedVetoableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedVetoable *DelayedVetoableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelayedVetoable.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableSession) Version() (string, error) {
	return _DelayedVetoable.Contract.Version(&_DelayedVetoable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedVetoable *DelayedVetoableCallerSession) Version() (string, error) {
	return _DelayedVetoable.Contract.Version(&_DelayedVetoable.CallOpts)
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableTransactor) Delay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "delay")
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableSession) Delay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Delay(&_DelayedVetoable.TransactOpts)
}

// Delay is a paid mutator transaction binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() returns(uint256 delay_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Delay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Delay(&_DelayedVetoable.TransactOpts)
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableTransactor) Initiator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "initiator")
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableSession) Initiator() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Initiator(&_DelayedVetoable.TransactOpts)
}

// Initiator is a paid mutator transaction binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() returns(address initiator_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Initiator() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Initiator(&_DelayedVetoable.TransactOpts)
}

// OperatingDelay is a paid mutator transaction binding the contract method 0x2750a0bc.
//
// Solidity: function operatingDelay() returns(uint256 operatingDelay_)
func (_DelayedVetoable *DelayedVetoableTransactor) OperatingDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "operatingDelay")
}

// OperatingDelay is a paid mutator transaction binding the contract method 0x2750a0bc.
//
// Solidity: function operatingDelay() returns(uint256 operatingDelay_)
func (_DelayedVetoable *DelayedVetoableSession) OperatingDelay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.OperatingDelay(&_DelayedVetoable.TransactOpts)
}

// OperatingDelay is a paid mutator transaction binding the contract method 0x2750a0bc.
//
// Solidity: function operatingDelay() returns(uint256 operatingDelay_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) OperatingDelay() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.OperatingDelay(&_DelayedVetoable.TransactOpts)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableTransactor) QueuedAt(opts *bind.TransactOpts, callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "queuedAt", callHash)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableSession) QueuedAt(callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.QueuedAt(&_DelayedVetoable.TransactOpts, callHash)
}

// QueuedAt is a paid mutator transaction binding the contract method 0xb912de5d.
//
// Solidity: function queuedAt(bytes32 callHash) returns(uint256 queuedAt_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) QueuedAt(callHash [32]byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.QueuedAt(&_DelayedVetoable.TransactOpts, callHash)
}

// SuperchainConfig is a paid mutator transaction binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() returns(address superchainConfig_)
func (_DelayedVetoable *DelayedVetoableTransactor) SuperchainConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "superchainConfig")
}

// SuperchainConfig is a paid mutator transaction binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() returns(address superchainConfig_)
func (_DelayedVetoable *DelayedVetoableSession) SuperchainConfig() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.SuperchainConfig(&_DelayedVetoable.TransactOpts)
}

// SuperchainConfig is a paid mutator transaction binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() returns(address superchainConfig_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) SuperchainConfig() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.SuperchainConfig(&_DelayedVetoable.TransactOpts)
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableTransactor) Target(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "target")
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableSession) Target() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Target(&_DelayedVetoable.TransactOpts)
}

// Target is a paid mutator transaction binding the contract method 0xd4b83992.
//
// Solidity: function target() returns(address target_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Target() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Target(&_DelayedVetoable.TransactOpts)
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableTransactor) Vetoer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedVetoable.contract.Transact(opts, "vetoer")
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableSession) Vetoer() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Vetoer(&_DelayedVetoable.TransactOpts)
}

// Vetoer is a paid mutator transaction binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() returns(address vetoer_)
func (_DelayedVetoable *DelayedVetoableTransactorSession) Vetoer() (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Vetoer(&_DelayedVetoable.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Fallback(&_DelayedVetoable.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DelayedVetoable *DelayedVetoableTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DelayedVetoable.Contract.Fallback(&_DelayedVetoable.TransactOpts, calldata)
}

// DelayedVetoableDelayActivatedIterator is returned from FilterDelayActivated and is used to iterate over the raw logs and unpacked data for DelayActivated events raised by the DelayedVetoable contract.
type DelayedVetoableDelayActivatedIterator struct {
	Event *DelayedVetoableDelayActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableDelayActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableDelayActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableDelayActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableDelayActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableDelayActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableDelayActivated represents a DelayActivated event raised by the DelayedVetoable contract.
type DelayedVetoableDelayActivated struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelayActivated is a free log retrieval operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterDelayActivated(opts *bind.FilterOpts) (*DelayedVetoableDelayActivatedIterator, error) {

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "DelayActivated")
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableDelayActivatedIterator{contract: _DelayedVetoable.contract, event: "DelayActivated", logs: logs, sub: sub}, nil
}

// WatchDelayActivated is a free log subscription operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchDelayActivated(opts *bind.WatchOpts, sink chan<- *DelayedVetoableDelayActivated) (event.Subscription, error) {

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "DelayActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableDelayActivated)
				if err := _DelayedVetoable.contract.UnpackLog(event, "DelayActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayActivated is a log parse operation binding the contract event 0xebf28bfb587e28dfffd9173cf71c32ba5d3f0544a0117b5539c9b274a5bba2a8.
//
// Solidity: event DelayActivated(uint256 delay)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseDelayActivated(log types.Log) (*DelayedVetoableDelayActivated, error) {
	event := new(DelayedVetoableDelayActivated)
	if err := _DelayedVetoable.contract.UnpackLog(event, "DelayActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableForwardedIterator is returned from FilterForwarded and is used to iterate over the raw logs and unpacked data for Forwarded events raised by the DelayedVetoable contract.
type DelayedVetoableForwardedIterator struct {
	Event *DelayedVetoableForwarded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableForwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableForwarded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableForwarded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableForwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableForwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableForwarded represents a Forwarded event raised by the DelayedVetoable contract.
type DelayedVetoableForwarded struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterForwarded is a free log retrieval operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterForwarded(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableForwardedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Forwarded", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableForwardedIterator{contract: _DelayedVetoable.contract, event: "Forwarded", logs: logs, sub: sub}, nil
}

// WatchForwarded is a free log subscription operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchForwarded(opts *bind.WatchOpts, sink chan<- *DelayedVetoableForwarded, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Forwarded", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableForwarded)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Forwarded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForwarded is a log parse operation binding the contract event 0x4c109d85bcd0bb5c735b4be850953d652afe4cd9aa2e0b1426a65a4dcb2e1229.
//
// Solidity: event Forwarded(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseForwarded(log types.Log) (*DelayedVetoableForwarded, error) {
	event := new(DelayedVetoableForwarded)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Forwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the DelayedVetoable contract.
type DelayedVetoableInitiatedIterator struct {
	Event *DelayedVetoableInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableInitiated represents a Initiated event raised by the DelayedVetoable contract.
type DelayedVetoableInitiated struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterInitiated(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableInitiatedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Initiated", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableInitiatedIterator{contract: _DelayedVetoable.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *DelayedVetoableInitiated, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Initiated", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableInitiated)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Initiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiated is a log parse operation binding the contract event 0x87a332a414acbc7da074543639ce7ae02ff1ea72e88379da9f261b080beb5a13.
//
// Solidity: event Initiated(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseInitiated(log types.Log) (*DelayedVetoableInitiated, error) {
	event := new(DelayedVetoableInitiated)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Initiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedVetoableVetoedIterator is returned from FilterVetoed and is used to iterate over the raw logs and unpacked data for Vetoed events raised by the DelayedVetoable contract.
type DelayedVetoableVetoedIterator struct {
	Event *DelayedVetoableVetoed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelayedVetoableVetoedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedVetoableVetoed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelayedVetoableVetoed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelayedVetoableVetoedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedVetoableVetoedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedVetoableVetoed represents a Vetoed event raised by the DelayedVetoable contract.
type DelayedVetoableVetoed struct {
	CallHash [32]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVetoed is a free log retrieval operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) FilterVetoed(opts *bind.FilterOpts, callHash [][32]byte) (*DelayedVetoableVetoedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.FilterLogs(opts, "Vetoed", callHashRule)
	if err != nil {
		return nil, err
	}
	return &DelayedVetoableVetoedIterator{contract: _DelayedVetoable.contract, event: "Vetoed", logs: logs, sub: sub}, nil
}

// WatchVetoed is a free log subscription operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) WatchVetoed(opts *bind.WatchOpts, sink chan<- *DelayedVetoableVetoed, callHash [][32]byte) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}

	logs, sub, err := _DelayedVetoable.contract.WatchLogs(opts, "Vetoed", callHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedVetoableVetoed)
				if err := _DelayedVetoable.contract.UnpackLog(event, "Vetoed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVetoed is a log parse operation binding the contract event 0xbede6852c1d97d93ff557f676de76670cd0dec861e7fe8beb13aa0ba2b0ab040.
//
// Solidity: event Vetoed(bytes32 indexed callHash, bytes data)
func (_DelayedVetoable *DelayedVetoableFilterer) ParseVetoed(log types.Log) (*DelayedVetoableVetoed, error) {
	event := new(DelayedVetoableVetoed)
	if err := _DelayedVetoable.contract.UnpackLog(event, "Vetoed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
