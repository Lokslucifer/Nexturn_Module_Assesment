package db
import (
	"database/sql"
	"log"
	_"modernc.org/sqlite"
)
var DB *sql.DB

func InitialiseDatabase(){
	var err error
	DB,err =sql.Open("sqlite","./Blog.db")
	if err!=nil{
		log.Fatal(err)
	}
	_,err=DB.Exec(`	CREATE TABLE IF NOT EXISTS 	blogs(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	author TEXT,
	content TEXT,
	timestamp DATETIME
	);`)
	if err!=nil {
		log.Fatal(err)
		
	}
	
}
func GetDB()*sql.DB{
	return DB
}
