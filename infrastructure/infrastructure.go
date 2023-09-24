package infrastructure

import (
	"clean-arch/config"
	"clean-arch/domain/entity"
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.Config) *Database {
	var err error
	var sqlDB *sql.DB

	gormDB, err := getDatabaseInstance(config)
	if err != nil {
		for i := 0; i < 5; i++ {
			gormDB, err = getDatabaseInstance(config)
			if err == nil {
				break
			}
		}
	}
	// try to connect again

	db := &Database{gormDB}

	db.RegisterTables()

	sqlDB, err = db.DB.DB()
	if err != nil {
		log.Fatal("sqlDB connection error ", err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func getDatabaseInstance(config *config.Config) (db *gorm.DB, err error) {
	switch config.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Name,
			config.Database.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}

		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	}
	return db, nil
}

func (d Database) RegisterTables() {
	err := d.DB.AutoMigrate(
		&entity.Article{},
	)

	if err != nil {
		log.Fatal("Database migration error", err.Error())
		os.Exit(0)
	}
	log.Print("Database migration success")
}
