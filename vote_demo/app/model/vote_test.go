package model

import (
	"fmt"
	"testing"
)

func TestGetVotes(t *testing.T) {
	NewMysql()

	r := GetVotes()
	fmt.Printf("ret: %+v", r)
	Close()
}

func TestGetVote(t *testing.T) {
	NewMysql()

	r := GetVote(1)
	fmt.Printf("ret: %+v", r)
	Close()
}

func TestDoVote(t *testing.T) {
	NewMysql()

	r := DoVote(2, 1, []int64{3, 4})
	fmt.Printf("ret: %+v", r)
	Close()
}
