package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var REQUIRED_ENV_VARS = []string{"HASURA_ENDPOINT", "POLYGON_ENDPOINT", "HASURA_ADMIN_SECRET", "RAT_ADDR", "CLOSET_ADDR"}

var RAT_ADDR string
var CLOSET_ADDR string
var HASURA_ENDPOINT string
var POLYGON_ENDPOINT string
var HASURA_ADMIN_SECRET string

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
	HASURA_ENDPOINT = os.Getenv("HASURA_ENDPOINT")
	POLYGON_ENDPOINT = os.Getenv("POLYGON_ENDPOINT")
	HASURA_ADMIN_SECRET = os.Getenv("HASURA_ADMIN_SECRET")
	RAT_ADDR = os.Getenv("RAT_ADDR")
	CLOSET_ADDR = os.Getenv("CLOSET_ADDR")
}
