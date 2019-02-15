package dbwrapper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // here
)

/*
	sqlStatement := `
		DROP TABLE web_links;
		CREATE TABLE web_links (
		id SERIAL PRIMARY KEY,
		title TEXT UNIQUE,
		link TEXT NOT NULL
		);`
*/

func connectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[DB]Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[DB] Successfully connected!")
	return db
}

/*
 */
func ExecuteQuery(sqlStatement string) {
	db := connectDB()
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("[DB] Successfully Executed!")

}

/*
 */
func CreateWebLink(title, link string) {
	db := connectDB()
	defer db.Close()
	sqlStatement := `
	INSERT INTO web_links (title,link)
	VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, title, link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[DB] Created web link successfully")
}
