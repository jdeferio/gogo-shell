package main

import (
	"fmt"
	"testing"
)

func TestCollectInput(t *testing.T) {

	tests := []struct {
		in        string
		want_cmd  string
		want_args []string
	}{
		{"cd", "cd", []string{}},
		{"ls -latr", "ls", []string{"-latr"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("CollectInput(%d)", tt.in)
		t.Run(testname, func(t *testing.T) {
			got_cmd, got_args := collectInput(tt.in)
			if (got_cmd != tt.want_cmd) & (got_args != tt.want_args) {
				t.Errorf("CollectInput(%d) = %d, %d; want %d, %d", tt.in, got_cmd, got_args, tt.cmd, tt.args)
			}
		})
	}
}
