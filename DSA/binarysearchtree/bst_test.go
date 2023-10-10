package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("mahesh")

	if got != "hi mahesh, how are you?" {
		t.Errorf("got %s", got)
	}
}
