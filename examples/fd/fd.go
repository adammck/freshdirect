package main

import (
	"fmt"
	"os"

	"github.com/adammck/freshdirect/cart"
	"github.com/adammck/freshdirect/session"
)

func main() {
	k := "FRESHDIRECT_SESSION_ID"
	sid := os.Getenv(k)
	if sid == "" {
		fmt.Printf("%s must be set\n")
		os.Exit(1)
	}

	s := session.New(sid)

	if len(os.Args) < 2 {
		fmt.Printf("usage: fd sku [sku..]\n")
		os.Exit(1)
	}

	for i := 0; i < len(os.Args); i++ {
		cart.Add(s, os.Args[i])
	}
}
