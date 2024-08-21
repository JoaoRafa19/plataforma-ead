package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func GetDb() *sql.DB {
	godotenv.Load(".env")
	
	
		
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		)
	
	db,err := sql.Open("postgres", connString)
	
	if err != nil {
		log.Panic(err)
	}
	
	return db
}