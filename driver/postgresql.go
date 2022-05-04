package driver

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	rs     *sql.DB
	rsOnce sync.Once
)

func NewPostgresql() (*sql.DB, error) {
	var err error
	rsOnce.Do(func() {
		err = newPostgresql()
	})
	return rs, err
}

func newPostgresql() error {
	username := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	dbName := os.Getenv("POSTGRESQL_DBNAME")

	url := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		username,
		password,
		host,
		port,
		dbName)

	var err error
	if rs, err = sql.Open("postgres", url); err != nil {
		return fmt.Errorf("postgresql connect error : (%v)", err)
	}

	if err = rs.Ping(); err != nil {
		return fmt.Errorf("postgresql ping error : (%v)", err)
	}

	return nil
}
