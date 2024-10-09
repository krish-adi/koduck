package main

import (
	"github.com/krish-adi/koduck/cmd"
	"github.com/krish-adi/koduck/db"
	"github.com/krish-adi/koduck/paths"
)

func main() {
	paths.InitPaths()

	db.InitDB()
	db.AttachMD()

	cmd.Execute()

	defer db.CloseDB()
}
