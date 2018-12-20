// +build models

package main

import "testing"

func TestReply(t *testing.T) {
	ins := []string{
		"World",
		"You",
	}
	exps := []string{
		"Hello world",
		"Hello you",
	}
	for i, input := range ins {
		act := Reply(input)
		if exps[i] != act {
			t.Errorf("expected output '%v'; got '%v'", exps[i], act)
		}
	}
}
