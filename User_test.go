package govalid
import (
  "testing"
)

var expectedResults = map[string]bool{
	"testy@tes.test":  true,
	"Chuck Norries":   false,
  "code@monkey.se":   true,
}

func TestUserEmail(t *testing.T){
  for k, v := range expectedResults {
    if result := Email(k); result != v {
			t.Errorf("Email(%s) returned %t, expected: %t", k, result, v)
		}
  }
}
