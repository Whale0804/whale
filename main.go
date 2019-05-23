package main

import (
	"github.com/githinkcn/whale/cmd"
	_ "github.com/githinkcn/whale/routers"
)

func main() {
	cmd.Execute()
}
