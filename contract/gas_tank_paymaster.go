// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// GasTankPaymasterMetaData contains all meta data concerning the GasTankPaymaster contract.
var GasTankPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deduct\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"PostOp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"PostOpReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsedBeforePostOp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postOpGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceMarkupBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualTokenCost\",\"type\":\"uint256\"}],\"name\":\"SponsorReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"sponsorMode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTokenCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxGasCost\",\"type\":\"uint256\"}],\"name\":\"Validate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SPONSOR_MODE_CREDIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPONSOR_MODE_REFUND\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositTokenTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getPaymasterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSignerAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postOpGasOverhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceMarkupBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPostOpGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPriceMarkupBps\",\"type\":\"uint256\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawTokenTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GasTankPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use GasTankPaymasterMetaData.ABI instead.
var GasTankPaymasterABI = GasTankPaymasterMetaData.ABI

// GasTankPaymaster is an auto generated Go binding around an Ethereum contract.
type GasTankPaymaster struct {
	GasTankPaymasterCaller     // Read-only binding to the contract
	GasTankPaymasterTransactor // Write-only binding to the contract
	GasTankPaymasterFilterer   // Log filterer for contract events
}

// GasTankPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasTankPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTankPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasTankPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTankPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasTankPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasTankPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasTankPaymasterSession struct {
	Contract     *GasTankPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasTankPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasTankPaymasterCallerSession struct {
	Contract *GasTankPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GasTankPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasTankPaymasterTransactorSession struct {
	Contract     *GasTankPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GasTankPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasTankPaymasterRaw struct {
	Contract *GasTankPaymaster // Generic contract binding to access the raw methods on
}

// GasTankPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasTankPaymasterCallerRaw struct {
	Contract *GasTankPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// GasTankPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasTankPaymasterTransactorRaw struct {
	Contract *GasTankPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasTankPaymaster creates a new instance of GasTankPaymaster, bound to a specific deployed contract.
func NewGasTankPaymaster(address common.Address, backend bind.ContractBackend) (*GasTankPaymaster, error) {
	contract, err := bindGasTankPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymaster{GasTankPaymasterCaller: GasTankPaymasterCaller{contract: contract}, GasTankPaymasterTransactor: GasTankPaymasterTransactor{contract: contract}, GasTankPaymasterFilterer: GasTankPaymasterFilterer{contract: contract}}, nil
}

// NewGasTankPaymasterCaller creates a new read-only instance of GasTankPaymaster, bound to a specific deployed contract.
func NewGasTankPaymasterCaller(address common.Address, caller bind.ContractCaller) (*GasTankPaymasterCaller, error) {
	contract, err := bindGasTankPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterCaller{contract: contract}, nil
}

// NewGasTankPaymasterTransactor creates a new write-only instance of GasTankPaymaster, bound to a specific deployed contract.
func NewGasTankPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*GasTankPaymasterTransactor, error) {
	contract, err := bindGasTankPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterTransactor{contract: contract}, nil
}

// NewGasTankPaymasterFilterer creates a new log filterer instance of GasTankPaymaster, bound to a specific deployed contract.
func NewGasTankPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*GasTankPaymasterFilterer, error) {
	contract, err := bindGasTankPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterFilterer{contract: contract}, nil
}

// bindGasTankPaymaster binds a generic wrapper to an already deployed contract.
func bindGasTankPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasTankPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasTankPaymaster *GasTankPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasTankPaymaster.Contract.GasTankPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasTankPaymaster *GasTankPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.GasTankPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasTankPaymaster *GasTankPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.GasTankPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasTankPaymaster *GasTankPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasTankPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasTankPaymaster *GasTankPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasTankPaymaster *GasTankPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.contract.Transact(opts, method, params...)
}

// SPONSORMODECREDIT is a free data retrieval call binding the contract method 0x21cf78ca.
//
// Solidity: function SPONSOR_MODE_CREDIT() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterCaller) SPONSORMODECREDIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "SPONSOR_MODE_CREDIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SPONSORMODECREDIT is a free data retrieval call binding the contract method 0x21cf78ca.
//
// Solidity: function SPONSOR_MODE_CREDIT() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterSession) SPONSORMODECREDIT() (uint8, error) {
	return _GasTankPaymaster.Contract.SPONSORMODECREDIT(&_GasTankPaymaster.CallOpts)
}

// SPONSORMODECREDIT is a free data retrieval call binding the contract method 0x21cf78ca.
//
// Solidity: function SPONSOR_MODE_CREDIT() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) SPONSORMODECREDIT() (uint8, error) {
	return _GasTankPaymaster.Contract.SPONSORMODECREDIT(&_GasTankPaymaster.CallOpts)
}

// SPONSORMODEREFUND is a free data retrieval call binding the contract method 0x29ed1c50.
//
// Solidity: function SPONSOR_MODE_REFUND() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterCaller) SPONSORMODEREFUND(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "SPONSOR_MODE_REFUND")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SPONSORMODEREFUND is a free data retrieval call binding the contract method 0x29ed1c50.
//
// Solidity: function SPONSOR_MODE_REFUND() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterSession) SPONSORMODEREFUND() (uint8, error) {
	return _GasTankPaymaster.Contract.SPONSORMODEREFUND(&_GasTankPaymaster.CallOpts)
}

// SPONSORMODEREFUND is a free data retrieval call binding the contract method 0x29ed1c50.
//
// Solidity: function SPONSOR_MODE_REFUND() view returns(uint8)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) SPONSORMODEREFUND() (uint8, error) {
	return _GasTankPaymaster.Contract.SPONSORMODEREFUND(&_GasTankPaymaster.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "balanceOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) BalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GasTankPaymaster.Contract.BalanceOf(&_GasTankPaymaster.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) BalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GasTankPaymaster.Contract.BalanceOf(&_GasTankPaymaster.CallOpts, arg0, arg1)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterSession) EntryPoint() (common.Address, error) {
	return _GasTankPaymaster.Contract.EntryPoint(&_GasTankPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _GasTankPaymaster.Contract.EntryPoint(&_GasTankPaymaster.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) GetBalance() (*big.Int, error) {
	return _GasTankPaymaster.Contract.GetBalance(&_GasTankPaymaster.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) GetBalance() (*big.Int, error) {
	return _GasTankPaymaster.Contract.GetBalance(&_GasTankPaymaster.CallOpts)
}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_GasTankPaymaster *GasTankPaymasterCaller) GetPaymasterHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "getPaymasterHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_GasTankPaymaster *GasTankPaymasterSession) GetPaymasterHash(userOp PackedUserOperation) ([32]byte, error) {
	return _GasTankPaymaster.Contract.GetPaymasterHash(&_GasTankPaymaster.CallOpts, userOp)
}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) GetPaymasterHash(userOp PackedUserOperation) ([32]byte, error) {
	return _GasTankPaymaster.Contract.GetPaymasterHash(&_GasTankPaymaster.CallOpts, userOp)
}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCaller) IsSignerAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "isSignerAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterSession) IsSignerAllowed(arg0 common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsSignerAllowed(&_GasTankPaymaster.CallOpts, arg0)
}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) IsSignerAllowed(arg0 common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsSignerAllowed(&_GasTankPaymaster.CallOpts, arg0)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCaller) IsTokenAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "isTokenAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterSession) IsTokenAllowed(arg0 common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsTokenAllowed(&_GasTankPaymaster.CallOpts, arg0)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address ) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) IsTokenAllowed(arg0 common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsTokenAllowed(&_GasTankPaymaster.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterSession) Owner() (common.Address, error) {
	return _GasTankPaymaster.Contract.Owner(&_GasTankPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Owner() (common.Address, error) {
	return _GasTankPaymaster.Contract.Owner(&_GasTankPaymaster.CallOpts)
}

// PostOpGasOverhead is a free data retrieval call binding the contract method 0x6ec5f681.
//
// Solidity: function postOpGasOverhead() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) PostOpGasOverhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "postOpGasOverhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostOpGasOverhead is a free data retrieval call binding the contract method 0x6ec5f681.
//
// Solidity: function postOpGasOverhead() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) PostOpGasOverhead() (*big.Int, error) {
	return _GasTankPaymaster.Contract.PostOpGasOverhead(&_GasTankPaymaster.CallOpts)
}

// PostOpGasOverhead is a free data retrieval call binding the contract method 0x6ec5f681.
//
// Solidity: function postOpGasOverhead() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) PostOpGasOverhead() (*big.Int, error) {
	return _GasTankPaymaster.Contract.PostOpGasOverhead(&_GasTankPaymaster.CallOpts)
}

// PriceMarkupBps is a free data retrieval call binding the contract method 0x627b9b51.
//
// Solidity: function priceMarkupBps() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) PriceMarkupBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "priceMarkupBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceMarkupBps is a free data retrieval call binding the contract method 0x627b9b51.
//
// Solidity: function priceMarkupBps() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) PriceMarkupBps() (*big.Int, error) {
	return _GasTankPaymaster.Contract.PriceMarkupBps(&_GasTankPaymaster.CallOpts)
}

// PriceMarkupBps is a free data retrieval call binding the contract method 0x627b9b51.
//
// Solidity: function priceMarkupBps() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) PriceMarkupBps() (*big.Int, error) {
	return _GasTankPaymaster.Contract.PriceMarkupBps(&_GasTankPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasTankPaymaster *GasTankPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.AddStake(&_GasTankPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.AddStake(&_GasTankPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasTankPaymaster *GasTankPaymasterSession) Deposit() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Deposit(&_GasTankPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Deposit(&_GasTankPaymaster.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) DepositToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "depositToken", token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.DepositToken(&_GasTankPaymaster.TransactOpts, token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.DepositToken(&_GasTankPaymaster.TransactOpts, token, amount)
}

// DepositTokenTo is a paid mutator transaction binding the contract method 0x72d6db07.
//
// Solidity: function depositTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) DepositTokenTo(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "depositTokenTo", token, amount, recipient)
}

// DepositTokenTo is a paid mutator transaction binding the contract method 0x72d6db07.
//
// Solidity: function depositTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) DepositTokenTo(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.DepositTokenTo(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
}

// DepositTokenTo is a paid mutator transaction binding the contract method 0x72d6db07.
//
// Solidity: function depositTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) DepositTokenTo(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.DepositTokenTo(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.PostOp(&_GasTankPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.PostOp(&_GasTankPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasTankPaymaster *GasTankPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.RenounceOwnership(&_GasTankPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.RenounceOwnership(&_GasTankPaymaster.TransactOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x1e34c585.
//
// Solidity: function setConfig(uint256 newPostOpGasOverhead, uint256 newPriceMarkupBps) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) SetConfig(opts *bind.TransactOpts, newPostOpGasOverhead *big.Int, newPriceMarkupBps *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "setConfig", newPostOpGasOverhead, newPriceMarkupBps)
}

// SetConfig is a paid mutator transaction binding the contract method 0x1e34c585.
//
// Solidity: function setConfig(uint256 newPostOpGasOverhead, uint256 newPriceMarkupBps) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) SetConfig(newPostOpGasOverhead *big.Int, newPriceMarkupBps *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetConfig(&_GasTankPaymaster.TransactOpts, newPostOpGasOverhead, newPriceMarkupBps)
}

// SetConfig is a paid mutator transaction binding the contract method 0x1e34c585.
//
// Solidity: function setConfig(uint256 newPostOpGasOverhead, uint256 newPriceMarkupBps) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) SetConfig(newPostOpGasOverhead *big.Int, newPriceMarkupBps *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetConfig(&_GasTankPaymaster.TransactOpts, newPostOpGasOverhead, newPriceMarkupBps)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) SetSigner(opts *bind.TransactOpts, signer common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "setSigner", signer, allowed)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) SetSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetSigner(&_GasTankPaymaster.TransactOpts, signer, allowed)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) SetSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetSigner(&_GasTankPaymaster.TransactOpts, signer, allowed)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) SetToken(opts *bind.TransactOpts, token common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "setToken", token, allowed)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) SetToken(token common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetToken(&_GasTankPaymaster.TransactOpts, token, allowed)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool allowed) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) SetToken(token common.Address, allowed bool) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetToken(&_GasTankPaymaster.TransactOpts, token, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.TransferOwnership(&_GasTankPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.TransferOwnership(&_GasTankPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasTankPaymaster *GasTankPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.UnlockStake(&_GasTankPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.UnlockStake(&_GasTankPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasTankPaymaster *GasTankPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasTankPaymaster *GasTankPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.ValidatePaymasterUserOp(&_GasTankPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.ValidatePaymasterUserOp(&_GasTankPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawStake(&_GasTankPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawStake(&_GasTankPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTo(&_GasTankPaymaster.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTo(&_GasTankPaymaster.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawToken(&_GasTankPaymaster.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawToken(&_GasTankPaymaster.TransactOpts, token, amount)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x54ad4179.
//
// Solidity: function withdrawTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawTokenTo(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawTokenTo", token, amount, recipient)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x54ad4179.
//
// Solidity: function withdrawTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawTokenTo(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTokenTo(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x54ad4179.
//
// Solidity: function withdrawTokenTo(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawTokenTo(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTokenTo(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasTankPaymaster *GasTankPaymasterSession) Receive() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Receive(&_GasTankPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Receive(&_GasTankPaymaster.TransactOpts)
}

// GasTankPaymasterBadDebtIterator is returned from FilterBadDebt and is used to iterate over the raw logs and unpacked data for BadDebt events raised by the GasTankPaymaster contract.
type GasTankPaymasterBadDebtIterator struct {
	Event *GasTankPaymasterBadDebt // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterBadDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterBadDebt)
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
		it.Event = new(GasTankPaymasterBadDebt)
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
func (it *GasTankPaymasterBadDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterBadDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterBadDebt represents a BadDebt event raised by the GasTankPaymaster contract.
type GasTankPaymasterBadDebt struct {
	UserOpHash [32]byte
	Sender     common.Address
	Token      common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBadDebt is a free log retrieval operation binding the contract event 0x6ce6649cba7d6b40ab14e94a265681110ba823079f09819a9f0e9e07aafc7a35.
//
// Solidity: event BadDebt(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterBadDebt(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, token []common.Address) (*GasTankPaymasterBadDebtIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "BadDebt", userOpHashRule, senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterBadDebtIterator{contract: _GasTankPaymaster.contract, event: "BadDebt", logs: logs, sub: sub}, nil
}

// WatchBadDebt is a free log subscription operation binding the contract event 0x6ce6649cba7d6b40ab14e94a265681110ba823079f09819a9f0e9e07aafc7a35.
//
// Solidity: event BadDebt(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchBadDebt(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterBadDebt, userOpHash [][32]byte, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "BadDebt", userOpHashRule, senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterBadDebt)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "BadDebt", log); err != nil {
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

// ParseBadDebt is a log parse operation binding the contract event 0x6ce6649cba7d6b40ab14e94a265681110ba823079f09819a9f0e9e07aafc7a35.
//
// Solidity: event BadDebt(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseBadDebt(log types.Log) (*GasTankPaymasterBadDebt, error) {
	event := new(GasTankPaymasterBadDebt)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "BadDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterDeductIterator is returned from FilterDeduct and is used to iterate over the raw logs and unpacked data for Deduct events raised by the GasTankPaymaster contract.
type GasTankPaymasterDeductIterator struct {
	Event *GasTankPaymasterDeduct // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterDeductIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterDeduct)
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
		it.Event = new(GasTankPaymasterDeduct)
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
func (it *GasTankPaymasterDeductIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterDeductIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterDeduct represents a Deduct event raised by the GasTankPaymaster contract.
type GasTankPaymasterDeduct struct {
	Sender  common.Address
	Token   common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeduct is a free log retrieval operation binding the contract event 0xe5c3f657921839006d873b44df8feececaee703cad36110006c25f71b47222bd.
//
// Solidity: event Deduct(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterDeduct(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*GasTankPaymasterDeductIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Deduct", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterDeductIterator{contract: _GasTankPaymaster.contract, event: "Deduct", logs: logs, sub: sub}, nil
}

// WatchDeduct is a free log subscription operation binding the contract event 0xe5c3f657921839006d873b44df8feececaee703cad36110006c25f71b47222bd.
//
// Solidity: event Deduct(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchDeduct(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterDeduct, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Deduct", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterDeduct)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Deduct", log); err != nil {
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

// ParseDeduct is a log parse operation binding the contract event 0xe5c3f657921839006d873b44df8feececaee703cad36110006c25f71b47222bd.
//
// Solidity: event Deduct(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseDeduct(log types.Log) (*GasTankPaymasterDeduct, error) {
	event := new(GasTankPaymasterDeduct)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Deduct", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GasTankPaymaster contract.
type GasTankPaymasterDepositIterator struct {
	Event *GasTankPaymasterDeposit // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterDeposit)
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
		it.Event = new(GasTankPaymasterDeposit)
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
func (it *GasTankPaymasterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterDeposit represents a Deposit event raised by the GasTankPaymaster contract.
type GasTankPaymasterDeposit struct {
	Operator  common.Address
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Balance   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5fe47ed6d4225326d3303476197d782ded5a4e9c14f479dc9ec4992af4e85d59.
//
// Solidity: event Deposit(address indexed operator, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterDeposit(opts *bind.FilterOpts, operator []common.Address, token []common.Address, recipient []common.Address) (*GasTankPaymasterDepositIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Deposit", operatorRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterDepositIterator{contract: _GasTankPaymaster.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5fe47ed6d4225326d3303476197d782ded5a4e9c14f479dc9ec4992af4e85d59.
//
// Solidity: event Deposit(address indexed operator, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterDeposit, operator []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Deposit", operatorRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterDeposit)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5fe47ed6d4225326d3303476197d782ded5a4e9c14f479dc9ec4992af4e85d59.
//
// Solidity: event Deposit(address indexed operator, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseDeposit(log types.Log) (*GasTankPaymasterDeposit, error) {
	event := new(GasTankPaymasterDeposit)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasTankPaymaster contract.
type GasTankPaymasterOwnershipTransferredIterator struct {
	Event *GasTankPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterOwnershipTransferred)
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
		it.Event = new(GasTankPaymasterOwnershipTransferred)
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
func (it *GasTankPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the GasTankPaymaster contract.
type GasTankPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasTankPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterOwnershipTransferredIterator{contract: _GasTankPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterOwnershipTransferred)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*GasTankPaymasterOwnershipTransferred, error) {
	event := new(GasTankPaymasterOwnershipTransferred)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterPostOpIterator is returned from FilterPostOp and is used to iterate over the raw logs and unpacked data for PostOp events raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOpIterator struct {
	Event *GasTankPaymasterPostOp // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterPostOpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterPostOp)
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
		it.Event = new(GasTankPaymasterPostOp)
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
func (it *GasTankPaymasterPostOpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterPostOpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterPostOp represents a PostOp event raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOp struct {
	UserOpHash            [32]byte
	Sender                common.Address
	Mode                  uint8
	ActualGasCost         *big.Int
	ActualUserOpFeePerGas *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPostOp is a free log retrieval operation binding the contract event 0x0f337296fc63a313242c3bf5647e81d9eca3511978e568c5543d8ce00226c595.
//
// Solidity: event PostOp(bytes32 indexed userOpHash, address indexed sender, uint8 mode, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterPostOp(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GasTankPaymasterPostOpIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "PostOp", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterPostOpIterator{contract: _GasTankPaymaster.contract, event: "PostOp", logs: logs, sub: sub}, nil
}

// WatchPostOp is a free log subscription operation binding the contract event 0x0f337296fc63a313242c3bf5647e81d9eca3511978e568c5543d8ce00226c595.
//
// Solidity: event PostOp(bytes32 indexed userOpHash, address indexed sender, uint8 mode, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchPostOp(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterPostOp, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "PostOp", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterPostOp)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOp", log); err != nil {
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

// ParsePostOp is a log parse operation binding the contract event 0x0f337296fc63a313242c3bf5647e81d9eca3511978e568c5543d8ce00226c595.
//
// Solidity: event PostOp(bytes32 indexed userOpHash, address indexed sender, uint8 mode, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParsePostOp(log types.Log) (*GasTankPaymasterPostOp, error) {
	event := new(GasTankPaymasterPostOp)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterPostOpRevertedIterator is returned from FilterPostOpReverted and is used to iterate over the raw logs and unpacked data for PostOpReverted events raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOpRevertedIterator struct {
	Event *GasTankPaymasterPostOpReverted // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterPostOpRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterPostOpReverted)
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
		it.Event = new(GasTankPaymasterPostOpReverted)
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
func (it *GasTankPaymasterPostOpRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterPostOpRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterPostOpReverted represents a PostOpReverted event raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOpReverted struct {
	UserOpHash            [32]byte
	Sender                common.Address
	ActualGasCost         *big.Int
	ActualUserOpFeePerGas *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPostOpReverted is a free log retrieval operation binding the contract event 0x305c3003dbdd1d2ddb85afedcc7c9414a770cff42aa07a4017540ba2c58718e4.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, address indexed sender, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterPostOpReverted(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*GasTankPaymasterPostOpRevertedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "PostOpReverted", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterPostOpRevertedIterator{contract: _GasTankPaymaster.contract, event: "PostOpReverted", logs: logs, sub: sub}, nil
}

// WatchPostOpReverted is a free log subscription operation binding the contract event 0x305c3003dbdd1d2ddb85afedcc7c9414a770cff42aa07a4017540ba2c58718e4.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, address indexed sender, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchPostOpReverted(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterPostOpReverted, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "PostOpReverted", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterPostOpReverted)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOpReverted", log); err != nil {
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

// ParsePostOpReverted is a log parse operation binding the contract event 0x305c3003dbdd1d2ddb85afedcc7c9414a770cff42aa07a4017540ba2c58718e4.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, address indexed sender, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParsePostOpReverted(log types.Log) (*GasTankPaymasterPostOpReverted, error) {
	event := new(GasTankPaymasterPostOpReverted)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOpReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the GasTankPaymaster contract.
type GasTankPaymasterRefundIterator struct {
	Event *GasTankPaymasterRefund // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterRefund)
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
		it.Event = new(GasTankPaymasterRefund)
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
func (it *GasTankPaymasterRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterRefund represents a Refund event raised by the GasTankPaymaster contract.
type GasTankPaymasterRefund struct {
	Sender  common.Address
	Token   common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x82c4addd7df9bb5b801dcdeb0a67eb8bda3d9e213af78965d584b0be8cb63660.
//
// Solidity: event Refund(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterRefund(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*GasTankPaymasterRefundIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Refund", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterRefundIterator{contract: _GasTankPaymaster.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x82c4addd7df9bb5b801dcdeb0a67eb8bda3d9e213af78965d584b0be8cb63660.
//
// Solidity: event Refund(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterRefund, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Refund", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterRefund)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x82c4addd7df9bb5b801dcdeb0a67eb8bda3d9e213af78965d584b0be8cb63660.
//
// Solidity: event Refund(address indexed sender, address indexed token, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseRefund(log types.Log) (*GasTankPaymasterRefund, error) {
	event := new(GasTankPaymasterRefund)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the GasTankPaymaster contract.
type GasTankPaymasterSignerUpdatedIterator struct {
	Event *GasTankPaymasterSignerUpdated // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterSignerUpdated)
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
		it.Event = new(GasTankPaymasterSignerUpdated)
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
func (it *GasTankPaymasterSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterSignerUpdated represents a SignerUpdated event raised by the GasTankPaymaster contract.
type GasTankPaymasterSignerUpdated struct {
	Signer  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0xfcaa24b1276bfa7dbf77797c0a984b9df924acbeaabd48cd2f1b0eca379b78fa.
//
// Solidity: event SignerUpdated(address indexed signer, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterSignerUpdated(opts *bind.FilterOpts, signer []common.Address) (*GasTankPaymasterSignerUpdatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "SignerUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterSignerUpdatedIterator{contract: _GasTankPaymaster.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0xfcaa24b1276bfa7dbf77797c0a984b9df924acbeaabd48cd2f1b0eca379b78fa.
//
// Solidity: event SignerUpdated(address indexed signer, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterSignerUpdated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "SignerUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterSignerUpdated)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0xfcaa24b1276bfa7dbf77797c0a984b9df924acbeaabd48cd2f1b0eca379b78fa.
//
// Solidity: event SignerUpdated(address indexed signer, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseSignerUpdated(log types.Log) (*GasTankPaymasterSignerUpdated, error) {
	event := new(GasTankPaymasterSignerUpdated)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterSponsorReceiptIterator is returned from FilterSponsorReceipt and is used to iterate over the raw logs and unpacked data for SponsorReceipt events raised by the GasTankPaymaster contract.
type GasTankPaymasterSponsorReceiptIterator struct {
	Event *GasTankPaymasterSponsorReceipt // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterSponsorReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterSponsorReceipt)
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
		it.Event = new(GasTankPaymasterSponsorReceipt)
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
func (it *GasTankPaymasterSponsorReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterSponsorReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterSponsorReceipt represents a SponsorReceipt event raised by the GasTankPaymaster contract.
type GasTankPaymasterSponsorReceipt struct {
	ActualGasUsedBeforePostOp *big.Int
	ActualGasPrice            *big.Int
	PostOpGas                 *big.Int
	PriceMarkupBps            *big.Int
	ActualTokenCost           *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSponsorReceipt is a free log retrieval operation binding the contract event 0x3c59d84d6a99ce741141de0a99b6fd31a9e55a2c8ee45f3a4435e19e9b04d0a4.
//
// Solidity: event SponsorReceipt(uint256 actualGasUsedBeforePostOp, uint256 actualGasPrice, uint256 postOpGas, uint256 priceMarkupBps, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterSponsorReceipt(opts *bind.FilterOpts) (*GasTankPaymasterSponsorReceiptIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "SponsorReceipt")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterSponsorReceiptIterator{contract: _GasTankPaymaster.contract, event: "SponsorReceipt", logs: logs, sub: sub}, nil
}

// WatchSponsorReceipt is a free log subscription operation binding the contract event 0x3c59d84d6a99ce741141de0a99b6fd31a9e55a2c8ee45f3a4435e19e9b04d0a4.
//
// Solidity: event SponsorReceipt(uint256 actualGasUsedBeforePostOp, uint256 actualGasPrice, uint256 postOpGas, uint256 priceMarkupBps, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchSponsorReceipt(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterSponsorReceipt) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "SponsorReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterSponsorReceipt)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "SponsorReceipt", log); err != nil {
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

// ParseSponsorReceipt is a log parse operation binding the contract event 0x3c59d84d6a99ce741141de0a99b6fd31a9e55a2c8ee45f3a4435e19e9b04d0a4.
//
// Solidity: event SponsorReceipt(uint256 actualGasUsedBeforePostOp, uint256 actualGasPrice, uint256 postOpGas, uint256 priceMarkupBps, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseSponsorReceipt(log types.Log) (*GasTankPaymasterSponsorReceipt, error) {
	event := new(GasTankPaymasterSponsorReceipt)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "SponsorReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterTokenUpdatedIterator is returned from FilterTokenUpdated and is used to iterate over the raw logs and unpacked data for TokenUpdated events raised by the GasTankPaymaster contract.
type GasTankPaymasterTokenUpdatedIterator struct {
	Event *GasTankPaymasterTokenUpdated // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterTokenUpdated)
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
		it.Event = new(GasTankPaymasterTokenUpdated)
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
func (it *GasTankPaymasterTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterTokenUpdated represents a TokenUpdated event raised by the GasTankPaymaster contract.
type GasTankPaymasterTokenUpdated struct {
	Token   common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdated is a free log retrieval operation binding the contract event 0xdcb2804db02b95bdd568fd11a31c5577ffdf36538c0f670e92930d9c1e8518ab.
//
// Solidity: event TokenUpdated(address indexed token, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*GasTankPaymasterTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterTokenUpdatedIterator{contract: _GasTankPaymaster.contract, event: "TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenUpdated is a free log subscription operation binding the contract event 0xdcb2804db02b95bdd568fd11a31c5577ffdf36538c0f670e92930d9c1e8518ab.
//
// Solidity: event TokenUpdated(address indexed token, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchTokenUpdated(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterTokenUpdated)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
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

// ParseTokenUpdated is a log parse operation binding the contract event 0xdcb2804db02b95bdd568fd11a31c5577ffdf36538c0f670e92930d9c1e8518ab.
//
// Solidity: event TokenUpdated(address indexed token, bool allowed)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseTokenUpdated(log types.Log) (*GasTankPaymasterTokenUpdated, error) {
	event := new(GasTankPaymasterTokenUpdated)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterValidateIterator is returned from FilterValidate and is used to iterate over the raw logs and unpacked data for Validate events raised by the GasTankPaymaster contract.
type GasTankPaymasterValidateIterator struct {
	Event *GasTankPaymasterValidate // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterValidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterValidate)
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
		it.Event = new(GasTankPaymasterValidate)
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
func (it *GasTankPaymasterValidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterValidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterValidate represents a Validate event raised by the GasTankPaymaster contract.
type GasTankPaymasterValidate struct {
	SponsorMode  uint8
	Token        common.Address
	MaxTokenCost *big.Int
	MaxGasCost   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidate is a free log retrieval operation binding the contract event 0x84498fb0f390f47c296ca89b9abfa7e3f50af93d45399e0c23c47af0c83019e0.
//
// Solidity: event Validate(uint8 indexed sponsorMode, address indexed token, uint256 maxTokenCost, uint256 maxGasCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterValidate(opts *bind.FilterOpts, sponsorMode []uint8, token []common.Address) (*GasTankPaymasterValidateIterator, error) {

	var sponsorModeRule []interface{}
	for _, sponsorModeItem := range sponsorMode {
		sponsorModeRule = append(sponsorModeRule, sponsorModeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Validate", sponsorModeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterValidateIterator{contract: _GasTankPaymaster.contract, event: "Validate", logs: logs, sub: sub}, nil
}

// WatchValidate is a free log subscription operation binding the contract event 0x84498fb0f390f47c296ca89b9abfa7e3f50af93d45399e0c23c47af0c83019e0.
//
// Solidity: event Validate(uint8 indexed sponsorMode, address indexed token, uint256 maxTokenCost, uint256 maxGasCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchValidate(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterValidate, sponsorMode []uint8, token []common.Address) (event.Subscription, error) {

	var sponsorModeRule []interface{}
	for _, sponsorModeItem := range sponsorMode {
		sponsorModeRule = append(sponsorModeRule, sponsorModeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Validate", sponsorModeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterValidate)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Validate", log); err != nil {
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

// ParseValidate is a log parse operation binding the contract event 0x84498fb0f390f47c296ca89b9abfa7e3f50af93d45399e0c23c47af0c83019e0.
//
// Solidity: event Validate(uint8 indexed sponsorMode, address indexed token, uint256 maxTokenCost, uint256 maxGasCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseValidate(log types.Log) (*GasTankPaymasterValidate, error) {
	event := new(GasTankPaymasterValidate)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Validate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawIterator struct {
	Event *GasTankPaymasterWithdraw // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterWithdraw)
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
		it.Event = new(GasTankPaymasterWithdraw)
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
func (it *GasTankPaymasterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterWithdraw represents a Withdraw event raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdraw struct {
	Sender    common.Address
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Balance   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, token []common.Address, recipient []common.Address) (*GasTankPaymasterWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Withdraw", senderRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterWithdrawIterator{contract: _GasTankPaymaster.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterWithdraw, sender []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Withdraw", senderRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterWithdraw)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed token, address indexed recipient, uint256 amount, uint256 balance)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseWithdraw(log types.Log) (*GasTankPaymasterWithdraw, error) {
	event := new(GasTankPaymasterWithdraw)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
