package main

import (
	stdlog "log"
	"os"

	kitlog "github.com/go-kit/kit/log"
)

func main() {
	logger := kitlog.NewJSONLogger(os.Stdout)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
	stdlog.Print("I sure like pie")
}
