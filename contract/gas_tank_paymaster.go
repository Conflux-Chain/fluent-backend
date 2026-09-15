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

// // PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
// type PackedUserOperation struct {
// 	Sender             common.Address
// 	Nonce              *big.Int
// 	InitCode           []byte
// 	CallData           []byte
// 	AccountGasLimits   [32]byte
// 	PreVerificationGas *big.Int
// 	GasFees            [32]byte
// 	PaymasterAndData   []byte
// 	Signature          []byte
// }

// GasTankPaymasterMetaData contains all meta data concerning the GasTankPaymaster contract.
var GasTankPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deduct\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPostOpGasOverhead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPostOpGasOverhead\",\"type\":\"uint256\"}],\"name\":\"PostOpGasOverheadUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"PostOpReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTokenCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postOpGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualTokenCost\",\"type\":\"uint256\"}],\"name\":\"Sponsored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"WithdrawalCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"oldDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"WithdrawalDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"withdrawableAt\",\"type\":\"uint48\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint208\",\"name\":\"withdrawAmount\",\"type\":\"uint208\"},{\"internalType\":\"uint48\",\"name\":\"withdrawableAt\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"cancelWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositTokenTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getPaymasterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSignerAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"payment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postOpGasOverhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"amount\",\"type\":\"uint208\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPostOpGasOverhead\",\"type\":\"uint256\"}],\"name\":\"setPostOpGasOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"setWithdrawalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawTokenTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Accounts is a free data retrieval call binding the contract method 0xad74b775.
//
// Solidity: function accounts(address account, address token) view returns(uint256 balance, uint208 withdrawAmount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterCaller) Accounts(opts *bind.CallOpts, account common.Address, token common.Address) (struct {
	Balance        *big.Int
	WithdrawAmount *big.Int
	WithdrawableAt *big.Int
}, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "accounts", account, token)

	outstruct := new(struct {
		Balance        *big.Int
		WithdrawAmount *big.Int
		WithdrawableAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawableAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0xad74b775.
//
// Solidity: function accounts(address account, address token) view returns(uint256 balance, uint208 withdrawAmount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterSession) Accounts(account common.Address, token common.Address) (struct {
	Balance        *big.Int
	WithdrawAmount *big.Int
	WithdrawableAt *big.Int
}, error) {
	return _GasTankPaymaster.Contract.Accounts(&_GasTankPaymaster.CallOpts, account, token)
}

// Accounts is a free data retrieval call binding the contract method 0xad74b775.
//
// Solidity: function accounts(address account, address token) view returns(uint256 balance, uint208 withdrawAmount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Accounts(account common.Address, token common.Address) (struct {
	Balance        *big.Int
	WithdrawAmount *big.Int
	WithdrawableAt *big.Int
}, error) {
	return _GasTankPaymaster.Contract.Accounts(&_GasTankPaymaster.CallOpts, account, token)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) Balance() (*big.Int, error) {
	return _GasTankPaymaster.Contract.Balance(&_GasTankPaymaster.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Balance() (*big.Int, error) {
	return _GasTankPaymaster.Contract.Balance(&_GasTankPaymaster.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GasTankPaymaster *GasTankPaymasterCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GasTankPaymaster *GasTankPaymasterSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GasTankPaymaster.Contract.Eip712Domain(&_GasTankPaymaster.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GasTankPaymaster.Contract.Eip712Domain(&_GasTankPaymaster.CallOpts)
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
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCaller) IsTokenAllowed(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "isTokenAllowed", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsTokenAllowed(&_GasTankPaymaster.CallOpts, token)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _GasTankPaymaster.Contract.IsTokenAllowed(&_GasTankPaymaster.CallOpts, token)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterSession) Paused() (bool, error) {
	return _GasTankPaymaster.Contract.Paused(&_GasTankPaymaster.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Paused() (bool, error) {
	return _GasTankPaymaster.Contract.Paused(&_GasTankPaymaster.CallOpts)
}

// Payment is a free data retrieval call binding the contract method 0x3b92f3df.
//
// Solidity: function payment(address token) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCaller) Payment(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "payment", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Payment is a free data retrieval call binding the contract method 0x3b92f3df.
//
// Solidity: function payment(address token) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterSession) Payment(token common.Address) (*big.Int, error) {
	return _GasTankPaymaster.Contract.Payment(&_GasTankPaymaster.CallOpts, token)
}

// Payment is a free data retrieval call binding the contract method 0x3b92f3df.
//
// Solidity: function payment(address token) view returns(uint256)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) Payment(token common.Address) (*big.Int, error) {
	return _GasTankPaymaster.Contract.Payment(&_GasTankPaymaster.CallOpts, token)
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

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint48)
func (_GasTankPaymaster *GasTankPaymasterCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasTankPaymaster.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint48)
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawalDelay() (*big.Int, error) {
	return _GasTankPaymaster.Contract.WithdrawalDelay(&_GasTankPaymaster.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint48)
func (_GasTankPaymaster *GasTankPaymasterCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _GasTankPaymaster.Contract.WithdrawalDelay(&_GasTankPaymaster.CallOpts)
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

// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
//
// Solidity: function cancelWithdraw(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) CancelWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "cancelWithdraw", token)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
//
// Solidity: function cancelWithdraw(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) CancelWithdraw(token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.CancelWithdraw(&_GasTankPaymaster.TransactOpts, token)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0xe9919629.
//
// Solidity: function cancelWithdraw(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) CancelWithdraw(token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.CancelWithdraw(&_GasTankPaymaster.TransactOpts, token)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasTankPaymaster *GasTankPaymasterSession) Pause() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Pause(&_GasTankPaymaster.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) Pause() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Pause(&_GasTankPaymaster.TransactOpts)
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

// RequestWithdraw is a paid mutator transaction binding the contract method 0x0d970d6f.
//
// Solidity: function requestWithdraw(address token, uint208 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) RequestWithdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "requestWithdraw", token, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x0d970d6f.
//
// Solidity: function requestWithdraw(address token, uint208 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) RequestWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.RequestWithdraw(&_GasTankPaymaster.TransactOpts, token, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x0d970d6f.
//
// Solidity: function requestWithdraw(address token, uint208 amount) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) RequestWithdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.RequestWithdraw(&_GasTankPaymaster.TransactOpts, token, amount)
}

// SetPostOpGasOverhead is a paid mutator transaction binding the contract method 0xfb79777d.
//
// Solidity: function setPostOpGasOverhead(uint256 newPostOpGasOverhead) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) SetPostOpGasOverhead(opts *bind.TransactOpts, newPostOpGasOverhead *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "setPostOpGasOverhead", newPostOpGasOverhead)
}

// SetPostOpGasOverhead is a paid mutator transaction binding the contract method 0xfb79777d.
//
// Solidity: function setPostOpGasOverhead(uint256 newPostOpGasOverhead) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) SetPostOpGasOverhead(newPostOpGasOverhead *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetPostOpGasOverhead(&_GasTankPaymaster.TransactOpts, newPostOpGasOverhead)
}

// SetPostOpGasOverhead is a paid mutator transaction binding the contract method 0xfb79777d.
//
// Solidity: function setPostOpGasOverhead(uint256 newPostOpGasOverhead) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) SetPostOpGasOverhead(newPostOpGasOverhead *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetPostOpGasOverhead(&_GasTankPaymaster.TransactOpts, newPostOpGasOverhead)
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

// SetWithdrawalDelay is a paid mutator transaction binding the contract method 0x4baaee46.
//
// Solidity: function setWithdrawalDelay(uint48 newDelay) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) SetWithdrawalDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "setWithdrawalDelay", newDelay)
}

// SetWithdrawalDelay is a paid mutator transaction binding the contract method 0x4baaee46.
//
// Solidity: function setWithdrawalDelay(uint48 newDelay) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) SetWithdrawalDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetWithdrawalDelay(&_GasTankPaymaster.TransactOpts, newDelay)
}

// SetWithdrawalDelay is a paid mutator transaction binding the contract method 0x4baaee46.
//
// Solidity: function setWithdrawalDelay(uint48 newDelay) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) SetWithdrawalDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.SetWithdrawalDelay(&_GasTankPaymaster.TransactOpts, newDelay)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasTankPaymaster *GasTankPaymasterSession) Unpause() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Unpause(&_GasTankPaymaster.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) Unpause() (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.Unpause(&_GasTankPaymaster.TransactOpts)
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

// WithdrawPayment is a paid mutator transaction binding the contract method 0xfa192d99.
//
// Solidity: function withdrawPayment(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawPayment(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawPayment", token, amount, recipient)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xfa192d99.
//
// Solidity: function withdrawPayment(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawPayment(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawPayment(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0xfa192d99.
//
// Solidity: function withdrawPayment(address token, uint256 amount, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawPayment(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawPayment(&_GasTankPaymaster.TransactOpts, token, amount, recipient)
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

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawToken", token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawToken(&_GasTankPaymaster.TransactOpts, token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawToken(&_GasTankPaymaster.TransactOpts, token)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0xf19fe69b.
//
// Solidity: function withdrawTokenTo(address token, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactor) WithdrawTokenTo(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.contract.Transact(opts, "withdrawTokenTo", token, recipient)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0xf19fe69b.
//
// Solidity: function withdrawTokenTo(address token, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterSession) WithdrawTokenTo(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTokenTo(&_GasTankPaymaster.TransactOpts, token, recipient)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0xf19fe69b.
//
// Solidity: function withdrawTokenTo(address token, address recipient) returns()
func (_GasTankPaymaster *GasTankPaymasterTransactorSession) WithdrawTokenTo(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _GasTankPaymaster.Contract.WithdrawTokenTo(&_GasTankPaymaster.TransactOpts, token, recipient)
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

// GasTankPaymasterEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GasTankPaymaster contract.
type GasTankPaymasterEIP712DomainChangedIterator struct {
	Event *GasTankPaymasterEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterEIP712DomainChanged)
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
		it.Event = new(GasTankPaymasterEIP712DomainChanged)
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
func (it *GasTankPaymasterEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterEIP712DomainChanged represents a EIP712DomainChanged event raised by the GasTankPaymaster contract.
type GasTankPaymasterEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GasTankPaymasterEIP712DomainChangedIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterEIP712DomainChangedIterator{contract: _GasTankPaymaster.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterEIP712DomainChanged)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseEIP712DomainChanged(log types.Log) (*GasTankPaymasterEIP712DomainChanged, error) {
	event := new(GasTankPaymasterEIP712DomainChanged)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// GasTankPaymasterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GasTankPaymaster contract.
type GasTankPaymasterPausedIterator struct {
	Event *GasTankPaymasterPaused // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterPaused)
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
		it.Event = new(GasTankPaymasterPaused)
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
func (it *GasTankPaymasterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterPaused represents a Paused event raised by the GasTankPaymaster contract.
type GasTankPaymasterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterPaused(opts *bind.FilterOpts) (*GasTankPaymasterPausedIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterPausedIterator{contract: _GasTankPaymaster.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterPaused) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterPaused)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParsePaused(log types.Log) (*GasTankPaymasterPaused, error) {
	event := new(GasTankPaymasterPaused)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterPostOpGasOverheadUpdatedIterator is returned from FilterPostOpGasOverheadUpdated and is used to iterate over the raw logs and unpacked data for PostOpGasOverheadUpdated events raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOpGasOverheadUpdatedIterator struct {
	Event *GasTankPaymasterPostOpGasOverheadUpdated // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterPostOpGasOverheadUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterPostOpGasOverheadUpdated)
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
		it.Event = new(GasTankPaymasterPostOpGasOverheadUpdated)
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
func (it *GasTankPaymasterPostOpGasOverheadUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterPostOpGasOverheadUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterPostOpGasOverheadUpdated represents a PostOpGasOverheadUpdated event raised by the GasTankPaymaster contract.
type GasTankPaymasterPostOpGasOverheadUpdated struct {
	OldPostOpGasOverhead *big.Int
	NewPostOpGasOverhead *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPostOpGasOverheadUpdated is a free log retrieval operation binding the contract event 0x4864bffd33a5e428131277b247b43e552b814f2eee7eba6b68a3ebcc45a23fbd.
//
// Solidity: event PostOpGasOverheadUpdated(uint256 oldPostOpGasOverhead, uint256 newPostOpGasOverhead)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterPostOpGasOverheadUpdated(opts *bind.FilterOpts) (*GasTankPaymasterPostOpGasOverheadUpdatedIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "PostOpGasOverheadUpdated")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterPostOpGasOverheadUpdatedIterator{contract: _GasTankPaymaster.contract, event: "PostOpGasOverheadUpdated", logs: logs, sub: sub}, nil
}

// WatchPostOpGasOverheadUpdated is a free log subscription operation binding the contract event 0x4864bffd33a5e428131277b247b43e552b814f2eee7eba6b68a3ebcc45a23fbd.
//
// Solidity: event PostOpGasOverheadUpdated(uint256 oldPostOpGasOverhead, uint256 newPostOpGasOverhead)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchPostOpGasOverheadUpdated(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterPostOpGasOverheadUpdated) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "PostOpGasOverheadUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterPostOpGasOverheadUpdated)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOpGasOverheadUpdated", log); err != nil {
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

// ParsePostOpGasOverheadUpdated is a log parse operation binding the contract event 0x4864bffd33a5e428131277b247b43e552b814f2eee7eba6b68a3ebcc45a23fbd.
//
// Solidity: event PostOpGasOverheadUpdated(uint256 oldPostOpGasOverhead, uint256 newPostOpGasOverhead)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParsePostOpGasOverheadUpdated(log types.Log) (*GasTankPaymasterPostOpGasOverheadUpdated, error) {
	event := new(GasTankPaymasterPostOpGasOverheadUpdated)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "PostOpGasOverheadUpdated", log); err != nil {
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
	ActualGasCost         *big.Int
	ActualUserOpFeePerGas *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPostOpReverted is a free log retrieval operation binding the contract event 0xcedbbce311228041ee21adf55ebf281e4f143971bde7070bcc1311c0fcac7af3.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterPostOpReverted(opts *bind.FilterOpts, userOpHash [][32]byte) (*GasTankPaymasterPostOpRevertedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "PostOpReverted", userOpHashRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterPostOpRevertedIterator{contract: _GasTankPaymaster.contract, event: "PostOpReverted", logs: logs, sub: sub}, nil
}

// WatchPostOpReverted is a free log subscription operation binding the contract event 0xcedbbce311228041ee21adf55ebf281e4f143971bde7070bcc1311c0fcac7af3.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchPostOpReverted(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterPostOpReverted, userOpHash [][32]byte) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "PostOpReverted", userOpHashRule)
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

// ParsePostOpReverted is a log parse operation binding the contract event 0xcedbbce311228041ee21adf55ebf281e4f143971bde7070bcc1311c0fcac7af3.
//
// Solidity: event PostOpReverted(bytes32 indexed userOpHash, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
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

// GasTankPaymasterSponsoredIterator is returned from FilterSponsored and is used to iterate over the raw logs and unpacked data for Sponsored events raised by the GasTankPaymaster contract.
type GasTankPaymasterSponsoredIterator struct {
	Event *GasTankPaymasterSponsored // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterSponsoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterSponsored)
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
		it.Event = new(GasTankPaymasterSponsored)
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
func (it *GasTankPaymasterSponsoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterSponsoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterSponsored represents a Sponsored event raised by the GasTankPaymaster contract.
type GasTankPaymasterSponsored struct {
	UserOpHash            [32]byte
	Sender                common.Address
	Token                 common.Address
	MaxTokenCost          *big.Int
	MaxGasCost            *big.Int
	Success               bool
	ActualGasCost         *big.Int
	ActualUserOpFeePerGas *big.Int
	PostOpGas             *big.Int
	ActualTokenCost       *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSponsored is a free log retrieval operation binding the contract event 0xf97799dc487e2f824bc4b8f06c9e6e3218224d69a4a15befa111babc208ebe6e.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 maxTokenCost, uint256 maxGasCost, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas, uint256 postOpGas, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterSponsored(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, token []common.Address) (*GasTankPaymasterSponsoredIterator, error) {

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

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Sponsored", userOpHashRule, senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterSponsoredIterator{contract: _GasTankPaymaster.contract, event: "Sponsored", logs: logs, sub: sub}, nil
}

// WatchSponsored is a free log subscription operation binding the contract event 0xf97799dc487e2f824bc4b8f06c9e6e3218224d69a4a15befa111babc208ebe6e.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 maxTokenCost, uint256 maxGasCost, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas, uint256 postOpGas, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchSponsored(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterSponsored, userOpHash [][32]byte, sender []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Sponsored", userOpHashRule, senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterSponsored)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Sponsored", log); err != nil {
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

// ParseSponsored is a log parse operation binding the contract event 0xf97799dc487e2f824bc4b8f06c9e6e3218224d69a4a15befa111babc208ebe6e.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, address indexed sender, address indexed token, uint256 maxTokenCost, uint256 maxGasCost, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas, uint256 postOpGas, uint256 actualTokenCost)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseSponsored(log types.Log) (*GasTankPaymasterSponsored, error) {
	event := new(GasTankPaymasterSponsored)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Sponsored", log); err != nil {
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

// GasTankPaymasterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GasTankPaymaster contract.
type GasTankPaymasterUnpausedIterator struct {
	Event *GasTankPaymasterUnpaused // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterUnpaused)
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
		it.Event = new(GasTankPaymasterUnpaused)
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
func (it *GasTankPaymasterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterUnpaused represents a Unpaused event raised by the GasTankPaymaster contract.
type GasTankPaymasterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GasTankPaymasterUnpausedIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterUnpausedIterator{contract: _GasTankPaymaster.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterUnpaused) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterUnpaused)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseUnpaused(log types.Log) (*GasTankPaymasterUnpaused, error) {
	event := new(GasTankPaymasterUnpaused)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// GasTankPaymasterWithdrawPaymentIterator is returned from FilterWithdrawPayment and is used to iterate over the raw logs and unpacked data for WithdrawPayment events raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawPaymentIterator struct {
	Event *GasTankPaymasterWithdrawPayment // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterWithdrawPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterWithdrawPayment)
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
		it.Event = new(GasTankPaymasterWithdrawPayment)
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
func (it *GasTankPaymasterWithdrawPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterWithdrawPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterWithdrawPayment represents a WithdrawPayment event raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawPayment struct {
	Operator  common.Address
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawPayment is a free log retrieval operation binding the contract event 0x7408a9dadbab0577dee7cc1b110ca1d07ce941651371ec7ccf9635e3cf758d4c.
//
// Solidity: event WithdrawPayment(address indexed operator, address indexed token, address indexed recipient, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterWithdrawPayment(opts *bind.FilterOpts, operator []common.Address, token []common.Address, recipient []common.Address) (*GasTankPaymasterWithdrawPaymentIterator, error) {

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

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "WithdrawPayment", operatorRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterWithdrawPaymentIterator{contract: _GasTankPaymaster.contract, event: "WithdrawPayment", logs: logs, sub: sub}, nil
}

// WatchWithdrawPayment is a free log subscription operation binding the contract event 0x7408a9dadbab0577dee7cc1b110ca1d07ce941651371ec7ccf9635e3cf758d4c.
//
// Solidity: event WithdrawPayment(address indexed operator, address indexed token, address indexed recipient, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchWithdrawPayment(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterWithdrawPayment, operator []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "WithdrawPayment", operatorRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterWithdrawPayment)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawPayment", log); err != nil {
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

// ParseWithdrawPayment is a log parse operation binding the contract event 0x7408a9dadbab0577dee7cc1b110ca1d07ce941651371ec7ccf9635e3cf758d4c.
//
// Solidity: event WithdrawPayment(address indexed operator, address indexed token, address indexed recipient, uint256 amount)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseWithdrawPayment(log types.Log) (*GasTankPaymasterWithdrawPayment, error) {
	event := new(GasTankPaymasterWithdrawPayment)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterWithdrawalCancelledIterator is returned from FilterWithdrawalCancelled and is used to iterate over the raw logs and unpacked data for WithdrawalCancelled events raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalCancelledIterator struct {
	Event *GasTankPaymasterWithdrawalCancelled // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterWithdrawalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterWithdrawalCancelled)
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
		it.Event = new(GasTankPaymasterWithdrawalCancelled)
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
func (it *GasTankPaymasterWithdrawalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterWithdrawalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterWithdrawalCancelled represents a WithdrawalCancelled event raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalCancelled struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCancelled is a free log retrieval operation binding the contract event 0x06788d6037cfde4c7d6701eaa3ddf0875dfe514ed5d6ea4ca14e36c24b107841.
//
// Solidity: event WithdrawalCancelled(address indexed sender, address indexed token)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterWithdrawalCancelled(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*GasTankPaymasterWithdrawalCancelledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "WithdrawalCancelled", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterWithdrawalCancelledIterator{contract: _GasTankPaymaster.contract, event: "WithdrawalCancelled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCancelled is a free log subscription operation binding the contract event 0x06788d6037cfde4c7d6701eaa3ddf0875dfe514ed5d6ea4ca14e36c24b107841.
//
// Solidity: event WithdrawalCancelled(address indexed sender, address indexed token)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchWithdrawalCancelled(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterWithdrawalCancelled, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "WithdrawalCancelled", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterWithdrawalCancelled)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
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

// ParseWithdrawalCancelled is a log parse operation binding the contract event 0x06788d6037cfde4c7d6701eaa3ddf0875dfe514ed5d6ea4ca14e36c24b107841.
//
// Solidity: event WithdrawalCancelled(address indexed sender, address indexed token)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseWithdrawalCancelled(log types.Log) (*GasTankPaymasterWithdrawalCancelled, error) {
	event := new(GasTankPaymasterWithdrawalCancelled)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterWithdrawalDelayUpdatedIterator is returned from FilterWithdrawalDelayUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalDelayUpdated events raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalDelayUpdatedIterator struct {
	Event *GasTankPaymasterWithdrawalDelayUpdated // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterWithdrawalDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterWithdrawalDelayUpdated)
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
		it.Event = new(GasTankPaymasterWithdrawalDelayUpdated)
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
func (it *GasTankPaymasterWithdrawalDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterWithdrawalDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterWithdrawalDelayUpdated represents a WithdrawalDelayUpdated event raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalDelayUpdated struct {
	OldDelay *big.Int
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayUpdated is a free log retrieval operation binding the contract event 0x14f1d1f27725572a5713cdc18afb00e27e109dbfafd3a9991dc1be64370c4e64.
//
// Solidity: event WithdrawalDelayUpdated(uint48 oldDelay, uint48 newDelay)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterWithdrawalDelayUpdated(opts *bind.FilterOpts) (*GasTankPaymasterWithdrawalDelayUpdatedIterator, error) {

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "WithdrawalDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterWithdrawalDelayUpdatedIterator{contract: _GasTankPaymaster.contract, event: "WithdrawalDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayUpdated is a free log subscription operation binding the contract event 0x14f1d1f27725572a5713cdc18afb00e27e109dbfafd3a9991dc1be64370c4e64.
//
// Solidity: event WithdrawalDelayUpdated(uint48 oldDelay, uint48 newDelay)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchWithdrawalDelayUpdated(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterWithdrawalDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "WithdrawalDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterWithdrawalDelayUpdated)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalDelayUpdated", log); err != nil {
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

// ParseWithdrawalDelayUpdated is a log parse operation binding the contract event 0x14f1d1f27725572a5713cdc18afb00e27e109dbfafd3a9991dc1be64370c4e64.
//
// Solidity: event WithdrawalDelayUpdated(uint48 oldDelay, uint48 newDelay)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseWithdrawalDelayUpdated(log types.Log) (*GasTankPaymasterWithdrawalDelayUpdated, error) {
	event := new(GasTankPaymasterWithdrawalDelayUpdated)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasTankPaymasterWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalRequestedIterator struct {
	Event *GasTankPaymasterWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *GasTankPaymasterWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasTankPaymasterWithdrawalRequested)
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
		it.Event = new(GasTankPaymasterWithdrawalRequested)
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
func (it *GasTankPaymasterWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasTankPaymasterWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasTankPaymasterWithdrawalRequested represents a WithdrawalRequested event raised by the GasTankPaymaster contract.
type GasTankPaymasterWithdrawalRequested struct {
	Sender         common.Address
	Token          common.Address
	Amount         *big.Int
	WithdrawableAt *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xa3b042426ba46a9e8338fb225f9c5812ff35cd03b5945b082789e9475311d1a1.
//
// Solidity: event WithdrawalRequested(address indexed sender, address indexed token, uint256 amount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*GasTankPaymasterWithdrawalRequestedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.FilterLogs(opts, "WithdrawalRequested", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GasTankPaymasterWithdrawalRequestedIterator{contract: _GasTankPaymaster.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xa3b042426ba46a9e8338fb225f9c5812ff35cd03b5945b082789e9475311d1a1.
//
// Solidity: event WithdrawalRequested(address indexed sender, address indexed token, uint256 amount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *GasTankPaymasterWithdrawalRequested, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GasTankPaymaster.contract.WatchLogs(opts, "WithdrawalRequested", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasTankPaymasterWithdrawalRequested)
				if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xa3b042426ba46a9e8338fb225f9c5812ff35cd03b5945b082789e9475311d1a1.
//
// Solidity: event WithdrawalRequested(address indexed sender, address indexed token, uint256 amount, uint48 withdrawableAt)
func (_GasTankPaymaster *GasTankPaymasterFilterer) ParseWithdrawalRequested(log types.Log) (*GasTankPaymasterWithdrawalRequested, error) {
	event := new(GasTankPaymasterWithdrawalRequested)
	if err := _GasTankPaymaster.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
