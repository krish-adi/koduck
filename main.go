package main

import (
	"github.com/krish-adi/koduck/cmd"
	"github.com/krish-adi/koduck/db"
	"github.com/krish-adi/koduck/paths"
	"github.com/krish-adi/koduck/prompt"
)

func main() {
	paths.InitPaths()
	prompt.InitClients()

	db.InitDB()
	db.AttachMD()

	cmd.Execute()

	defer db.CloseDB()
}
