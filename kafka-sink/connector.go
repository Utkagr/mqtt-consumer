package main


type connector string

func (c connector) Sink(msg string) {
	
}

// exported
var Connector connector