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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Target   common.Address
	Value    *big.Int
	CallData []byte
}

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

// SimpleSmartAccount7702MetaData contains all meta data concerning the SimpleSmartAccount7702 contract.
var SimpleSmartAccount7702MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ENTRY_POINT\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structExecution[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SimpleSmartAccount7702ABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleSmartAccount7702MetaData.ABI instead.
var SimpleSmartAccount7702ABI = SimpleSmartAccount7702MetaData.ABI

// SimpleSmartAccount7702 is an auto generated Go binding around an Ethereum contract.
type SimpleSmartAccount7702 struct {
	SimpleSmartAccount7702Caller     // Read-only binding to the contract
	SimpleSmartAccount7702Transactor // Write-only binding to the contract
	SimpleSmartAccount7702Filterer   // Log filterer for contract events
}

// SimpleSmartAccount7702Caller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleSmartAccount7702Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSmartAccount7702Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleSmartAccount7702Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSmartAccount7702Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleSmartAccount7702Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSmartAccount7702Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSmartAccount7702Session struct {
	Contract     *SimpleSmartAccount7702 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SimpleSmartAccount7702CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleSmartAccount7702CallerSession struct {
	Contract *SimpleSmartAccount7702Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SimpleSmartAccount7702TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleSmartAccount7702TransactorSession struct {
	Contract     *SimpleSmartAccount7702Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SimpleSmartAccount7702Raw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleSmartAccount7702Raw struct {
	Contract *SimpleSmartAccount7702 // Generic contract binding to access the raw methods on
}

// SimpleSmartAccount7702CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleSmartAccount7702CallerRaw struct {
	Contract *SimpleSmartAccount7702Caller // Generic read-only contract binding to access the raw methods on
}

// SimpleSmartAccount7702TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleSmartAccount7702TransactorRaw struct {
	Contract *SimpleSmartAccount7702Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleSmartAccount7702 creates a new instance of SimpleSmartAccount7702, bound to a specific deployed contract.
func NewSimpleSmartAccount7702(address common.Address, backend bind.ContractBackend) (*SimpleSmartAccount7702, error) {
	contract, err := bindSimpleSmartAccount7702(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleSmartAccount7702{SimpleSmartAccount7702Caller: SimpleSmartAccount7702Caller{contract: contract}, SimpleSmartAccount7702Transactor: SimpleSmartAccount7702Transactor{contract: contract}, SimpleSmartAccount7702Filterer: SimpleSmartAccount7702Filterer{contract: contract}}, nil
}

// NewSimpleSmartAccount7702Caller creates a new read-only instance of SimpleSmartAccount7702, bound to a specific deployed contract.
func NewSimpleSmartAccount7702Caller(address common.Address, caller bind.ContractCaller) (*SimpleSmartAccount7702Caller, error) {
	contract, err := bindSimpleSmartAccount7702(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleSmartAccount7702Caller{contract: contract}, nil
}

// NewSimpleSmartAccount7702Transactor creates a new write-only instance of SimpleSmartAccount7702, bound to a specific deployed contract.
func NewSimpleSmartAccount7702Transactor(address common.Address, transactor bind.ContractTransactor) (*SimpleSmartAccount7702Transactor, error) {
	contract, err := bindSimpleSmartAccount7702(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleSmartAccount7702Transactor{contract: contract}, nil
}

// NewSimpleSmartAccount7702Filterer creates a new log filterer instance of SimpleSmartAccount7702, bound to a specific deployed contract.
func NewSimpleSmartAccount7702Filterer(address common.Address, filterer bind.ContractFilterer) (*SimpleSmartAccount7702Filterer, error) {
	contract, err := bindSimpleSmartAccount7702(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleSmartAccount7702Filterer{contract: contract}, nil
}

// bindSimpleSmartAccount7702 binds a generic wrapper to an already deployed contract.
func bindSimpleSmartAccount7702(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleSmartAccount7702MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleSmartAccount7702.Contract.SimpleSmartAccount7702Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.SimpleSmartAccount7702Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.SimpleSmartAccount7702Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleSmartAccount7702.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.contract.Transact(opts, method, params...)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Caller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleSmartAccount7702.contract.Call(opts, &out, "ENTRY_POINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) ENTRYPOINT() (common.Address, error) {
	return _SimpleSmartAccount7702.Contract.ENTRYPOINT(&_SimpleSmartAccount7702.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0x94430fa5.
//
// Solidity: function ENTRY_POINT() view returns(address)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702CallerSession) ENTRYPOINT() (common.Address, error) {
	return _SimpleSmartAccount7702.Contract.ENTRYPOINT(&_SimpleSmartAccount7702.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Caller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleSmartAccount7702.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) GetNonce() (*big.Int, error) {
	return _SimpleSmartAccount7702.Contract.GetNonce(&_SimpleSmartAccount7702.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702CallerSession) GetNonce() (*big.Int, error) {
	return _SimpleSmartAccount7702.Contract.GetNonce(&_SimpleSmartAccount7702.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Caller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _SimpleSmartAccount7702.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _SimpleSmartAccount7702.Contract.IsValidSignature(&_SimpleSmartAccount7702.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702CallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _SimpleSmartAccount7702.Contract.IsValidSignature(&_SimpleSmartAccount7702.CallOpts, hash, signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SimpleSmartAccount7702.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleSmartAccount7702.Contract.SupportsInterface(&_SimpleSmartAccount7702.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleSmartAccount7702.Contract.SupportsInterface(&_SimpleSmartAccount7702.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.Execute(&_SimpleSmartAccount7702.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.Execute(&_SimpleSmartAccount7702.TransactOpts, target, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) ExecuteBatch(opts *bind.TransactOpts, calls []Execution) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) ExecuteBatch(calls []Execution) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.ExecuteBatch(&_SimpleSmartAccount7702.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) ExecuteBatch(calls []Execution) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.ExecuteBatch(&_SimpleSmartAccount7702.TransactOpts, calls)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC1155BatchReceived(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC1155BatchReceived(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC1155Received(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC1155Received(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC721Received(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.OnERC721Received(&_SimpleSmartAccount7702.TransactOpts, arg0, arg1, arg2, arg3)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.ValidateUserOp(&_SimpleSmartAccount7702.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.ValidateUserOp(&_SimpleSmartAccount7702.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleSmartAccount7702.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702Session) Receive() (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.Receive(&_SimpleSmartAccount7702.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleSmartAccount7702 *SimpleSmartAccount7702TransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleSmartAccount7702.Contract.Receive(&_SimpleSmartAccount7702.TransactOpts)
}
