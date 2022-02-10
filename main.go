package main

import (
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

	if cmd.CurrCommand.Watch {
		watcher.Watch()
	} else {
		historical.Query()
	}
}
