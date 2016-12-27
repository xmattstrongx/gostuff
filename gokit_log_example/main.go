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
	stdlog.Print("I wish commits to forked repos showed up in github profiles")
}
