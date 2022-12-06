package initializer

import (
	"ic-indexer-service/app/config"
	"ice-cream-indexer-service/app/config/locales/local_config"
	"log"
)

func Initialize() {

	initializeConfig()

}

func initializeConfig() {

	err := config.Initialize()
	local_config.InitializeLocalizer()
	config.InitializeDecoderAndValidator()
	config.InitializeElasticSearchClient()

	if err != nil {
		log.Fatal(nil, "error initializing Config :", err)
		return
	}

}
