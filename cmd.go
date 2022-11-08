package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XJreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XJreOption, "XJreOption", "", "path to jre")

	flag.Parse()

	return &cmd
}

func printUsage() {
	fmt.Printf("Usage : %s [-options] class [args...]", os.Args[0])
}
