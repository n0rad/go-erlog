package main

import (
	_ "github.com/n0rad/go-erlog/register"
	"github.com/n0rad/go-erlog/log"
	"github.com/n0rad/go-erlog/with"
	"os"
	"github.com/n0rad/go-erlog"
)

func main() {
	logger := log.GetLog("newlog") // another logger

	logger.(*erlog.ErlogLogger).Appenders[0].(*erlog.ErlogWriterAppender).Out = os.Stdout


	path := "/toto/config"
	if err := os.Mkdir(path, 0777); err != nil {
		logger.LogEntry(&log.Entry{
			Fields:  with.Field("dir", path),
			Level:   log.INFO,
			Error:   err,
			Message: "Salut !1",
		})
	}

	logger.Info("Salut !2")
	logger.Info("Salut !3")

}
