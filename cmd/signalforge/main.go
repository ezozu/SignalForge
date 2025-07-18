// cmd/signalforge/main.go
package main

import (
"flag"
"log"
"os"

"signalforge/internal/signalforge"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := signalforge.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
