package cmd

import (
	"log"
	"os"
	"strings"
)

type Command struct {
	Watch        bool
	Historical   bool
	Rats         bool
	ClosetPieces bool
	ClosetTokens bool
	Sync         bool
	IgnoreCache  bool
}

var CurrCommand Command

func Init() {
	if len(os.Args) < 1 {
		log.Println()
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
	CurrCommand.CheckSubCommands(subCommands)
	CurrCommand.CheckFlags(flags)
	if CurrCommand.Historical && CurrCommand.Watch {
		log.Fatal("please select either historical or watch")
	}
}

func (c *Command) CheckSubCommands(subCommands []string) {
	for _, subCommand := range subCommands {
		switch subCommand {
		case "historical":
			CurrCommand.Historical = true
		case "watch":
			CurrCommand.Watch = true
		}
	}
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
