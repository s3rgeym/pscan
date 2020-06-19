package main

import (
	"os"
	"pscan/internal"

	"github.com/jessevdk/go-flags"
)

func main() {
	opts := &internal.Options{}
	p := flags.NewParser(opts, flags.Default)
	p.Usage = "[options] arguments\n\nTCP port scanner."
	args, err := p.Parse()
	if err != nil {
		os.Exit(1)
	}
	internal.Run(opts, args)
}
