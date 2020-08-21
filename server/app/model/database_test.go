package model

import (
	"fmt"
	"testing"
)

func TestTableAdmins(t *testing.T) {
	if err := TableAdmins(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}
