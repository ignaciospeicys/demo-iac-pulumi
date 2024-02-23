package setup

import (
	"fmt"
	"github.com/joeshaw/envdecode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type postgresConfig struct {
	Host     string `env:"db_postgres_db_host,default=127.0.0.1"`
	Port     string `env:"db_postgres_db_port,default=5432"`
	Name     string `env:"db_postgres_db_name,default=pulumi_resource_db"`
	User     string `env:"db_postgres_db_user,default=postgres"`
	Password string `env:"db_postgres_db_password,default=123456"`
	PoolIdle int    `env:"db_postgres_pool_idle,default=5"`
	PoolMax  int    `env:"db_postgres_pool_max,default=60"`
}

var params = postgresParams()

func postgresParams() postgresConfig {
	config := postgresConfig{}
	err := envdecode.Decode(&config)
	if err != nil {
		log.Fatalf("--Config:postgres:PostgresParams --Error [message=%s]", err)
	}
	return config
}

func CreateGormDBConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", params.Host, params.User, params.Password, params.Name, params.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	// Optionally, configure the underlying SQL DB for connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get database: %v", err)
	}
	sqlDB.SetMaxIdleConns(params.PoolIdle)
	sqlDB.SetMaxOpenConns(params.PoolMax)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}
