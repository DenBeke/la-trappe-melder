package register

import (
	"fmt"

	"github.com/xo/dburl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Register is the wrapper struct around the batch database
type Register struct {
	db *gorm.DB
}

// New creates a new register database
func New(dbURL string) (*Register, error) {

	u, err := dburl.Parse(dbURL)

	if u.Driver != "sqlite3" {
		return nil, fmt.Errorf("unsupported schema: %s", u.Driver)
	}

	db, err := gorm.Open(sqlite.Open(u.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&Batch{})
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return &Register{db: db}, nil

}

// Close closes the database connection
func (r *Register) Close() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
