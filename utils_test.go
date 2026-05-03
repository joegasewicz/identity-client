package identity_client

import "testing"

func TestAddTokenPrefix(t *testing.T) {
	token := "<TOKEN>"
	result := AddTokenPrefix(token)
	expected := "Bearer " + token
	if result != expected {
		t.Logf("expected result to be %s but got %s", expected, result)
		t.Fail()
	}

	token = "Bearer <TOKEN>"
	result = AddTokenPrefix(token)
	if result != expected {
		t.Logf("expected result to be %s but got %s", token, result)
		t.Fail()
	}
}
