package configdb

import (
	"fmt"
	"os"

	"github.com/arfan21/golang-mygram/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Ari23 dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(entity.User{}, entity.Photo{}, entity.Comment{}, entity.SocialMedia{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type pgConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func newPGConfig() *pgConfig {
	dbConfig := pgConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func (dbConfig *pgConfig) String() string {
	mode := os.Getenv("MODE")
	dsn := ""
	if mode == "production" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DBName,
			dbConfig.Port,
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DBName,
			dbConfig.Port,
		)
	}

	return dsn
}
