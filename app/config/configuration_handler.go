package config

import (
	"github.com/gorilla/schema"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
	"strings"
)

var (
	config                 *Config
	requestParamsValidator *validator.Validate
	requestParamsDecoder   *schema.Decoder
	elasticClient          *elastic.Client
)

type Config struct {
	Server        ServerConfig
	ElasticSearch ElasticSearchConfiguration
}

type ESIndex struct {
	IndexName string
	Type      string
}

type ElasticSearchConfiguration struct {
	URL     string
	Indices struct {
		IcecreamIndex ESIndex
	}
}

type ServerConfig struct {
	Port string
}

func GetConfig() *Config {
	return config
}

func InitializeElasticSearchClient() {
	log.Print("Initializing Elastic Search connection")
	var err error

	elasticClient, err = elastic.NewClient(
		elastic.SetURL(config.ElasticSearch.URL),
		elastic.SetSniff(false), //TODO: Decide whether to enable/disable. https://github.com/olivere/elastic/wiki/Sniffing
		elastic.SetHttpClient(nil),
	)
	if err != nil {
		errorMsg := "SERVER_STARTUP : Error during initialization of ES: %s"
		log.Fatalf(errorMsg, err)
	}
	log.Print("Elasticsearch initiated successfully.")

	esversion, err := elasticClient.ElasticsearchVersion(config.ElasticSearch.URL)
	if err != nil {
		errorMsg := "SERVER_STARTUP : Error during validation of ES: %s"
		log.Fatalf(errorMsg, err)
	}
	log.Print("Elasticsearch version ", esversion)
}

func Initialize() error {
	viper.SetConfigFile(getConfigFile())
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv((strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))))
		}
	}
	return viper.Unmarshal(&config)
}

func GetESConnection() *elastic.Client {
	return elasticClient
}

func GetReqParamsValidator() *validator.Validate {
	return requestParamsValidator
}

func GetReqParamsDecoder() *schema.Decoder {
	return requestParamsDecoder
}

func getConfigFile() string {
	return "app/config/" + os.Getenv("ENV") + ".yml"
}

func InitializeDecoderAndValidator() {
	log.Print("Initializing params decoder & validator")
	requestParamsDecoder = schema.NewDecoder()
	requestParamsValidator = validator.New()
	requestParamsDecoder.IgnoreUnknownKeys(true)
}
