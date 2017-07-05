package main

import (
	"os"
	"github.com/SRI-CSL/gllvm/shared"
)

func main() {
	// Parse command line
	var args = os.Args
	args = args[1:]

	
	exitCode := shared.Compile(args, "clang++")
	

	shared.LogInfo("Calling %v returned %v\n", os.Args, exitCode)

	//important to pretend to look like the actual wrapped command
	os.Exit(exitCode)

}