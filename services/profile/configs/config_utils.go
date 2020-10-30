package configs

var cfg = Config{
	Database: dbConfig{
		Address:  "localhost",
		Username: "postgres",
		Password: "admin",
		//Database:   "profile",
		Sslmode:    "disable",
		Drivername: "postgres",
	},
}

// GetTestConfig returns config for tests
func GetTestConfig() (Config, error) {
	return cfg, nil
}
