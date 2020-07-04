package models

import (
	"fiber-rest-api/lib"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB { return db }

func InitDB() {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(connString), nil)
	lib.CheckErr(err, "trying to connect pg database")

	sync(&syncOptions{Force: true})
}

type syncOptions struct {
	Force bool
}

func sync(options *syncOptions) {
	if options.Force {
		err := db.Migrator().DropTable(new(UserModel))
		lib.CheckErr(err, "trying to drop tables")
	}

	err := db.Migrator().AutoMigrate(new(UserModel))
	lib.CheckErr(err, "migrating models")
}

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
