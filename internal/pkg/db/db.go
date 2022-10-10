package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ziqiongg/hackernews/internal/links"
	"github.com/ziqiongg/hackernews/internal/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// // Db is our database struct used for interacting with the database
// var Db *sql.DB

// // New makes a new database using the connection string and
// // returns it, otherwise returns the error
// func New(connString string){
// 	db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check that our connection is good
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Db = db
// }

// func CloseDB() error{
// 	return Db.Close()
// }

// // ConnString returns a connection string based on the parameters it's given
// // This would normally also contain the password, however we're not using one
// func ConnString(host string, port int, user string, dbName string) string {
// 	return fmt.Sprintf(
// 		"host=%s port=%d user=%s dbname=%s sslmode=disable",
// 		host, port, user, dbName,
// 	)
// }

func NewConnection() (*gorm.DB,error) {
	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=carolli password=password dbname=hackernews sslmode=disable",
		// config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&users.User{},&links.Link{})

	return DB,nil
}


