package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/yornifpaz/back_noteapp/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IConfig interface {
	Init() (db *gorm.DB, errors []error)
	GetDefaultEmailConfig() (emailConfig models.EmailConfig)
	GetDefaultDatabaseConfig() (config models.DatabaseConfig)
	ConfigurationDatabase(config models.DatabaseConfig) (db *gorm.DB, err error)
}
type ConfigurationApplication struct {
}

// GetDefaultDatabaseConfig implements IConfig.
func (configurationApplication *ConfigurationApplication) GetDefaultDatabaseConfig() (config models.DatabaseConfig) {
	return getDefaultDatabaseConfig()
}

// GetDefaultEmailConfig implements IConfig.
func (configurationApplication *ConfigurationApplication) GetDefaultEmailConfig() (emailConfig models.EmailConfig) {
	return getDefaultEmailConfig()
}

// init implements IConfig.
func (configurationApplication *ConfigurationApplication) Init() (db *gorm.DB, errors []error) {
	return configurationApplication.init()

}

// ConfigurationDatabase implements IConfig.
func (configurationApplication *ConfigurationApplication) ConfigurationDatabase(config models.DatabaseConfig) (db *gorm.DB, err error) {
	return configurationApplication.configurationDatabase(config)

}

// Init implements IConfig.
func (configurationApplication *ConfigurationApplication) init() (db *gorm.DB, errors []error) {
	if os.Getenv("ENV") != "production" {
		err := configurationApplication.loadEnvironments()
		if err != nil {
			errors = append(errors, err)
			panic(err.Error())
		}
	}
	configDatabase := getDefaultDatabaseConfig()
	db, errDatabase := configurationApplication.configurationDatabase(configDatabase)
	if errDatabase != nil {
		errors = append(errors, errDatabase)
		panic(errDatabase.Error())
	}
	errMigration := configurationApplication.migrateModels(db, initialModels...)
	if errMigration != nil {
		errors = append(errors, errMigration)
	}

	return
}

// configurationDatabase implements IConfig.
func (configurationApplication *ConfigurationApplication) configurationDatabase(config models.DatabaseConfig) (db *gorm.DB, err error) {

	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing environment variables for database configuration")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("Failed to connect to database : " + err.Error())

	}
	return
}

// loadEnvironments implements IConfig.
func (configurationApplication *ConfigurationApplication) loadEnvironments() (err error) {
	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file : " + err.Error())

	}
	return
}

// syncDatabase implements IConfig.
func (configurationApplication *ConfigurationApplication) migrateModels(db *gorm.DB, dst ...interface{}) (err error) {
	err = db.AutoMigrate(dst...)
	if err != nil {
		return fmt.Errorf("Failed to sync database : " + err.Error())
	}
	return
}

func NewConfigurationApplication() IConfig {
	return &ConfigurationApplication{}
}
