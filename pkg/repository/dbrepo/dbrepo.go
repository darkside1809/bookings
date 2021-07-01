package dbrepo

import (
	// built in Golang packages
	"database/sql"
	// My own packages
	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/repository"
)

// testDBRepo holds pointers to configs and DB for testing
type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
// postgresDBRepo holds pointers to configs and DB
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
// NewPostgresRepo constructs new connection for DB, app configs, and returns DatabaseRepo interface
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo {
		App: a,
		DB: conn,
	}
}
// NewTestingRepo app configs, and returns DatabaseRepo interface for testing
func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo {
		App: a,
	}
}