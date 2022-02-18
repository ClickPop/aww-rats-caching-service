package cmd

import (
	"log"
	"os"
	"strings"
)

type Command struct {
	Cmd          string
	Rats         bool
	ClosetPieces bool
	ClosetTokens bool
	Sync         bool
	IgnoreCache  bool
}

var CurrCommand Command

func Init() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a command")
	}
	args := os.Args[1:]
	subCommands := make([]string, 0)
	flags := make([]string, 0)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			flags = append(flags, arg)
		} else {
			subCommands = append(subCommands, arg)
		}
	}
	CurrCommand.Cmd = subCommands[0]
	CurrCommand.CheckFlags(flags)
}

func (c *Command) CheckFlags(flags []string) {
	for _, flag := range flags {
		switch flag {
		case "--rats":
			c.Rats = true
		case "--closet-pieces":
			c.ClosetPieces = true
		case "--closet-tokens":
			c.ClosetTokens = true
		case "--sync":
			c.Sync = true
		case "--ignore-cache":
			c.IgnoreCache = true
		}
	}
}
