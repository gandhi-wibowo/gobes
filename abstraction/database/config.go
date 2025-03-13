package database

import (
	"gobes/abstraction/config"
	"time"
)

type Database struct {
	Dialect      string        // Database.Dialect
	LogMode      bool          // Database.LogMode
	Separation   bool          // Database.Separation
	MaxOpenConns int           // Database.MaxOpenConns
	MaxIdleConns int           // Database.MaxIdleConns
	MaxLifetime  time.Duration // Database.MaxLifetime

	Parameters []string // Database.Parameters

	Primary Cluster // Database.Primary
}

type Cluster struct {
	DBName   string // Database.[Primary|Secondary].DBName
	Host     string // Database.[Primary|Secondary].Host
	Port     int    // Database.[Primary|Secondary].Database
	Username string // Database.[Primary|Secondary].Username
	Password string // Database.[Primary|Secondary].Password
}

func newDatabaseConfig() Database {
	cfg := config.NewViperConfig(nil)
	var params []string
	cfg.GetList("Database.Parameters", params)

	return Database{
		Dialect:      cfg.GetString("Database.Dialect", "postgres"),
		LogMode:      cfg.GetBool("Database.LogMode", true),
		Separation:   cfg.GetBool("Database.Separation", true),
		MaxOpenConns: cfg.GetInt("Database.MaxOpenConns", 100),
		MaxIdleConns: cfg.GetInt("Database.MaxIdleConns", 5),
		MaxLifetime:  cfg.GetDuration("Database.MaxLifetime", 210*time.Minute),
		Parameters:   params,
	}
}

func NewDatabaseConfig(path *string) Database {
	cfg := config.NewViperConfig(nil)
	basePath := "Database.Primary"
	if path != nil {
		basePath = *path
	}

	database := newDatabaseConfig()
	database.Primary = Cluster{
		DBName:   cfg.GetString(basePath+".DBName", ""),
		Host:     cfg.GetString(basePath+".Host", ""),
		Port:     cfg.GetInt(basePath+".Port", 5432),
		Username: cfg.GetString(basePath+".Username", ""),
		Password: cfg.GetString(basePath+".Password", ""),
	}
	return database
}
