package main

import (
	"fmt"
)

type connector string

func (c connector) Sink(msg string) {
	fmt.Println(msg)
}

// exported
var Connector connector