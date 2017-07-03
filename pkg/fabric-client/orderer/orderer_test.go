/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package orderer

import (
	"fmt"
	"net"
	"testing"
	"time"

	api "github.com/tuxago/fabric-sdk-go/api"
	client "github.com/tuxago/fabric-sdk-go/pkg/fabric-client"
	mocks "github.com/tuxago/fabric-sdk-go/pkg/fabric-client/mocks"

	ab "github.com/hyperledger/fabric/protos/orderer"
	"google.golang.org/grpc"
)

var testOrdererURL = "0.0.0.0:4584"

//
// Orderer via chain setOrderer/getOrderer
//
// Set the orderer URL through the chain setOrderer method. Verify that the
// orderer URL was set correctly through the getOrderer method. Repeat the
// process by updating the orderer URL to a different address.
//
func TestOrdererViaChain(t *testing.T) {
	client := client.NewClient(mocks.NewMockConfig())
	chain, err := client.NewChannel("testChain-orderer-member")
	if err != nil {
		t.Fatalf("error from NewChain %v", err)
	}
	orderer, _ := NewOrderer("localhost:7050", "", "", mocks.NewMockConfig())
	err = chain.AddOrderer(orderer)
	if err != nil {
		t.Fatalf("Error adding orderer: %v", err)
	}

	orderers := chain.GetOrderers()
	if orderers == nil || len(orderers) != 1 || orderers[0].GetURL() != "localhost:7050" {
		t.Fatalf("Failed to retieve the new orderer URL from the chain")
	}
	chain.RemoveOrderer(orderer)
	orderer2, err := NewOrderer("localhost:7054", "", "", mocks.NewMockConfig())
	if err != nil {
		t.Fatalf("Failed to create NewOrderer error(%v)", err)
	}
	err = chain.AddOrderer(orderer2)
	if err != nil {
		t.Fatalf("Error adding orderer: %v", err)
	}
	orderers = chain.GetOrderers()

	if orderers == nil || len(orderers) != 1 || orderers[0].GetURL() != "localhost:7054" {
		t.Fatalf("Failed to retieve the new orderer URL from the chain")
	}

}

//
// Orderer via chain missing orderer
//
// Attempt to send a request to the orderer with the sendTransaction method
// before the orderer URL was set. Verify that an error is reported when tying
// to send the request.
//
func TestPeerViaChainMissingOrderer(t *testing.T) {
	client := client.NewClient(mocks.NewMockConfig())
	chain, err := client.NewChannel("testChain-orderer-member2")
	if err != nil {
		t.Fatalf("error from NewChain %v", err)
	}
	_, err = chain.SendTransaction(nil)
	if err == nil {
		t.Fatalf("SendTransaction didn't return error")
	}
	if err.Error() != "orderers is nil" {
		t.Fatalf("SendTransaction didn't return right error")
	}

}

//
// Orderer via chain nil data
//
// Attempt to send a request to the orderer with the sendTransaction method
// with the data set to null. Verify that an error is reported when tying
// to send null data.
//
func TestOrdererViaChainNilData(t *testing.T) {
	client := client.NewClient(mocks.NewMockConfig())
	chain, err := client.NewChannel("testChain-orderer-member2")
	if err != nil {
		t.Fatalf("error from NewChain %v", err)
	}
	orderer, err := NewOrderer("localhost:7050", "", "", mocks.NewMockConfig())
	if err != nil {
		t.Fatalf("Failed to create NewOrderer error(%v)", err)
	}
	err = chain.AddOrderer(orderer)
	if err != nil {
		t.Fatalf("Error adding orderer: %v", err)
	}

	_, err = chain.SendTransaction(nil)
	if err == nil {
		t.Fatalf("SendTransaction didn't return error")
	}
	if err.Error() != "proposal is nil" {
		t.Fatalf("SendTransaction didn't return right error")
	}
}

func TestSendDeliver(t *testing.T) {
	mockServer := startMockServer(t)
	orderer, _ := NewOrderer(testOrdererURL, "", "", mocks.NewMockConfig())
	// Test deliver happy path
	blocks, errors := orderer.SendDeliver(&api.SignedEnvelope{})
	select {
	case block := <-blocks:
		if string(block.Data.Data[0]) != "test" {
			t.Fatalf("Expected test block got: %#v", block)
		}
	case err := <-errors:
		t.Fatalf("Unexpected error from SendDeliver(): %s", err)
	case <-time.After(time.Second * 5):
		t.Fatalf("Did not receive block or error from SendDeliver")
	}

	// Test deliver without valid envelope
	blocks, errors = orderer.SendDeliver(nil)
	select {
	case block := <-blocks:
		t.Fatalf("Expected error got block: %#v", block)
	case err := <-errors:
		if err == nil {
			t.Fatalf("Expected error with nil envelope")
		}
	case <-time.After(time.Second * 5):
		t.Fatalf("Did not receive block or error from SendDeliver")
	}

	// Test deliver with deliver error from OS
	testError := fmt.Errorf("test error")
	mockServer.DeliverError = testError
	blocks, errors = orderer.SendDeliver(&api.SignedEnvelope{})
	select {
	case block := <-blocks:
		t.Fatalf("Expected error got block: %#v", block)
	case err := <-errors:
		if err == nil {
			t.Fatalf("Expected test error when OS Recv() fails, got: %s", err)
		}
	case <-time.After(time.Second * 5):
		t.Fatalf("Did not receive block or error from SendDeliver")
	}
}

func startMockServer(t *testing.T) *mocks.MockBroadcastServer {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", testOrdererURL)
	broadcastServer := new(mocks.MockBroadcastServer)
	ab.RegisterAtomicBroadcastServer(grpcServer, broadcastServer)
	if err != nil {
		fmt.Printf("Error starting test server %s", err)
		t.FailNow()
	}
	fmt.Printf("Starting test server\n")
	go grpcServer.Serve(lis)

	return broadcastServer
}
