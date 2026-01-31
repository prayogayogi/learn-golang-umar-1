package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	if connectionString == "" {
		log.Println("ERROR: Database connection string is empty. Please check your .env file.")
		return nil, sql.ErrConnDone
	}

	log.Println("Attempting to connect to database...")
	
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("ERROR: Failed to open database connection: %v\n", err)
		return nil, err
	}

	log.Println("Testing database connection...")
	err = db.Ping()
	if err != nil {
		log.Printf("ERROR: Failed to ping database: %v\n", err)
		log.Println("Please verify:")
		log.Println("  1. Your database credentials are correct")
		log.Println("  2. The database server is accessible")
		log.Println("  3. Your connection string includes ?sslmode=require for Supabase")
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
