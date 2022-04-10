package main

import (
	"azura-test/src/business/domain"
	"azura-test/src/business/usecases"
	"azura-test/src/handler/rest"
	"azura-test/src/utils/config"
	"azura-test/src/utils/configreader"
	logger "azura-test/src/utils/logger"
	"azura-test/src/utils/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		Name: "conf",
		Type: "yaml",
		Path: "../etc/cfg",
	})
	configreader.ReadConfig(&cfg)

	log := logger.Init()

	db := sql.Init(cfg.SQL, log)

	d := domain.Init(log, db)

	uc := usecases.Init(log, d)

	r := rest.Init(cfg, configreader, log, uc)
	r.Run()
}
