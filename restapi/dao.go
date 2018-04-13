package restapi

import "os"

const (
	ENV_DB_URL = "DB_URL"
	ENV_DB_NAME = "DB_NAME"
)

var (
	dbUrl = os.Getenv(ENV_DB_URL)
)