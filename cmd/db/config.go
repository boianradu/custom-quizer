package db

import "os"

var PGGlobal = map[string]string{}

func initDBVariables() {
	PGGlobal["PG_HOST"] = os.Getenv("PQL_HOST")
	PGGlobal["PG_PORT"] = os.Getenv("PQL_PORT")
	PGGlobal["PG_USER"] = os.Getenv("PQL_USER")
	PGGlobal["PG_PASSWORD"] = os.Getenv("PQL_PASSWORD")
	PGGlobal["PG_DBNAME"] = os.Getenv("PQL_DBNAME")
}
