package main

import (
	_ "github.com/n0rad/go-erlog"
	"github.com/n0rad/go-erlog/log"
	"github.com/n0rad/go-erlog/with"
	"os"
)

func main() {
	path := "/toto/config"
	if err := os.Mkdir(path, 0777); err != nil {
		log.LogEntry(&log.Entry{
			Fields:  with.Field("dir", path),
			Level:   log.INFO,
			Error:   err,
			Message: "Salut !1",
		})
	}

	log.GetDefaultLog().Info("Salut !2")
	log.GetLog("other").Info("Salut !3")

}
