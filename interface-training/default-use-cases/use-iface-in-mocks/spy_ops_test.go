package main

import (
	"fmt"
	"testing"
)

/* Here we want to test Eavesdrop function
   But we dont want to deploy whole SpyOrganization Service!
   We need to mock it somehow and forward the stub into Eavesdrop!
   thats why we need iface as input argument for Eavesdrop!!
*/

// MockSpyOrganization is a fake SpyOrganization
type MockSpyOrganization struct{}

// MockSpyToolTransmitter is a fake SpyToolTransmitter
type MockSpyToolTransmitter struct {
	Client MockSpyOrganization
}

// Listen implements original Listen func for SpyToolTransmitter, but its faked!
func (mt MockSpyToolTransmitter) Listen(target string) (string, error) {
	fmt.Printf("SHHHHSHS....FAKE: %v...\n", target)

	return "FAKE: secret information", nil
}

// TestEavesdrop tests generation of a spy report
func TestEavesdrop(t *testing.T) {
	var fakeSpyT MockSpyToolTransmitter

	result, err := Eavesdrop(fakeSpyT, "fake-president")
	if err != nil {
		t.Errorf("Testing error: %v", err)
	}

	t.Logf("Test result: %v\n", result)
}
