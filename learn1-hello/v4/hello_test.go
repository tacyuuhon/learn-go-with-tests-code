package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hello func test",
			args: args{name: "World"},
			want: "Hello, World",
		},
		{
			name: "Hello func test",
			args: args{name: "Panda"},
			want: "Hello, Panda",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
