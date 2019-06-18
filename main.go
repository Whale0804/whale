package main

import (
	"fmt"
	"github.com/githinkcn/whale/cmd"
	_ "github.com/githinkcn/whale/routers"
	"github.com/githinkcn/whale/utils"
)

func main() {
	fmt.Println(utils.GetCurrentPath())
	cmd.Execute()
}
