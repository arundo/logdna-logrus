package setup_logdna

import (
	"fmt"

	"github.com/arundo/fabric-service-message-processor/src/utils"
	"github.com/gogap/logrus_mate"
	log "github.com/sirupsen/logrus"

	_ "github.com/arundo/logdna-logrus"
)

func Setup(logLevel string, apiKey string, appName string) {
	utils.SetLogLevel(logLevel)
	log.Infoln("Log level:", log.GetLevel())

	if apiKey != "" {
		log.Println("Setting up LogDNA logger")

		config := fmt.Sprintf(`elon {
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
            app = "%s"
            flush = 1s
            json = true
        }
    }
}`,
			apiKey,
			appName,
		)

		mate, _ := logrus_mate.NewLogrusMate(
			logrus_mate.ConfigString(config),
			//logrus_mate.ConfigFile("logdna.conf"),
			//logrus_mate.ConfigProvider(&config.HOCONConfigProvider{}), // default provider
		)

		err := mate.Hijack(
			log.StandardLogger(),
			"elon",
		)
		if err != nil {
			log.WithError(err).Error("Error when configuring LogDNA logger")
		}
	} else {
		fmt.Println("Logging to std")
	}
}
