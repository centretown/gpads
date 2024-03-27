package main

import (
	"flag"
	"fmt"
	"gpads/gcmd"
)

func init() {
	gcmd.SetFlagsVars()
}

func main() {
	flag.Parse()
	cmds := gcmd.NewCmds()
	gcmd.RunJoyCmds(cmds)
	fmt.Println("done!")

}
