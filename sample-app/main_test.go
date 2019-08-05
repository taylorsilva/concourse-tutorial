package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello world!"
	output := getHelloWorld()
	if output != expected {
		t.Errorf("Got %s when expecting %s", output, expected)
	}
}
