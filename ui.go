package main

import (
	"fmt"
	"os"
)

func statusf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "dino: "+format+"\n", args...)
}
