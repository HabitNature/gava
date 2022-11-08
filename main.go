package main

import "fmt"

var version = 1.0

func main() {
	cmd := parseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		fmt.Printf("version %f", version)
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// todo
}
