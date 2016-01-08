package register

import (
	"github.com/n0rad/go-erlog"
	"github.com/n0rad/go-erlog/log"
)

func init() {
	log.RegisterLoggerFactory(&erlog.ErrLogsFactory{})
}
