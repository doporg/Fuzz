/*
Copyright 2014 Google Inc. All rights reserved.

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
package unit

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func TestFunc(t *testing.T) {
	fmt.Println("---------------------Test begin---------------------")
	cc := new(SimpleChaincode)                 // 创建Chaincode对象
	stub := shimtest.NewMockStub("source", cc) // 创建MockStub对象
	res1 := stub.MockInit("1", [][]byte{[]byte("")})
	if res1.Status != 200 {
		fmt.Println("init failed")
	}
	res2 := stub.MockInvoke("1", [][]byte{[]byte("initMarble"), []byte("asdf"), []byte("blue"), []byte("35"), []byte("bob")})
	if res2.Status != 200 {
		fmt.Println("initmock failed")
	}

	fmt.Println("---------------------Test finished---------------------")
}
