package database

import "fmt"

type Config struct {
	ServerName string
	User       string
	Pass       string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User,
		config.Pass,
		config.ServerName,
		config.DB)

	return connectionString
}
