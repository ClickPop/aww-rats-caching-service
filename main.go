package main

import (
	"log"

	"github.com/clickpop/aww-rats-caching-service/blockchain"
	"github.com/clickpop/aww-rats-caching-service/cmd"
	"github.com/clickpop/aww-rats-caching-service/env"
	"github.com/clickpop/aww-rats-caching-service/historical"
	"github.com/clickpop/aww-rats-caching-service/watcher"
)

func initialize() {
	env.Init()
	blockchain.Init()
	cmd.Init()
}

func main() {
	initialize()

	switch {
	case cmd.CurrCommand.Cmd == "watch":
		watcher.Watch(0)
	case cmd.CurrCommand.Cmd == "historical":
		historical.Query()
	default:
		log.Fatal("please provide a command")
	}
}
