package main

import "fmt"

/* Another use case of iface:
   You can use iface type as input for function
   to simplify mocking of an input argument, for ex. Database client!
   Its extremely useful in scope of unit testing!
   Tools like `mockgen` already implements such feature during mock
   autogeneration!
*/

// SpyOrganization is super secret service,
// thats really hard to deploy during unit tests
type SpyOrganization struct{}

// SpyToolTransmitter is one of manymore spy tools,
// that have API client to the SpyOrganization
// thats really hard to deploy during unit tests
type SpyToolTransmitter struct {
	Client SpyOrganization
}

// Listen represents how transmitter listen voices
func (t SpyToolTransmitter) Listen(target string) (string, error) {
	fmt.Printf("SHHHHSHS....listening target: %v...\n", target)

	return "secret information", nil
}

// Listener is iface that can allow us to union all objects that
// implements function Listen!
type Listener interface {
	Listen(string) (string, error)
}

// Report represents spy report like recorded information after Eavesdrop
type Report struct {
	Value string
}

/* we may declare `(t SpyToolTransmitter)` as input argument by default
   but its hard to mock in Unit test because of complex SpyOrganization
   so we declaring `(liface Listener)` to have an ability of simple exchange
   real client with the fake one!!
*/

// Eavesdrop represent Linten&Recording process, that Spy is doing
func Eavesdrop(liface Listener, target string) (Report, error) {
	var r Report

	s, err := liface.Listen(target)
	if err != nil {
		return r, err
	}
	r.Value = s

	return r, nil
}

func main() {
	// note: below is hard to mock!
	var spyT SpyToolTransmitter

	result, err := Eavesdrop(spyT, "president")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result is: %v\n", result)
}
