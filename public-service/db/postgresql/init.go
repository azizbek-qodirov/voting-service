package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	cf "public-service/config"
	"public-service/db/postgresql/managers"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB      *sql.DB
	PartyI  PartyI
	PublicI PublicI
}

func ConnectDB(config cf.Config) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	party_manager := managers.NewPartyManager(db)
	public_manager := managers.NewPublicManager(db)
	log.Println("Successfully connected to the db pgsql!")
	return &Storage{
		DB:      db,
		PartyI:  party_manager,
		PublicI: public_manager,
	}, nil
}
