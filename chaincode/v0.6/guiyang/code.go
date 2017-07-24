/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//募资结构

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//初始化的时候传入参数有6个：募资结构编号，计划募资总金额，第一顺位（json字符串），第二顺位，第三顺位，操作人编号。顺序以这个为准。
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var fundRaisingID string	//募资结构编号
	var Sum string	//计划募资总金额
	var Prority1 string	//第一顺位
	var Prority2 string	//第二顺位
	var Prority3 string	//第三顺位

	var err error

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}

	// Initialize the chaincode
	fundRaisingID = args[0]
	Sum = args[1]
	Prority1 = args[2]
	Prority2 = args[3]
	Prority3 = args[4]

	// Write the state to the ledger
	err = stub.PutState("fundRaisingID", []byte(fundRaisingID))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Sum", []byte(Sum))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority1", []byte(Prority1))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority2", []byte(Prority2))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority3", []byte(Prority3))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "update" {
		return t.update(stub, args)
	}

	return nil, errors.New("no such a method on this chaincode")
}

//更新募资结构传入参数有6个：募资结构编号，计划募资总金额，第一顺位（json字符串），第二顺位，第三顺位，操作人编号。
func (t *SimpleChaincode) update(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var newFundRaisingID string	//募资结构编号
	var newSum string	//计划募资总金额
	var newPrority1 string	//第一顺位
	var newPrority2 string	//第二顺位
	var newPrority3 string	//第三顺位

	var err error

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}

	// Initialize the chaincode
	newFundRaisingID = args[0]
	newSum = args[1]
	newPrority1 = args[2]
	newPrority2 = args[3]
	newPrority3 = args[4]

	// Write the state to the ledger
	err = stub.PutState("fundRaisingID", []byte(newFundRaisingID))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Sum", []byte(newSum))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority1", []byte(newPrority1))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority2", []byte(newPrority2))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("Prority3", []byte(newPrority3))
	if err != nil {
		return nil, err
	}

	return nil, nil
}


// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
