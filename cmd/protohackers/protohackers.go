package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chigley/protohackers/challenge/smoke"
	"golang.org/x/sync/errgroup"
)

type ListenFunc func(address string) error

var challenges = []ListenFunc{
	smoke.Listen,
}

func main() {
	var (
		addr      = flag.String("addr", "localhost", "listen address")
		startPort = flag.Int("startPort", 8080, "listen port to be used for challenge 0")
	)
	flag.Parse()

	var g errgroup.Group

	for i, f := range challenges {
		f := f

		addr := fmt.Sprintf("%s:%d", *addr, *startPort+i)
		g.Go(func() error {
			return f(addr)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
