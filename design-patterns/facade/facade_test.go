package facade

import "testing"

var expect = "A, B"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Call()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
