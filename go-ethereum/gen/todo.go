// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110fa806100606000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630f560cd7146100675780634cc8221514610085578063751ef753146100a15780639507d39a146100bd578063b0c8f9dc146100ed578063f745630f14610109575b600080fd5b61006f610125565b60405161007c9190610921565b60405180910390f35b61009f600480360381019061009a919061098d565b610231565b005b6100bb60048036038101906100b6919061098d565b6103ac565b005b6100d760048036038101906100d2919061098d565b61048a565b6040516100e491906109f7565b60405180910390f35b61010760048036038101906101029190610b4e565b610586565b005b610123600480360381019061011e9190610b97565b61065c565b005b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610228578382906000526020600020906002020160405180604001604052908160008201805461017c90610c22565b80601f01602080910402602001604051908101604052809291908181526020018280546101a890610c22565b80156101f55780601f106101ca576101008083540402835291602001916101f5565b820191906000526020600020905b8154815290600101906020018083116101d857829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610149565b50505050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028957600080fd5b80600180549050811061029b57600080fd5b60008290505b600180805490506102b29190610c82565b81101561035857600180826102c79190610cb6565b815481106102d8576102d7610cea565b5b9060005260206000209060020201600182815481106102fa576102f9610cea565b5b90600052602060002090600202016000820181600001908161031c9190610edb565b506001820160009054906101000a900460ff168160010160006101000a81548160ff0219169083151502179055509050506001810190506102a1565b50600180548061036b5761036a610fc3565b5b60019003818190600052602060002090600202016000808201600061039091906106fe565b6001820160006101000a81549060ff0219169055505090555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461040457600080fd5b80600180549050811061041657600080fd5b6001828154811061042a57610429610cea565b5b906000526020600020906002020160010160009054906101000a900460ff16156001838154811061045e5761045d610cea565b5b906000526020600020906002020160010160006101000a81548160ff0219169083151502179055505050565b61049261073e565b8160018054905081106104a457600080fd5b600183815481106104b8576104b7610cea565b5b90600052602060002090600202016040518060400160405290816000820180546104e190610c22565b80601f016020809104026020016040519081016040528092919081815260200182805461050d90610c22565b801561055a5780601f1061052f5761010080835404028352916020019161055a565b820191906000526020600020905b81548152906001019060200180831161053d57829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581525050915050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105de57600080fd5b6001604051806040016040528083815260200160001515815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190816106369190610ff2565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106b457600080fd5b8160018054905081106106c657600080fd5b81600184815481106106db576106da610cea565b5b906000526020600020906002020160000190816106f89190610ff2565b50505050565b50805461070a90610c22565b6000825580601f1061071c575061073b565b601f01602090049060005260206000209081019061073a919061075a565b5b50565b6040518060400160405280606081526020016000151581525090565b5b8082111561077357600081600090555060010161075b565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156107dd5780820151818401526020810190506107c2565b60008484015250505050565b6000601f19601f8301169050919050565b6000610805826107a3565b61080f81856107ae565b935061081f8185602086016107bf565b610828816107e9565b840191505092915050565b60008115159050919050565b61084881610833565b82525050565b6000604083016000830151848203600086015261086b82826107fa565b9150506020830151610880602086018261083f565b508091505092915050565b6000610897838361084e565b905092915050565b6000602082019050919050565b60006108b782610777565b6108c18185610782565b9350836020820285016108d385610793565b8060005b8581101561090f57848403895281516108f0858261088b565b94506108fb8361089f565b925060208a019950506001810190506108d7565b50829750879550505050505092915050565b6000602082019050818103600083015261093b81846108ac565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61096a81610957565b811461097557600080fd5b50565b60008135905061098781610961565b92915050565b6000602082840312156109a3576109a261094d565b5b60006109b184828501610978565b91505092915050565b600060408301600083015184820360008601526109d782826107fa565b91505060208301516109ec602086018261083f565b508091505092915050565b60006020820190508181036000830152610a1181846109ba565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a5b826107e9565b810181811067ffffffffffffffff82111715610a7a57610a79610a23565b5b80604052505050565b6000610a8d610943565b9050610a998282610a52565b919050565b600067ffffffffffffffff821115610ab957610ab8610a23565b5b610ac2826107e9565b9050602081019050919050565b82818337600083830152505050565b6000610af1610aec84610a9e565b610a83565b905082815260208101848484011115610b0d57610b0c610a1e565b5b610b18848285610acf565b509392505050565b600082601f830112610b3557610b34610a19565b5b8135610b45848260208601610ade565b91505092915050565b600060208284031215610b6457610b6361094d565b5b600082013567ffffffffffffffff811115610b8257610b81610952565b5b610b8e84828501610b20565b91505092915050565b60008060408385031215610bae57610bad61094d565b5b6000610bbc85828601610978565b925050602083013567ffffffffffffffff811115610bdd57610bdc610952565b5b610be985828601610b20565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c3a57607f821691505b602082108103610c4d57610c4c610bf3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c8d82610957565b9150610c9883610957565b9250828203905081811115610cb057610caf610c53565b5b92915050565b6000610cc182610957565b9150610ccc83610957565b9250828201905080821115610ce457610ce3610c53565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081549050610d2881610c22565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610d917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d54565b610d9b8683610d54565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610dd8610dd3610dce84610957565b610db3565b610957565b9050919050565b6000819050919050565b610df283610dbd565b610e06610dfe82610ddf565b848454610d61565b825550505050565b600090565b610e1b610e0e565b610e26818484610de9565b505050565b5b81811015610e4a57610e3f600082610e13565b600181019050610e2c565b5050565b601f821115610e8f57610e6081610d2f565b610e6984610d44565b81016020851015610e78578190505b610e8c610e8485610d44565b830182610e2b565b50505b505050565b600082821c905092915050565b6000610eb260001984600802610e94565b1980831691505092915050565b6000610ecb8383610ea1565b9150826002028217905092915050565b818103610ee9575050610fc1565b610ef282610d19565b67ffffffffffffffff811115610f0b57610f0a610a23565b5b610f158254610c22565b610f20828285610e4e565b6000601f831160018114610f4f5760008415610f3d578287015490505b610f478582610ebf565b865550610fba565b601f198416610f5d87610d2f565b9650610f6886610d2f565b60005b82811015610f9057848901548255600182019150600185019450602081019050610f6b565b86831015610fad5784890154610fa9601f891682610ea1565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610ffb826107a3565b67ffffffffffffffff81111561101457611013610a23565b5b61101e8254610c22565b611029828285610e4e565b600060209050601f83116001811461105c576000841561104a578287015190505b6110548582610ebf565b8655506110bc565b601f19841661106a86610d2f565b60005b828110156110925784890151825560018201915060208501945060208101905061106d565b868310156110af57848901516110ab601f891682610ea1565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220ca75d424273027525def9b009d55b76f20608fdd14992be606adea8075b0041664736f6c63430008110033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
