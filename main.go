package main

import (
	"fmt"
	"log"

	"github.com/OzkrOssa/mac-acl-go/devices"
)

func main() {
	mk, err := devices.Mikrotik("", "", "")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mk)

}
