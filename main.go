package main

import (
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	now := fmt.Sprint(time.Now().Unix())
	_, _ = fmt.Fprintln(os.Stdout, now)
	if len(os.Args) != 2 || os.Args[1] != "--no-copy" {
		err := clipboard.WriteAll(now)
		if err != nil {
			panic(err)
		}
	}
}
