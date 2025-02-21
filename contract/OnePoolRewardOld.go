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

// OnePoolRewardOldMetaData contains all meta data concerning the OnePoolRewardOld contract.
var OnePoolRewardOldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lifetimeSeconds_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accumulatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"claimMineReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardSectors\",\"type\":\"uint256\"}],\"name\":\"fillReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mine_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lifetimeInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextChunkDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeoutRecords\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"numPriceChunks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeoutTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"donation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610ca9380380610ca983398101604081905261002f91610037565b608052610050565b60006020828403121561004957600080fd5b5051919050565b608051610c30610079600039600081816102fe0152818161056c0152610a300152610c306000f3fe6080604052600436106101025760003560e01c806359e967001161009557806380fad3251161006457806380fad325146102b657806399f4b251146102cc578063b7375ddb146102ec578063b7a3c04c14610320578063f8ac93e81461034057600080fd5b806359e967001461023a5780636baff3d31461024d57806377bd602b1461026357806380f556051461027957600080fd5b806322ee4cb8116100d157806322ee4cb8146101d657806328fedefe146101ec578063485cc95514610202578063514ccef01461022457600080fd5b80630dbbe0a21461012657806314bcec9f14610172578063158ef93e1461019657806318ca1b2b146101c057600080fd5b366101215734600960008282546101199190610ab6565b925050819055005b600080fd5b34801561013257600080fd5b50610146610141366004610acf565b610355565b6040805167ffffffffffffffff9485168152939092166020840152908201526060015b60405180910390f35b34801561017e57600080fd5b5061018860085481565b604051908152602001610169565b3480156101a257600080fd5b506000546101b09060ff1681565b6040519015158152602001610169565b3480156101cc57600080fd5b5061018860055481565b3480156101e257600080fd5b5061018860045481565b3480156101f857600080fd5b5061018860035481565b34801561020e57600080fd5b5061022261021d366004610afd565b61039a565b005b34801561023057600080fd5b5061018860075481565b610222610248366004610b36565b61043d565b34801561025957600080fd5b5061018860095481565b34801561026f57600080fd5b50610188600a5481565b34801561028557600080fd5b5060005461029e9061010090046001600160a01b031681565b6040516001600160a01b039091168152602001610169565b3480156102c257600080fd5b5061018860065481565b3480156102d857600080fd5b5060015461029e906001600160a01b031681565b3480156102f857600080fd5b506101887f000000000000000000000000000000000000000000000000000000000000000081565b34801561032c57600080fd5b5061022261033b366004610b58565b610668565b34801561034c57600080fd5b50610222610789565b6002818154811061036557600080fd5b60009182526020909120600290910201805460019091015467ffffffffffffffff8083169350600160401b9092049091169083565b60005460ff16156103fe5760405162461bcd60e51b8152602060048201526024808201527f5a67496e697469616c697a61626c653a20616c726561647920696e697469616c6044820152631a5e995960e21b60648201526084015b60405180910390fd5b600080546001600160a81b0319166101006001600160a01b039485160217600190811790915580546001600160a01b0319169190921617905542600855565b60005461010090046001600160a01b0316336001600160a01b0316146104a55760405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e0060448201526064016103f5565b6104ad610789565b60006104b98284610ab6565b905060006101006104cc61040080610b90565b6104d890610400610b90565b6104e3906008610b90565b6104ed9190610baf565b6104f79085610baf565b9050600061010061050a61040080610b90565b61051690610400610b90565b610521906008610b90565b61052b9190610baf565b6105359084610baf565b905081811115610661576000604051806060016040528084846105589190610bd1565b67ffffffffffffffff1681526020016105917f000000000000000000000000000000000000000000000000000000000000000042610ab6565b67ffffffffffffffff9081168252600980546020938401526002805460018101825560008281528651919092027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81018054968801518616600160401b026001600160801b031990971692909516919091179490941790925560408401517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf9093019290925560058590559054600a80549394509092909190610655908490610ab6565b90915550506000600955505b5050505050565b6001546001600160a01b0316336001600160a01b0316146106cb5760405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e0060448201526064016103f5565b6004548310156106da57505050565b6106e2610789565b60006007546006546106f49190610bd1565b9050478111156107015750475b8015610783576040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561073d573d6000803e3d6000fd5b50826001600160a01b0316847f83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c8360405161077a91815260200190565b60405180910390a35b50505050565b60025460008190036107a7574260088190556107a490610964565b50565b426002600354815481106107bd576107bd610be4565b6000918252602090912060029091020154600160401b900467ffffffffffffffff161161095f576108236002600354815481106107fc576107fc610be4565b6000918252602090912060029091020154600160401b900467ffffffffffffffff16610964565b60026003548154811061083857610838610be4565b600091825260208220600290910201546004805467ffffffffffffffff909216929091610866908490610ab6565b9250508190555060026003548154811061088257610882610be4565b906000526020600020906002020160010154600a60008282546108a59190610bd1565b925050819055506040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152506002600354815481106108f3576108f3610be4565b60009182526020808320845160029093020180549185015167ffffffffffffffff908116600160401b026001600160801b031990931693169290921717815560409092015160019283015560038054909190610950908490610ab6565b909155505060035481036107a7575b6107a4425b60085481116109705750565b600061097e601f600c610b90565b61098b9062015180610b90565b61099761040080610b90565b6109a390610400610b90565b670de0b6b3a7640000600a6109ba61040080610b90565b6109c690610400610b90565b6109d1906008610b90565b6004546005546109e19190610bd1565b6008546109ee9089610bd1565b6109f89190610b90565b610a029190610b90565b610a0c9190610b90565b610a169190610b90565b610a209190610baf565b610a2a9190610baf565b905060007f000000000000000000000000000000000000000000000000000000000000000060085484610a5d9190610bd1565b600a54610a6a9190610b90565b610a749190610baf565b9050610a808183610ab6565b60066000828254610a919190610ab6565b90915550505060089190915550565b634e487b7160e01b600052601160045260246000fd5b80820180821115610ac957610ac9610aa0565b92915050565b600060208284031215610ae157600080fd5b5035919050565b6001600160a01b03811681146107a457600080fd5b60008060408385031215610b1057600080fd5b8235610b1b81610ae8565b91506020830135610b2b81610ae8565b809150509250929050565b60008060408385031215610b4957600080fd5b50508035926020909101359150565b600080600060608486031215610b6d57600080fd5b833592506020840135610b7f81610ae8565b929592945050506040919091013590565b6000816000190483118215151615610baa57610baa610aa0565b500290565b600082610bcc57634e487b7160e01b600052601260045260246000fd5b500490565b81810381811115610ac957610ac9610aa0565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220d6ca2ec514951e1ad6290de24655b92882f5ccedf2f60814880e79a62acac9f164736f6c63430008100033",
}

// OnePoolRewardOldABI is the input ABI used to generate the binding from.
// Deprecated: Use OnePoolRewardOldMetaData.ABI instead.
var OnePoolRewardOldABI = OnePoolRewardOldMetaData.ABI

// OnePoolRewardOldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnePoolRewardOldMetaData.Bin instead.
var OnePoolRewardOldBin = OnePoolRewardOldMetaData.Bin

// DeployOnePoolRewardOld deploys a new Ethereum contract, binding an instance of OnePoolRewardOld to it.
func DeployOnePoolRewardOld(auth *bind.TransactOpts, backend bind.ContractBackend, lifetimeSeconds_ *big.Int) (common.Address, *types.Transaction, *OnePoolRewardOld, error) {
	parsed, err := OnePoolRewardOldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnePoolRewardOldBin), backend, lifetimeSeconds_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnePoolRewardOld{OnePoolRewardOldCaller: OnePoolRewardOldCaller{contract: contract}, OnePoolRewardOldTransactor: OnePoolRewardOldTransactor{contract: contract}, OnePoolRewardOldFilterer: OnePoolRewardOldFilterer{contract: contract}}, nil
}

// OnePoolRewardOld is an auto generated Go binding around an Ethereum contract.
type OnePoolRewardOld struct {
	OnePoolRewardOldCaller     // Read-only binding to the contract
	OnePoolRewardOldTransactor // Write-only binding to the contract
	OnePoolRewardOldFilterer   // Log filterer for contract events
}

// OnePoolRewardOldCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnePoolRewardOldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardOldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnePoolRewardOldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardOldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnePoolRewardOldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardOldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnePoolRewardOldSession struct {
	Contract     *OnePoolRewardOld // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnePoolRewardOldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnePoolRewardOldCallerSession struct {
	Contract *OnePoolRewardOldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OnePoolRewardOldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnePoolRewardOldTransactorSession struct {
	Contract     *OnePoolRewardOldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OnePoolRewardOldRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnePoolRewardOldRaw struct {
	Contract *OnePoolRewardOld // Generic contract binding to access the raw methods on
}

// OnePoolRewardOldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnePoolRewardOldCallerRaw struct {
	Contract *OnePoolRewardOldCaller // Generic read-only contract binding to access the raw methods on
}

// OnePoolRewardOldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnePoolRewardOldTransactorRaw struct {
	Contract *OnePoolRewardOldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnePoolRewardOld creates a new instance of OnePoolRewardOld, bound to a specific deployed contract.
func NewOnePoolRewardOld(address common.Address, backend bind.ContractBackend) (*OnePoolRewardOld, error) {
	contract, err := bindOnePoolRewardOld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardOld{OnePoolRewardOldCaller: OnePoolRewardOldCaller{contract: contract}, OnePoolRewardOldTransactor: OnePoolRewardOldTransactor{contract: contract}, OnePoolRewardOldFilterer: OnePoolRewardOldFilterer{contract: contract}}, nil
}

// NewOnePoolRewardOldCaller creates a new read-only instance of OnePoolRewardOld, bound to a specific deployed contract.
func NewOnePoolRewardOldCaller(address common.Address, caller bind.ContractCaller) (*OnePoolRewardOldCaller, error) {
	contract, err := bindOnePoolRewardOld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardOldCaller{contract: contract}, nil
}

// NewOnePoolRewardOldTransactor creates a new write-only instance of OnePoolRewardOld, bound to a specific deployed contract.
func NewOnePoolRewardOldTransactor(address common.Address, transactor bind.ContractTransactor) (*OnePoolRewardOldTransactor, error) {
	contract, err := bindOnePoolRewardOld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardOldTransactor{contract: contract}, nil
}

// NewOnePoolRewardOldFilterer creates a new log filterer instance of OnePoolRewardOld, bound to a specific deployed contract.
func NewOnePoolRewardOldFilterer(address common.Address, filterer bind.ContractFilterer) (*OnePoolRewardOldFilterer, error) {
	contract, err := bindOnePoolRewardOld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardOldFilterer{contract: contract}, nil
}

// bindOnePoolRewardOld binds a generic wrapper to an already deployed contract.
func bindOnePoolRewardOld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnePoolRewardOldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolRewardOld *OnePoolRewardOldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolRewardOld.Contract.OnePoolRewardOldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolRewardOld *OnePoolRewardOldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.OnePoolRewardOldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolRewardOld *OnePoolRewardOldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.OnePoolRewardOldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolRewardOld *OnePoolRewardOldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolRewardOld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolRewardOld *OnePoolRewardOldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolRewardOld *OnePoolRewardOldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) AccumulatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "accumulatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.AccumulatedReward(&_OnePoolRewardOld.CallOpts)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.AccumulatedReward(&_OnePoolRewardOld.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) ActiveDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "activeDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.ActiveDonation(&_OnePoolRewardOld.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.ActiveDonation(&_OnePoolRewardOld.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) ClaimedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "claimedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.ClaimedReward(&_OnePoolRewardOld.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.ClaimedReward(&_OnePoolRewardOld.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) FirstValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "firstValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.FirstValidChunk(&_OnePoolRewardOld.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.FirstValidChunk(&_OnePoolRewardOld.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolRewardOld *OnePoolRewardOldSession) Initialized() (bool, error) {
	return _OnePoolRewardOld.Contract.Initialized(&_OnePoolRewardOld.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) Initialized() (bool, error) {
	return _OnePoolRewardOld.Contract.Initialized(&_OnePoolRewardOld.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) LastUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "lastUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LastUpdateTimestamp(&_OnePoolRewardOld.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LastUpdateTimestamp(&_OnePoolRewardOld.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) LastValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "lastValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LastValidChunk(&_OnePoolRewardOld.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LastValidChunk(&_OnePoolRewardOld.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) LifetimeInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "lifetimeInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LifetimeInSeconds(&_OnePoolRewardOld.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.LifetimeInSeconds(&_OnePoolRewardOld.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldSession) Market() (common.Address, error) {
	return _OnePoolRewardOld.Contract.Market(&_OnePoolRewardOld.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) Market() (common.Address, error) {
	return _OnePoolRewardOld.Contract.Market(&_OnePoolRewardOld.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) Mine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "mine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldSession) Mine() (common.Address, error) {
	return _OnePoolRewardOld.Contract.Mine(&_OnePoolRewardOld.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) Mine() (common.Address, error) {
	return _OnePoolRewardOld.Contract.Mine(&_OnePoolRewardOld.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) NextChunkDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "nextChunkDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.NextChunkDonation(&_OnePoolRewardOld.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.NextChunkDonation(&_OnePoolRewardOld.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) TimeoutHead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "timeoutHead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.TimeoutHead(&_OnePoolRewardOld.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolRewardOld.Contract.TimeoutHead(&_OnePoolRewardOld.CallOpts)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolRewardOld *OnePoolRewardOldCaller) TimeoutRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	var out []interface{}
	err := _OnePoolRewardOld.contract.Call(opts, &out, "timeoutRecords", arg0)

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
func (_OnePoolRewardOld *OnePoolRewardOldSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolRewardOld.Contract.TimeoutRecords(&_OnePoolRewardOld.CallOpts, arg0)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolRewardOld *OnePoolRewardOldCallerSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolRewardOld.Contract.TimeoutRecords(&_OnePoolRewardOld.CallOpts, arg0)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactor) ClaimMineReward(opts *bind.TransactOpts, pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _OnePoolRewardOld.contract.Transact(opts, "claimMineReward", pricingIndex, beneficiary, arg2)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_OnePoolRewardOld *OnePoolRewardOldSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.ClaimMineReward(&_OnePoolRewardOld.TransactOpts, pricingIndex, beneficiary, arg2)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactorSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.ClaimMineReward(&_OnePoolRewardOld.TransactOpts, pricingIndex, beneficiary, arg2)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactor) FillReward(opts *bind.TransactOpts, beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolRewardOld.contract.Transact(opts, "fillReward", beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.FillReward(&_OnePoolRewardOld.TransactOpts, beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactorSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.FillReward(&_OnePoolRewardOld.TransactOpts, beforeLength, rewardSectors)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactor) Initialize(opts *bind.TransactOpts, market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolRewardOld.contract.Transact(opts, "initialize", market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolRewardOld *OnePoolRewardOldSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Initialize(&_OnePoolRewardOld.TransactOpts, market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactorSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Initialize(&_OnePoolRewardOld.TransactOpts, market_, mine_)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactor) Refresh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolRewardOld.contract.Transact(opts, "refresh")
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolRewardOld *OnePoolRewardOldSession) Refresh() (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Refresh(&_OnePoolRewardOld.TransactOpts)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactorSession) Refresh() (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Refresh(&_OnePoolRewardOld.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolRewardOld.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldSession) Receive() (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Receive(&_OnePoolRewardOld.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolRewardOld *OnePoolRewardOldTransactorSession) Receive() (*types.Transaction, error) {
	return _OnePoolRewardOld.Contract.Receive(&_OnePoolRewardOld.TransactOpts)
}

// OnePoolRewardOldDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the OnePoolRewardOld contract.
type OnePoolRewardOldDistributeRewardIterator struct {
	Event *OnePoolRewardOldDistributeReward // Event containing the contract specifics and raw log

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
func (it *OnePoolRewardOldDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnePoolRewardOldDistributeReward)
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
		it.Event = new(OnePoolRewardOldDistributeReward)
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
func (it *OnePoolRewardOldDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnePoolRewardOldDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnePoolRewardOldDistributeReward represents a DistributeReward event raised by the OnePoolRewardOld contract.
type OnePoolRewardOldDistributeReward struct {
	PricingIndex *big.Int
	Beneficiary  common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolRewardOld *OnePoolRewardOldFilterer) FilterDistributeReward(opts *bind.FilterOpts, pricingIndex []*big.Int, beneficiary []common.Address) (*OnePoolRewardOldDistributeRewardIterator, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _OnePoolRewardOld.contract.FilterLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardOldDistributeRewardIterator{contract: _OnePoolRewardOld.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolRewardOld *OnePoolRewardOldFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *OnePoolRewardOldDistributeReward, pricingIndex []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _OnePoolRewardOld.contract.WatchLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnePoolRewardOldDistributeReward)
				if err := _OnePoolRewardOld.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolRewardOld *OnePoolRewardOldFilterer) ParseDistributeReward(log types.Log) (*OnePoolRewardOldDistributeReward, error) {
	event := new(OnePoolRewardOldDistributeReward)
	if err := _OnePoolRewardOld.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
