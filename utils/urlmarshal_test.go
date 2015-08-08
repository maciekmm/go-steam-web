package utils

import "testing"

func TestURLMarshal(t *testing.T) {
	str := struct {
		Boolean  bool
		String   string
		Integer  int64
		UInteger uint64 `uval:"uint"`
	}{
		true, "hey", -10, 10,
	}
	marshaled := ToURLValues(&str)
	if marshaled.Get("Boolean") != "1" {
		t.Error("Error marshaling boolean")
	}
	if marshaled.Get("String") != "hey" {
		t.Error("Error marshaling string")
	}
	if marshaled.Get("Integer") != "-10" {
		t.Error("Error marshaling integer")
	}
	if marshaled.Get("uint") != "10" {
		t.Error("Error marshaling uint or tag")
	}
}
