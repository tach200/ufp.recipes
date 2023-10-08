package db

import "github.com/jmoiron/sqlx"

type DB struct {
	driver string
	opts   string
	Client *sqlx.DB
}

func NewDB(driver, opts string) (*DB, error) {
	dbClient, err := sqlx.Connect(driver, opts)
	if err != nil {
		return nil, err
	}

	DB := &DB{
		driver: driver,
		opts:   opts,
		Client: dbClient,
	}

	return DB, nil
}
