package iredigo

import (
	"testing"

	"fmt"

	"github.com/garyburd/redigo/redis"
)

func TestSet(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1",
			args: args{
				name:  "test1",
				value: "ribenren",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args.name, tt.args.value, 30); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
			result := Get("test1")
			fmt.Printf("result = %s\n", result)
		})
	}
}

func TestSetFromConn(t *testing.T) {
	type args struct {
		c     redis.Conn
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetFromConn(tt.args.c, tt.args.name, tt.args.value, 30); got != tt.want {
				t.Errorf("SetFromConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
