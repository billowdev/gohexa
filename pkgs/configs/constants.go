package configs

import "os"

var (
	TEMPLATE_URL = "https://github.com/rapidstellar/gohexa-template/releases/download/1.0.0/templates.zip"
	DATABASE     = os.Getenv("DATABASE")
	ORM          = os.Getenv("ORM")
	DB_ADAPTER   = os.Getenv("DB_ADAPTER")
)
