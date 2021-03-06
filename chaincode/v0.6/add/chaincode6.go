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

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var Value int
	var Time int64

	Value = 0
	Time = time.Now().UnixNano()

	// Write the state to the ledger
	err = stub.PutState("Value", []byte(strconv.Itoa(Value)))
	if err != nil {
		return nil, err
	}
	// Write the state to the ledger
	err = stub.PutState("Time", []byte(strconv.FormatInt(Time, 10)))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "add" {
		return t.add(stub, args)
	} else if function == "reset" {
		return t.reset(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var newTime int64
	temV, _ := stub.GetState("Value")
	V, _ := strconv.Atoi(string(temV))
	V++
	// Write the state to the ledger
	err = stub.PutState("Value", []byte(strconv.Itoa(V)))
	if err != nil {
		return nil, err
	}
	newTime = time.Now().UnixNano()
	// Write the state to the ledger
	err = stub.PutState("Time", []byte(strconv.FormatInt(newTime, 10)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) reset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	T := time.Now().UnixNano()
	V := 0
	//reset v = 0

	// Write the state to the ledger
	err = stub.PutState("Value", []byte(strconv.Itoa(V)))
	if err != nil {
		return nil, err
	}
	// Write the state to the ledger
	err = stub.PutState("Time", []byte(strconv.FormatInt(T, 10)))
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
