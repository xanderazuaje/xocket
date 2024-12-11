package main

import (
	"flag"
	"fmt"
	"github.com/xanderazuaje/xocket/flags"
)

func main() {
	flag.Parse()
	fmt.Printf("%+v\n", flags.This)
	flags.ValidateFlags()
}
