package logdna

import (
	"fmt"
	"github.com/arundo/fabric-service-message-processor/src/utils"
	"github.com/gogap/logrus_mate"
	log "github.com/sirupsen/logrus"
)

func Setup(logLevel string, apiKey string) {
	utils.SetLogLevel(logLevel)
	log.Infoln("Log level:", log.GetLevel())

	if apiKey != "" {
		log.Println("Setting up LogDNA logger")

		config := fmt.Sprintf(`{
    out.name = "stdout"
    level = "debug"

    formatter.name = "json"
    formatter.options  {
        force-colors      = false
        disable-colors    = true
        disable-timestamp = false
        full-timestamp    = false
        timestamp-format  = "2006-01-02 15:04:05"
        disable-sorting   = false
    }

    hooks {
        logdna {
            api-key = "%s"
            app = "tag-monitor"
            flush = 1s
            json = true
        }
    }
}`,
			apiKey,
		)

		err := logrus_mate.Hijack(
			log.StandardLogger(),
			logrus_mate.ConfigString(config),
		)
		if err != nil {
			log.WithError(err).Error("Failed to configure LogDNA logger")
		}
	}
}
