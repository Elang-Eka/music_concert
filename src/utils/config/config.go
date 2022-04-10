package config

import "azura-test/src/utils/sql"

type Application struct {
	Gin  GinConfig
	SQL  sql.Config
	Meta ApplicationMeta
}

func Init() Application {
	return Application{}
}

type GinConfig struct {
	Port    string
	Swagger SwaggerConfig
}

type ApplicationMeta struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
}

type SwaggerConfig struct {
	Enabled   bool
	Path      string
	BasicAuth BasicAuthConf
}

type BasicAuthConf struct {
	Username string
	Password string
}
