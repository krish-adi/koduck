package main

import (
	"github.com/krish-adi/koduck/cmd"
	"github.com/krish-adi/koduck/db"
	"github.com/krish-adi/koduck/paths"
)

func main() {
	db.InitDB()

	paths.InitPaths()
	cmd.Execute()

	defer db.CloseDB()
}
