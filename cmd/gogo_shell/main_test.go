package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollectInput(t *testing.T) {
	tests := []struct {
		input    string
		wantCmd  string
		wantArgs []string
	}{
		{"cd", "cd", []string{}},
		{"ls -latr", "ls", []string{"-latr"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("parseInput(%s)", tt.input)
		t.Run(testname, func(t *testing.T) {
			gotCmd, gotArgs := parseInput(tt.input)
			if (gotCmd != tt.wantCmd) && (!reflect.DeepEqual(gotArgs, tt.wantArgs)) {
				t.Errorf("parseInput(%s) = %s, %s; want %s, %s", tt.input, gotCmd, gotArgs, tt.wantCmd, tt.wantArgs)
			}
		})
	}
}
