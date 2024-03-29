package main

import "testing"

func Test_add(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("1 + 2 should be 3")
	}
}
