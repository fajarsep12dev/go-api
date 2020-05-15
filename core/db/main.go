package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Connection Provider
	"github.com/qor/audited"
)

// ORM Struct
type ORM struct {
	DB     *gorm.DB
}

// Initialize the database instance
func (orm *ORM) Initialize() {
	var err error

	dbDriver(
		orm,
		os.Getenv("DB_PG_DRIVER"),
		os.Getenv("DB_PG_USER"),
		os.Getenv("DB_PG_PASSWORD"),
		os.Getenv("DB_PG_PORT"),
		os.Getenv("DB_PG_HOST"),
		os.Getenv("DB_PG_NAME"),
	)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return os.Getenv("DB_TABLE_PREFIX") + defaultTableName
	}

	orm.DB.SingularTable(true)
	audited.RegisterCallbacks(orm.DB)
}

func dbDriver (orm *ORM, Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		orm.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		orm.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	
}