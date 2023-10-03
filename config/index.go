package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IConfig interface {
	Init() (db *gorm.DB, errors []error)
	loadEnvironments() (err error)
	migrateModels(db *gorm.DB, dst ...interface{}) (err error)
	configurationDatabase() (db *gorm.DB, err error)
}
type ConfigurationApplication struct {
}

// Init implements IConfig.
func (configurationApplication *ConfigurationApplication) Init() (db *gorm.DB, errors []error) {
	if os.Getenv("ENV") != "production" {
		err := configurationApplication.loadEnvironments()
		if err != nil {
			errors = append(errors, err)
		}
	}
	db, errDatabase := configurationApplication.configurationDatabase()
	if errDatabase != nil {
		errors = append(errors, errDatabase)
	}
	errMigration := configurationApplication.migrateModels(db, initialModels...)
	if errMigration != nil {
		errors = append(errors, errMigration)
	}

	return
}

// configurationDatabase implements IConfig.
func (configurationApplication *ConfigurationApplication) configurationDatabase() (db *gorm.DB, err error) {
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	user := string(os.Getenv("DB_USERNAME"))
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	return
}

// loadEnvironments implements IConfig.
func (configurationApplication *ConfigurationApplication) loadEnvironments() (err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file ", err.Error())
	}
	return
}

// syncDatabase implements IConfig.
func (configurationApplication *ConfigurationApplication) migrateModels(db *gorm.DB, dst ...interface{}) (err error) {
	err = db.AutoMigrate(dst...)
	if err != nil {
		fmt.Println("Failed to sync database", err.Error())
	}
	return
}

func NewConfigurationApplication() IConfig {
	return &ConfigurationApplication{}
}
