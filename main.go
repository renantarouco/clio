package main

import (
	"fmt"
	"log"
	"os"

	be "github.com/renantarouco/clio/backend"
	"github.com/renantarouco/clio/kv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	command := args[1]
	commandArgs := args[2:]

	fsBackend := be.NewFSBackend("./kv.db")

	kv, err := kv.NewKV(fsBackend)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "set":
		if len(commandArgs) < 2 {
			log.Fatal("not enough arguments for 'set' command")
		}

		key, value := commandArgs[0], commandArgs[1]

		kv.Set(key, value)
	case "get":
		if len(commandArgs) < 1 {
			log.Fatal("not enough arguments for 'get' command")
		}

		key := commandArgs[0]

		value := kv.Get(key)

		fmt.Println(value)
	default:
		log.Fatal("unsupported command")
	}
}
