package config

// AppConfiguration is the set of global application configuration constants.
type AppConfiguration struct {
	// AppName name of the software
	Name string
	// AppVersion version number of the software
	Version  string
	EnvDBURL string
}

var AppConfig = AppConfiguration{
	Name:     "pg_featureserv",
	Version:  "0.1",
	EnvDBURL: "PG_FS_DATABASE_URL",
}
