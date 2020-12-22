package simpleFactory

import "testing"

func TestNetApiFetch(t *testing.T) {
	api := NewApi("net")
	res := api.Fetch("net")
	if res != "hi netApi net" {
		t.Fatal("TestNetApi_Fetch test fail")
	}
}