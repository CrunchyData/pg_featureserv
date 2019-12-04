package config

// Configuration for system
var Configuration Config

func init() {
	Configuration = Config{
		Server: Server{
			BindHost: "",
			BindPort: 10000,
		},
		Metadata: Metadata{
			Title:       "PG-FeatureServ",
			Description: "A PostGIS Feature Server",
		},
		Database: Database{
			ConnectString: "",
		},
	}
}

// Config for system
type Config struct {
	Server   Server
	Metadata Metadata
	Database Database
}

// Server config
type Server struct {
	BindHost string `toml:"bind_host"`
	BindPort int    `toml:"bind_port"`
}

// Database config
type Database struct {
	ConnectString string
}

// Metadata config
type Metadata struct {
	Title       string
	Description string
}
