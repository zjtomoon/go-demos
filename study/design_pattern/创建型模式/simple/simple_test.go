package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("tom")
	if s != "Hi,tom" {
		t.Fatal("type1 test fail")
	}
}
func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("tom")
	if s != "Hello,tom" {
		t.Fatal("type2 test fail")
	}
}
