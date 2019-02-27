package main

import (
	"log"
)

func main() {
	s := thing.NewStuff() 
	log.Fatal(s.TestThing().Error())
}
