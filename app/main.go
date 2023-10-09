package main

import (
	"log"
	"w8protector/binding"
)

func main() {
	// example of binding
	err := binding.BindMachine()
	if err != nil {
		log.Fatalf("FATAL ERR: failed to bind machine err:%v", err)
	}
}
