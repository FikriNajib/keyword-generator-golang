package config

import (
	"encoding/json"
	"fmt"
	"github.com/sahalazain/go-common/config"
	"os"
)

var DefaultConfig = map[string]interface{}{
	"PORT":                      "",
	"REDIS_HOST":                "127.0.0.1:6379",
	"REDIS_PASSWORD":            "",
	"REDIS_DB":                  8,
	"MONGO_URL":                 "",
	"MONGO_DB":                  "userbehaviour",
	"MONGO_COLLECTION":          "behaviour",
	"MONGO_COLLECTION_PROCESS":  "behaviour_processed",
	"MONGO_COLLECTION_ARCHIVED": "behaviour_archived",
	"ELASTIC_SEARCH_HOST":       "",
	"ELASTIC_SEARCH_USER":       "elastic",
	"ELASTIC_SEARCH_PASS":       "",
	"VIEW_DURATION":             "10",
	"MONGO_TTL":                 172800,
	"REDIS_TTL":                 7,
	"DB_VIDEO_DRIVER":           "mysql",
	"DB_VIDEO_USER":             "root",
	"DB_VIDEO_PASSWORD":         "",
	"DB_VIDEO_HOST":             "localhost",
	"DB_VIDEO_DB":               "",
	"DB_HOT_DRIVER":             "mysql",
	"DB_HOT_USER":               "root",
	"DB_HOT_PASSWORD":           "",
	"DB_HOT_HOST":               "localhost",
	"DB_HOT_DB":                 "",
	"ELASTIC_APM_ACTIVE":        true,
	"ELASTIC_APM_ENVIRONMENT":   "dev",
	"ELASTIC_APM_SERVICE_NAME":  "keyword-generator",
	"ELASTIC_APM_SERVER_URL":    "",
	"ELASTIC_APM_SECRET_TOKEN":  "",
}

var Config config.Getter
var Url string

func Load() error {
	conf, err := os.ReadFile("config.json")
	var configmap map[string]interface{}

	// Unmarshal the JSON data into the data variable
	err = json.Unmarshal(conf, &configmap)
	if err != nil {
		fmt.Println("Error:", err)
	}
	cfgClient, err := config.Load(configmap, Url)
	if err != nil {
		return err
	}

	Config = cfgClient

	return nil
}
