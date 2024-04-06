package main

import (
	"fmt"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"github.com/n0rad/go-erlog/logs"
)
import _ "github.com/n0rad/go-erlog/register-json"

func main() {
	logs.SetLevel(logs.TRACE)

	logs.Trace("I'm trace")
	logs.Debug("I'm debug")
	logs.Info("I'm info")
	logs.Warn("I'm warn")
	logs.WithE(errs.WithE(errs.WithF(data.WithField("name", "value"), "source2"), "source")).WithField("file", "toto").Error("I'm error")
	logs.WithE(fmt.Errorf("genre")).WithField("file", "toto").Error("I'm error")

	func() {
		defer func() { recover() }()
		func() { logs.Panic("I'm panic") }()
	}()

	logs.Fatal("I'm fatal")
}
