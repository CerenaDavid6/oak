package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
)

//go:embed cmd/version.oak
var cmdversion string

//go:embed cmd/help.oak
var cmdhelp string

//go:embed cmd/cat.oak
var cmdcat string

//go:embed cmd/fmt.oak
var cmdfmt string

//go:embed cmd/build.oak
var cmdbuild string

var cliCommands = map[string]string{
	"version": cmdversion,
	"help":    cmdhelp,
	"cat":     cmdcat,
	"fmt":     cmdfmt,
	"build":   cmdbuild,
}

func performCommandIfExists(command string) bool {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get working directory")
		os.Exit(1)
	}

	ctx := NewContext(path.Dir(cwd))
	ctx.LoadBuiltins()
	defer ctx.Wait()

	commandProgram, ok := cliCommands[command]
	if !ok {
		return false
	}

	_, err = ctx.Eval(strings.NewReader(commandProgram))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return true
}
