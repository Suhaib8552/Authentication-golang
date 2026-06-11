package migrations

import (
	"auth-go/database"
	"log"
)

func CreateTables() {

	if database.DB == nil {
		panic("database.DB is nil")
	}

	usersTable := `
	CREATE TABLE IF NOT EXISTS USERS(
	id SERIAL PRIMARY KEY,
	username VARCHAR(100) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	tokenBlacklistTable := `
	CREATE TABLE IF NOT EXISTS token_blacklist(
	id SERIAL PRIMARY KEY,
	token TEXT NOT NULL UNIQUE,
	expires_at TIMESTAMP NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := database.DB.Exec(usersTable)

	if err != nil {
		log.Fatal("failed to create users table:", err)
	}

	_, err = database.DB.Exec(tokenBlacklistTable)

	if err != nil {
		log.Fatal("failed to create token_blacklist table:", err)
	}

	log.Println("Tables created successfully")

}
