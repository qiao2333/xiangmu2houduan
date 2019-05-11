package mytest

import (
	"fmt"
	"project/project3/sql/impl"
	"testing"
)

func Test_cao(test *testing.T) {
	state := new(impl.AddressstateImpl)
	result := state.GetState()
	fmt.Println(result)
}
