package main

import (
	"flag"
	"github.com/goccy/go-yaml"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"io"
	"log"
	"os"
)

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
	parsing.GetFilter("<integer:max=2>")
}
