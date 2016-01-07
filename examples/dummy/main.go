package main
import (
	"github.com/n0rad/go-erlog/log"
)

// go-erlog/log is an api, by default it uses a very basic dummy logger
func main() {
	log.Trace("I'm trace")
	log.Debug("I'm debug")
	log.Info("I'm info")
	log.Warn("I'm warn")
	log.Error("I'm error")

	func() {
		defer func() {recover()}()
		func() {log.Panic("I'm panic")}()
	}()

	log.Fatal("I'm fatal")
}

