package main

import (
	"os"
)

func main() {
    app := NewApp(os.Args[1:])
    app.Run()
}
