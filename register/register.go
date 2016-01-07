package register
import (
	"github.com/n0rad/go-erlog/log"
	"github.com/n0rad/go-erlog"
)

func init() {
	log.RegisterLoggerFactory(&erlog.ErrLogsFactory{})
}