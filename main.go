package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintln(os.Stdout, time.Now().Unix())
}
