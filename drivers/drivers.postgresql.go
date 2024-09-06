package drivers

import (
	"fmt"
	"log"
	"os"
	"test-pari/schemas"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	itemModel "test-pari/modules/items/model"
)

/*
? I use the GORM framework for database operations.
? If you compare it to a bare Go SQL driver, the bare driver is indeed faster.
? However, using GORM allows me to accomplish tasks that would take much longer with
? just the bare driver, such as automating migrations and writing simpler query code.
*/

func SetupDBSQL(config schemas.SchemaEnvironment) (*gorm.DB, error) {
	logrus.Debug("🔌Starting Create Database Postgres")
	CreateDB(config)
	logrus.Debug("🔌Finished Create Database Postgres")
	// panic(1)

	logrus.Debug("🔌 Connecting into Database Postgres")
	dbHost := config.DB_HOST
	dbUsername := config.DB_USER
	dbPassword := config.DB_PASS
	dbName := config.DB_NAME
	dbPort := config.DB_PORT
	dbSSLMode := config.DB_SSLMODE
	timezone := config.TIMEZONE

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		defer logrus.Errorln("❌ Error Connect into Database Postgres", err.Error())

		return nil, err
	}

	postgreSQL, err := db.DB()
	// Set connection pool parameters
	postgreSQL.SetMaxOpenConns(10)   // Maximum number of open connections
	postgreSQL.SetMaxIdleConns(5)    // Maximum number of idle connections
	postgreSQL.SetConnMaxLifetime(0) // Connection lifetime (0 means connections are reused indefinitely)

	if os.Getenv("GO_ENV") == "development" {
		db.Debug()
	}

	if err != nil {
		logrus.Errorln("❌ Error Migrate ", err.Error())
		return nil, err
	}

	fmt.Println("💚 Connect into Database Postgres Success")

	AutoMigrate(db)

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	// err = db.AutoMigrate(
	// 	//mTasklistDetail.TasklistDetail{},
	// 	//model3.TRVisit{},
	// 	//&entities.EntityVenue{},
	// 	sopModel.SOP{},
	// 	instructionModel.Instruction{},
	// 	confModel.Config{},
	// )

	//? Transaction for create table
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := itemModel.Items.Migrate(itemModel.Items{}, tx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Println(err)
		panic("Fail to Create Table")
	}

	//? Transaction for Insert table
	if err := db.Transaction(func(tx *gorm.DB) error {

		var count int64
		if err := tx.Raw(`SELECT count(id) FROM "items"`).Scan(&count).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if count == 0 {

			if err := tx.Exec(`INSERT INTO "items" ("CreatedTime","Name","Description","Price","Qty") VALUES (?,?,?,?,?)`, time.Now(), "Magicom", "Tempat Memasak Nasi", 300000, 10).Error; err != nil {
				logrus.Errorln("❌ Error Insert users : ", err.Error())
			}
			if err := tx.Exec(`INSERT INTO "items" ("CreatedTime","Name","Description","Price","Qty") VALUES (?,?,?,?,?)`, time.Now(), "Kompor", "Alat Memasak", 100000, 15).Error; err != nil {
				logrus.Errorln("❌ Error Insert users : ", err.Error())
			}
		}

		return nil
	}); err != nil {
		log.Println(err)
		panic("Fail to Create Table")
	}

}

func CreateDB(config schemas.SchemaEnvironment) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s", config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_PORT, config.DB_SSLMODE)
	// dsn := "host=localhost user=postgres password=mysecretpassword port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to the database: %v", err)
	}

	createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", config.DB_NAME)
	if err := db.Exec(createDBSQL).Error; err != nil {
		log.Println("failed to create database: %v", err)
		CloseDB(db)
	}
}

func CloseDB(db *gorm.DB) {

	sqlDB, err := db.DB() // Get the underlying sql.DB object
	if err != nil {
		log.Println("failed to get sql.DB from gorm.DB: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("failed to close the database connection: %v", err)
	}
}
