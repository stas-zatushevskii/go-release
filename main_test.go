package main

import "testing"

func TestPrintHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Hello World", want: "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintHello(); got != tt.want {
				t.Errorf("PrintHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
