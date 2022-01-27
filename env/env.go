package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var REQUIRED_ENV_VARS = []string{"ENDPOINT_URL", "POLYGON_ENDPOINT", "HASURA_API_KEY", "RAT_ADDR", "CLOSET_ADDR"}

var RAT_ADDR string
var CLOSET_ADDR string
var ENDPOINT_URL string
var POLYGON_ENDPOINT string
var HASURA_API_KEY string

func Init() {
	godotenv.Load()
	missingEnvs := make([]string, 0)
	for _, val := range REQUIRED_ENV_VARS {
		if os.Getenv(val) == "" {
			missingEnvs = append(missingEnvs, val)
		}
	}
	if len(missingEnvs) != 0 {
		log.Fatalf("missing %v env var(s)", missingEnvs)
	}
	ENDPOINT_URL = os.Getenv("ENDPOINT_URL")
	POLYGON_ENDPOINT = os.Getenv("POLYGON_ENDPOINT")
	HASURA_API_KEY = os.Getenv("HASURA_API_KEY")
	RAT_ADDR = os.Getenv("RAT_ADDR")
	CLOSET_ADDR = os.Getenv("CLOSET_ADDR")
}
