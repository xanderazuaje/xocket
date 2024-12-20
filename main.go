package main

import (
	"flag"
	"github.com/goccy/go-yaml"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"github.com/xanderazuaje/xocket/run"
	"io"
	"log"
	"os"
)

var ExitCode = 0

func main() {
	flag.Parse()
	flags.ValidateFlags()
	file, err := os.Open(flags.This.ConfigFile)
	if err != nil {
		log.Fatal("error opening "+flags.This.ConfigFile+": ", err.Error())
	}
	raw, err := io.ReadAll(file)
	_ = file.Close()
	var program parsing.Program
	err = yaml.Unmarshal(raw, &program)
	if err != nil {
		log.Fatal(err.Error())
	}
	if program.Endpoint == "" {
		log.Fatal("endpoint is required")
	}
	ok := run.All(program)
	if !ok {
		ExitCode = 1
	}
	os.Exit(ExitCode)
}
