package main
import (
	"github.com/n0rad/go-erlog/log"
	_ "github.com/n0rad/go-erlog"
	"os"
	"github.com/n0rad/go-erlog/with"
)

func main() {
	path := "/toto/config"
	if err := os.Mkdir(path, 0777); err != nil {
		log.LogEntry(&log.Entry{
			Fields: with.Field("dir", path),
			Level: log.INFO,
			Error: err,
			Message: "Salut !1",
		})
	}

	log.GetDefaultLog().Info("Salut !2")
	log.GetLog("other").Info("Salut !3")

}
