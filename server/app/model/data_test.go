package model

import (
	"fmt"
	"testing"
)

func TestAdmins(t *testing.T) {
	if err := DataAdmins(); err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func TestCasbinRule(t *testing.T) {
	if err := DataCasbinRule(); err != nil {
		panic(err)
	}
	fmt.Println("success")
}
