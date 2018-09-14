package healthcheck

import "testing"

// func TestLookUp(t *testing.T) {
// 	err := LookUp("http://google.com")
// 	if err != nil {
// 		t.Fail()
// 	}
// }

func TestGet(t *testing.T) {
	c := TracingClient{}
	resp := Get(c, "http://google.com")
	if resp == nil {
		t.Fail()
	}
}
