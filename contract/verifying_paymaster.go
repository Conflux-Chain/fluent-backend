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

// VerifyingPaymasterMetaData contains all meta data concerning the VerifyingPaymaster contract.
var VerifyingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SmartAccountWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"Sponsored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getPaymasterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSignerAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartAccount\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setSmartAccountAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartAccount\",\"type\":\"address\"}],\"name\":\"smartAccountWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VerifyingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyingPaymasterMetaData.ABI instead.
var VerifyingPaymasterABI = VerifyingPaymasterMetaData.ABI

// VerifyingPaymaster is an auto generated Go binding around an Ethereum contract.
type VerifyingPaymaster struct {
	VerifyingPaymasterCaller     // Read-only binding to the contract
	VerifyingPaymasterTransactor // Write-only binding to the contract
	VerifyingPaymasterFilterer   // Log filterer for contract events
}

// VerifyingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyingPaymasterSession struct {
	Contract     *VerifyingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifyingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyingPaymasterCallerSession struct {
	Contract *VerifyingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifyingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyingPaymasterTransactorSession struct {
	Contract     *VerifyingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifyingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyingPaymasterRaw struct {
	Contract *VerifyingPaymaster // Generic contract binding to access the raw methods on
}

// VerifyingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCallerRaw struct {
	Contract *VerifyingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactorRaw struct {
	Contract *VerifyingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyingPaymaster creates a new instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymaster(address common.Address, backend bind.ContractBackend) (*VerifyingPaymaster, error) {
	contract, err := bindVerifyingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymaster{VerifyingPaymasterCaller: VerifyingPaymasterCaller{contract: contract}, VerifyingPaymasterTransactor: VerifyingPaymasterTransactor{contract: contract}, VerifyingPaymasterFilterer: VerifyingPaymasterFilterer{contract: contract}}, nil
}

// NewVerifyingPaymasterCaller creates a new read-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*VerifyingPaymasterCaller, error) {
	contract, err := bindVerifyingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterCaller{contract: contract}, nil
}

// NewVerifyingPaymasterTransactor creates a new write-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyingPaymasterTransactor, error) {
	contract, err := bindVerifyingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterTransactor{contract: contract}, nil
}

// NewVerifyingPaymasterFilterer creates a new log filterer instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyingPaymasterFilterer, error) {
	contract, err := bindVerifyingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterFilterer{contract: contract}, nil
}

// bindVerifyingPaymaster binds a generic wrapper to an already deployed contract.
func bindVerifyingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyingPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterSession) Balance() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.Balance(&_VerifyingPaymaster.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Balance() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.Balance(&_VerifyingPaymaster.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "eip712Domain")

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
func (_VerifyingPaymaster *VerifyingPaymasterSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VerifyingPaymaster.Contract.Eip712Domain(&_VerifyingPaymaster.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VerifyingPaymaster.Contract.Eip712Domain(&_VerifyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetPaymasterHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getPaymasterHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetPaymasterHash(userOp PackedUserOperation) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetPaymasterHash(&_VerifyingPaymaster.CallOpts, userOp)
}

// GetPaymasterHash is a free data retrieval call binding the contract method 0xf11a7bea.
//
// Solidity: function getPaymasterHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetPaymasterHash(userOp PackedUserOperation) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetPaymasterHash(&_VerifyingPaymaster.CallOpts, userOp)
}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) IsSignerAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "isSignerAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_VerifyingPaymaster *VerifyingPaymasterSession) IsSignerAllowed(arg0 common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.IsSignerAllowed(&_VerifyingPaymaster.CallOpts, arg0)
}

// IsSignerAllowed is a free data retrieval call binding the contract method 0xdddf6ff9.
//
// Solidity: function isSignerAllowed(address ) view returns(bool)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) IsSignerAllowed(arg0 common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.IsSignerAllowed(&_VerifyingPaymaster.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// SmartAccountWhitelist is a free data retrieval call binding the contract method 0x72c6b323.
//
// Solidity: function smartAccountWhitelist(address smartAccount) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) SmartAccountWhitelist(opts *bind.CallOpts, smartAccount common.Address) (bool, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "smartAccountWhitelist", smartAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SmartAccountWhitelist is a free data retrieval call binding the contract method 0x72c6b323.
//
// Solidity: function smartAccountWhitelist(address smartAccount) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterSession) SmartAccountWhitelist(smartAccount common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.SmartAccountWhitelist(&_VerifyingPaymaster.CallOpts, smartAccount)
}

// SmartAccountWhitelist is a free data retrieval call binding the contract method 0x72c6b323.
//
// Solidity: function smartAccountWhitelist(address smartAccount) view returns(bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) SmartAccountWhitelist(smartAccount common.Address) (bool, error) {
	return _VerifyingPaymaster.Contract.SmartAccountWhitelist(&_VerifyingPaymaster.CallOpts, smartAccount)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) view returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "validatePaymasterUserOp", userOp, userOpHash, arg2)

	outstruct := new(struct {
		Context        []byte
		ValidationData *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) view returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.CallOpts, userOp, userOpHash, arg2)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 ) view returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, arg2 *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.CallOpts, userOp, userOpHash, arg2)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) SetSigner(opts *bind.TransactOpts, signer common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "setSigner", signer, allowed)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) SetSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetSigner(&_VerifyingPaymaster.TransactOpts, signer, allowed)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address signer, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) SetSigner(signer common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetSigner(&_VerifyingPaymaster.TransactOpts, signer, allowed)
}

// SetSmartAccountAllowed is a paid mutator transaction binding the contract method 0x3f24bc2f.
//
// Solidity: function setSmartAccountAllowed(address smartAccount, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) SetSmartAccountAllowed(opts *bind.TransactOpts, smartAccount common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "setSmartAccountAllowed", smartAccount, allowed)
}

// SetSmartAccountAllowed is a paid mutator transaction binding the contract method 0x3f24bc2f.
//
// Solidity: function setSmartAccountAllowed(address smartAccount, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) SetSmartAccountAllowed(smartAccount common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetSmartAccountAllowed(&_VerifyingPaymaster.TransactOpts, smartAccount, allowed)
}

// SetSmartAccountAllowed is a paid mutator transaction binding the contract method 0x3f24bc2f.
//
// Solidity: function setSmartAccountAllowed(address smartAccount, bool allowed) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) SetSmartAccountAllowed(smartAccount common.Address, allowed bool) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.SetSmartAccountAllowed(&_VerifyingPaymaster.TransactOpts, smartAccount, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Receive() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Receive(&_VerifyingPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Receive(&_VerifyingPaymaster.TransactOpts)
}

// VerifyingPaymasterEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterEIP712DomainChangedIterator struct {
	Event *VerifyingPaymasterEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterEIP712DomainChanged)
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
		it.Event = new(VerifyingPaymasterEIP712DomainChanged)
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
func (it *VerifyingPaymasterEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterEIP712DomainChanged represents a EIP712DomainChanged event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*VerifyingPaymasterEIP712DomainChangedIterator, error) {

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterEIP712DomainChangedIterator{contract: _VerifyingPaymaster.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterEIP712DomainChanged)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseEIP712DomainChanged(log types.Log) (*VerifyingPaymasterEIP712DomainChanged, error) {
	event := new(VerifyingPaymasterEIP712DomainChanged)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferredIterator struct {
	Event *VerifyingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
		it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterOwnershipTransferredIterator{contract: _VerifyingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterOwnershipTransferred)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*VerifyingPaymasterOwnershipTransferred, error) {
	event := new(VerifyingPaymasterOwnershipTransferred)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSignerUpdatedIterator struct {
	Event *VerifyingPaymasterSignerUpdated // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterSignerUpdated)
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
		it.Event = new(VerifyingPaymasterSignerUpdated)
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
func (it *VerifyingPaymasterSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterSignerUpdated represents a SignerUpdated event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSignerUpdated struct {
	Signer  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0xfcaa24b1276bfa7dbf77797c0a984b9df924acbeaabd48cd2f1b0eca379b78fa.
//
// Solidity: event SignerUpdated(address indexed signer, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterSignerUpdated(opts *bind.FilterOpts, signer []common.Address) (*VerifyingPaymasterSignerUpdatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "SignerUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterSignerUpdatedIterator{contract: _VerifyingPaymaster.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0xfcaa24b1276bfa7dbf77797c0a984b9df924acbeaabd48cd2f1b0eca379b78fa.
//
// Solidity: event SignerUpdated(address indexed signer, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterSignerUpdated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "SignerUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterSignerUpdated)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseSignerUpdated(log types.Log) (*VerifyingPaymasterSignerUpdated, error) {
	event := new(VerifyingPaymasterSignerUpdated)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterSmartAccountWhitelistUpdatedIterator is returned from FilterSmartAccountWhitelistUpdated and is used to iterate over the raw logs and unpacked data for SmartAccountWhitelistUpdated events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSmartAccountWhitelistUpdatedIterator struct {
	Event *VerifyingPaymasterSmartAccountWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterSmartAccountWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterSmartAccountWhitelistUpdated)
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
		it.Event = new(VerifyingPaymasterSmartAccountWhitelistUpdated)
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
func (it *VerifyingPaymasterSmartAccountWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterSmartAccountWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterSmartAccountWhitelistUpdated represents a SmartAccountWhitelistUpdated event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSmartAccountWhitelistUpdated struct {
	SmartAccount common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSmartAccountWhitelistUpdated is a free log retrieval operation binding the contract event 0xf08180ad383eb03c0e052f1756659253e8ab36ff79adca05c56b95b7d3dda211.
//
// Solidity: event SmartAccountWhitelistUpdated(address indexed smartAccount, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterSmartAccountWhitelistUpdated(opts *bind.FilterOpts, smartAccount []common.Address) (*VerifyingPaymasterSmartAccountWhitelistUpdatedIterator, error) {

	var smartAccountRule []interface{}
	for _, smartAccountItem := range smartAccount {
		smartAccountRule = append(smartAccountRule, smartAccountItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "SmartAccountWhitelistUpdated", smartAccountRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterSmartAccountWhitelistUpdatedIterator{contract: _VerifyingPaymaster.contract, event: "SmartAccountWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchSmartAccountWhitelistUpdated is a free log subscription operation binding the contract event 0xf08180ad383eb03c0e052f1756659253e8ab36ff79adca05c56b95b7d3dda211.
//
// Solidity: event SmartAccountWhitelistUpdated(address indexed smartAccount, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchSmartAccountWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterSmartAccountWhitelistUpdated, smartAccount []common.Address) (event.Subscription, error) {

	var smartAccountRule []interface{}
	for _, smartAccountItem := range smartAccount {
		smartAccountRule = append(smartAccountRule, smartAccountItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "SmartAccountWhitelistUpdated", smartAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterSmartAccountWhitelistUpdated)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "SmartAccountWhitelistUpdated", log); err != nil {
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

// ParseSmartAccountWhitelistUpdated is a log parse operation binding the contract event 0xf08180ad383eb03c0e052f1756659253e8ab36ff79adca05c56b95b7d3dda211.
//
// Solidity: event SmartAccountWhitelistUpdated(address indexed smartAccount, bool allowed)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseSmartAccountWhitelistUpdated(log types.Log) (*VerifyingPaymasterSmartAccountWhitelistUpdated, error) {
	event := new(VerifyingPaymasterSmartAccountWhitelistUpdated)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "SmartAccountWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifyingPaymasterSponsoredIterator is returned from FilterSponsored and is used to iterate over the raw logs and unpacked data for Sponsored events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSponsoredIterator struct {
	Event *VerifyingPaymasterSponsored // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterSponsoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterSponsored)
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
		it.Event = new(VerifyingPaymasterSponsored)
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
func (it *VerifyingPaymasterSponsoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterSponsoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterSponsored represents a Sponsored event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterSponsored struct {
	UserOpHash            [32]byte
	Success               bool
	ActualGasCost         *big.Int
	ActualUserOpFeePerGas *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSponsored is a free log retrieval operation binding the contract event 0x9e4d0db315fe76b23fdbfa2d345dd700140cae22039688064727b295782b0204.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterSponsored(opts *bind.FilterOpts, userOpHash [][32]byte) (*VerifyingPaymasterSponsoredIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "Sponsored", userOpHashRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterSponsoredIterator{contract: _VerifyingPaymaster.contract, event: "Sponsored", logs: logs, sub: sub}, nil
}

// WatchSponsored is a free log subscription operation binding the contract event 0x9e4d0db315fe76b23fdbfa2d345dd700140cae22039688064727b295782b0204.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchSponsored(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterSponsored, userOpHash [][32]byte) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "Sponsored", userOpHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterSponsored)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "Sponsored", log); err != nil {
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

// ParseSponsored is a log parse operation binding the contract event 0x9e4d0db315fe76b23fdbfa2d345dd700140cae22039688064727b295782b0204.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseSponsored(log types.Log) (*VerifyingPaymasterSponsored, error) {
	event := new(VerifyingPaymasterSponsored)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "Sponsored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
