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

// OnePoolRewardMetaData contains all meta data concerning the OnePoolReward contract.
var OnePoolRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lifetimeSeconds_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"minerId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accumulatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"}],\"name\":\"claimMineReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardSectors\",\"type\":\"uint256\"}],\"name\":\"fillReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mine_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lifetimeInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextChunkDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeoutRecords\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"numPriceChunks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeoutTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"donation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a034606557601f610a4c38819003918201601f19168301916001600160401b03831184841017606a57808492602094604052833981010312606557516080526040516109cb90816100818239608051818181610156015281816102a901526109560152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561002a575b361561001957600080fd5b610025346009546105f7565b600955005b60003560e01c80630dbbe0a21461052457806314bcec9f14610506578063158ef93e146104e357806318ca1b2b146104c557806322ee4cb8146104a757806328fedefe14610489578063485cc95514610365578063514ccef01461034757806359e96700146102245780636baff3d31461020657806377bd602b146101e857806380f55605146101be57806380fad325146101a057806399f4b25114610179578063b7375ddb1461013e578063b7a3c04c1461010a5763f8ac93e80361000e5734610105576000366003190112610105576101036107ca565b005b600080fd5b34610105576060366003190112610105576024356001600160a01b0381168103610105576101039060443590600435610701565b346101055760003660031901126101055760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101055760003660031901126101055760206001600160a01b0360015416604051908152f35b34610105576000366003190112610105576020600654604051908152f35b346101055760003660031901126101055760206001600160a01b0360005460081c16604051908152f35b34610105576000366003190112610105576020600a54604051908152f35b34610105576000366003190112610105576020600954604051908152f35b60403660031901126101055760043561024c6001600160a01b0360005460081c1633146105ac565b6102546107ca565b61027d610263602435836105f7565b9161027764020000000060081c809261062d565b9261062d565b9080821161028757005b61029a67ffffffffffffffff918361064d565b1667ffffffffffffffff6102ce7f0000000000000000000000000000000000000000000000000000000000000000426105f7565b16600954906102db61065a565b928352602083015260408201526002549068010000000000000000821015610331576103108260016103169401600255610577565b9061067a565b600555610327600954600a546105f7565b600a556000600955005b634e487b7160e01b600052604160045260246000fd5b34610105576000366003190112610105576020600754604051908152f35b34610105576040366003190112610105576004356001600160a01b038116810361010557602435906001600160a01b0382168092036101055760005460ff8116610420577fffffffffffffffffffffff00000000000000000000000000000000000000000074ffffffffffffffffffffffffffffffffffffffff0060019360081b16911617176000557fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015542600855600080f35b608460405162461bcd60e51b8152602060048201526024808201527f5a67496e697469616c697a61626c653a20616c726561647920696e697469616c60448201527f697a6564000000000000000000000000000000000000000000000000000000006064820152fd5b34610105576000366003190112610105576020600354604051908152f35b34610105576000366003190112610105576020600454604051908152f35b34610105576000366003190112610105576020600554604051908152f35b3461010557600036600319011261010557602060ff600054166040519015158152f35b34610105576000366003190112610105576020600854604051908152f35b34610105576020366003190112610105576004356002548110156101055761054d606091610577565b506001815491015467ffffffffffffffff60405192818116845260401c1660208301526040820152f35b60025481101561059657600260005260206000209060011b0190600090565b634e487b7160e01b600052603260045260246000fd5b156105b357565b606460405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e006044820152fd5b9190820180921161060457565b634e487b7160e01b600052601160045260246000fd5b8181029291811591840414171561060457565b8115610637570490565b634e487b7160e01b600052601260045260246000fd5b9190820391821161060457565b604051906060820182811067ffffffffffffffff82111761033157604052565b91906106eb5760408167ffffffffffffffff60019351168454908067ffffffffffffffff1983161786557fffffffffffffffffffffffffffffffff000000000000000000000000000000006fffffffffffffffff00000000000000006020850151861b169216171784550151910155565b634e487b7160e01b600052600060045260246000fd5b6107176001600160a01b036001541633146105ac565b60045481106107c5576107286107ca565b6107376006546007549061064d565b4781116107be575b8015928315610750575b5050505050565b6001600160a01b03169281846000926107b4575b600092839283928392f1156107a85760207ff099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b091604051908152a43880808080610749565b6040513d6000823e3d90fd5b6108fc9250610764565b504761073f565b505050565b60025480156108ab575b6003546107e081610577565b5067ffffffffffffffff42915460401c16116108a45767ffffffffffffffff61080b61081692610577565b505460401c166108b5565b61087e60035461083c67ffffffffffffffff61083183610577565b5054166004546105f7565b600455610858600161084d83610577565b500154600a5461064d565b600a5561031061086661065a565b91600083526000602084015260006040840152610577565b6003546001810180911161060457808291600355036107d457505b6108a2426108b5565b565b5050610899565b50426008556108a2425b60085480821115610991576108e26108cd828461064d565b6108dc6005546004549061064d565b9061061a565b6108f76340000000916402000000009061061a565b91600a830292808404600a148115171561060457678ac7230489e800000292808404670de0b6b3a764000014901517156106045781156106375761097b6109546301ea6e0061098194610989960404926108dc600a54918861064d565b7f00000000000000000000000000000000000000000000000000000000000000009061062d565b906105f7565b6006546105f7565b600655600855565b505056fea264697066735822122075c1a02052068f38393b2270d0c4f8dc0e805fff83f9f2a6bba7a9157c76847d64736f6c634300081b0033",
}

// OnePoolRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use OnePoolRewardMetaData.ABI instead.
var OnePoolRewardABI = OnePoolRewardMetaData.ABI

// OnePoolRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnePoolRewardMetaData.Bin instead.
var OnePoolRewardBin = OnePoolRewardMetaData.Bin

// DeployOnePoolReward deploys a new Ethereum contract, binding an instance of OnePoolReward to it.
func DeployOnePoolReward(auth *bind.TransactOpts, backend bind.ContractBackend, lifetimeSeconds_ *big.Int) (common.Address, *types.Transaction, *OnePoolReward, error) {
	parsed, err := OnePoolRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnePoolRewardBin), backend, lifetimeSeconds_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnePoolReward{OnePoolRewardCaller: OnePoolRewardCaller{contract: contract}, OnePoolRewardTransactor: OnePoolRewardTransactor{contract: contract}, OnePoolRewardFilterer: OnePoolRewardFilterer{contract: contract}}, nil
}

// OnePoolReward is an auto generated Go binding around an Ethereum contract.
type OnePoolReward struct {
	OnePoolRewardCaller     // Read-only binding to the contract
	OnePoolRewardTransactor // Write-only binding to the contract
	OnePoolRewardFilterer   // Log filterer for contract events
}

// OnePoolRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnePoolRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnePoolRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnePoolRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnePoolRewardSession struct {
	Contract     *OnePoolReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnePoolRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnePoolRewardCallerSession struct {
	Contract *OnePoolRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OnePoolRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnePoolRewardTransactorSession struct {
	Contract     *OnePoolRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OnePoolRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnePoolRewardRaw struct {
	Contract *OnePoolReward // Generic contract binding to access the raw methods on
}

// OnePoolRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnePoolRewardCallerRaw struct {
	Contract *OnePoolRewardCaller // Generic read-only contract binding to access the raw methods on
}

// OnePoolRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnePoolRewardTransactorRaw struct {
	Contract *OnePoolRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnePoolReward creates a new instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolReward(address common.Address, backend bind.ContractBackend) (*OnePoolReward, error) {
	contract, err := bindOnePoolReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnePoolReward{OnePoolRewardCaller: OnePoolRewardCaller{contract: contract}, OnePoolRewardTransactor: OnePoolRewardTransactor{contract: contract}, OnePoolRewardFilterer: OnePoolRewardFilterer{contract: contract}}, nil
}

// NewOnePoolRewardCaller creates a new read-only instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardCaller(address common.Address, caller bind.ContractCaller) (*OnePoolRewardCaller, error) {
	contract, err := bindOnePoolReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardCaller{contract: contract}, nil
}

// NewOnePoolRewardTransactor creates a new write-only instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*OnePoolRewardTransactor, error) {
	contract, err := bindOnePoolReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardTransactor{contract: contract}, nil
}

// NewOnePoolRewardFilterer creates a new log filterer instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*OnePoolRewardFilterer, error) {
	contract, err := bindOnePoolReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardFilterer{contract: contract}, nil
}

// bindOnePoolReward binds a generic wrapper to an already deployed contract.
func bindOnePoolReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnePoolRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolReward *OnePoolRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolReward.Contract.OnePoolRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolReward *OnePoolRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.Contract.OnePoolRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolReward *OnePoolRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolReward.Contract.OnePoolRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolReward *OnePoolRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolReward *OnePoolRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolReward *OnePoolRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolReward.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) AccumulatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "accumulatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.AccumulatedReward(&_OnePoolReward.CallOpts)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.AccumulatedReward(&_OnePoolReward.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) ActiveDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "activeDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.ActiveDonation(&_OnePoolReward.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.ActiveDonation(&_OnePoolReward.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) ClaimedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "claimedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.ClaimedReward(&_OnePoolReward.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.ClaimedReward(&_OnePoolReward.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) FirstValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "firstValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.FirstValidChunk(&_OnePoolReward.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.FirstValidChunk(&_OnePoolReward.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolReward *OnePoolRewardCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolReward *OnePoolRewardSession) Initialized() (bool, error) {
	return _OnePoolReward.Contract.Initialized(&_OnePoolReward.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolReward *OnePoolRewardCallerSession) Initialized() (bool, error) {
	return _OnePoolReward.Contract.Initialized(&_OnePoolReward.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LastUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lastUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolReward.Contract.LastUpdateTimestamp(&_OnePoolReward.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolReward.Contract.LastUpdateTimestamp(&_OnePoolReward.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LastValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lastValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.LastValidChunk(&_OnePoolReward.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.LastValidChunk(&_OnePoolReward.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LifetimeInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lifetimeInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolReward.Contract.LifetimeInSeconds(&_OnePoolReward.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolReward.Contract.LifetimeInSeconds(&_OnePoolReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolReward *OnePoolRewardCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolReward *OnePoolRewardSession) Market() (common.Address, error) {
	return _OnePoolReward.Contract.Market(&_OnePoolReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolReward *OnePoolRewardCallerSession) Market() (common.Address, error) {
	return _OnePoolReward.Contract.Market(&_OnePoolReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolReward *OnePoolRewardCaller) Mine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "mine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolReward *OnePoolRewardSession) Mine() (common.Address, error) {
	return _OnePoolReward.Contract.Mine(&_OnePoolReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolReward *OnePoolRewardCallerSession) Mine() (common.Address, error) {
	return _OnePoolReward.Contract.Mine(&_OnePoolReward.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) NextChunkDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "nextChunkDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.NextChunkDonation(&_OnePoolReward.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.NextChunkDonation(&_OnePoolReward.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) TimeoutHead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "timeoutHead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolReward.Contract.TimeoutHead(&_OnePoolReward.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolReward.Contract.TimeoutHead(&_OnePoolReward.CallOpts)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardCaller) TimeoutRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "timeoutRecords", arg0)

	outstruct := new(struct {
		NumPriceChunks   uint64
		TimeoutTimestamp uint64
		Donation         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumPriceChunks = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Donation = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolReward.Contract.TimeoutRecords(&_OnePoolReward.CallOpts, arg0)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardCallerSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolReward.Contract.TimeoutRecords(&_OnePoolReward.CallOpts, arg0)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_OnePoolReward *OnePoolRewardTransactor) ClaimMineReward(opts *bind.TransactOpts, pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "claimMineReward", pricingIndex, beneficiary, minerID)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_OnePoolReward *OnePoolRewardSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _OnePoolReward.Contract.ClaimMineReward(&_OnePoolReward.TransactOpts, pricingIndex, beneficiary, minerID)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 minerID) returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, minerID [32]byte) (*types.Transaction, error) {
	return _OnePoolReward.Contract.ClaimMineReward(&_OnePoolReward.TransactOpts, pricingIndex, beneficiary, minerID)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardTransactor) FillReward(opts *bind.TransactOpts, beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "fillReward", beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.Contract.FillReward(&_OnePoolReward.TransactOpts, beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.Contract.FillReward(&_OnePoolReward.TransactOpts, beforeLength, rewardSectors)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolReward *OnePoolRewardTransactor) Initialize(opts *bind.TransactOpts, market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "initialize", market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolReward *OnePoolRewardSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolReward.Contract.Initialize(&_OnePoolReward.TransactOpts, market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolReward.Contract.Initialize(&_OnePoolReward.TransactOpts, market_, mine_)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardTransactor) Refresh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "refresh")
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardSession) Refresh() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Refresh(&_OnePoolReward.TransactOpts)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) Refresh() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Refresh(&_OnePoolReward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardSession) Receive() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Receive(&_OnePoolReward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) Receive() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Receive(&_OnePoolReward.TransactOpts)
}

// OnePoolRewardDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the OnePoolReward contract.
type OnePoolRewardDistributeRewardIterator struct {
	Event *OnePoolRewardDistributeReward // Event containing the contract specifics and raw log

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
func (it *OnePoolRewardDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnePoolRewardDistributeReward)
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
		it.Event = new(OnePoolRewardDistributeReward)
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
func (it *OnePoolRewardDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnePoolRewardDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnePoolRewardDistributeReward represents a DistributeReward event raised by the OnePoolReward contract.
type OnePoolRewardDistributeReward struct {
	PricingIndex *big.Int
	Beneficiary  common.Address
	MinerId      [32]byte
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) FilterDistributeReward(opts *bind.FilterOpts, pricingIndex []*big.Int, beneficiary []common.Address, minerId [][32]byte) (*OnePoolRewardDistributeRewardIterator, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _OnePoolReward.contract.FilterLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardDistributeRewardIterator{contract: _OnePoolReward.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *OnePoolRewardDistributeReward, pricingIndex []*big.Int, beneficiary []common.Address, minerId [][32]byte) (event.Subscription, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _OnePoolReward.contract.WatchLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnePoolRewardDistributeReward)
				if err := _OnePoolReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0xf099de0c527342bff1397d08daae31c3de66104f48886094fbe9126a21ebf4b0.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, bytes32 indexed minerId, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) ParseDistributeReward(log types.Log) (*OnePoolRewardDistributeReward, error) {
	event := new(OnePoolRewardDistributeReward)
	if err := _OnePoolReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
