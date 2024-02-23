package setup

import (
	"database/sql"
	"fmt"
	"github.com/joeshaw/envdecode"
	"github.com/lib/pq"
	sqltrace "github.com/signalfx/signalfx-go-tracing/contrib/database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type postgresConfig struct {
	Host     string `env:"db_postgres_db_host,default=127.0.0.1"`
	Port     string `env:"db_postgres_db_port,default=5432"`
	Name     string `env:"db_postgres_db_name,default=rpp-ms-resources"`
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

func CreateDBConnection() *sql.DB {
	sqltrace.Register("postgres", &pq.Driver{})

	db, err := sqltrace.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", params.User, params.Password, params.Host, params.Port, params.Name))
	if err != nil {
		log.Fatalf("--Config:postgres:postgresDatabase Could not open %v database: %v", params, err)
	}
	db.SetMaxIdleConns(params.PoolIdle)
	db.SetMaxOpenConns(params.PoolMax)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}

func InitGormDB(sqlDB *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not initialize GORM DB: %v", err)
	}
	return gormDB
}
