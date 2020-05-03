package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Connection Provider
)

// ORM Struct
type ORM struct {
	DB     *gorm.DB
}

// Initialize the database instance
func (orm *ORM) Initialize() {
	var err error

	orm.DB, err = gorm.Open(os.Getenv("DB_TYPE"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return os.Getenv("DB_TABLE_PREFIX") + defaultTableName
	}

	orm.DB.SingularTable(true)
}